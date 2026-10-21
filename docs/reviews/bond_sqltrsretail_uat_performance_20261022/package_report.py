"""Render workload-oriented appendices and selectable proposed DDL; no database access."""
import collections
import csv
import json
import pathlib
from analyze import norm, writecsv

ROOT=pathlib.Path(__file__).resolve().parent
def load(name):
    with (ROOT/name).open(encoding='utf-8-sig',newline='') as f:return list(csv.DictReader(f))

def main():
    candidates=json.loads((ROOT/'recommendations.json').read_text(encoding='utf8'))
    access=load('predicate_access.csv'); indexes=load('existing_indexes.csv'); cols=load('table_columns.csv')
    conversions=load('implicit_conversions.csv');patterns=load('query_patterns.csv');coverage=load('module_coverage.csv')
    workload=[];candidate_rows=[]
    sql=["-- PROPOSAL ONLY, generated 2026-10-22. No statement has been applied.",
         "-- Choose ONE candidate and its database after reviewing REPORT.md and live index definitions.",
         "-- A/B/C are review order, not measured performance benefit.",
         "-- SQL Server version/edition, replication policy, build window, log/tempdb space remain unverified.",
         "DECLARE @CandidateId varchar(10) = NULL; -- e.g. 'IX01'; IX17/IX18 require SQL_REPLICATE context",
         "DECLARE @ExpectedRetailDatabase sysname = NULL; -- actual retail UAT DB name, required for IX01..IX16",
         "IF @CandidateId IS NULL THROW 51000, 'Set one reviewed CandidateId; this script defaults to no DDL.', 1;",
         "IF @CandidateId NOT IN ("+', '.join("'"+x['id']+"'" for x in candidates)+") THROW 51000, 'Unknown CandidateId.', 1;",""]
    md=['# Kandidat index gabungan seluruh workload','',
        'Tanggal: 22 Oktober 2026. Status seluruh kandidat: **PROPOSED, belum dieksekusi/diukur**. A = tinjau awal, B = setelah rewrite/ukur kebutuhan, C = kondisional. Angka procedure adalah cakupan source, bukan frekuensi eksekusi.','',
        'Existing prefix di bawah mengabaikan INCLUDE saat menentukan search key. Daftar prefix bukan bukti coverage penuh/aktif di live DB. Predicate dalam OR tidak dianggap otomatis equality wajib.','',
        '| ID | Prioritas | Tabel dan key | Procedure dengan equality key pertama | Existing prefix terkait |',
        '|---|---|---|---:|---|']
    detailed=[]
    for x in candidates:
        names=[k.split()[0] for k in x['keys']];lower=[norm(k) for k in names]
        schema=[r for r in cols if r['table']==x['table']];known={norm(r['column']) for r in schema}
        assert all(norm(k) in known for k in names+x['include']),x['id']
        relevant=[r for r in access if r['table']==x['table'] and norm(r['column'])==lower[0] and r['operator']=='Equals' and r['file'].startswith('procedures/')]
        procedures=sorted({r['file'] for r in relevant})
        existing=[r for r in indexes if r['table']==x['table']]
        prefix=[r for r in existing if [norm(k) for k in r['keys'].split('|')][:min(len(lower),len(r['keys'].split('|')))]==lower[:min(len(lower),len(r['keys'].split('|')))]]
        fullprefix=[r for r in existing if [norm(k) for k in r['keys'].split('|')][:len(lower)]==lower]
        assert not fullprefix,(x['id'],'duplicate proposed search-key prefix',fullprefix)
        indexname='IX_Audit20261022_'+x['object'].split('.')[-1]+'_'+'_'.join(names)
        assert len(indexname)<=128
        obj='.'.join('['+a+']' for a in x['object'].split('.'))
        keysql=', '.join('['+k.split()[0]+'] '+('DESC' if k.endswith(' DESC') else 'ASC') for k in x['keys'])
        ddl=f'CREATE NONCLUSTERED INDEX [{indexname}]\nON {obj} ({keysql})\nINCLUDE ('+', '.join('['+c+']' for c in x['include'])+');'
        for r in relevant:workload.append(dict(candidate=x['id'],**r))
        # IX18 is contingent on a proposed pushdown, so its base equality frequency is intentionally zero.
        candidate_rows.append({'id':x['id'],'priority':x['priority'],'table':x['table'],'keys':'|'.join(x['keys']),'include':'|'.join(x['include']),'procedure_count_key1_equality':len(procedures),'predicate_occurrences_key1_equality':len(relevant),'existing_prefixes':'|'.join(r['name'] for r in prefix),'reason':x['reason'],'prerequisite':x['prerequisite'],'status':'PROPOSED_NOT_MEASURED'})
        md.append(f"| {x['id']} | {x['priority']} | `{x['object']} ({', '.join(x['keys'])})` | {len(procedures)} | {', '.join(r['name'] for r in prefix) or 'Tidak ada'} |")
        detailed += [f"## {x['id']} — {x['object']}",'',x['reason'],'',f"**Prasyarat:** {x['prerequisite']}",'',
                     f"**Bukti workload:** {len(relevant)} predicate kesetaraan key pertama pada {len(procedures)} procedure; detail seluruh lokasi pada `index_workload.csv`.",'',
                     '**Contoh konsumen:** '+('; '.join(f"`{r['file']}:{r['line']}`" for r in relevant[:5]) or '`procedures/dbo.TRSRPTMasterNasabahPerubahanObligasi2.sql:251` (date pushdown yang diusulkan).'),'',
                     '**Existing index pada tabel:** '+('; '.join(f"{r['name']} ({r['keys']}) INCLUDE ({r['includes']})" for r in existing) or 'Tidak ada index dalam snapshot.'),'','```sql',ddl,'```','']
        sql += [f"-- {x['id']}: {x['reason']}",f"-- Prerequisite: {x['prerequisite']}",f"IF @CandidateId = '{x['id']}'",'BEGIN']
        if x['table'].startswith('sql_replicate.'):
            sql.append("    IF DB_NAME() <> N'SQL_REPLICATE' THROW 51000, 'Use SQL_REPLICATE database for this candidate.', 1;")
        else:sql.append("    IF @ExpectedRetailDatabase IS NULL OR DB_NAME() <> @ExpectedRetailDatabase THROW 51000, 'Set and select the verified retail UAT database.', 1;")
        sql += [f"    IF OBJECT_ID(N'{x['object']}', N'U') IS NULL THROW 51000, 'Expected base table missing.', 1;",
                f"    IF EXISTS (SELECT 1 FROM sys.indexes WHERE object_id = OBJECT_ID(N'{x['object']}') AND name = N'{indexname}') THROW 51000, 'Candidate name already exists; compare definitions.', 1;",
                "    -- Compare equivalent/prefix keys and INCLUDE in catalog_validation.sql before proceeding.",
                "    EXEC sys.sp_executesql N'"+ddl.replace("'","''")+"';",'END;','']
    writecsv(ROOT/'index_candidates.csv',candidate_rows);writecsv(ROOT/'index_workload.csv',workload)
    (ROOT/'index_candidates.md').write_text('\n'.join(md+['']+detailed),encoding='utf8')
    (ROOT/'recommended_indexes.sql').write_text('\n'.join(sql),encoding='utf8')
    proc=[]
    for c in coverage:
        if c['database']!='retail' or not c['file'].startswith('procedures/'):continue
        rs=[r for r in conversions if r['database']=='retail' and r['file']==c['file']]
        ps=[r for r in patterns if r['database']=='retail' and r['file']==c['file']]
        ids=[x['id'] for x in candidates if any(r['candidate']==x['id'] and r['file']==c['file'] for r in workload)]
        proc.append({'procedure':pathlib.Path(c['file']).stem,'source':c['file'],'parse_errors':c['parse_errors'],
                     'text_to_numeric_column':sum(r['category']=='COLUMN_TEXT_TO_NUMERIC' and r['context'] in {'WHERE','JOIN_ON','DML_CONDITION'} for r in rs),
                     'text_to_date_column':sum(r['category']=='COLUMN_TEXT_TO_DATE' and r['context'] in {'WHERE','JOIN_ON','DML_CONDITION'} for r in rs),
                     'column_to_unicode':sum(r['category']=='COLUMN_TO_UNICODE' and r['context'] in {'WHERE','JOIN_ON','DML_CONDITION'} for r in rs),
                     'patterns':'|'.join(f'{k}:{v}' for k,v in sorted(collections.Counter(r['pattern'] for r in ps).items())),
                     'index_candidates':'|'.join(ids),'assessment':'STATIC_REVIEW_CANDIDATES' if rs or ps else 'NO_MATCH_BY_BOUNDED_RULES_NOT_CERTIFIED'})
    writecsv(ROOT/'procedure_summary.csv',proc)
    risks=sorted([r for r in proc if sum(int(r[k]) for k in ['text_to_numeric_column','text_to_date_column','column_to_unicode'])],key=lambda r:(-sum(int(r[k]) for k in ['text_to_numeric_column','text_to_date_column','column_to_unicode']),r['procedure']))
    listing=['# Listing procedure dengan risiko konversi kolom pada filter/join','',
             'Daftar ini mencakup kandidat dari binding statis, termasuk temporary table; belum membuktikan scan atau biaya runtime. Perbandingan CASE/projection, perbedaan numeric family, dan CHAR/VARCHAR biasa tidak dihitung di tabel ini. Semua 1.063 procedure tercakup dalam `procedure_summary.csv`. Lokasi dan arah konversi lengkap pada `implicit_conversions.csv`.','',
             '| Procedure | Teks → numerik | Teks → tanggal | ANSI → Unicode | Kandidat index |','|---|---:|---:|---:|---|']
    listing += [f"| [{r['procedure']}](../../../bond-sqltrsretail-uat/db_object/{r['source']}) | {r['text_to_numeric_column']} | {r['text_to_date_column']} | {r['column_to_unicode']} | {r['index_candidates']} |" for r in risks]
    (ROOT/'procedure_listing.md').write_text('\n'.join(listing)+'\n',encoding='utf8')
    print(json.dumps({'index_candidates':len(candidates),'workload_rows':len(workload),'procedure_summary':len(proc),'risk_procedure_listing':len(risks)},indent=2))

if __name__=='__main__':main()
