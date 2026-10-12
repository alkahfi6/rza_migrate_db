package verify

import (
	"fmt"
	"reflect"
	"sync"
	"time"

	"rza_migrate_db/internal/output"
)

// Orchestrator drives validation execution order.
type Orchestrator struct {
	IncludeWrapper      bool
	MigrationValidation string
	FailFast            bool
	EnableParallelRun   bool
	ParallelWorkers     int
	EnablePerfNotes     bool
}

// MigrationsCallbacks binds validations for migrations-domain.
type MigrationsCallbacks struct {
	ValidateMigrations         func(*output.Report) (any, error)
	ReconstructHistory         func(*output.Report) (any, error)
	ValidateCoreAuditContracts func(*output.Report) error
	ValidateForbiddenSQL       func(*output.Report) error
}

// PatchCallbacks binds validations for patch-domain.
type PatchCallbacks struct {
	ValidatePatchFolders         func(*output.Report) error
	ValidatePatchSourceIntegrity func(*output.Report, any) error
}

// DbObjectCallbacks binds validations for db_object-domain.
type DbObjectCallbacks struct {
	ValidateTableInlineDefault     func(*output.Report) error
	ValidateRepoHygiene            func(*output.Report) error
	ValidateDbObjectConsistency    func(*output.Report) error
	CollectDesiredObjects          func(*output.Report) (any, error)
	ValidateDesiredHistoryDrift    func(*output.Report, any, any) error
	ValidateProgrammabilityHeaders func(*output.Report, any) error
}

// WrapperCallbacks binds validations for wrapper-domain.
type WrapperCallbacks struct {
	ValidateWrapperSubset func(*output.Report, any) error
}

// DomainCallbacks groups callback contracts per validation domain.
type DomainCallbacks struct {
	Migrations MigrationsCallbacks
	Patch      PatchCallbacks
	DbObject   DbObjectCallbacks
	Wrapper    WrapperCallbacks
}

// Run executes verify flow deterministically and exits on first runtime error.
func (o Orchestrator) Run(report *output.Report, cbs DomainCallbacks) error {
	if err := validateCallbacks(cbs, o.IncludeWrapper); err != nil {
		return err
	}

	migrationsStarted := time.Now()
	metas, history, err := runMigrationsDomain(report, cbs.Migrations, o)
	recordPerf(report, o.EnablePerfNotes, "verify.domain.migrations", migrationsStarted)
	if err != nil {
		return err
	}
	if o.FailFast && report.Fail.Count > 0 {
		return nil
	}

	patchStarted := time.Now()
	if err := runPatchDomain(report, cbs.Patch, metas, o); err != nil {
		recordPerf(report, o.EnablePerfNotes, "verify.domain.patch", patchStarted)
		return err
	}
	recordPerf(report, o.EnablePerfNotes, "verify.domain.patch", patchStarted)
	if o.FailFast && report.Fail.Count > 0 {
		return nil
	}

	dbObjectStarted := time.Now()
	if err := runDbObjectDomain(report, cbs.DbObject, history, metas, o); err != nil {
		recordPerf(report, o.EnablePerfNotes, "verify.domain.db_object", dbObjectStarted)
		return err
	}
	recordPerf(report, o.EnablePerfNotes, "verify.domain.db_object", dbObjectStarted)
	if o.FailFast && report.Fail.Count > 0 {
		return nil
	}

	if o.IncludeWrapper {
		wrapperStarted := time.Now()
		if err := runWrapperDomain(report, cbs.Wrapper, metas); err != nil {
			recordPerf(report, o.EnablePerfNotes, "verify.domain.wrapper", wrapperStarted)
			return err
		}
		recordPerf(report, o.EnablePerfNotes, "verify.domain.wrapper", wrapperStarted)
	}
	return nil
}

func recordPerf(report *output.Report, enabled bool, key string, started time.Time) {
	if report == nil || !enabled {
		return
	}
	entry, ok := report.Notes["perf_ms"]
	if !ok {
		perf := map[string]interface{}{
			key: time.Since(started).Milliseconds(),
		}
		report.Notes["perf_ms"] = perf
		return
	}
	perf, ok := entry.(map[string]interface{})
	if !ok {
		perf = map[string]interface{}{}
		report.Notes["perf_ms"] = perf
	}
	perf[key] = time.Since(started).Milliseconds()
}

func runMigrationsDomain(report *output.Report, cbs MigrationsCallbacks, opts Orchestrator) (any, any, error) {
	mode := effectiveMigrationValidation(opts.MigrationValidation)
	if mode == "off" {
		return nil, nil, nil
	}
	metas, err := cbs.ValidateMigrations(report)
	if err != nil {
		return nil, nil, err
	}
	if opts.FailFast && report.Fail.Count > 0 {
		return metas, nil, nil
	}
	if mode == "integrity" {
		return metas, nil, nil
	}

	var history any
	if report.Fail.Count == 0 {
		history, err = cbs.ReconstructHistory(report)
		if err != nil {
			return nil, nil, err
		}
	}
	if opts.FailFast && report.Fail.Count > 0 {
		return metas, history, nil
	}

	if !opts.EnableParallelRun {
		if err := cbs.ValidateCoreAuditContracts(report); err != nil {
			return nil, nil, err
		}
		if opts.FailFast && report.Fail.Count > 0 {
			return metas, history, nil
		}
		if err := cbs.ValidateForbiddenSQL(report); err != nil {
			return nil, nil, err
		}
		return metas, history, nil
	}

	type callbackTask struct {
		name string
		fn   func(*output.Report) error
	}
	type callbackResult struct {
		name   string
		report output.Report
		err    error
	}

	tasks := []callbackTask{
		{
			name: "core_audit",
			fn:   cbs.ValidateCoreAuditContracts,
		},
		{
			name: "forbidden_sql",
			fn:   cbs.ValidateForbiddenSQL,
		},
	}

	results := make(map[string]callbackResult, len(tasks))
	workers := opts.ParallelWorkers
	if workers < 1 {
		workers = 1
	}
	if workers > len(tasks) {
		workers = len(tasks)
	}
	if workers <= 1 {
		for _, task := range tasks {
			sub := output.NewReport(report.NewMigrations.Naming.Pattern, report.NewMigrations.Naming.LineEndings)
			err := task.fn(&sub)
			results[task.name] = callbackResult{name: task.name, report: sub, err: err}
		}
	} else {
		jobs := make(chan callbackTask)
		out := make(chan callbackResult, len(tasks))
		var wg sync.WaitGroup
		for i := 0; i < workers; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for task := range jobs {
					sub := output.NewReport(report.NewMigrations.Naming.Pattern, report.NewMigrations.Naming.LineEndings)
					err := task.fn(&sub)
					out <- callbackResult{name: task.name, report: sub, err: err}
				}
			}()
		}
		for _, task := range tasks {
			jobs <- task
		}
		close(jobs)
		wg.Wait()
		close(out)
		for res := range out {
			results[res.name] = res
		}
	}

	mergeOrder := []string{"core_audit", "forbidden_sql"}
	for _, name := range mergeOrder {
		res, ok := results[name]
		if !ok {
			return nil, nil, fmt.Errorf("verify orchestrator internal: missing result for %s", name)
		}
		mergeReport(report, res.report)
		if res.err != nil {
			return nil, nil, res.err
		}
		if opts.FailFast && report.Fail.Count > 0 {
			return metas, history, nil
		}
	}
	return metas, history, nil
}

func runPatchDomain(report *output.Report, cbs PatchCallbacks, metas any, opts Orchestrator) error {
	if err := cbs.ValidatePatchFolders(report); err != nil {
		return err
	}
	if opts.FailFast && report.Fail.Count > 0 {
		return nil
	}
	if effectiveMigrationValidation(opts.MigrationValidation) == "off" {
		return nil
	}
	if err := cbs.ValidatePatchSourceIntegrity(report, metas); err != nil {
		return err
	}
	return nil
}

func runDbObjectDomain(report *output.Report, cbs DbObjectCallbacks, history, metas any, opts Orchestrator) error {
	if opts.EnableParallelRun && !opts.FailFast {
		return runDbObjectDomainParallel(report, cbs, history, metas, opts)
	}
	if err := cbs.ValidateTableInlineDefault(report); err != nil {
		return err
	}
	if opts.FailFast && report.Fail.Count > 0 {
		return nil
	}
	if err := cbs.ValidateRepoHygiene(report); err != nil {
		return err
	}
	if opts.FailFast && report.Fail.Count > 0 {
		return nil
	}
	if err := cbs.ValidateDbObjectConsistency(report); err != nil {
		return err
	}
	if opts.FailFast && report.Fail.Count > 0 {
		return nil
	}

	desiredObjects, err := cbs.CollectDesiredObjects(report)
	if err != nil {
		return err
	}
	if opts.FailFast && report.Fail.Count > 0 {
		return nil
	}
	if effectiveMigrationValidation(opts.MigrationValidation) == "full" {
		if err := cbs.ValidateDesiredHistoryDrift(report, history, desiredObjects); err != nil {
			return err
		}
		if opts.FailFast && report.Fail.Count > 0 {
			return nil
		}
	}
	if err := cbs.ValidateProgrammabilityHeaders(report, metas); err != nil {
		return err
	}
	return nil
}

func runDbObjectDomainParallel(report *output.Report, cbs DbObjectCallbacks, history, metas any, opts Orchestrator) error {
	type callbackTask struct {
		name string
		fn   func(*output.Report) error
	}
	type callbackResult struct {
		report output.Report
		err    error
	}
	tasks := []callbackTask{
		{name: "inline_default", fn: cbs.ValidateTableInlineDefault},
		{name: "repo_hygiene", fn: cbs.ValidateRepoHygiene},
		{name: "consistency", fn: cbs.ValidateDbObjectConsistency},
		{name: "headers", fn: func(r *output.Report) error { return cbs.ValidateProgrammabilityHeaders(r, metas) }},
	}
	results := make(map[string]callbackResult, len(tasks))
	workers := opts.ParallelWorkers
	if workers < 1 {
		workers = 1
	}
	if workers > len(tasks) {
		workers = len(tasks)
	}
	jobs := make(chan callbackTask)
	out := make(chan struct {
		name string
		res  callbackResult
	}, len(tasks))
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range jobs {
				sub := output.NewReport("", "")
				out <- struct {
					name string
					res  callbackResult
				}{task.name, callbackResult{report: sub, err: task.fn(&sub)}}
			}
		}()
	}
	for _, task := range tasks {
		jobs <- task
	}
	close(jobs)
	wg.Wait()
	close(out)
	for item := range out {
		results[item.name] = item.res
	}

	for _, name := range []string{"inline_default", "repo_hygiene", "consistency"} {
		res := results[name]
		mergeReport(report, res.report)
		if res.err != nil {
			return res.err
		}
	}
	desiredObjects, err := cbs.CollectDesiredObjects(report)
	if err != nil {
		return err
	}
	if effectiveMigrationValidation(opts.MigrationValidation) == "full" {
		if err := cbs.ValidateDesiredHistoryDrift(report, history, desiredObjects); err != nil {
			return err
		}
	}
	headerResult := results["headers"]
	mergeReport(report, headerResult.report)
	return headerResult.err
}

func effectiveMigrationValidation(value string) string {
	switch value {
	case "off", "integrity", "full":
		return value
	default:
		return "full"
	}
}

func runWrapperDomain(report *output.Report, cbs WrapperCallbacks, metas any) error {
	return cbs.ValidateWrapperSubset(report, metas)
}

func validateCallbacks(cbs DomainCallbacks, includeWrapper bool) error {
	required := []struct {
		name string
		fn   any
	}{
		{name: "Migrations.ValidateMigrations", fn: cbs.Migrations.ValidateMigrations},
		{name: "Migrations.ReconstructHistory", fn: cbs.Migrations.ReconstructHistory},
		{name: "Migrations.ValidateCoreAuditContracts", fn: cbs.Migrations.ValidateCoreAuditContracts},
		{name: "Migrations.ValidateForbiddenSQL", fn: cbs.Migrations.ValidateForbiddenSQL},
		{name: "Patch.ValidatePatchFolders", fn: cbs.Patch.ValidatePatchFolders},
		{name: "Patch.ValidatePatchSourceIntegrity", fn: cbs.Patch.ValidatePatchSourceIntegrity},
		{name: "DbObject.ValidateTableInlineDefault", fn: cbs.DbObject.ValidateTableInlineDefault},
		{name: "DbObject.ValidateRepoHygiene", fn: cbs.DbObject.ValidateRepoHygiene},
		{name: "DbObject.ValidateDbObjectConsistency", fn: cbs.DbObject.ValidateDbObjectConsistency},
		{name: "DbObject.CollectDesiredObjects", fn: cbs.DbObject.CollectDesiredObjects},
		{name: "DbObject.ValidateDesiredHistoryDrift", fn: cbs.DbObject.ValidateDesiredHistoryDrift},
		{name: "DbObject.ValidateProgrammabilityHeaders", fn: cbs.DbObject.ValidateProgrammabilityHeaders},
	}
	if includeWrapper {
		required = append(required, struct {
			name string
			fn   any
		}{
			name: "Wrapper.ValidateWrapperSubset",
			fn:   cbs.Wrapper.ValidateWrapperSubset,
		})
	}

	for _, item := range required {
		if item.fn == nil {
			return fmt.Errorf("verify orchestrator callback %s is required", item.name)
		}
		value := reflect.ValueOf(item.fn)
		if value.Kind() == reflect.Func && value.IsNil() {
			return fmt.Errorf("verify orchestrator callback %s is required", item.name)
		}
	}
	return nil
}

func mergeReport(dst *output.Report, src output.Report) {
	for _, item := range src.Fail.Items {
		dst.AddFail(item)
	}
	for _, item := range src.Pass.Items {
		dst.AddPass(item)
	}
}
