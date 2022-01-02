// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gandi

import (
	"fmt"
	"path/filepath"

	"github.com/go-gandi/terraform-provider-gandi/gandi"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/vincentbernat/pulumi-gandi/provider/pkg/version"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "gandi"
	// modules:
	domainMod        = "domain"
	livednsMod       = "livedns"
	emailMod         = "email"
	simplehostingMod = "simplehosting"
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(gandi.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:                    p,
		Name:                 "gandi",
		Description:          "A Pulumi package for creating and managing gandi cloud resources.",
		Keywords:             []string{"pulumi", "gandi"},
		License:              "Apache-2.0",
		Homepage:             "https://pulumi.io",
		Repository:           "https://github.com/vincentbernat/pulumi-gandi",
		Config:               map[string]*tfbridge.SchemaInfo{},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"gandi_dnssec_key":             {Tok: tfbridge.MakeResource(mainPkg, domainMod, "DnssecKey")},
			"gandi_domain":                 {Tok: tfbridge.MakeResource(mainPkg, domainMod, "Domain")},
			"gandi_nameservers":            {Tok: tfbridge.MakeResource(mainPkg, domainMod, "Nameservers")},
			"gandi_glue_record":            {Tok: tfbridge.MakeResource(mainPkg, domainMod, "GlueRecord")},
			"gandi_livedns_domain":         {Tok: tfbridge.MakeResource(mainPkg, livednsMod, "Domain")},
			"gandi_livedns_record":         {Tok: tfbridge.MakeResource(mainPkg, livednsMod, "Record")},
			"gandi_mailbox":                {Tok: tfbridge.MakeResource(mainPkg, emailMod, "Mailbox")},
			"gandi_email_forwarding":       {Tok: tfbridge.MakeResource(mainPkg, emailMod, "Forwarding")},
			"gandi_simplehosting_instance": {Tok: tfbridge.MakeResource(mainPkg, simplehostingMod, "Instance")},
			"gandi_simplehosting_vhost":    {Tok: tfbridge.MakeResource(mainPkg, simplehostingMod, "Vhost")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"gandi_domain":            {Tok: tfbridge.MakeDataSource(mainPkg, domainMod, "getDomain")},
			"gandi_glue_record":       {Tok: tfbridge.MakeDataSource(mainPkg, domainMod, "getGlueRecord")},
			"gandi_livedns_domain":    {Tok: tfbridge.MakeDataSource(mainPkg, livednsMod, "getDomain")},
			"gandi_livedns_domain_ns": {Tok: tfbridge.MakeDataSource(mainPkg, livednsMod, "getNameservers")},
			"gandi_mailbox":           {Tok: tfbridge.MakeDataSource(mainPkg, emailMod, "getMailbox")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
