package verify

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"rza_migrate_db/internal/output"
)

func TestOrchestratorRun_OrderByDomainWithWrapper(t *testing.T) {
	report := output.NewReport("p", "lf")
	sequence := []string{}

	cbs := DomainCallbacks{
		Migrations: MigrationsCallbacks{
			ValidateMigrations: func(r *output.Report) (any, error) {
				sequence = append(sequence, "migrations.validate")
				return "metas", nil
			},
			ReconstructHistory: func(r *output.Report) (any, error) {
				sequence = append(sequence, "migrations.reconstruct")
				return "history", nil
			},
			ValidateCoreAuditContracts: func(r *output.Report) error {
				sequence = append(sequence, "migrations.core_audit")
				return nil
			},
			ValidateForbiddenSQL: func(r *output.Report) error {
				sequence = append(sequence, "migrations.forbidden_sql")
				return nil
			},
		},
		Patch: PatchCallbacks{
			ValidatePatchFolders: func(r *output.Report) error {
				sequence = append(sequence, "patch.folders")
				return nil
			},
			ValidatePatchSourceIntegrity: func(r *output.Report, metas any) error {
				sequence = append(sequence, fmt.Sprintf("patch.integrity:%v", metas))
				return nil
			},
		},
		DbObject: DbObjectCallbacks{
			ValidateTableInlineDefault: func(r *output.Report) error {
				sequence = append(sequence, "db.inline_default")
				return nil
			},
			ValidateRepoHygiene: func(r *output.Report) error {
				sequence = append(sequence, "db.repo_hygiene")
				return nil
			},
			ValidateDbObjectConsistency: func(r *output.Report) error {
				sequence = append(sequence, "db.consistency")
				return nil
			},
			CollectDesiredObjects: func(r *output.Report) (any, error) {
				sequence = append(sequence, "db.collect_desired")
				return "desired", nil
			},
			ValidateDesiredHistoryDrift: func(r *output.Report, history any, desired any) error {
				sequence = append(sequence, fmt.Sprintf("db.drift:%v:%v", history, desired))
				return nil
			},
			ValidateProgrammabilityHeaders: func(r *output.Report, metas any) error {
				sequence = append(sequence, fmt.Sprintf("db.programmability:%v", metas))
				return nil
			},
		},
		Wrapper: WrapperCallbacks{
			ValidateWrapperSubset: func(r *output.Report, metas any) error {
				sequence = append(sequence, fmt.Sprintf("wrapper.validate:%v", metas))
				return nil
			},
		},
	}

	err := Orchestrator{IncludeWrapper: true}.Run(&report, cbs)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	want := []string{
		"migrations.validate",
		"migrations.reconstruct",
		"migrations.core_audit",
		"migrations.forbidden_sql",
		"patch.folders",
		"patch.integrity:metas",
		"db.inline_default",
		"db.repo_hygiene",
		"db.consistency",
		"db.collect_desired",
		"db.drift:history:desired",
		"db.programmability:metas",
		"wrapper.validate:metas",
	}
	if !reflect.DeepEqual(sequence, want) {
		t.Fatalf("unexpected sequence:\nwant: %#v\ngot:  %#v", want, sequence)
	}
}

func TestOrchestratorRun_SkipReconstructWhenFailExists(t *testing.T) {
	report := output.NewReport("p", "lf")
	sequence := []string{}
	cbs := fullNoopDomainCallbacks()
	cbs.Migrations.ValidateMigrations = func(r *output.Report) (any, error) {
		sequence = append(sequence, "migrations.validate")
		r.AddFail(output.Item{
			Rule:     "X",
			Severity: 3,
			File:     ".",
			Message:  "x",
			Reason:   "REQUIRES_MANUAL",
			Action:   "FAIL",
		})
		return "metas", nil
	}
	cbs.Migrations.ReconstructHistory = func(r *output.Report) (any, error) {
		sequence = append(sequence, "migrations.reconstruct")
		return "history", nil
	}

	err := Orchestrator{IncludeWrapper: true}.Run(&report, cbs)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	if len(sequence) != 1 || sequence[0] != "migrations.validate" {
		t.Fatalf("reconstruct should be skipped when fail count > 0, sequence=%v", sequence)
	}
}

func TestOrchestratorRun_NoWrapperWhenDisabled(t *testing.T) {
	report := output.NewReport("p", "lf")
	wrapperCalls := 0
	cbs := fullNoopDomainCallbacks()
	cbs.Wrapper.ValidateWrapperSubset = func(r *output.Report, metas any) error {
		wrapperCalls++
		return nil
	}

	err := Orchestrator{IncludeWrapper: false}.Run(&report, cbs)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	if wrapperCalls != 0 {
		t.Fatalf("wrapper should not be called when include_wrapper=false")
	}
}

func TestOrchestratorRun_MissingCallback(t *testing.T) {
	report := output.NewReport("p", "lf")
	cbs := fullNoopDomainCallbacks()
	cbs.Patch.ValidatePatchFolders = nil

	err := Orchestrator{IncludeWrapper: true}.Run(&report, cbs)
	if err == nil {
		t.Fatalf("expected missing callback error")
	}
}

func TestOrchestratorRun_FailFastStopsAfterFirstFail(t *testing.T) {
	report := output.NewReport("p", "lf")
	sequence := []string{}

	cbs := fullNoopDomainCallbacks()
	cbs.Migrations.ValidateMigrations = func(r *output.Report) (any, error) {
		sequence = append(sequence, "migrations.validate")
		r.AddFail(output.Item{
			Rule:     "X",
			Severity: 3,
			File:     ".",
			Message:  "x",
			Reason:   "REQUIRES_MANUAL",
			Action:   "FAIL",
		})
		return "metas", nil
	}
	cbs.Migrations.ValidateCoreAuditContracts = func(r *output.Report) error {
		sequence = append(sequence, "migrations.core_audit")
		return nil
	}
	cbs.Migrations.ValidateForbiddenSQL = func(r *output.Report) error {
		sequence = append(sequence, "migrations.forbidden_sql")
		return nil
	}
	cbs.Patch.ValidatePatchFolders = func(r *output.Report) error {
		sequence = append(sequence, "patch.folders")
		return nil
	}
	cbs.DbObject.ValidateRepoHygiene = func(r *output.Report) error {
		sequence = append(sequence, "db.repo_hygiene")
		return nil
	}
	cbs.Wrapper.ValidateWrapperSubset = func(r *output.Report, metas any) error {
		sequence = append(sequence, "wrapper.validate")
		return nil
	}

	err := Orchestrator{IncludeWrapper: true, FailFast: true}.Run(&report, cbs)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	want := []string{"migrations.validate"}
	if !reflect.DeepEqual(sequence, want) {
		t.Fatalf("unexpected sequence in fail-fast mode:\nwant: %#v\ngot:  %#v", want, sequence)
	}
}

func TestOrchestratorRun_ParallelMigrationsChecksDeterministicMergeOrder(t *testing.T) {
	report := output.NewReport("p", "lf")

	cbs := fullNoopDomainCallbacks()
	cbs.Migrations.ValidateCoreAuditContracts = func(r *output.Report) error {
		time.Sleep(30 * time.Millisecond)
		r.AddFail(output.Item{
			Rule:     "CORE_FAIL",
			Severity: 3,
			File:     "core.sql",
			Message:  "core",
			Reason:   "REQUIRES_MANUAL",
			Action:   "FAIL",
		})
		return nil
	}
	cbs.Migrations.ValidateForbiddenSQL = func(r *output.Report) error {
		r.AddFail(output.Item{
			Rule:     "FORBIDDEN_FAIL",
			Severity: 3,
			File:     "forbidden.sql",
			Message:  "forbidden",
			Reason:   "REQUIRES_MANUAL",
			Action:   "FAIL",
		})
		return nil
	}

	err := Orchestrator{
		IncludeWrapper:    false,
		EnableParallelRun: true,
		ParallelWorkers:   2,
	}.Run(&report, cbs)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	if len(report.Fail.Items) < 2 {
		t.Fatalf("expected at least 2 fail items, got %d", len(report.Fail.Items))
	}
	if report.Fail.Items[0].Rule != "CORE_FAIL" {
		t.Fatalf("expected first fail from core merge order, got %s", report.Fail.Items[0].Rule)
	}
	if report.Fail.Items[1].Rule != "FORBIDDEN_FAIL" {
		t.Fatalf("expected second fail from forbidden merge order, got %s", report.Fail.Items[1].Rule)
	}
}

func fullNoopDomainCallbacks() DomainCallbacks {
	return DomainCallbacks{
		Migrations: MigrationsCallbacks{
			ValidateMigrations:         func(r *output.Report) (any, error) { return "metas", nil },
			ReconstructHistory:         func(r *output.Report) (any, error) { return "history", nil },
			ValidateCoreAuditContracts: func(r *output.Report) error { return nil },
			ValidateForbiddenSQL:       func(r *output.Report) error { return nil },
		},
		Patch: PatchCallbacks{
			ValidatePatchFolders:         func(r *output.Report) error { return nil },
			ValidatePatchSourceIntegrity: func(r *output.Report, metas any) error { return nil },
		},
		DbObject: DbObjectCallbacks{
			ValidateTableInlineDefault:     func(r *output.Report) error { return nil },
			ValidateRepoHygiene:            func(r *output.Report) error { return nil },
			ValidateDbObjectConsistency:    func(r *output.Report) error { return nil },
			CollectDesiredObjects:          func(r *output.Report) (any, error) { return "desired", nil },
			ValidateDesiredHistoryDrift:    func(r *output.Report, history any, desired any) error { return nil },
			ValidateProgrammabilityHeaders: func(r *output.Report, metas any) error { return nil },
		},
		Wrapper: WrapperCallbacks{
			ValidateWrapperSubset: func(r *output.Report, metas any) error { return nil },
		},
	}
}
