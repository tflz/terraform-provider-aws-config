// Copyright IBM Corp. 2021, 2025
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure AWSLandingZoneConfigProvider satisfies various provider interfaces.
var _ provider.Provider = &AWSLandingZoneConfigProvider{}
var _ provider.ProviderWithFunctions = &AWSLandingZoneConfigProvider{}
var _ provider.ProviderWithEphemeralResources = &AWSLandingZoneConfigProvider{}
var _ provider.ProviderWithActions = &AWSLandingZoneConfigProvider{}

// AWSLandingZoneConfigProvider defines the provider implementation.
type AWSLandingZoneConfigProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// AWSLandingZoneConfigProviderModel describes the provider data model.
type AWSLandingZoneConfigProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
}

func (p *AWSLandingZoneConfigProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "aws-landing-zone-config"
	resp.Version = p.version
}

func (p *AWSLandingZoneConfigProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "Example provider attribute",
				Optional:            true,
			},
		},
	}
}

func (p *AWSLandingZoneConfigProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data AWSLandingZoneConfigProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Configuration values are now available.
	// if data.Endpoint.IsNull() { /* ... */ }

	// Example client configuration for data sources and resources
	client := http.DefaultClient
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *AWSLandingZoneConfigProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *AWSLandingZoneConfigProvider) EphemeralResources(ctx context.Context) []func() ephemeral.EphemeralResource {
	return []func() ephemeral.EphemeralResource{}
}

func (p *AWSLandingZoneConfigProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *AWSLandingZoneConfigProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		NewLoadOrganizationConfigFunction,
	}
}

func (p *AWSLandingZoneConfigProvider) Actions(ctx context.Context) []func() action.Action {
	return []func() action.Action{}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &AWSLandingZoneConfigProvider{
			version: version,
		}
	}
}
