package config

import (
	"encoding/json"
	"reflect"

	"github.com/cloudnative-pg/cnpg-i-machinery/pkg/pluginhelper/common"
	"github.com/cloudnative-pg/cnpg-i-machinery/pkg/pluginhelper/validation"
	"github.com/cloudnative-pg/cnpg-i/pkg/operator"
)

const (
	labelsParameter     = "labels"
	annotationParameter = "annotations"
)

// Configuration represents the plugin configuration parameters
type Configuration struct {
	Labels      map[string]string
	Annotations map[string]string
}

// FromParameters builds a plugin configuration from the configuration parameters
func FromParameters(
	helper *common.Plugin,
) (*Configuration, []*operator.ValidationError) {
	validationErrors := make([]*operator.ValidationError, 0)

	var labels map[string]string
	if helper.Parameters[labelsParameter] != "" {
		if err := json.Unmarshal([]byte(helper.Parameters[labelsParameter]), &labels); err != nil {
			validationErrors = append(
				validationErrors,
				validation.BuildErrorForParameter(helper, labelsParameter, err.Error()),
			)
		}
	}

	var annotations map[string]string
	if helper.Parameters[annotationParameter] != "" {
		if err := json.Unmarshal([]byte(helper.Parameters[annotationParameter]), &annotations); err != nil {
			validationErrors = append(
				validationErrors,
				validation.BuildErrorForParameter(helper, annotationParameter, err.Error()),
			)
		}
	}

	configuration := &Configuration{
		Labels:      labels,
		Annotations: annotations,
	}

	configuration.applyDefaults()

	return configuration, validationErrors
}

// ValidateChanges validates the changes between the old configuration to the
// new configuration
func ValidateChanges(
	oldConfiguration *Configuration,
	newConfiguration *Configuration,
	helper *common.Plugin,
) []*operator.ValidationError {
	validationErrors := make([]*operator.ValidationError, 0)

	if !reflect.DeepEqual(oldConfiguration.Labels, newConfiguration.Labels) {
		validationErrors = append(
			validationErrors,
			validation.BuildErrorForParameter(helper, labelsParameter, "Labels cannot be changed"))
	}

	return validationErrors
}

// ToParameters serialize the configuration to a map of plugin parameters
func (config *Configuration) ToParameters() (map[string]string, error) {
	result := make(map[string]string)
	serializedLabels, err := json.Marshal(config.Labels)
	if err != nil {
		return nil, err
	}
	serializedAnnotations, err := json.Marshal(config.Annotations)
	if err != nil {
		return nil, err
	}
	result[labelsParameter] = string(serializedLabels)
	result[annotationParameter] = string(serializedAnnotations)

	return result, nil
}

// applyDefaults fills the configuration with the defaults
func (config *Configuration) applyDefaults() {
	if len(config.Labels) == 0 {
		config.Labels = map[string]string{
			"plugin-metadata": "default",
		}
	}
	if len(config.Annotations) == 0 {
		config.Annotations = map[string]string{
			"plugin-metadata": "default",
		}
	}
}
