// Package main is the entrypoint of the application
package main

import (
	"fmt"
	"os"

	"github.com/cloudnative-pg/machinery/pkg/log"
	"github.com/spf13/cobra"

	"github.com/cloudnative-pg/cnpg-i-hello-world/cmd/plugin"
)

func main() {
	cobra.EnableTraverseRunHooks = true

	logFlags := &log.Flags{}
	rootCmd := &cobra.Command{
		Use:   "cnpg-i-hello-world",
		Short: "A plugin example",
		PersistentPreRun: func(cmd *cobra.Command, _ []string) {
			logFlags.ConfigureLogging()
			cmd.SetContext(log.IntoContext(cmd.Context(), log.GetLogger()))
		},
	}

	logFlags.AddFlags(rootCmd.PersistentFlags())

	rootCmd.AddCommand(plugin.NewCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
