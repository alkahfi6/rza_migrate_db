package engine_test

import (
	"bytes"
	"crypto/sha256"
	"debug/buildinfo"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestMigrateCheckMRGateContract(t *testing.T) {
	t.Parallel()
	repo := repoRoot(t)

	ciRaw, err := os.ReadFile(filepath.Join(repo, ".gitlab-ci.yml"))
	if err != nil {
		t.Fatalf("read .gitlab-ci.yml: %v", err)
	}
	var ci map[string]interface{}
	if err := yaml.Unmarshal(ciRaw, &ci); err != nil {
		t.Fatalf("parse .gitlab-ci.yml: %v", err)
	}
	for _, forbidden := range []string{"stages", "workflow", "default", "rza_migrate_check_gate"} {
		if _, exists := ci[forbidden]; exists {
			t.Fatalf("root CI must not override central template key %q", forbidden)
		}
	}
	variables, ok := ci["variables"].(map[string]interface{})
	if !ok {
		t.Fatalf("office variables are missing: %#v", ci["variables"])
	}
	if variables["PROJECT_NAME"] != "wealth" || variables["SERVICE_NAME"] != "db-reksa-sqlreksa2" || variables["ALLOWED_BUILD"] != "/(sit|uat|tags)/g" {
		t.Fatalf("office pipeline variables changed unexpectedly: %#v", variables)
	}
	includes, ok := ci["include"].([]interface{})
	if !ok || len(includes) != 3 {
		t.Fatalf("office pipeline must have two central and one isolated checker include: %#v", ci["include"])
	}
	localFound := false
	centralFiles := map[string]bool{"/main.yaml": false, "/templates/golang.yaml": false}
	for _, rawInclude := range includes {
		include, ok := rawInclude.(map[string]interface{})
		if !ok {
			t.Fatalf("invalid include entry: %#v", rawInclude)
		}
		if include["local"] == "/.gitlab/rza_migrate_check.gitlab-ci.yml" {
			localFound = true
		}
		if include["project"] == "ocbc/devops/pipeline" {
			if file, ok := include["file"].(string); ok {
				centralFiles[file] = true
			}
		}
	}
	if !localFound || !centralFiles["/main.yaml"] || !centralFiles["/templates/golang.yaml"] {
		t.Fatalf("pipeline include contract mismatch: checker=%t central=%#v", localFound, centralFiles)
	}
	for _, jobName := range []string{"docker-release-dev", "docker-release-dc1", "docker-release-dc2", "helm-deploy-dev", "helm-deploy-dc1", "helm-deploy-dc2"} {
		if _, ok := ci[jobName].(map[string]interface{}); !ok {
			t.Fatalf("office job %s is missing", jobName)
		}
	}

	gateRaw, err := os.ReadFile(filepath.Join(repo, ".gitlab", "rza_migrate_check.gitlab-ci.yml"))
	if err != nil {
		t.Fatalf("read local checker include: %v", err)
	}
	var gateConfig map[string]interface{}
	if err := yaml.Unmarshal(gateRaw, &gateConfig); err != nil {
		t.Fatalf("parse local checker include: %v", err)
	}
	if len(gateConfig) != 1 {
		t.Fatalf("local checker include must define exactly one job: %#v", gateConfig)
	}
	job, ok := gateConfig["rza_migrate_check_gate"].(map[string]interface{})
	if !ok {
		t.Fatalf("rza_migrate_check_gate is missing or invalid: %#v", gateConfig["rza_migrate_check_gate"])
	}
	if job["stage"] != ".pre" || job["allow_failure"] != false {
		t.Fatalf("MR checker must be blocking: %#v", job)
	}
	if job["image"] != "nexus.appsnp.ocbcnisp.com:5004/alpine:3.22" || job["interruptible"] != true || job["timeout"] != "10m" {
		t.Fatalf("MR checker runtime contract mismatch: %#v", job)
	}
	retry, ok := job["retry"].(map[string]interface{})
	if !ok || retry["max"] != 2 {
		t.Fatalf("MR checker infrastructure retry contract mismatch: %#v", job["retry"])
	}
	retryWhen := yamlStrings(t, retry["when"], "rza_migrate_check_gate.retry.when")
	if strings.Join(retryWhen, ",") != "runner_system_failure,stuck_or_timeout_failure" {
		t.Fatalf("MR checker must retry infrastructure failures only: %#v", retryWhen)
	}
	script := strings.Join(yamlStrings(t, job["script"], "rza_migrate_check_gate.script"), "\n")
	for _, required := range []string{
		"test \"$(uname -s)\" = \"Linux\"",
		"test \"$(uname -m)\" = \"x86_64\"",
		"test -f tools/license.key",
		"sha256sum -c tools/linux-amd64/release.sha256",
		"sha256sum -c tools/linux-amd64/rza_migrate_check.sha256",
		"test \"$(tools/linux-amd64/rza_migrate_check -build-info)\" = \"$(cat tools/linux-amd64/build-info.txt)\"",
		"sh scripts/verify_release_source.sh .",
		"tools/linux-amd64/rza_migrate_check -repo . -check -profile authoring -migration-validation integrity -pause-on-exit false",
		"test -s out/output.json",
	} {
		if !strings.Contains(script, required) {
			t.Fatalf("CI checker script is missing %q:\n%s", required, script)
		}
	}
	beforeScript := strings.Join(yamlStrings(t, job["before_script"], "rza_migrate_check_gate.before_script"), "\n")
	if beforeScript != "" {
		t.Fatalf("MR checker must isolate itself from central before_script: %q", beforeScript)
	}
	rules, ok := job["rules"].([]interface{})
	if !ok || len(rules) != 3 || !strings.Contains(string(gateRaw), "merge_request_event") ||
		!strings.Contains(string(gateRaw), "CI_COMMIT_BRANCH") || !strings.Contains(string(gateRaw), "CI_COMMIT_TAG") {
		t.Fatalf("MR/deploy rules contract mismatch: %#v", job["rules"])
	}

	binaryPath := filepath.Join(repo, "tools", "linux-amd64", "rza_migrate_check")
	binary, err := os.ReadFile(binaryPath)
	if err != nil {
		t.Fatalf("read Linux checker: %v", err)
	}
	if len(binary) < 4 || string(binary[:4]) != "\x7fELF" {
		t.Fatalf("Linux checker is not an ELF binary: %x", binary[:min(4, len(binary))])
	}
	info, err := buildinfo.ReadFile(binaryPath)
	if err != nil {
		t.Fatalf("read Linux checker build info: %v", err)
	}
	settings := map[string]string{}
	for _, setting := range info.Settings {
		settings[setting.Key] = setting.Value
	}
	if settings["GOOS"] != "linux" || settings["GOARCH"] != "amd64" || settings["CGO_ENABLED"] != "0" {
		t.Fatalf("Linux checker target mismatch: GOOS=%q GOARCH=%q CGO_ENABLED=%q", settings["GOOS"], settings["GOARCH"], settings["CGO_ENABLED"])
	}

	manifestRaw, err := os.ReadFile(binaryPath + ".sha256")
	if err != nil {
		t.Fatalf("read checker checksum: %v", err)
	}
	fields := strings.Fields(string(manifestRaw))
	if len(fields) != 2 || fields[1] != "tools/linux-amd64/rza_migrate_check" {
		t.Fatalf("invalid checker checksum manifest: %q", manifestRaw)
	}
	sum := sha256.Sum256(binary)
	if fields[0] != hex.EncodeToString(sum[:]) {
		t.Fatalf("checker checksum mismatch: manifest=%s actual=%s", fields[0], hex.EncodeToString(sum[:]))
	}

	buildInfoRaw, err := os.ReadFile(filepath.Join(repo, "tools", "linux-amd64", "build-info.txt"))
	if err != nil {
		t.Fatalf("read release build identity: %v", err)
	}
	sourceDigest := buildInfoValue(t, string(buildInfoRaw), "source_sha256")
	if len(sourceDigest) != 64 {
		t.Fatalf("invalid release source digest: %q", sourceDigest)
	}
	actualSourceDigest := releaseSourceDigest(t, repo)
	if sourceDigest != actualSourceDigest {
		t.Fatalf("Linux checker release is stale: embedded metadata=%s current source=%s; run build_all.bat", sourceDigest, actualSourceDigest)
	}
	if !bytes.Contains(binary, []byte(sourceDigest)) {
		t.Fatalf("Linux checker does not embed release source digest %s", sourceDigest)
	}
	validateReleaseManifestHashes(t, repo)

	licensePath := filepath.Join(repo, "tools", "license.key")
	if _, err := os.Stat(licensePath); err != nil {
		licensePath = filepath.Join(repo, "license.key")
	}
	license, err := os.ReadFile(licensePath)
	if err != nil || len(strings.TrimSpace(string(license))) == 0 {
		t.Fatalf("repository license.key is missing or empty: %v", err)
	}
	batchRaw, err := os.ReadFile(filepath.Join(repo, "scripts", "build_migrate_check_linux.bat"))
	if err != nil {
		t.Fatalf("read Linux checker build batch: %v", err)
	}
	batch := string(batchRaw)
	for _, required := range []string{"build_all.bat", "RZA_LICENSE_OWNER", "RZA_LICENSE_MONTHS"} {
		if !strings.Contains(batch, required) {
			t.Fatalf("build batch is missing %q", required)
		}
	}
}

func buildInfoValue(t *testing.T, text, key string) string {
	t.Helper()
	prefix := key + "="
	for _, line := range strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n") {
		if strings.HasPrefix(line, prefix) {
			return strings.TrimPrefix(line, prefix)
		}
	}
	t.Fatalf("build info is missing %s", key)
	return ""
}

func releaseSourceDigest(t *testing.T, repo string) string {
	t.Helper()
	files := make([]string, 0)
	for _, root := range []string{"cmd/rza_migrate_db", "cmd/rza_migrate_check", "cmd/rza_migrate_report", "internal"} {
		err := filepath.WalkDir(filepath.Join(repo, filepath.FromSlash(root)), func(path string, entry os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !entry.IsDir() && strings.EqualFold(filepath.Ext(entry.Name()), ".go") {
				relative, err := filepath.Rel(repo, path)
				if err != nil {
					return err
				}
				files = append(files, filepath.ToSlash(relative))
			}
			return nil
		})
		if err != nil {
			t.Fatalf("enumerate release sources: %v", err)
		}
	}
	files = append(files,
		"go.mod",
		"go.sum",
		"build_all.bat",
		"scripts/release_artifacts.ps1",
		"scripts/verify_release_source.sh",
		".gitlab/rza_migrate_check.gitlab-ci.yml",
	)
	sort.Strings(files)
	var inventory strings.Builder
	for _, relative := range files {
		raw, err := os.ReadFile(filepath.Join(repo, filepath.FromSlash(relative)))
		if err != nil {
			t.Fatalf("read release source %s: %v", relative, err)
		}
		raw = bytes.ReplaceAll(raw, []byte("\r\n"), []byte("\n"))
		sum := sha256.Sum256(raw)
		fmt.Fprintf(&inventory, "%s  %s\n", hex.EncodeToString(sum[:]), relative)
	}
	sum := sha256.Sum256([]byte(inventory.String()))
	return hex.EncodeToString(sum[:])
}

func validateReleaseManifestHashes(t *testing.T, repo string) {
	t.Helper()
	raw, err := os.ReadFile(filepath.Join(repo, "tools", "linux-amd64", "release.sha256"))
	if err != nil {
		t.Fatalf("read complete release checksum manifest: %v", err)
	}
	entries := 0
	for _, line := range strings.Split(strings.TrimSpace(string(raw)), "\n") {
		parts := strings.SplitN(strings.TrimSuffix(line, "\r"), "  ", 2)
		if len(parts) != 2 || len(parts[0]) != 64 {
			t.Fatalf("invalid release checksum entry: %q", line)
		}
		path := filepath.Clean(filepath.Join(repo, filepath.FromSlash(parts[1])))
		relative, err := filepath.Rel(repo, path)
		if err != nil || strings.HasPrefix(relative, "..") || filepath.IsAbs(relative) {
			t.Fatalf("unsafe release checksum path: %q", parts[1])
		}
		payload, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("read release payload %s: %v", parts[1], err)
		}
		sum := sha256.Sum256(payload)
		if hex.EncodeToString(sum[:]) != parts[0] {
			t.Fatalf("release checksum mismatch for %s", parts[1])
		}
		entries++
	}
	if entries < 10 {
		t.Fatalf("complete release manifest has only %d entries", entries)
	}
}

func yamlStrings(t *testing.T, value interface{}, label string) []string {
	t.Helper()
	items, ok := value.([]interface{})
	if !ok {
		t.Fatalf("%s must be a YAML list, got %#v", label, value)
	}
	out := make([]string, 0, len(items))
	for _, item := range items {
		text, ok := item.(string)
		if !ok {
			t.Fatalf("%s contains non-string value %#v", label, item)
		}
		out = append(out, text)
	}
	return out
}
