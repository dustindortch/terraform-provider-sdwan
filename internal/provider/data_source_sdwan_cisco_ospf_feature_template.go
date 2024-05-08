// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"fmt"
	"net/url"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &CiscoOSPFFeatureTemplateDataSource{}
	_ datasource.DataSourceWithConfigure = &CiscoOSPFFeatureTemplateDataSource{}
)

func NewCiscoOSPFFeatureTemplateDataSource() datasource.DataSource {
	return &CiscoOSPFFeatureTemplateDataSource{}
}

type CiscoOSPFFeatureTemplateDataSource struct {
	client *sdwan.Client
}

func (d *CiscoOSPFFeatureTemplateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cisco_ospf_feature_template"
}

func (d *CiscoOSPFFeatureTemplateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Cisco OSPF feature template.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the feature template",
				Optional:            true,
				Computed:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the feature template",
				Computed:            true,
			},
			"template_type": schema.StringAttribute{
				MarkdownDescription: "The template type",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the feature template",
				Optional:            true,
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the feature template",
				Computed:            true,
			},
			"device_types": schema.SetAttribute{
				MarkdownDescription: "List of supported device types",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"router_id": schema.StringAttribute{
				MarkdownDescription: "Set OSPF router ID to override system IP address",
				Computed:            true,
			},
			"router_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"auto_cost_reference_bandwidth": schema.Int64Attribute{
				MarkdownDescription: "Set reference bandwidth method to assign OSPF cost",
				Computed:            true,
			},
			"auto_cost_reference_bandwidth_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"compatible_rfc1583": schema.BoolAttribute{
				MarkdownDescription: "Calculate summary route cost based on RFC 1583",
				Computed:            true,
			},
			"compatible_rfc1583_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"default_information_originate": schema.BoolAttribute{
				MarkdownDescription: "Distribute default external route into OSPF",
				Computed:            true,
			},
			"default_information_originate_always": schema.BoolAttribute{
				MarkdownDescription: "Always advertise default route",
				Computed:            true,
			},
			"default_information_originate_always_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"default_information_originate_metric": schema.Int64Attribute{
				MarkdownDescription: "Set metric used to generate default route <0..16777214>",
				Computed:            true,
			},
			"default_information_originate_metric_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"default_information_originate_metric_type": schema.StringAttribute{
				MarkdownDescription: "Set default route type",
				Computed:            true,
			},
			"default_information_originate_metric_type_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"distance_external": schema.Int64Attribute{
				MarkdownDescription: "Set distance for external routes",
				Computed:            true,
			},
			"distance_external_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"distance_inter_area": schema.Int64Attribute{
				MarkdownDescription: "Set distance for inter-area routes",
				Computed:            true,
			},
			"distance_inter_area_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"distance_intra_area": schema.Int64Attribute{
				MarkdownDescription: "Set distance for intra-area routes",
				Computed:            true,
			},
			"distance_intra_area_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"timers_spf_delay": schema.Int64Attribute{
				MarkdownDescription: "Set delay from first change received until performing SPF calculation",
				Computed:            true,
			},
			"timers_spf_delay_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"timers_spf_initial_hold": schema.Int64Attribute{
				MarkdownDescription: "Set initial hold time between consecutive SPF calculations",
				Computed:            true,
			},
			"timers_spf_initial_hold_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"timers_spf_max_hold": schema.Int64Attribute{
				MarkdownDescription: "Set maximum hold time between consecutive SPF calculations",
				Computed:            true,
			},
			"timers_spf_max_hold_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"redistribute": schema.ListNestedAttribute{
				MarkdownDescription: "Redistribute routes",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"protocol": schema.StringAttribute{
							MarkdownDescription: "Set the protocol",
							Computed:            true,
						},
						"protocol_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"route_policy": schema.StringAttribute{
							MarkdownDescription: "Set route policy to apply to redistributed routes",
							Computed:            true,
						},
						"route_policy_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"nat_dia": schema.BoolAttribute{
							MarkdownDescription: "Enable NAT DIA for redistributed routes",
							Computed:            true,
						},
						"nat_dia_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
			"max_metric_router_lsa": schema.ListNestedAttribute{
				MarkdownDescription: "Advertise own router LSA with infinite distance",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ad_type": schema.StringAttribute{
							MarkdownDescription: "Set the router LSA advertisement type",
							Computed:            true,
						},
						"time": schema.Int64Attribute{
							MarkdownDescription: "Set how long to advertise maximum metric after router starts up",
							Computed:            true,
						},
						"time_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
			"route_policies": schema.ListNestedAttribute{
				MarkdownDescription: "Set route policy to apply",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"direction": schema.StringAttribute{
							MarkdownDescription: "Set direction to apply policy",
							Computed:            true,
						},
						"direction_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"policy_name": schema.StringAttribute{
							MarkdownDescription: "Name of route policy",
							Computed:            true,
						},
						"policy_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
			"areas": schema.ListNestedAttribute{
				MarkdownDescription: "Configure OSPF area",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"area_number": schema.Int64Attribute{
							MarkdownDescription: "Set OSPF area number",
							Computed:            true,
						},
						"area_number_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"stub": schema.BoolAttribute{
							MarkdownDescription: "Stub area",
							Computed:            true,
						},
						"stub_no_summary": schema.BoolAttribute{
							MarkdownDescription: "Do not inject interarea routes into stub",
							Computed:            true,
						},
						"stub_no_summary_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"nssa": schema.BoolAttribute{
							MarkdownDescription: "NSSA area",
							Computed:            true,
						},
						"nssa_no_summary": schema.BoolAttribute{
							MarkdownDescription: "Do not inject interarea routes into NSSA",
							Computed:            true,
						},
						"nssa_no_summary_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"interfaces": schema.ListNestedAttribute{
							MarkdownDescription: "Set OSPF interface parameters",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Set interface name",
										Computed:            true,
									},
									"name_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"hello_interval": schema.Int64Attribute{
										MarkdownDescription: "Set interval between OSPF hello packets",
										Computed:            true,
									},
									"hello_interval_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"dead_interval": schema.Int64Attribute{
										MarkdownDescription: "Set interval after which neighbor is declared to be down",
										Computed:            true,
									},
									"dead_interval_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"retransmit_interval": schema.Int64Attribute{
										MarkdownDescription: "Set time between retransmitting LSAs",
										Computed:            true,
									},
									"retransmit_interval_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"cost": schema.Int64Attribute{
										MarkdownDescription: "Set cost of OSPF interface",
										Computed:            true,
									},
									"cost_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"priority": schema.Int64Attribute{
										MarkdownDescription: "Set router’s priority to be elected as designated router",
										Computed:            true,
									},
									"priority_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"network": schema.StringAttribute{
										MarkdownDescription: "Set the OSPF network type",
										Computed:            true,
									},
									"network_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"passive_interface": schema.BoolAttribute{
										MarkdownDescription: "Set the interface to advertise its address, but not to actively run OSPF",
										Computed:            true,
									},
									"passive_interface_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"authentication_type": schema.StringAttribute{
										MarkdownDescription: "Set OSPF interface authentication type",
										Computed:            true,
									},
									"authentication_type_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"authentication_message_digest_key_id": schema.Int64Attribute{
										MarkdownDescription: "Set MD5 message digest key",
										Computed:            true,
									},
									"authentication_message_digest_key_id_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"authentication_message_digest_key": schema.StringAttribute{
										MarkdownDescription: "Set MD5 authentication key",
										Computed:            true,
									},
									"authentication_message_digest_key_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"optional": schema.BoolAttribute{
										MarkdownDescription: "Indicates if list item is considered optional.",
										Computed:            true,
									},
								},
							},
						},
						"ranges": schema.ListNestedAttribute{
							MarkdownDescription: "Summarize OSPF routes at an area boundary",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										MarkdownDescription: "Set matching prefix",
										Computed:            true,
									},
									"address_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"cost": schema.Int64Attribute{
										MarkdownDescription: "Set cost for this range",
										Computed:            true,
									},
									"cost_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"no_advertise": schema.BoolAttribute{
										MarkdownDescription: "Do not advertise this range",
										Computed:            true,
									},
									"no_advertise_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"optional": schema.BoolAttribute{
										MarkdownDescription: "Indicates if list item is considered optional.",
										Computed:            true,
									},
								},
							},
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *CiscoOSPFFeatureTemplateDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *CiscoOSPFFeatureTemplateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *CiscoOSPFFeatureTemplateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config CiscoOSPF

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	if config.Id.IsNull() && !config.Name.IsNull() {
		res, err := d.client.Get("/template/feature")
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve existing templates, got error: %s", err))
			return
		}
		if value := res.Get("data"); len(value.Array()) > 0 {
			value.ForEach(func(k, v gjson.Result) bool {
				if config.Name.ValueString() == v.Get("templateName").String() {
					config.Id = types.StringValue(v.Get("templateId").String())
					tflog.Debug(ctx, fmt.Sprintf("%s: Found feature template with name '%v', id: %v", config.Id.String(), config.Name.ValueString(), config.Id.String()))
					return false
				}
				return true
			})
		}
		if config.Id.IsNull() {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find feature template with name: %s", config.Name.ValueString()))
			return
		}
	}

	res, err := d.client.Get("/template/feature/object/" + url.QueryEscape(config.Id.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Name.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
