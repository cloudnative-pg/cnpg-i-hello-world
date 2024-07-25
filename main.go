// Package main is the entrypoint of the application
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/cloudnative-pg/cnpg-i-hello-world/cmd/plugin"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "cnpg-i-hello-world",
		Short: "A plugin example",
	}

	rootCmd.AddCommand(plugin.NewCmd())

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
