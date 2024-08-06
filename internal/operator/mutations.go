package operator

import (
	"context"

	"github.com/cloudnative-pg/cnpg-i-machinery/pkg/pluginhelper"
	"github.com/cloudnative-pg/cnpg-i/pkg/operator"

	"github.com/cloudnative-pg/cnpg-i-hello-world/internal/config"
	"github.com/cloudnative-pg/cnpg-i-hello-world/pkg/metadata"
)

// MutateCluster is called to mutate a cluster with the defaulting webhook.
// This function is defaulting the "imagePullPolicy" plugin parameter
func (Implementation) MutateCluster(
	_ context.Context,
	request *operator.OperatorMutateClusterRequest,
) (*operator.OperatorMutateClusterResult, error) {
	helper, err := pluginhelper.NewDataBuilder(
		metadata.PluginName,
		request.GetDefinition(),
	).Build()
	if err != nil {
		return nil, err
	}

	config, valErrs := config.FromParameters(helper)
	if len(valErrs) > 0 {
		return nil, valErrs[0]
	}

	mutatedCluster := helper.GetCluster().DeepCopy()
	for i := range mutatedCluster.Spec.Plugins {
		if mutatedCluster.Spec.Plugins[i].Name != metadata.PluginName {
			continue
		}

		if mutatedCluster.Spec.Plugins[i].Parameters == nil {
			mutatedCluster.Spec.Plugins[i].Parameters = make(map[string]string)
		}

		mutatedCluster.Spec.Plugins[i].Parameters, err = config.ToParameters()
		if err != nil {
			return nil, err
		}
	}

	patch, err := helper.CreateClusterJSONPatch(*mutatedCluster)
	if err != nil {
		return nil, err
	}

	return &operator.OperatorMutateClusterResult{
		JsonPatch: patch,
	}, nil
}
