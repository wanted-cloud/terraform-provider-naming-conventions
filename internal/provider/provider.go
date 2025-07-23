package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider              = &NamingConventionsProvider{}
	_ provider.ProviderWithFunctions = &NamingConventionsProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &NamingConventionsProvider{
			version: version,
		}
	}
}

// NamingConventionsProvider is the provider implementation.
type NamingConventionsProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// Metadata returns the provider type name.
func (p *NamingConventionsProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "naming-conventions"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *NamingConventionsProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"metadata": schema.MapAttribute{
				Optional:    false,
				Required:    true,
				ElementType: types.StringType,
			},
		},
	}
}

// Configure prepares an API client for data sources and resources.
func (p *NamingConventionsProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

}

// DataSources defines the data sources implemented in the provider.
func (p *NamingConventionsProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewGenerateDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *NamingConventionsProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}

// Functions defines the functions implemented in the provider.
func (p *NamingConventionsProvider) Functions(_ context.Context) []func() function.Function {
	return []func() function.Function{
		NewGenerateFunction,
	}
}
