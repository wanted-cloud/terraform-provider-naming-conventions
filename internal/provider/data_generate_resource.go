package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func NewGenerateResourceDataSource() datasource.DataSource {
	return &GenerateResourceDataSource{}
}

// Ensure the implementation satisfies the desired interfaces.
var _ datasource.DataSource = &GenerateResourceDataSource{}

type GenerateResourceDataSource struct{}

type GenerateResourceDataSourceModel struct {
	ExampleAttribute types.String `tfsdk:"example_attribute"`
	ID               types.String `tfsdk:"id"`
}

func (d *GenerateResourceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "conventions_generate_resource"
}

func (d *GenerateResourceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"example_attribute": schema.StringAttribute{
				Required: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (d *GenerateResourceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data GenerateResourceDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	// Typically data sources will make external calls, however this example
	// hardcodes setting the id attribute to a specific value for brevity.
	data.ID = types.StringValue("example-id")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
