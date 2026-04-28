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

package operator

import (
	"context"

	"github.com/cloudnative-pg/cnpg-i-machinery/pkg/pluginhelper/common"
	"github.com/cloudnative-pg/cnpg-i-machinery/pkg/pluginhelper/decoder"
	"github.com/cloudnative-pg/cnpg-i/pkg/operator"

	"github.com/cloudnative-pg/cnpg-i-hello-world/internal/config"
	"github.com/cloudnative-pg/cnpg-i-hello-world/pkg/metadata"
)

// ValidateClusterCreate validates a cluster that is being created
func (Implementation) ValidateClusterCreate(
	_ context.Context,
	request *operator.OperatorValidateClusterCreateRequest,
) (*operator.OperatorValidateClusterCreateResult, error) {
	cluster, err := decoder.DecodeClusterLenient(request.GetDefinition())
	if err != nil {
		return nil, err
	}

	result := &operator.OperatorValidateClusterCreateResult{}

	helper := common.NewPlugin(
		*cluster,
		metadata.PluginName,
	)

	_, result.ValidationErrors = config.FromParameters(helper)

	return result, nil
}

// ValidateClusterChange validates a cluster that is being changed
func (Implementation) ValidateClusterChange(
	_ context.Context,
	request *operator.OperatorValidateClusterChangeRequest,
) (*operator.OperatorValidateClusterChangeResult, error) {
	result := &operator.OperatorValidateClusterChangeResult{}

	oldCluster, err := decoder.DecodeClusterLenient(request.GetOldCluster())
	if err != nil {
		return nil, err
	}

	newCluster, err := decoder.DecodeClusterLenient(request.GetNewCluster())
	if err != nil {
		return nil, err
	}

	oldClusterHelper := common.NewPlugin(
		*oldCluster,
		metadata.PluginName,
	)

	newClusterHelper := common.NewPlugin(
		*newCluster,
		metadata.PluginName,
	)

	var newConfiguration *config.Configuration
	newConfiguration, result.ValidationErrors = config.FromParameters(newClusterHelper)
	oldConfiguration, _ := config.FromParameters(oldClusterHelper)
	result.ValidationErrors = config.ValidateChanges(oldConfiguration, newConfiguration, newClusterHelper)

	return result, nil
}
