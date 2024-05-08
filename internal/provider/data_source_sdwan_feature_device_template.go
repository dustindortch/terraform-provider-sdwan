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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &FeatureDeviceTemplateDataSource{}
	_ datasource.DataSourceWithConfigure = &FeatureDeviceTemplateDataSource{}
)

func NewFeatureDeviceTemplateDataSource() datasource.DataSource {
	return &FeatureDeviceTemplateDataSource{}
}

type FeatureDeviceTemplateDataSource struct {
	client *sdwan.Client
}

func (d *FeatureDeviceTemplateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_feature_device_template"
}

func (d *FeatureDeviceTemplateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Feature Device Template .",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the object",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the device template",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the device template",
				Computed:            true,
			},
			"device_type": schema.StringAttribute{
				MarkdownDescription: "The device type (e.g., `vedge-ISR-4331`)",
				Computed:            true,
			},
			"device_role": schema.StringAttribute{
				MarkdownDescription: "The device role",
				Computed:            true,
			},
			"policy_id": schema.StringAttribute{
				MarkdownDescription: "The policy ID",
				Computed:            true,
			},
			"policy_version": schema.Int64Attribute{
				MarkdownDescription: "The policy version",
				Computed:            true,
			},
			"security_policy_id": schema.StringAttribute{
				MarkdownDescription: "The security policy ID",
				Computed:            true,
			},
			"security_policy_version": schema.Int64Attribute{
				MarkdownDescription: "The security policy version",
				Computed:            true,
			},
			"general_templates": schema.SetNestedAttribute{
				MarkdownDescription: "List of general templates",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Feature template ID",
							Computed:            true,
						},
						"version": schema.Int64Attribute{
							MarkdownDescription: "Feature template version",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "Feature template type",
							Computed:            true,
						},
						"sub_templates": schema.SetNestedAttribute{
							MarkdownDescription: "List of sub templates",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Feature template ID",
										Computed:            true,
									},
									"version": schema.Int64Attribute{
										MarkdownDescription: "Feature template version",
										Computed:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: "Feature template type",
										Computed:            true,
									},
									"sub_templates": schema.SetNestedAttribute{
										MarkdownDescription: "List of sub templates",
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													MarkdownDescription: "Feature template ID",
													Computed:            true,
												},
												"version": schema.Int64Attribute{
													MarkdownDescription: "Feature template version",
													Computed:            true,
												},
												"type": schema.StringAttribute{
													MarkdownDescription: "Feature template type",
													Computed:            true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *FeatureDeviceTemplateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *FeatureDeviceTemplateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config FeatureDeviceTemplate

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	res, err := d.client.Get("/template/device/object/" + url.QueryEscape(config.Id.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
