package cmd

import (
	"fmt"

	"github.com/kirill-a-belov/test-sg/services/server"
	server2 "github.com/kirill-a-belov/test-sg/storages/pg"
	"github.com/spf13/pflag"
)

type Config struct {
	listenerConfig server.ServerConfig
	postgresConfig server2.PostgresConfig
}

func (c *Config) Flags() *pflag.FlagSet {
	f := pflag.NewFlagSet("MainConfig", pflag.PanicOnError)
	f.AddFlagSet(c.listenerConfig.Flags())
	f.AddFlagSet(c.postgresConfig.Flags())

	return f
}

func (c *Config) Validate() error {
	if err := c.listenerConfig.Validate(); err != nil {
		return fmt.Errorf("listener: %v", err)
	}
	if err := c.postgresConfig.Validate(); err != nil {
		return fmt.Errorf("postgres: %v", err)
	}

	return nil
}
