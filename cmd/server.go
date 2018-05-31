package cmd

import (
	"log"

	"github.com/kirill-a-belov/test-sg/services/server"
	"github.com/kirill-a-belov/test-sg/storages/pg"
	"github.com/spf13/cobra"
)

var config Config

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Configurator service start",
	Run: func(cmd *cobra.Command, args []string) {
		log.Print("***configurator service***")

		if err := config.Validate(); err != nil {
			log.Fatal(err)
		}

		log.Print("init storage")
		storage, err := pg.NewPGStorage(&config.postgresConfig)
		if err != nil {
			log.Fatalf("storage error %v:", err)
		}

		listener := server.NewServer(&config.listenerConfig, storage)
		log.Print("starting listener...")
		listener.Serve()
	},
}

func init() {
	// init config
	RootCmd.AddCommand(serverCmd)
	serverCmd.Flags().AddFlagSet(config.Flags())
}
