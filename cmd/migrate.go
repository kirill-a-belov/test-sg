package cmd

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
	"github.com/kirill-a-belov/test-sg/storages/pg"
)

var migrations = &migrate.MemoryMigrationSource{
	Migrations: []*migrate.Migration{
		{
			Id: "Tables creation",
			Up: []string{`
				create table requests
				(
					id serial not null
						constraint requests_pkey
							primary key,
					created_at timestamp with time zone,
					updated_at timestamp with time zone,
					deleted_at timestamp with time zone,
					p_id text not null
						constraint requests_p_id_key
							unique,
					url   text,
					title text,
					price text,
					image text,
					is_in_stock boolean
				);
				
				create index idx_requests_deleted_at
					on requests (deleted_at);
									`},
			Down: []string{`
				drop index if exists index idx_requests_deleted_at;
				drop table if exists requests;

			`},
		},
	},
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Configurator service DB migration",
	Run: func(cmd *cobra.Command, args []string) {
		log.Print("***scrapper migrations***")

		if err := config.postgresConfig.Validate(); err != nil {
			log.Fatalf("invalid postgres configuration %v:", config.postgresConfig)
		}

		storage, err := pg.NewPGStorage(&config.postgresConfig)
		if err != nil {
			log.Fatalf("storage %v:", config.postgresConfig)
		}

		n, err := migrate.Exec(storage.GetDB(), "postgres", migrations, migrate.Up)
		if err != nil {
			log.Fatalf("migration applying error: %v", err)
		}

		storage.Close()
		log.Printf("applied %d migrations!\n", n)
	},
}

func init() {
	// init config
	RootCmd.AddCommand(migrateCmd)
	migrateCmd.Flags().AddFlagSet(config.Flags())
}
