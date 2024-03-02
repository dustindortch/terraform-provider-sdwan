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

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

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

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &CiscoDHCPServerFeatureTemplateDataSource{}
	_ datasource.DataSourceWithConfigure = &CiscoDHCPServerFeatureTemplateDataSource{}
)

func NewCiscoDHCPServerFeatureTemplateDataSource() datasource.DataSource {
	return &CiscoDHCPServerFeatureTemplateDataSource{}
}

type CiscoDHCPServerFeatureTemplateDataSource struct {
	client *sdwan.Client
}

func (d *CiscoDHCPServerFeatureTemplateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cisco_dhcp_server_feature_template"
}

func (d *CiscoDHCPServerFeatureTemplateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Cisco DHCP Server feature template.",

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
			"address_pool": schema.StringAttribute{
				MarkdownDescription: "Configure IPv4 prefix range of the DHCP address pool",
				Computed:            true,
			},
			"address_pool_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"exclude_addresses": schema.SetAttribute{
				MarkdownDescription: "Configure IPv4 address to exclude from DHCP address pool",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"exclude_addresses_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"lease_time": schema.Int64Attribute{
				MarkdownDescription: "Configure how long a DHCP-assigned IP address is valid",
				Computed:            true,
			},
			"lease_time_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"interface_mtu": schema.Int64Attribute{
				MarkdownDescription: "Set MTU on interface to DHCP client",
				Computed:            true,
			},
			"interface_mtu_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"domain_name": schema.StringAttribute{
				MarkdownDescription: "Set domain name client uses to resolve hostnames",
				Computed:            true,
			},
			"domain_name_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"default_gateway": schema.StringAttribute{
				MarkdownDescription: "Set IP address of default gateway",
				Computed:            true,
			},
			"default_gateway_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"dns_servers": schema.SetAttribute{
				MarkdownDescription: "Configure one or more DNS server IP addresses",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"dns_servers_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tftp_servers": schema.SetAttribute{
				MarkdownDescription: "Configure TFTP server IP addresses",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"tftp_servers_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"static_leases": schema.ListNestedAttribute{
				MarkdownDescription: "Configure static IP addresses",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"mac_address": schema.StringAttribute{
							MarkdownDescription: "Set MAC address of client",
							Computed:            true,
						},
						"mac_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ip_address": schema.StringAttribute{
							MarkdownDescription: "Set client’s static IP address",
							Computed:            true,
						},
						"ip_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"hostname": schema.StringAttribute{
							MarkdownDescription: "Set client’s hostname",
							Computed:            true,
						},
						"hostname_variable": schema.StringAttribute{
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
			"options": schema.ListNestedAttribute{
				MarkdownDescription: "Configure Options Code",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"option_code": schema.Int64Attribute{
							MarkdownDescription: "Set Option Code",
							Computed:            true,
						},
						"option_code_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ascii": schema.StringAttribute{
							MarkdownDescription: "Set ASCII value",
							Computed:            true,
						},
						"ascii_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"hex": schema.StringAttribute{
							MarkdownDescription: "Set HEX value",
							Computed:            true,
						},
						"hex_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ip_address": schema.SetAttribute{
							MarkdownDescription: "Set ip address",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"ip_address_variable": schema.StringAttribute{
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
		},
	}
}

func (d *CiscoDHCPServerFeatureTemplateDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *CiscoDHCPServerFeatureTemplateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

func (d *CiscoDHCPServerFeatureTemplateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config CiscoDHCPServer

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

	res, err := d.client.Get("/template/feature/object/" + config.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Name.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
