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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type SecurityPolicy struct {
	Id                         types.String                `tfsdk:"id"`
	Version                    types.Int64                 `tfsdk:"version"`
	Name                       types.String                `tfsdk:"name"`
	Description                types.String                `tfsdk:"description"`
	Mode                       types.String                `tfsdk:"mode"`
	UseCase                    types.String                `tfsdk:"use_case"`
	Definitions                []SecurityPolicyDefinitions `tfsdk:"definitions"`
	DirectInternetApplications types.String                `tfsdk:"direct_internet_applications"`
	TcpSynFloodLimit           types.String                `tfsdk:"tcp_syn_flood_limit"`
	AuditTrail                 types.String                `tfsdk:"audit_trail"`
	MatchStatisticsPerFilter   types.String                `tfsdk:"match_statistics_per_filter"`
	FailureMode                types.String                `tfsdk:"failure_mode"`
	HighSpeedLoggingServerIp   types.String                `tfsdk:"high_speed_logging_server_ip"`
	HighSpeedLoggingVpn        types.String                `tfsdk:"high_speed_logging_vpn"`
	HighSpeedLoggingServerPort types.String                `tfsdk:"high_speed_logging_server_port"`
	Logging                    []SecurityPolicyLogging     `tfsdk:"logging"`
}

type SecurityPolicyDefinitions struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}

type SecurityPolicyLogging struct {
	ExternalSyslogServerIp  types.String `tfsdk:"external_syslog_server_ip"`
	ExternalSyslogServerVpn types.String `tfsdk:"external_syslog_server_vpn"`
}

func (data SecurityPolicy) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "policyType", "feature")
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "policyName", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "policyDescription", data.Description.ValueString())
	}
	if !data.Mode.IsNull() {
		body, _ = sjson.Set(body, "policyMode", data.Mode.ValueString())
	}
	if !data.UseCase.IsNull() {
		body, _ = sjson.Set(body, "policyUseCase", data.UseCase.ValueString())
	}
	if true {
		body, _ = sjson.Set(body, "policyDefinition.assembly", []interface{}{})
		for _, item := range data.Definitions {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "definitionId", item.Id.ValueString())
			}
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.Type.ValueString())
			}
			body, _ = sjson.SetRaw(body, "policyDefinition.assembly.-1", itemBody)
		}
	}
	if !data.DirectInternetApplications.IsNull() {
		body, _ = sjson.Set(body, "policyDefinition.settings.zoneToNozoneInternet", data.DirectInternetApplications.ValueString())
	}
	if !data.TcpSynFloodLimit.IsNull() {
		body, _ = sjson.Set(body, "policyDefinition.settings.tcpSynFloodLimit", data.TcpSynFloodLimit.ValueString())
	}
	if !data.AuditTrail.IsNull() {
		body, _ = sjson.Set(body, "policyDefinition.settings.auditTrail", data.AuditTrail.ValueString())
	}
	if !data.MatchStatisticsPerFilter.IsNull() {
		body, _ = sjson.Set(body, "policyDefinition.settings.platformMatch", data.MatchStatisticsPerFilter.ValueString())
	}
	if !data.FailureMode.IsNull() {
		body, _ = sjson.Set(body, "policyDefinition.settings.failureMode", data.FailureMode.ValueString())
	}
	if !data.HighSpeedLoggingServerIp.IsNull() {
		body, _ = sjson.Set(body, "policyDefinition.settings.highSpeedLogging.serverIp", data.HighSpeedLoggingServerIp.ValueString())
	}
	if !data.HighSpeedLoggingVpn.IsNull() {
		body, _ = sjson.Set(body, "policyDefinition.settings.highSpeedLogging.vrf", data.HighSpeedLoggingVpn.ValueString())
	}
	if !data.HighSpeedLoggingServerPort.IsNull() {
		body, _ = sjson.Set(body, "policyDefinition.settings.highSpeedLogging.port", data.HighSpeedLoggingServerPort.ValueString())
	}
	if true {
		body, _ = sjson.Set(body, "policyDefinition.settings.logging", []interface{}{})
		for _, item := range data.Logging {
			itemBody := ""
			if !item.ExternalSyslogServerIp.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "serverIP", item.ExternalSyslogServerIp.ValueString())
			}
			if !item.ExternalSyslogServerVpn.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "vpn", item.ExternalSyslogServerVpn.ValueString())
			}
			body, _ = sjson.SetRaw(body, "policyDefinition.settings.logging.-1", itemBody)
		}
	}
	return body
}

func (data *SecurityPolicy) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("policyName"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("policyDescription"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("policyMode"); value.Exists() {
		data.Mode = types.StringValue(value.String())
	} else {
		data.Mode = types.StringNull()
	}
	if value := res.Get("policyUseCase"); value.Exists() {
		data.UseCase = types.StringValue(value.String())
	} else {
		data.UseCase = types.StringNull()
	}
	if value := res.Get("policyDefinition.assembly"); value.Exists() && len(value.Array()) > 0 {
		data.Definitions = make([]SecurityPolicyDefinitions, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SecurityPolicyDefinitions{}
			if cValue := v.Get("definitionId"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			if cValue := v.Get("type"); cValue.Exists() {
				item.Type = types.StringValue(cValue.String())
			} else {
				item.Type = types.StringNull()
			}
			data.Definitions = append(data.Definitions, item)
			return true
		})
	} else {
		data.Definitions = []SecurityPolicyDefinitions{}
	}
	if value := res.Get("policyDefinition.settings.zoneToNozoneInternet"); value.Exists() {
		data.DirectInternetApplications = types.StringValue(value.String())
	} else {
		data.DirectInternetApplications = types.StringNull()
	}
	if value := res.Get("policyDefinition.settings.tcpSynFloodLimit"); value.Exists() {
		data.TcpSynFloodLimit = types.StringValue(value.String())
	} else {
		data.TcpSynFloodLimit = types.StringNull()
	}
	if value := res.Get("policyDefinition.settings.auditTrail"); value.Exists() {
		data.AuditTrail = types.StringValue(value.String())
	} else {
		data.AuditTrail = types.StringNull()
	}
	if value := res.Get("policyDefinition.settings.platformMatch"); value.Exists() {
		data.MatchStatisticsPerFilter = types.StringValue(value.String())
	} else {
		data.MatchStatisticsPerFilter = types.StringNull()
	}
	if value := res.Get("policyDefinition.settings.failureMode"); value.Exists() {
		data.FailureMode = types.StringValue(value.String())
	} else {
		data.FailureMode = types.StringNull()
	}
	if value := res.Get("policyDefinition.settings.highSpeedLogging.serverIp"); value.Exists() {
		data.HighSpeedLoggingServerIp = types.StringValue(value.String())
	} else {
		data.HighSpeedLoggingServerIp = types.StringNull()
	}
	if value := res.Get("policyDefinition.settings.highSpeedLogging.vrf"); value.Exists() {
		data.HighSpeedLoggingVpn = types.StringValue(value.String())
	} else {
		data.HighSpeedLoggingVpn = types.StringNull()
	}
	if value := res.Get("policyDefinition.settings.highSpeedLogging.port"); value.Exists() {
		data.HighSpeedLoggingServerPort = types.StringValue(value.String())
	} else {
		data.HighSpeedLoggingServerPort = types.StringNull()
	}
	if value := res.Get("policyDefinition.settings.logging"); value.Exists() && len(value.Array()) > 0 {
		data.Logging = make([]SecurityPolicyLogging, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SecurityPolicyLogging{}
			if cValue := v.Get("serverIP"); cValue.Exists() {
				item.ExternalSyslogServerIp = types.StringValue(cValue.String())
			} else {
				item.ExternalSyslogServerIp = types.StringNull()
			}
			if cValue := v.Get("vpn"); cValue.Exists() {
				item.ExternalSyslogServerVpn = types.StringValue(cValue.String())
			} else {
				item.ExternalSyslogServerVpn = types.StringNull()
			}
			data.Logging = append(data.Logging, item)
			return true
		})
	} else {
		data.Logging = []SecurityPolicyLogging{}
	}
}

func (data *SecurityPolicy) hasChanges(ctx context.Context, state *SecurityPolicy) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.Description.Equal(state.Description) {
		hasChanges = true
	}
	if !data.Mode.Equal(state.Mode) {
		hasChanges = true
	}
	if !data.UseCase.Equal(state.UseCase) {
		hasChanges = true
	}
	if len(data.Definitions) != len(state.Definitions) {
		hasChanges = true
	} else {
		for i := range data.Definitions {
			if !data.Definitions[i].Id.Equal(state.Definitions[i].Id) {
				hasChanges = true
			}
			if !data.Definitions[i].Type.Equal(state.Definitions[i].Type) {
				hasChanges = true
			}
		}
	}
	if !data.DirectInternetApplications.Equal(state.DirectInternetApplications) {
		hasChanges = true
	}
	if !data.TcpSynFloodLimit.Equal(state.TcpSynFloodLimit) {
		hasChanges = true
	}
	if !data.AuditTrail.Equal(state.AuditTrail) {
		hasChanges = true
	}
	if !data.MatchStatisticsPerFilter.Equal(state.MatchStatisticsPerFilter) {
		hasChanges = true
	}
	if !data.FailureMode.Equal(state.FailureMode) {
		hasChanges = true
	}
	if !data.HighSpeedLoggingServerIp.Equal(state.HighSpeedLoggingServerIp) {
		hasChanges = true
	}
	if !data.HighSpeedLoggingVpn.Equal(state.HighSpeedLoggingVpn) {
		hasChanges = true
	}
	if !data.HighSpeedLoggingServerPort.Equal(state.HighSpeedLoggingServerPort) {
		hasChanges = true
	}
	if len(data.Logging) != len(state.Logging) {
		hasChanges = true
	} else {
		for i := range data.Logging {
			if !data.Logging[i].ExternalSyslogServerIp.Equal(state.Logging[i].ExternalSyslogServerIp) {
				hasChanges = true
			}
			if !data.Logging[i].ExternalSyslogServerVpn.Equal(state.Logging[i].ExternalSyslogServerVpn) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}
