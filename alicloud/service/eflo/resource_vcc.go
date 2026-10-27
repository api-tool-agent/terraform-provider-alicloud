// Package eflo. This file is generated automatically. Please do not modify it manually, thank you!
// Nested attributes require Terraform Plugin Protocol 6.
package eflo

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/errs"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/provider/fwadapt"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const vccTypeName = "alicloud_eflo_vcc"

var (
	_ resource.Resource                = &vccResource{}
	_ resource.ResourceWithImportState = &vccResource{}
)

func NewVccResource() resource.Resource {
	return &vccResource{}
}

type vccResource struct {
	fwadapt.ResourceBase
}

type vccResourceModel struct {
	Id                 types.String `tfsdk:"id"`
	FwTimeouts         types.Object `tfsdk:"timeouts"`
	AccessCouldService types.Bool   `tfsdk:"access_could_service"`
	Bandwidth          types.Int64  `tfsdk:"bandwidth"`
	BgpAsn             types.Int64  `tfsdk:"bgp_asn"`
	BgpCidr            types.String `tfsdk:"bgp_cidr"`
	CenId              types.String `tfsdk:"cen_id"`
	CenOwnerId         types.String `tfsdk:"cen_owner_id"`
	ConnectionType     types.String `tfsdk:"connection_type"`
	CreateTime         types.String `tfsdk:"create_time"`
	Description        types.String `tfsdk:"description"`
	GmtModified        types.String `tfsdk:"gmt_modified"`
	Message            types.String `tfsdk:"message"`
	ResourceGroupId    types.String `tfsdk:"resource_group_id"`
	Spec               types.String `tfsdk:"spec"`
	Status             types.String `tfsdk:"status"`
	Tags               types.Map    `tfsdk:"tags"`
	VSwitchId          types.String `tfsdk:"vswitch_id"`
	VccName            types.String `tfsdk:"vcc_name"`
	VpcId              types.String `tfsdk:"vpc_id"`
	ZoneId             types.String `tfsdk:"zone_id"`
}

func (r *vccResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = vccTypeName
}

func (r *vccResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Blocks: map[string]schema.Block{"timeouts": fwTimeoutBlock()},
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"access_could_service": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
				Description: "Start access to cloud service, optional value:\n\n-**true**: enable access to cloud services\n-**false**: does not enable access to cloud services",
			},
			"bandwidth": schema.Int64Attribute{
				Required:    true,
				Description: "Bandwidth, in Mbps. The minimum value is 1000, representing 1Gbps bandwidth; The maximum value is 400000, representing 400Gbps bandwidth.\n\n> 1Gbps = 1000Mbps",
			},
			"bgp_asn": schema.Int64Attribute{
				Optional:    true,
				Description: "Bgp as number",
			},
			"bgp_cidr": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "BGP network segment.",
			},
			"cen_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "CEN",
			},
			"cen_owner_id": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Account to which cen belongs",
			},
			"connection_type": schema.StringAttribute{
				Required:    true,
				Description: "Connection method, optional value:\n-**VPC**\n-**CENTR**",
			},
			"create_time": schema.StringAttribute{
				Computed:    true,
				Description: "Creation time",
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Description: "Description",
			},
			"gmt_modified": schema.StringAttribute{
				Computed:    true,
				Description: "Modification time",
			},
			"message": schema.StringAttribute{
				Computed:    true,
				Description: "Information returned when the call fails.",
			},
			"resource_group_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Resource group id",
			},
			"spec": schema.StringAttribute{
				Computed:    true,
				Description: "Specifications",
			},
			"status": schema.StringAttribute{
				Computed:    true,
				Description: "Status",
			},
			"tags": schema.MapAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "The tag of the resource",
			},
			"vswitch_id": schema.StringAttribute{
				Required:    true,
				Description: "vSwitch",
			},
			"vcc_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Lingjun Connection Name",
			},
			"vpc_id": schema.StringAttribute{
				Required:    true,
				Description: "VPC",
			},
			"zone_id": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Availability Zone",
			},
		},
	}
}

func (r *vccResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan vccResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	operationTimeout, timeoutDiagnostics := fwTimeoutDuration(plan.FwTimeouts, "create", 22*time.Minute)
	resp.Diagnostics.Append(timeoutDiagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := map[string]interface{}{
		"ClientToken": buildClientToken("CreateVcc"),
	}
	// gap#9: RPC APIs require RegionId as a request param (separate from endpoint resolution);
	// the SDK RpcPost resolves the endpoint by region but does NOT inject RegionId into query/body.
	// v1 sets request["RegionId"] = client.RegionId authoritatively (provider region wins).
	request["RegionId"] = r.Client().RegionId
	createValues := map[string]interface{}{"RegionId": r.Client().RegionId}
	createValues["ResourceType"] = "Vcc"
	if !plan.AccessCouldService.IsNull() && !plan.AccessCouldService.IsUnknown() {
		createValues["AccessCouldService"] = fwValueToInterface(plan.AccessCouldService, map[string]fwWireField{}, false)
	}
	if !plan.Bandwidth.IsNull() && !plan.Bandwidth.IsUnknown() {
		createValues["Bandwidth"] = fwValueToInterface(plan.Bandwidth, map[string]fwWireField{}, false)
	}
	if !plan.BgpAsn.IsNull() && !plan.BgpAsn.IsUnknown() {
		createValues["BgpAsn"] = fwValueToInterface(plan.BgpAsn, map[string]fwWireField{}, false)
	}
	if !plan.BgpCidr.IsNull() && !plan.BgpCidr.IsUnknown() {
		createValues["BgpCidr"] = fwValueToInterface(plan.BgpCidr, map[string]fwWireField{}, false)
	}
	if !plan.CenId.IsNull() && !plan.CenId.IsUnknown() {
		createValues["CenId"] = fwValueToInterface(plan.CenId, map[string]fwWireField{}, false)
	}
	if !plan.CenOwnerId.IsNull() && !plan.CenOwnerId.IsUnknown() {
		createValues["CenOwnerId"] = fwValueToInterface(plan.CenOwnerId, map[string]fwWireField{}, false)
	}
	if !plan.ConnectionType.IsNull() && !plan.ConnectionType.IsUnknown() {
		createValues["ConnectionType"] = fwValueToInterface(plan.ConnectionType, map[string]fwWireField{}, false)
	}
	if !plan.CreateTime.IsNull() && !plan.CreateTime.IsUnknown() {
		createValues["CreateTime"] = fwValueToInterface(plan.CreateTime, map[string]fwWireField{}, false)
	}
	if !plan.Description.IsNull() && !plan.Description.IsUnknown() {
		createValues["Description"] = fwValueToInterface(plan.Description, map[string]fwWireField{}, false)
	}
	if !plan.GmtModified.IsNull() && !plan.GmtModified.IsUnknown() {
		createValues["GmtModified"] = fwValueToInterface(plan.GmtModified, map[string]fwWireField{}, false)
	}
	if !plan.Message.IsNull() && !plan.Message.IsUnknown() {
		createValues["Message"] = fwValueToInterface(plan.Message, map[string]fwWireField{}, false)
	}
	if !plan.ResourceGroupId.IsNull() && !plan.ResourceGroupId.IsUnknown() {
		createValues["ResourceGroupId"] = fwValueToInterface(plan.ResourceGroupId, map[string]fwWireField{}, false)
	}
	if !plan.Spec.IsNull() && !plan.Spec.IsUnknown() {
		createValues["Spec"] = fwValueToInterface(plan.Spec, map[string]fwWireField{}, false)
	}
	if !plan.Status.IsNull() && !plan.Status.IsUnknown() {
		createValues["Status"] = fwValueToInterface(plan.Status, map[string]fwWireField{}, false)
	}
	if !plan.Tags.IsNull() && !plan.Tags.IsUnknown() {
		createValues["Tags"] = fwTagsToAttributes(plan.Tags)
	}
	if !plan.VSwitchId.IsNull() && !plan.VSwitchId.IsUnknown() {
		createValues["VSwitchId"] = fwValueToInterface(plan.VSwitchId, map[string]fwWireField{}, false)
	}
	if !plan.VccName.IsNull() && !plan.VccName.IsUnknown() {
		createValues["VccName"] = fwValueToInterface(plan.VccName, map[string]fwWireField{}, false)
	}
	if !plan.VpcId.IsNull() && !plan.VpcId.IsUnknown() {
		createValues["VpcId"] = fwValueToInterface(plan.VpcId, map[string]fwWireField{}, false)
	}
	if !plan.ZoneId.IsNull() && !plan.ZoneId.IsUnknown() {
		createValues["ZoneId"] = fwValueToInterface(plan.ZoneId, map[string]fwWireField{}, false)
	}
	if err := fwApplyMappings(request, createValues, []fwMapping{
		{Source: "$.VccName", Target: "VccName", Encoding: "", Constant: "", Repeat: false, Values: map[string]string{}},
		{Source: "$.CenId", Target: "CenId", Encoding: "", Constant: "", Repeat: false, Values: map[string]string{}},
		{Source: "$.ResourceGroupId", Target: "ResourceGroupId", Encoding: "", Constant: "", Repeat: false, Values: map[string]string{}},
		{Source: "$.Bandwidth", Target: "Bandwidth", Encoding: "", Constant: "", Repeat: false, Values: map[string]string{}},
		{Source: "$.Description", Target: "Description", Encoding: "", Constant: "", Repeat: false, Values: map[string]string{}},
		{Source: "$.VSwitchId", Target: "VSwitchId", Encoding: "", Constant: "", Repeat: false, Values: map[string]string{}},
		{Source: "$.Tags[*].TagValue", Target: "Tag[*].Value", Encoding: "", Constant: "", Repeat: true, Values: map[string]string{}},
		{Source: "$.BgpCidr", Target: "BgpCidr", Encoding: "", Constant: "", Repeat: false, Values: map[string]string{}},
		{Source: "$.CenOwnerId", Target: "CenOwnerId", Encoding: "", Constant: "", Repeat: false, Values: map[string]string{}},
		{Source: "$.RegionId", Target: "RegionId", Encoding: "", Constant: "", Repeat: false, Values: map[string]string{}},
		{Source: "$.ConnectionType", Target: "ConnectionType", Encoding: "", Constant: "", Repeat: false, Values: map[string]string{}},
		{Source: "$.BgpAsn", Target: "BgpAsn", Encoding: "", Constant: "", Repeat: false, Values: map[string]string{}},
		{Source: "$.VpcId", Target: "VpcId", Encoding: "", Constant: "", Repeat: false, Values: map[string]string{}},
		{Source: "$.Tags[*].TagKey", Target: "Tag[*].Key", Encoding: "", Constant: "", Repeat: true, Values: map[string]string{}},
		{Source: "$.AccessCouldService", Target: "AccessCouldService", Encoding: "", Constant: "", Repeat: false, Values: map[string]string{}},
		{Source: "$.ZoneId", Target: "ZoneId", Encoding: "", Constant: "", Repeat: false, Values: map[string]string{}},
	}, false); err != nil {
		resp.Diagnostics.AddError("Building create request", err.Error())
		return
	}

	response, err := retryRpcPost(r.Client(), "eflo", "2022-05-30", "CreateVcc", request, operationTimeout)
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Creating %s: calling CreateVcc", vccTypeName),
			errs.WrapErrorf(err, "Creating %s: calling CreateVcc", vccTypeName).Error(),
		)
		return
	}

	identityParts := []string{}
	identityValues0 := fwWalkMappedPath(response, fwMappingPath("$.Content.VccId"), nil)
	if len(identityValues0) == 0 || identityValues0[0].Value == nil || fmt.Sprint(identityValues0[0].Value) == "" {
		resp.Diagnostics.AddError("Creating alicloud_eflo_vcc: missing identity", "The create operation did not return or accept every identity component")
		return
	}
	identityParts = append(identityParts, fmt.Sprint(identityValues0[0].Value))
	id := strings.Join(identityParts, ":")
	if id == "" || id == "<nil>" {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Creating %s: calling CreateVcc", vccTypeName),
			fmt.Sprintf("The response carried no VccId: %v", response),
		)
		return
	}

	// Persist the allocated identity before any asynchronous or configuration stage.
	trackedPlan := plan
	trackedPlan.Id = types.StringValue(id)
	trackedPlan.AccessCouldService = fwKnownValue(trackedPlan.AccessCouldService).(types.Bool)
	trackedPlan.Bandwidth = fwKnownValue(trackedPlan.Bandwidth).(types.Int64)
	trackedPlan.BgpAsn = fwKnownValue(trackedPlan.BgpAsn).(types.Int64)
	trackedPlan.BgpCidr = fwKnownValue(trackedPlan.BgpCidr).(types.String)
	trackedPlan.CenId = fwKnownValue(trackedPlan.CenId).(types.String)
	trackedPlan.CenOwnerId = fwKnownValue(trackedPlan.CenOwnerId).(types.String)
	trackedPlan.ConnectionType = fwKnownValue(trackedPlan.ConnectionType).(types.String)
	trackedPlan.CreateTime = fwKnownValue(trackedPlan.CreateTime).(types.String)
	trackedPlan.Description = fwKnownValue(trackedPlan.Description).(types.String)
	trackedPlan.GmtModified = fwKnownValue(trackedPlan.GmtModified).(types.String)
	trackedPlan.Message = fwKnownValue(trackedPlan.Message).(types.String)
	trackedPlan.ResourceGroupId = fwKnownValue(trackedPlan.ResourceGroupId).(types.String)
	trackedPlan.Spec = fwKnownValue(trackedPlan.Spec).(types.String)
	trackedPlan.Status = fwKnownValue(trackedPlan.Status).(types.String)
	trackedPlan.Tags = fwKnownValue(trackedPlan.Tags).(types.Map)
	trackedPlan.VSwitchId = fwKnownValue(trackedPlan.VSwitchId).(types.String)
	trackedPlan.VccName = fwKnownValue(trackedPlan.VccName).(types.String)
	trackedPlan.VpcId = fwKnownValue(trackedPlan.VpcId).(types.String)
	trackedPlan.ZoneId = fwKnownValue(trackedPlan.ZoneId).(types.String)
	resp.Diagnostics.Append(resp.State.Set(ctx, trackedPlan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	object, err := vccWaitForState(r.Client(), id, "Status", []string{"Available"}, nil, operationTimeout)
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Creating %s: waiting for Available", vccTypeName),
			errs.WrapErrorf(err, "Creating %s: waiting for Available (%s)", vccTypeName, id).Error(),
		)
		return
	}
	object, err = vccGet(r.Client(), id)
	if err != nil {
		resp.Diagnostics.AddError("Reading secondary resource attributes", err.Error())
		return
	}
	if object == nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Creating %s: after create", vccTypeName),
			fmt.Sprintf("The resource %s is not visible via GetVcc right after creation", id),
		)
		return
	}
	if !plan.Tags.IsNull() && !plan.Tags.IsUnknown() {
		createdTags := vccState(object, &plan).Tags
		if !plan.Tags.Equal(createdTags) {
			if err := vccSetResourceTags(r.Client(), id, "vcc", plan.Tags, createdTags, false); err != nil {
				resp.Diagnostics.AddError("Setting tags after creation", err.Error())
				return
			}
			object, err = vccGet(r.Client(), id)
			if err != nil || object == nil {
				resp.Diagnostics.AddError("Reading tags after creation", fmt.Sprint(err))
				return
			}
		}
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, vccState(object, &plan))...)
}

func (r *vccResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state vccResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	object, err := vccGet(r.Client(), state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Reading %s: calling GetVcc", vccTypeName),
			errs.WrapErrorf(err, "Reading %s: calling GetVcc", vccTypeName).Error(),
		)
		return
	}
	if object == nil {
		resp.State.RemoveResource(ctx)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, vccState(object, &state))...)
}

func (r *vccResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.State = req.State
	var plan, state vccResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	operationTimeout, timeoutDiagnostics := fwTimeoutDuration(plan.FwTimeouts, "update", 62*time.Minute)
	resp.Diagnostics.Append(timeoutDiagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}

	planValues := map[string]interface{}{"RegionId": r.Client().RegionId}
	planValues["ResourceType"] = "Vcc"
	if !plan.AccessCouldService.IsNull() && !plan.AccessCouldService.IsUnknown() {
		planValues["AccessCouldService"] = fwValueToInterface(plan.AccessCouldService, map[string]fwWireField{}, false)
	}
	if !plan.Bandwidth.IsNull() && !plan.Bandwidth.IsUnknown() {
		planValues["Bandwidth"] = fwValueToInterface(plan.Bandwidth, map[string]fwWireField{}, false)
	}
	if !plan.BgpAsn.IsNull() && !plan.BgpAsn.IsUnknown() {
		planValues["BgpAsn"] = fwValueToInterface(plan.BgpAsn, map[string]fwWireField{}, false)
	}
	if !plan.BgpCidr.IsNull() && !plan.BgpCidr.IsUnknown() {
		planValues["BgpCidr"] = fwValueToInterface(plan.BgpCidr, map[string]fwWireField{}, false)
	}
	if !plan.CenId.IsNull() && !plan.CenId.IsUnknown() {
		planValues["CenId"] = fwValueToInterface(plan.CenId, map[string]fwWireField{}, false)
	}
	if !plan.CenOwnerId.IsNull() && !plan.CenOwnerId.IsUnknown() {
		planValues["CenOwnerId"] = fwValueToInterface(plan.CenOwnerId, map[string]fwWireField{}, false)
	}
	if !plan.ConnectionType.IsNull() && !plan.ConnectionType.IsUnknown() {
		planValues["ConnectionType"] = fwValueToInterface(plan.ConnectionType, map[string]fwWireField{}, false)
	}
	if !plan.CreateTime.IsNull() && !plan.CreateTime.IsUnknown() {
		planValues["CreateTime"] = fwValueToInterface(plan.CreateTime, map[string]fwWireField{}, false)
	}
	if !plan.Description.IsNull() && !plan.Description.IsUnknown() {
		planValues["Description"] = fwValueToInterface(plan.Description, map[string]fwWireField{}, false)
	}
	if !plan.GmtModified.IsNull() && !plan.GmtModified.IsUnknown() {
		planValues["GmtModified"] = fwValueToInterface(plan.GmtModified, map[string]fwWireField{}, false)
	}
	if !plan.Message.IsNull() && !plan.Message.IsUnknown() {
		planValues["Message"] = fwValueToInterface(plan.Message, map[string]fwWireField{}, false)
	}
	if !plan.ResourceGroupId.IsNull() && !plan.ResourceGroupId.IsUnknown() {
		planValues["ResourceGroupId"] = fwValueToInterface(plan.ResourceGroupId, map[string]fwWireField{}, false)
	}
	if !plan.Spec.IsNull() && !plan.Spec.IsUnknown() {
		planValues["Spec"] = fwValueToInterface(plan.Spec, map[string]fwWireField{}, false)
	}
	if !plan.Status.IsNull() && !plan.Status.IsUnknown() {
		planValues["Status"] = fwValueToInterface(plan.Status, map[string]fwWireField{}, false)
	}
	if !plan.Tags.IsNull() && !plan.Tags.IsUnknown() {
		planValues["Tags"] = fwTagsToAttributes(plan.Tags)
	}
	if !plan.VSwitchId.IsNull() && !plan.VSwitchId.IsUnknown() {
		planValues["VSwitchId"] = fwValueToInterface(plan.VSwitchId, map[string]fwWireField{}, false)
	}
	if !plan.VccName.IsNull() && !plan.VccName.IsUnknown() {
		planValues["VccName"] = fwValueToInterface(plan.VccName, map[string]fwWireField{}, false)
	}
	if !plan.VpcId.IsNull() && !plan.VpcId.IsUnknown() {
		planValues["VpcId"] = fwValueToInterface(plan.VpcId, map[string]fwWireField{}, false)
	}
	if !plan.ZoneId.IsNull() && !plan.ZoneId.IsUnknown() {
		planValues["ZoneId"] = fwValueToInterface(plan.ZoneId, map[string]fwWireField{}, false)
	}
	stateValues := map[string]interface{}{"RegionId": r.Client().RegionId}
	stateValues["ResourceType"] = "Vcc"
	if !state.AccessCouldService.IsNull() && !state.AccessCouldService.IsUnknown() {
		stateValues["AccessCouldService"] = fwValueToInterface(state.AccessCouldService, map[string]fwWireField{}, false)
	}
	if !state.Bandwidth.IsNull() && !state.Bandwidth.IsUnknown() {
		stateValues["Bandwidth"] = fwValueToInterface(state.Bandwidth, map[string]fwWireField{}, false)
	}
	if !state.BgpAsn.IsNull() && !state.BgpAsn.IsUnknown() {
		stateValues["BgpAsn"] = fwValueToInterface(state.BgpAsn, map[string]fwWireField{}, false)
	}
	if !state.BgpCidr.IsNull() && !state.BgpCidr.IsUnknown() {
		stateValues["BgpCidr"] = fwValueToInterface(state.BgpCidr, map[string]fwWireField{}, false)
	}
	if !state.CenId.IsNull() && !state.CenId.IsUnknown() {
		stateValues["CenId"] = fwValueToInterface(state.CenId, map[string]fwWireField{}, false)
	}
	if !state.CenOwnerId.IsNull() && !state.CenOwnerId.IsUnknown() {
		stateValues["CenOwnerId"] = fwValueToInterface(state.CenOwnerId, map[string]fwWireField{}, false)
	}
	if !state.ConnectionType.IsNull() && !state.ConnectionType.IsUnknown() {
		stateValues["ConnectionType"] = fwValueToInterface(state.ConnectionType, map[string]fwWireField{}, false)
	}
	if !state.CreateTime.IsNull() && !state.CreateTime.IsUnknown() {
		stateValues["CreateTime"] = fwValueToInterface(state.CreateTime, map[string]fwWireField{}, false)
	}
	if !state.Description.IsNull() && !state.Description.IsUnknown() {
		stateValues["Description"] = fwValueToInterface(state.Description, map[string]fwWireField{}, false)
	}
	if !state.GmtModified.IsNull() && !state.GmtModified.IsUnknown() {
		stateValues["GmtModified"] = fwValueToInterface(state.GmtModified, map[string]fwWireField{}, false)
	}
	if !state.Message.IsNull() && !state.Message.IsUnknown() {
		stateValues["Message"] = fwValueToInterface(state.Message, map[string]fwWireField{}, false)
	}
	if !state.ResourceGroupId.IsNull() && !state.ResourceGroupId.IsUnknown() {
		stateValues["ResourceGroupId"] = fwValueToInterface(state.ResourceGroupId, map[string]fwWireField{}, false)
	}
	if !state.Spec.IsNull() && !state.Spec.IsUnknown() {
		stateValues["Spec"] = fwValueToInterface(state.Spec, map[string]fwWireField{}, false)
	}
	if !state.Status.IsNull() && !state.Status.IsUnknown() {
		stateValues["Status"] = fwValueToInterface(state.Status, map[string]fwWireField{}, false)
	}
	if !state.Tags.IsNull() && !state.Tags.IsUnknown() {
		stateValues["Tags"] = fwTagsToAttributes(state.Tags)
	}
	if !state.VSwitchId.IsNull() && !state.VSwitchId.IsUnknown() {
		stateValues["VSwitchId"] = fwValueToInterface(state.VSwitchId, map[string]fwWireField{}, false)
	}
	if !state.VccName.IsNull() && !state.VccName.IsUnknown() {
		stateValues["VccName"] = fwValueToInterface(state.VccName, map[string]fwWireField{}, false)
	}
	if !state.VpcId.IsNull() && !state.VpcId.IsUnknown() {
		stateValues["VpcId"] = fwValueToInterface(state.VpcId, map[string]fwWireField{}, false)
	}
	if !state.ZoneId.IsNull() && !state.ZoneId.IsUnknown() {
		stateValues["ZoneId"] = fwValueToInterface(state.ZoneId, map[string]fwWireField{}, false)
	}
	updated0 := false
	if fwPlannedChanged(plan.VccName, state.VccName) {
		updated0 = true
	}
	if fwPlannedChanged(plan.Bandwidth, state.Bandwidth) {
		updated0 = true
	}
	if updated0 {
		operationValues0 := planValues
		request0 := map[string]interface{}{
			"VccId":       state.Id.ValueString(),
			"ClientToken": buildClientToken("UpdateVcc"),
		}
		request0["RegionId"] = r.Client().RegionId
		if err := fwApplyMappings(request0, operationValues0, fwChangedUpdateMappings([]fwMapping{
			{Source: "$.VccName", Target: "VccName", Encoding: "", Constant: "", Repeat: false, Values: map[string]string{}},
			{Source: "$.Bandwidth", Target: "Bandwidth", Encoding: "", Constant: "", Repeat: false, Values: map[string]string{}},
			{Source: "$.RegionId", Target: "RegionId", Encoding: "", Constant: "", Repeat: false, Values: map[string]string{}},
			{Source: "$.VccId", Target: "VccId", Encoding: "", Constant: "", Repeat: false, Values: map[string]string{}},
		}, operationValues0, stateValues, []string{"$.VccId"}), false); err != nil {
			resp.Diagnostics.AddError("Building UpdateVcc request", err.Error())
			return
		}
		if _, err := retryRpcPost(r.Client(), "eflo", "2022-05-30", "UpdateVcc", request0, operationTimeout); err != nil {
			resp.Diagnostics.AddError("Updating alicloud_eflo_vcc: UpdateVcc", err.Error())
			return
		}
		if _, err := vccWaitForState(r.Client(), state.Id.ValueString(), "Status", []string{"Available"}, []string{"Update Failed"}, operationTimeout); err != nil {
			resp.Diagnostics.AddError("Waiting after UpdateVcc", err.Error())
			return
		}
	}
	updated1 := false
	if fwPlannedChanged(plan.ResourceGroupId, state.ResourceGroupId) {
		updated1 = true
	}
	if updated1 {
		operationValues1 := planValues
		request1 := map[string]interface{}{
			"ResourceId":  state.Id.ValueString(),
			"ClientToken": buildClientToken("ChangeResourceGroup"),
		}
		request1["RegionId"] = r.Client().RegionId
		if err := fwApplyMappings(request1, operationValues1, fwChangedUpdateMappings([]fwMapping{
			{Source: "$.ResourceGroupId", Target: "ResourceGroupId", Encoding: "", Constant: "", Repeat: false, Values: map[string]string{}},
			{Source: "$.RegionId", Target: "ResourceRegionId", Encoding: "", Constant: "", Repeat: false, Values: map[string]string{}},
			{Source: "$.VccId", Target: "ResourceId", Encoding: "", Constant: "", Repeat: false, Values: map[string]string{}},
			{Source: "$.ResourceType", Target: "ResourceType", Encoding: "", Constant: "\"Vcc\"", Repeat: false, Values: map[string]string{}},
		}, operationValues1, stateValues, []string{"$.VccId", "$.ResourceGroupId", "$.RegionId", "$.ResourceType"}), false); err != nil {
			resp.Diagnostics.AddError("Building ChangeResourceGroup request", err.Error())
			return
		}
		if _, err := retryRpcPost(r.Client(), "eflo-controller", "2022-12-15", "ChangeResourceGroup", request1, operationTimeout); err != nil {
			resp.Diagnostics.AddError("Updating alicloud_eflo_vcc: ChangeResourceGroup", err.Error())
			return
		}
		waitValues1_0 := fwWalkMappedPath(operationValues1, fwMappingPath("$.ResourceGroupId"), nil)
		if len(waitValues1_0) == 0 || waitValues1_0[0].Value == nil {
			resp.Diagnostics.AddError("Waiting after ChangeResourceGroup", "Desired async target $.ResourceGroupId is missing or unknown")
			return
		}
		if _, err := vccWaitForState(r.Client(), state.Id.ValueString(), "ResourceGroupId", []string{fmt.Sprint(waitValues1_0[0].Value)}, nil, operationTimeout); err != nil {
			resp.Diagnostics.AddError("Waiting after ChangeResourceGroup", err.Error())
			return
		}
	}
	updated2 := false
	if fwPlannedChanged(plan.CenId, state.CenId) {
		updated2 = true
	}
	if fwPlannedChanged(plan.VpcId, state.VpcId) {
		updated2 = true
	}
	if fwPlannedChanged(plan.VSwitchId, state.VSwitchId) {
		updated2 = true
	}
	if fwPlannedChanged(plan.ConnectionType, state.ConnectionType) {
		updated2 = true
	}
	if updated2 {
		operationValues2 := planValues
		request2 := map[string]interface{}{
			"VccId":       state.Id.ValueString(),
			"ClientToken": buildClientToken("SwitchVccConnection"),
		}
		request2["RegionId"] = r.Client().RegionId
		if err := fwApplyMappings(request2, operationValues2, fwChangedUpdateMappings([]fwMapping{
			{Source: "$.CenId", Target: "CenId", Encoding: "", Constant: "", Repeat: false, Values: map[string]string{}},
			{Source: "$.VpcId", Target: "VpcId", Encoding: "", Constant: "", Repeat: false, Values: map[string]string{}},
			{Source: "$.VSwitchId", Target: "VSwitchId", Encoding: "", Constant: "", Repeat: false, Values: map[string]string{}},
			{Source: "$.RegionId", Target: "RegionId", Encoding: "", Constant: "", Repeat: false, Values: map[string]string{}},
			{Source: "$.VccId", Target: "VccId", Encoding: "", Constant: "", Repeat: false, Values: map[string]string{}},
			{Source: "$.ConnectionType", Target: "ConnectionType", Encoding: "", Constant: "", Repeat: false, Values: map[string]string{}},
		}, operationValues2, stateValues, []string{"$.VccId", "$.RegionId"}), false); err != nil {
			resp.Diagnostics.AddError("Building SwitchVccConnection request", err.Error())
			return
		}
		if _, err := retryRpcPost(r.Client(), "eflo", "2022-05-30", "SwitchVccConnection", request2, operationTimeout); err != nil {
			resp.Diagnostics.AddError("Updating alicloud_eflo_vcc: SwitchVccConnection", err.Error())
			return
		}
		if _, err := vccWaitForState(r.Client(), state.Id.ValueString(), "Status", []string{"Available"}, []string{"Switch Failed"}, operationTimeout); err != nil {
			resp.Diagnostics.AddError("Waiting after SwitchVccConnection", err.Error())
			return
		}
	}
	// RetryVcc is marked as a separate association resource in metadata.
	// Shared input fields do not make it part of this resource's update contract.
	// RenewVcc is marked as a separate association resource in metadata.
	// Shared input fields do not make it part of this resource's update contract.
	if !plan.Tags.IsUnknown() && !plan.Tags.Equal(state.Tags) {
		var dryRun bool
		if err := vccSetResourceTags(r.Client(), state.Id.ValueString(), "vcc", plan.Tags, state.Tags, dryRun); err != nil {
			resp.Diagnostics.AddError(
				fmt.Sprintf("Updating %s: setting tags", vccTypeName),
				errs.WrapErrorf(err, "Updating %s: setting tags", vccTypeName).Error(),
			)
			return
		}
	}
	object, err := vccGet(r.Client(), state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Reading %s: calling GetVcc", vccTypeName),
			errs.WrapErrorf(err, "Reading %s: calling GetVcc", vccTypeName).Error(),
		)
		return
	}
	if object == nil {
		resp.State.RemoveResource(ctx)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, vccState(object, &plan))...)
}

func (r *vccResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state vccResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	operationTimeout, timeoutDiagnostics := fwTimeoutDuration(state.FwTimeouts, "delete", 62*time.Minute)
	resp.Diagnostics.Append(timeoutDiagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}

	deleteValues := map[string]interface{}{"RegionId": r.Client().RegionId}
	deleteValues["ResourceType"] = "Vcc"
	if !state.AccessCouldService.IsNull() && !state.AccessCouldService.IsUnknown() {
		deleteValues["AccessCouldService"] = fwValueToInterface(state.AccessCouldService, map[string]fwWireField{}, false)
	}
	if !state.Bandwidth.IsNull() && !state.Bandwidth.IsUnknown() {
		deleteValues["Bandwidth"] = fwValueToInterface(state.Bandwidth, map[string]fwWireField{}, false)
	}
	if !state.BgpAsn.IsNull() && !state.BgpAsn.IsUnknown() {
		deleteValues["BgpAsn"] = fwValueToInterface(state.BgpAsn, map[string]fwWireField{}, false)
	}
	if !state.BgpCidr.IsNull() && !state.BgpCidr.IsUnknown() {
		deleteValues["BgpCidr"] = fwValueToInterface(state.BgpCidr, map[string]fwWireField{}, false)
	}
	if !state.CenId.IsNull() && !state.CenId.IsUnknown() {
		deleteValues["CenId"] = fwValueToInterface(state.CenId, map[string]fwWireField{}, false)
	}
	if !state.CenOwnerId.IsNull() && !state.CenOwnerId.IsUnknown() {
		deleteValues["CenOwnerId"] = fwValueToInterface(state.CenOwnerId, map[string]fwWireField{}, false)
	}
	if !state.ConnectionType.IsNull() && !state.ConnectionType.IsUnknown() {
		deleteValues["ConnectionType"] = fwValueToInterface(state.ConnectionType, map[string]fwWireField{}, false)
	}
	if !state.CreateTime.IsNull() && !state.CreateTime.IsUnknown() {
		deleteValues["CreateTime"] = fwValueToInterface(state.CreateTime, map[string]fwWireField{}, false)
	}
	if !state.Description.IsNull() && !state.Description.IsUnknown() {
		deleteValues["Description"] = fwValueToInterface(state.Description, map[string]fwWireField{}, false)
	}
	if !state.GmtModified.IsNull() && !state.GmtModified.IsUnknown() {
		deleteValues["GmtModified"] = fwValueToInterface(state.GmtModified, map[string]fwWireField{}, false)
	}
	if !state.Message.IsNull() && !state.Message.IsUnknown() {
		deleteValues["Message"] = fwValueToInterface(state.Message, map[string]fwWireField{}, false)
	}
	if !state.ResourceGroupId.IsNull() && !state.ResourceGroupId.IsUnknown() {
		deleteValues["ResourceGroupId"] = fwValueToInterface(state.ResourceGroupId, map[string]fwWireField{}, false)
	}
	if !state.Spec.IsNull() && !state.Spec.IsUnknown() {
		deleteValues["Spec"] = fwValueToInterface(state.Spec, map[string]fwWireField{}, false)
	}
	if !state.Status.IsNull() && !state.Status.IsUnknown() {
		deleteValues["Status"] = fwValueToInterface(state.Status, map[string]fwWireField{}, false)
	}
	if !state.Tags.IsNull() && !state.Tags.IsUnknown() {
		deleteValues["Tags"] = fwTagsToAttributes(state.Tags)
	}
	if !state.VSwitchId.IsNull() && !state.VSwitchId.IsUnknown() {
		deleteValues["VSwitchId"] = fwValueToInterface(state.VSwitchId, map[string]fwWireField{}, false)
	}
	if !state.VccName.IsNull() && !state.VccName.IsUnknown() {
		deleteValues["VccName"] = fwValueToInterface(state.VccName, map[string]fwWireField{}, false)
	}
	if !state.VpcId.IsNull() && !state.VpcId.IsUnknown() {
		deleteValues["VpcId"] = fwValueToInterface(state.VpcId, map[string]fwWireField{}, false)
	}
	if !state.ZoneId.IsNull() && !state.ZoneId.IsUnknown() {
		deleteValues["ZoneId"] = fwValueToInterface(state.ZoneId, map[string]fwWireField{}, false)
	}
	deleteValues["VccId"] = state.Id.ValueString()
	needsRefresh := false
	if needsRefresh {
		object, err := vccGet(r.Client(), state.Id.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Reading deletion conditions", err.Error())
			return
		}
		if object == nil {
			return
		}
		for key, value := range object {
			if value != nil {
				deleteValues[key] = value
			}
		}
	}
	selected := false
	deleteEligible0, deleteConditionErr0 := fwDeleteEligible(deleteValues, []fwDeleteCondition{})
	if deleteConditionErr0 != nil {
		resp.Diagnostics.AddError("Selecting deletion operation", deleteConditionErr0.Error())
		return
	}
	selected = selected || deleteEligible0
	if !selected {
		resp.Diagnostics.AddError("Selecting deletion operation", "No deletion operation matches the resource state; the resource remains tracked")
		return
	}
	if deleteEligible0 {
		request0 := map[string]interface{}{"ClientToken": buildClientToken("DeleteVcc")}
		request0["RegionId"] = r.Client().RegionId
		if err := fwApplyMappings(request0, deleteValues, []fwMapping{
			{Source: "$.RegionId", Target: "RegionId", Encoding: "", Constant: "", Repeat: false, Values: map[string]string{}},
			{Source: "$.VccId", Target: "VccId", Encoding: "", Constant: "", Repeat: false, Values: map[string]string{}},
		}, false); err != nil {
			resp.Diagnostics.AddError("Building DeleteVcc request", err.Error())
			return
		}
		if _, err := retryRpcPost(r.Client(), "eflo", "2022-05-30", "DeleteVcc", request0, operationTimeout); err != nil && !isExpectedErrors(err, []string{"1011"}) {
			resp.Diagnostics.AddError("Deleting alicloud_eflo_vcc: DeleteVcc", err.Error())
			return
		}
		if _, err := vccWaitForState(r.Client(), state.Id.ValueString(), "Status", []string{}, nil, operationTimeout); err != nil {
			resp.Diagnostics.AddError("Waiting after DeleteVcc", err.Error())
			return
		}
	}
	// gap: v1 Delete 用 BuildStateConf(target=[]) 轮询 StateRefreshFunc 直到 Get 返回 nil（gone），
	// 解决 gateway-route-table 异步删除后 DeleteVpc 竞态 DependencyViolation.GatewayRouteTable。
	// v2 WaitForState：空 deleteTargetStatus（gone-wait）发 Target=[]string{}（空 slice, len 0），
	// 镜像 v1 BuildStateConf(target=[])。hashicorp retry 当 Refresh 返回 (nil,"",nil)（gone）且
	// Target 空 → success（done）；若误用 Target=[""]（len 1, 非空），gone 时 helper 报
	// NotFoundError("couldn't find resource (N retries)") 而非 success——此前 route_table destroy
	// 即栽此。非空 deleteTargetStatus 则发 Target=["<status>"] 等待该终态。Get 在 not-found 返回
	// (nil,nil)（非 error），故 wait 不挂；同步删除资源首检即 gone（nil,"",nil→Target 空→done），无 hang。
	if _, err := vccWaitForState(r.Client(), state.Id.ValueString(), "Status", []string{}, nil, operationTimeout); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Deleting %s: waiting for gone", vccTypeName),
			errs.WrapErrorf(err, "Deleting %s: waiting for gone (%s)", vccTypeName, state.Id.ValueString()).Error(),
		)
		return
	}
}

func (r *vccResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// vccState maps the Get response object onto the resource model. Fields present in the
// response are read via their responseKey (e.g. status -> ListenerStatus); request-only fields
// (responseKey empty, e.g. dry_run) are preserved from the prior plan/state instead.
func vccState(object map[string]interface{}, prior *vccResourceModel) vccResourceModel {
	// Get has already merged and normalized every response mapping.
	str := func(key string) types.String {
		cur := interface{}(object)
		for _, seg := range strings.Split(key, ".") {
			m, ok := cur.(map[string]interface{})
			if !ok {
				return types.StringNull()
			}
			cur, ok = m[seg]
			if !ok {
				return types.StringNull()
			}
		}
		// gap#7c: present-but-empty strings must return ValueString(""), not Null.
		// v1 d.Set(field, "") treats empty as set (not absence); TestCheckResourceAttrSet
		// expects a non-null value for API-returned "" (e.g. DhcpOptionsSetId when VPC has
		// no DHCP set). The earlier && v != "" collapsed present-empty to Null, failing Set
		// checks. Absent fields still return Null via the !ok branches above.
		if v, ok := cur.(string); ok {
			return types.StringValue(v)
		}
		return types.StringNull()
	}
	intVal := func(key string) types.Int64 {
		cur := interface{}(object)
		for _, seg := range strings.Split(key, ".") {
			m, ok := cur.(map[string]interface{})
			if !ok {
				return types.Int64Null()
			}
			cur, ok = m[seg]
			if !ok {
				return types.Int64Null()
			}
		}
		switch v := cur.(type) {
		case float64:
			return types.Int64Value(int64(v))
		case int:
			return types.Int64Value(int64(v))
		case int64:
			return types.Int64Value(v)
		case json.Number:
			if n, err := v.Int64(); err == nil {
				return types.Int64Value(n)
			}
		case string:
			if n, err := strconv.ParseInt(v, 10, 64); err == nil {
				return types.Int64Value(n)
			}
		}
		return types.Int64Null()
	}
	// lookup traverses a dotted responseKey (e.g. "RegionStatuses.RegionStatus") segment by segment
	// from the Get response object, returning the raw value (array/map/scalar) or nil. gap#5: array
	// fields' responseKey is the array's true path from the read-tree [*] child, which may be a
	// two-level unwrap (object["RegionStatuses"]["RegionStatus"]) that a literal object[key] lookup
	// cannot reach. Single-segment keys (e.g. "IpWhiteList") traverse one step == object[key].
	lookup := func(key string) interface{} {
		cur := interface{}(object)
		for _, seg := range strings.Split(key, ".") {
			m, ok := cur.(map[string]interface{})
			if !ok {
				return nil
			}
			cur, ok = m[seg]
			if !ok {
				return nil
			}
		}
		return cur
	}
	m := vccResourceModel{
		FwTimeouts:      types.ObjectNull(fwTimeoutTypes()),
		Id:              str("VccId"),
		Bandwidth:       intVal("Bandwidth"),
		BgpCidr:         str("BgpCidr"),
		CenId:           str("CenId"),
		CenOwnerId:      str("CenOwnerId"),
		ConnectionType:  str("ConnectionType"),
		CreateTime:      str("CreateTime"),
		GmtModified:     str("GmtModified"),
		Message:         str("Message"),
		ResourceGroupId: str("ResourceGroupId"),
		Spec:            str("Spec"),
		Status:          str("Status"),
		Tags:            tagsToTypeMap(lookup("Tags")),
		VSwitchId:       str("VSwitchId"),
		VccName:         str("VccName"),
		VpcId:           str("VpcId"),
	}
	if prior != nil {
		m.FwTimeouts = prior.FwTimeouts
		m.AccessCouldService = prior.AccessCouldService
		m.BgpAsn = prior.BgpAsn
		m.Description = prior.Description
		m.ZoneId = prior.ZoneId
	}
	return m
}
