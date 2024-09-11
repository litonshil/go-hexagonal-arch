// Package cmd holds all the commands to run this application.
package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hexagonal-arch",
	Short: "hexagonal-arch is a sample boilerplate project",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err.Error())
	}
}
