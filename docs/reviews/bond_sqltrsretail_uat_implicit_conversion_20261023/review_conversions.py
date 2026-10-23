"""Whole-snapshot implicit-conversion review, with table/column-first evidence.

Offline only. Reuses the pinned, bounded ScriptDom binder from the 2026-10-22
audit and extends observations to write boundaries and other SQL expressions.
No business SQL, migrations, or historical report artifacts are modified.
"""
import argparse
import collections
import csv
import hashlib
import json
import pathlib
import sys
import tempfile

HERE=pathlib.Path(__file__).resolve().parent
PREVIOUS=HERE.parent/'bond_sqltrsretail_uat_performance_20261022'
sys.path.insert(0,str(PREVIOUS))
from analyze import Audit, Doc, TEXT, NUM, DATE, PRECEDENCE, norm, compact, writecsv

FILTERS={'WHERE','JOIN_ON','DML_CONDITION','HAVING'}

class ReviewAudit(Audit):
    def __init__(self,docs):
        super().__init__(docs)
        self.signatures={};self.returntypes={};self.extra=[];self.skips=[];self.visits=collections.Counter()
        for d in docs:d.output_type_cache={};d.output_type_busy=set()
        # UPDATE/DELETE alias targets are not an additional FROM source. A target
        # alias may also happen to spell a real table name (UPDATE b FROM dbo.B b).
        for d in docs:
            for typ in ['UpdateSpecification','DeleteSpecification']:
                for spec in d.bytype[typ]:
                    target=d.child(spec,'Target')
                    if target is None or d.nodes[target]['t']!='NamedTableReference':continue
                    parts=d.identifiers(d.child(target,'SchemaObject'))
                    if len(parts)!=1:continue
                    peers=[r for r in d.refs[spec] if r!=target and d.nodes[r]['t']=='NamedTableReference']
                    if any(norm(d.value(r,'Alias') or (d.identifiers(d.child(r,'SchemaObject')) or [''])[-1])==norm(parts[0]) for r in peers):
                        d.refs[spec]=[r for r in d.refs[spec] if r!=target]
        for d in docs:
            if d.path.split('/')[0] not in {'functions','procedures'}:continue
            key=d.db+'.'+norm(pathlib.Path(d.path).stem)
            self.signatures[key]=[(d.value(p,'VariableName'),d.datatype(d.child(p,'DataType'))) for p in d.bytype['ProcedureParameter']]
            for rt in d.bytype['ScalarFunctionReturnType']:
                value=d.datatype(d.child(rt,'DataType'))
                if value:self.returntypes[key]=dict(value,origin='FUNCTION_RESULT',plain=False)
    def outputnames(self,d,q):
        if q is None:return []
        if d.nodes[q]['t']=='BinaryQueryExpression':return self.outputnames(d,d.child(q,'FirstQueryExpression'))
        if d.nodes[q]['t']=='QueryParenthesisExpression':return self.outputnames(d,d.child(q,'QueryExpression'))
        if d.nodes[q]['t']!='QuerySpecification':return []
        result=[]
        for sel in d.kids(q,'SelectElements'):
            if d.nodes[sel]['t']!='SelectScalarExpression':return []
            expr=d.child(sel,'Expression');name=d.value(sel,'ColumnName')
            if not name and expr is not None and d.nodes[expr]['t']=='ColumnReferenceExpression':
                parts=d.parts(expr);name=parts[-1] if parts else ''
            result.append(name)
        return result
    def querycols(self,d,i):
        if i is not None and d.nodes[i]['t']=='BinaryQueryExpression':
            names=self.outputnames(d,i);types=self.outputtypes(d,i)
            if names and len(names)==len(types):return {norm(n):v for n,v in zip(names,types) if n and v}
            return {}
        return super().querycols(d,i)
    def objectcols(self,key):
        obj=self.objects.get(key)
        if not obj or obj['kind'] not in {'VIEWS','FUNCTIONS'}:return super().objectcols(key)
        if 'cols' in obj:return obj['cols']
        if obj.get('busy'):return None
        obj['busy']=True;d=obj['doc'];result={}
        if obj['kind']=='VIEWS':
            nodes=d.bytype['CreateOrAlterViewStatement']+d.bytype['CreateViewStatement']+d.bytype['AlterViewStatement']
            if len(nodes)==1:
                sel=d.child(nodes[0],'SelectStatement');q=d.child(sel,'QueryExpression') if sel is not None else None
                result=self.querycols(d,q)
                explicit=[d.nodes[c].get('Value','') for c in d.kids(nodes[0],'Columns')]
                if explicit:
                    types=self.outputtypes(d,q)
                    result={norm(n):v for n,v in zip(explicit,types) if n and v} if len(explicit)==len(types) else {}
        else:
            for rt in d.bytype['TableValuedFunctionReturnType']:
                result=d.columns(rt,'FUNCTION_RETURN',key)
            for rt in d.bytype['SelectFunctionReturnType']:
                sel=d.child(rt,'SelectStatement')
                if sel is not None:result=self.querycols(d,d.child(sel,'QueryExpression'))
        obj['busy']=False;obj['cols']=result
        return result
    def ref(self,d,i):
        n=d.nodes[i]
        if n['t']=='NamedTableReference' and n['r']=='Target':
            parts=d.identifiers(d.child(i,'SchemaObject'));spec=n['p']
            if len(parts)==1:
                matches=[r for r in d.refs[spec] if r!=i and d.nodes[r]['t']=='NamedTableReference' and norm(d.value(r,'Alias') or (d.identifiers(d.child(r,'SchemaObject')) or [''])[-1])==norm(parts[0])]
                if len(matches)==1:
                    a,cs,key,state=self.ref(d,matches[0]);return a,cs,key,'DML_TARGET_ALIAS'
        if n['t']=='SchemaObjectFunctionTableReference':
            parts=d.identifiers(d.child(i,'SchemaObject'));key=self.key(parts,d.db)
            if key in self.objects and self.objects[key]['kind']=='FUNCTIONS':
                cols=self.objectcols(key)
                if cols:return d.value(i,'Alias') or parts[-1],cols,key,'FUNCTION_RETURN_TYPED'
        if n['t']=='InlineDerivedTable':
            names=[d.nodes[c].get('Value','') for c in d.kids(i,'Columns')];rows=d.kids(i,'RowValues');cols={}
            values=[d.kids(r,'ColumnValues') for r in rows]
            if names and values and all(len(v)==len(names) for v in values):
                for ordinal,name in enumerate(names):
                    types=[self.expr(d,row[ordinal],d.scope[row[ordinal]]) for row in values if d.nodes[row[ordinal]]['t']!='NullLiteral']
                    high=self.dominant(types)
                    if name and high:cols[norm(name)]=dict(high,origin='INLINE_VALUES',table='',column='',plain=False)
                return d.value(i,'Alias'),cols,'inline_values','INLINE_VALUES_TYPED'
        return super().ref(d,i)
    def synthetic(self,t,d,i):
        return {'type':t,'sql_type':t.upper(),'origin':'EXPRESSION','table':'','column':'','source':f'{d.path}:{d.nodes[i]["l"]}','collation':'','plain':False}
    def assignment_target(self,d,i,col):
        if col is None:return None
        spec=d.ancestor(i,{'UpdateSpecification'})
        if spec is None:return self.expr(d,col,d.scope[col])
        target=d.child(spec,'Target')
        if target is None:return None
        parts=d.parts(col)
        if not parts:return None
        alias,cols,key,state=self.ref(d,target)
        raw=d.identifiers(d.child(target,'SchemaObject')) if d.nodes[target]['t']=='NamedTableReference' else []
        if len(raw)==1:
            peers=[self.ref(d,r) for r in d.refs[spec] if r!=target]
            matches=[r for r in peers if norm(r[0])==norm(raw[0])]
            if len(matches)==1:alias,cols,key,state=matches[0]
        if len(parts)>1 and norm(parts[-2]) not in {norm(alias),norm(key.split('.')[-1])}:return None
        value=cols.get(norm(parts[-1])) if cols else None
        return dict(value,plain=True) if value else None
    def dominant(self,values):
        if not values or any(v is None or v['type'] not in PRECEDENCE for v in values):return None
        high=min(values,key=lambda v:PRECEDENCE.index(v['type']))
        return dict(high,sql_type=high['type'].upper()+' (combined result; facets unknown)')
    def outputexprs(self,d,q):
        if q is None:return []
        if d.nodes[q]['t']=='QueryParenthesisExpression':return self.outputexprs(d,d.child(q,'QueryExpression'))
        if d.nodes[q]['t']!='QuerySpecification':return []
        sels=d.kids(q,'SelectElements')
        if any(d.nodes[s]['t']!='SelectScalarExpression' for s in sels):return []
        return [d.child(s,'Expression') for s in sels]
    def outputtypes(self,d,q):
        if q is None:return []
        if q in d.output_type_cache:return d.output_type_cache[q]
        if q in d.output_type_busy:return []
        d.output_type_busy.add(q)
        values=self.compute_outputtypes(d,q)
        d.output_type_busy.remove(q)
        d.output_type_cache[q]=values
        return values
    def compute_outputtypes(self,d,q):
        if d.nodes[q]['t']=='BinaryQueryExpression':
            left=self.outputtypes(d,d.child(q,'FirstQueryExpression'));right=self.outputtypes(d,d.child(q,'SecondQueryExpression'))
            if not left or len(left)!=len(right):return []
            result=[]
            for a,b in zip(left,right):
                high=self.dominant([a,b]);result.append(dict(high,origin='EXPRESSION',table='',column='',plain=False) if high else None)
            return result
        return [self.expr(d,e,d.scope[e]) for e in self.outputexprs(d,q)]
    def expr(self,d,i,scope):
        if i is None:return None
        n=d.nodes[i];t=n['t']
        if t=='NullIfExpression':
            first=self.expr(d,d.child(i,'FirstExpression'),scope)
            return dict(first,origin='EXPRESSION',plain=False) if first else None
        if t=='ScalarSubquery':
            values=self.outputtypes(d,d.child(i,'QueryExpression'))
            return dict(values[0],plain=False) if len(values)==1 and values[0] else None
        if t in {'LeftFunctionCall','RightFunctionCall'}:
            args=d.kids(i,'Parameters');value=self.expr(d,args[0],scope) if args else None
            if value and value['type'] in TEXT|NUM|DATE:
                typ='nvarchar' if value['type'] in {'nvarchar','nchar','ntext'} else 'varchar'
                return dict(value,type=typ,sql_type=typ.upper()+' (string function; length unknown)',origin='EXPRESSION',plain=False)
            return None
        if t in {'SearchedCaseExpression','SimpleCaseExpression','CoalesceExpression'}:
            branches=d.kids(i,'Expressions') if t=='CoalesceExpression' else [d.child(w,'ThenExpression') for w in d.kids(i,'WhenClauses')]+([d.child(i,'ElseExpression')] if d.child(i,'ElseExpression') is not None else [])
            values=[self.expr(d,b,scope) for b in branches if d.nodes[b]['t']!='NullLiteral']
            high=self.dominant(values)
            return dict(high,origin='EXPRESSION',table='',column='',plain=False) if high else None
        if t=='FunctionCall':
            name=norm(d.value(i,'FunctionName'));args=d.kids(i,'Parameters')
            target=d.child(i,'CallTarget')
            if target is not None and d.nodes[target]['t']=='MultiPartIdentifierCallTarget':
                key=self.key(d.identifiers(d.child(target,'MultiPartIdentifier'))+[d.value(i,'FunctionName')],d.db)
                if key in self.returntypes:return dict(self.returntypes[key])
            if name in {'sum','avg'} and args:
                source=self.expr(d,args[0],scope)
                if source and source['type'] in NUM-{'bit'}:
                    typ='int' if source['type'] in {'tinyint','smallint','int'} else 'money' if source['type'] in {'money','smallmoney'} else 'float' if source['type'] in {'real','float'} else source['type']
                    return dict(source,type=typ,sql_type=typ.upper()+' (aggregate; facets unknown)',origin='EXPRESSION',plain=False)
            if name=='dateadd':
                if len(args)!=3:return None
                v=self.expr(d,args[2],scope)
                if d.nodes[args[2]]['t']=='StringLiteral':return self.synthetic('datetime',d,i)
                return dict(v,origin='EXPRESSION',table='',column='',plain=False) if v and v['type'] in DATE else None
            if name=='isnull' and args and d.nodes[args[0]]['t']=='NullLiteral':return self.expr(d,args[1],scope) if len(args)>1 else None
            if name=='replace' and args:
                values=[self.expr(d,arg,scope) for arg in args]
                if any(v is None for v in values):return None
                typ='nvarchar' if any(v['type'] in {'nvarchar','nchar','ntext'} for v in values) else 'varchar'
                return dict(values[0],type=typ,sql_type=typ.upper()+' (expression; length unknown)',plain=False,origin='EXPRESSION')
            if name in {'ltrim','rtrim','trim','upper','lower','left','right','substring','replicate'} and args:
                v=self.expr(d,args[0],scope)
                if not v:return None
                typ='nvarchar' if v['type'] in {'nvarchar','nchar','ntext'} else 'varchar'
                # Length/precision are intentionally not inferred for function results.
                return dict(v,type=typ,sql_type=typ.upper()+' (expression; length unknown)',plain=False,origin='EXPRESSION')
        if t=='BinaryExpression':
            a=self.expr(d,d.child(i,'FirstExpression'),scope);b=self.expr(d,d.child(i,'SecondExpression'),scope)
            if a and b and a['type'] in NUM|TEXT and b['type'] in NUM|TEXT:
                high=self.dominant([a,b])
                if high and (high['type'] in NUM or n.get('BinaryExpressionType')=='Add'):
                    origins={(v.get('table',''),v.get('column','')) for v in [a,b] if v.get('table')}
                    table,column=next(iter(origins)) if len(origins)==1 else ('','')
                    return dict(high,origin='EXPRESSION',table=table,column=column,plain=False)
            # Do not ask the base binder to recurse over the same unknown binary
            # subtree again: long dynamic-SQL concatenations grow exponentially.
            return None
        return super().expr(d,i,scope)
    def temp_select_into(self,d):
        d.output_type_cache.clear();d.querycache.clear()
        for i in d.bytype['SelectStatement']:
            parts=d.identifiers(d.child(i,'Into'))
            if parts and parts[-1].startswith('#'):
                cols=self.querycols(d,d.child(i,'QueryExpression'))
                if cols:d.temps[norm(parts[-1])]={k:dict(v,origin='TEMP_SELECT_INTO',table=parts[-1],column=k) for k,v in cols.items()}
                # A later SELECT INTO may depend on this newly available schema.
                d.output_type_cache.clear();d.querycache.clear()
    def emit(self,d,i,left,a,right,b,form,method='PRECEDENCE',converted=None,target=None,context=None):
        self.visits[form]+=1
        if not a or not b:
            self.skips.append({'database':d.db,'file':d.path,'line':d.nodes[i]['l'],'form':form,'reason':'Operand type, target, or supported result shape unresolved','code':compact(d.text(i))[:600]});return
        if a['type']==b['type']:return
        if a['type'] not in PRECEDENCE or b['type'] not in PRECEDENCE:return
        if not a.get('table') and not b.get('table'):return
        if converted is None:converted='left' if PRECEDENCE.index(a['type'])>PRECEDENCE.index(b['type']) else 'right'
        low=a if converted=='left' else b;high=b if converted=='left' else a
        dest=target or (high['sql_type'] if method in {'ASSIGNMENT','ISNULL_RETURN','UDF_PARAMETER'} else high['type'].upper())
        if method in {'ASSIGNMENT','ISNULL_RETURN','UDF_PARAMETER'}:cat='BOUNDARY_'+low['type'].upper()+'_TO_'+high['type'].upper()
        elif method=='LIKE_CHARACTER':cat='LIKE_NONCHAR_TO_CHARACTER'
        elif low['type'] in TEXT and high['type'] in NUM:cat='COLUMN_TEXT_TO_NUMERIC' if low.get('plain') else 'OPERAND_TEXT_TO_NUMERIC'
        elif low['type'] in {'varchar','char'} and high['type'] in {'nvarchar','nchar'}:cat='COLUMN_TO_UNICODE' if low.get('plain') else 'OPERAND_TO_UNICODE'
        elif low['type'] in TEXT and high['type'] in DATE:cat='COLUMN_TEXT_TO_DATE' if low.get('plain') else 'OPERAND_TEXT_TO_DATE'
        elif low['type'] in TEXT and high['type'] in TEXT:cat='TEXT_FAMILY_ONLY'
        elif low['type'] in NUM and high['type'] in NUM:cat='NUMERIC_FAMILY'
        else:cat='OTHER_TYPE_DIFFERENCE'
        self.extra.append({'database':d.db,'file':d.path,'line':d.nodes[i]['l'],'context':context or self.context(d,i),'form':form,'category':cat,'predicate':compact(d.text(i)) if form not in {'INSERT_SELECT','INSERT_VALUES'} else f'{left} <- {right}',
                           'left':left,'left_type':a['sql_type'],'left_origin':a['table']+'.'+a['column'] if a['table'] else a['origin'],'left_declaration':a['source'],
                           'right':right,'right_type':b['sql_type'],'right_origin':b['table']+'.'+b['column'] if b['table'] else b['origin'],'right_declaration':b['source'],
                           'converted_side':converted,'target_type':dest,'column_side_conversion':bool(low.get('plain')),'existing_leading_indexes':self.leading(low['table'],low['column']) if low['table'] else '',
                           'method':method,'confidence':'STATIC_TYPE_RULE; runtime plan/data not observed'})
    def supplemental(self):
        for d in self.docs:
            if d.path.split('/')[0] not in {'procedures','functions','views'}:continue
            self.temp_select_into(d);d.querycache.clear()
            for typ in ['BinaryExpression','LikePredicate','InPredicate','AssignmentSetClause','SelectSetVariable','SetVariableStatement','FunctionCall','SearchedCaseExpression','SimpleCaseExpression','CoalesceExpression','NullIfExpression','BinaryQueryExpression','InsertSpecification']:
                for i in d.bytype[typ]:
                    n=d.nodes[i];sc=d.scope[i]
                    def inferred(j):return self.expr(d,j,d.scope[j]) if j is not None else None
                    def txt(j):return compact(d.text(j)) if j is not None else ''
                    if typ=='BinaryExpression':
                        l=d.child(i,'FirstExpression');r=d.child(i,'SecondExpression')
                        a=inferred(l);b=inferred(r)
                        if a and b and (a['type'] in DATE or b['type'] in DATE):
                            self.skips.append({'database':d.db,'file':d.path,'line':n['l'],'form':'BINARY_TEMPORAL','reason':'Temporal arithmetic uses operator-specific rules; do not infer conversion from generic precedence','code':compact(d.text(i))[:600]})
                        else:self.emit(d,i,txt(l),a,txt(r),b,'BINARY_'+n.get('BinaryExpressionType',''))
                    elif typ=='LikePredicate':
                        l=d.child(i,'FirstExpression');r=d.child(i,'SecondExpression');a=inferred(l);b=inferred(r)
                        if a and b and a['type'] not in TEXT and b['type'] in TEXT:
                            self.emit(d,i,txt(l),a,txt(r),b,'LIKE','LIKE_CHARACTER',converted='left',target=b['sql_type'])
                        elif a and b and a['type'] in TEXT and b['type'] in TEXT and ({a['type'],b['type']}&{'nvarchar','nchar'}):self.emit(d,i,txt(l),a,txt(r),b,'LIKE')
                        elif not a or not b:self.emit(d,i,txt(l),a,txt(r),b,'LIKE')
                    elif typ=='InPredicate':
                        sub=d.child(i,'Subquery')
                        if sub is not None:
                            l=d.child(i,'Expression');self.emit(d,i,txt(l),inferred(l),txt(sub),inferred(sub),'IN_SUBQUERY')
                    elif typ in {'AssignmentSetClause','SelectSetVariable','SetVariableStatement'}:
                        if n.get('AssignmentKind','Equals')!='Equals':continue
                        l=d.child(i,'Column' if typ=='AssignmentSetClause' else 'Variable');r=d.child(i,'NewValue' if typ=='AssignmentSetClause' else 'Expression')
                        destination=self.assignment_target(d,i,l) if typ=='AssignmentSetClause' else inferred(l)
                        self.emit(d,i,txt(l),destination,txt(r),inferred(r),typ.upper(),'ASSIGNMENT',converted='right',context='WRITE_ASSIGNMENT' if typ=='AssignmentSetClause' else 'VARIABLE_ASSIGNMENT')
                    elif typ=='FunctionCall':
                        name=norm(d.value(i,'FunctionName'));args=d.kids(i,'Parameters')
                        if name=='isnull' and len(args)==2 and d.nodes[args[0]]['t']!='NullLiteral':
                            l,r=args;self.emit(d,i,txt(l),inferred(l),txt(r),inferred(r),'ISNULL_REPLACEMENT','ISNULL_RETURN',converted='right',context='EXPRESSION')
                        target=d.child(i,'CallTarget')
                        if target is not None and d.nodes[target]['t']=='MultiPartIdentifierCallTarget':
                            parts=d.identifiers(d.child(target,'MultiPartIdentifier'))+[d.value(i,'FunctionName')];key=self.key(parts,d.db)
                            signature=self.signatures.get(key)
                            if signature and len(args)==len(signature):
                                for arg,(pn,dt) in zip(args,signature):self.emit(d,i,key+'.'+pn,dt,txt(arg),inferred(arg),'UDF_ARGUMENT','UDF_PARAMETER',converted='right',context='FUNCTION_CALL')
                    elif typ in {'SearchedCaseExpression','SimpleCaseExpression','CoalesceExpression'}:
                        if typ=='SimpleCaseExpression':
                            input=d.child(i,'InputExpression')
                            for when in d.kids(i,'WhenClauses'):
                                value=d.child(when,'WhenExpression')
                                self.emit(d,when,txt(input),inferred(input),txt(value),inferred(value),'SIMPLE_CASE_MATCH',context='CASE_TEST')
                        branches=d.kids(i,'Expressions') if typ=='CoalesceExpression' else [d.child(w,'ThenExpression') for w in d.kids(i,'WhenClauses')]+([d.child(i,'ElseExpression')] if d.child(i,'ElseExpression') is not None else [])
                        values=[inferred(b) for b in branches if d.nodes[b]['t']!='NullLiteral'];high=self.dominant(values)
                        for branch in branches:
                            v=inferred(branch)
                            if v and high:self.emit(d,i,'result type',dict(high,table='',column='',origin='RESULT_TYPE'),txt(branch),v,'COALESCE_BRANCH' if typ=='CoalesceExpression' else 'CASE_BRANCH',converted='right',context='EXPRESSION')
                    elif typ=='NullIfExpression':
                        left=d.child(i,'FirstExpression');right=d.child(i,'SecondExpression')
                        self.emit(d,i,txt(left),inferred(left),txt(right),inferred(right),'NULLIF_COMPARE',context='EXPRESSION')
                    elif typ=='BinaryQueryExpression':
                        l=d.child(i,'FirstQueryExpression');r=d.child(i,'SecondQueryExpression');a=self.outputtypes(d,l);b=self.outputtypes(d,r)
                        if a and len(a)==len(b):
                            for ordinal,(av,bv) in enumerate(zip(a,b),1):self.emit(d,i,f'left output #{ordinal}',av,f'right output #{ordinal}',bv,n.get('BinaryQueryExpressionType','SET').upper()+'_OUTPUT',context='SET_OUTPUT')
                        else:self.skips.append({'database':d.db,'file':d.path,'line':n['l'],'form':'SET_OUTPUT','reason':'Result ordinal/type not fully resolved; no guessed UNION mapping','code':compact(d.text(i))[:600]})
                    elif typ=='InsertSpecification':
                        target=d.child(i,'Target');columns=d.kids(i,'Columns');src=d.child(i,'InsertSource')
                        if target is None or src is None:continue
                        alias,cs,key,state=self.ref(d,target)
                        mappings=[]
                        if d.nodes[src]['t']=='SelectInsertSource':
                            exprs=self.outputexprs(d,d.child(src,'Select'))
                            if exprs:mappings.append(('INSERT_SELECT',exprs))
                        elif d.nodes[src]['t']=='ValuesInsertSource':
                            mappings=[('INSERT_VALUES',d.kids(row,'ColumnValues')) for row in d.kids(src,'RowValues')]
                        if not cs or not columns or not mappings or any(len(es)!=len(columns) for form,es in mappings):
                            self.skips.append({'database':d.db,'file':d.path,'line':n['l'],'form':'INSERT','reason':'Missing explicit column list, unresolved target, INSERT EXEC, star/UNION/unsupported source; no guessed ordinal mapping','code':compact(d.text(i))[:600]});continue
                        for form,exprs in mappings:
                            for col,expr in zip(columns,exprs):
                                parts=d.parts(col);cv=cs.get(norm(parts[-1])) if parts else None
                                self.emit(d,i,key+'.'+(parts[-1] if parts else '?'),cv,txt(expr),inferred(expr),form,'ASSIGNMENT',converted='right',context='WRITE_ASSIGNMENT')

def load_csv(p):
    with p.open(encoding='utf-8-sig',newline='') as f:return list(csv.DictReader(f))

def main():
    ap=argparse.ArgumentParser();ap.add_argument('--workspace',required=True,type=pathlib.Path);ap.add_argument('--retail-ast',required=True,type=pathlib.Path);ap.add_argument('--replicate-ast',required=True,type=pathlib.Path);args=ap.parse_args();workspace=args.workspace.resolve()
    docs=[];manifest=[]
    for ast,db,rel in [(args.retail_ast,'retail','bond-sqltrsretail-uat/db_object'),(args.replicate_ast,'sql_replicate','all_object/bond/bond-sqlreplicate/db_object')]:
        for line in ast.open(encoding='utf8'):
            d=Doc(json.loads(line),workspace/rel,db);docs.append(d);manifest.append({'path':d.full.relative_to(workspace).as_posix(),'sha256':hashlib.sha256(d.full.read_bytes()).hexdigest()})
    audit=ReviewAudit(docs)
    # Preserve the previous dated report. Fresh base outputs exist only in temp.
    with tempfile.TemporaryDirectory(prefix='rza-conversion-base-') as tmp:
        tmp=pathlib.Path(tmp);audit.run(tmp)
        base=load_csv(tmp/'implicit_conversions.csv');unresolved=load_csv(tmp/'unresolved_predicates.csv');coverage=load_csv(tmp/'module_coverage.csv');references=load_csv(tmp/'object_references.csv');columns=load_csv(tmp/'table_columns.csv');indexes=load_csv(tmp/'existing_indexes.csv');oldsummary=json.loads((tmp/'summary.json').read_text())
    audit.supplemental()
    sites=[]
    for r in base:
        r.update(method='PRECEDENCE')
        # Comparison target family is defensible; exact computed length/precision is not.
        import re
        match=re.match(r'\[?([A-Za-z0-9_]+)',r['target_type'])
        if match:r['target_type']=match.group(1).upper()
    # Only procedure query occurrences are central; functions/views remain a supporting appendix.
    for i,r in enumerate(sorted(base+audit.extra,key=lambda r:(r['database'],r['file'],int(r['line']),r['form'],r['left'],r['right'])),1):
        r['site_id']=f'IC-{i:06}'
        r['module_kind']=r['file'].split('/')[0]
        r['source_path']=('bond-sqltrsretail-uat/db_object/' if r['database']=='retail' else 'all_object/bond/bond-sqlreplicate/db_object/')+r['file']
        sites.append(r)
    # Physical column identities, not column-name or alias guesses.
    identities={norm(r['table']+'.'+r['column']):r for r in columns}
    endpoints=[]
    for r in sites:
        for side in ['left','right']:
            identity=identities.get(norm(r[side+'_origin']))
            if not identity:continue
            converted=side==r['converted_side'];method=r['method'];context=r['context']
            if method=='ASSIGNMENT':role='SOURCE_TO_DESTINATION' if side=='right' else 'DESTINATION_COLUMN'
            elif method in {'UDF_PARAMETER','ISNULL_RETURN'}:role='CONVERTED_ARGUMENT' if converted else 'TYPE_ANCHOR'
            else:role='CONVERTED_COLUMN' if converted and str(r['column_side_conversion']).lower()=='true' else 'CONVERTED_COLUMN_EXPRESSION' if converted else 'COUNTERPART_COLUMN_NOT_CONVERTED'
            highrisk=r['category'] in {'COLUMN_TEXT_TO_NUMERIC','COLUMN_TEXT_TO_DATE','COLUMN_TO_UNICODE','LIKE_NONCHAR_TO_CHARACTER'}
            indexrisk=converted and str(r['column_side_conversion']).lower()=='true' and context in FILTERS and method in {'PRECEDENCE','LIKE_CHARACTER'} and highrisk
            if indexrisk:tier='A_COLUMN_PREDICATE'
            elif converted and r['category'] not in {'NUMERIC_FAMILY','TEXT_FAMILY_ONLY'}:tier='B_EXPRESSION_OR_BOUNDARY'
            elif not converted:tier='PARTNER_OR_DESTINATION'
            else:tier='C_FAMILY_OR_ROUTINE'
            endpoints.append({'site_id':r['site_id'],'table':identity['table'],'column':identity['column'],'column_type':identity['type'],'role':role,'tier':tier,'form':r['form'],'context':context,'pattern':r['category'],'method':method,'module_database':r['database'],'file':r['file'],'source_path':r['source_path'],'line':r['line'],'predicate':r['predicate'],'column_operand':r[side],'operand_type':r[side+'_type'],'other_operand':r['right' if side=='left' else 'left'],'other_type':r['right_type' if side=='left' else 'left_type'],'converted_operand':r[r['converted_side']],'target_type':r['target_type'],'column_declaration':identity['source'],'existing_leading_indexes':audit.leading(identity['table'],identity['column'])})
    writecsv(HERE/'conversion_sites.csv',sites);writecsv(HERE/'column_occurrences.csv',endpoints)
    writecsv(HERE/'source_manifest.csv',manifest);writecsv(HERE/'table_columns.csv',columns);writecsv(HERE/'existing_indexes.csv',indexes)
    writecsv(HERE/'module_coverage.csv',coverage);writecsv(HERE/'object_references.csv',references)
    writecsv(HERE/'unresolved_predicates.csv',unresolved);writecsv(HERE/'unresolved_extensions.csv',audit.skips)
    physical=[r for r in endpoints if r['module_database']=='retail' and r['file'].startswith('procedures/')]
    dangerous=[r for r in physical if r['tier']=='A_COLUMN_PREDICATE']
    summary={'date':'2026-10-23','input_files':len(manifest),'retail_procedures':sum(d.db=='retail' and d.path.startswith('procedures/') for d in docs),'tables':sum(d.path.startswith('tables/') for d in docs),'base_comparison_binding':oldsummary['predicate_binding'],'all_sites':len(sites),'base_sites':len(base),'supplementary_sites':len(audit.extra),'sites_by_form':dict(collections.Counter(r['form'] for r in sites)),'physical_column_endpoints':len(endpoints),'retail_procedure_physical_endpoints':len(physical),'tables_in_procedure_conversion_context':len({r['table'] for r in physical}),'columns_in_procedure_conversion_context':len({(r['table'],r['column']) for r in physical}),'procedure_column_predicate_risk_sites':len({r['site_id'] for r in dangerous}),'procedure_column_predicate_risk_tables':len({r['table'] for r in dangerous}),'procedure_column_predicate_risk_columns':len({(r['table'],r['column']) for r in dangerous}),'procedure_column_predicate_risk_procedures':len({r['file'] for r in dangerous}),'supplemental_visits':dict(audit.visits),'unresolved_extension_observations':len(audit.skips),'parse_errors':oldsummary['parse_errors'],'limits':['Static type differences and rule-based direction only; no actual plans/data observed.','Equal base types with differing lengths/precision/scale are not inferred as separate conversions.','Unsupported/ambiguous mappings stay in unresolved files; all files scanned is not all SQL paths fully bound.','Inherited snapshot binder is casefold-based and does not model live collation or branch-specific temp schemas.']}
    (HERE/'summary.json').write_text(json.dumps(summary,indent=2),encoding='utf8')
    print(json.dumps(summary,indent=2))

if __name__=='__main__':main()
