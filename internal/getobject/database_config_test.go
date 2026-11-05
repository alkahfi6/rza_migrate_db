package getobject

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/microsoft/go-mssqldb/msdsn"
)

func TestDatabaseConnectionStringPreservesLiteralCredentials(t *testing.T) {
	for _, password := range []string{"p@ssword", "@:/?#[]%&=+;\\\"'", "literal%40password", " space + unicode 密碼 "} {
		t.Run(password, func(t *testing.T) {
			for _, host := range []string{"sql.example.test", "127.0.0.1", "2001:db8::1"} {
				cfg := DatabaseConfig{Host: host, Port: 1444, User: `DOMAIN\user@name`, Password: password, InitialCatalog: "Sales & Ops+#?=資料", Encrypt: "disable"}
				dsn, err := cfg.ConnectionString()
				if err != nil {
					t.Fatal(err)
				}
				// Use the actual SQL Server driver's parser as the consumer oracle.
				parsed, err := msdsn.Parse(dsn)
				if err != nil {
					t.Fatal(err)
				}
				if parsed.User != cfg.User || parsed.Password != password || parsed.Database != cfg.InitialCatalog || parsed.Host != host || parsed.Port != 1444 {
					t.Fatal("driver did not receive the original credentials/server/catalog")
				}
			}
		})
	}
}

func TestLoadYAMLDatabaseConfig(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.yaml")
	raw := "database:\n  host: localhost\n  user: 'test@user'\n  password: ' p@ss:#?/%40&+''word '\n  initial_catalog: 'Sales & Ops'\nrepo_root: './target'\nrender_workers: 7\nstrip_module_comments: false\n"
	if err := os.WriteFile(path, []byte(raw), 0o600); err != nil {
		t.Fatal(err)
	}
	cfg, err := LoadAppConfig(path)
	if err != nil {
		t.Fatal(err)
	}
	parsed, err := msdsn.Parse(cfg.ConnectionString)
	if err != nil {
		t.Fatal(err)
	}
	if parsed.Password != " p@ss:#?/%40&+'word " || parsed.User != "test@user" || parsed.Database != "Sales & Ops" || parsed.Port != 1433 {
		t.Fatal("YAML credentials or default port changed")
	}
	if cfg.RepoRoot != "target" || cfg.RenderWorkers != 7 || cfg.StripModuleComments || cfg.QueryTimeoutSeconds != 120 {
		t.Fatal("YAML options/defaults were not retained")
	}
}

func TestLoadYAMLRejectsInvalidConfigWithoutSecrets(t *testing.T) {
	valid := "database:\n  host: localhost\n  port: 1433\n  user: test\n  password: 'SECRET@value'\n  initial_catalog: master\n"
	cases := []string{
		strings.Replace(valid, "port: 1433", "port: SECRET@value", 1),
		strings.Replace(valid, "port: 1433", "port: -1", 1),
		strings.Replace(valid, "port: 1433", "port: 65536", 1),
		strings.Replace(valid, "user: test", "user: ''", 1),
		strings.Replace(valid, "initial_catalog: master", "initial_catalog: ''", 1),
		strings.Replace(valid, "password: 'SECRET@value'", "password: ''", 1),
		strings.Replace(valid, "host: localhost", "host: 'SECRET@value/instance'", 1),
		valid + "connection_string: 'SECRET@value'\n",
		valid + "unknown_field: 'SECRET@value'\n",
		valid + "---\npassword: 'SECRET@value'\n",
		valid + "database: 'SECRET@value'\n",
	}
	for i, raw := range cases {
		path := filepath.Join(t.TempDir(), "config.yaml")
		if err := os.WriteFile(path, []byte(raw), 0o600); err != nil {
			t.Fatal(err)
		}
		_, err := LoadAppConfig(path)
		if err == nil {
			t.Fatalf("case %d: invalid config accepted", i)
		}
		if strings.Contains(err.Error(), "SECRET") {
			t.Fatalf("case %d: error disclosed secret", i)
		}
	}
}

func TestWriteYAMLConfigTemplateRoundTrip(t *testing.T) {
	cfg := DefaultAppConfig(".")
	cfg.ConnectionString = "sqlserver://user:pass@localhost:1433?database=master&encrypt=disable"
	path := filepath.Join(t.TempDir(), "template.yml")
	if err := WriteAppConfigTemplate(path, cfg); err != nil {
		t.Fatal(err)
	}
	loaded, err := LoadAppConfig(path)
	if err != nil {
		t.Fatal(err)
	}
	if loaded.ConnectionString != cfg.ConnectionString || loaded.QueryTimeoutSeconds != cfg.QueryTimeoutSeconds {
		t.Fatal("template round trip changed config")
	}
}
