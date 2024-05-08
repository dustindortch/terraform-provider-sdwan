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
type CiscoSecurity struct {
	Id                         types.String             `tfsdk:"id"`
	Version                    types.Int64              `tfsdk:"version"`
	TemplateType               types.String             `tfsdk:"template_type"`
	Name                       types.String             `tfsdk:"name"`
	Description                types.String             `tfsdk:"description"`
	DeviceTypes                types.Set                `tfsdk:"device_types"`
	RekeyInterval              types.Int64              `tfsdk:"rekey_interval"`
	RekeyIntervalVariable      types.String             `tfsdk:"rekey_interval_variable"`
	ReplayWindow               types.String             `tfsdk:"replay_window"`
	ReplayWindowVariable       types.String             `tfsdk:"replay_window_variable"`
	ExtendedArWindow           types.Int64              `tfsdk:"extended_ar_window"`
	ExtendedArWindowVariable   types.String             `tfsdk:"extended_ar_window_variable"`
	AuthenticationType         types.Set                `tfsdk:"authentication_type"`
	AuthenticationTypeVariable types.String             `tfsdk:"authentication_type_variable"`
	IntegrityType              types.Set                `tfsdk:"integrity_type"`
	IntegrityTypeVariable      types.String             `tfsdk:"integrity_type_variable"`
	PairwiseKeying             types.Bool               `tfsdk:"pairwise_keying"`
	PairwiseKeyingVariable     types.String             `tfsdk:"pairwise_keying_variable"`
	Keychains                  []CiscoSecurityKeychains `tfsdk:"keychains"`
	Keys                       []CiscoSecurityKeys      `tfsdk:"keys"`
}

type CiscoSecurityKeychains struct {
	Optional types.Bool   `tfsdk:"optional"`
	Name     types.String `tfsdk:"name"`
	KeyId    types.Int64  `tfsdk:"key_id"`
}

type CiscoSecurityKeys struct {
	Optional                       types.Bool   `tfsdk:"optional"`
	Id                             types.String `tfsdk:"id"`
	ChainName                      types.String `tfsdk:"chain_name"`
	SendId                         types.Int64  `tfsdk:"send_id"`
	SendIdVariable                 types.String `tfsdk:"send_id_variable"`
	ReceiveId                      types.Int64  `tfsdk:"receive_id"`
	ReceiveIdVariable              types.String `tfsdk:"receive_id_variable"`
	CryptoAlgorithm                types.String `tfsdk:"crypto_algorithm"`
	KeyString                      types.String `tfsdk:"key_string"`
	KeyStringVariable              types.String `tfsdk:"key_string_variable"`
	SendLifetimeLocal              types.Bool   `tfsdk:"send_lifetime_local"`
	SendLifetimeLocalVariable      types.String `tfsdk:"send_lifetime_local_variable"`
	SendLifetimeStartTime          types.String `tfsdk:"send_lifetime_start_time"`
	SendLifetimeEndTimeFormat      types.String `tfsdk:"send_lifetime_end_time_format"`
	SendLifetimeDuration           types.Int64  `tfsdk:"send_lifetime_duration"`
	SendLifetimeDurationVariable   types.String `tfsdk:"send_lifetime_duration_variable"`
	SendLifetimeEndTime            types.String `tfsdk:"send_lifetime_end_time"`
	SendLifetimeInfinite           types.Bool   `tfsdk:"send_lifetime_infinite"`
	SendLifetimeInfiniteVariable   types.String `tfsdk:"send_lifetime_infinite_variable"`
	AcceptLifetimeLocal            types.Bool   `tfsdk:"accept_lifetime_local"`
	AcceptLifetimeLocalVariable    types.String `tfsdk:"accept_lifetime_local_variable"`
	AcceptLifetimeStartTime        types.String `tfsdk:"accept_lifetime_start_time"`
	AcceptLifetimeEndTimeFormat    types.String `tfsdk:"accept_lifetime_end_time_format"`
	AcceptLifetimeDuration         types.Int64  `tfsdk:"accept_lifetime_duration"`
	AcceptLifetimeDurationVariable types.String `tfsdk:"accept_lifetime_duration_variable"`
	AcceptLifetimeEndTime          types.String `tfsdk:"accept_lifetime_end_time"`
	AcceptLifetimeInfinite         types.Bool   `tfsdk:"accept_lifetime_infinite"`
	AcceptLifetimeInfiniteVariable types.String `tfsdk:"accept_lifetime_infinite_variable"`
	IncludeTcpOptions              types.Bool   `tfsdk:"include_tcp_options"`
	IncludeTcpOptionsVariable      types.String `tfsdk:"include_tcp_options_variable"`
	AcceptAoMismatch               types.Bool   `tfsdk:"accept_ao_mismatch"`
	AcceptAoMismatchVariable       types.String `tfsdk:"accept_ao_mismatch_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data CiscoSecurity) getModel() string {
	return "cisco_security"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data CiscoSecurity) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "cisco_security")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."

	if !data.RekeyIntervalVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ipsec.rekey."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipsec.rekey."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ipsec.rekey."+"vipVariableName", data.RekeyIntervalVariable.ValueString())
	} else if data.RekeyInterval.IsNull() {
		body, _ = sjson.Set(body, path+"ipsec.rekey."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipsec.rekey."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ipsec.rekey."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipsec.rekey."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ipsec.rekey."+"vipValue", data.RekeyInterval.ValueInt64())
	}

	if !data.ReplayWindowVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ipsec.replay-window."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipsec.replay-window."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ipsec.replay-window."+"vipVariableName", data.ReplayWindowVariable.ValueString())
	} else if data.ReplayWindow.IsNull() {
		body, _ = sjson.Set(body, path+"ipsec.replay-window."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipsec.replay-window."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ipsec.replay-window."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipsec.replay-window."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ipsec.replay-window."+"vipValue", data.ReplayWindow.ValueString())
	}

	if !data.ExtendedArWindowVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ipsec.extended-ar-window."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipsec.extended-ar-window."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ipsec.extended-ar-window."+"vipVariableName", data.ExtendedArWindowVariable.ValueString())
	} else if data.ExtendedArWindow.IsNull() {
		body, _ = sjson.Set(body, path+"ipsec.extended-ar-window."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipsec.extended-ar-window."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ipsec.extended-ar-window."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipsec.extended-ar-window."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ipsec.extended-ar-window."+"vipValue", data.ExtendedArWindow.ValueInt64())
	}

	if !data.AuthenticationTypeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ipsec.authentication-type."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"ipsec.authentication-type."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ipsec.authentication-type."+"vipVariableName", data.AuthenticationTypeVariable.ValueString())
	} else if data.AuthenticationType.IsNull() {
		body, _ = sjson.Set(body, path+"ipsec.authentication-type."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"ipsec.authentication-type."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ipsec.authentication-type."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"ipsec.authentication-type."+"vipType", "constant")
		var values []string
		data.AuthenticationType.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, path+"ipsec.authentication-type."+"vipValue", values)
	}

	if !data.IntegrityTypeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ipsec.integrity-type."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"ipsec.integrity-type."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ipsec.integrity-type."+"vipVariableName", data.IntegrityTypeVariable.ValueString())
	} else if data.IntegrityType.IsNull() {
		body, _ = sjson.Set(body, path+"ipsec.integrity-type."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"ipsec.integrity-type."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ipsec.integrity-type."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"ipsec.integrity-type."+"vipType", "constant")
		var values []string
		data.IntegrityType.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, path+"ipsec.integrity-type."+"vipValue", values)
	}

	if !data.PairwiseKeyingVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ipsec.pairwise-keying."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipsec.pairwise-keying."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ipsec.pairwise-keying."+"vipVariableName", data.PairwiseKeyingVariable.ValueString())
	} else if data.PairwiseKeying.IsNull() {
		body, _ = sjson.Set(body, path+"ipsec.pairwise-keying."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipsec.pairwise-keying."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ipsec.pairwise-keying."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipsec.pairwise-keying."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ipsec.pairwise-keying."+"vipValue", strconv.FormatBool(data.PairwiseKeying.ValueBool()))
	}
	if len(data.Keychains) > 0 {
		body, _ = sjson.Set(body, path+"trustsec.keychain."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"trustsec.keychain."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"trustsec.keychain."+"vipPrimaryKey", []string{"name", "keyid"})
		body, _ = sjson.Set(body, path+"trustsec.keychain."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"trustsec.keychain."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"trustsec.keychain."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"trustsec.keychain."+"vipPrimaryKey", []string{"name", "keyid"})
		body, _ = sjson.Set(body, path+"trustsec.keychain."+"vipValue", []interface{}{})
	}
	for _, item := range data.Keychains {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "name")
		if item.Name.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipValue", item.Name.ValueString())
		}
		itemAttributes = append(itemAttributes, "keyid")
		if item.KeyId.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "keyid."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "keyid."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "keyid."+"vipValue", item.KeyId.ValueInt64())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"trustsec.keychain."+"vipValue.-1", itemBody)
	}
	if len(data.Keys) > 0 {
		body, _ = sjson.Set(body, path+"key."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"key."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"key."+"vipPrimaryKey", []string{"id", "chain-name"})
		body, _ = sjson.Set(body, path+"key."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"key."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"key."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"key."+"vipPrimaryKey", []string{"id", "chain-name"})
		body, _ = sjson.Set(body, path+"key."+"vipValue", []interface{}{})
	}
	for _, item := range data.Keys {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "id")
		if item.Id.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "id."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "id."+"vipValue", item.Id.ValueString())
		}
		itemAttributes = append(itemAttributes, "chain-name")
		if item.ChainName.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "chain-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "chain-name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "chain-name."+"vipValue", item.ChainName.ValueString())
		}
		itemAttributes = append(itemAttributes, "send-id")

		if !item.SendIdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "send-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "send-id."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "send-id."+"vipVariableName", item.SendIdVariable.ValueString())
		} else if item.SendId.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "send-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "send-id."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "send-id."+"vipValue", item.SendId.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "recv-id")

		if !item.ReceiveIdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "recv-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "recv-id."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "recv-id."+"vipVariableName", item.ReceiveIdVariable.ValueString())
		} else if item.ReceiveId.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "recv-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "recv-id."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "recv-id."+"vipValue", item.ReceiveId.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "cryptographic-algorithm-choice")
		if item.CryptoAlgorithm.IsNull() {
			if !gjson.Get(itemBody, "cryptographic-algorithm-choice").Exists() {
				itemBody, _ = sjson.Set(itemBody, "cryptographic-algorithm-choice", map[string]interface{}{})
			}
		} else {
			itemBody, _ = sjson.Set(itemBody, "cryptographic-algorithm-choice.tcp."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "cryptographic-algorithm-choice.tcp."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "cryptographic-algorithm-choice.tcp."+"vipValue", item.CryptoAlgorithm.ValueString())
		}
		itemAttributes = append(itemAttributes, "key-string")

		if !item.KeyStringVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "key-string."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "key-string."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "key-string."+"vipVariableName", item.KeyStringVariable.ValueString())
		} else if item.KeyString.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "key-string."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "key-string."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "key-string."+"vipValue", item.KeyString.ValueString())
		}
		itemAttributes = append(itemAttributes, "send-lifetime")

		if !item.SendLifetimeLocalVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "send-lifetime.lifetime-group-v1.local."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "send-lifetime.lifetime-group-v1.local."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "send-lifetime.lifetime-group-v1.local."+"vipVariableName", item.SendLifetimeLocalVariable.ValueString())
		} else if item.SendLifetimeLocal.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "send-lifetime.lifetime-group-v1.local."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "send-lifetime.lifetime-group-v1.local."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "send-lifetime.lifetime-group-v1.local."+"vipValue", strconv.FormatBool(item.SendLifetimeLocal.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "send-lifetime")
		if item.SendLifetimeStartTime.IsNull() {
			if !gjson.Get(itemBody, "send-lifetime.lifetime-group-v1").Exists() {
				itemBody, _ = sjson.Set(itemBody, "send-lifetime.lifetime-group-v1", map[string]interface{}{})
			}
		} else {
			itemBody, _ = sjson.Set(itemBody, "send-lifetime.lifetime-group-v1.start-epoch."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "send-lifetime.lifetime-group-v1.start-epoch."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "send-lifetime.lifetime-group-v1.start-epoch."+"vipValue", item.SendLifetimeStartTime.ValueString())
		}
		itemAttributes = append(itemAttributes, "send-lifetime")
		if item.SendLifetimeEndTimeFormat.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "send-lifetime.lifetime-group-v1.end-choice."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "send-lifetime.lifetime-group-v1.end-choice."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "send-lifetime.lifetime-group-v1.end-choice."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "send-lifetime.lifetime-group-v1.end-choice."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "send-lifetime.lifetime-group-v1.end-choice."+"vipValue", item.SendLifetimeEndTimeFormat.ValueString())
		}
		itemAttributes = append(itemAttributes, "send-lifetime")

		if !item.SendLifetimeDurationVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "send-lifetime.lifetime-group-v1.duration."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "send-lifetime.lifetime-group-v1.duration."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "send-lifetime.lifetime-group-v1.duration."+"vipVariableName", item.SendLifetimeDurationVariable.ValueString())
		} else if item.SendLifetimeDuration.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "send-lifetime.lifetime-group-v1.duration."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "send-lifetime.lifetime-group-v1.duration."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "send-lifetime.lifetime-group-v1.duration."+"vipValue", item.SendLifetimeDuration.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "send-lifetime")
		if item.SendLifetimeEndTime.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "send-lifetime.lifetime-group-v1.end-epoch."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "send-lifetime.lifetime-group-v1.end-epoch."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "send-lifetime.lifetime-group-v1.end-epoch."+"vipValue", item.SendLifetimeEndTime.ValueString())
		}
		itemAttributes = append(itemAttributes, "send-lifetime")

		if !item.SendLifetimeInfiniteVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "send-lifetime.lifetime-group-v1.infinite."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "send-lifetime.lifetime-group-v1.infinite."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "send-lifetime.lifetime-group-v1.infinite."+"vipVariableName", item.SendLifetimeInfiniteVariable.ValueString())
		} else if item.SendLifetimeInfinite.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "send-lifetime.lifetime-group-v1.infinite."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "send-lifetime.lifetime-group-v1.infinite."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "send-lifetime.lifetime-group-v1.infinite."+"vipValue", strconv.FormatBool(item.SendLifetimeInfinite.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "accept-lifetime")

		if !item.AcceptLifetimeLocalVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "accept-lifetime.lifetime-group-v1.local."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "accept-lifetime.lifetime-group-v1.local."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "accept-lifetime.lifetime-group-v1.local."+"vipVariableName", item.AcceptLifetimeLocalVariable.ValueString())
		} else if item.AcceptLifetimeLocal.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "accept-lifetime.lifetime-group-v1.local."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "accept-lifetime.lifetime-group-v1.local."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "accept-lifetime.lifetime-group-v1.local."+"vipValue", strconv.FormatBool(item.AcceptLifetimeLocal.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "accept-lifetime")
		if item.AcceptLifetimeStartTime.IsNull() {
			if !gjson.Get(itemBody, "accept-lifetime.lifetime-group-v1").Exists() {
				itemBody, _ = sjson.Set(itemBody, "accept-lifetime.lifetime-group-v1", map[string]interface{}{})
			}
		} else {
			itemBody, _ = sjson.Set(itemBody, "accept-lifetime.lifetime-group-v1.start-epoch."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "accept-lifetime.lifetime-group-v1.start-epoch."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "accept-lifetime.lifetime-group-v1.start-epoch."+"vipValue", item.AcceptLifetimeStartTime.ValueString())
		}
		itemAttributes = append(itemAttributes, "accept-lifetime")
		if item.AcceptLifetimeEndTimeFormat.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "accept-lifetime.lifetime-group-v1.end-choice."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "accept-lifetime.lifetime-group-v1.end-choice."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "accept-lifetime.lifetime-group-v1.end-choice."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "accept-lifetime.lifetime-group-v1.end-choice."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "accept-lifetime.lifetime-group-v1.end-choice."+"vipValue", item.AcceptLifetimeEndTimeFormat.ValueString())
		}
		itemAttributes = append(itemAttributes, "accept-lifetime")

		if !item.AcceptLifetimeDurationVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "accept-lifetime.lifetime-group-v1.duration."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "accept-lifetime.lifetime-group-v1.duration."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "accept-lifetime.lifetime-group-v1.duration."+"vipVariableName", item.AcceptLifetimeDurationVariable.ValueString())
		} else if item.AcceptLifetimeDuration.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "accept-lifetime.lifetime-group-v1.duration."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "accept-lifetime.lifetime-group-v1.duration."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "accept-lifetime.lifetime-group-v1.duration."+"vipValue", item.AcceptLifetimeDuration.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "accept-lifetime")
		if item.AcceptLifetimeEndTime.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "accept-lifetime.lifetime-group-v1.end-epoch."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "accept-lifetime.lifetime-group-v1.end-epoch."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "accept-lifetime.lifetime-group-v1.end-epoch."+"vipValue", item.AcceptLifetimeEndTime.ValueString())
		}
		itemAttributes = append(itemAttributes, "accept-lifetime")

		if !item.AcceptLifetimeInfiniteVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "accept-lifetime.lifetime-group-v1.infinite."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "accept-lifetime.lifetime-group-v1.infinite."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "accept-lifetime.lifetime-group-v1.infinite."+"vipVariableName", item.AcceptLifetimeInfiniteVariable.ValueString())
		} else if item.AcceptLifetimeInfinite.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "accept-lifetime.lifetime-group-v1.infinite."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "accept-lifetime.lifetime-group-v1.infinite."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "accept-lifetime.lifetime-group-v1.infinite."+"vipValue", strconv.FormatBool(item.AcceptLifetimeInfinite.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "include-tcp-options")

		if !item.IncludeTcpOptionsVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "include-tcp-options."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "include-tcp-options."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "include-tcp-options."+"vipVariableName", item.IncludeTcpOptionsVariable.ValueString())
		} else if item.IncludeTcpOptions.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "include-tcp-options."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "include-tcp-options."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "include-tcp-options."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "include-tcp-options."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "include-tcp-options."+"vipValue", strconv.FormatBool(item.IncludeTcpOptions.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "accept-ao-mismatch")

		if !item.AcceptAoMismatchVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "accept-ao-mismatch."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "accept-ao-mismatch."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "accept-ao-mismatch."+"vipVariableName", item.AcceptAoMismatchVariable.ValueString())
		} else if item.AcceptAoMismatch.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "accept-ao-mismatch."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "accept-ao-mismatch."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "accept-ao-mismatch."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "accept-ao-mismatch."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "accept-ao-mismatch."+"vipValue", strconv.FormatBool(item.AcceptAoMismatch.ValueBool()))
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"key."+"vipValue.-1", itemBody)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *CiscoSecurity) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(path + "ipsec.rekey.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.RekeyInterval = types.Int64Null()

			v := res.Get(path + "ipsec.rekey.vipVariableName")
			data.RekeyIntervalVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.RekeyInterval = types.Int64Null()
			data.RekeyIntervalVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ipsec.rekey.vipValue")
			data.RekeyInterval = types.Int64Value(v.Int())
			data.RekeyIntervalVariable = types.StringNull()
		}
	} else {
		data.RekeyInterval = types.Int64Null()
		data.RekeyIntervalVariable = types.StringNull()
	}
	if value := res.Get(path + "ipsec.replay-window.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.ReplayWindow = types.StringNull()

			v := res.Get(path + "ipsec.replay-window.vipVariableName")
			data.ReplayWindowVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.ReplayWindow = types.StringNull()
			data.ReplayWindowVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ipsec.replay-window.vipValue")
			data.ReplayWindow = types.StringValue(v.String())
			data.ReplayWindowVariable = types.StringNull()
		}
	} else {
		data.ReplayWindow = types.StringNull()
		data.ReplayWindowVariable = types.StringNull()
	}
	if value := res.Get(path + "ipsec.extended-ar-window.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.ExtendedArWindow = types.Int64Null()

			v := res.Get(path + "ipsec.extended-ar-window.vipVariableName")
			data.ExtendedArWindowVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.ExtendedArWindow = types.Int64Null()
			data.ExtendedArWindowVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ipsec.extended-ar-window.vipValue")
			data.ExtendedArWindow = types.Int64Value(v.Int())
			data.ExtendedArWindowVariable = types.StringNull()
		}
	} else {
		data.ExtendedArWindow = types.Int64Null()
		data.ExtendedArWindowVariable = types.StringNull()
	}
	if value := res.Get(path + "ipsec.authentication-type.vipType"); len(value.Array()) > 0 {
		if value.String() == "variableName" {
			data.AuthenticationType = types.SetNull(types.StringType)

			v := res.Get(path + "ipsec.authentication-type.vipVariableName")
			data.AuthenticationTypeVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.AuthenticationType = types.SetNull(types.StringType)
			data.AuthenticationTypeVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ipsec.authentication-type.vipValue")
			data.AuthenticationType = helpers.GetStringSet(v.Array())
			data.AuthenticationTypeVariable = types.StringNull()
		}
	} else {
		data.AuthenticationType = types.SetNull(types.StringType)
		data.AuthenticationTypeVariable = types.StringNull()
	}
	if value := res.Get(path + "ipsec.integrity-type.vipType"); len(value.Array()) > 0 {
		if value.String() == "variableName" {
			data.IntegrityType = types.SetNull(types.StringType)

			v := res.Get(path + "ipsec.integrity-type.vipVariableName")
			data.IntegrityTypeVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.IntegrityType = types.SetNull(types.StringType)
			data.IntegrityTypeVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ipsec.integrity-type.vipValue")
			data.IntegrityType = helpers.GetStringSet(v.Array())
			data.IntegrityTypeVariable = types.StringNull()
		}
	} else {
		data.IntegrityType = types.SetNull(types.StringType)
		data.IntegrityTypeVariable = types.StringNull()
	}
	if value := res.Get(path + "ipsec.pairwise-keying.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.PairwiseKeying = types.BoolNull()

			v := res.Get(path + "ipsec.pairwise-keying.vipVariableName")
			data.PairwiseKeyingVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.PairwiseKeying = types.BoolNull()
			data.PairwiseKeyingVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ipsec.pairwise-keying.vipValue")
			data.PairwiseKeying = types.BoolValue(v.Bool())
			data.PairwiseKeyingVariable = types.StringNull()
		}
	} else {
		data.PairwiseKeying = types.BoolNull()
		data.PairwiseKeyingVariable = types.StringNull()
	}
	if value := res.Get(path + "trustsec.keychain.vipValue"); len(value.Array()) > 0 {
		data.Keychains = make([]CiscoSecurityKeychains, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoSecurityKeychains{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Name = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.Name = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("name.vipValue")
					item.Name = types.StringValue(cv.String())

				}
			} else {
				item.Name = types.StringNull()

			}
			if cValue := v.Get("keyid.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.KeyId = types.Int64Null()

				} else if cValue.String() == "ignore" {
					item.KeyId = types.Int64Null()

				} else if cValue.String() == "constant" {
					cv := v.Get("keyid.vipValue")
					item.KeyId = types.Int64Value(cv.Int())

				}
			} else {
				item.KeyId = types.Int64Null()

			}
			data.Keychains = append(data.Keychains, item)
			return true
		})
	} else {
		if len(data.Keychains) > 0 {
			data.Keychains = []CiscoSecurityKeychains{}
		}
	}
	if value := res.Get(path + "key.vipValue"); len(value.Array()) > 0 {
		data.Keys = make([]CiscoSecurityKeys, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoSecurityKeys{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("id.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Id = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.Id = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("id.vipValue")
					item.Id = types.StringValue(cv.String())

				}
			} else {
				item.Id = types.StringNull()

			}
			if cValue := v.Get("chain-name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ChainName = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.ChainName = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("chain-name.vipValue")
					item.ChainName = types.StringValue(cv.String())

				}
			} else {
				item.ChainName = types.StringNull()

			}
			if cValue := v.Get("send-id.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SendId = types.Int64Null()

					cv := v.Get("send-id.vipVariableName")
					item.SendIdVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SendId = types.Int64Null()
					item.SendIdVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("send-id.vipValue")
					item.SendId = types.Int64Value(cv.Int())
					item.SendIdVariable = types.StringNull()
				}
			} else {
				item.SendId = types.Int64Null()
				item.SendIdVariable = types.StringNull()
			}
			if cValue := v.Get("recv-id.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ReceiveId = types.Int64Null()

					cv := v.Get("recv-id.vipVariableName")
					item.ReceiveIdVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.ReceiveId = types.Int64Null()
					item.ReceiveIdVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("recv-id.vipValue")
					item.ReceiveId = types.Int64Value(cv.Int())
					item.ReceiveIdVariable = types.StringNull()
				}
			} else {
				item.ReceiveId = types.Int64Null()
				item.ReceiveIdVariable = types.StringNull()
			}
			if cValue := v.Get("cryptographic-algorithm-choice.tcp.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.CryptoAlgorithm = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.CryptoAlgorithm = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("cryptographic-algorithm-choice.tcp.vipValue")
					item.CryptoAlgorithm = types.StringValue(cv.String())

				}
			} else {
				item.CryptoAlgorithm = types.StringNull()

			}
			if cValue := v.Get("key-string.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.KeyString = types.StringNull()

					cv := v.Get("key-string.vipVariableName")
					item.KeyStringVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.KeyString = types.StringNull()
					item.KeyStringVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("key-string.vipValue")
					item.KeyString = types.StringValue(cv.String())
					item.KeyStringVariable = types.StringNull()
				}
			} else {
				item.KeyString = types.StringNull()
				item.KeyStringVariable = types.StringNull()
			}
			if cValue := v.Get("send-lifetime.lifetime-group-v1.local.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SendLifetimeLocal = types.BoolNull()

					cv := v.Get("send-lifetime.lifetime-group-v1.local.vipVariableName")
					item.SendLifetimeLocalVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SendLifetimeLocal = types.BoolNull()
					item.SendLifetimeLocalVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("send-lifetime.lifetime-group-v1.local.vipValue")
					item.SendLifetimeLocal = types.BoolValue(cv.Bool())
					item.SendLifetimeLocalVariable = types.StringNull()
				}
			} else {
				item.SendLifetimeLocal = types.BoolNull()
				item.SendLifetimeLocalVariable = types.StringNull()
			}
			if cValue := v.Get("send-lifetime.lifetime-group-v1.start-epoch.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SendLifetimeStartTime = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.SendLifetimeStartTime = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("send-lifetime.lifetime-group-v1.start-epoch.vipValue")
					item.SendLifetimeStartTime = types.StringValue(cv.String())

				}
			} else {
				item.SendLifetimeStartTime = types.StringNull()

			}
			if cValue := v.Get("send-lifetime.lifetime-group-v1.end-choice.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SendLifetimeEndTimeFormat = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.SendLifetimeEndTimeFormat = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("send-lifetime.lifetime-group-v1.end-choice.vipValue")
					item.SendLifetimeEndTimeFormat = types.StringValue(cv.String())

				}
			} else {
				item.SendLifetimeEndTimeFormat = types.StringNull()

			}
			if cValue := v.Get("send-lifetime.lifetime-group-v1.duration.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SendLifetimeDuration = types.Int64Null()

					cv := v.Get("send-lifetime.lifetime-group-v1.duration.vipVariableName")
					item.SendLifetimeDurationVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SendLifetimeDuration = types.Int64Null()
					item.SendLifetimeDurationVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("send-lifetime.lifetime-group-v1.duration.vipValue")
					item.SendLifetimeDuration = types.Int64Value(cv.Int())
					item.SendLifetimeDurationVariable = types.StringNull()
				}
			} else {
				item.SendLifetimeDuration = types.Int64Null()
				item.SendLifetimeDurationVariable = types.StringNull()
			}
			if cValue := v.Get("send-lifetime.lifetime-group-v1.end-epoch.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SendLifetimeEndTime = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.SendLifetimeEndTime = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("send-lifetime.lifetime-group-v1.end-epoch.vipValue")
					item.SendLifetimeEndTime = types.StringValue(cv.String())

				}
			} else {
				item.SendLifetimeEndTime = types.StringNull()

			}
			if cValue := v.Get("send-lifetime.lifetime-group-v1.infinite.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SendLifetimeInfinite = types.BoolNull()

					cv := v.Get("send-lifetime.lifetime-group-v1.infinite.vipVariableName")
					item.SendLifetimeInfiniteVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SendLifetimeInfinite = types.BoolNull()
					item.SendLifetimeInfiniteVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("send-lifetime.lifetime-group-v1.infinite.vipValue")
					item.SendLifetimeInfinite = types.BoolValue(cv.Bool())
					item.SendLifetimeInfiniteVariable = types.StringNull()
				}
			} else {
				item.SendLifetimeInfinite = types.BoolNull()
				item.SendLifetimeInfiniteVariable = types.StringNull()
			}
			if cValue := v.Get("accept-lifetime.lifetime-group-v1.local.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.AcceptLifetimeLocal = types.BoolNull()

					cv := v.Get("accept-lifetime.lifetime-group-v1.local.vipVariableName")
					item.AcceptLifetimeLocalVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.AcceptLifetimeLocal = types.BoolNull()
					item.AcceptLifetimeLocalVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("accept-lifetime.lifetime-group-v1.local.vipValue")
					item.AcceptLifetimeLocal = types.BoolValue(cv.Bool())
					item.AcceptLifetimeLocalVariable = types.StringNull()
				}
			} else {
				item.AcceptLifetimeLocal = types.BoolNull()
				item.AcceptLifetimeLocalVariable = types.StringNull()
			}
			if cValue := v.Get("accept-lifetime.lifetime-group-v1.start-epoch.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.AcceptLifetimeStartTime = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.AcceptLifetimeStartTime = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("accept-lifetime.lifetime-group-v1.start-epoch.vipValue")
					item.AcceptLifetimeStartTime = types.StringValue(cv.String())

				}
			} else {
				item.AcceptLifetimeStartTime = types.StringNull()

			}
			if cValue := v.Get("accept-lifetime.lifetime-group-v1.end-choice.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.AcceptLifetimeEndTimeFormat = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.AcceptLifetimeEndTimeFormat = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("accept-lifetime.lifetime-group-v1.end-choice.vipValue")
					item.AcceptLifetimeEndTimeFormat = types.StringValue(cv.String())

				}
			} else {
				item.AcceptLifetimeEndTimeFormat = types.StringNull()

			}
			if cValue := v.Get("accept-lifetime.lifetime-group-v1.duration.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.AcceptLifetimeDuration = types.Int64Null()

					cv := v.Get("accept-lifetime.lifetime-group-v1.duration.vipVariableName")
					item.AcceptLifetimeDurationVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.AcceptLifetimeDuration = types.Int64Null()
					item.AcceptLifetimeDurationVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("accept-lifetime.lifetime-group-v1.duration.vipValue")
					item.AcceptLifetimeDuration = types.Int64Value(cv.Int())
					item.AcceptLifetimeDurationVariable = types.StringNull()
				}
			} else {
				item.AcceptLifetimeDuration = types.Int64Null()
				item.AcceptLifetimeDurationVariable = types.StringNull()
			}
			if cValue := v.Get("accept-lifetime.lifetime-group-v1.end-epoch.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.AcceptLifetimeEndTime = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.AcceptLifetimeEndTime = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("accept-lifetime.lifetime-group-v1.end-epoch.vipValue")
					item.AcceptLifetimeEndTime = types.StringValue(cv.String())

				}
			} else {
				item.AcceptLifetimeEndTime = types.StringNull()

			}
			if cValue := v.Get("accept-lifetime.lifetime-group-v1.infinite.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.AcceptLifetimeInfinite = types.BoolNull()

					cv := v.Get("accept-lifetime.lifetime-group-v1.infinite.vipVariableName")
					item.AcceptLifetimeInfiniteVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.AcceptLifetimeInfinite = types.BoolNull()
					item.AcceptLifetimeInfiniteVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("accept-lifetime.lifetime-group-v1.infinite.vipValue")
					item.AcceptLifetimeInfinite = types.BoolValue(cv.Bool())
					item.AcceptLifetimeInfiniteVariable = types.StringNull()
				}
			} else {
				item.AcceptLifetimeInfinite = types.BoolNull()
				item.AcceptLifetimeInfiniteVariable = types.StringNull()
			}
			if cValue := v.Get("include-tcp-options.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.IncludeTcpOptions = types.BoolNull()

					cv := v.Get("include-tcp-options.vipVariableName")
					item.IncludeTcpOptionsVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.IncludeTcpOptions = types.BoolNull()
					item.IncludeTcpOptionsVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("include-tcp-options.vipValue")
					item.IncludeTcpOptions = types.BoolValue(cv.Bool())
					item.IncludeTcpOptionsVariable = types.StringNull()
				}
			} else {
				item.IncludeTcpOptions = types.BoolNull()
				item.IncludeTcpOptionsVariable = types.StringNull()
			}
			if cValue := v.Get("accept-ao-mismatch.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.AcceptAoMismatch = types.BoolNull()

					cv := v.Get("accept-ao-mismatch.vipVariableName")
					item.AcceptAoMismatchVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.AcceptAoMismatch = types.BoolNull()
					item.AcceptAoMismatchVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("accept-ao-mismatch.vipValue")
					item.AcceptAoMismatch = types.BoolValue(cv.Bool())
					item.AcceptAoMismatchVariable = types.StringNull()
				}
			} else {
				item.AcceptAoMismatch = types.BoolNull()
				item.AcceptAoMismatchVariable = types.StringNull()
			}
			data.Keys = append(data.Keys, item)
			return true
		})
	} else {
		if len(data.Keys) > 0 {
			data.Keys = []CiscoSecurityKeys{}
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *CiscoSecurity) hasChanges(ctx context.Context, state *CiscoSecurity) bool {
	hasChanges := false
	if !data.RekeyInterval.Equal(state.RekeyInterval) {
		hasChanges = true
	}
	if !data.ReplayWindow.Equal(state.ReplayWindow) {
		hasChanges = true
	}
	if !data.ExtendedArWindow.Equal(state.ExtendedArWindow) {
		hasChanges = true
	}
	if !data.AuthenticationType.Equal(state.AuthenticationType) {
		hasChanges = true
	}
	if !data.IntegrityType.Equal(state.IntegrityType) {
		hasChanges = true
	}
	if !data.PairwiseKeying.Equal(state.PairwiseKeying) {
		hasChanges = true
	}
	if len(data.Keychains) != len(state.Keychains) {
		hasChanges = true
	} else {
		for i := range data.Keychains {
			if !data.Keychains[i].Name.Equal(state.Keychains[i].Name) {
				hasChanges = true
			}
			if !data.Keychains[i].KeyId.Equal(state.Keychains[i].KeyId) {
				hasChanges = true
			}
		}
	}
	if len(data.Keys) != len(state.Keys) {
		hasChanges = true
	} else {
		for i := range data.Keys {
			if !data.Keys[i].Id.Equal(state.Keys[i].Id) {
				hasChanges = true
			}
			if !data.Keys[i].ChainName.Equal(state.Keys[i].ChainName) {
				hasChanges = true
			}
			if !data.Keys[i].SendId.Equal(state.Keys[i].SendId) {
				hasChanges = true
			}
			if !data.Keys[i].ReceiveId.Equal(state.Keys[i].ReceiveId) {
				hasChanges = true
			}
			if !data.Keys[i].CryptoAlgorithm.Equal(state.Keys[i].CryptoAlgorithm) {
				hasChanges = true
			}
			if !data.Keys[i].KeyString.Equal(state.Keys[i].KeyString) {
				hasChanges = true
			}
			if !data.Keys[i].SendLifetimeLocal.Equal(state.Keys[i].SendLifetimeLocal) {
				hasChanges = true
			}
			if !data.Keys[i].SendLifetimeStartTime.Equal(state.Keys[i].SendLifetimeStartTime) {
				hasChanges = true
			}
			if !data.Keys[i].SendLifetimeEndTimeFormat.Equal(state.Keys[i].SendLifetimeEndTimeFormat) {
				hasChanges = true
			}
			if !data.Keys[i].SendLifetimeDuration.Equal(state.Keys[i].SendLifetimeDuration) {
				hasChanges = true
			}
			if !data.Keys[i].SendLifetimeEndTime.Equal(state.Keys[i].SendLifetimeEndTime) {
				hasChanges = true
			}
			if !data.Keys[i].SendLifetimeInfinite.Equal(state.Keys[i].SendLifetimeInfinite) {
				hasChanges = true
			}
			if !data.Keys[i].AcceptLifetimeLocal.Equal(state.Keys[i].AcceptLifetimeLocal) {
				hasChanges = true
			}
			if !data.Keys[i].AcceptLifetimeStartTime.Equal(state.Keys[i].AcceptLifetimeStartTime) {
				hasChanges = true
			}
			if !data.Keys[i].AcceptLifetimeEndTimeFormat.Equal(state.Keys[i].AcceptLifetimeEndTimeFormat) {
				hasChanges = true
			}
			if !data.Keys[i].AcceptLifetimeDuration.Equal(state.Keys[i].AcceptLifetimeDuration) {
				hasChanges = true
			}
			if !data.Keys[i].AcceptLifetimeEndTime.Equal(state.Keys[i].AcceptLifetimeEndTime) {
				hasChanges = true
			}
			if !data.Keys[i].AcceptLifetimeInfinite.Equal(state.Keys[i].AcceptLifetimeInfinite) {
				hasChanges = true
			}
			if !data.Keys[i].IncludeTcpOptions.Equal(state.Keys[i].IncludeTcpOptions) {
				hasChanges = true
			}
			if !data.Keys[i].AcceptAoMismatch.Equal(state.Keys[i].AcceptAoMismatch) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges
