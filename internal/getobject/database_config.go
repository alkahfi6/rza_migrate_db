package getobject

import (
	"errors"
	"net"
	"net/url"
	"strconv"
	"strings"
)

// DatabaseConfig holds literal SQL Server credentials, never pre-escaped values.
type DatabaseConfig struct {
	Host           string `json:"host" yaml:"host"`
	Port           int    `json:"port" yaml:"port"`
	User           string `json:"user" yaml:"user"`
	Password       string `json:"password" yaml:"password"`
	InitialCatalog string `json:"initial_catalog" yaml:"initial_catalog"`
	Encrypt        string `json:"encrypt" yaml:"encrypt"`
}

// ConnectionString encodes userinfo and query parameters independently so reserved
// characters cannot change the selected server, credentials, or database.
func (c DatabaseConfig) ConnectionString() (string, error) {
	host := strings.TrimSpace(c.Host)
	if host == "" || strings.ContainsAny(host, "/\\@?# \t\r\n") || (strings.Contains(host, ":") && net.ParseIP(host) == nil) {
		return "", errors.New("database.host must be a hostname or IP address without a port or instance")
	}
	if c.Port == 0 {
		c.Port = 1433
	}
	if c.Port < 1 || c.Port > 65535 {
		return "", errors.New("database.port must be between 1 and 65535")
	}
	if strings.TrimSpace(c.User) == "" || c.Password == "" || strings.TrimSpace(c.InitialCatalog) == "" {
		return "", errors.New("database.user, database.password and database.initial_catalog are required")
	}
	encrypt := strings.ToLower(strings.TrimSpace(c.Encrypt))
	if encrypt == "" {
		encrypt = "disable"
	}
	switch encrypt {
	case "disable", "false", "true", "strict":
	default:
		return "", errors.New("database.encrypt must be disable, false, true or strict")
	}
	query := url.Values{}
	query.Set("database", c.InitialCatalog)
	query.Set("encrypt", encrypt)
	dsn := url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(c.User, c.Password),
		Host:     net.JoinHostPort(host, strconv.Itoa(c.Port)),
		RawQuery: query.Encode(),
	}
	return dsn.String(), nil
}
