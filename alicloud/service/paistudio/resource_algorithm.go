package paistudio

import (
	"context"
	"fmt"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/errs"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/provider/fwadapt"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const algorithmTypeName = "alicloud_paistudio_algorithm"

// PaiStudio is an independent POP, distinct from AIWorkSpace (2021-02-04)
// and PAIFlow. Its RESTful/ROA API version is 2022-01-12. Exported so the
// external _test package can build a standalone client for fixtures.
const (
	PaiStudioProduct = "PaiStudio"
	PaiStudioVersion = "2022-01-12"
)

var (
	_ resource.Resource                = &algorithmResource{}
	_ resource.ResourceWithImportState = &algorithmResource{}
)

func NewAlgorithmResource() resource.Resource {
	return &algorithmResource{}
}

type algorithmResource struct {
	fwadapt.ResourceBase
}

type algorithmResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	AlgorithmName        types.String `tfsdk:"algorithm_name"`
	WorkspaceId          types.String `tfsdk:"workspace_id"`
	DisplayName          types.String `tfsdk:"display_name"`
	AlgorithmDescription types.String `tfsdk:"algorithm_description"`
	AlgorithmProvider    types.String `tfsdk:"algorithm_provider"`
	UserId               types.String `tfsdk:"user_id"`
	TenantId             types.String `tfsdk:"tenant_id"`
	GmtCreateTime        types.String `tfsdk:"gmt_create_time"`
	GmtModifiedTime      types.String `tfsdk:"gmt_modified_time"`
}

func (r *algorithmResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = algorithmTypeName
}

func (r *algorithmResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Provides a PaiStudio Algorithm resource.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the Algorithm.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"algorithm_name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the Algorithm. The UpdateAlgorithm API only accepts the description and display name, so changing the name recreates the Algorithm.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"workspace_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The ID of the workspace that the Algorithm belongs to. If not set, the API assigns the default workspace; changing it recreates the Algorithm.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplace(),
				},
			},
			"display_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The display name of the Algorithm.",
			},
			"algorithm_description": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The description of the Algorithm.",
			},
			"algorithm_provider": schema.StringAttribute{
				Computed:    true,
				Description: "The provider of the Algorithm.",
			},
			"user_id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the user who owns the Algorithm.",
			},
			"tenant_id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the tenant that the Algorithm belongs to.",
			},
			"gmt_create_time": schema.StringAttribute{
				Computed:    true,
				Description: "The time when the Algorithm was created.",
			},
			"gmt_modified_time": schema.StringAttribute{
				Computed:    true,
				Description: "The time when the Algorithm was last modified.",
			},
		},
	}
}

func (r *algorithmResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan algorithmResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if plan.AlgorithmName.IsNull() || plan.AlgorithmName.IsUnknown() || plan.AlgorithmName.ValueString() == "" {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Creating %s: missing algorithm_name", algorithmTypeName),
			"The PaiStudio CreateAlgorithm API requires a non-empty AlgorithmName.",
		)
		return
	}

	body := algorithmCreateBody(plan)
	response, err := r.Client().RoaPost(PaiStudioProduct, PaiStudioVersion, algorithmPath(""), nil, nil, body, true)
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Creating %s: calling CreateAlgorithm", algorithmTypeName),
			errs.WrapErrorf(err, "Creating %s: calling CreateAlgorithm", algorithmTypeName).Error(),
		)
		return
	}
	algorithmId, _ := response["AlgorithmId"].(string)
	if algorithmId == "" {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Creating %s: CreateAlgorithm returned no AlgorithmId", algorithmTypeName),
			fmt.Sprintf("CreateAlgorithm returned %#v; the Algorithm may have been created without a usable ID and needs manual cleanup", response),
		)
		return
	}

	// Persist the created resource BEFORE reading. The id is only knowable
	// from the create response, so the anchor state (id + config, computed
	// fields null) keeps a successfully created Algorithm tracked even when
	// the first read fails or the object is transiently invisible.
	anchor := plan
	anchor.Id = types.StringValue(algorithmId)
	anchor.WorkspaceId = knownOrNull(plan.WorkspaceId)
	anchor.DisplayName = knownOrNull(plan.DisplayName)
	anchor.AlgorithmDescription = knownOrNull(plan.AlgorithmDescription)
	anchor.AlgorithmProvider = types.StringNull()
	anchor.UserId = types.StringNull()
	anchor.TenantId = types.StringNull()
	anchor.GmtCreateTime = types.StringNull()
	anchor.GmtModifiedTime = types.StringNull()
	resp.Diagnostics.Append(resp.State.Set(ctx, anchor)...)
	if resp.Diagnostics.HasError() {
		return
	}

	object, err := getAlgorithm(r.Client(), algorithmId)
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Reading %s: calling GetAlgorithm", algorithmTypeName),
			errs.WrapErrorf(err, "Reading %s: calling GetAlgorithm", algorithmTypeName).Error(),
		)
		return
	}
	if object == nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Creating %s: calling GetAlgorithm", algorithmTypeName),
			fmt.Sprintf("The algorithm %s is not visible via GetAlgorithm right after creation; the resource is persisted in state for the next refresh", algorithmId),
		)
		return
	}
	if err := validateAlgorithmResponse(object, algorithmId); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Creating %s: validating GetAlgorithm response", algorithmTypeName),
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, algorithmState(object))...)
}

func (r *algorithmResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state algorithmResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	object, err := getAlgorithm(r.Client(), state.Id.ValueString())
	if err != nil {
		// A transient read error must not drop the existing state; return the
		// error without Set/RemoveResource so the prior state is preserved.
		resp.Diagnostics.AddError(
			fmt.Sprintf("Reading %s: calling GetAlgorithm", algorithmTypeName),
			errs.WrapErrorf(err, "Reading %s: calling GetAlgorithm", algorithmTypeName).Error(),
		)
		return
	}
	if object == nil {
		resp.State.RemoveResource(ctx)
		return
	}
	// Validate the response identity before flattening onto state; a
	// mismatched response must not overwrite the resource's identity. On such
	// a malformed response, keep the prior state and surface the error.
	if err := validateAlgorithmResponse(object, state.Id.ValueString()); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Reading %s: validating GetAlgorithm response", algorithmTypeName),
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, algorithmState(object))...)
}

func (r *algorithmResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state algorithmResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	body := algorithmUpdateBody(plan)
	if _, err := r.Client().RoaPut(PaiStudioProduct, PaiStudioVersion, algorithmPath(state.Id.ValueString()), nil, nil, body, true); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Updating %s: calling UpdateAlgorithm", algorithmTypeName),
			errs.WrapErrorf(err, "Updating %s: calling UpdateAlgorithm", algorithmTypeName).Error(),
		)
		return
	}

	object, err := getAlgorithm(r.Client(), state.Id.ValueString())
	if err != nil {
		// Keep the prior state on a read failure after update.
		resp.Diagnostics.AddError(
			fmt.Sprintf("Reading %s: calling GetAlgorithm", algorithmTypeName),
			errs.WrapErrorf(err, "Reading %s: calling GetAlgorithm", algorithmTypeName).Error(),
		)
		return
	}
	if object == nil {
		resp.State.RemoveResource(ctx)
		return
	}
	if err := validateAlgorithmResponse(object, state.Id.ValueString()); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Updating %s: validating GetAlgorithm response", algorithmTypeName),
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, algorithmState(object))...)
}

func (r *algorithmResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state algorithmResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.Client().RoaDelete(PaiStudioProduct, PaiStudioVersion, algorithmPath(state.Id.ValueString()), nil, nil, nil, true)
	if err != nil && !isAlgorithmNotExist(err) {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Deleting %s: calling DeleteAlgorithm", algorithmTypeName),
			errs.WrapErrorf(err, "Deleting %s: calling DeleteAlgorithm", algorithmTypeName).Error(),
		)
		return
	}
}

func (r *algorithmResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if req.ID == "" {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Importing %s", algorithmTypeName),
			"The import ID must be the AlgorithmId; got an empty string",
		)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, algorithmResourceModel{
		Id: types.StringValue(req.ID),
	})...)
}

// algorithmPath builds the ROA path that addresses a single Algorithm; the
// create path has no trailing id. The id lives in the path, not the query.
func algorithmPath(algorithmId string) string {
	if algorithmId == "" {
		return "/api/v1/algorithms"
	}
	return fmt.Sprintf("/api/v1/algorithms/%s", algorithmId)
}

// algorithmCreateBody maps the plan onto the CreateAlgorithm form parameters.
// The body is sent as a JSON object: tea-roa serializes the RoaPost body as
// JSON, matching how the production SAE resource calls form-data ROA APIs.
// Optional parameters are only sent when the user set them.
func algorithmCreateBody(plan algorithmResourceModel) map[string]interface{} {
	body := map[string]interface{}{
		"AlgorithmName": plan.AlgorithmName.ValueString(),
	}
	if v := knownOrNull(plan.WorkspaceId); !v.IsNull() {
		body["WorkspaceId"] = v.ValueString()
	}
	if v := knownOrNull(plan.DisplayName); !v.IsNull() {
		body["DisplayName"] = v.ValueString()
	}
	if v := knownOrNull(plan.AlgorithmDescription); !v.IsNull() {
		body["AlgorithmDescription"] = v.ValueString()
	}
	return body
}

// algorithmUpdateBody maps the plan onto the UpdateAlgorithm parameters. The
// API contract only accepts AlgorithmDescription and DisplayName; algorithm
// name and workspace are immutable (RequiresReplace in the schema).
func algorithmUpdateBody(plan algorithmResourceModel) map[string]interface{} {
	body := map[string]interface{}{}
	if v := knownOrNull(plan.DisplayName); !v.IsNull() {
		body["DisplayName"] = v.ValueString()
	}
	if v := knownOrNull(plan.AlgorithmDescription); !v.IsNull() {
		body["AlgorithmDescription"] = v.ValueString()
	}
	return body
}

// getAlgorithm queries GetAlgorithm and returns the algorithm object, or nil
// when it no longer exists (NoSuchObject).
func getAlgorithm(client *connectivity.AliyunClient, algorithmId string) (map[string]interface{}, error) {
	response, err := client.RoaGet(PaiStudioProduct, PaiStudioVersion, algorithmPath(algorithmId), nil, nil, nil)
	if err != nil {
		if isAlgorithmNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	return response, nil
}

// algorithmState maps a GetAlgorithm response onto the resource model. Empty
// strings become null: optional fields the API omits read back as absent
// rather than "".
func algorithmState(object map[string]interface{}) algorithmResourceModel {
	str := func(key string) types.String {
		if v, ok := object[key].(string); ok && v != "" {
			return types.StringValue(v)
		}
		return types.StringNull()
	}
	return algorithmResourceModel{
		Id:                   str("AlgorithmId"),
		AlgorithmName:        str("AlgorithmName"),
		WorkspaceId:          str("WorkspaceId"),
		DisplayName:          str("DisplayName"),
		AlgorithmDescription: str("AlgorithmDescription"),
		AlgorithmProvider:    str("AlgorithmProvider"),
		UserId:               str("UserId"),
		TenantId:             str("TenantId"),
		GmtCreateTime:        str("GmtCreateTime"),
		GmtModifiedTime:      str("GmtModifiedTime"),
	}
}

// knownOrNull collapses an unknown or null value to null so a value can never
// be written to state while the plan still shows unknown.
func knownOrNull(value types.String) types.String {
	if value.IsNull() || value.IsUnknown() {
		return types.StringNull()
	}
	return value
}

// isAlgorithmNotExist reports whether err is the PaiStudio not-found error.
// The ROA call surfaces the product error contract on a *tea.SDKError whose
// Code is NoSuchObject; that is the only signal used here. A loose substring
// match is intentionally avoided so a non-NotFound message that merely
// contains the word "NoSuchObject" cannot be misread as not-found and
// silently drop the resource from state.
func isAlgorithmNotExist(err error) bool {
	if err == nil {
		return false
	}
	if e, ok := err.(*tea.SDKError); ok && e.Code != nil && *e.Code == "NoSuchObject" {
		return true
	}
	return false
}

// validateAlgorithmResponse guards the GetAlgorithm response before it
// reaches state. A missing or wrong-typed AlgorithmId would compose an empty
// id, and a mismatched response could overwrite the resource's identity. On
// any such failure the caller must keep the prior state and surface this
// error rather than writing a corrupt object.
func validateAlgorithmResponse(object map[string]interface{}, expectedAlgorithmId string) error {
	if object == nil {
		return fmt.Errorf("GetAlgorithm returned an empty response for %s", expectedAlgorithmId)
	}
	algorithmId, ok := object["AlgorithmId"].(string)
	if !ok || algorithmId == "" {
		return fmt.Errorf("GetAlgorithm response is missing the string AlgorithmId for %s: %#v", expectedAlgorithmId, object["AlgorithmId"])
	}
	if algorithmId != expectedAlgorithmId {
		return fmt.Errorf("GetAlgorithm response identity %s does not match the requested %s", algorithmId, expectedAlgorithmId)
	}
	return nil
}
