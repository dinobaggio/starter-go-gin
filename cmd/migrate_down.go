package cmd

import (
	"starter-go-gin/config"

	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

var migrateDownCmd = &cobra.Command{
	Use:   "migrate:down",
	Short: "down all table",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		sqlDB := config.SQLDBConn()
		defer sqlDB.Close()

		if err := goose.SetDialect("mysql"); err != nil {
			panic(err)
		}

		err := goose.Down(sqlDB, "database/migrations")
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateDownCmd)
}
