"""Bounded offline AST audit; produces evidence, never changes SQL source.

Usage: python analyze.py --workspace ROOT --ast TEMP/ast.jsonl --replicate-ast TEMP/replicate_ast.jsonl
ScriptDom extraction is separate. No claim of full SQL Server name/type binding.
"""
import argparse
import collections
import csv
import hashlib
import json
import pathlib
import re

SCOPES = {'QuerySpecification', 'UpdateSpecification', 'DeleteSpecification', 'MergeSpecification', 'InsertSpecification'}
TEXT = {'char', 'varchar', 'nchar', 'nvarchar', 'text', 'ntext'}
NUM = {'bigint', 'int', 'smallint', 'tinyint', 'bit', 'decimal', 'numeric', 'money', 'smallmoney', 'float', 'real'}
DATE = {'datetimeoffset', 'datetime2', 'datetime', 'smalldatetime', 'date', 'time'}
PRECEDENCE = ['sql_variant','xml','datetimeoffset','datetime2','datetime','smalldatetime','date','time','float','real','decimal','numeric','money','smallmoney','bigint','int','smallint','tinyint','bit','ntext','text','image','timestamp','uniqueidentifier','nvarchar','nchar','varchar','char','varbinary','binary']

def norm(s): return s.casefold()
def compact(s): return re.sub(r'\s+', ' ', s).strip()
def writecsv(path, rows, columns=None):
    if columns is None: columns = list(dict.fromkeys(k for r in rows for k in r))
    with path.open('w', encoding='utf-8-sig', newline='') as f:
        w=csv.DictWriter(f,fieldnames=columns); w.writeheader(); w.writerows(rows)

class Doc:
    def __init__(self, raw, root, db):
        self.path=raw['path']; self.db=db; self.nodes=raw['nodes']; self.errors=raw['errors']
        self.full=root/self.path
        if raw.get('sha256') != hashlib.sha256(self.full.read_bytes()).hexdigest():
            raise ValueError('AST source hash missing/stale: '+str(self.full))
        self.sql=self.full.read_bytes().decode('utf-8-sig')
        # ScriptDom offsets count UTF-16 code units. Map correctly even with non-BMP source.
        self.utf16=self.sql.encode('utf-16-le')
        self.children=collections.defaultdict(list); self.bytype=collections.defaultdict(list)
        self.scope={}; self.refs=collections.defaultdict(list); self.temps={}; self.variables={}; self.ctes=collections.defaultdict(list)
        for i,n in enumerate(self.nodes):
            self.children[n['p']].append(i); self.bytype[n['t']].append(i)
            self.scope[i]=i if n['t'] in SCOPES else self.scope.get(n['p'],-1)
            if n['t'] in {'NamedTableReference','QueryDerivedTable','VariableTableReference','SchemaObjectFunctionTableReference','OpenRowsetTableReference','OpenXmlTableReference','InlineDerivedTable'}:
                self.refs[self.scope[i]].append(i)
        for i in self.bytype['CommonTableExpression']:
            self.ctes[norm(self.value(i,'ExpressionName'))].append(i)
        self.querycache={}; self.busy=set()
    def child(self,i,role):
        return next((c for c in self.children[i] if self.nodes[c]['r']==role),None)
    def kids(self,i,role): return [c for c in self.children[i] if self.nodes[c]['r']==role]
    def value(self,i,role):
        j=self.child(i,role); return self.nodes[j].get('Value',self.nodes[j].get('Name','')) if j is not None else ''
    def text(self,i):
        n=self.nodes[i]; return self.utf16[n['o']*2:(n['o']+n['z'])*2].decode('utf-16-le')
    def descendants(self,i,types=None):
        for c in self.children[i]:
            if types is None or self.nodes[c]['t'] in types: yield c
            yield from self.descendants(c,types)
    def identifiers(self,i):
        if i is None:return []
        return [self.nodes[c]['Value'] for c in self.kids(i,'Identifiers')]
    def parts(self,i): return self.identifiers(self.child(i,'MultiPartIdentifier'))
    def ancestor(self,i,types):
        p=self.nodes[i]['p']
        while p>=0:
            if self.nodes[p]['t'] in types:return p
            p=self.nodes[p]['p']
        return None
    def datatype(self,i):
        if i is None:return None
        n=self.nodes[i]
        if n['t']!='SqlDataTypeReference': return None
        t=n['SqlDataTypeOption'].lower()
        return {'type':t,'sql_type':compact(self.text(i)), 'origin':'DECLARATION', 'table':'','column':'','source':f'{self.path}:{n["l"]}', 'collation':''}
    def decls(self):
        for t in ['ProcedureParameter','DeclareVariableElement']:
            for i in self.bytype[t]:
                v=norm(self.value(i,'VariableName')); dt=self.datatype(self.child(i,'DataType'))
                if dt:self.variables[v]=dict(dt,origin='VARIABLE',source=f'{self.path}:{self.nodes[i]["l"]}')
        for i in self.bytype['CreateTableStatement']:
            parts=self.identifiers(self.child(i,'SchemaObjectName'))
            if parts and parts[-1].startswith('#'):self.temps[norm(parts[-1])]=self.columns(i,'TEMP_TABLE',parts[-1])
        for i in self.bytype['DeclareTableVariableBody']:
            name=self.value(i,'VariableName')
            if name:self.temps[norm(name)]=self.columns(i,'TABLE_VARIABLE',name)
    def columns(self,i,origin,table):
        result={}
        for c in self.descendants(i,{'ColumnDefinition'}):
            name=self.value(c,'ColumnIdentifier'); dt=self.datatype(self.child(c,'DataType'))
            if dt:result[norm(name)]=dict(dt,origin=origin,table=table,column=name,source=f'{self.path}:{self.nodes[c]["l"]}',collation=self.value(c,'Collation'))
        return result

class Audit:
    def __init__(self,docs):
        self.docs=docs; self.objects={}; self.indexes=[]; self.unresolved=collections.Counter(); self.predstatus=collections.Counter()
        for d in docs:
            d.decls()
            kind=d.path.split('/')[0]
            name=norm(pathlib.Path(d.path).stem)
            key=d.db+'.'+name
            if kind=='tables':
                ts=d.bytype['CreateTableStatement']
                self.objects[key]={'kind':'TABLE','cols':d.columns(ts[0],'TABLE',key) if ts else {},'doc':d}
            elif kind in {'views','functions'}:self.objects[key]={'kind':kind.upper(),'doc':d}
        self.build_indexes()
    def key(self,parts,db):
        if len(parts)==1:return db+'.dbo.'+norm(parts[0])
        if len(parts)==2:return db+'.'+'.'.join(map(norm,parts))
        if len(parts)==3:return '.'.join(map(norm,[parts[0],parts[1] or 'dbo',parts[2]]))
        return '.'.join(map(norm,parts))
    def build_indexes(self):
        for d in self.docs:
            if d.path.split('/')[0] not in {'indexes','tables'}:continue
            for typ in ['CreateIndexStatement','UniqueConstraintDefinition']:
                for i in d.bytype[typ]:
                    if typ=='CreateIndexStatement':
                        table=self.key(d.identifiers(d.child(i,'OnName')),d.db); name=d.value(i,'Name'); cls=d.nodes[i].get('Clustered',''); unique=d.nodes[i].get('Unique','')
                    else:
                        table=d.db+'.'+norm(pathlib.Path(d.path).stem);name=d.value(i,'ConstraintIdentifier');cls=d.nodes[i].get('Clustered','');unique='True'
                    keys=[]
                    for c in d.kids(i,'Columns'):
                        col=d.child(c,'Column');parts=d.parts(col) if col is not None else []
                        if parts:keys.append(parts[-1])
                    if not keys:
                        par=d.nodes[i]['p']
                        if d.nodes[par]['t']=='ColumnDefinition':keys=[d.value(par,'ColumnIdentifier')]
                    includes=[d.parts(c)[-1] for c in d.kids(i,'IncludeColumns') if d.parts(c)]
                    filt=d.child(i,'FilterPredicate')
                    self.indexes.append({'table':table,'name':name,'keys':keys,'includes':includes,'clustered':cls,'unique':unique,'filter':d.text(filt) if filt is not None else '', 'source':d.path,'line':d.nodes[i]['l']})
    def objectcols(self,key):
        obj=self.objects.get(key)
        if not obj:return None
        if 'cols' in obj:return obj['cols']
        if obj.get('busy'):return None
        obj['busy']=True;d=obj['doc'];queries=[]
        for i in d.bytype['QuerySpecification']:
            if d.ancestor(i,{'ScalarSubquery','QueryDerivedTable','CommonTableExpression'}) is None:queries.append(i)
        # UNION branches/ambiguous top-level shapes remain unresolved.
        obj['cols']=self.qcols(d,queries[0]) if len(queries)==1 else {}
        obj['busy']=False
        return obj['cols']
    def ref(self,d,i):
        n=d.nodes[i];typ=n['t'];alias=d.value(i,'Alias')
        if typ=='NamedTableReference':
            parts=d.identifiers(d.child(i,'SchemaObject'));base=parts[-1] if parts else '';alias=alias or base
            if norm(base) in d.temps:return alias,d.temps[norm(base)],base,'TEMP'
            if len(parts)==1 and norm(base) in d.ctes:
                cs=d.ctes[norm(base)]
                # Same CTE name in separate statements must not be conflated.
                stmt=d.ancestor(i,{'SelectStatement','UpdateStatement','DeleteStatement','InsertStatement','MergeStatement'})
                cs=[c for c in cs if d.ancestor(c,{'SelectStatement','UpdateStatement','DeleteStatement','InsertStatement','MergeStatement'})==stmt]
                if len(cs)==1:
                    q=d.child(cs[0],'QueryExpression');cols=self.querycols(d,q)
                    explicit=[d.nodes[x].get('Value','') for x in d.kids(cs[0],'Columns')]
                    if explicit and len(explicit)==len(cols):cols=dict(zip(map(norm,explicit),cols.values()))
                    return alias,cols,base,'CTE'
                return alias,None,base,'AMBIGUOUS_CTE'
            key=self.key(parts,d.db);cols=self.objectcols(key)
            return alias,cols,key,'RESOLVED_STATIC' if cols is not None else 'NOT_FOUND_IN_SCOPE'
        if typ=='QueryDerivedTable':return alias,self.querycols(d,d.child(i,'QueryExpression')),alias,'DERIVED'
        if typ=='VariableTableReference':
            v=d.child(i,'Variable');name=d.nodes[v].get('Name','') if v is not None else ''
            return alias or name,d.temps.get(norm(name)),name,'TABLE_VARIABLE'
        return alias,None,d.text(i),'UNSUPPORTED_ROWSET'
    def scopechain(self,d,scope):
        while scope is not None and scope>=0:
            yield scope
            p=d.nodes[scope]['p'];scope=d.scope.get(p,-1)
    def resolvecol(self,d,i,scope):
        parts=d.parts(i)
        if not parts:return None
        col=norm(parts[-1]);qual=norm(parts[-2]) if len(parts)>1 else None
        for sc in self.scopechain(d,scope):
            refs=[self.ref(d,r) for r in d.refs[sc]]
            # UPDATE alias target is a reference to an existing FROM alias.
            aliases={norm(a) for a,cs,k,s in refs if cs is not None}
            refs=[r for r in refs if not(r[1] is None and norm(r[0]) in aliases)]
            if qual:
                matches=[r for r in refs if norm(r[0])==qual or norm(r[2].split('.')[-1])==qual]
                if matches:
                    if len(matches)==1 and matches[0][1] is not None:return matches[0][1].get(col)
                    return None
            else:
                known=[cs[col] for a,cs,k,s in refs if cs is not None and col in cs]
                if any(cs is None for a,cs,k,s in refs):return None
                if len(known)==1:return known[0]
                if len(known)>1:return None
        return None
    def expr(self,d,i,scope):
        if i is None:return None
        n=d.nodes[i];t=n['t']
        if t=='ColumnReferenceExpression':
            v=self.resolvecol(d,i,scope);return dict(v,plain=True) if v else None
        if t=='VariableReference':return d.variables.get(norm(n.get('Name','')))
        if t in {'StringLiteral','IntegerLiteral','NumericLiteral','RealLiteral','MoneyLiteral'}:
            typ={'IntegerLiteral':'int','NumericLiteral':'decimal','RealLiteral':'float','MoneyLiteral':'money'}.get(t,'nvarchar' if n.get('IsNational')=='True' else 'varchar')
            if t=='IntegerLiteral' and int(n.get('Value','0'))>2147483647:typ='decimal'
            return {'type':typ,'sql_type':typ,'origin':'LITERAL','table':'','column':'','source':f'{d.path}:{n["l"]}','collation':''}
        if t in {'CastCall','ConvertCall','TryCastCall','TryConvertCall'}:
            dt=d.datatype(d.child(i,'DataType'))
            return dict(dt,origin='EXPRESSION') if dt else None
        if t in {'ParenthesisExpression','UnaryExpression'}:return self.expr(d,d.child(i,'Expression'),scope)
        if t in {'FunctionCall','LeftFunctionCall','RightFunctionCall'}:
            name=norm(d.value(i,'FunctionName') or t.replace('FunctionCall',''))
            args=d.kids(i,'Parameters')
            if name in {'isnull','ltrim','rtrim','trim','upper','lower','replace','replicate','left','right','substring','min','max'} and args:
                v=self.expr(d,args[0],scope)
                return dict(v,plain=False,origin='EXPRESSION') if v else None
            if name in {'getdate','getutcdate','dateadd'}:typ='datetime'
            elif name in {'sysdatetime','sysutcdatetime'}:typ='datetime2'
            elif name in {'datediff','datepart','day','month','year','len','count','charindex'}:typ='int'
            elif name in {'count_big','row_number','rank','dense_rank'}:typ='bigint'
            else:return None
            return {'type':typ,'sql_type':typ,'origin':'EXPRESSION','table':'','column':'','source':f'{d.path}:{n["l"]}','collation':''}
        if t=='BinaryExpression':
            a=self.expr(d,d.child(i,'FirstExpression'),scope);b=self.expr(d,d.child(i,'SecondExpression'),scope)
            if a and b and a['type'] in TEXT and b['type'] in TEXT and n.get('BinaryExpressionType')=='Add':
                v=a if PRECEDENCE.index(a['type'])<PRECEDENCE.index(b['type']) else b
                return dict(v,plain=False,origin='EXPRESSION',table='',column='')
        return None
    def querycols(self,d,i):
        if i is None:return {}
        if d.nodes[i]['t']=='QuerySpecification':return self.qcols(d,i)
        if d.nodes[i]['t']=='QueryParenthesisExpression':return self.querycols(d,d.child(i,'QueryExpression'))
        return {}
    def qcols(self,d,q):
        if q in d.querycache:return d.querycache[q]
        if q in d.busy:return {}
        d.busy.add(q);result={}
        for sel in d.kids(q,'SelectElements'):
            if d.nodes[sel]['t']=='SelectScalarExpression':
                expr=d.child(sel,'Expression');name=d.value(sel,'ColumnName');v=self.expr(d,expr,q)
                if not name and expr is not None and d.nodes[expr]['t']=='ColumnReferenceExpression':
                    parts=d.parts(expr);name=parts[-1] if parts else ''
                if name and v:result[norm(name)]=v
            elif d.nodes[sel]['t']=='SelectStarExpression':
                quals=d.identifiers(d.child(sel,'Qualifier'))
                for r in d.refs[q]:
                    a,cs,k,s=self.ref(d,r)
                    if cs and (not quals or norm(a)==norm(quals[-1])):result.update(cs)
        d.busy.remove(q);d.querycache[q]=result
        return result
    def temp_select_into(self,d):
        for i in d.bytype['SelectStatement']:
            it=d.child(i,'Into');parts=d.identifiers(it)
            if parts and parts[-1].startswith('#'):
                cols=self.querycols(d,d.child(i,'QueryExpression'))
                if cols:d.temps[norm(parts[-1])]={k:dict(v,origin='TEMP_SELECT_INTO',table=parts[-1],column=k) for k,v in cols.items()}
    def leading(self,table,col):
        return '|'.join(x['name'] for x in self.indexes if x['table']==table and x['keys'] and norm(x['keys'][0])==norm(col))
    def context(self,d,i):
        scope=d.scope[i];p=d.nodes[i]['p']
        while p>=0 and p!=scope:
            t=d.nodes[p]['t']
            if t=='WhereClause':return 'WHERE'
            if t=='QualifiedJoin':return 'JOIN_ON'
            if t=='HavingClause':return 'HAVING'
            if t in {'SearchedWhenClause','SimpleWhenClause'}:return 'CASE'
            p=d.nodes[p]['p']
        if scope>=0 and d.nodes[scope]['t'] in {'UpdateSpecification','DeleteSpecification','MergeSpecification'}:return 'DML_CONDITION'
        return 'EXPRESSION_OR_CONTROL'
    def run(self,out):
        conversions=[];unbound=[];patterns=[];refs=[];cf=[];access=[];inventory=[]
        for key in list(self.objects):self.objectcols(key)
        for d in self.docs:
            if d.path.split('/')[0] not in {'procedures','views','functions'}:continue
            self.temp_select_into(d);d.querycache.clear()
            refs_in_doc=[]
            for typ in ['NamedTableReference','SchemaObjectFunctionTableReference']:
                for i in d.bytype[typ]:
                    alias,cols,key,status=self.ref(d,i)
                    if status=='NOT_FOUND_IN_SCOPE':
                        peers=[self.ref(d,j) for j in d.refs[d.scope[i]] if j!=i]
                        if any(norm(a)==norm(alias) and cs is not None for a,cs,k,s in peers):status='DML_TARGET_ALIAS'
                    row={'database':d.db,'file':d.path,'line':d.nodes[i]['l'],'object':key,'alias':alias,'resolution':status};refs.append(row);refs_in_doc.append(row)
            pairs=[]
            for i in d.bytype['BooleanComparisonExpression']:
                pairs.append((i,d.child(i,'FirstExpression'),d.child(i,'SecondExpression'),'COMPARISON'))
            for i in d.bytype['InPredicate']:
                left=d.child(i,'Expression');seen=set()
                for right in d.kids(i,'Values'):
                    v=self.expr(d,right,d.scope[i]);sig=v['type'] if v else 'unknown'
                    if sig not in seen:pairs.append((i,left,right,'IN_VALUE_TYPE'));seen.add(sig)
            for i in d.bytype['BooleanTernaryExpression']:
                left=d.child(i,'FirstExpression');seen=set()
                for role in ['SecondExpression','ThirdExpression']:
                    right=d.child(i,role);v=self.expr(d,right,d.scope[i]);sig=v['type'] if v else 'unknown'
                    if sig not in seen:pairs.append((i,left,right,'BETWEEN_BOUND_TYPE'));seen.add(sig)
            for i,left,right,form in pairs:
                n=d.nodes[i];sc=d.scope[i]
                context=self.context(d,i)
                a=self.expr(d,left,sc);b=self.expr(d,right,sc)
                cols=[c for c in d.descendants(i,{'ColumnReferenceExpression'})]
                if not cols:continue
                if not a or not b:
                    self.predstatus['UNRESOLVED_OR_COMPLEX']+=1
                    unbound.append({'database':d.db,'file':d.path,'line':n['l'],'predicate':compact(d.text(i)),'left_type':a['sql_type'] if a else '', 'right_type':b['sql_type'] if b else '', 'reason':'Complex expression, unknown column, ambiguous scope, unsupported rowset, or unknown external metadata'})
                else:
                    self.predstatus['TYPED']+=1
                    if a['type']!=b['type'] and a['type'] in PRECEDENCE and b['type'] in PRECEDENCE:
                        low=a if PRECEDENCE.index(a['type'])>PRECEDENCE.index(b['type']) else b
                        high=b if low is a else a
                        oncol=bool(low.get('plain'))
                        if low['type'] in TEXT and high['type'] in NUM:
                            category='COLUMN_TEXT_TO_NUMERIC' if oncol else 'OPERAND_TEXT_TO_NUMERIC';priority='P1' if oncol else 'P3'
                        elif low['type'] in {'varchar','char'} and high['type'] in {'nvarchar','nchar'}:
                            category='COLUMN_TO_UNICODE' if oncol else 'OPERAND_TO_UNICODE';priority='P2' if oncol else 'P3'
                        elif low['type'] in TEXT and high['type'] in DATE:
                            category='COLUMN_TEXT_TO_DATE' if oncol else 'OPERAND_TEXT_TO_DATE';priority='P1' if oncol else 'P3'
                        elif low['type'] in TEXT and high['type'] in TEXT:category='TEXT_FAMILY_ONLY';priority='INFO'
                        elif low['type'] in NUM and high['type'] in NUM:category='NUMERIC_FAMILY';priority='P3'
                        else:category='OTHER_TYPE_DIFFERENCE';priority='P3'
                        conversions.append({'database':d.db,'file':d.path,'line':n['l'],'context':context,'form':form,'priority':priority,'category':category,'predicate':compact(d.text(i)), 'left':compact(d.text(left)),'left_type':a['sql_type'],'left_origin':a['table']+'.'+a['column'] if a['table'] else a['origin'],'left_declaration':a['source'],'right':compact(d.text(right)),'right_type':b['sql_type'],'right_origin':b['table']+'.'+b['column'] if b['table'] else b['origin'],'right_declaration':b['source'],'converted_side':'left' if low is a else 'right','target_type':high['sql_type'],'column_side_conversion':oncol,'existing_leading_indexes':self.leading(low['table'],low['column']) if low['table'] else '', 'confidence':'STATIC_TYPE_BINDING; optimizer effect requires plan'})
                # Index workload: direct, resolved physical column equality/range only.
                for side,other in [(left,right),(right,left)]:
                    if side is None or d.nodes[side]['t']!='ColumnReferenceExpression' or form!='COMPARISON' or context not in {'WHERE','JOIN_ON','DML_CONDITION'}:continue
                    v=self.expr(d,side,sc)
                    if v and v['origin']=='TABLE':
                        access.append({'database':d.db,'file':d.path,'line':n['l'],'query_line':d.nodes[sc]['l'] if sc>=0 else n['l'],'context':context,'table':v['table'],'column':v['column'],'operator':n.get('ComparisonType',''),'other':compact(d.text(other)),'existing_leading_indexes':self.leading(v['table'],v['column']),'predicate':compact(d.text(i))})
                # CFMAST predicates: include all raw CFMAST column references, not alias-name guesses.
                for c in cols:
                    v=self.expr(d,c,sc)
                    if v and v['table']=='sql_replicate.dbo.cfmast':
                        cf.append({'file':d.path,'line':n['l'],'predicate':compact(d.text(i)),'column':v['column'],'type':v['sql_type'],'column_source':v['source'],'other_type': b['sql_type'] if c==left and b else a['sql_type'] if c==right and a else 'COMPLEX/UNRESOLVED','index':self.leading(v['table'],v['column'])})
            # AST pattern inventory: review candidates, not automatically proven slow queries.
            for i,n in enumerate(d.nodes):
                t=n['t'];code=None;advice=None
                if t=='DeclareCursorStatement':code='CURSOR';advice='Bandingkan rewrite set-based; pertahankan urutan/side effects per baris.'
                elif t=='WhileStatement':code='WHILE';advice='Tinjau batch/set-based; loop retry/control-flow belum tentu masalah.'
                elif t=='SelectStarExpression':code='SELECT_STAR';advice='Batasi projection pada result besar; SELECT * di EXISTS tidak bermasalah sendiri.'
                elif t=='QuerySpecification' and n.get('UniqueRowFilter')=='Distinct':code='DISTINCT';advice='Periksa fan-out JOIN; jangan hapus dedup tanpa bukti cardinality.'
                elif t=='BinaryQueryExpression' and n.get('All')=='False':code='UNION_DEDUP';advice='UNION ALL hanya bila dedup tidak diperlukan secara bisnis.'
                elif t=='BooleanBinaryExpression' and n.get('BinaryExpressionType')=='Or':code='OR_PREDICATE';advice='Periksa optional filter dan selektivitas; branch/parameterized SQL atau recompile terukur.'
                elif t=='LikePredicate':
                    second=d.child(i,'SecondExpression');txt=compact(d.text(second)) if second is not None else ''
                    first=d.child(i,'FirstExpression')
                    if first is not None and list(d.descendants(first,{'ColumnReferenceExpression'}))+([first] if d.nodes[first]['t']=='ColumnReferenceExpression' else []) and re.match(r"N?'%",txt,re.I):code='LEADING_WILDCARD';advice='B-tree tidak membantu contains search; prefix/full-text sesuai kebutuhan tokenization.'
                elif t in {'ExecutableStringList'}:code='DYNAMIC_SQL';advice='SQL dinamis tidak di-bind ulang; kumpulkan generated SQL dan parameter aktual.'
                elif t=='ExecutableProcedureReference':
                    txt=compact(d.text(i))
                    if re.search(r'\bsp_executesql\b',txt,re.I):code='DYNAMIC_SQL';advice='Periksa deklarasi tipe parameter dan query dinamis yang dieksekusi.'
                elif t in {'FunctionCall','CastCall','ConvertCall','TryCastCall','TryConvertCall','LeftFunctionCall','RightFunctionCall'}:
                    pred=d.ancestor(i,{'BooleanComparisonExpression','LikePredicate','InPredicate','BooleanTernaryExpression','BooleanIsNullExpression'})
                    if pred is not None and d.scope[pred]==d.scope[i] and self.context(d,pred) in {'WHERE','JOIN_ON','HAVING','DML_CONDITION'}:
                        colrefs=list(d.descendants(i,{'ColumnReferenceExpression'}))
                        if colrefs:
                            name=d.value(i,'FunctionName') or t
                            code='FUNCTION_ON_COLUMN_IN_PREDICATE';advice='Tinjau sargability: pindahkan konversi/fungsi ke nilai pencarian; gunakan rentang tanggal. CAST(date) kadang dynamic seek; ukur residual/read.'
                if code:
                    patterns.append({'database':d.db,'file':d.path,'line':n['l'],'pattern':code,'snippet':compact(d.text(i))[:700],'advice':advice,'status':'REVIEW_CANDIDATE'})
                if t=='FunctionCall':
                    target=d.child(i,'CallTarget');functionname=d.value(i,'FunctionName')
                    if target is not None and d.nodes[target]['t']=='MultiPartIdentifierCallTarget':
                        parts=d.identifiers(d.child(target,'MultiPartIdentifier'))+[functionname]
                        key=self.key(parts,d.db)
                        if key in self.objects and self.objects[key]['kind']=='FUNCTIONS':
                            patterns.append({'database':d.db,'file':d.path,'line':n['l'],'pattern':'SCALAR_UDF_CALL','snippet':compact(d.text(i))[:700],'advice':'Periksa isi function dan actual inlining; compatibility/version tidak diketahui. Set-based APPLY/JOIN hanya setelah cardinality dan urutan hasil terbukti.','status':'REVIEW_CANDIDATE'})
            inventory.append({'database':d.db,'file':d.path,'parse_errors':len(d.errors),'query_scopes':sum(len(d.bytype[x]) for x in SCOPES),'comparisons':len(d.bytype['BooleanComparisonExpression']),'references':len(refs_in_doc),'dynamic_sites':sum(p['pattern']=='DYNAMIC_SQL' for p in patterns if p['file']==d.path and p['database']==d.db)})
        writecsv(out/'implicit_conversions.csv',conversions)
        writecsv(out/'unresolved_predicates.csv',unbound)
        writecsv(out/'query_patterns.csv',patterns)
        writecsv(out/'object_references.csv',refs)
        writecsv(out/'cfmast_predicates.csv',cf)
        writecsv(out/'predicate_access.csv',access)
        writecsv(out/'module_coverage.csv',inventory)
        writecsv(out/'existing_indexes.csv',[dict(x,keys='|'.join(x['keys']),includes='|'.join(x['includes'])) for x in self.indexes])
        columns=[]
        for key,obj in self.objects.items():
            if obj['kind']=='TABLE':
                columns.extend({'table':key,'column':v['column'],'type':v['sql_type'],'collation':v['collation'],'source':v['source']} for v in obj['cols'].values())
        writecsv(out/'table_columns.csv',columns)
        counts={'sources':dict(collections.Counter(d.db+'/'+d.path.split('/')[0] for d in self.docs)), 'parser':'Microsoft.SqlServer.TransactSql.ScriptDom 180.140.0; TSql170Parser; QUOTED_IDENTIFIER=True', 'parse_errors':[{'database':d.db,'file':d.path,'errors':d.errors} for d in self.docs if d.errors], 'predicate_binding':dict(self.predstatus),'conversion_categories':dict(collections.Counter(x['category'] for x in conversions)),'conversion_procedures':len({(x['database'],x['file']) for x in conversions if x['priority']!='INFO' and x['file'].startswith('procedures/')}),'patterns':dict(collections.Counter(x['pattern'] for x in patterns)), 'indexes':len(self.indexes),'table_columns':len(columns),'cfmast_predicates':len(cf),'cfmast_modules':len({x['file'] for x in cf})}
        (out/'summary.json').write_text(json.dumps(counts,indent=2),encoding='utf8')
        print(json.dumps(counts,indent=2))

def main():
    p=argparse.ArgumentParser();p.add_argument('--workspace',type=pathlib.Path,required=True);p.add_argument('--ast',type=pathlib.Path,required=True);p.add_argument('--replicate-ast',type=pathlib.Path,required=True);p.add_argument('--output',type=pathlib.Path,default=pathlib.Path(__file__).parent)
    args=p.parse_args();docs=[];inputs=[]
    for ast,relative,db in [(args.ast,'bond-sqltrsretail-uat/db_object','retail'),(args.replicate_ast,'all_object/bond/bond-sqlreplicate/db_object','sql_replicate')]:
        root=args.workspace/relative
        with ast.open(encoding='utf8') as f:
            for line in f:
                d=Doc(json.loads(line),root,db);docs.append(d)
                inputs.append({'path':str(d.full.relative_to(args.workspace)).replace('\\','/'),'sha256':hashlib.sha256(d.full.read_bytes()).hexdigest()})
    args.output.mkdir(parents=True,exist_ok=True)
    writecsv(args.output/'source_manifest.csv',inputs)
    Audit(docs).run(args.output)

if __name__=='__main__':main()
