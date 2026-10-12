# Constraint Naming Guideline

Dokumen ini adalah guideline penamaan constraint pada `db_object/tables/*.sql`.

## Tujuan

- Menjaga konsistensi nama constraint antar tim/environment.
- Mempermudah diff, rollback, dan troubleshooting migration.

## Pola Nama Rekomendasi

- Primary Key: `PK_<Table>`
- Default: `DF_<Table>__<Column>`
- Unique: `UQ_<Table>__<ColumnOrGroup>`
- Check: `CK_<Table>__<Rule>`
- Foreign Key: `FK_<Table>__<RefTable>__<ColumnOrGroup>`

## Aturan Umum (Rekomendasi)

- Disarankan menggunakan keyword `CONSTRAINT <Name> ...`.
- Disarankan nama constraint memuat token nama tabel.
- Hindari nama random/non-deterministik bila memungkinkan.
- Jika rename constraint dilakukan, dokumentasikan di migration.
- `rza_migrate_report` memperlakukan drop+add dengan nama berbeda sebagai metadata-only hanya jika definisi constraint identik secara semantik. Migration sumber tetap immutable dan tetap terlihat di `migration_summary.csv`; perubahan expression, column/order, reference, atau atribut lain tetap dilaporkan sebagai operasi semantik.
- Drop+add dengan nama yang sama dan definisi berubah dilaporkan sebagai satu `ALTER_CONSTRAINT`/`ALTER_DEFAULT` yang menyimpan definisi lama dan baru, bukan dua operasi mekanis terpisah.

## Contoh

```sql
CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  status VARCHAR(10) NULL,
  CONSTRAINT PK_TableA PRIMARY KEY (id),
  CONSTRAINT DF_TableA__status DEFAULT ('NEW') FOR status
);
```
