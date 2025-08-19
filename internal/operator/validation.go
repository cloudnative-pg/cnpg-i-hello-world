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

	oldConfiguration, oldValidationErrors := config.FromParameters(oldClusterHelper)
	if len(oldValidationErrors) > 0 {
		// Combine validation errors
		result.ValidationErrors = append(result.ValidationErrors, oldValidationErrors...)
	}

	result.ValidationErrors = append(result.ValidationErrors,
		config.ValidateChanges(oldConfiguration, newConfiguration, newClusterHelper)...)

	return result, nil
}
