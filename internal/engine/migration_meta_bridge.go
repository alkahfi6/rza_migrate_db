package engine

import "rza_migrate_db/internal/verify"

func toVerifyMigrationMeta(meta migrationMeta) verify.MigrationMeta {
	return verify.MigrationMeta{
		FileName:  meta.FileName,
		RelPath:   meta.RelPath,
		Version:   meta.Version,
		Category:  meta.Category,
		Object:    meta.Object,
		Direction: meta.Direction,
		VK:        meta.VK,
		IsCore:    meta.IsCore,
	}
}

func toVerifyMigrationMetas(metas []migrationMeta) []verify.MigrationMeta {
	converted := make([]verify.MigrationMeta, 0, len(metas))
	for _, meta := range metas {
		converted = append(converted, toVerifyMigrationMeta(meta))
	}
	return converted
}

func toEngineMigrationMeta(meta verify.MigrationMeta) migrationMeta {
	return migrationMeta{
		FileName:  meta.FileName,
		RelPath:   meta.RelPath,
		Version:   meta.Version,
		Category:  meta.Category,
		Object:    meta.Object,
		Direction: meta.Direction,
		VK:        meta.VK,
		IsCore:    meta.IsCore,
	}
}

func toEngineMigrationMetas(metas []verify.MigrationMeta) []migrationMeta {
	converted := make([]migrationMeta, 0, len(metas))
	for _, meta := range metas {
		converted = append(converted, toEngineMigrationMeta(meta))
	}
	return converted
}
