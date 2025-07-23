package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var generateReturnAttrTypes = map[string]attr.Type{
	"name": types.StringType,
}

var _ function.Function = &GenerateFunction{}

type GenerateFunction struct{}

func NewGenerateFunction() function.Function {
	return &GenerateFunction{}
}

func (f *GenerateFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "generate"
}

func (f *GenerateFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Generate",
		Description: "Given an RFC3339 timestamp string, will parse and return an object representation of that date and time.",

		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "resource_type",
				Description: "The type of resource to generate a name for, e.g., 'azurerm_resource_group'.",
			},
			function.MapParameter{
				ElementType: types.StringType,
				Name:        "options",
				Description: "Options for the generate function",
			},
		},
		Return: function.ObjectReturn{
			AttributeTypes: generateReturnAttrTypes,
		},
	}
}

func (f *GenerateFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var resourceType string
	var options map[string]types.String

	resp.Error = req.Arguments.Get(ctx, &resourceType, &options)
	if resp.Error != nil {
		return
	}

	generateObj, diags := types.ObjectValue(
		generateReturnAttrTypes,
		map[string]attr.Value{
			"name": types.StringValue(fmt.Sprintf("%s %v", resourceType, options)),
		},
	)

	resp.Error = function.FuncErrorFromDiags(ctx, diags)
	if resp.Error != nil {
		return
	}

	resp.Error = resp.Result.Set(ctx, &generateObj)
}
