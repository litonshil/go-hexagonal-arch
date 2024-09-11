package cmd

import (
	"log"

	"hexagonal-arch/config"
	app "hexagonal-arch/internal/application"
	"hexagonal-arch/internal/framework/http"
	db "hexagonal-arch/internal/framework/mysql"

	"github.com/spf13/cobra"
)

// Run ...
// @title hexagonal-arch
// @version 1.0
// @description API Documentation for hexagonal-arch
// @contact.name API Support
// @host localhost:7766
// @BasePath /
// @schemes http
var runCmd = &cobra.Command{
	Use:   "serve",
	Short: "run server",
	Run: func(cmd *cobra.Command, args []string) {
		dbConfig := config.GetDBConfig()
		dbAdapter, err := db.NewAdapterWithConfig(dbConfig)
		if err != nil {
			log.Fatalln(err.Error())
		}

		application := app.NewApplication(dbAdapter)

		// Creating a new http adapter with the application.
		httpAdapter := http.NewAdapter(application)
		// Running the http server.
		httpAdapter.Run()
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		// Initializing the config.
		config.Init()
	},
}

func init() {
	// Adding the run command to the root command.
	rootCmd.AddCommand(runCmd)
}
