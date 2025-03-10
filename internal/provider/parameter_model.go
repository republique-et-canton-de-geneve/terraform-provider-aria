// Copyright (c) State of Geneva (Switzerland)
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ParameterModel describes the resource data model.
type ParameterModel struct {
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Type        types.String `tfsdk:"type"`
}

// ParameterAPIModel describes the resource API model.
type ParameterAPIModel struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

// Used to convert structure to a types.Object.
func (self ParameterModel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":        types.StringType,
		"description": types.StringType,
		"type":        types.StringType,
	}
}

func (self ParameterModel) String() string {
	return fmt.Sprintf(
		"Input Parameter %s (%s)",
		self.Name.ValueString(),
		self.Type.ValueString())
}

func (self *ParameterModel) FromAPI(raw ParameterAPIModel) {
	self.Name = types.StringValue(raw.Name)
	self.Description = types.StringValue(raw.Description)
	self.Type = types.StringValue(raw.Type)
}

func (self ParameterModel) ToAPI() ParameterAPIModel {
	return ParameterAPIModel{
		Name:        self.Name.ValueString(),
		Description: CleanString(self.Description.ValueString()),
		Type:        self.Type.ValueString(),
	}
}

// Utils -------------------------------------------------------------------------------------------

func ParameterModelListFromAPI(
	ctx context.Context,
	parametersRaw []ParameterAPIModel,
) (types.List, diag.Diagnostics) {
	// Convert input parameters from raw
	parameters := []ParameterModel{}
	for _, parameterRaw := range parametersRaw {
		parameter := ParameterModel{}
		parameter.FromAPI(parameterRaw)
		parameters = append(parameters, parameter)
	}

	// Store inputs parameters to list value
	parameterAttrs := types.ObjectType{AttrTypes: ParameterModel{}.AttributeTypes()}
	return types.ListValueFrom(ctx, parameterAttrs, parameters)
}

func ParameterModelListToAPI(
	ctx context.Context,
	parametersList types.List,
	name string,
) ([]ParameterAPIModel, diag.Diagnostics) {
	diags := diag.Diagnostics{}
	parametersRaw := []ParameterAPIModel{}

	// https://developer.hashicorp.com/terraform/plugin/framework/handling-data/types/list
	if parametersList.IsNull() || parametersList.IsUnknown() {
		diags.AddError(
			"Configuration error",
			fmt.Sprintf("Unable to manage %s is either null or unknown", name))
		return parametersRaw, diags
	}

	// Extract input parameters from list value and then convert to raw
	parameters := make([]ParameterModel, 0, len(parametersList.Elements()))
	diags.Append(parametersList.ElementsAs(ctx, &parameters, false)...)
	if !diags.HasError() {
		for _, parameter := range parameters {
			parametersRaw = append(parametersRaw, parameter.ToAPI())
		}
	}

	return parametersRaw, diags
}
