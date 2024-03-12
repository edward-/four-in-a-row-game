package cmd

import (
	"log"

	"github.com/edward-/four-in-a-row-game/pkg/migrate"
	"github.com/spf13/cobra"
)

func init() {
	commands.AddCommand(buildMigrateCommand())
}

// migrateCmd represents the migrate command.
func buildMigrateCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "Run migrations of database",
		Run: func(_ *cobra.Command, args []string) {
			if len(args) > 0 {
				switch args[0] {
				case "up":
					log.Println("migration upping...")
					migrate.Up()
				case "down":
					log.Println("migration downing...")
					migrate.Down()
				}
			}
		},
	}
}
