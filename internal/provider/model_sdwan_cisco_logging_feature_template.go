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
	"strconv"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type CiscoLogging struct {
	Id                   types.String              `tfsdk:"id"`
	Version              types.Int64               `tfsdk:"version"`
	TemplateType         types.String              `tfsdk:"template_type"`
	Name                 types.String              `tfsdk:"name"`
	Description          types.String              `tfsdk:"description"`
	DeviceTypes          types.Set                 `tfsdk:"device_types"`
	DiskLogging          types.Bool                `tfsdk:"disk_logging"`
	DiskLoggingVariable  types.String              `tfsdk:"disk_logging_variable"`
	MaxSize              types.Int64               `tfsdk:"max_size"`
	MaxSizeVariable      types.String              `tfsdk:"max_size_variable"`
	LogRotations         types.Int64               `tfsdk:"log_rotations"`
	LogRotationsVariable types.String              `tfsdk:"log_rotations_variable"`
	TlsProfiles          []CiscoLoggingTlsProfiles `tfsdk:"tls_profiles"`
	Ipv4Servers          []CiscoLoggingIpv4Servers `tfsdk:"ipv4_servers"`
	Ipv6Servers          []CiscoLoggingIpv6Servers `tfsdk:"ipv6_servers"`
}

type CiscoLoggingTlsProfiles struct {
	Optional                types.Bool   `tfsdk:"optional"`
	Name                    types.String `tfsdk:"name"`
	NameVariable            types.String `tfsdk:"name_variable"`
	Version                 types.String `tfsdk:"version"`
	VersionVariable         types.String `tfsdk:"version_variable"`
	AuthenticationType      types.String `tfsdk:"authentication_type"`
	CiphersuiteList         types.Set    `tfsdk:"ciphersuite_list"`
	CiphersuiteListVariable types.String `tfsdk:"ciphersuite_list_variable"`
}

type CiscoLoggingIpv4Servers struct {
	Optional                types.Bool   `tfsdk:"optional"`
	HostnameIp              types.String `tfsdk:"hostname_ip"`
	HostnameIpVariable      types.String `tfsdk:"hostname_ip_variable"`
	VpnId                   types.Int64  `tfsdk:"vpn_id"`
	VpnIdVariable           types.String `tfsdk:"vpn_id_variable"`
	SourceInterface         types.String `tfsdk:"source_interface"`
	SourceInterfaceVariable types.String `tfsdk:"source_interface_variable"`
	LoggingLevel            types.String `tfsdk:"logging_level"`
	LoggingLevelVariable    types.String `tfsdk:"logging_level_variable"`
	EnableTls               types.Bool   `tfsdk:"enable_tls"`
	EnableTlsVariable       types.String `tfsdk:"enable_tls_variable"`
	CustomProfile           types.Bool   `tfsdk:"custom_profile"`
	CustomProfileVariable   types.String `tfsdk:"custom_profile_variable"`
	Profile                 types.String `tfsdk:"profile"`
	ProfileVariable         types.String `tfsdk:"profile_variable"`
}

type CiscoLoggingIpv6Servers struct {
	Optional                types.Bool   `tfsdk:"optional"`
	HostnameIp              types.String `tfsdk:"hostname_ip"`
	HostnameIpVariable      types.String `tfsdk:"hostname_ip_variable"`
	VpnId                   types.Int64  `tfsdk:"vpn_id"`
	VpnIdVariable           types.String `tfsdk:"vpn_id_variable"`
	SourceInterface         types.String `tfsdk:"source_interface"`
	SourceInterfaceVariable types.String `tfsdk:"source_interface_variable"`
	LoggingLevel            types.String `tfsdk:"logging_level"`
	LoggingLevelVariable    types.String `tfsdk:"logging_level_variable"`
	EnableTls               types.Bool   `tfsdk:"enable_tls"`
	EnableTlsVariable       types.String `tfsdk:"enable_tls_variable"`
	CustomProfile           types.Bool   `tfsdk:"custom_profile"`
	CustomProfileVariable   types.String `tfsdk:"custom_profile_variable"`
	Profile                 types.String `tfsdk:"profile"`
	ProfileVariable         types.String `tfsdk:"profile_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data CiscoLogging) getModel() string {
	return "cisco_logging"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data CiscoLogging) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "cisco_logging")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."

	if !data.DiskLoggingVariable.IsNull() {
		body, _ = sjson.Set(body, path+"disk.enable."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"disk.enable."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"disk.enable."+"vipVariableName", data.DiskLoggingVariable.ValueString())
	} else if data.DiskLogging.IsNull() {
		body, _ = sjson.Set(body, path+"disk.enable."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"disk.enable."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"disk.enable."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"disk.enable."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"disk.enable."+"vipValue", strconv.FormatBool(data.DiskLogging.ValueBool()))
	}

	if !data.MaxSizeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"disk.file.size."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"disk.file.size."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"disk.file.size."+"vipVariableName", data.MaxSizeVariable.ValueString())
	} else if data.MaxSize.IsNull() {
		body, _ = sjson.Set(body, path+"disk.file.size."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"disk.file.size."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"disk.file.size."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"disk.file.size."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"disk.file.size."+"vipValue", data.MaxSize.ValueInt64())
	}

	if !data.LogRotationsVariable.IsNull() {
		body, _ = sjson.Set(body, path+"disk.file.rotate."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"disk.file.rotate."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"disk.file.rotate."+"vipVariableName", data.LogRotationsVariable.ValueString())
	} else if data.LogRotations.IsNull() {
		body, _ = sjson.Set(body, path+"disk.file.rotate."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"disk.file.rotate."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"disk.file.rotate."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"disk.file.rotate."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"disk.file.rotate."+"vipValue", data.LogRotations.ValueInt64())
	}
	if len(data.TlsProfiles) > 0 {
		body, _ = sjson.Set(body, path+"tls-profile."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"tls-profile."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tls-profile."+"vipPrimaryKey", []string{"profile"})
		body, _ = sjson.Set(body, path+"tls-profile."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"tls-profile."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"tls-profile."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"tls-profile."+"vipPrimaryKey", []string{"profile"})
		body, _ = sjson.Set(body, path+"tls-profile."+"vipValue", []interface{}{})
	}
	for _, item := range data.TlsProfiles {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "profile")

		if !item.NameVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "profile."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "profile."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "profile."+"vipVariableName", item.NameVariable.ValueString())
		} else if item.Name.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "profile."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "profile."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "profile."+"vipValue", item.Name.ValueString())
		}
		itemAttributes = append(itemAttributes, "tls-version")

		if !item.VersionVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tls-version.version."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tls-version.version."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "tls-version.version."+"vipVariableName", item.VersionVariable.ValueString())
		} else if item.Version.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tls-version.version."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tls-version.version."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "tls-version.version."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tls-version.version."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "tls-version.version."+"vipValue", item.Version.ValueString())
		}
		itemAttributes = append(itemAttributes, "auth-type")
		if item.AuthenticationType.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "auth-type."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "auth-type."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "auth-type."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "auth-type."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "auth-type."+"vipValue", item.AuthenticationType.ValueString())
		}
		itemAttributes = append(itemAttributes, "ciphersuite")

		if !item.CiphersuiteListVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ciphersuite.ciphersuite-list."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ciphersuite.ciphersuite-list."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "ciphersuite.ciphersuite-list."+"vipVariableName", item.CiphersuiteListVariable.ValueString())
		} else if item.CiphersuiteList.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ciphersuite.ciphersuite-list."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ciphersuite.ciphersuite-list."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "ciphersuite.ciphersuite-list."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ciphersuite.ciphersuite-list."+"vipType", "constant")
			var values []string
			item.CiphersuiteList.ElementsAs(ctx, &values, false)
			itemBody, _ = sjson.Set(itemBody, "ciphersuite.ciphersuite-list."+"vipValue", values)
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"tls-profile."+"vipValue.-1", itemBody)
	}
	if len(data.Ipv4Servers) > 0 {
		body, _ = sjson.Set(body, path+"server."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"server."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"server."+"vipPrimaryKey", []string{"name"})
		body, _ = sjson.Set(body, path+"server."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"server."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"server."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"server."+"vipPrimaryKey", []string{"name"})
		body, _ = sjson.Set(body, path+"server."+"vipValue", []interface{}{})
	}
	for _, item := range data.Ipv4Servers {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "name")

		if !item.HostnameIpVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipVariableName", item.HostnameIpVariable.ValueString())
		} else if item.HostnameIp.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipValue", item.HostnameIp.ValueString())
		}
		itemAttributes = append(itemAttributes, "vpn")

		if !item.VpnIdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipVariableName", item.VpnIdVariable.ValueString())
		} else if item.VpnId.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipValue", item.VpnId.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "source-interface")

		if !item.SourceInterfaceVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipVariableName", item.SourceInterfaceVariable.ValueString())
		} else if item.SourceInterface.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipValue", item.SourceInterface.ValueString())
		}
		itemAttributes = append(itemAttributes, "priority")

		if !item.LoggingLevelVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipVariableName", item.LoggingLevelVariable.ValueString())
		} else if item.LoggingLevel.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipValue", item.LoggingLevel.ValueString())
		}
		itemAttributes = append(itemAttributes, "tls")

		if !item.EnableTlsVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tls.enable-tls."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tls.enable-tls."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "tls.enable-tls."+"vipVariableName", item.EnableTlsVariable.ValueString())
		} else if item.EnableTls.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tls.enable-tls."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tls.enable-tls."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "tls.enable-tls."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tls.enable-tls."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "tls.enable-tls."+"vipValue", strconv.FormatBool(item.EnableTls.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "tls")

		if !item.CustomProfileVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tls.tls-properties.custom-profile."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tls.tls-properties.custom-profile."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "tls.tls-properties.custom-profile."+"vipVariableName", item.CustomProfileVariable.ValueString())
		} else if item.CustomProfile.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tls.tls-properties.custom-profile."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tls.tls-properties.custom-profile."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "tls.tls-properties.custom-profile."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tls.tls-properties.custom-profile."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "tls.tls-properties.custom-profile."+"vipValue", strconv.FormatBool(item.CustomProfile.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "tls")

		if !item.ProfileVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tls.tls-properties.profile."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tls.tls-properties.profile."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "tls.tls-properties.profile."+"vipVariableName", item.ProfileVariable.ValueString())
		} else if item.Profile.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tls.tls-properties.profile."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tls.tls-properties.profile."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "tls.tls-properties.profile."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tls.tls-properties.profile."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "tls.tls-properties.profile."+"vipValue", item.Profile.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"server."+"vipValue.-1", itemBody)
	}
	if len(data.Ipv6Servers) > 0 {
		body, _ = sjson.Set(body, path+"ipv6-server."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ipv6-server."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ipv6-server."+"vipPrimaryKey", []string{"name"})
		body, _ = sjson.Set(body, path+"ipv6-server."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"ipv6-server."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ipv6-server."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"ipv6-server."+"vipPrimaryKey", []string{"name"})
		body, _ = sjson.Set(body, path+"ipv6-server."+"vipValue", []interface{}{})
	}
	for _, item := range data.Ipv6Servers {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "name")

		if !item.HostnameIpVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipVariableName", item.HostnameIpVariable.ValueString())
		} else if item.HostnameIp.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipValue", item.HostnameIp.ValueString())
		}
		itemAttributes = append(itemAttributes, "vpn")

		if !item.VpnIdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipVariableName", item.VpnIdVariable.ValueString())
		} else if item.VpnId.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipValue", item.VpnId.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "source-interface")

		if !item.SourceInterfaceVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipVariableName", item.SourceInterfaceVariable.ValueString())
		} else if item.SourceInterface.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipValue", item.SourceInterface.ValueString())
		}
		itemAttributes = append(itemAttributes, "priority")

		if !item.LoggingLevelVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipVariableName", item.LoggingLevelVariable.ValueString())
		} else if item.LoggingLevel.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipValue", item.LoggingLevel.ValueString())
		}
		itemAttributes = append(itemAttributes, "tls")

		if !item.EnableTlsVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tls.enable-tls."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tls.enable-tls."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "tls.enable-tls."+"vipVariableName", item.EnableTlsVariable.ValueString())
		} else if item.EnableTls.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tls.enable-tls."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tls.enable-tls."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "tls.enable-tls."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tls.enable-tls."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "tls.enable-tls."+"vipValue", strconv.FormatBool(item.EnableTls.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "tls")

		if !item.CustomProfileVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tls.tls-properties.custom-profile."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tls.tls-properties.custom-profile."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "tls.tls-properties.custom-profile."+"vipVariableName", item.CustomProfileVariable.ValueString())
		} else if item.CustomProfile.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tls.tls-properties.custom-profile."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tls.tls-properties.custom-profile."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "tls.tls-properties.custom-profile."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tls.tls-properties.custom-profile."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "tls.tls-properties.custom-profile."+"vipValue", strconv.FormatBool(item.CustomProfile.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "tls")

		if !item.ProfileVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tls.tls-properties.profile."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tls.tls-properties.profile."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "tls.tls-properties.profile."+"vipVariableName", item.ProfileVariable.ValueString())
		} else if item.Profile.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tls.tls-properties.profile."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tls.tls-properties.profile."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "tls.tls-properties.profile."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tls.tls-properties.profile."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "tls.tls-properties.profile."+"vipValue", item.Profile.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"ipv6-server."+"vipValue.-1", itemBody)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *CiscoLogging) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("deviceType"); value.Exists() {
		data.DeviceTypes = helpers.GetStringSet(value.Array())
	} else {
		data.DeviceTypes = types.SetNull(types.StringType)
	}
	if value := res.Get("templateDescription"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("templateName"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("templateType"); value.Exists() {
		data.TemplateType = types.StringValue(value.String())
	} else {
		data.TemplateType = types.StringNull()
	}

	path := "templateDefinition."
	if value := res.Get(path + "disk.enable.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.DiskLogging = types.BoolNull()

			v := res.Get(path + "disk.enable.vipVariableName")
			data.DiskLoggingVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.DiskLogging = types.BoolNull()
			data.DiskLoggingVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "disk.enable.vipValue")
			data.DiskLogging = types.BoolValue(v.Bool())
			data.DiskLoggingVariable = types.StringNull()
		}
	} else {
		data.DiskLogging = types.BoolNull()
		data.DiskLoggingVariable = types.StringNull()
	}
	if value := res.Get(path + "disk.file.size.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.MaxSize = types.Int64Null()

			v := res.Get(path + "disk.file.size.vipVariableName")
			data.MaxSizeVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.MaxSize = types.Int64Null()
			data.MaxSizeVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "disk.file.size.vipValue")
			data.MaxSize = types.Int64Value(v.Int())
			data.MaxSizeVariable = types.StringNull()
		}
	} else {
		data.MaxSize = types.Int64Null()
		data.MaxSizeVariable = types.StringNull()
	}
	if value := res.Get(path + "disk.file.rotate.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.LogRotations = types.Int64Null()

			v := res.Get(path + "disk.file.rotate.vipVariableName")
			data.LogRotationsVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.LogRotations = types.Int64Null()
			data.LogRotationsVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "disk.file.rotate.vipValue")
			data.LogRotations = types.Int64Value(v.Int())
			data.LogRotationsVariable = types.StringNull()
		}
	} else {
		data.LogRotations = types.Int64Null()
		data.LogRotationsVariable = types.StringNull()
	}
	if value := res.Get(path + "tls-profile.vipValue"); len(value.Array()) > 0 {
		data.TlsProfiles = make([]CiscoLoggingTlsProfiles, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoLoggingTlsProfiles{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("profile.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Name = types.StringNull()

					cv := v.Get("profile.vipVariableName")
					item.NameVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Name = types.StringNull()
					item.NameVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("profile.vipValue")
					item.Name = types.StringValue(cv.String())
					item.NameVariable = types.StringNull()
				}
			} else {
				item.Name = types.StringNull()
				item.NameVariable = types.StringNull()
			}
			if cValue := v.Get("tls-version.version.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Version = types.StringNull()

					cv := v.Get("tls-version.version.vipVariableName")
					item.VersionVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Version = types.StringNull()
					item.VersionVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("tls-version.version.vipValue")
					item.Version = types.StringValue(cv.String())
					item.VersionVariable = types.StringNull()
				}
			} else {
				item.Version = types.StringNull()
				item.VersionVariable = types.StringNull()
			}
			if cValue := v.Get("auth-type.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.AuthenticationType = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.AuthenticationType = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("auth-type.vipValue")
					item.AuthenticationType = types.StringValue(cv.String())

				}
			} else {
				item.AuthenticationType = types.StringNull()

			}
			if cValue := v.Get("ciphersuite.ciphersuite-list.vipType"); len(cValue.Array()) > 0 {
				if cValue.String() == "variableName" {
					item.CiphersuiteList = types.SetNull(types.StringType)

					cv := v.Get("ciphersuite.ciphersuite-list.vipVariableName")
					item.CiphersuiteListVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.CiphersuiteList = types.SetNull(types.StringType)
					item.CiphersuiteListVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("ciphersuite.ciphersuite-list.vipValue")
					item.CiphersuiteList = helpers.GetStringSet(cv.Array())
					item.CiphersuiteListVariable = types.StringNull()
				}
			} else {
				item.CiphersuiteList = types.SetNull(types.StringType)
				item.CiphersuiteListVariable = types.StringNull()
			}
			data.TlsProfiles = append(data.TlsProfiles, item)
			return true
		})
	} else {
		if len(data.TlsProfiles) > 0 {
			data.TlsProfiles = []CiscoLoggingTlsProfiles{}
		}
	}
	if value := res.Get(path + "server.vipValue"); len(value.Array()) > 0 {
		data.Ipv4Servers = make([]CiscoLoggingIpv4Servers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoLoggingIpv4Servers{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.HostnameIp = types.StringNull()

					cv := v.Get("name.vipVariableName")
					item.HostnameIpVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.HostnameIp = types.StringNull()
					item.HostnameIpVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("name.vipValue")
					item.HostnameIp = types.StringValue(cv.String())
					item.HostnameIpVariable = types.StringNull()
				}
			} else {
				item.HostnameIp = types.StringNull()
				item.HostnameIpVariable = types.StringNull()
			}
			if cValue := v.Get("vpn.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.VpnId = types.Int64Null()

					cv := v.Get("vpn.vipVariableName")
					item.VpnIdVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.VpnId = types.Int64Null()
					item.VpnIdVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("vpn.vipValue")
					item.VpnId = types.Int64Value(cv.Int())
					item.VpnIdVariable = types.StringNull()
				}
			} else {
				item.VpnId = types.Int64Null()
				item.VpnIdVariable = types.StringNull()
			}
			if cValue := v.Get("source-interface.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SourceInterface = types.StringNull()

					cv := v.Get("source-interface.vipVariableName")
					item.SourceInterfaceVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SourceInterface = types.StringNull()
					item.SourceInterfaceVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("source-interface.vipValue")
					item.SourceInterface = types.StringValue(cv.String())
					item.SourceInterfaceVariable = types.StringNull()
				}
			} else {
				item.SourceInterface = types.StringNull()
				item.SourceInterfaceVariable = types.StringNull()
			}
			if cValue := v.Get("priority.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.LoggingLevel = types.StringNull()

					cv := v.Get("priority.vipVariableName")
					item.LoggingLevelVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.LoggingLevel = types.StringNull()
					item.LoggingLevelVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("priority.vipValue")
					item.LoggingLevel = types.StringValue(cv.String())
					item.LoggingLevelVariable = types.StringNull()
				}
			} else {
				item.LoggingLevel = types.StringNull()
				item.LoggingLevelVariable = types.StringNull()
			}
			if cValue := v.Get("tls.enable-tls.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.EnableTls = types.BoolNull()

					cv := v.Get("tls.enable-tls.vipVariableName")
					item.EnableTlsVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.EnableTls = types.BoolNull()
					item.EnableTlsVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("tls.enable-tls.vipValue")
					item.EnableTls = types.BoolValue(cv.Bool())
					item.EnableTlsVariable = types.StringNull()
				}
			} else {
				item.EnableTls = types.BoolNull()
				item.EnableTlsVariable = types.StringNull()
			}
			if cValue := v.Get("tls.tls-properties.custom-profile.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.CustomProfile = types.BoolNull()

					cv := v.Get("tls.tls-properties.custom-profile.vipVariableName")
					item.CustomProfileVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.CustomProfile = types.BoolNull()
					item.CustomProfileVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("tls.tls-properties.custom-profile.vipValue")
					item.CustomProfile = types.BoolValue(cv.Bool())
					item.CustomProfileVariable = types.StringNull()
				}
			} else {
				item.CustomProfile = types.BoolNull()
				item.CustomProfileVariable = types.StringNull()
			}
			if cValue := v.Get("tls.tls-properties.profile.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Profile = types.StringNull()

					cv := v.Get("tls.tls-properties.profile.vipVariableName")
					item.ProfileVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Profile = types.StringNull()
					item.ProfileVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("tls.tls-properties.profile.vipValue")
					item.Profile = types.StringValue(cv.String())
					item.ProfileVariable = types.StringNull()
				}
			} else {
				item.Profile = types.StringNull()
				item.ProfileVariable = types.StringNull()
			}
			data.Ipv4Servers = append(data.Ipv4Servers, item)
			return true
		})
	} else {
		if len(data.Ipv4Servers) > 0 {
			data.Ipv4Servers = []CiscoLoggingIpv4Servers{}
		}
	}
	if value := res.Get(path + "ipv6-server.vipValue"); len(value.Array()) > 0 {
		data.Ipv6Servers = make([]CiscoLoggingIpv6Servers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoLoggingIpv6Servers{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.HostnameIp = types.StringNull()

					cv := v.Get("name.vipVariableName")
					item.HostnameIpVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.HostnameIp = types.StringNull()
					item.HostnameIpVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("name.vipValue")
					item.HostnameIp = types.StringValue(cv.String())
					item.HostnameIpVariable = types.StringNull()
				}
			} else {
				item.HostnameIp = types.StringNull()
				item.HostnameIpVariable = types.StringNull()
			}
			if cValue := v.Get("vpn.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.VpnId = types.Int64Null()

					cv := v.Get("vpn.vipVariableName")
					item.VpnIdVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.VpnId = types.Int64Null()
					item.VpnIdVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("vpn.vipValue")
					item.VpnId = types.Int64Value(cv.Int())
					item.VpnIdVariable = types.StringNull()
				}
			} else {
				item.VpnId = types.Int64Null()
				item.VpnIdVariable = types.StringNull()
			}
			if cValue := v.Get("source-interface.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SourceInterface = types.StringNull()

					cv := v.Get("source-interface.vipVariableName")
					item.SourceInterfaceVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SourceInterface = types.StringNull()
					item.SourceInterfaceVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("source-interface.vipValue")
					item.SourceInterface = types.StringValue(cv.String())
					item.SourceInterfaceVariable = types.StringNull()
				}
			} else {
				item.SourceInterface = types.StringNull()
				item.SourceInterfaceVariable = types.StringNull()
			}
			if cValue := v.Get("priority.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.LoggingLevel = types.StringNull()

					cv := v.Get("priority.vipVariableName")
					item.LoggingLevelVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.LoggingLevel = types.StringNull()
					item.LoggingLevelVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("priority.vipValue")
					item.LoggingLevel = types.StringValue(cv.String())
					item.LoggingLevelVariable = types.StringNull()
				}
			} else {
				item.LoggingLevel = types.StringNull()
				item.LoggingLevelVariable = types.StringNull()
			}
			if cValue := v.Get("tls.enable-tls.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.EnableTls = types.BoolNull()

					cv := v.Get("tls.enable-tls.vipVariableName")
					item.EnableTlsVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.EnableTls = types.BoolNull()
					item.EnableTlsVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("tls.enable-tls.vipValue")
					item.EnableTls = types.BoolValue(cv.Bool())
					item.EnableTlsVariable = types.StringNull()
				}
			} else {
				item.EnableTls = types.BoolNull()
				item.EnableTlsVariable = types.StringNull()
			}
			if cValue := v.Get("tls.tls-properties.custom-profile.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.CustomProfile = types.BoolNull()

					cv := v.Get("tls.tls-properties.custom-profile.vipVariableName")
					item.CustomProfileVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.CustomProfile = types.BoolNull()
					item.CustomProfileVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("tls.tls-properties.custom-profile.vipValue")
					item.CustomProfile = types.BoolValue(cv.Bool())
					item.CustomProfileVariable = types.StringNull()
				}
			} else {
				item.CustomProfile = types.BoolNull()
				item.CustomProfileVariable = types.StringNull()
			}
			if cValue := v.Get("tls.tls-properties.profile.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Profile = types.StringNull()

					cv := v.Get("tls.tls-properties.profile.vipVariableName")
					item.ProfileVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Profile = types.StringNull()
					item.ProfileVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("tls.tls-properties.profile.vipValue")
					item.Profile = types.StringValue(cv.String())
					item.ProfileVariable = types.StringNull()
				}
			} else {
				item.Profile = types.StringNull()
				item.ProfileVariable = types.StringNull()
			}
			data.Ipv6Servers = append(data.Ipv6Servers, item)
			return true
		})
	} else {
		if len(data.Ipv6Servers) > 0 {
			data.Ipv6Servers = []CiscoLoggingIpv6Servers{}
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *CiscoLogging) hasChanges(ctx context.Context, state *CiscoLogging) bool {
	hasChanges := false
	if !data.DiskLogging.Equal(state.DiskLogging) {
		hasChanges = true
	}
	if !data.MaxSize.Equal(state.MaxSize) {
		hasChanges = true
	}
	if !data.LogRotations.Equal(state.LogRotations) {
		hasChanges = true
	}
	if len(data.TlsProfiles) != len(state.TlsProfiles) {
		hasChanges = true
	} else {
		for i := range data.TlsProfiles {
			if !data.TlsProfiles[i].Name.Equal(state.TlsProfiles[i].Name) {
				hasChanges = true
			}
			if !data.TlsProfiles[i].Version.Equal(state.TlsProfiles[i].Version) {
				hasChanges = true
			}
			if !data.TlsProfiles[i].AuthenticationType.Equal(state.TlsProfiles[i].AuthenticationType) {
				hasChanges = true
			}
			if !data.TlsProfiles[i].CiphersuiteList.Equal(state.TlsProfiles[i].CiphersuiteList) {
				hasChanges = true
			}
		}
	}
	if len(data.Ipv4Servers) != len(state.Ipv4Servers) {
		hasChanges = true
	} else {
		for i := range data.Ipv4Servers {
			if !data.Ipv4Servers[i].HostnameIp.Equal(state.Ipv4Servers[i].HostnameIp) {
				hasChanges = true
			}
			if !data.Ipv4Servers[i].VpnId.Equal(state.Ipv4Servers[i].VpnId) {
				hasChanges = true
			}
			if !data.Ipv4Servers[i].SourceInterface.Equal(state.Ipv4Servers[i].SourceInterface) {
				hasChanges = true
			}
			if !data.Ipv4Servers[i].LoggingLevel.Equal(state.Ipv4Servers[i].LoggingLevel) {
				hasChanges = true
			}
			if !data.Ipv4Servers[i].EnableTls.Equal(state.Ipv4Servers[i].EnableTls) {
				hasChanges = true
			}
			if !data.Ipv4Servers[i].CustomProfile.Equal(state.Ipv4Servers[i].CustomProfile) {
				hasChanges = true
			}
			if !data.Ipv4Servers[i].Profile.Equal(state.Ipv4Servers[i].Profile) {
				hasChanges = true
			}
		}
	}
	if len(data.Ipv6Servers) != len(state.Ipv6Servers) {
		hasChanges = true
	} else {
		for i := range data.Ipv6Servers {
			if !data.Ipv6Servers[i].HostnameIp.Equal(state.Ipv6Servers[i].HostnameIp) {
				hasChanges = true
			}
			if !data.Ipv6Servers[i].VpnId.Equal(state.Ipv6Servers[i].VpnId) {
				hasChanges = true
			}
			if !data.Ipv6Servers[i].SourceInterface.Equal(state.Ipv6Servers[i].SourceInterface) {
				hasChanges = true
			}
			if !data.Ipv6Servers[i].LoggingLevel.Equal(state.Ipv6Servers[i].LoggingLevel) {
				hasChanges = true
			}
			if !data.Ipv6Servers[i].EnableTls.Equal(state.Ipv6Servers[i].EnableTls) {
				hasChanges = true
			}
			if !data.Ipv6Servers[i].CustomProfile.Equal(state.Ipv6Servers[i].CustomProfile) {
				hasChanges = true
			}
			if !data.Ipv6Servers[i].Profile.Equal(state.Ipv6Servers[i].Profile) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges
