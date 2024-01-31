package cmd

import (
	"starter-go-gin/config"

	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "for db migration",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		sqlDB := config.SQLDBConn()

		if err := goose.SetDialect("mysql"); err != nil {
			panic(err)
		}

		err := goose.Up(sqlDB, "database/migrations")
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
