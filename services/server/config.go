package server

import (
	"fmt"
	"regexp"

	"github.com/spf13/pflag"
)

type ServerConfig struct {
	Address string
	Port    int
}

func (c *ServerConfig) Flags() *pflag.FlagSet {
	f := pflag.NewFlagSet("ServerServerConfig", pflag.PanicOnError)
	f.StringVar(&c.Address,
		"address",
		"0.0.0.0",
		"Address for binding listener")
	f.IntVar(&c.Port,
		"port",
		8080,
		"Port for binding listener")

	return f
}

func (c *ServerConfig) Validate() error {
	re, err := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	if err != nil {
		return fmt.Errorf("configuration validation error: %v", err)
	}

	if !re.MatchString(c.Address) {
		return fmt.Errorf("configuration validation error: not valid adress %v", c.Address)
	}

	if c.Port < 0 || c.Port > 49150 {
		return fmt.Errorf("configuration validation error: not valid port %v", c.Port)
	}

	return nil
}
