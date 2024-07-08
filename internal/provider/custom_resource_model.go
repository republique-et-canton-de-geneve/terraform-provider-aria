// Copyright (c) State of Geneva (Switzerland)
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// CustomResourceModel describes the resource data model.
type CustomResourceModel struct {
	Id           types.String `tfsdk:"id"`
	DisplayName  types.String `tfsdk:"display_name"`
	Description  types.String `tfsdk:"description"`
	ResourceType types.String `tfsdk:"resource_type"`
	SchemaType   types.String `tfsdk:"schema_type"`
	Status       types.String `tfsdk:"status"`

	Properties []PropertyModel `tfsdk:"properties"`

	Create CustomResourceActionModel `tfsdk:"create"`
	Read   CustomResourceActionModel `tfsdk:"read"`
	Update CustomResourceActionModel `tfsdk:"update"`
	Delete CustomResourceActionModel `tfsdk:"delete"`

	ProjectId types.String `tfsdk:"project_id"`
	OrgId     types.String `tfsdk:"org_id"`
}

// CustomResourcePropertiesAPIModel describes the resource API model.
type CustomResourcePropertiesAPIModel struct {
	Properties PropertiesAPIModel `json:"properties"`
}

// CustomResourceAPIModel describes the resource API model.
type CustomResourceAPIModel struct {
	Id           string `json:"id"`
	DisplayName  string `json:"displayName"`
	Description  string `json:"description"`
	ResourceType string `json:"resourceType"`
	SchemaType   string `json:"schemaType"`
	Status       string `json:"status"`

	Properties CustomResourcePropertiesAPIModel `json:"properties"`

	MainActions map[string]CustomResourceActionAPIModel `json:"mainActions"`

	ProjectId string `json:"projectId"`
	OrgId     string `json:"orgId"`
}

func (self *CustomResourceModel) String() string {
	return fmt.Sprintf(
		"ABX Custom Resource %s (%s)",
		self.Id.ValueString(),
		self.DisplayName.ValueString())
}

func (self *CustomResourceModel) FromAPI(
	ctx context.Context,
	raw CustomResourceAPIModel,
) diag.Diagnostics {

	diags := diag.Diagnostics{}

	self.Id = types.StringValue(raw.Id)
	self.DisplayName = types.StringValue(raw.DisplayName)
	self.Description = types.StringValue(raw.Description)
	self.ResourceType = types.StringValue(raw.ResourceType)
	self.SchemaType = types.StringValue(raw.SchemaType)
	self.Status = types.StringValue(raw.Status)
	self.ProjectId = types.StringValue(raw.ProjectId)
	self.OrgId = types.StringValue(raw.OrgId)

	self.Properties = []PropertyModel{}
	for _, propertyItem := range raw.Properties.Properties.Items() {
		property := PropertyModel{}
		diags.Append(property.FromAPI(ctx, propertyItem.Name, propertyItem.Property)...)
		self.Properties = append(self.Properties, property)
	}

	self.Create = CustomResourceActionModel{}
	diags.Append(self.Create.FromAPI(ctx, raw.MainActions["create"])...)

	self.Read = CustomResourceActionModel{}
	diags.Append(self.Read.FromAPI(ctx, raw.MainActions["read"])...)

	self.Update = CustomResourceActionModel{}
	diags.Append(self.Update.FromAPI(ctx, raw.MainActions["update"])...)

	self.Delete = CustomResourceActionModel{}
	diags.Append(self.Delete.FromAPI(ctx, raw.MainActions["delete"])...)

	return diags
}

func (self *CustomResourceModel) ToAPI(
	ctx context.Context,
) (CustomResourceAPIModel, diag.Diagnostics) {

	diags := diag.Diagnostics{}

	propertiesRaw := PropertiesAPIModel{}
	propertiesRaw.Init()
	for _, property := range self.Properties {
		propertyName, propertyRaw, propertyDiags := property.ToAPI(ctx)
		propertiesRaw.Set(propertyName, propertyRaw)
		diags.Append(propertyDiags...)
	}

	createRaw, createDiags := self.Create.ToAPI(ctx)
	diags.Append(createDiags...)

	readRaw, readDiags := self.Read.ToAPI(ctx)
	diags.Append(readDiags...)

	updateRaw, updateDiags := self.Update.ToAPI(ctx)
	diags.Append(updateDiags...)

	deleteRaw, deleteDiags := self.Delete.ToAPI(ctx)
	diags.Append(deleteDiags...)

	raw := CustomResourceAPIModel{
		DisplayName:  self.DisplayName.ValueString(),
		Description:  CleanString(self.Description.ValueString()),
		ResourceType: self.ResourceType.ValueString(),
		SchemaType:   self.SchemaType.ValueString(),
		Status:       self.Status.ValueString(),
		ProjectId:    self.ProjectId.ValueString(),
		OrgId:        self.OrgId.ValueString(),
		Properties: CustomResourcePropertiesAPIModel{
			Properties: propertiesRaw,
		},
		MainActions: map[string]CustomResourceActionAPIModel{
			"create": createRaw,
			"read":   readRaw,
			"update": updateRaw,
			"delete": deleteRaw,
		},
	}

	// When updating resource
	if !self.Id.IsNull() {
		raw.Id = self.Id.ValueString()
	}

	return raw, diags
}
