package plugin

import (
	"github.com/cloudnative-pg/cnpg-i-machinery/pkg/pluginhelper"
	"github.com/cloudnative-pg/cnpg-i/pkg/lifecycle"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	"github.com/cloudnative-pg/cnpg-i-hello-world/internal/identity"
	lifecycleImpl "github.com/cloudnative-pg/cnpg-i-hello-world/internal/lifecycle"
)

// NewCmd creates the `plugin` command
func NewCmd() *cobra.Command {
	cmd := pluginhelper.CreateMainCmd(identity.Implementation{}, func(server *grpc.Server) error {
		// Register the declared implementations
		lifecycle.RegisterOperatorLifecycleServer(server, lifecycleImpl.Implementation{})
		return nil
	})

	cmd.Use = "plugin"

	return cmd
}
