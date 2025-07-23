package provider

import (
	"bytes"
	"context"
	"html/template"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func NewGenerateDataSource() datasource.DataSource {
	return &GenerateDataSource{}
}

// Ensure the implementation satisfies the desired interfaces.
var _ datasource.DataSource = &GenerateDataSource{}

type GenerateDataSource struct{}

type GenerateDataSourceModel struct {
	Name     types.String `tfsdk:"name"`
	Type     types.String `tfsdk:"type"`
	ID       types.String `tfsdk:"id"`
	FullName types.String `tfsdk:"full_name"`
}

func (d *GenerateDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "conventions_generate"
}

func (d *GenerateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required: true,
			},
			"type": schema.StringAttribute{
				Required: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"full_name": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (d *GenerateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data GenerateDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	// Typically data sources will make external calls, however this example
	// hardcodes setting the id attribute to a specific value for brevity.
	data.ID = types.StringValue("example-id")

	// Load and Parse template
	tmpl := `{{.name}}-rg-{{.location}}`
	t := template.Must(template.New("msg").Parse(tmpl))
	var buf bytes.Buffer
	_ = t.Execute(&buf, map[string]string{
		"name":     data.Name.ValueString(),
		"location": "eastus", // This could be dynamic based on your requirements
	})
	data.FullName = types.StringValue(buf.String())

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
