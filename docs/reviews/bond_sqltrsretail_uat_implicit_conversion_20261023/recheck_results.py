"""Compare recheck to preserved evidence and inventory remaining blind spots."""
import argparse
import collections
import hashlib
import csv
import gzip
import io
import json
import pathlib
from review_conversions import HERE, Doc, load_csv, norm, compact, writecsv

def location(r):return tuple(r[k] for k in ['database','file','line','context','form','left','right'])
def typing(r):return tuple(r[k] for k in ['left_type','right_type','converted_side','target_type','left_origin','right_origin'])
def riskkey(r):return tuple(r[k] for k in ['table','column','file','line','context','form','column_operand','other_operand'])
def baseline_csv(root,name):
    path=root/name
    if path.exists():return load_csv(path)
    return list(csv.DictReader(io.StringIO(gzip.decompress((root/(name+'.gz')).read_bytes()).decode('utf-8-sig'))))

def main():
    ap=argparse.ArgumentParser();ap.add_argument('--workspace',required=True,type=pathlib.Path);ap.add_argument('--baseline',required=True,type=pathlib.Path);ap.add_argument('--retail-ast',required=True,type=pathlib.Path);ap.add_argument('--replicate-ast',required=True,type=pathlib.Path);args=ap.parse_args();w=args.workspace.resolve()
    before=baseline_csv(args.baseline,'conversion_sites.csv');after=load_csv(HERE/'conversion_sites.csv')
    oldmap=collections.defaultdict(list);newmap=collections.defaultdict(list)
    for r in before:oldmap[location(r)].append(r)
    for r in after:newmap[location(r)].append(r)
    added=[];removed=[]
    for r in after:
        prior=oldmap.get(location(r),[])
        if not prior:added.append(dict(change='NEW_BINDING_OR_OPERATOR',**r))
        elif typing(r) not in {typing(p) for p in prior}:added.append(dict(change='TYPE_OR_ORIGIN_RECLASSIFIED',**r))
    for r in before:
        if location(r) not in newmap:removed.append(dict(change='NO_LONGER_REPORTED_AT_THIS_PAIR',**r))
    writecsv(HERE/'recheck_added_or_retyped.csv',added);writecsv(HERE/'recheck_removed_pairs.csv',removed)
    def primary_a(rows):return [r for r in rows if r['tier']=='A_COLUMN_PREDICATE' and r['module_database']=='retail' and r['file'].startswith('procedures/')]
    olda=primary_a(baseline_csv(args.baseline,'column_occurrences.csv'));newa=primary_a(load_csv(HERE/'column_occurrences.csv'))
    oldkeys={riskkey(r) for r in olda};newkeys={riskkey(r) for r in newa}
    ar=[dict(change='NEW_PHYSICAL_COLUMN_PREDICATE',**r) for r in newa if riskkey(r) not in oldkeys]
    writecsv(HERE/'recheck_new_column_predicates.csv',ar)
    # Do not infer that dynamic execution text is static merely because it has literals.
    dynamic=[];temp=[];reverse=collections.defaultdict(set)
    modulepaths={}
    for r in load_csv(HERE/'source_manifest.csv'):
        obj=r['path'].split('/db_object/',1)[1]
        if obj.split('/')[0] not in {'functions','views','procedures'}:continue
        db='sql_replicate' if r['path'].startswith('all_object/') else 'retail'
        modulepaths[db+'.'+norm(pathlib.Path(obj).stem)]={'path':obj,'source_path':r['path'],'database':db}
    for r in load_csv(HERE/'object_references.csv'):
        source=r['database']+'.'+norm(pathlib.Path(r['file']).stem)
        if r['object'] in modulepaths:reverse[r['object']].add(source)
    for ast,db,root in [(args.retail_ast,'retail',w/'bond-sqltrsretail-uat/db_object'),(args.replicate_ast,'sql_replicate',w/'all_object/bond/bond-sqlreplicate/db_object')]:
        for line in ast.open(encoding='utf8'):
            d=Doc(json.loads(line),root,db)
            if d.path.split('/')[0] not in {'functions','views','procedures'}:continue
            source=db+'.'+norm(pathlib.Path(d.path).stem)
            for i in d.bytype['FunctionCall']:
                t=d.child(i,'CallTarget')
                if t is not None and d.nodes[t]['t']=='MultiPartIdentifierCallTarget':
                    parts=d.identifiers(d.child(t,'MultiPartIdentifier'))+[d.value(i,'FunctionName')]
                    key=db+'.dbo.'+norm(parts[0]) if len(parts)==1 else db+'.'+'.'.join(map(norm,parts)) if len(parts)==2 else '.'.join(map(norm,parts))
                    if key in modulepaths:reverse[key].add(source)
            if not d.path.startswith('procedures/') or db!='retail':continue
            for t in ['ExecutableStringList','ExecutableProcedureReference']:
                for i in d.bytype[t]:
                    code=compact(d.text(i));dynamic_kind=None
                    if t=='ExecutableStringList':dynamic_kind='EXEC_STRING_OR_VARIABLE'
                    elif 'sp_executesql' in code.lower():dynamic_kind='SP_EXECUTESQL'
                    elif code.startswith('@'):dynamic_kind='DYNAMIC_PROCEDURE_NAME'
                    if dynamic_kind:
                        dynamic.append({'file':d.path,'line':d.nodes[i]['l'],'kind':dynamic_kind,'status':'UNRESOLVED_DYNAMIC_RUNTIME_TEXT_OR_BINDINGS','execution_site':code,'limitation':'Reaching assignments/branches/parameter declarations not fully resolved; not certified conversion-free'})
            definitions=collections.defaultdict(list)
            for i in d.bytype['CreateTableStatement']:
                parts=d.identifiers(d.child(i,'SchemaObjectName'))
                if parts and parts[-1].startswith('#'):
                    cs=d.columns(i,'TEMP',parts[-1]);sig=tuple((k,v['sql_type'].lower(),v['collation']) for k,v in cs.items())
                    definitions[norm(parts[-1])].append(('DECLARE',d.nodes[i]['l'],sig))
            for i in d.bytype['SelectStatement']:
                parts=d.identifiers(d.child(i,'Into'))
                if parts and parts[-1].startswith('#'):definitions[norm(parts[-1])].append(('SELECT_INTO',d.nodes[i]['l'],()))
            for name,defs in definitions.items():
                if len(defs)<2:continue
                declared={sig for typ,line,sig in defs if typ=='DECLARE'}
                status='DIVERGENT_DECLARED_SCHEMA' if len(declared)>1 else 'MULTIPLE_DEFINITIONS_BRANCH_REVIEW' if any(typ=='SELECT_INTO' for typ,line,sig in defs) else 'REPEATED_SAME_DECLARED_SCHEMA'
                temp.append({'file':d.path,'temp_object':name,'definition_lines':'|'.join(str(line) for typ,line,sig in defs),'definition_kinds':'|'.join(typ for typ,line,sig in defs),'status':status,'limitation':'Current binder is not a full path-sensitive temp lifecycle/dataflow analysis'})
    writecsv(HERE/'remaining_dynamic_sql.csv',dynamic)
    writecsv(HERE/'remaining_temp_lifecycle.csv',temp,columns=['file','temp_object','definition_lines','definition_kinds','status','limitation'])
    impacts=[];truncated=0
    for site in added:
        if site['module_kind'] not in {'views','functions'}:continue
        key=site['database']+'.'+norm(pathlib.Path(site['file']).stem);todo=collections.deque([(key,[key])]);seen={key}
        while todo:
            node,path=todo.popleft()
            if len(path)>12:truncated+=1;continue
            for caller in sorted(reverse[node]):
                if caller in seen:continue
                seen.add(caller);info=modulepaths.get(caller)
                if not info:continue
                nextpath=path+[caller]
                if info['database']=='retail' and info['path'].startswith('procedures/'):
                    impacts.append({'site_id':site['site_id'],'body_file':site['file'],'body_line':site['line'],'form':site['form'],'category':site['category'],'predicate':site['predicate'],'consumer_procedure':info['path'],'call_path':' -> '.join(reversed(nextpath)),'reachability':'POSSIBLE_STATIC_CALL_PATH; branch/execution frequency not observed'})
                else:todo.append((caller,nextpath))
    writecsv(HERE/'recheck_supporting_impacts.csv',impacts)
    before_summary=json.loads((args.baseline/'summary.json').read_text());after_summary=json.loads((HERE/'summary.json').read_text())
    result={'baseline':before_summary,'after':after_summary,'new_location_pairs':sum(r['change']=='NEW_BINDING_OR_OPERATOR' for r in added),'retyped_or_new_origin_pairs':sum(r['change']=='TYPE_OR_ORIGIN_RECLASSIFIED' for r in added),'removed_location_pairs':len(removed),'new_physical_column_predicate_endpoints':len(ar),'removed_physical_column_predicate_endpoints':sum(riskkey(r) not in newkeys for r in olda),'new_physical_column_predicate_procedures':len({r['file'] for r in ar}),'new_physical_column_predicate_procedure_names':sorted({r['file'] for r in ar}),'new_predicate_forms':dict(collections.Counter(r['form'] for r in added)),'dynamic_execution_sites':len(dynamic),'dynamic_procedures':len({r['file'] for r in dynamic}),'temp_multiple_definition_objects':len(temp),'temp_divergent_declared_objects':sum(r['status']=='DIVERGENT_DECLARED_SCHEMA' for r in temp),'supporting_impacts':len(impacts),'supporting_affected_procedures':len({r['consumer_procedure'] for r in impacts}),'supporting_path_truncations':truncated,'notes':['New/retyped rows include family/routine conversions, not all severe findings.','Predicate A counts only procedure-local physical endpoints; supporting impacts are potential reachable effects, counted separately.','Dynamic, ambiguous temp flow, parser failure and other unsupported expressions remain explicitly incomplete.']}
    (HERE/'recheck_delta.json').write_text(json.dumps(result,indent=2),encoding='utf8')
    print(json.dumps({k:v for k,v in result.items() if k not in {'baseline','after','new_predicate_forms','new_physical_column_predicate_procedure_names'}},indent=2))

if __name__=='__main__':main()
