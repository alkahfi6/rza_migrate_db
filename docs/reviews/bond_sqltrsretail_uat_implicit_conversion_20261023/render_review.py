"""Render table-first study documents from complete, untruncated evidence CSVs."""
import collections
import csv
import json
import os
import pathlib
import re
from urllib.parse import quote

HERE=pathlib.Path(__file__).resolve().parent
WORKSPACE=HERE.parents[2]
PATTERNS={
 'P01':('Kolom teks bertemu numerik','Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.'),
 'P02':('ANSI bertemu Unicode','Samakan tipe/kolasi staging dengan key tujuan sesudah validasi panjang dan karakter. Jangan menurunkan Unicode jika karakter akan hilang; perubahan storage memerlukan review tersendiri.'),
 'P03':('Tanggal teks bertemu date/time','Tentukan kontrak YYYYMMDD/hari/timestamp. Bentuk nilai pencarian dalam tipe kolom atau gunakan rentang native date; jangan konversi kolom seluruh baris tanpa alasan.'),
 'P04':('Nilai/parameter dikonversi; kolom menjadi pasangan','Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.'),
 'P05':('Aritmetika/concatenation dengan tipe campuran','Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.'),
 'P06':('LIKE dan representasi karakter','LIKE pada key numerik memerlukan representasi karakter. Equality numeric yang tervalidasi cocok untuk lookup tepat; prefix/substring membutuhkan requirement pencarian tersendiri, bukan cast kolom sebagai solusi index.'),
 'P07':('CASE matching dan tipe hasil CASE/COALESCE/UNION','Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.'),
 'P08':('Assignment/INSERT/UPDATE mengikuti tipe tujuan','Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.'),
 'P09':('ISNULL dan NULLIF mempunyai aturan berbeda','ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.'),
 'P10':('Argumen function mengikuti signature','Konversi mengikuti parameter function yang dideklarasikan. Cocokkan signature dan pemanggil; function per row tetap membutuhkan review terpisah, tidak selalu masalah index.'),
 'P11':('Perbedaan keluarga numerik','BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.'),
 'P12':('CHAR/VARCHAR dan keluarga teks biasa','Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.'),
 'P13':('Perbedaan lain yang memerlukan verifikasi','Periksa apakah konversi tipe tersebut diizinkan pada operator yang dipakai. Tanpa oracle SQL Server, sebagian kasus dapat merupakan type incompatibility/error alih-alih conversion yang sukses.'),
}
def read(name):
    with (HERE/name).open(encoding='utf-8-sig',newline='') as f:return list(csv.DictReader(f))
def csvout(name,rows):
    fields=list(dict.fromkeys(k for r in rows for k in r))
    with (HERE/name).open('w',encoding='utf-8-sig',newline='') as f:
        w=csv.DictWriter(f,fieldnames=fields);w.writeheader();w.writerows(rows)
def md(s):return str(s).replace('|','&#124;').replace('\n',' ').replace('`','&#96;')
def link(fromfile,target,line=None):
    dest=os.path.relpath(target,fromfile.parent).replace('\\','/')
    return quote(dest,safe='/.:')+('#L'+str(line) if line else '')
def anchor(s):return re.sub('[^a-z0-9]+','-',s.lower()).strip('-')
def pattern(r):
    f=r['form'];cat=r.get('pattern',r.get('category',''));method=r['method']
    if method=='ASSIGNMENT':return 'P08'
    if method=='ISNULL_RETURN':return 'P09'
    if f=='NULLIF_COMPARE':return 'P09'
    if method=='UDF_PARAMETER':return 'P10'
    if f.startswith('BINARY_'):return 'P05'
    if f=='LIKE':return 'P06'
    if f in {'SIMPLE_CASE_MATCH','CASE_BRANCH','COALESCE_BRANCH','UNION_OUTPUT','EXCEPT_OUTPUT','INTERSECT_OUTPUT'}:return 'P07'
    if cat=='COLUMN_TEXT_TO_NUMERIC':return 'P01'
    if cat=='COLUMN_TO_UNICODE':return 'P02'
    if cat=='COLUMN_TEXT_TO_DATE':return 'P03'
    if cat.startswith('OPERAND_'):return 'P04'
    if cat=='NUMERIC_FAMILY':return 'P11'
    if cat=='TEXT_FAMILY_ONLY':return 'P12'
    return 'P13'
def proc(r):return r['module_database']=='retail' and r['file'].startswith('procedures/')
def nproc(rows):return len({r['file'] for r in rows})
def nsites(rows):return len({r['site_id'] for r in rows})
def tierorder(r):return ['A_COLUMN_PREDICATE','B_EXPRESSION_OR_BOUNDARY','C_FAMILY_OR_ROUTINE','PARTNER_OR_DESTINATION'].index(r['tier'])

def main():
    occurrences=read('column_occurrences.csv');sites=read('conversion_sites.csv');columns=read('table_columns.csv');refs=read('object_references.csv');coverage=read('module_coverage.csv')
    rows=[r for r in occurrences if proc(r)];risk=[r for r in rows if r['tier']=='A_COLUMN_PREDICATE']
    for r in occurrences:r['pattern_id']=pattern(r)
    groups=collections.defaultdict(list)
    for r in rows:groups[r['table']].append(r)
    catalogue=collections.defaultdict(list)
    for r in columns:catalogue[r['table']].append(r)
    # Keep original identifier spelling in all user-facing table headings.
    def label(table):
        decl=catalogue[table][0]['source'].rsplit(':',1)[0]
        return ('SQL_REPLICATE' if table.startswith('sql_replicate.') else 'Retail UAT')+'.'+pathlib.Path(decl).stem
    def page(table):return HERE/'tables'/(table+'.md')
    def declaration(r):
        base='all_object/bond/bond-sqlreplicate/db_object' if r['table'].startswith('sql_replicate.') else 'bond-sqltrsretail-uat/db_object'
        f,line=r['column_declaration'].rsplit(':',1)
        return WORKSPACE/base/f,int(line)
    table_summary=[];column_summary=[]
    for table,cols in catalogue.items():
        xs=groups[table];a=[r for r in xs if r['tier']=='A_COLUMN_PREDICATE'];b=[r for r in xs if r['tier']=='B_EXPRESSION_OR_BOUNDARY'];c=[r for r in xs if r['tier']=='C_FAMILY_OR_ROUTINE'];partners=[r for r in xs if r['tier']=='PARTNER_OR_DESTINATION']
        table_summary.append({'table':table,'display_table':label(table),'schema_columns':len(cols),'conversion_context_columns':len({r['column'] for r in xs}),'all_context_sites':nsites(xs),'all_context_procedures':nproc(xs),'a_sites':nsites(a),'a_procedures':nproc(a),'b_sites':nsites(b),'c_sites':nsites(c),'partner_destination_sites':nsites(partners),'status':'HAS_STATIC_CONVERSION_CONTEXT' if xs else 'NO_BOUND_MATCH_NOT_CERTIFIED','detail_file':page(table).relative_to(HERE).as_posix() if xs else ''})
        for col in sorted({r['column'] for r in xs}):
            cs=[r for r in xs if r['column']==col];aa=[r for r in cs if r['tier']=='A_COLUMN_PREDICATE']
            column_summary.append({'table':table,'column':col,'type':cs[0]['column_type'],'sites':nsites(cs),'procedures':nproc(cs),'a_sites':nsites(aa),'a_procedures':nproc(aa),'patterns':'|'.join(sorted({r['pattern_id'] for r in cs})),'roles':'|'.join(sorted({r['role'] for r in cs})),'other_types':'|'.join(sorted({r['other_type'] for r in cs})),'index_keys':cs[0]['existing_leading_indexes']})
    table_summary.sort(key=lambda r:(-r['a_procedures'],-r['a_sites'],-r['all_context_procedures'],r['table']))
    column_summary.sort(key=lambda r:(-r['a_procedures'],-r['a_sites'],-r['procedures'],r['table'],r['column']))
    csvout('table_summary.csv',table_summary);csvout('column_summary.csv',column_summary)
    psummary=[]
    for pid,(title,advice) in PATTERNS.items():
        ss=[r for r in sites if r['database']=='retail' and r['module_kind']=='procedures' and pattern(r)==pid]
        es=[r for r in rows if r['pattern_id']==pid];aa=[r for r in es if r['tier']=='A_COLUMN_PREDICATE']
        psummary.append({'pattern_id':pid,'name':title,'sites_all_operands':len(ss),'sites_with_physical_columns':nsites(es),'physical_tables':len({r['table'] for r in es}),'procedures_all_operands':len({r['file'] for r in ss}),'a_column_sites':nsites(aa),'advice':advice})
    csvout('pattern_summary.csv',psummary)
    tabledoc=HERE/'TABLE_COLUMN_SUMMARY.md'
    lines=['# Ringkasan seluruh tabel dan kolom: konversi implisit','',
           'Sumber 23 Oktober 2026. Unit frekuensi adalah **site pasangan operand/ordinal** dan **procedure berbeda**, bukan execution count. Satu statement dapat memiliki beberapa site konversi. Tabel sumber: 676 Retail UAT + 165 SQL_REPLICATE. Semua tabel tercatat pada `table_summary.csv`, termasuk yang tidak mempunyai match.','',
           'A = konversi pada kolom fisik dalam filter/join/operasi predicate yang berpotensi mengganggu akses index. B = expression/assignment/function boundary. C = perbedaan family rutin. Partner/destination dicatat untuk lineage dan **tidak berarti kolom tersebut ikut dikonversi**. Tidak ada runtime performance certification.','',
           '## Kolom yang paling perlu ditinjau untuk akses index','',
           'Tabel ini memuat **seluruh** kolom A, diurutkan jumlah procedure. Klik nama tabel untuk semua procedure dan query pada kolom tersebut.','',
           '| Tabel | Kolom (tipe source) | Lokasi A | SP A | Pola | Existing leading index |',
           '|---|---|---:|---:|---|---|']
    for c in column_summary:
        if not c['a_sites']:continue
        lines.append(f"| [{md(label(c['table']))}]({link(tabledoc,page(c['table']))}#kolom-{anchor(c['column'])}) | `{md(c['column'])}` (`{md(c['type'])}`) | {c['a_sites']} | {c['a_procedures']} | {', '.join(c['patterns'].split('|'))} | {md(c['index_keys'] or 'Tidak ada key pertama di snapshot')} |")
    lines += ['','## Semua tabel yang muncul dalam conversion context','',
              'Bagian ini juga mencakup destination INSERT/UPDATE dan kolom pasangan yang **tidak dikonversi**. Gunakan angka A untuk memprioritaskan risiko index; jangan mengurutkan pekerjaan hanya dari total semua context.','',
              '| Tabel | Kolom terkait | Seluruh lokasi | Seluruh SP | Lokasi A | SP A | Lokasi B | Lokasi C | Partner/destination |',
              '|---|---:|---:|---:|---:|---:|---:|---:|---:|']
    for r in table_summary:
        if r['all_context_sites']:lines.append(f"| [{md(r['display_table'])}]({link(tabledoc,page(r['table']))}) | {r['conversion_context_columns']} | {r['all_context_sites']} | {r['all_context_procedures']} | {r['a_sites']} | {r['a_procedures']} | {r['b_sites']} | {r['c_sites']} | {r['partner_destination_sites']} |")
    lines += ['','## Tabel tanpa bound match','',
              'Daftar lengkap berada di `table_summary.csv`, status NO_BOUND_MATCH_NOT_CERTIFIED. Ini mencakup tabel tanpa pemakaian terkait, tipe sama, maupun jalur yang belum berhasil di-bind; bukan bukti semua query terhadap tabel tersebut bebas conversion.','']
    tabledoc.write_text('\n'.join(lines),encoding='utf8')
    (HERE/'tables').mkdir(exist_ok=True)
    for table,xs in groups.items():
        if not xs:continue
        file=page(table);a=[r for r in xs if r['tier']=='A_COLUMN_PREDICATE'];b=[r for r in xs if r['tier']=='B_EXPRESSION_OR_BOUNDARY']
        content=[f'# {label(table)} — review konversi implisit','',f"[Kembali ke rangkuman]({link(file,tabledoc)}) · [Panduan pola]({link(file,HERE/'PATTERN_GUIDE.md')})",'',
                 f'**Cakupan procedure UAT:** {nsites(xs)} lokasi source, {nproc(xs)} procedure, {len({r["column"] for r in xs})} kolom. **A:** {nsites(a)} lokasi pada {nproc(a)} procedure. **B:** {nsites(b)} lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.','',
                 'Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.','']
        for col in sorted({r['column'] for r in xs},key=lambda col:(-sum(r['tier']=='A_COLUMN_PREDICATE' for r in xs if r['column']==col),col)):
            cs=[r for r in xs if r['column']==col];aa=[r for r in cs if r['tier']=='A_COLUMN_PREDICATE'];decl,ln=declaration(cs[0]);patterns=sorted({r['pattern_id'] for r in cs})
            content += [f'<a id="kolom-{anchor(col)}"></a>',f'## {col}','',
                        f"**Tipe dasar:** `{md(cs[0]['column_type'])}`. [Deklarasi kolom]({link(file,decl,ln)}). **Existing leading index:** {md(cs[0]['existing_leading_indexes'] or 'Tidak ada yang tercatat')}.",'',
                        f"**Frekuensi source:** {nsites(cs)} lokasi / {nproc(cs)} SP. **A:** {nsites(aa)} lokasi / {nproc(aa)} SP. **Peran:** {', '.join(sorted({r['role'] for r in cs}))}.",'',
                        '**Pola dan arah review:**','']
            for p in patterns:content.append(f'- **{p} — {PATTERNS[p][0]}:** {PATTERNS[p][1]}')
            content += ['','### Daftar procedure dan bagian query','',
                        '| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |',
                        '|---|---|---|---|---|---|---|---|']
            for r in sorted(cs,key=lambda r:(tierorder(r),r['file'],int(r['line']),r['site_id'],r['role'])):
                src=WORKSPACE/r['source_path'];pr=pathlib.Path(r['file']).stem
                content.append(f"| {r['site_id']} | {r['tier']} / {r['role']} | [{md(pr)}:{r['line']}]({link(file,src,int(r['line']))}) | {r['context']} / {r['pattern_id']} / {r['form']} | `{md(r['column_operand'])}` : `{md(r['operand_type'])}` | `{md(r['other_operand'])}` : `{md(r['other_type'])}` | `{md(r['converted_operand'])}` → `{md(r['target_type'])}` | `{md(r['predicate'])}` |")
            content += ['']
        file.write_text('\n'.join(content),encoding='utf8')
    # Exactly one row per UAT procedure, including those with no supported mismatch.
    pcoverage=[]
    for c in coverage:
        if c['database']!='retail' or not c['file'].startswith('procedures/'):continue
        ss=[r for r in sites if r['database']=='retail' and r['file']==c['file']];es=[r for r in rows if r['file']==c['file']]
        pcoverage.append(dict(c,typed_difference_sites=len(ss),physical_tables=len({r['table'] for r in es}),a_sites=nsites([r for r in es if r['tier']=='A_COLUMN_PREDICATE']),patterns='|'.join(sorted({pattern(r) for r in ss})),status='PARTIAL_PARSE_ERROR' if int(c['parse_errors']) else 'SCANNED_BOUNDED_STATIC_NOT_CERTIFIED'))
    csvout('procedure_coverage.csv',pcoverage)
    auxiliary=HERE/'SUPPORTING_MODULES.md';aux=['# Function dan view pendukung','',
        'Temuan pada body function/view dicatat terpisah dari lokasi di procedure. Pemanggilan UDF langsung dengan argumen kolom ada dalam daftar procedure; pemanggilan transitif tidak dianggap otomatis mengeksekusi setiap branch function. Semua detail tetap tersedia di conversion_sites.csv.','',
        '| Database / module | Lokasi beda tipe | Context / form |', '|---|---:|---|']
    mg=collections.defaultdict(list)
    for r in sites:
        if r['module_kind']!='procedures':mg[(r['database'],r['file'])].append(r)
    for (db,f),ss in sorted(mg.items()):aux.append(f"| {db} / [{md(pathlib.Path(f).stem)}]({link(auxiliary,WORKSPACE/ss[0]['source_path'])}) | {len(ss)} | {', '.join(sorted({r['form'] for r in ss}))} |")
    auxiliary.write_text('\n'.join(aux)+'\n',encoding='utf8')
    print(json.dumps({'table_documents':sum(bool(v) for v in groups.values()),'table_inventory':len(table_summary),'column_summary':len(column_summary),'procedure_coverage':len(pcoverage),'pattern_groups':len(PATTERNS)},indent=2))

if __name__=='__main__':main()
