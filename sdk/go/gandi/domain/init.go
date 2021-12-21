// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package domain

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-gandi/sdk/go/gandi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "gandi:domain/dnssecKey:DnssecKey":
		r = &DnssecKey{}
	case "gandi:domain/domain:Domain":
		r = &Domain{}
	case "gandi:domain/glueRecord:GlueRecord":
		r = &GlueRecord{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := gandi.PkgVersion()
	if err != nil {
		fmt.Printf("failed to determine package version. defaulting to v1: %v\n", err)
	}
	pulumi.RegisterResourceModule(
		"gandi",
		"domain/dnssecKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gandi",
		"domain/domain",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gandi",
		"domain/glueRecord",
		&module{version},
	)
}