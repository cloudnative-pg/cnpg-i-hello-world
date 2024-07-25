// Package lifecycle implements the lifecycle hooks
package lifecycle

import (
	"context"
	"fmt"

	"github.com/cloudnative-pg/cnpg-i-machinery/pkg/logging"
	"github.com/cloudnative-pg/cnpg-i-machinery/pkg/pluginhelper"
	"github.com/cloudnative-pg/cnpg-i/pkg/lifecycle"

	"github.com/cloudnative-pg/cnpg-i-hello-world/internal/config"
	"github.com/cloudnative-pg/cnpg-i-hello-world/internal/utils"
	"github.com/cloudnative-pg/cnpg-i-hello-world/pkg/metadata"
)

// Implementation is the implementation of the lifecycle handler
type Implementation struct {
	lifecycle.UnimplementedOperatorLifecycleServer
}

// GetCapabilities exposes the lifecycle capabilities
func (impl Implementation) GetCapabilities(
	_ context.Context,
	_ *lifecycle.OperatorLifecycleCapabilitiesRequest,
) (*lifecycle.OperatorLifecycleCapabilitiesResponse, error) {
	return &lifecycle.OperatorLifecycleCapabilitiesResponse{
		LifecycleCapabilities: []*lifecycle.OperatorLifecycleCapabilities{
			{
				Group: "",
				Kind:  "Pod",
				OperationTypes: []*lifecycle.OperatorOperationType{
					{
						Type: lifecycle.OperatorOperationType_TYPE_CREATE,
					},
					{
						Type: lifecycle.OperatorOperationType_TYPE_PATCH,
					},
				},
			},
		},
	}, nil
}

// LifecycleHook is called when creating Kubernetes services
func (impl Implementation) LifecycleHook(
	ctx context.Context,
	request *lifecycle.OperatorLifecycleRequest,
) (*lifecycle.OperatorLifecycleResponse, error) {
	kind, err := utils.GetKind(request.GetObjectDefinition())
	if err != nil {
		return nil, err
	}
	operation := request.GetOperationType().GetType().Enum()
	if operation == nil {
		return nil, fmt.Errorf("no operation set")
	}

	//nolint: gocritic
	switch kind {
	case "Pod":
		switch *operation {
		case lifecycle.OperatorOperationType_TYPE_CREATE, lifecycle.OperatorOperationType_TYPE_PATCH,
			lifecycle.OperatorOperationType_TYPE_UPDATE:
			return impl.reconcileMetadata(ctx, request)
		}
		// add any other custom logic to execute based on the operation
	}

	return &lifecycle.OperatorLifecycleResponse{}, nil
}

// LifecycleHook is called when creating Kubernetes services
func (impl Implementation) reconcileMetadata(
	ctx context.Context,
	request *lifecycle.OperatorLifecycleRequest,
) (*lifecycle.OperatorLifecycleResponse, error) {
	logger := logging.FromContext(ctx).WithName("cnpg_i_example_lifecyle")
	helper, err := pluginhelper.NewDataBuilder(
		metadata.PluginName,
		request.GetClusterDefinition(),
	).WithPod(request.GetObjectDefinition()).Build()
	if err != nil {
		return nil, err
	}
	configuration, valErrs := config.FromParameters(helper)
	if len(valErrs) > 0 {
		return nil, valErrs[0]
	}
	mutatedPod := helper.GetPod().DeepCopy()

	helper.InjectPluginVolume(mutatedPod)

	// Apply any custom logic needed here, in this example we just add some metadata to the pod

	for key, value := range configuration.Labels {
		mutatedPod.Labels[key] = value
	}
	for key, value := range configuration.Annotations {
		mutatedPod.Annotations[key] = value
	}

	patch, err := helper.CreatePodJSONPatch(*mutatedPod)
	if err != nil {
		return nil, err
	}

	logger.V(0).Info("generated patch", "content", string(patch), "configuration", configuration)

	return &lifecycle.OperatorLifecycleResponse{
		JsonPatch: patch,
	}, nil
}
