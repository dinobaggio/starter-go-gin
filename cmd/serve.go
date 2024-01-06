package cmd

import (
	"fmt"
	"starter-go-gin/bin"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "for running server",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		app := bin.NewApp()
		port := "3000"
		app.Run(fmt.Sprint(":", port))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
