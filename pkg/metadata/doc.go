// Package metadata contains the metadata of this plugin
package metadata

import "github.com/cloudnative-pg/cnpg-i/pkg/identity"

// PluginName is the name of the plugin.
const PluginName = "cnpg-i-hello-world.cloudnative-pg.io"

// Data is the metadata of this plugin.
var Data = identity.GetPluginMetadataResponse{
	Name:          PluginName,
	Version:       "0.0.1",
	DisplayName:   "Plugin feature showcase",
	ProjectUrl:    "https://github.com/cloudnative-pg/cnpg-i-hello-world",
	RepositoryUrl: "https://github.com/cloudnative-pg/cnpg-i-hello-world",
	License:       "Proprietary",
	LicenseUrl:    "https://github.com/cloudnative-pg/cnpg-i-hello-world/LICENSE",
	Maturity:      "alpha",
}
