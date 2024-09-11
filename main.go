// Package main is the entry point of the application
package main

import (
	"hexagonal-arch/cmd"
	_ "hexagonal-arch/docs"

	_ "github.com/spf13/viper/remote"
)

func main() {
	cmd.Execute()
}
