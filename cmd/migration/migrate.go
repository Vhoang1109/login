package migrate

import (
	"fmt"
	"log"

	"github.com/Vhoang1109/login/migration" 
	"github.com/Vhoang1109/share-module/config"
	"github.com/Vhoang1109/share-module/system"
	"github.com/spf13/cobra"
)

var migrate = &cobra.Command{ 
	Use:   "migrate",
	Short: "Manage database migrations",
}

var up = &cobra.Command{
	Use:   "up",
	Short: "Apply database migrations to update the schema",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadAppConfig(".")
		if err != nil {
			return fmt.Errorf("error loading config: %w", err)
		}
		sys := system.New(cfg, cmd.Parent().Name())

		if err := sys.MigrateDB(migration.FS); err != nil {
			return fmt.Errorf("migration failed: %w", err)
		}
		log.Println("Migration completed successfully")
		return nil
	},
}

var down = &cobra.Command{
	Use:   "down",
	Short: "Rollback the latest database migration",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadAppConfig(".")
		if err != nil {
			log.Fatalf("error loading config: %v", err)
		}
		sys := system.New(cfg, cmd.Parent().Name())

		if err := sys.RollbackDB(); err != nil {
			log.Fatalf("Rollback failed: %v", err)
		}
		log.Println("Rollback completed successfully")
	},
}

func RegisterMigrate(root *cobra.Command) {
	migrate.AddCommand(up, down)
	root.AddCommand(migrate)
}
