package engine

import (
	"fmt"

	"rza_migrate_db/internal/output"
	"rza_migrate_db/internal/verify"
)

func (e Engine) verifyCallbacks(cache *genPrecheckHistoryCache) verify.DomainCallbacks {
	return verify.DomainCallbacks{
		Migrations: verify.MigrationsCallbacks{
			ValidateMigrations: func(r *output.Report) (any, error) {
				metas, err := e.validateMigrations(r)
				if err != nil {
					return nil, err
				}
				return metas, nil
			},
			ReconstructHistory: func(r *output.Report) (any, error) {
				historyMap, existingFiles, maxVersion, err := e.reconstructHistory(r)
				if err != nil {
					return nil, err
				}
				if cache != nil && r.Fail.Count == 0 {
					cache.Ready = true
					cache.History = cloneHistoryStateMap(historyMap)
					cache.ExistingFiles = cloneStringSet(existingFiles)
					cache.MaxVersion = maxVersion
				}
				return historyMap, nil
			},
			ValidateCoreAuditContracts: e.validateCoreAuditContracts,
			ValidateForbiddenSQL:       e.validateForbiddenSQL,
		},
		Patch: verify.PatchCallbacks{
			ValidatePatchFolders:         e.validatePatchFolders,
			ValidatePatchSourceIntegrity: e.verifyValidatePatchSourceIntegrity,
		},
		DbObject: verify.DbObjectCallbacks{
			ValidateTableInlineDefault:  e.validateTableInlineDefault,
			ValidateRepoHygiene:         e.validateRepoHygiene,
			ValidateDbObjectConsistency: e.validateDbObjectConsistency,
			CollectDesiredObjects: func(r *output.Report) (any, error) {
				desiredObjects, err := e.collectDesiredObjects(r)
				if err != nil {
					return nil, err
				}
				if cache != nil && r.Fail.Count == 0 {
					cache.DesiredReady = true
					cache.Desired = cloneDesiredObjects(desiredObjects)
				}
				return desiredObjects, nil
			},
			ValidateDesiredHistoryDrift:    e.verifyValidateDesiredHistoryDrift,
			ValidateProgrammabilityHeaders: e.verifyValidateProgrammabilityHeaders,
		},
		Wrapper: verify.WrapperCallbacks{
			ValidateWrapperSubset: e.verifyValidateWrapperSubset,
		},
	}
}

func (e Engine) verifyValidatePatchSourceIntegrity(report *output.Report, metasAny any) error {
	metas, err := verifyMetas(metasAny)
	if err != nil {
		return err
	}
	return e.validatePatchSourceIntegrity(report, metas)
}

func (e Engine) verifyValidateDesiredHistoryDrift(report *output.Report, historyAny, desiredAny any) error {
	historyMap, err := verifyHistoryMap(historyAny)
	if err != nil {
		return err
	}
	desiredObjects, err := verifyDesiredObjects(desiredAny)
	if err != nil {
		return err
	}
	return e.validateDesiredHistoryDrift(report, historyMap, desiredObjects)
}

func (e Engine) verifyValidateProgrammabilityHeaders(report *output.Report, metasAny any) error {
	metas, err := verifyMetas(metasAny)
	if err != nil {
		return err
	}
	return e.validateProgrammabilityHeaders(report, metas)
}

func (e Engine) verifyValidateWrapperSubset(report *output.Report, metasAny any) error {
	metas, err := verifyMetas(metasAny)
	if err != nil {
		return err
	}
	return e.validateWrapperSubset(report, metas)
}

func verifyMetas(input any) ([]migrationMeta, error) {
	if input == nil {
		return []migrationMeta{}, nil
	}
	metas, ok := input.([]migrationMeta)
	if !ok {
		return nil, fmt.Errorf("verify bridge: invalid metas payload type %T", input)
	}
	return metas, nil
}

func verifyHistoryMap(input any) (map[string]historyObject, error) {
	if input == nil {
		return map[string]historyObject{}, nil
	}
	historyMap, ok := input.(map[string]historyObject)
	if !ok {
		return nil, fmt.Errorf("verify bridge: invalid history payload type %T", input)
	}
	return historyMap, nil
}

func verifyDesiredObjects(input any) ([]desiredObject, error) {
	desiredObjects, ok := input.([]desiredObject)
	if !ok {
		return nil, fmt.Errorf("verify bridge: invalid desired payload type %T", input)
	}
	return desiredObjects, nil
}
