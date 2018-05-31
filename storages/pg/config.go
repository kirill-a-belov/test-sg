package pg

import (
	"fmt"

	"github.com/spf13/pflag"
)

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func (c *PostgresConfig) Flags() *pflag.FlagSet {
	f := pflag.NewFlagSet("PostgresConfig", pflag.PanicOnError)
	f.StringVar(&c.Host,
		"host",
		"postgres",
		"PG host")
	f.IntVar(&c.Port,
		"port",
		5432,
		"PG port")
	f.StringVar(&c.User,
		"user",
		"postgres",
		"PG user")
	f.StringVar(&c.Password,
		"passwords",
		"postgres",
		"PG password")
	f.StringVar(&c.DBName,
		"dbname",
		"postgres",
		"PG database name")
	return f
}

func (c *PostgresConfig) Validate() error {
	if c.Host == "" {
		return fmt.Errorf("configuration validation error: not valid host %v", c.Host)
	}

	if c.User == "" {
		return fmt.Errorf("configuration validation error: not valid host %v", c.User)
	}

	if c.Password == "" {
		return fmt.Errorf("configuration validation error: not valid host %v", c.Password)
	}

	if c.DBName == "" {
		return fmt.Errorf("configuration validation error: not valid host %v", c.DBName)
	}

	if c.Port < 0 || c.Port > 49150 {
		return fmt.Errorf("configuration validation error: not valid port %v", c.Port)
	}

	return nil
}
