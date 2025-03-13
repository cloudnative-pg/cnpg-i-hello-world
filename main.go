// Package main is the entrypoint of the application
package main

import (
	"fmt"
	"os"

	"github.com/cloudnative-pg/machinery/pkg/log"
	"github.com/spf13/cobra"
	corelog "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	"github.com/cloudnative-pg/cnpg-i-hello-world/cmd/plugin"
)

func main() {
	logFlags := &log.Flags{}
	rootCmd := &cobra.Command{
		Use:   "cnpg-i-hello-world",
		Short: "A plugin example",
		PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
			logFlags.ConfigureLogging()
			return nil
		},
	}

	corelog.SetLogger(zap.New())
	logFlags.AddFlags(rootCmd.PersistentFlags())

	rootCmd.AddCommand(plugin.NewCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
