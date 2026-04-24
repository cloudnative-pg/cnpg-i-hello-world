/*
Copyright © contributors to CloudNativePG, established as
CloudNativePG a Series of LF Projects, LLC.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

SPDX-License-Identifier: Apache-2.0
*/

package plugin

import (
	"github.com/cloudnative-pg/cnpg-i-machinery/pkg/pluginhelper/http"
	"github.com/cloudnative-pg/cnpg-i/pkg/lifecycle"
	"github.com/cloudnative-pg/cnpg-i/pkg/operator"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	"github.com/cloudnative-pg/cnpg-i-hello-world/internal/identity"
	lifecycleImpl "github.com/cloudnative-pg/cnpg-i-hello-world/internal/lifecycle"
	operatorImpl "github.com/cloudnative-pg/cnpg-i-hello-world/internal/operator"
)

// NewCmd creates the `plugin` command
func NewCmd() *cobra.Command {
	cmd := http.CreateMainCmd(identity.Implementation{}, func(server *grpc.Server) error {
		// Register the declared implementations
		operator.RegisterOperatorServer(server, operatorImpl.Implementation{})
		lifecycle.RegisterOperatorLifecycleServer(server, lifecycleImpl.Implementation{})
		return nil
	})

	// If you want to provide your own logr.Logger here, inject it into a context.Context
	// with logr.NewContext(ctx, logger) and pass it to cmd.SetContext(ctx)

	// Additional custom behaviour can be added by wrapping cmd.PersistentPreRun or cmd.Run

	cmd.Use = "plugin"

	return cmd
}
