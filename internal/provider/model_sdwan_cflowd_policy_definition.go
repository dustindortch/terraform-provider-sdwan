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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type CflowdPolicyDefinition struct {
	Id                  types.String                       `tfsdk:"id"`
	Version             types.Int64                        `tfsdk:"version"`
	Name                types.String                       `tfsdk:"name"`
	Description         types.String                       `tfsdk:"description"`
	ActiveFlowTimeout   types.Int64                        `tfsdk:"active_flow_timeout"`
	InactiveFlowTimeout types.Int64                        `tfsdk:"inactive_flow_timeout"`
	SamplingInterval    types.Int64                        `tfsdk:"sampling_interval"`
	FlowRefresh         types.Int64                        `tfsdk:"flow_refresh"`
	Protocol            types.String                       `tfsdk:"protocol"`
	Tos                 types.Bool                         `tfsdk:"tos"`
	RemarkedDscp        types.Bool                         `tfsdk:"remarked_dscp"`
	Collectors          []CflowdPolicyDefinitionCollectors `tfsdk:"collectors"`
}

type CflowdPolicyDefinitionCollectors struct {
	VpnId           types.Int64  `tfsdk:"vpn_id"`
	IpAddress       types.String `tfsdk:"ip_address"`
	Port            types.Int64  `tfsdk:"port"`
	Transport       types.String `tfsdk:"transport"`
	SourceInterface types.String `tfsdk:"source_interface"`
	ExportSpreading types.String `tfsdk:"export_spreading"`
}

func (data CflowdPolicyDefinition) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "type", "cflowd")
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.ActiveFlowTimeout.IsNull() {
		body, _ = sjson.Set(body, "definition.flowActiveTimeout", data.ActiveFlowTimeout.ValueInt64())
	}
	if !data.InactiveFlowTimeout.IsNull() {
		body, _ = sjson.Set(body, "definition.flowInactiveTimeout", data.InactiveFlowTimeout.ValueInt64())
	}
	if !data.SamplingInterval.IsNull() {
		body, _ = sjson.Set(body, "definition.flowSamplingInterval", data.SamplingInterval.ValueInt64())
	}
	if !data.FlowRefresh.IsNull() {
		body, _ = sjson.Set(body, "definition.templateRefresh", data.FlowRefresh.ValueInt64())
	}
	if !data.Protocol.IsNull() {
		body, _ = sjson.Set(body, "definition.protocol", data.Protocol.ValueString())
	}
	if !data.Tos.IsNull() {
		if false && data.Tos.ValueBool() {
			body, _ = sjson.Set(body, "definition.customizedIpv4RecordFields.collectTos", "")
		} else {
			body, _ = sjson.Set(body, "definition.customizedIpv4RecordFields.collectTos", data.Tos.ValueBool())
		}
	}
	if !data.RemarkedDscp.IsNull() {
		if false && data.RemarkedDscp.ValueBool() {
			body, _ = sjson.Set(body, "definition.customizedIpv4RecordFields.collectDscpOutput", "")
		} else {
			body, _ = sjson.Set(body, "definition.customizedIpv4RecordFields.collectDscpOutput", data.RemarkedDscp.ValueBool())
		}
	}
	if true {
		body, _ = sjson.Set(body, "definition.collectors", []interface{}{})
		for _, item := range data.Collectors {
			itemBody := ""
			if !item.VpnId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "vpn", fmt.Sprint(item.VpnId.ValueInt64()))
			}
			if !item.IpAddress.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "address", item.IpAddress.ValueString())
			}
			if !item.Port.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "port", item.Port.ValueInt64())
			}
			if !item.Transport.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "transport", item.Transport.ValueString())
			}
			if !item.SourceInterface.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sourceInterface", item.SourceInterface.ValueString())
			}
			if !item.ExportSpreading.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "exportSpread", item.ExportSpreading.ValueString())
			}
			body, _ = sjson.SetRaw(body, "definition.collectors.-1", itemBody)
		}
	}
	return body
}

func (data *CflowdPolicyDefinition) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("definition.flowActiveTimeout"); value.Exists() {
		data.ActiveFlowTimeout = types.Int64Value(value.Int())
	} else {
		data.ActiveFlowTimeout = types.Int64Null()
	}
	if value := res.Get("definition.flowInactiveTimeout"); value.Exists() {
		data.InactiveFlowTimeout = types.Int64Value(value.Int())
	} else {
		data.InactiveFlowTimeout = types.Int64Null()
	}
	if value := res.Get("definition.flowSamplingInterval"); value.Exists() {
		data.SamplingInterval = types.Int64Value(value.Int())
	} else {
		data.SamplingInterval = types.Int64Null()
	}
	if value := res.Get("definition.templateRefresh"); value.Exists() {
		data.FlowRefresh = types.Int64Value(value.Int())
	} else {
		data.FlowRefresh = types.Int64Null()
	}
	if value := res.Get("definition.protocol"); value.Exists() {
		data.Protocol = types.StringValue(value.String())
	} else {
		data.Protocol = types.StringNull()
	}
	if value := res.Get("definition.customizedIpv4RecordFields.collectTos"); value.Exists() {
		if false && value.String() == "" {
			data.Tos = types.BoolValue(true)
		} else {
			data.Tos = types.BoolValue(value.Bool())
		}
	} else {
		data.Tos = types.BoolNull()
	}
	if value := res.Get("definition.customizedIpv4RecordFields.collectDscpOutput"); value.Exists() {
		if false && value.String() == "" {
			data.RemarkedDscp = types.BoolValue(true)
		} else {
			data.RemarkedDscp = types.BoolValue(value.Bool())
		}
	} else {
		data.RemarkedDscp = types.BoolNull()
	}
	if value := res.Get("definition.collectors"); value.Exists() && len(value.Array()) > 0 {
		data.Collectors = make([]CflowdPolicyDefinitionCollectors, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CflowdPolicyDefinitionCollectors{}
			if cValue := v.Get("vpn"); cValue.Exists() {
				item.VpnId = types.Int64Value(cValue.Int())
			} else {
				item.VpnId = types.Int64Null()
			}
			if cValue := v.Get("address"); cValue.Exists() {
				item.IpAddress = types.StringValue(cValue.String())
			} else {
				item.IpAddress = types.StringNull()
			}
			if cValue := v.Get("port"); cValue.Exists() {
				item.Port = types.Int64Value(cValue.Int())
			} else {
				item.Port = types.Int64Null()
			}
			if cValue := v.Get("transport"); cValue.Exists() {
				item.Transport = types.StringValue(cValue.String())
			} else {
				item.Transport = types.StringNull()
			}
			if cValue := v.Get("sourceInterface"); cValue.Exists() {
				item.SourceInterface = types.StringValue(cValue.String())
			} else {
				item.SourceInterface = types.StringNull()
			}
			if cValue := v.Get("exportSpread"); cValue.Exists() {
				item.ExportSpreading = types.StringValue(cValue.String())
			} else {
				item.ExportSpreading = types.StringNull()
			}
			data.Collectors = append(data.Collectors, item)
			return true
		})
	} else {
		data.Collectors = []CflowdPolicyDefinitionCollectors{}
	}
}

func (data *CflowdPolicyDefinition) hasChanges(ctx context.Context, state *CflowdPolicyDefinition) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.Description.Equal(state.Description) {
		hasChanges = true
	}
	if !data.ActiveFlowTimeout.Equal(state.ActiveFlowTimeout) {
		hasChanges = true
	}
	if !data.InactiveFlowTimeout.Equal(state.InactiveFlowTimeout) {
		hasChanges = true
	}
	if !data.SamplingInterval.Equal(state.SamplingInterval) {
		hasChanges = true
	}
	if !data.FlowRefresh.Equal(state.FlowRefresh) {
		hasChanges = true
	}
	if !data.Protocol.Equal(state.Protocol) {
		hasChanges = true
	}
	if !data.Tos.Equal(state.Tos) {
		hasChanges = true
	}
	if !data.RemarkedDscp.Equal(state.RemarkedDscp) {
		hasChanges = true
	}
	if len(data.Collectors) != len(state.Collectors) {
		hasChanges = true
	} else {
		for i := range data.Collectors {
			if !data.Collectors[i].VpnId.Equal(state.Collectors[i].VpnId) {
				hasChanges = true
			}
			if !data.Collectors[i].IpAddress.Equal(state.Collectors[i].IpAddress) {
				hasChanges = true
			}
			if !data.Collectors[i].Port.Equal(state.Collectors[i].Port) {
				hasChanges = true
			}
			if !data.Collectors[i].Transport.Equal(state.Collectors[i].Transport) {
				hasChanges = true
			}
			if !data.Collectors[i].SourceInterface.Equal(state.Collectors[i].SourceInterface) {
				hasChanges = true
			}
			if !data.Collectors[i].ExportSpreading.Equal(state.Collectors[i].ExportSpreading) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}
