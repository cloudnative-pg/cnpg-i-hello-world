package operator

import (
	"context"
	"fmt"

	"github.com/cloudnative-pg/cnpg-i-machinery/pkg/pluginhelper"
	"github.com/cloudnative-pg/cnpg-i/pkg/operator"

	"github.com/cloudnative-pg/cnpg-i-hello-world/internal/config"
	"github.com/cloudnative-pg/cnpg-i-hello-world/pkg/metadata"
)

// ValidateClusterCreate validates a cluster that is being created
func (Implementation) ValidateClusterCreate(
	_ context.Context,
	request *operator.OperatorValidateClusterCreateRequest,
) (*operator.OperatorValidateClusterCreateResult, error) {
	result := &operator.OperatorValidateClusterCreateResult{}

	helper, err := pluginhelper.NewDataBuilder(
		metadata.PluginName,
		request.GetDefinition(),
	).Build()
	if err != nil {
		return nil, err
	}

	_, result.ValidationErrors = config.FromParameters(helper)

	return result, nil
}

// ValidateClusterChange validates a cluster that is being changed
func (Implementation) ValidateClusterChange(
	_ context.Context,
	request *operator.OperatorValidateClusterChangeRequest,
) (*operator.OperatorValidateClusterChangeResult, error) {
	result := &operator.OperatorValidateClusterChangeResult{}

	oldClusterHelper, err := pluginhelper.NewDataBuilder(
		metadata.PluginName,
		request.GetOldCluster(),
	).Build()
	if err != nil {
		return nil, fmt.Errorf("while parsing old cluster: %w", err)
	}

	newClusterHelper, err := pluginhelper.NewDataBuilder(
		metadata.PluginName,
		request.GetNewCluster(),
	).Build()
	if err != nil {
		return nil, fmt.Errorf("while parsing new cluster: %w", err)
	}

	var newConfiguration *config.Configuration
	newConfiguration, result.ValidationErrors = config.FromParameters(newClusterHelper)
	oldConfiguration, _ := config.FromParameters(oldClusterHelper)
	result.ValidationErrors = config.ValidateChanges(oldConfiguration, newConfiguration, newClusterHelper)

	return result, nil
}
