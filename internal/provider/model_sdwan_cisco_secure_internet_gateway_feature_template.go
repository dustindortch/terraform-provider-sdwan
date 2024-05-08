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
type CiscoSecureInternetGateway struct {
	Id                      types.String                           `tfsdk:"id"`
	Version                 types.Int64                            `tfsdk:"version"`
	TemplateType            types.String                           `tfsdk:"template_type"`
	Name                    types.String                           `tfsdk:"name"`
	Description             types.String                           `tfsdk:"description"`
	DeviceTypes             types.Set                              `tfsdk:"device_types"`
	VpnId                   types.Int64                            `tfsdk:"vpn_id"`
	Interfaces              []CiscoSecureInternetGatewayInterfaces `tfsdk:"interfaces"`
	Services                []CiscoSecureInternetGatewayServices   `tfsdk:"services"`
	TrackerSourceIp         types.String                           `tfsdk:"tracker_source_ip"`
	TrackerSourceIpVariable types.String                           `tfsdk:"tracker_source_ip_variable"`
	Trackers                []CiscoSecureInternetGatewayTrackers   `tfsdk:"trackers"`
}

type CiscoSecureInternetGatewayInterfaces struct {
	Optional                           types.Bool   `tfsdk:"optional"`
	Name                               types.String `tfsdk:"name"`
	NameVariable                       types.String `tfsdk:"name_variable"`
	AutoTunnelMode                     types.Bool   `tfsdk:"auto_tunnel_mode"`
	Shutdown                           types.Bool   `tfsdk:"shutdown"`
	Description                        types.String `tfsdk:"description"`
	DescriptionVariable                types.String `tfsdk:"description_variable"`
	IpUnnumbered                       types.Bool   `tfsdk:"ip_unnumbered"`
	Ipv4Address                        types.String `tfsdk:"ipv4_address"`
	Ipv4AddressVariable                types.String `tfsdk:"ipv4_address_variable"`
	TunnelSource                       types.String `tfsdk:"tunnel_source"`
	TunnelSourceVariable               types.String `tfsdk:"tunnel_source_variable"`
	TunnelSourceInterface              types.String `tfsdk:"tunnel_source_interface"`
	TunnelSourceInterfaceVariable      types.String `tfsdk:"tunnel_source_interface_variable"`
	TunnelRouteVia                     types.String `tfsdk:"tunnel_route_via"`
	TunnelRouteViaVariable             types.String `tfsdk:"tunnel_route_via_variable"`
	TunnelDestination                  types.String `tfsdk:"tunnel_destination"`
	TunnelDestinationVariable          types.String `tfsdk:"tunnel_destination_variable"`
	Application                        types.String `tfsdk:"application"`
	SigProvider                        types.String `tfsdk:"sig_provider"`
	TunnelDcPreference                 types.String `tfsdk:"tunnel_dc_preference"`
	TcpMss                             types.Int64  `tfsdk:"tcp_mss"`
	TcpMssVariable                     types.String `tfsdk:"tcp_mss_variable"`
	Mtu                                types.Int64  `tfsdk:"mtu"`
	MtuVariable                        types.String `tfsdk:"mtu_variable"`
	DeadPeerDetectionInterval          types.Int64  `tfsdk:"dead_peer_detection_interval"`
	DeadPeerDetectionIntervalVariable  types.String `tfsdk:"dead_peer_detection_interval_variable"`
	DeadPeerDetectionRetries           types.Int64  `tfsdk:"dead_peer_detection_retries"`
	DeadPeerDetectionRetriesVariable   types.String `tfsdk:"dead_peer_detection_retries_variable"`
	IkeVersion                         types.Int64  `tfsdk:"ike_version"`
	IkeVersionVariable                 types.String `tfsdk:"ike_version_variable"`
	IkePreSharedKey                    types.String `tfsdk:"ike_pre_shared_key"`
	IkePreSharedKeyVariable            types.String `tfsdk:"ike_pre_shared_key_variable"`
	IkeRekeyInterval                   types.Int64  `tfsdk:"ike_rekey_interval"`
	IkeRekeyIntervalVariable           types.String `tfsdk:"ike_rekey_interval_variable"`
	IkeCiphersuite                     types.String `tfsdk:"ike_ciphersuite"`
	IkeCiphersuiteVariable             types.String `tfsdk:"ike_ciphersuite_variable"`
	IkeGroup                           types.String `tfsdk:"ike_group"`
	IkeGroupVariable                   types.String `tfsdk:"ike_group_variable"`
	IkePreSharedKeyDynamic             types.Bool   `tfsdk:"ike_pre_shared_key_dynamic"`
	IkePreSharedKeyLocalId             types.String `tfsdk:"ike_pre_shared_key_local_id"`
	IkePreSharedKeyLocalIdVariable     types.String `tfsdk:"ike_pre_shared_key_local_id_variable"`
	IkePreSharedKeyRemoteId            types.String `tfsdk:"ike_pre_shared_key_remote_id"`
	IkePreSharedKeyRemoteIdVariable    types.String `tfsdk:"ike_pre_shared_key_remote_id_variable"`
	IpsecRekeyInterval                 types.Int64  `tfsdk:"ipsec_rekey_interval"`
	IpsecRekeyIntervalVariable         types.String `tfsdk:"ipsec_rekey_interval_variable"`
	IpsecReplayWindow                  types.Int64  `tfsdk:"ipsec_replay_window"`
	IpsecReplayWindowVariable          types.String `tfsdk:"ipsec_replay_window_variable"`
	IpsecCiphersuite                   types.String `tfsdk:"ipsec_ciphersuite"`
	IpsecCiphersuiteVariable           types.String `tfsdk:"ipsec_ciphersuite_variable"`
	IpsecPerfectForwardSecrecy         types.String `tfsdk:"ipsec_perfect_forward_secrecy"`
	IpsecPerfectForwardSecrecyVariable types.String `tfsdk:"ipsec_perfect_forward_secrecy_variable"`
	Tracker                            types.String `tfsdk:"tracker"`
	TrackEnable                        types.Bool   `tfsdk:"track_enable"`
	TunnelPublicIp                     types.String `tfsdk:"tunnel_public_ip"`
	TunnelPublicIpVariable             types.String `tfsdk:"tunnel_public_ip_variable"`
}

type CiscoSecureInternetGatewayServices struct {
	Optional                                  types.Bool                                         `tfsdk:"optional"`
	ServiceType                               types.String                                       `tfsdk:"service_type"`
	InterfacePairs                            []CiscoSecureInternetGatewayServicesInterfacePairs `tfsdk:"interface_pairs"`
	ZscalerAuthenticationRequired             types.Bool                                         `tfsdk:"zscaler_authentication_required"`
	ZscalerXffForward                         types.Bool                                         `tfsdk:"zscaler_xff_forward"`
	ZscalerFirewallEnabled                    types.Bool                                         `tfsdk:"zscaler_firewall_enabled"`
	ZscalerIpsControlEnabled                  types.Bool                                         `tfsdk:"zscaler_ips_control_enabled"`
	ZscalerCautionEnabled                     types.Bool                                         `tfsdk:"zscaler_caution_enabled"`
	ZscalerPrimaryDataCenter                  types.String                                       `tfsdk:"zscaler_primary_data_center"`
	ZscalerPrimaryDataCenterVariable          types.String                                       `tfsdk:"zscaler_primary_data_center_variable"`
	ZscalerSecondaryDataCenter                types.String                                       `tfsdk:"zscaler_secondary_data_center"`
	ZscalerSecondaryDataCenterVariable        types.String                                       `tfsdk:"zscaler_secondary_data_center_variable"`
	ZscalerSurrogateIp                        types.Bool                                         `tfsdk:"zscaler_surrogate_ip"`
	ZscalerSurrogateIdleTime                  types.Int64                                        `tfsdk:"zscaler_surrogate_idle_time"`
	ZscalerSurrogateDisplayTimeUnit           types.String                                       `tfsdk:"zscaler_surrogate_display_time_unit"`
	ZscalerSurrogateIpEnforceForKnownBrowsers types.Bool                                         `tfsdk:"zscaler_surrogate_ip_enforce_for_known_browsers"`
	ZscalerSurrogateRefreshTime               types.Int64                                        `tfsdk:"zscaler_surrogate_refresh_time"`
	ZscalerSurrogateRefreshTimeUnit           types.String                                       `tfsdk:"zscaler_surrogate_refresh_time_unit"`
	ZscalerAupEnabled                         types.Bool                                         `tfsdk:"zscaler_aup_enabled"`
	ZscalerAupBlockInternetUntilAccepted      types.Bool                                         `tfsdk:"zscaler_aup_block_internet_until_accepted"`
	ZscalerAupForceSslInspection              types.Bool                                         `tfsdk:"zscaler_aup_force_ssl_inspection"`
	ZscalerAupTimeout                         types.Int64                                        `tfsdk:"zscaler_aup_timeout"`
	ZscalerLocationName                       types.String                                       `tfsdk:"zscaler_location_name"`
	ZscalerLocationNameVariable               types.String                                       `tfsdk:"zscaler_location_name_variable"`
	UmbrellaPrimaryDataCenter                 types.String                                       `tfsdk:"umbrella_primary_data_center"`
	UmbrellaPrimaryDataCenterVariable         types.String                                       `tfsdk:"umbrella_primary_data_center_variable"`
	UmbrellaSecondaryDataCenter               types.String                                       `tfsdk:"umbrella_secondary_data_center"`
	UmbrellaSecondaryDataCenterVariable       types.String                                       `tfsdk:"umbrella_secondary_data_center_variable"`
}

type CiscoSecureInternetGatewayTrackers struct {
	Optional               types.Bool   `tfsdk:"optional"`
	Name                   types.String `tfsdk:"name"`
	NameVariable           types.String `tfsdk:"name_variable"`
	EndpointApiUrl         types.String `tfsdk:"endpoint_api_url"`
	EndpointApiUrlVariable types.String `tfsdk:"endpoint_api_url_variable"`
	Threshold              types.Int64  `tfsdk:"threshold"`
	ThresholdVariable      types.String `tfsdk:"threshold_variable"`
	Interval               types.Int64  `tfsdk:"interval"`
	IntervalVariable       types.String `tfsdk:"interval_variable"`
	Multiplier             types.Int64  `tfsdk:"multiplier"`
	MultiplierVariable     types.String `tfsdk:"multiplier_variable"`
	TrackerType            types.String `tfsdk:"tracker_type"`
}

type CiscoSecureInternetGatewayServicesInterfacePairs struct {
	Optional              types.Bool   `tfsdk:"optional"`
	ActiveInterface       types.String `tfsdk:"active_interface"`
	ActiveInterfaceWeight types.Int64  `tfsdk:"active_interface_weight"`
	BackupInterface       types.String `tfsdk:"backup_interface"`
	BackupInterfaceWeight types.Int64  `tfsdk:"backup_interface_weight"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data CiscoSecureInternetGateway) getModel() string {
	return "cisco_secure_internet_gateway"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data CiscoSecureInternetGateway) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "cisco_secure_internet_gateway")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."
	if data.VpnId.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"vpn-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"vpn-id."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"vpn-id."+"vipValue", data.VpnId.ValueInt64())
	}
	if len(data.Interfaces) > 0 {
		body, _ = sjson.Set(body, path+"interface."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"interface."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"interface."+"vipPrimaryKey", []string{"if-name"})
		body, _ = sjson.Set(body, path+"interface."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"interface."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"interface."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"interface."+"vipPrimaryKey", []string{"if-name"})
		body, _ = sjson.Set(body, path+"interface."+"vipValue", []interface{}{})
	}
	for _, item := range data.Interfaces {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "if-name")

		if !item.NameVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipVariableName", item.NameVariable.ValueString())
		} else if item.Name.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipValue", item.Name.ValueString())
		}
		itemAttributes = append(itemAttributes, "auto")
		if item.AutoTunnelMode.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "auto."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "auto."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "auto."+"vipValue", strconv.FormatBool(item.AutoTunnelMode.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "shutdown")
		if item.Shutdown.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "shutdown."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "shutdown."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "shutdown."+"vipValue", strconv.FormatBool(item.Shutdown.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "description")

		if !item.DescriptionVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "description."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "description."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "description."+"vipVariableName", item.DescriptionVariable.ValueString())
		} else if item.Description.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "description."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "description."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "description."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "description."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "description."+"vipValue", item.Description.ValueString())
		}
		itemAttributes = append(itemAttributes, "ip")
		if item.IpUnnumbered.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "ip.unnumbered."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "ip.unnumbered."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "ip.unnumbered."+"vipValue", strconv.FormatBool(item.IpUnnumbered.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "ip")

		if !item.Ipv4AddressVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ip.address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ip.address."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "ip.address."+"vipVariableName", item.Ipv4AddressVariable.ValueString())
		} else if item.Ipv4Address.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "ip.address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ip.address."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "ip.address."+"vipValue", item.Ipv4Address.ValueString())
		}
		itemAttributes = append(itemAttributes, "tunnel-source")

		if !item.TunnelSourceVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tunnel-source."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tunnel-source."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "tunnel-source."+"vipVariableName", item.TunnelSourceVariable.ValueString())
		} else if item.TunnelSource.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "tunnel-source."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tunnel-source."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "tunnel-source."+"vipValue", item.TunnelSource.ValueString())
		}
		itemAttributes = append(itemAttributes, "tunnel-source-interface")

		if !item.TunnelSourceInterfaceVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tunnel-source-interface."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tunnel-source-interface."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "tunnel-source-interface."+"vipVariableName", item.TunnelSourceInterfaceVariable.ValueString())
		} else if item.TunnelSourceInterface.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "tunnel-source-interface."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tunnel-source-interface."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "tunnel-source-interface."+"vipValue", item.TunnelSourceInterface.ValueString())
		}
		itemAttributes = append(itemAttributes, "tunnel-route-via")

		if !item.TunnelRouteViaVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tunnel-route-via."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tunnel-route-via."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "tunnel-route-via."+"vipVariableName", item.TunnelRouteViaVariable.ValueString())
		} else if item.TunnelRouteVia.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "tunnel-route-via."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tunnel-route-via."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "tunnel-route-via."+"vipValue", item.TunnelRouteVia.ValueString())
		}
		itemAttributes = append(itemAttributes, "tunnel-destination")

		if !item.TunnelDestinationVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tunnel-destination."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tunnel-destination."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "tunnel-destination."+"vipVariableName", item.TunnelDestinationVariable.ValueString())
		} else if item.TunnelDestination.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "tunnel-destination."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tunnel-destination."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "tunnel-destination."+"vipValue", item.TunnelDestination.ValueString())
		}
		itemAttributes = append(itemAttributes, "application")
		if item.Application.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "application."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "application."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "application."+"vipValue", item.Application.ValueString())
		}
		itemAttributes = append(itemAttributes, "tunnel-set")
		if item.SigProvider.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "tunnel-set."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tunnel-set."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "tunnel-set."+"vipValue", item.SigProvider.ValueString())
		}
		itemAttributes = append(itemAttributes, "tunnel-dc-preference")
		if item.TunnelDcPreference.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "tunnel-dc-preference."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tunnel-dc-preference."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "tunnel-dc-preference."+"vipValue", item.TunnelDcPreference.ValueString())
		}
		itemAttributes = append(itemAttributes, "tcp-mss-adjust")

		if !item.TcpMssVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tcp-mss-adjust."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tcp-mss-adjust."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "tcp-mss-adjust."+"vipVariableName", item.TcpMssVariable.ValueString())
		} else if item.TcpMss.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tcp-mss-adjust."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tcp-mss-adjust."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "tcp-mss-adjust."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tcp-mss-adjust."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "tcp-mss-adjust."+"vipValue", item.TcpMss.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "mtu")

		if !item.MtuVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "mtu."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "mtu."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "mtu."+"vipVariableName", item.MtuVariable.ValueString())
		} else if item.Mtu.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "mtu."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "mtu."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "mtu."+"vipValue", item.Mtu.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "dead-peer-detection")

		if !item.DeadPeerDetectionIntervalVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dead-peer-detection.dpd-interval."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dead-peer-detection.dpd-interval."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "dead-peer-detection.dpd-interval."+"vipVariableName", item.DeadPeerDetectionIntervalVariable.ValueString())
		} else if item.DeadPeerDetectionInterval.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dead-peer-detection.dpd-interval."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dead-peer-detection.dpd-interval."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "dead-peer-detection.dpd-interval."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dead-peer-detection.dpd-interval."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "dead-peer-detection.dpd-interval."+"vipValue", item.DeadPeerDetectionInterval.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "dead-peer-detection")

		if !item.DeadPeerDetectionRetriesVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dead-peer-detection.dpd-retries."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dead-peer-detection.dpd-retries."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "dead-peer-detection.dpd-retries."+"vipVariableName", item.DeadPeerDetectionRetriesVariable.ValueString())
		} else if item.DeadPeerDetectionRetries.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dead-peer-detection.dpd-retries."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dead-peer-detection.dpd-retries."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "dead-peer-detection.dpd-retries."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dead-peer-detection.dpd-retries."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "dead-peer-detection.dpd-retries."+"vipValue", item.DeadPeerDetectionRetries.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "ike")

		if !item.IkeVersionVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ike.ike-version."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ike.ike-version."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "ike.ike-version."+"vipVariableName", item.IkeVersionVariable.ValueString())
		} else if item.IkeVersion.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "ike.ike-version."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ike.ike-version."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "ike.ike-version."+"vipValue", item.IkeVersion.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "ike")

		if !item.IkePreSharedKeyVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ike.authentication-type.pre-shared-key.pre-shared-secret."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ike.authentication-type.pre-shared-key.pre-shared-secret."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "ike.authentication-type.pre-shared-key.pre-shared-secret."+"vipVariableName", item.IkePreSharedKeyVariable.ValueString())
		} else if item.IkePreSharedKey.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "ike.authentication-type.pre-shared-key.pre-shared-secret."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ike.authentication-type.pre-shared-key.pre-shared-secret."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "ike.authentication-type.pre-shared-key.pre-shared-secret."+"vipValue", item.IkePreSharedKey.ValueString())
		}
		itemAttributes = append(itemAttributes, "ike")

		if !item.IkeRekeyIntervalVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ike.ike-rekey-interval."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ike.ike-rekey-interval."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "ike.ike-rekey-interval."+"vipVariableName", item.IkeRekeyIntervalVariable.ValueString())
		} else if item.IkeRekeyInterval.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "ike.ike-rekey-interval."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ike.ike-rekey-interval."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "ike.ike-rekey-interval."+"vipValue", item.IkeRekeyInterval.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "ike")

		if !item.IkeCiphersuiteVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ike.ike-ciphersuite."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ike.ike-ciphersuite."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "ike.ike-ciphersuite."+"vipVariableName", item.IkeCiphersuiteVariable.ValueString())
		} else if item.IkeCiphersuite.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "ike.ike-ciphersuite."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ike.ike-ciphersuite."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "ike.ike-ciphersuite."+"vipValue", item.IkeCiphersuite.ValueString())
		}
		itemAttributes = append(itemAttributes, "ike")

		if !item.IkeGroupVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ike.ike-group."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ike.ike-group."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "ike.ike-group."+"vipVariableName", item.IkeGroupVariable.ValueString())
		} else if item.IkeGroup.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "ike.ike-group."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ike.ike-group."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "ike.ike-group."+"vipValue", item.IkeGroup.ValueString())
		}
		itemAttributes = append(itemAttributes, "ike")
		if item.IkePreSharedKeyDynamic.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "ike.authentication-type.pre-shared-key-dynamic."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "ike.authentication-type.pre-shared-key-dynamic."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "ike.authentication-type.pre-shared-key-dynamic."+"vipValue", strconv.FormatBool(item.IkePreSharedKeyDynamic.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "ike")

		if !item.IkePreSharedKeyLocalIdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ike.authentication-type.pre-shared-key.ike-local-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ike.authentication-type.pre-shared-key.ike-local-id."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "ike.authentication-type.pre-shared-key.ike-local-id."+"vipVariableName", item.IkePreSharedKeyLocalIdVariable.ValueString())
		} else if item.IkePreSharedKeyLocalId.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "ike.authentication-type.pre-shared-key.ike-local-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ike.authentication-type.pre-shared-key.ike-local-id."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "ike.authentication-type.pre-shared-key.ike-local-id."+"vipValue", item.IkePreSharedKeyLocalId.ValueString())
		}
		itemAttributes = append(itemAttributes, "ike")

		if !item.IkePreSharedKeyRemoteIdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ike.authentication-type.pre-shared-key.ike-remote-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ike.authentication-type.pre-shared-key.ike-remote-id."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "ike.authentication-type.pre-shared-key.ike-remote-id."+"vipVariableName", item.IkePreSharedKeyRemoteIdVariable.ValueString())
		} else if item.IkePreSharedKeyRemoteId.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "ike.authentication-type.pre-shared-key.ike-remote-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ike.authentication-type.pre-shared-key.ike-remote-id."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "ike.authentication-type.pre-shared-key.ike-remote-id."+"vipValue", item.IkePreSharedKeyRemoteId.ValueString())
		}
		itemAttributes = append(itemAttributes, "ipsec")

		if !item.IpsecRekeyIntervalVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ipsec.ipsec-rekey-interval."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ipsec.ipsec-rekey-interval."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "ipsec.ipsec-rekey-interval."+"vipVariableName", item.IpsecRekeyIntervalVariable.ValueString())
		} else if item.IpsecRekeyInterval.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "ipsec.ipsec-rekey-interval."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ipsec.ipsec-rekey-interval."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "ipsec.ipsec-rekey-interval."+"vipValue", item.IpsecRekeyInterval.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "ipsec")

		if !item.IpsecReplayWindowVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ipsec.ipsec-replay-window."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ipsec.ipsec-replay-window."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "ipsec.ipsec-replay-window."+"vipVariableName", item.IpsecReplayWindowVariable.ValueString())
		} else if item.IpsecReplayWindow.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "ipsec.ipsec-replay-window."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ipsec.ipsec-replay-window."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "ipsec.ipsec-replay-window."+"vipValue", item.IpsecReplayWindow.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "ipsec")

		if !item.IpsecCiphersuiteVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ipsec.ipsec-ciphersuite."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ipsec.ipsec-ciphersuite."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "ipsec.ipsec-ciphersuite."+"vipVariableName", item.IpsecCiphersuiteVariable.ValueString())
		} else if item.IpsecCiphersuite.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "ipsec.ipsec-ciphersuite."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ipsec.ipsec-ciphersuite."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "ipsec.ipsec-ciphersuite."+"vipValue", item.IpsecCiphersuite.ValueString())
		}
		itemAttributes = append(itemAttributes, "ipsec")

		if !item.IpsecPerfectForwardSecrecyVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ipsec.perfect-forward-secrecy."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ipsec.perfect-forward-secrecy."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "ipsec.perfect-forward-secrecy."+"vipVariableName", item.IpsecPerfectForwardSecrecyVariable.ValueString())
		} else if item.IpsecPerfectForwardSecrecy.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "ipsec.perfect-forward-secrecy."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ipsec.perfect-forward-secrecy."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "ipsec.perfect-forward-secrecy."+"vipValue", item.IpsecPerfectForwardSecrecy.ValueString())
		}
		itemAttributes = append(itemAttributes, "tracker")
		if item.Tracker.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tracker."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tracker."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "tracker."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tracker."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "tracker."+"vipValue", item.Tracker.ValueString())
		}
		itemAttributes = append(itemAttributes, "track-enable")
		if item.TrackEnable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "track-enable."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "track-enable."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "track-enable."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "track-enable."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "track-enable."+"vipValue", strconv.FormatBool(item.TrackEnable.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "tunnel-public-ip")

		if !item.TunnelPublicIpVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tunnel-public-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tunnel-public-ip."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "tunnel-public-ip."+"vipVariableName", item.TunnelPublicIpVariable.ValueString())
		} else if item.TunnelPublicIp.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tunnel-public-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tunnel-public-ip."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "tunnel-public-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tunnel-public-ip."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "tunnel-public-ip."+"vipValue", item.TunnelPublicIp.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"interface."+"vipValue.-1", itemBody)
	}
	if len(data.Services) > 0 {
		body, _ = sjson.Set(body, path+"service."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"service."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"service."+"vipPrimaryKey", []string{"svc-type"})
		body, _ = sjson.Set(body, path+"service."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"service."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"service."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"service."+"vipPrimaryKey", []string{"svc-type"})
		body, _ = sjson.Set(body, path+"service."+"vipValue", []interface{}{})
	}
	for _, item := range data.Services {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "svc-type")
		if item.ServiceType.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "svc-type."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "svc-type."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "svc-type."+"vipValue", item.ServiceType.ValueString())
		}
		itemAttributes = append(itemAttributes, "ha-pairs")
		if len(item.InterfacePairs) > 0 {
			itemBody, _ = sjson.Set(itemBody, "ha-pairs.interface-pair."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "ha-pairs.interface-pair."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "ha-pairs.interface-pair."+"vipPrimaryKey", []string{"active-interface", "backup-interface"})
			itemBody, _ = sjson.Set(itemBody, "ha-pairs.interface-pair."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "ha-pairs.interface-pair."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "ha-pairs.interface-pair."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "ha-pairs.interface-pair."+"vipPrimaryKey", []string{"active-interface", "backup-interface"})
			itemBody, _ = sjson.Set(itemBody, "ha-pairs.interface-pair."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.InterfacePairs {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "active-interface")
			if childItem.ActiveInterface.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "active-interface."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "active-interface."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "active-interface."+"vipValue", childItem.ActiveInterface.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "active-interface-weight")
			if childItem.ActiveInterfaceWeight.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "active-interface-weight."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "active-interface-weight."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "active-interface-weight."+"vipValue", childItem.ActiveInterfaceWeight.ValueInt64())
			}
			itemChildAttributes = append(itemChildAttributes, "backup-interface")
			if childItem.BackupInterface.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "backup-interface."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "backup-interface."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "backup-interface."+"vipValue", childItem.BackupInterface.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "backup-interface-weight")
			if childItem.BackupInterfaceWeight.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "backup-interface-weight."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "backup-interface-weight."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "backup-interface-weight."+"vipValue", childItem.BackupInterfaceWeight.ValueInt64())
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "ha-pairs.interface-pair."+"vipValue.-1", itemChildBody)
		}
		itemAttributes = append(itemAttributes, "zscaler-location-settings")
		if item.ZscalerAuthenticationRequired.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.auth-required."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.auth-required."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.auth-required."+"vipValue", strconv.FormatBool(item.ZscalerAuthenticationRequired.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "zscaler-location-settings")
		if item.ZscalerXffForward.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.xff-forward-enabled."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.xff-forward-enabled."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.xff-forward-enabled."+"vipValue", strconv.FormatBool(item.ZscalerXffForward.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "zscaler-location-settings")
		if item.ZscalerFirewallEnabled.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.ofw-enabled."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.ofw-enabled."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.ofw-enabled."+"vipValue", strconv.FormatBool(item.ZscalerFirewallEnabled.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "zscaler-location-settings")
		if item.ZscalerIpsControlEnabled.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.ips-control."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.ips-control."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.ips-control."+"vipValue", strconv.FormatBool(item.ZscalerIpsControlEnabled.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "zscaler-location-settings")
		if item.ZscalerCautionEnabled.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.caution-enabled."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.caution-enabled."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.caution-enabled."+"vipValue", strconv.FormatBool(item.ZscalerCautionEnabled.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "zscaler-location-settings")

		if !item.ZscalerPrimaryDataCenterVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.datacenters.primary-data-center."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.datacenters.primary-data-center."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.datacenters.primary-data-center."+"vipVariableName", item.ZscalerPrimaryDataCenterVariable.ValueString())
		} else if item.ZscalerPrimaryDataCenter.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.datacenters.primary-data-center."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.datacenters.primary-data-center."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.datacenters.primary-data-center."+"vipValue", item.ZscalerPrimaryDataCenter.ValueString())
		}
		itemAttributes = append(itemAttributes, "zscaler-location-settings")

		if !item.ZscalerSecondaryDataCenterVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.datacenters.secondary-data-center."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.datacenters.secondary-data-center."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.datacenters.secondary-data-center."+"vipVariableName", item.ZscalerSecondaryDataCenterVariable.ValueString())
		} else if item.ZscalerSecondaryDataCenter.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.datacenters.secondary-data-center."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.datacenters.secondary-data-center."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.datacenters.secondary-data-center."+"vipValue", item.ZscalerSecondaryDataCenter.ValueString())
		}
		itemAttributes = append(itemAttributes, "zscaler-location-settings")
		if item.ZscalerSurrogateIp.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.surrogate.ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.surrogate.ip."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.surrogate.ip."+"vipValue", strconv.FormatBool(item.ZscalerSurrogateIp.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "zscaler-location-settings")
		if item.ZscalerSurrogateIdleTime.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.surrogate.idle-time."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.surrogate.idle-time."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.surrogate.idle-time."+"vipValue", item.ZscalerSurrogateIdleTime.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "zscaler-location-settings")
		if item.ZscalerSurrogateDisplayTimeUnit.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.surrogate.display-time-unit."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.surrogate.display-time-unit."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.surrogate.display-time-unit."+"vipValue", item.ZscalerSurrogateDisplayTimeUnit.ValueString())
		}
		itemAttributes = append(itemAttributes, "zscaler-location-settings")
		if item.ZscalerSurrogateIpEnforceForKnownBrowsers.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.surrogate.ip-enforced-for-known-browsers."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.surrogate.ip-enforced-for-known-browsers."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.surrogate.ip-enforced-for-known-browsers."+"vipValue", strconv.FormatBool(item.ZscalerSurrogateIpEnforceForKnownBrowsers.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "zscaler-location-settings")
		if item.ZscalerSurrogateRefreshTime.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.surrogate.refresh-time."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.surrogate.refresh-time."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.surrogate.refresh-time."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.surrogate.refresh-time."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.surrogate.refresh-time."+"vipValue", item.ZscalerSurrogateRefreshTime.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "zscaler-location-settings")
		if item.ZscalerSurrogateRefreshTimeUnit.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.surrogate.refresh-time-unit."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.surrogate.refresh-time-unit."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.surrogate.refresh-time-unit."+"vipValue", item.ZscalerSurrogateRefreshTimeUnit.ValueString())
		}
		itemAttributes = append(itemAttributes, "zscaler-location-settings")
		if item.ZscalerAupEnabled.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.aup.enabled."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.aup.enabled."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.aup.enabled."+"vipValue", strconv.FormatBool(item.ZscalerAupEnabled.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "zscaler-location-settings")
		if item.ZscalerAupBlockInternetUntilAccepted.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.aup.block-internet-until-accepted."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.aup.block-internet-until-accepted."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.aup.block-internet-until-accepted."+"vipValue", strconv.FormatBool(item.ZscalerAupBlockInternetUntilAccepted.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "zscaler-location-settings")
		if item.ZscalerAupForceSslInspection.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.aup.force-ssl-inspection."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.aup.force-ssl-inspection."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.aup.force-ssl-inspection."+"vipValue", strconv.FormatBool(item.ZscalerAupForceSslInspection.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "zscaler-location-settings")
		if item.ZscalerAupTimeout.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.aup.timeout."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.aup.timeout."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.aup.timeout."+"vipValue", item.ZscalerAupTimeout.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "zscaler-location-settings")

		if !item.ZscalerLocationNameVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.location-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.location-name."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.location-name."+"vipVariableName", item.ZscalerLocationNameVariable.ValueString())
		} else if item.ZscalerLocationName.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.location-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.location-name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "zscaler-location-settings.location-name."+"vipValue", item.ZscalerLocationName.ValueString())
		}
		itemAttributes = append(itemAttributes, "umbrella-data-center")

		if !item.UmbrellaPrimaryDataCenterVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "umbrella-data-center.data-center-primary."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "umbrella-data-center.data-center-primary."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "umbrella-data-center.data-center-primary."+"vipVariableName", item.UmbrellaPrimaryDataCenterVariable.ValueString())
		} else if item.UmbrellaPrimaryDataCenter.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "umbrella-data-center.data-center-primary."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "umbrella-data-center.data-center-primary."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "umbrella-data-center.data-center-primary."+"vipValue", item.UmbrellaPrimaryDataCenter.ValueString())
		}
		itemAttributes = append(itemAttributes, "umbrella-data-center")

		if !item.UmbrellaSecondaryDataCenterVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "umbrella-data-center.data-center-secondary."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "umbrella-data-center.data-center-secondary."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "umbrella-data-center.data-center-secondary."+"vipVariableName", item.UmbrellaSecondaryDataCenterVariable.ValueString())
		} else if item.UmbrellaSecondaryDataCenter.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "umbrella-data-center.data-center-secondary."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "umbrella-data-center.data-center-secondary."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "umbrella-data-center.data-center-secondary."+"vipValue", item.UmbrellaSecondaryDataCenter.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"service."+"vipValue.-1", itemBody)
	}

	if !data.TrackerSourceIpVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tracker-src-ip."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tracker-src-ip."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tracker-src-ip."+"vipVariableName", data.TrackerSourceIpVariable.ValueString())
	} else if data.TrackerSourceIp.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tracker-src-ip."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tracker-src-ip."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tracker-src-ip."+"vipValue", data.TrackerSourceIp.ValueString())
	}
	if len(data.Trackers) > 0 {
		body, _ = sjson.Set(body, path+"tracker."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"tracker."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tracker."+"vipPrimaryKey", []string{"tracker-type", "name"})
		body, _ = sjson.Set(body, path+"tracker."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"tracker."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"tracker."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"tracker."+"vipPrimaryKey", []string{"tracker-type", "name"})
		body, _ = sjson.Set(body, path+"tracker."+"vipValue", []interface{}{})
	}
	for _, item := range data.Trackers {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "name")

		if !item.NameVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipVariableName", item.NameVariable.ValueString())
		} else if item.Name.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipValue", item.Name.ValueString())
		}
		itemAttributes = append(itemAttributes, "endpoint-api-url")

		if !item.EndpointApiUrlVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "endpoint-api-url."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "endpoint-api-url."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "endpoint-api-url."+"vipVariableName", item.EndpointApiUrlVariable.ValueString())
		} else if item.EndpointApiUrl.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "endpoint-api-url."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "endpoint-api-url."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "endpoint-api-url."+"vipValue", item.EndpointApiUrl.ValueString())
		}
		itemAttributes = append(itemAttributes, "threshold")

		if !item.ThresholdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "threshold."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "threshold."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "threshold."+"vipVariableName", item.ThresholdVariable.ValueString())
		} else if item.Threshold.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "threshold."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "threshold."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "threshold."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "threshold."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "threshold."+"vipValue", item.Threshold.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "interval")

		if !item.IntervalVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "interval."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "interval."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "interval."+"vipVariableName", item.IntervalVariable.ValueString())
		} else if item.Interval.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "interval."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "interval."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "interval."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "interval."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "interval."+"vipValue", item.Interval.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "multiplier")

		if !item.MultiplierVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "multiplier."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "multiplier."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "multiplier."+"vipVariableName", item.MultiplierVariable.ValueString())
		} else if item.Multiplier.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "multiplier."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "multiplier."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "multiplier."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "multiplier."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "multiplier."+"vipValue", item.Multiplier.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "tracker-type")
		if item.TrackerType.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "tracker-type."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tracker-type."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "tracker-type."+"vipValue", item.TrackerType.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"tracker."+"vipValue.-1", itemBody)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *CiscoSecureInternetGateway) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(path + "vpn-id.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.VpnId = types.Int64Null()

		} else if value.String() == "ignore" {
			data.VpnId = types.Int64Null()

		} else if value.String() == "constant" {
			v := res.Get(path + "vpn-id.vipValue")
			data.VpnId = types.Int64Value(v.Int())

		}
	} else {
		data.VpnId = types.Int64Null()

	}
	if value := res.Get(path + "interface.vipValue"); len(value.Array()) > 0 {
		data.Interfaces = make([]CiscoSecureInternetGatewayInterfaces, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoSecureInternetGatewayInterfaces{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("if-name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Name = types.StringNull()

					cv := v.Get("if-name.vipVariableName")
					item.NameVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Name = types.StringNull()
					item.NameVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("if-name.vipValue")
					item.Name = types.StringValue(cv.String())
					item.NameVariable = types.StringNull()
				}
			} else {
				item.Name = types.StringNull()
				item.NameVariable = types.StringNull()
			}
			if cValue := v.Get("auto.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.AutoTunnelMode = types.BoolNull()

				} else if cValue.String() == "ignore" {
					item.AutoTunnelMode = types.BoolNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("auto.vipValue")
					item.AutoTunnelMode = types.BoolValue(cv.Bool())

				}
			} else {
				item.AutoTunnelMode = types.BoolNull()

			}
			if cValue := v.Get("shutdown.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Shutdown = types.BoolNull()

				} else if cValue.String() == "ignore" {
					item.Shutdown = types.BoolNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("shutdown.vipValue")
					item.Shutdown = types.BoolValue(cv.Bool())

				}
			} else {
				item.Shutdown = types.BoolNull()

			}
			if cValue := v.Get("description.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Description = types.StringNull()

					cv := v.Get("description.vipVariableName")
					item.DescriptionVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Description = types.StringNull()
					item.DescriptionVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("description.vipValue")
					item.Description = types.StringValue(cv.String())
					item.DescriptionVariable = types.StringNull()
				}
			} else {
				item.Description = types.StringNull()
				item.DescriptionVariable = types.StringNull()
			}
			if cValue := v.Get("ip.unnumbered.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.IpUnnumbered = types.BoolNull()

				} else if cValue.String() == "ignore" {
					item.IpUnnumbered = types.BoolNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("ip.unnumbered.vipValue")
					item.IpUnnumbered = types.BoolValue(cv.Bool())

				}
			} else {
				item.IpUnnumbered = types.BoolNull()

			}
			if cValue := v.Get("ip.address.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Ipv4Address = types.StringNull()

					cv := v.Get("ip.address.vipVariableName")
					item.Ipv4AddressVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Ipv4Address = types.StringNull()
					item.Ipv4AddressVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("ip.address.vipValue")
					item.Ipv4Address = types.StringValue(cv.String())
					item.Ipv4AddressVariable = types.StringNull()
				}
			} else {
				item.Ipv4Address = types.StringNull()
				item.Ipv4AddressVariable = types.StringNull()
			}
			if cValue := v.Get("tunnel-source.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TunnelSource = types.StringNull()

					cv := v.Get("tunnel-source.vipVariableName")
					item.TunnelSourceVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TunnelSource = types.StringNull()
					item.TunnelSourceVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("tunnel-source.vipValue")
					item.TunnelSource = types.StringValue(cv.String())
					item.TunnelSourceVariable = types.StringNull()
				}
			} else {
				item.TunnelSource = types.StringNull()
				item.TunnelSourceVariable = types.StringNull()
			}
			if cValue := v.Get("tunnel-source-interface.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TunnelSourceInterface = types.StringNull()

					cv := v.Get("tunnel-source-interface.vipVariableName")
					item.TunnelSourceInterfaceVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TunnelSourceInterface = types.StringNull()
					item.TunnelSourceInterfaceVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("tunnel-source-interface.vipValue")
					item.TunnelSourceInterface = types.StringValue(cv.String())
					item.TunnelSourceInterfaceVariable = types.StringNull()
				}
			} else {
				item.TunnelSourceInterface = types.StringNull()
				item.TunnelSourceInterfaceVariable = types.StringNull()
			}
			if cValue := v.Get("tunnel-route-via.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TunnelRouteVia = types.StringNull()

					cv := v.Get("tunnel-route-via.vipVariableName")
					item.TunnelRouteViaVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TunnelRouteVia = types.StringNull()
					item.TunnelRouteViaVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("tunnel-route-via.vipValue")
					item.TunnelRouteVia = types.StringValue(cv.String())
					item.TunnelRouteViaVariable = types.StringNull()
				}
			} else {
				item.TunnelRouteVia = types.StringNull()
				item.TunnelRouteViaVariable = types.StringNull()
			}
			if cValue := v.Get("tunnel-destination.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TunnelDestination = types.StringNull()

					cv := v.Get("tunnel-destination.vipVariableName")
					item.TunnelDestinationVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TunnelDestination = types.StringNull()
					item.TunnelDestinationVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("tunnel-destination.vipValue")
					item.TunnelDestination = types.StringValue(cv.String())
					item.TunnelDestinationVariable = types.StringNull()
				}
			} else {
				item.TunnelDestination = types.StringNull()
				item.TunnelDestinationVariable = types.StringNull()
			}
			if cValue := v.Get("application.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Application = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.Application = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("application.vipValue")
					item.Application = types.StringValue(cv.String())

				}
			} else {
				item.Application = types.StringNull()

			}
			if cValue := v.Get("tunnel-set.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SigProvider = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.SigProvider = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("tunnel-set.vipValue")
					item.SigProvider = types.StringValue(cv.String())

				}
			} else {
				item.SigProvider = types.StringNull()

			}
			if cValue := v.Get("tunnel-dc-preference.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TunnelDcPreference = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.TunnelDcPreference = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("tunnel-dc-preference.vipValue")
					item.TunnelDcPreference = types.StringValue(cv.String())

				}
			} else {
				item.TunnelDcPreference = types.StringNull()

			}
			if cValue := v.Get("tcp-mss-adjust.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TcpMss = types.Int64Null()

					cv := v.Get("tcp-mss-adjust.vipVariableName")
					item.TcpMssVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TcpMss = types.Int64Null()
					item.TcpMssVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("tcp-mss-adjust.vipValue")
					item.TcpMss = types.Int64Value(cv.Int())
					item.TcpMssVariable = types.StringNull()
				}
			} else {
				item.TcpMss = types.Int64Null()
				item.TcpMssVariable = types.StringNull()
			}
			if cValue := v.Get("mtu.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Mtu = types.Int64Null()

					cv := v.Get("mtu.vipVariableName")
					item.MtuVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Mtu = types.Int64Null()
					item.MtuVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("mtu.vipValue")
					item.Mtu = types.Int64Value(cv.Int())
					item.MtuVariable = types.StringNull()
				}
			} else {
				item.Mtu = types.Int64Null()
				item.MtuVariable = types.StringNull()
			}
			if cValue := v.Get("dead-peer-detection.dpd-interval.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.DeadPeerDetectionInterval = types.Int64Null()

					cv := v.Get("dead-peer-detection.dpd-interval.vipVariableName")
					item.DeadPeerDetectionIntervalVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.DeadPeerDetectionInterval = types.Int64Null()
					item.DeadPeerDetectionIntervalVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("dead-peer-detection.dpd-interval.vipValue")
					item.DeadPeerDetectionInterval = types.Int64Value(cv.Int())
					item.DeadPeerDetectionIntervalVariable = types.StringNull()
				}
			} else {
				item.DeadPeerDetectionInterval = types.Int64Null()
				item.DeadPeerDetectionIntervalVariable = types.StringNull()
			}
			if cValue := v.Get("dead-peer-detection.dpd-retries.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.DeadPeerDetectionRetries = types.Int64Null()

					cv := v.Get("dead-peer-detection.dpd-retries.vipVariableName")
					item.DeadPeerDetectionRetriesVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.DeadPeerDetectionRetries = types.Int64Null()
					item.DeadPeerDetectionRetriesVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("dead-peer-detection.dpd-retries.vipValue")
					item.DeadPeerDetectionRetries = types.Int64Value(cv.Int())
					item.DeadPeerDetectionRetriesVariable = types.StringNull()
				}
			} else {
				item.DeadPeerDetectionRetries = types.Int64Null()
				item.DeadPeerDetectionRetriesVariable = types.StringNull()
			}
			if cValue := v.Get("ike.ike-version.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.IkeVersion = types.Int64Null()

					cv := v.Get("ike.ike-version.vipVariableName")
					item.IkeVersionVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.IkeVersion = types.Int64Null()
					item.IkeVersionVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("ike.ike-version.vipValue")
					item.IkeVersion = types.Int64Value(cv.Int())
					item.IkeVersionVariable = types.StringNull()
				}
			} else {
				item.IkeVersion = types.Int64Null()
				item.IkeVersionVariable = types.StringNull()
			}
			if cValue := v.Get("ike.authentication-type.pre-shared-key.pre-shared-secret.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.IkePreSharedKey = types.StringNull()

					cv := v.Get("ike.authentication-type.pre-shared-key.pre-shared-secret.vipVariableName")
					item.IkePreSharedKeyVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.IkePreSharedKey = types.StringNull()
					item.IkePreSharedKeyVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("ike.authentication-type.pre-shared-key.pre-shared-secret.vipValue")
					item.IkePreSharedKey = types.StringValue(cv.String())
					item.IkePreSharedKeyVariable = types.StringNull()
				}
			} else {
				item.IkePreSharedKey = types.StringNull()
				item.IkePreSharedKeyVariable = types.StringNull()
			}
			if cValue := v.Get("ike.ike-rekey-interval.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.IkeRekeyInterval = types.Int64Null()

					cv := v.Get("ike.ike-rekey-interval.vipVariableName")
					item.IkeRekeyIntervalVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.IkeRekeyInterval = types.Int64Null()
					item.IkeRekeyIntervalVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("ike.ike-rekey-interval.vipValue")
					item.IkeRekeyInterval = types.Int64Value(cv.Int())
					item.IkeRekeyIntervalVariable = types.StringNull()
				}
			} else {
				item.IkeRekeyInterval = types.Int64Null()
				item.IkeRekeyIntervalVariable = types.StringNull()
			}
			if cValue := v.Get("ike.ike-ciphersuite.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.IkeCiphersuite = types.StringNull()

					cv := v.Get("ike.ike-ciphersuite.vipVariableName")
					item.IkeCiphersuiteVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.IkeCiphersuite = types.StringNull()
					item.IkeCiphersuiteVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("ike.ike-ciphersuite.vipValue")
					item.IkeCiphersuite = types.StringValue(cv.String())
					item.IkeCiphersuiteVariable = types.StringNull()
				}
			} else {
				item.IkeCiphersuite = types.StringNull()
				item.IkeCiphersuiteVariable = types.StringNull()
			}
			if cValue := v.Get("ike.ike-group.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.IkeGroup = types.StringNull()

					cv := v.Get("ike.ike-group.vipVariableName")
					item.IkeGroupVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.IkeGroup = types.StringNull()
					item.IkeGroupVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("ike.ike-group.vipValue")
					item.IkeGroup = types.StringValue(cv.String())
					item.IkeGroupVariable = types.StringNull()
				}
			} else {
				item.IkeGroup = types.StringNull()
				item.IkeGroupVariable = types.StringNull()
			}
			if cValue := v.Get("ike.authentication-type.pre-shared-key-dynamic.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.IkePreSharedKeyDynamic = types.BoolNull()

				} else if cValue.String() == "ignore" {
					item.IkePreSharedKeyDynamic = types.BoolNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("ike.authentication-type.pre-shared-key-dynamic.vipValue")
					item.IkePreSharedKeyDynamic = types.BoolValue(cv.Bool())

				}
			} else {
				item.IkePreSharedKeyDynamic = types.BoolNull()

			}
			if cValue := v.Get("ike.authentication-type.pre-shared-key.ike-local-id.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.IkePreSharedKeyLocalId = types.StringNull()

					cv := v.Get("ike.authentication-type.pre-shared-key.ike-local-id.vipVariableName")
					item.IkePreSharedKeyLocalIdVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.IkePreSharedKeyLocalId = types.StringNull()
					item.IkePreSharedKeyLocalIdVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("ike.authentication-type.pre-shared-key.ike-local-id.vipValue")
					item.IkePreSharedKeyLocalId = types.StringValue(cv.String())
					item.IkePreSharedKeyLocalIdVariable = types.StringNull()
				}
			} else {
				item.IkePreSharedKeyLocalId = types.StringNull()
				item.IkePreSharedKeyLocalIdVariable = types.StringNull()
			}
			if cValue := v.Get("ike.authentication-type.pre-shared-key.ike-remote-id.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.IkePreSharedKeyRemoteId = types.StringNull()

					cv := v.Get("ike.authentication-type.pre-shared-key.ike-remote-id.vipVariableName")
					item.IkePreSharedKeyRemoteIdVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.IkePreSharedKeyRemoteId = types.StringNull()
					item.IkePreSharedKeyRemoteIdVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("ike.authentication-type.pre-shared-key.ike-remote-id.vipValue")
					item.IkePreSharedKeyRemoteId = types.StringValue(cv.String())
					item.IkePreSharedKeyRemoteIdVariable = types.StringNull()
				}
			} else {
				item.IkePreSharedKeyRemoteId = types.StringNull()
				item.IkePreSharedKeyRemoteIdVariable = types.StringNull()
			}
			if cValue := v.Get("ipsec.ipsec-rekey-interval.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.IpsecRekeyInterval = types.Int64Null()

					cv := v.Get("ipsec.ipsec-rekey-interval.vipVariableName")
					item.IpsecRekeyIntervalVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.IpsecRekeyInterval = types.Int64Null()
					item.IpsecRekeyIntervalVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("ipsec.ipsec-rekey-interval.vipValue")
					item.IpsecRekeyInterval = types.Int64Value(cv.Int())
					item.IpsecRekeyIntervalVariable = types.StringNull()
				}
			} else {
				item.IpsecRekeyInterval = types.Int64Null()
				item.IpsecRekeyIntervalVariable = types.StringNull()
			}
			if cValue := v.Get("ipsec.ipsec-replay-window.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.IpsecReplayWindow = types.Int64Null()

					cv := v.Get("ipsec.ipsec-replay-window.vipVariableName")
					item.IpsecReplayWindowVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.IpsecReplayWindow = types.Int64Null()
					item.IpsecReplayWindowVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("ipsec.ipsec-replay-window.vipValue")
					item.IpsecReplayWindow = types.Int64Value(cv.Int())
					item.IpsecReplayWindowVariable = types.StringNull()
				}
			} else {
				item.IpsecReplayWindow = types.Int64Null()
				item.IpsecReplayWindowVariable = types.StringNull()
			}
			if cValue := v.Get("ipsec.ipsec-ciphersuite.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.IpsecCiphersuite = types.StringNull()

					cv := v.Get("ipsec.ipsec-ciphersuite.vipVariableName")
					item.IpsecCiphersuiteVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.IpsecCiphersuite = types.StringNull()
					item.IpsecCiphersuiteVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("ipsec.ipsec-ciphersuite.vipValue")
					item.IpsecCiphersuite = types.StringValue(cv.String())
					item.IpsecCiphersuiteVariable = types.StringNull()
				}
			} else {
				item.IpsecCiphersuite = types.StringNull()
				item.IpsecCiphersuiteVariable = types.StringNull()
			}
			if cValue := v.Get("ipsec.perfect-forward-secrecy.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.IpsecPerfectForwardSecrecy = types.StringNull()

					cv := v.Get("ipsec.perfect-forward-secrecy.vipVariableName")
					item.IpsecPerfectForwardSecrecyVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.IpsecPerfectForwardSecrecy = types.StringNull()
					item.IpsecPerfectForwardSecrecyVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("ipsec.perfect-forward-secrecy.vipValue")
					item.IpsecPerfectForwardSecrecy = types.StringValue(cv.String())
					item.IpsecPerfectForwardSecrecyVariable = types.StringNull()
				}
			} else {
				item.IpsecPerfectForwardSecrecy = types.StringNull()
				item.IpsecPerfectForwardSecrecyVariable = types.StringNull()
			}
			if cValue := v.Get("tracker.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Tracker = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.Tracker = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("tracker.vipValue")
					item.Tracker = types.StringValue(cv.String())

				}
			} else {
				item.Tracker = types.StringNull()

			}
			if cValue := v.Get("track-enable.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TrackEnable = types.BoolNull()

				} else if cValue.String() == "ignore" {
					item.TrackEnable = types.BoolNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("track-enable.vipValue")
					item.TrackEnable = types.BoolValue(cv.Bool())

				}
			} else {
				item.TrackEnable = types.BoolNull()

			}
			if cValue := v.Get("tunnel-public-ip.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TunnelPublicIp = types.StringNull()

					cv := v.Get("tunnel-public-ip.vipVariableName")
					item.TunnelPublicIpVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TunnelPublicIp = types.StringNull()
					item.TunnelPublicIpVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("tunnel-public-ip.vipValue")
					item.TunnelPublicIp = types.StringValue(cv.String())
					item.TunnelPublicIpVariable = types.StringNull()
				}
			} else {
				item.TunnelPublicIp = types.StringNull()
				item.TunnelPublicIpVariable = types.StringNull()
			}
			data.Interfaces = append(data.Interfaces, item)
			return true
		})
	} else {
		if len(data.Interfaces) > 0 {
			data.Interfaces = []CiscoSecureInternetGatewayInterfaces{}
		}
	}
	if value := res.Get(path + "service.vipValue"); len(value.Array()) > 0 {
		data.Services = make([]CiscoSecureInternetGatewayServices, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoSecureInternetGatewayServices{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("svc-type.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ServiceType = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.ServiceType = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("svc-type.vipValue")
					item.ServiceType = types.StringValue(cv.String())

				}
			} else {
				item.ServiceType = types.StringNull()

			}
			if cValue := v.Get("ha-pairs.interface-pair.vipValue"); len(cValue.Array()) > 0 {
				item.InterfacePairs = make([]CiscoSecureInternetGatewayServicesInterfacePairs, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CiscoSecureInternetGatewayServicesInterfacePairs{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("active-interface.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.ActiveInterface = types.StringNull()

						} else if ccValue.String() == "ignore" {
							cItem.ActiveInterface = types.StringNull()

						} else if ccValue.String() == "constant" {
							ccv := cv.Get("active-interface.vipValue")
							cItem.ActiveInterface = types.StringValue(ccv.String())

						}
					} else {
						cItem.ActiveInterface = types.StringNull()

					}
					if ccValue := cv.Get("active-interface-weight.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.ActiveInterfaceWeight = types.Int64Null()

						} else if ccValue.String() == "ignore" {
							cItem.ActiveInterfaceWeight = types.Int64Null()

						} else if ccValue.String() == "constant" {
							ccv := cv.Get("active-interface-weight.vipValue")
							cItem.ActiveInterfaceWeight = types.Int64Value(ccv.Int())

						}
					} else {
						cItem.ActiveInterfaceWeight = types.Int64Null()

					}
					if ccValue := cv.Get("backup-interface.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.BackupInterface = types.StringNull()

						} else if ccValue.String() == "ignore" {
							cItem.BackupInterface = types.StringNull()

						} else if ccValue.String() == "constant" {
							ccv := cv.Get("backup-interface.vipValue")
							cItem.BackupInterface = types.StringValue(ccv.String())

						}
					} else {
						cItem.BackupInterface = types.StringNull()

					}
					if ccValue := cv.Get("backup-interface-weight.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.BackupInterfaceWeight = types.Int64Null()

						} else if ccValue.String() == "ignore" {
							cItem.BackupInterfaceWeight = types.Int64Null()

						} else if ccValue.String() == "constant" {
							ccv := cv.Get("backup-interface-weight.vipValue")
							cItem.BackupInterfaceWeight = types.Int64Value(ccv.Int())

						}
					} else {
						cItem.BackupInterfaceWeight = types.Int64Null()

					}
					item.InterfacePairs = append(item.InterfacePairs, cItem)
					return true
				})
			} else {
				if len(item.InterfacePairs) > 0 {
					item.InterfacePairs = []CiscoSecureInternetGatewayServicesInterfacePairs{}
				}
			}
			if cValue := v.Get("zscaler-location-settings.auth-required.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ZscalerAuthenticationRequired = types.BoolNull()

				} else if cValue.String() == "ignore" {
					item.ZscalerAuthenticationRequired = types.BoolNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("zscaler-location-settings.auth-required.vipValue")
					item.ZscalerAuthenticationRequired = types.BoolValue(cv.Bool())

				}
			} else {
				item.ZscalerAuthenticationRequired = types.BoolNull()

			}
			if cValue := v.Get("zscaler-location-settings.xff-forward-enabled.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ZscalerXffForward = types.BoolNull()

				} else if cValue.String() == "ignore" {
					item.ZscalerXffForward = types.BoolNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("zscaler-location-settings.xff-forward-enabled.vipValue")
					item.ZscalerXffForward = types.BoolValue(cv.Bool())

				}
			} else {
				item.ZscalerXffForward = types.BoolNull()

			}
			if cValue := v.Get("zscaler-location-settings.ofw-enabled.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ZscalerFirewallEnabled = types.BoolNull()

				} else if cValue.String() == "ignore" {
					item.ZscalerFirewallEnabled = types.BoolNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("zscaler-location-settings.ofw-enabled.vipValue")
					item.ZscalerFirewallEnabled = types.BoolValue(cv.Bool())

				}
			} else {
				item.ZscalerFirewallEnabled = types.BoolNull()

			}
			if cValue := v.Get("zscaler-location-settings.ips-control.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ZscalerIpsControlEnabled = types.BoolNull()

				} else if cValue.String() == "ignore" {
					item.ZscalerIpsControlEnabled = types.BoolNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("zscaler-location-settings.ips-control.vipValue")
					item.ZscalerIpsControlEnabled = types.BoolValue(cv.Bool())

				}
			} else {
				item.ZscalerIpsControlEnabled = types.BoolNull()

			}
			if cValue := v.Get("zscaler-location-settings.caution-enabled.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ZscalerCautionEnabled = types.BoolNull()

				} else if cValue.String() == "ignore" {
					item.ZscalerCautionEnabled = types.BoolNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("zscaler-location-settings.caution-enabled.vipValue")
					item.ZscalerCautionEnabled = types.BoolValue(cv.Bool())

				}
			} else {
				item.ZscalerCautionEnabled = types.BoolNull()

			}
			if cValue := v.Get("zscaler-location-settings.datacenters.primary-data-center.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ZscalerPrimaryDataCenter = types.StringNull()

					cv := v.Get("zscaler-location-settings.datacenters.primary-data-center.vipVariableName")
					item.ZscalerPrimaryDataCenterVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.ZscalerPrimaryDataCenter = types.StringNull()
					item.ZscalerPrimaryDataCenterVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("zscaler-location-settings.datacenters.primary-data-center.vipValue")
					item.ZscalerPrimaryDataCenter = types.StringValue(cv.String())
					item.ZscalerPrimaryDataCenterVariable = types.StringNull()
				}
			} else {
				item.ZscalerPrimaryDataCenter = types.StringNull()
				item.ZscalerPrimaryDataCenterVariable = types.StringNull()
			}
			if cValue := v.Get("zscaler-location-settings.datacenters.secondary-data-center.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ZscalerSecondaryDataCenter = types.StringNull()

					cv := v.Get("zscaler-location-settings.datacenters.secondary-data-center.vipVariableName")
					item.ZscalerSecondaryDataCenterVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.ZscalerSecondaryDataCenter = types.StringNull()
					item.ZscalerSecondaryDataCenterVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("zscaler-location-settings.datacenters.secondary-data-center.vipValue")
					item.ZscalerSecondaryDataCenter = types.StringValue(cv.String())
					item.ZscalerSecondaryDataCenterVariable = types.StringNull()
				}
			} else {
				item.ZscalerSecondaryDataCenter = types.StringNull()
				item.ZscalerSecondaryDataCenterVariable = types.StringNull()
			}
			if cValue := v.Get("zscaler-location-settings.surrogate.ip.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ZscalerSurrogateIp = types.BoolNull()

				} else if cValue.String() == "ignore" {
					item.ZscalerSurrogateIp = types.BoolNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("zscaler-location-settings.surrogate.ip.vipValue")
					item.ZscalerSurrogateIp = types.BoolValue(cv.Bool())

				}
			} else {
				item.ZscalerSurrogateIp = types.BoolNull()

			}
			if cValue := v.Get("zscaler-location-settings.surrogate.idle-time.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ZscalerSurrogateIdleTime = types.Int64Null()

				} else if cValue.String() == "ignore" {
					item.ZscalerSurrogateIdleTime = types.Int64Null()

				} else if cValue.String() == "constant" {
					cv := v.Get("zscaler-location-settings.surrogate.idle-time.vipValue")
					item.ZscalerSurrogateIdleTime = types.Int64Value(cv.Int())

				}
			} else {
				item.ZscalerSurrogateIdleTime = types.Int64Null()

			}
			if cValue := v.Get("zscaler-location-settings.surrogate.display-time-unit.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ZscalerSurrogateDisplayTimeUnit = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.ZscalerSurrogateDisplayTimeUnit = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("zscaler-location-settings.surrogate.display-time-unit.vipValue")
					item.ZscalerSurrogateDisplayTimeUnit = types.StringValue(cv.String())

				}
			} else {
				item.ZscalerSurrogateDisplayTimeUnit = types.StringNull()

			}
			if cValue := v.Get("zscaler-location-settings.surrogate.ip-enforced-for-known-browsers.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ZscalerSurrogateIpEnforceForKnownBrowsers = types.BoolNull()

				} else if cValue.String() == "ignore" {
					item.ZscalerSurrogateIpEnforceForKnownBrowsers = types.BoolNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("zscaler-location-settings.surrogate.ip-enforced-for-known-browsers.vipValue")
					item.ZscalerSurrogateIpEnforceForKnownBrowsers = types.BoolValue(cv.Bool())

				}
			} else {
				item.ZscalerSurrogateIpEnforceForKnownBrowsers = types.BoolNull()

			}
			if cValue := v.Get("zscaler-location-settings.surrogate.refresh-time.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ZscalerSurrogateRefreshTime = types.Int64Null()

				} else if cValue.String() == "ignore" {
					item.ZscalerSurrogateRefreshTime = types.Int64Null()

				} else if cValue.String() == "constant" {
					cv := v.Get("zscaler-location-settings.surrogate.refresh-time.vipValue")
					item.ZscalerSurrogateRefreshTime = types.Int64Value(cv.Int())

				}
			} else {
				item.ZscalerSurrogateRefreshTime = types.Int64Null()

			}
			if cValue := v.Get("zscaler-location-settings.surrogate.refresh-time-unit.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ZscalerSurrogateRefreshTimeUnit = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.ZscalerSurrogateRefreshTimeUnit = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("zscaler-location-settings.surrogate.refresh-time-unit.vipValue")
					item.ZscalerSurrogateRefreshTimeUnit = types.StringValue(cv.String())

				}
			} else {
				item.ZscalerSurrogateRefreshTimeUnit = types.StringNull()

			}
			if cValue := v.Get("zscaler-location-settings.aup.enabled.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ZscalerAupEnabled = types.BoolNull()

				} else if cValue.String() == "ignore" {
					item.ZscalerAupEnabled = types.BoolNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("zscaler-location-settings.aup.enabled.vipValue")
					item.ZscalerAupEnabled = types.BoolValue(cv.Bool())

				}
			} else {
				item.ZscalerAupEnabled = types.BoolNull()

			}
			if cValue := v.Get("zscaler-location-settings.aup.block-internet-until-accepted.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ZscalerAupBlockInternetUntilAccepted = types.BoolNull()

				} else if cValue.String() == "ignore" {
					item.ZscalerAupBlockInternetUntilAccepted = types.BoolNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("zscaler-location-settings.aup.block-internet-until-accepted.vipValue")
					item.ZscalerAupBlockInternetUntilAccepted = types.BoolValue(cv.Bool())

				}
			} else {
				item.ZscalerAupBlockInternetUntilAccepted = types.BoolNull()

			}
			if cValue := v.Get("zscaler-location-settings.aup.force-ssl-inspection.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ZscalerAupForceSslInspection = types.BoolNull()

				} else if cValue.String() == "ignore" {
					item.ZscalerAupForceSslInspection = types.BoolNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("zscaler-location-settings.aup.force-ssl-inspection.vipValue")
					item.ZscalerAupForceSslInspection = types.BoolValue(cv.Bool())

				}
			} else {
				item.ZscalerAupForceSslInspection = types.BoolNull()

			}
			if cValue := v.Get("zscaler-location-settings.aup.timeout.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ZscalerAupTimeout = types.Int64Null()

				} else if cValue.String() == "ignore" {
					item.ZscalerAupTimeout = types.Int64Null()

				} else if cValue.String() == "constant" {
					cv := v.Get("zscaler-location-settings.aup.timeout.vipValue")
					item.ZscalerAupTimeout = types.Int64Value(cv.Int())

				}
			} else {
				item.ZscalerAupTimeout = types.Int64Null()

			}
			if cValue := v.Get("zscaler-location-settings.location-name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ZscalerLocationName = types.StringNull()

					cv := v.Get("zscaler-location-settings.location-name.vipVariableName")
					item.ZscalerLocationNameVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.ZscalerLocationName = types.StringNull()
					item.ZscalerLocationNameVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("zscaler-location-settings.location-name.vipValue")
					item.ZscalerLocationName = types.StringValue(cv.String())
					item.ZscalerLocationNameVariable = types.StringNull()
				}
			} else {
				item.ZscalerLocationName = types.StringNull()
				item.ZscalerLocationNameVariable = types.StringNull()
			}
			if cValue := v.Get("umbrella-data-center.data-center-primary.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.UmbrellaPrimaryDataCenter = types.StringNull()

					cv := v.Get("umbrella-data-center.data-center-primary.vipVariableName")
					item.UmbrellaPrimaryDataCenterVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.UmbrellaPrimaryDataCenter = types.StringNull()
					item.UmbrellaPrimaryDataCenterVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("umbrella-data-center.data-center-primary.vipValue")
					item.UmbrellaPrimaryDataCenter = types.StringValue(cv.String())
					item.UmbrellaPrimaryDataCenterVariable = types.StringNull()
				}
			} else {
				item.UmbrellaPrimaryDataCenter = types.StringNull()
				item.UmbrellaPrimaryDataCenterVariable = types.StringNull()
			}
			if cValue := v.Get("umbrella-data-center.data-center-secondary.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.UmbrellaSecondaryDataCenter = types.StringNull()

					cv := v.Get("umbrella-data-center.data-center-secondary.vipVariableName")
					item.UmbrellaSecondaryDataCenterVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.UmbrellaSecondaryDataCenter = types.StringNull()
					item.UmbrellaSecondaryDataCenterVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("umbrella-data-center.data-center-secondary.vipValue")
					item.UmbrellaSecondaryDataCenter = types.StringValue(cv.String())
					item.UmbrellaSecondaryDataCenterVariable = types.StringNull()
				}
			} else {
				item.UmbrellaSecondaryDataCenter = types.StringNull()
				item.UmbrellaSecondaryDataCenterVariable = types.StringNull()
			}
			data.Services = append(data.Services, item)
			return true
		})
	} else {
		if len(data.Services) > 0 {
			data.Services = []CiscoSecureInternetGatewayServices{}
		}
	}
	if value := res.Get(path + "tracker-src-ip.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TrackerSourceIp = types.StringNull()

			v := res.Get(path + "tracker-src-ip.vipVariableName")
			data.TrackerSourceIpVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TrackerSourceIp = types.StringNull()
			data.TrackerSourceIpVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tracker-src-ip.vipValue")
			data.TrackerSourceIp = types.StringValue(v.String())
			data.TrackerSourceIpVariable = types.StringNull()
		}
	} else {
		data.TrackerSourceIp = types.StringNull()
		data.TrackerSourceIpVariable = types.StringNull()
	}
	if value := res.Get(path + "tracker.vipValue"); len(value.Array()) > 0 {
		data.Trackers = make([]CiscoSecureInternetGatewayTrackers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoSecureInternetGatewayTrackers{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Name = types.StringNull()

					cv := v.Get("name.vipVariableName")
					item.NameVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Name = types.StringNull()
					item.NameVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("name.vipValue")
					item.Name = types.StringValue(cv.String())
					item.NameVariable = types.StringNull()
				}
			} else {
				item.Name = types.StringNull()
				item.NameVariable = types.StringNull()
			}
			if cValue := v.Get("endpoint-api-url.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.EndpointApiUrl = types.StringNull()

					cv := v.Get("endpoint-api-url.vipVariableName")
					item.EndpointApiUrlVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.EndpointApiUrl = types.StringNull()
					item.EndpointApiUrlVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("endpoint-api-url.vipValue")
					item.EndpointApiUrl = types.StringValue(cv.String())
					item.EndpointApiUrlVariable = types.StringNull()
				}
			} else {
				item.EndpointApiUrl = types.StringNull()
				item.EndpointApiUrlVariable = types.StringNull()
			}
			if cValue := v.Get("threshold.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Threshold = types.Int64Null()

					cv := v.Get("threshold.vipVariableName")
					item.ThresholdVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Threshold = types.Int64Null()
					item.ThresholdVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("threshold.vipValue")
					item.Threshold = types.Int64Value(cv.Int())
					item.ThresholdVariable = types.StringNull()
				}
			} else {
				item.Threshold = types.Int64Null()
				item.ThresholdVariable = types.StringNull()
			}
			if cValue := v.Get("interval.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Interval = types.Int64Null()

					cv := v.Get("interval.vipVariableName")
					item.IntervalVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Interval = types.Int64Null()
					item.IntervalVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("interval.vipValue")
					item.Interval = types.Int64Value(cv.Int())
					item.IntervalVariable = types.StringNull()
				}
			} else {
				item.Interval = types.Int64Null()
				item.IntervalVariable = types.StringNull()
			}
			if cValue := v.Get("multiplier.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Multiplier = types.Int64Null()

					cv := v.Get("multiplier.vipVariableName")
					item.MultiplierVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Multiplier = types.Int64Null()
					item.MultiplierVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("multiplier.vipValue")
					item.Multiplier = types.Int64Value(cv.Int())
					item.MultiplierVariable = types.StringNull()
				}
			} else {
				item.Multiplier = types.Int64Null()
				item.MultiplierVariable = types.StringNull()
			}
			if cValue := v.Get("tracker-type.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TrackerType = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.TrackerType = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("tracker-type.vipValue")
					item.TrackerType = types.StringValue(cv.String())

				}
			} else {
				item.TrackerType = types.StringNull()

			}
			data.Trackers = append(data.Trackers, item)
			return true
		})
	} else {
		if len(data.Trackers) > 0 {
			data.Trackers = []CiscoSecureInternetGatewayTrackers{}
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *CiscoSecureInternetGateway) hasChanges(ctx context.Context, state *CiscoSecureInternetGateway) bool {
	hasChanges := false
	if !data.VpnId.Equal(state.VpnId) {
		hasChanges = true
	}
	if len(data.Interfaces) != len(state.Interfaces) {
		hasChanges = true
	} else {
		for i := range data.Interfaces {
			if !data.Interfaces[i].Name.Equal(state.Interfaces[i].Name) {
				hasChanges = true
			}
			if !data.Interfaces[i].AutoTunnelMode.Equal(state.Interfaces[i].AutoTunnelMode) {
				hasChanges = true
			}
			if !data.Interfaces[i].Shutdown.Equal(state.Interfaces[i].Shutdown) {
				hasChanges = true
			}
			if !data.Interfaces[i].Description.Equal(state.Interfaces[i].Description) {
				hasChanges = true
			}
			if !data.Interfaces[i].IpUnnumbered.Equal(state.Interfaces[i].IpUnnumbered) {
				hasChanges = true
			}
			if !data.Interfaces[i].Ipv4Address.Equal(state.Interfaces[i].Ipv4Address) {
				hasChanges = true
			}
			if !data.Interfaces[i].TunnelSource.Equal(state.Interfaces[i].TunnelSource) {
				hasChanges = true
			}
			if !data.Interfaces[i].TunnelSourceInterface.Equal(state.Interfaces[i].TunnelSourceInterface) {
				hasChanges = true
			}
			if !data.Interfaces[i].TunnelRouteVia.Equal(state.Interfaces[i].TunnelRouteVia) {
				hasChanges = true
			}
			if !data.Interfaces[i].TunnelDestination.Equal(state.Interfaces[i].TunnelDestination) {
				hasChanges = true
			}
			if !data.Interfaces[i].Application.Equal(state.Interfaces[i].Application) {
				hasChanges = true
			}
			if !data.Interfaces[i].SigProvider.Equal(state.Interfaces[i].SigProvider) {
				hasChanges = true
			}
			if !data.Interfaces[i].TunnelDcPreference.Equal(state.Interfaces[i].TunnelDcPreference) {
				hasChanges = true
			}
			if !data.Interfaces[i].TcpMss.Equal(state.Interfaces[i].TcpMss) {
				hasChanges = true
			}
			if !data.Interfaces[i].Mtu.Equal(state.Interfaces[i].Mtu) {
				hasChanges = true
			}
			if !data.Interfaces[i].DeadPeerDetectionInterval.Equal(state.Interfaces[i].DeadPeerDetectionInterval) {
				hasChanges = true
			}
			if !data.Interfaces[i].DeadPeerDetectionRetries.Equal(state.Interfaces[i].DeadPeerDetectionRetries) {
				hasChanges = true
			}
			if !data.Interfaces[i].IkeVersion.Equal(state.Interfaces[i].IkeVersion) {
				hasChanges = true
			}
			if !data.Interfaces[i].IkePreSharedKey.Equal(state.Interfaces[i].IkePreSharedKey) {
				hasChanges = true
			}
			if !data.Interfaces[i].IkeRekeyInterval.Equal(state.Interfaces[i].IkeRekeyInterval) {
				hasChanges = true
			}
			if !data.Interfaces[i].IkeCiphersuite.Equal(state.Interfaces[i].IkeCiphersuite) {
				hasChanges = true
			}
			if !data.Interfaces[i].IkeGroup.Equal(state.Interfaces[i].IkeGroup) {
				hasChanges = true
			}
			if !data.Interfaces[i].IkePreSharedKeyDynamic.Equal(state.Interfaces[i].IkePreSharedKeyDynamic) {
				hasChanges = true
			}
			if !data.Interfaces[i].IkePreSharedKeyLocalId.Equal(state.Interfaces[i].IkePreSharedKeyLocalId) {
				hasChanges = true
			}
			if !data.Interfaces[i].IkePreSharedKeyRemoteId.Equal(state.Interfaces[i].IkePreSharedKeyRemoteId) {
				hasChanges = true
			}
			if !data.Interfaces[i].IpsecRekeyInterval.Equal(state.Interfaces[i].IpsecRekeyInterval) {
				hasChanges = true
			}
			if !data.Interfaces[i].IpsecReplayWindow.Equal(state.Interfaces[i].IpsecReplayWindow) {
				hasChanges = true
			}
			if !data.Interfaces[i].IpsecCiphersuite.Equal(state.Interfaces[i].IpsecCiphersuite) {
				hasChanges = true
			}
			if !data.Interfaces[i].IpsecPerfectForwardSecrecy.Equal(state.Interfaces[i].IpsecPerfectForwardSecrecy) {
				hasChanges = true
			}
			if !data.Interfaces[i].Tracker.Equal(state.Interfaces[i].Tracker) {
				hasChanges = true
			}
			if !data.Interfaces[i].TrackEnable.Equal(state.Interfaces[i].TrackEnable) {
				hasChanges = true
			}
			if !data.Interfaces[i].TunnelPublicIp.Equal(state.Interfaces[i].TunnelPublicIp) {
				hasChanges = true
			}
		}
	}
	if len(data.Services) != len(state.Services) {
		hasChanges = true
	} else {
		for i := range data.Services {
			if !data.Services[i].ServiceType.Equal(state.Services[i].ServiceType) {
				hasChanges = true
			}
			if len(data.Services[i].InterfacePairs) != len(state.Services[i].InterfacePairs) {
				hasChanges = true
			} else {
				for ii := range data.Services[i].InterfacePairs {
					if !data.Services[i].InterfacePairs[ii].ActiveInterface.Equal(state.Services[i].InterfacePairs[ii].ActiveInterface) {
						hasChanges = true
					}
					if !data.Services[i].InterfacePairs[ii].ActiveInterfaceWeight.Equal(state.Services[i].InterfacePairs[ii].ActiveInterfaceWeight) {
						hasChanges = true
					}
					if !data.Services[i].InterfacePairs[ii].BackupInterface.Equal(state.Services[i].InterfacePairs[ii].BackupInterface) {
						hasChanges = true
					}
					if !data.Services[i].InterfacePairs[ii].BackupInterfaceWeight.Equal(state.Services[i].InterfacePairs[ii].BackupInterfaceWeight) {
						hasChanges = true
					}
				}
			}
			if !data.Services[i].ZscalerAuthenticationRequired.Equal(state.Services[i].ZscalerAuthenticationRequired) {
				hasChanges = true
			}
			if !data.Services[i].ZscalerXffForward.Equal(state.Services[i].ZscalerXffForward) {
				hasChanges = true
			}
			if !data.Services[i].ZscalerFirewallEnabled.Equal(state.Services[i].ZscalerFirewallEnabled) {
				hasChanges = true
			}
			if !data.Services[i].ZscalerIpsControlEnabled.Equal(state.Services[i].ZscalerIpsControlEnabled) {
				hasChanges = true
			}
			if !data.Services[i].ZscalerCautionEnabled.Equal(state.Services[i].ZscalerCautionEnabled) {
				hasChanges = true
			}
			if !data.Services[i].ZscalerPrimaryDataCenter.Equal(state.Services[i].ZscalerPrimaryDataCenter) {
				hasChanges = true
			}
			if !data.Services[i].ZscalerSecondaryDataCenter.Equal(state.Services[i].ZscalerSecondaryDataCenter) {
				hasChanges = true
			}
			if !data.Services[i].ZscalerSurrogateIp.Equal(state.Services[i].ZscalerSurrogateIp) {
				hasChanges = true
			}
			if !data.Services[i].ZscalerSurrogateIdleTime.Equal(state.Services[i].ZscalerSurrogateIdleTime) {
				hasChanges = true
			}
			if !data.Services[i].ZscalerSurrogateDisplayTimeUnit.Equal(state.Services[i].ZscalerSurrogateDisplayTimeUnit) {
				hasChanges = true
			}
			if !data.Services[i].ZscalerSurrogateIpEnforceForKnownBrowsers.Equal(state.Services[i].ZscalerSurrogateIpEnforceForKnownBrowsers) {
				hasChanges = true
			}
			if !data.Services[i].ZscalerSurrogateRefreshTime.Equal(state.Services[i].ZscalerSurrogateRefreshTime) {
				hasChanges = true
			}
			if !data.Services[i].ZscalerSurrogateRefreshTimeUnit.Equal(state.Services[i].ZscalerSurrogateRefreshTimeUnit) {
				hasChanges = true
			}
			if !data.Services[i].ZscalerAupEnabled.Equal(state.Services[i].ZscalerAupEnabled) {
				hasChanges = true
			}
			if !data.Services[i].ZscalerAupBlockInternetUntilAccepted.Equal(state.Services[i].ZscalerAupBlockInternetUntilAccepted) {
				hasChanges = true
			}
			if !data.Services[i].ZscalerAupForceSslInspection.Equal(state.Services[i].ZscalerAupForceSslInspection) {
				hasChanges = true
			}
			if !data.Services[i].ZscalerAupTimeout.Equal(state.Services[i].ZscalerAupTimeout) {
				hasChanges = true
			}
			if !data.Services[i].ZscalerLocationName.Equal(state.Services[i].ZscalerLocationName) {
				hasChanges = true
			}
			if !data.Services[i].UmbrellaPrimaryDataCenter.Equal(state.Services[i].UmbrellaPrimaryDataCenter) {
				hasChanges = true
			}
			if !data.Services[i].UmbrellaSecondaryDataCenter.Equal(state.Services[i].UmbrellaSecondaryDataCenter) {
				hasChanges = true
			}
		}
	}
	if !data.TrackerSourceIp.Equal(state.TrackerSourceIp) {
		hasChanges = true
	}
	if len(data.Trackers) != len(state.Trackers) {
		hasChanges = true
	} else {
		for i := range data.Trackers {
			if !data.Trackers[i].Name.Equal(state.Trackers[i].Name) {
				hasChanges = true
			}
			if !data.Trackers[i].EndpointApiUrl.Equal(state.Trackers[i].EndpointApiUrl) {
				hasChanges = true
			}
			if !data.Trackers[i].Threshold.Equal(state.Trackers[i].Threshold) {
				hasChanges = true
			}
			if !data.Trackers[i].Interval.Equal(state.Trackers[i].Interval) {
				hasChanges = true
			}
			if !data.Trackers[i].Multiplier.Equal(state.Trackers[i].Multiplier) {
				hasChanges = true
			}
			if !data.Trackers[i].TrackerType.Equal(state.Trackers[i].TrackerType) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges
