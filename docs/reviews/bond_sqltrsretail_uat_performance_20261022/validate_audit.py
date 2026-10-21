"""Independent artifact checks and small synthetic scope/type regression corpus.

No SQL Server access; all SQL is parsed offline. This is audit validation, not
proof of actual optimizer plans or a test of the repository migration engine.
"""
import argparse
import collections
import csv
import hashlib
import json
import pathlib
import subprocess
import tempfile
from analyze import Audit, Doc, compact, writecsv

HERE=pathlib.Path(__file__).resolve().parent
def readcsv(name):
    with (HERE/name).open(encoding='utf-8-sig',newline='') as f:return list(csv.DictReader(f))
def sha(path):return hashlib.sha256(path.read_bytes()).hexdigest()
def protected_paths(workspace):
    for relative in ['bond-sqltrsretail-uat/db_object','bond-sqltrsretail-uat/migrations','all_object/bond/bond-sqlreplicate/db_object','all_object/bond/bond-sqlreplicate/migrations']:
        yield from sorted((workspace/relative).rglob('*.sql'))

def main():
    ap=argparse.ArgumentParser();ap.add_argument('--workspace',required=True,type=pathlib.Path);ap.add_argument('--dll',required=True,type=pathlib.Path);ap.add_argument('--capture-protected',action='store_true');args=ap.parse_args();workspace=args.workspace.resolve()
    checks=[]
    def check(name,condition,detail):
        checks.append({'id':name,'pass':bool(condition),'detail':detail})
        if not condition:raise AssertionError(name+': '+str(detail))
    protected=HERE/'protected_manifest.csv'
    if args.capture_protected:
        if protected.exists():raise ValueError('Will not replace existing protected baseline')
        writecsv(protected,[{'path':p.relative_to(workspace).as_posix(),'sha256':sha(p)} for p in protected_paths(workspace)])
    baseline=readcsv('protected_manifest.csv')
    current={p.relative_to(workspace).as_posix():sha(p) for p in protected_paths(workspace)}
    expected={r['path']:r['sha256'] for r in baseline}
    check('protected_sql_unchanged',current==expected,{'files':len(current),'migration_files':sum('/migrations/' in p for p in current),'boundary':'No source/migration mutation since protected baseline capture'})
    manifest=readcsv('source_manifest.csv')
    check('all_input_hashes_match',len(manifest)==2756 and all(sha(workspace/r['path'])==r['sha256'] for r in manifest),{'sources':len(manifest)})
    rows=readcsv('implicit_conversions.csv');texts={}
    for r in rows:
        base='bond-sqltrsretail-uat/db_object' if r['database']=='retail' else 'all_object/bond/bond-sqlreplicate/db_object'
        key=(base,r['file'])
        if key not in texts:texts[key]=compact((workspace/base/r['file']).read_bytes().decode('utf-8-sig'))
        check_ok=r['predicate'] in texts[key]
        if not check_ok:raise AssertionError('Snippet does not occur in source: '+str(r))
    check('conversion_evidence_snippets_exist',True,{'rows':len(rows),'note':'Every compact predicate matches exact case-preserved source text'})
    oracle=[
        ('procedures/dbo.TRSRPTMasterNasabahPerubahanObligasi2.sql',45,'COLUMN_TEXT_TO_NUMERIC','CHAR(19)','BIGINT','left'),
        ('procedures/dbo.TRSUpdateDataMasterNasabah.sql',121,'COLUMN_TEXT_TO_NUMERIC','BIGINT','CHAR(19)','right'),
        ('procedures/dbo.TRSSyncMasterRiskProfileData.sql',24,'COLUMN_TEXT_TO_NUMERIC','BIGINT','VARCHAR(19)','right'),
        ('procedures/dbo.PTOSProsesDataJobNasabah.sql',94,'COLUMN_TO_UNICODE','NVARCHAR(255)','CHAR(40)','right'),
        ('procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql',108,'COLUMN_TEXT_TO_NUMERIC','CHAR(7)','INT','left'),
        ('procedures/dbo.TRSHasilRevenueBondJurnalGL.sql',134,'COLUMN_TEXT_TO_DATE','VARCHAR(8)','DATETIME','left'),
        ('procedures/dbo.ORISBNIPopulateEarlyRedemption.sql',33,'COLUMN_TEXT_TO_NUMERIC','VARCHAR(100)','INT','left'),
    ]
    for f,line,cat,lt,rt,side in oracle:
        check('oracle_'+pathlib.Path(f).stem+'_'+str(line),any(r['file']==f and int(r['line'])==line and r['category']==cat and r['left_type']==lt and r['right_type']==rt and r['converted_side']==side for r in rows),{'file':f,'line':line,'expected':[cat,lt,rt,side]})
    check('case_not_misreported_as_filter',any(r['file']=='procedures/dbo.TRSSyncMasterIdAndFunc.sql' and r['context']=='CASE' for r in rows),'CFIEX7 CASE conversions have separate context')
    check('cfmast_padded_text_not_numeric_risk',not any(r['file']=='procedures/dbo.TRSGetJenisNasabahCIF.sql' and r['category']=='COLUMN_TEXT_TO_NUMERIC' for r in rows),'CFCIF is CHAR(19), not CFCIF# BIGINT')
    check('all_retail_procedures_in_summary',len(readcsv('procedure_summary.csv'))==1063,'One coverage row per procedure, including parse failure')
    idx=readcsv('existing_indexes.csv')
    check('clustered_and_constraints_included',any(r['table']=='retail.dbo.securitymaster_tm' and r['keys']=='SecId' and r['unique']=='True' for r in idx),'PK index counted separately from index files')
    candidate_rows=readcsv('index_candidates.csv')
    check('securityno_workload_reconciles',next(r for r in candidate_rows if r['id']=='IX01')['procedure_count_key1_equality']=='148','148 source procedures, not execution frequency')
    # Independent synthetic inputs. CRLF and non-BMP comment test UTF-16 offsets.
    fixtures={
        'tables/dbo.AuditChar.sql':"CREATE TABLE dbo.AuditChar (CFCIF char(19), [CFCIF#] bigint, Status varchar(3));",
        'tables/dbo.AuditOther.sql':"CREATE TABLE dbo.AuditOther (CIFNumber bigint);",
        'indexes/dbo.AuditChar.sql':"CREATE INDEX IX_Later ON dbo.AuditChar ([CFCIF#],CFCIF);",
        'views/dbo.AuditV.sql':"CREATE OR ALTER VIEW dbo.AuditV AS SELECT CFCIF,[CFCIF#],Status FROM dbo.AuditChar;",
        'procedures/dbo.AuditFixture.sql':"""-- Unicode offset sentinel: 😀
CREATE OR ALTER PROCEDURE dbo.AuditFixture @p bigint, @s varchar(19) AS
SELECT 1 FROM dbo.AuditV c WHERE c.CFCIF = @p;
SELECT 1 FROM dbo.AuditV c WHERE c.[CFCIF#] = @s;
SELECT CASE WHEN c.CFCIF = @p THEN 1 ELSE 0 END FROM dbo.AuditV c;
SELECT c.CFCIF FROM dbo.AuditV c WHERE EXISTS (SELECT 1 FROM dbo.AuditOther c WHERE c.CIFNumber = @p) AND c.CFCIF = @s;
WITH q AS (SELECT [CFCIF#] AS id FROM dbo.AuditChar) SELECT 1 FROM q WHERE id = @s;
CREATE TABLE #Stage (Id nvarchar(40));
SELECT 1 FROM dbo.AuditV c JOIN #Stage t ON c.CFCIF = t.Id;
SELECT CFCIF INTO #Derived FROM dbo.AuditChar;
SELECT 1 FROM #Derived x WHERE x.CFCIF = @p;
SELECT 1 FROM dbo.AuditV c WHERE c.Status IN (1,2);
EXEC('SELECT 1 FROM dbo.AuditChar WHERE CFCIF = 3');
""",
    }
    with tempfile.TemporaryDirectory(prefix='rza-uat-audit-validation-') as tmp:
        tmp=pathlib.Path(tmp);root=tmp/'input';root.mkdir()
        for rel,content in fixtures.items():
            target=root/rel;target.parent.mkdir(parents=True,exist_ok=True);target.write_bytes(content.replace('\n','\r\n').encode('utf8'))
        ast=tmp/'fixture_ast.jsonl'
        cmd=['pwsh','-NoProfile','-File',str(HERE/'extract_ast.ps1'),'-SourceRoot',str(root),'-ScriptDomDll',str(args.dll.resolve()),'-OutputFile',str(ast)]
        result=subprocess.run(cmd,capture_output=True,text=True,check=True)
        docs=[Doc(json.loads(s),root,'retail') for s in ast.read_text(encoding='utf8').splitlines()]
        check('synthetic_fixture_parse',all(not d.errors for d in docs),result.stdout.strip())
        audit=Audit(docs);out=tmp/'output';out.mkdir();audit.run(out)
        with (out/'implicit_conversions.csv').open(encoding='utf-8-sig') as f:rs=list(csv.DictReader(f))
        def exists(pred,cat,context=None):return any(r['predicate']==pred and r['category']==cat and (context is None or r['context']==context) for r in rs)
        check('synthetic_numeric_column_direction',exists('c.CFCIF = @p','COLUMN_TEXT_TO_NUMERIC','WHERE'),'View lineage + parameter precedence')
        check('synthetic_parameter_direction',exists('c.[CFCIF#] = @s','OPERAND_TEXT_TO_NUMERIC'),'No column conversion for BIGINT column vs VARCHAR parameter')
        check('synthetic_shadowed_alias',not any(r['predicate']=='c.CIFNumber = @p' for r in rs),'Nested alias resolves AuditOther.CIFNumber BIGINT, not outer CHAR table')
        check('synthetic_cte_lineage',exists('id = @s','OPERAND_TEXT_TO_NUMERIC'),'CTE alias output id inherits BIGINT')
        check('synthetic_unicode_direction',exists('c.CFCIF = t.Id','COLUMN_TO_UNICODE'),'CHAR key vs NVARCHAR temporary key')
        check('synthetic_case_context',exists('c.CFCIF = @p','COLUMN_TEXT_TO_NUMERIC','CASE'),'CASE is not a WHERE/JOIN access predicate')
        check('synthetic_select_into_lineage',exists('x.CFCIF = @p','COLUMN_TEXT_TO_NUMERIC'),'SELECT INTO inherits CHAR, keeps temp identity')
        check('synthetic_in_numeric',exists('c.Status IN (1,2)','COLUMN_TEXT_TO_NUMERIC'),'IN list grouped by operand type')
        check('synthetic_index_order',audit.leading('retail.dbo.auditchar','CFCIF')=='' and audit.leading('retail.dbo.auditchar','CFCIF#')=='IX_Later','Second index key does not count as leading-key search')
        check('synthetic_dynamic_not_bound',not any('CFCIF = 3' in r['predicate'] for r in rs),'SQL strings are dynamic boundaries, not static SQL comparisons')
        check('synthetic_utf16_crlf_spans',all(r['predicate'] in compact((root/r['file']).read_bytes().decode('utf8')) for r in rs),'CRLF and non-BMP text before AST spans')
        # Parse both proposal batch and actual CREATE statements within dynamic strings.
        sqlroot=tmp/'sql';sqlroot.mkdir();(sqlroot/'batch.sql').write_bytes((HERE/'recommended_indexes.sql').read_bytes())
        import re
        proposal=(HERE/'recommended_indexes.sql').read_text(encoding='utf8')
        statements=re.findall(r"EXEC sys\.sp_executesql N'((?:[^']|'')*)';",proposal)
        for no,stmt in enumerate(statements):(sqlroot/f'index_{no:02}.sql').write_text(stmt.replace("''","'"),encoding='utf8')
        (sqlroot/'catalog.sql').write_bytes((HERE/'catalog_validation.sql').read_bytes())
        result=subprocess.run(['pwsh','-NoProfile','-File',str(HERE/'extract_ast.ps1'),'-SourceRoot',str(sqlroot),'-ScriptDomDll',str(args.dll.resolve()),'-OutputFile',str(tmp/'proposal_ast.jsonl')],capture_output=True,text=True,check=True)
        proposal_ast=[json.loads(s) for s in (tmp/'proposal_ast.jsonl').read_text().splitlines()]
        check('proposed_sql_syntax',len(statements)==18 and all(not d['errors'] for d in proposal_ast),{'create_statements':len(statements),'parsed_files':len(proposal_ast),'runtime':'NOT_RUN'})
    # Idempotence: rerender outputs and compare byte hashes of all declared generated artifacts.
    generated=['index_candidates.csv','index_candidates.md','index_workload.csv','recommended_indexes.sql','procedure_summary.csv','procedure_listing.md']
    before={p:sha(HERE/p) for p in generated}
    subprocess.run(['python',str(HERE/'package_report.py')],check=True,capture_output=True,text=True)
    check('render_idempotent',before=={p:sha(HERE/p) for p in generated},generated)
    check('protected_sql_still_unchanged',current=={p.relative_to(workspace).as_posix():sha(p) for p in protected_paths(workspace)},'After synthetic tests and report rerender')
    head=subprocess.run(['git','rev-parse','HEAD'],cwd=workspace,capture_output=True,text=True,check=True).stdout.strip()
    receipt={'success':True,'date':'2026-10-22','head_at_validation':head,'scriptdom_dll_sha256':sha(args.dll),'checks':checks,'limitations':{'SQL_Server_runtime':'NOT_RUN','actual_execution_plans':'NOT_AVAILABLE','engine_go_tests_build':'NOT_RUN_NOT_APPLICABLE_REPORT_ONLY','audit_independent_agent_evaluation':'NOT_RUN'},'artifact_sha256':{p:sha(HERE/p) for p in generated+['REPORT.md','analyze.py','extract_ast.ps1','package_report.py','validate_audit.py','recommendations.json','source_manifest.csv']}}
    (HERE/'validation.json').write_text(json.dumps(receipt,indent=2),encoding='utf8')
    print(json.dumps({'success':True,'checks':len(checks),'head':head,'protected_sql_files':len(current)},indent=2))

if __name__=='__main__':main()
