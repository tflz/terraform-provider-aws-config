package provider

import (
	"context"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type OrganizationConfig struct {
	HomeRegion string `hcl:"home_region"`
}

// Ensure the implementation satisfies the desired interfaces.
var _ function.Function = &LoadOrganizationConfigFunction{}

type LoadOrganizationConfigFunction struct{}

func NewLoadOrganizationConfigFunction() function.Function {
	return &LoadOrganizationConfigFunction{}
}

func (f *LoadOrganizationConfigFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "load_organization_config"
}

func (f *LoadOrganizationConfigFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Organization Config",
		Description: "Load and parse the organization config",
		Parameters:  []function.Parameter{},
		Return: function.ObjectReturn{
			AttributeTypes: map[string]attr.Type{
				"home_region": types.StringType,
			},
		},
	}
}

func (f *LoadOrganizationConfigFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	parser := hclparse.NewParser()
	configFile, err := parser.ParseHCLFile("organization.tflz")
	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError(err.Error()))
		return
	}

	var config OrganizationConfig
	diag := gohcl.DecodeBody(configFile.Body, nil, &config)
	if diag.HasErrors() {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError(diag.Error()))
		return
	}

	result := struct {
		HomeRegion string `tfsdk:"home_region"`
	}{
		HomeRegion: config.HomeRegion,
	}

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, &result))
}
