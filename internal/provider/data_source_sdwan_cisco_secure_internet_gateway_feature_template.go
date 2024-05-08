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
	_ datasource.DataSource              = &CiscoSecureInternetGatewayFeatureTemplateDataSource{}
	_ datasource.DataSourceWithConfigure = &CiscoSecureInternetGatewayFeatureTemplateDataSource{}
)

func NewCiscoSecureInternetGatewayFeatureTemplateDataSource() datasource.DataSource {
	return &CiscoSecureInternetGatewayFeatureTemplateDataSource{}
}

type CiscoSecureInternetGatewayFeatureTemplateDataSource struct {
	client *sdwan.Client
}

func (d *CiscoSecureInternetGatewayFeatureTemplateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cisco_secure_internet_gateway_feature_template"
}

func (d *CiscoSecureInternetGatewayFeatureTemplateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Cisco Secure Internet Gateway feature template.",

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
			"vpn_id": schema.Int64Attribute{
				MarkdownDescription: "List of VPN instances",
				Computed:            true,
			},
			"interfaces": schema.ListNestedAttribute{
				MarkdownDescription: "Interface name: IPsec when present",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Interface name: IPsec when present",
							Computed:            true,
						},
						"name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"auto_tunnel_mode": schema.BoolAttribute{
							MarkdownDescription: "Auto Tunnel Mode",
							Computed:            true,
						},
						"shutdown": schema.BoolAttribute{
							MarkdownDescription: "Administrative state",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "Interface description",
							Computed:            true,
						},
						"description_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ip_unnumbered": schema.BoolAttribute{
							MarkdownDescription: "Unnumbered interface",
							Computed:            true,
						},
						"ipv4_address": schema.StringAttribute{
							MarkdownDescription: "Assign IPv4 address",
							Computed:            true,
						},
						"ipv4_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"tunnel_source": schema.StringAttribute{
							MarkdownDescription: "Tunnel source IP Address",
							Computed:            true,
						},
						"tunnel_source_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"tunnel_source_interface": schema.StringAttribute{
							MarkdownDescription: "<1..32 characters> Interface name: ge0/<0-..> or ge0/<0-..>.vlanid",
							Computed:            true,
						},
						"tunnel_source_interface_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"tunnel_route_via": schema.StringAttribute{
							MarkdownDescription: "<1..32 characters> Interface name: ge0/<0-..> or ge0/<0-..>.vlanid",
							Computed:            true,
						},
						"tunnel_route_via_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"tunnel_destination": schema.StringAttribute{
							MarkdownDescription: "Tunnel destination IP address",
							Computed:            true,
						},
						"tunnel_destination_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"application": schema.StringAttribute{
							MarkdownDescription: "Enable Application Tunnel Type",
							Computed:            true,
						},
						"sig_provider": schema.StringAttribute{
							MarkdownDescription: "SIG Tunnel Provider",
							Computed:            true,
						},
						"tunnel_dc_preference": schema.StringAttribute{
							MarkdownDescription: "SIG Tunnel Data Center",
							Computed:            true,
						},
						"tcp_mss": schema.Int64Attribute{
							MarkdownDescription: "TCP MSS on SYN packets, in bytes",
							Computed:            true,
						},
						"tcp_mss_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"mtu": schema.Int64Attribute{
							MarkdownDescription: "Interface MTU <576..2000>, in bytes",
							Computed:            true,
						},
						"mtu_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"dead_peer_detection_interval": schema.Int64Attribute{
							MarkdownDescription: "IKE keepalive interval (seconds)",
							Computed:            true,
						},
						"dead_peer_detection_interval_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"dead_peer_detection_retries": schema.Int64Attribute{
							MarkdownDescription: "IKE keepalive retries",
							Computed:            true,
						},
						"dead_peer_detection_retries_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ike_version": schema.Int64Attribute{
							MarkdownDescription: "IKE Version <1..2>",
							Computed:            true,
						},
						"ike_version_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ike_pre_shared_key": schema.StringAttribute{
							MarkdownDescription: "Use preshared key to authenticate IKE peer",
							Computed:            true,
						},
						"ike_pre_shared_key_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ike_rekey_interval": schema.Int64Attribute{
							MarkdownDescription: "IKE rekey interval <300..1209600> seconds",
							Computed:            true,
						},
						"ike_rekey_interval_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ike_ciphersuite": schema.StringAttribute{
							MarkdownDescription: "IKE identity the IKE preshared secret belongs to",
							Computed:            true,
						},
						"ike_ciphersuite_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ike_group": schema.StringAttribute{
							MarkdownDescription: "IKE Diffie Hellman Groups",
							Computed:            true,
						},
						"ike_group_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ike_pre_shared_key_dynamic": schema.BoolAttribute{
							MarkdownDescription: "Use preshared key to authenticate IKE peer",
							Computed:            true,
						},
						"ike_pre_shared_key_local_id": schema.StringAttribute{
							MarkdownDescription: "IKE ID for the local endpoint. Input IPv4 address, domain name, or email address",
							Computed:            true,
						},
						"ike_pre_shared_key_local_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ike_pre_shared_key_remote_id": schema.StringAttribute{
							MarkdownDescription: "IKE ID for the remote endpoint. Input IPv4 address, domain name, or email address",
							Computed:            true,
						},
						"ike_pre_shared_key_remote_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ipsec_rekey_interval": schema.Int64Attribute{
							MarkdownDescription: "IPsec rekey interval <300..1209600> seconds",
							Computed:            true,
						},
						"ipsec_rekey_interval_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ipsec_replay_window": schema.Int64Attribute{
							MarkdownDescription: "Replay window size 32..8192 (must be a power of 2)",
							Computed:            true,
						},
						"ipsec_replay_window_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ipsec_ciphersuite": schema.StringAttribute{
							MarkdownDescription: "IPsec(ESP) encryption and integrity protocol",
							Computed:            true,
						},
						"ipsec_ciphersuite_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ipsec_perfect_forward_secrecy": schema.StringAttribute{
							MarkdownDescription: "IPsec perfect forward secrecy settings",
							Computed:            true,
						},
						"ipsec_perfect_forward_secrecy_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"tracker": schema.StringAttribute{
							MarkdownDescription: "Enable tracker for this interface",
							Computed:            true,
						},
						"track_enable": schema.BoolAttribute{
							MarkdownDescription: "Enable/disable SIG tracking",
							Computed:            true,
						},
						"tunnel_public_ip": schema.StringAttribute{
							MarkdownDescription: "Public IP required to setup GRE tunnel to Zscaler",
							Computed:            true,
						},
						"tunnel_public_ip_variable": schema.StringAttribute{
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
			"services": schema.ListNestedAttribute{
				MarkdownDescription: "Configure services",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"service_type": schema.StringAttribute{
							MarkdownDescription: "Service Type",
							Computed:            true,
						},
						"interface_pairs": schema.ListNestedAttribute{
							MarkdownDescription: "Interface Pair for active and backup",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"active_interface": schema.StringAttribute{
										MarkdownDescription: "Active Tunnel Interface for SIG",
										Computed:            true,
									},
									"active_interface_weight": schema.Int64Attribute{
										MarkdownDescription: "Active Tunnel Interface Weight",
										Computed:            true,
									},
									"backup_interface": schema.StringAttribute{
										MarkdownDescription: "Backup Tunnel Interface for SIG",
										Computed:            true,
									},
									"backup_interface_weight": schema.Int64Attribute{
										MarkdownDescription: "Backup Tunnel Interface Weight",
										Computed:            true,
									},
									"optional": schema.BoolAttribute{
										MarkdownDescription: "Indicates if list item is considered optional.",
										Computed:            true,
									},
								},
							},
						},
						"zscaler_authentication_required": schema.BoolAttribute{
							MarkdownDescription: "Enforce Authentication",
							Computed:            true,
						},
						"zscaler_xff_forward": schema.BoolAttribute{
							MarkdownDescription: "XFF forwarding enabled",
							Computed:            true,
						},
						"zscaler_firewall_enabled": schema.BoolAttribute{
							MarkdownDescription: "Firewall enabled",
							Computed:            true,
						},
						"zscaler_ips_control_enabled": schema.BoolAttribute{
							MarkdownDescription: "Enable IPS Control",
							Computed:            true,
						},
						"zscaler_caution_enabled": schema.BoolAttribute{
							MarkdownDescription: "Enable Caution",
							Computed:            true,
						},
						"zscaler_primary_data_center": schema.StringAttribute{
							MarkdownDescription: "Custom Primary Datacenter",
							Computed:            true,
						},
						"zscaler_primary_data_center_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"zscaler_secondary_data_center": schema.StringAttribute{
							MarkdownDescription: "Custom Secondary Datacenter",
							Computed:            true,
						},
						"zscaler_secondary_data_center_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"zscaler_surrogate_ip": schema.BoolAttribute{
							MarkdownDescription: "Enable Surrogate IP",
							Computed:            true,
						},
						"zscaler_surrogate_idle_time": schema.Int64Attribute{
							MarkdownDescription: "Idle time to disassociation",
							Computed:            true,
						},
						"zscaler_surrogate_display_time_unit": schema.StringAttribute{
							MarkdownDescription: "Display time unit",
							Computed:            true,
						},
						"zscaler_surrogate_ip_enforce_for_known_browsers": schema.BoolAttribute{
							MarkdownDescription: "Enforce Surrogate IP for known browsers",
							Computed:            true,
						},
						"zscaler_surrogate_refresh_time": schema.Int64Attribute{
							MarkdownDescription: "Refresh time for re-validation of surrogacy in minutes",
							Computed:            true,
						},
						"zscaler_surrogate_refresh_time_unit": schema.StringAttribute{
							MarkdownDescription: "Refresh Time unit",
							Computed:            true,
						},
						"zscaler_aup_enabled": schema.BoolAttribute{
							MarkdownDescription: "Enable Acceptable User Policy",
							Computed:            true,
						},
						"zscaler_aup_block_internet_until_accepted": schema.BoolAttribute{
							MarkdownDescription: "For first-time Acceptable User Policy behavior, block Internet access",
							Computed:            true,
						},
						"zscaler_aup_force_ssl_inspection": schema.BoolAttribute{
							MarkdownDescription: "For first-time Acceptable User Policy behavior, force SSL inspection",
							Computed:            true,
						},
						"zscaler_aup_timeout": schema.Int64Attribute{
							MarkdownDescription: "Custom Acceptable User Policy frequency in days",
							Computed:            true,
						},
						"zscaler_location_name": schema.StringAttribute{
							MarkdownDescription: "Zscaler location name (optional)",
							Computed:            true,
						},
						"zscaler_location_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"umbrella_primary_data_center": schema.StringAttribute{
							MarkdownDescription: "Umbrella Primary Datacenter",
							Computed:            true,
						},
						"umbrella_primary_data_center_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"umbrella_secondary_data_center": schema.StringAttribute{
							MarkdownDescription: "Umbrella Secondary Datacenter",
							Computed:            true,
						},
						"umbrella_secondary_data_center_variable": schema.StringAttribute{
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
			"tracker_source_ip": schema.StringAttribute{
				MarkdownDescription: "Source IP address for Tracker",
				Computed:            true,
			},
			"tracker_source_ip_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"trackers": schema.ListNestedAttribute{
				MarkdownDescription: "Tracker configuration",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Tracker name",
							Computed:            true,
						},
						"name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"endpoint_api_url": schema.StringAttribute{
							MarkdownDescription: "API url of endpoint",
							Computed:            true,
						},
						"endpoint_api_url_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"threshold": schema.Int64Attribute{
							MarkdownDescription: "Probe Timeout threshold <100..1000> milliseconds",
							Computed:            true,
						},
						"threshold_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"interval": schema.Int64Attribute{
							MarkdownDescription: "Probe interval <10..600> seconds",
							Computed:            true,
						},
						"interval_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"multiplier": schema.Int64Attribute{
							MarkdownDescription: "Probe failure multiplier <1..10> failed attempts",
							Computed:            true,
						},
						"multiplier_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"tracker_type": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
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

func (d *CiscoSecureInternetGatewayFeatureTemplateDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *CiscoSecureInternetGatewayFeatureTemplateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *CiscoSecureInternetGatewayFeatureTemplateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config CiscoSecureInternetGateway

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
