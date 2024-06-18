// Copyright (c) State of Geneva (Switzerland)
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ABXActionModel describes the resource data model.
type ABXActionModel struct {
	Id           types.String `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	Description  types.String `tfsdk:"description"`
	FAASProvider types.String `tfsdk:"faas_provider"`
	Type         types.String `tfsdk:"type"`

	RuntimeName    types.String `tfsdk:"runtime_name"`
	RuntimeVersion types.String `tfsdk:"runtime_version"`
	MemoryInMB     types.Int64  `tfsdk:"memory_in_mb"`
	TimeoutSeconds types.Int64  `tfsdk:"timeout_seconds"`
	Entrypoint     types.String `tfsdk:"entrypoint"`
	Dependencies   types.List   `tfsdk:"dependencies"`
	// Constants types.List[String] `tfsdk:"constants"`
	// Secrets types.List[String] `tfsdk:"secrets"`

	Source types.String `tfsdk:"source"`

	ProjectId types.String `tfsdk:"project_id"`
	OrgId     types.String `tfsdk:"org_id"`
}

// ABXActionAPIModel describes the resource API model.
type ABXActionAPIModel struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	FAASProvider string `json:"provider"`
	Type         string `json:"actionType"`

	RuntimeName    string            `json:"runtime"`
	RuntimeVersion string            `json:"runtimeVersion"`
	MemoryInMB     int64             `json:"memoryInMB"`
	TimeoutSeconds int64             `json:"timeoutSeconds"`
	Entrypoint     string            `json:"entrypoint"`
	Dependencies   string            `json:"dependencies"`
	Inputs         map[string]string `json:"inputs"`

	Source string `json:"source"`

	ProjectId string `json:"projectId"`
	OrgId     string `json:"orgId"`
}

func (self *ABXActionModel) FromAPI(
	ctx context.Context,
	raw ABXActionAPIModel,
) diag.Diagnostics {

	// https://go.dev/blog/maps
	// inputs := map[string]string{}
	// for key, value := range self.Constants.Elemets() {
	//     inputs["secret:"+key] = value
	// }
	// for key, value := range self.Secrets.Elements() {
	//     inputs["psecret:"+key] = value
	// }

	dependencies, diags := types.ListValueFrom(
		ctx,
		types.StringType,
		SkipEmpty(strings.Split(raw.Dependencies, "\n")),
	)

	self.Id = types.StringValue(raw.Id)
	self.Name = types.StringValue(raw.Name)
	self.Description = types.StringValue(raw.Description)
	self.FAASProvider = types.StringValue(strings.Replace(raw.FAASProvider, "", "auto", 1))
	self.RuntimeName = types.StringValue(raw.RuntimeName)
	self.RuntimeVersion = types.StringValue(raw.RuntimeVersion)
	self.MemoryInMB = types.Int64Value(raw.MemoryInMB)
	self.TimeoutSeconds = types.Int64Value(raw.TimeoutSeconds)
	self.Entrypoint = types.StringValue(raw.Entrypoint)
	self.Dependencies = dependencies
	self.Source = types.StringValue(CleanString(raw.Source))
	self.ProjectId = types.StringValue(raw.ProjectId)
	self.OrgId = types.StringValue(raw.OrgId)

	return diags
}

func (self *ABXActionModel) ToAPI(ctx context.Context) (ABXActionAPIModel, diag.Diagnostics) {

	var diags diag.Diagnostics

	// https://developer.hashicorp.com/terraform/plugin/framework/handling-data/types/list
	if self.Dependencies.IsNull() || self.Dependencies.IsUnknown() {
		diags.AddError(
			"Configuration error",
			fmt.Sprintf(
				"Unable to manage subscription %s, project_ids is either null or unknown",
				self.Id.ValueString()))
		return ABXActionAPIModel{}, diags
	}

	dependencies := make([]string, 0, len(self.Dependencies.Elements()))
	diags = self.Dependencies.ElementsAs(ctx, &dependencies, false)
	if diags.HasError() {
		return ABXActionAPIModel{}, diags
	}

	// https://go.dev/blog/maps
	// inputs := map[string]string{}
	// for key, value := range self.Constants.Elemets() {
	//     inputs["secret:"+key] = value
	// }
	// for key, value := range self.Secrets.Elements() {
	//     inputs["psecret:"+key] = value
	// }

	return ABXActionAPIModel{
		Name:           self.Name.ValueString(),
		Description:    self.Description.ValueString(),
		FAASProvider:   strings.Replace(self.FAASProvider.ValueString(), "auto", "", 1),
		Type:           self.Type.ValueString(),
		RuntimeName:    self.RuntimeName.ValueString(),
		RuntimeVersion: self.RuntimeVersion.ValueString(),
		MemoryInMB:     self.MemoryInMB.ValueInt64(),
		TimeoutSeconds: self.TimeoutSeconds.ValueInt64(),
		Entrypoint:     self.Entrypoint.ValueString(),
		Dependencies:   strings.Join(SkipEmpty(dependencies), "\n"),
		Inputs:         map[string]string{}, // FIXME
		Source:         CleanString(self.Source.ValueString()),
		ProjectId:      self.ProjectId.ValueString(),
	}, diags
}

// ABXConstantModel describes the resource data model.
type ABXConstantModel struct {
	Id        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	Value     types.String `tfsdk:"value"`
	Encrypted types.Bool   `tfsdk:"encrypted"`
	OrgId     types.String `tfsdk:"org_id"`
}

// ABXConstantAPIModel describes the resource API model.
type ABXConstantAPIModel struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Value         string `json:"value"`
	Encrypted     bool   `json:"encrypted"`
	OrgId         string `json:"orgId"`
	CreatedMillis uint64 `json:"createdMillis"`
}

func (self *ABXConstantModel) FromAPI(
	ctx context.Context,
	raw ABXConstantAPIModel,
) diag.Diagnostics {
	self.Id = types.StringValue(raw.Id)
	self.Name = types.StringValue(raw.Name)
	self.Value = types.StringValue(raw.Value)
	self.Encrypted = types.BoolValue(raw.Encrypted)
	self.OrgId = types.StringValue(raw.OrgId)
	return diag.Diagnostics{}
}

func (self *ABXConstantModel) ToAPI() ABXConstantAPIModel {
	return ABXConstantAPIModel{
		Name:      self.Name.ValueString(),
		Value:     self.Value.ValueString(),
		Encrypted: self.Encrypted.ValueBool(),
	}
}

// ABXSecretModel describes the resource data model.
type ABXSecretModel struct {
	Id        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	Value     types.String `tfsdk:"value"`
	Encrypted types.Bool   `tfsdk:"encrypted"`
	OrgId     types.String `tfsdk:"org_id"`
}

// ABXSecretAPIModel describes the resource API model.
type ABXSecretAPIModel struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Value         string `json:"value"`
	Encrypted     bool   `json:"encrypted"`
	OrgId         string `json:"orgId"`
	CreatedMillis uint64 `json:"createdMillis"`
}

func (self *ABXSecretModel) FromAPI(
	ctx context.Context,
	raw ABXSecretAPIModel,
) diag.Diagnostics {
	self.Id = types.StringValue(raw.Id)
	self.Name = types.StringValue(raw.Name)
	// The value is returned with the following '*****' awesome :)
	// self.Value = types.StringValue(raw.Value)
	self.Encrypted = types.BoolValue(raw.Encrypted)
	self.OrgId = types.StringValue(raw.OrgId)
	return diag.Diagnostics{}
}

func (self *ABXSecretModel) ToAPI() ABXSecretAPIModel {
	return ABXSecretAPIModel{
		Name:      self.Name.ValueString(),
		Value:     self.Value.ValueString(),
		Encrypted: self.Encrypted.ValueBool(),
	}
}

// CatalogTypeModel describes the catalog type model.
type CatalogTypeModel struct {
	Id        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	BaseURI   types.String `tfsdk:"base_uri"`
	CreatedAt types.String `tfsdk:"created_at"`
	CreatedBy types.String `tfsdk:"created_by"`
	IconId    types.String `tfsdk:"icon_id"`
}

// CatalogTypeAPIModel describes the catalog type API model.
type CatalogTypeAPIModel struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	BaseURI   string `json:"baseUri"`
	CreatedAt string `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
	IconId    string `json:"iconId"`
}

func (self *CatalogTypeModel) FromAPI(
	ctx context.Context,
	raw CatalogTypeAPIModel,
) diag.Diagnostics {
	self.Id = types.StringValue(raw.Id)
	self.Name = types.StringValue(raw.Name)
	self.BaseURI = types.StringValue(raw.BaseURI)
	self.CreatedAt = types.StringValue(raw.CreatedAt)
	self.CreatedBy = types.StringValue(raw.CreatedBy)
	self.IconId = types.StringValue(raw.IconId)
	return diag.Diagnostics{}
}

// IconModel describes the resource data model.
type IconModel struct {
	Id      types.String `tfsdk:"id"`
	Content types.String `tfsdk:"content"`
}

// SubscriptionModel describes the resource data model.
type SubscriptionModel struct {
	Id                  types.String `tfsdk:"id"`
	Name                types.String `tfsdk:"name"`
	Description         types.String `tfsdk:"description"`
	Type                types.String `tfsdk:"type"`
	RunnableType        types.String `tfsdk:"runnable_type"`
	RunnableId          types.String `tfsdk:"runnable_id"`
	RecoverRunnableType types.String `tfsdk:"recover_runnable_type"`
	RecoverRunnableId   types.String `tfsdk:"recover_runnable_id"`
	EventTopicId        types.String `tfsdk:"event_topic_id"`

	ProjectIds types.Set `tfsdk:"project_ids"`

	Blocking   types.Bool   `tfsdk:"blocking"`
	Broadcast  types.Bool   `tfsdk:"broadcast"`
	Contextual types.Bool   `tfsdk:"contextual"`
	Criteria   types.String `tfsdk:"criteria"`
	Disabled   types.Bool   `tfsdk:"disabled"`
	Priority   types.Int64  `tfsdk:"priority"`
	System     types.Bool   `tfsdk:"system"`
	Timeout    types.Int64  `tfsdk:"timeout"`

	OrgId        types.String `tfsdk:"org_id"`
	OwnerId      types.String `tfsdk:"owner_id"`
	SubscriberId types.String `tfsdk:"subscriber_id"`
}

// SubscriptionAPIModel describes the resource API model.
type SubscriptionAPIModel struct {
	Id                  string `json:"id"`
	Name                string `json:"name"`
	Description         string `json:"description"`
	Type                string `json:"type"`
	RunnableType        string `json:"runnableType"`
	RunnableId          string `json:"runnableId"`
	RecoverRunnableType string `json:"recoverRunnableType"`
	RecoverRunnableId   string `json:"recoverRunnableId"`
	EventTopicId        string `json:"eventTopicId"`

	Constraints map[string][]string `json:"constraints"`

	Blocking   bool   `json:"blocking"`
	Broadcast  bool   `json:"broadcast"`
	Contextual bool   `json:"contextual"`
	Criteria   string `json:"criteria"`
	Disabled   bool   `json:"disabled"`
	Priority   int64  `json:"priority"`
	System     bool   `json:"system"`
	Timeout    int64  `json:"timeout"`

	OrgId        string `json:"orgId"`
	OwnerId      string `json:"ownerId"`
	SubscriberId string `json:"subscriberId"`
}

func (self *SubscriptionModel) GenerateId() {
	if len(self.Id.ValueString()) == 0 {
		self.Id = types.StringValue(uuid.New().String())
	}
}

func (self *SubscriptionModel) FromAPI(
	ctx context.Context,
	raw SubscriptionAPIModel,
) diag.Diagnostics {
	projectIds, diags := types.SetValueFrom(ctx, types.StringType, raw.Constraints["projectId"])

	self.Id = types.StringValue(raw.Id)
	self.Name = types.StringValue(raw.Name)
	self.Description = types.StringValue(raw.Description)
	self.Type = types.StringValue(raw.Type)
	self.RunnableType = types.StringValue(raw.RunnableType)
	self.RunnableId = types.StringValue(raw.RunnableId)
	self.RecoverRunnableType = StringOrNullValue(raw.RecoverRunnableType)
	self.RecoverRunnableId = StringOrNullValue(raw.RecoverRunnableId)
	self.EventTopicId = types.StringValue(raw.EventTopicId)
	self.ProjectIds = projectIds
	self.Blocking = types.BoolValue(raw.Blocking)
	self.Broadcast = types.BoolValue(raw.Broadcast)
	self.Contextual = types.BoolValue(raw.Contextual)
	self.Criteria = types.StringValue(raw.Criteria)
	self.Disabled = types.BoolValue(raw.Disabled)
	self.Priority = types.Int64Value(raw.Priority)
	self.System = types.BoolValue(raw.System)
	self.Timeout = types.Int64Value(raw.Timeout)
	self.OrgId = types.StringValue(raw.OrgId)
	self.OwnerId = types.StringValue(raw.OwnerId)
	self.SubscriberId = types.StringValue(raw.SubscriberId)

	return diags
}

func (self *SubscriptionModel) ToAPI(ctx context.Context) (SubscriptionAPIModel, diag.Diagnostics) {

	var diags diag.Diagnostics

	// https://developer.hashicorp.com/terraform/plugin/framework/handling-data/types/set
	if self.ProjectIds.IsNull() || self.ProjectIds.IsUnknown() {
		diags.AddError(
			"Configuration error",
			fmt.Sprintf(
				"Unable to manage subscription %s, project_ids is either null or unknown",
				self.Id.ValueString()))
		return SubscriptionAPIModel{}, diags
	}

	projectIds := make([]string, 0, len(self.ProjectIds.Elements()))
	diags = self.ProjectIds.ElementsAs(ctx, &projectIds, false)
	if diags.HasError() {
		return SubscriptionAPIModel{}, diags
	}

	return SubscriptionAPIModel{
		Id:                  self.Id.ValueString(),
		Name:                self.Name.ValueString(),
		Description:         self.Description.ValueString(),
		Type:                self.Type.ValueString(),
		RunnableType:        self.RunnableType.ValueString(),
		RunnableId:          self.RunnableId.ValueString(),
		RecoverRunnableType: self.RecoverRunnableType.ValueString(),
		RecoverRunnableId:   self.RecoverRunnableId.ValueString(),
		EventTopicId:        self.EventTopicId.ValueString(),
		Constraints:         map[string][]string{"projectId": projectIds},
		Blocking:            self.Blocking.ValueBool(),
		Contextual:          self.Contextual.ValueBool(),
		Criteria:            self.Criteria.ValueString(),
		Disabled:            self.Disabled.ValueBool(),
		Priority:            self.Priority.ValueInt64(),
		Timeout:             self.Timeout.ValueInt64(),
	}, diags
}
