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
type CiscoVPN struct {
	Id                           types.String                      `tfsdk:"id"`
	Version                      types.Int64                       `tfsdk:"version"`
	TemplateType                 types.String                      `tfsdk:"template_type"`
	Name                         types.String                      `tfsdk:"name"`
	Description                  types.String                      `tfsdk:"description"`
	DeviceTypes                  types.Set                         `tfsdk:"device_types"`
	VpnId                        types.Int64                       `tfsdk:"vpn_id"`
	VpnName                      types.String                      `tfsdk:"vpn_name"`
	VpnNameVariable              types.String                      `tfsdk:"vpn_name_variable"`
	TenantVpnId                  types.Int64                       `tfsdk:"tenant_vpn_id"`
	OrganizationName             types.String                      `tfsdk:"organization_name"`
	OmpAdminDistanceIpv4         types.Int64                       `tfsdk:"omp_admin_distance_ipv4"`
	OmpAdminDistanceIpv4Variable types.String                      `tfsdk:"omp_admin_distance_ipv4_variable"`
	OmpAdminDistanceIpv6         types.Int64                       `tfsdk:"omp_admin_distance_ipv6"`
	OmpAdminDistanceIpv6Variable types.String                      `tfsdk:"omp_admin_distance_ipv6_variable"`
	EnhanceEcmpKeying            types.Bool                        `tfsdk:"enhance_ecmp_keying"`
	EnhanceEcmpKeyingVariable    types.String                      `tfsdk:"enhance_ecmp_keying_variable"`
	DnsIpv4Servers               []CiscoVPNDnsIpv4Servers          `tfsdk:"dns_ipv4_servers"`
	DnsIpv6Servers               []CiscoVPNDnsIpv6Servers          `tfsdk:"dns_ipv6_servers"`
	DnsHosts                     []CiscoVPNDnsHosts                `tfsdk:"dns_hosts"`
	Services                     []CiscoVPNServices                `tfsdk:"services"`
	Ipv4StaticServiceRoutes      []CiscoVPNIpv4StaticServiceRoutes `tfsdk:"ipv4_static_service_routes"`
	Ipv4StaticRoutes             []CiscoVPNIpv4StaticRoutes        `tfsdk:"ipv4_static_routes"`
	Ipv6StaticRoutes             []CiscoVPNIpv6StaticRoutes        `tfsdk:"ipv6_static_routes"`
	Ipv4StaticGreRoutes          []CiscoVPNIpv4StaticGreRoutes     `tfsdk:"ipv4_static_gre_routes"`
	Ipv4StaticIpsecRoutes        []CiscoVPNIpv4StaticIpsecRoutes   `tfsdk:"ipv4_static_ipsec_routes"`
	OmpAdvertiseIpv4Routes       []CiscoVPNOmpAdvertiseIpv4Routes  `tfsdk:"omp_advertise_ipv4_routes"`
	OmpAdvertiseIpv6Routes       []CiscoVPNOmpAdvertiseIpv6Routes  `tfsdk:"omp_advertise_ipv6_routes"`
	Nat64Pools                   []CiscoVPNNat64Pools              `tfsdk:"nat64_pools"`
	NatPools                     []CiscoVPNNatPools                `tfsdk:"nat_pools"`
	StaticNatRules               []CiscoVPNStaticNatRules          `tfsdk:"static_nat_rules"`
	StaticNatSubnetRules         []CiscoVPNStaticNatSubnetRules    `tfsdk:"static_nat_subnet_rules"`
	PortForwardRules             []CiscoVPNPortForwardRules        `tfsdk:"port_forward_rules"`
	RouteGlobalImports           []CiscoVPNRouteGlobalImports      `tfsdk:"route_global_imports"`
	RouteVpnImports              []CiscoVPNRouteVpnImports         `tfsdk:"route_vpn_imports"`
	RouteGlobalExports           []CiscoVPNRouteGlobalExports      `tfsdk:"route_global_exports"`
}

type CiscoVPNDnsIpv4Servers struct {
	Optional        types.Bool   `tfsdk:"optional"`
	Address         types.String `tfsdk:"address"`
	AddressVariable types.String `tfsdk:"address_variable"`
	Role            types.String `tfsdk:"role"`
	RoleVariable    types.String `tfsdk:"role_variable"`
}

type CiscoVPNDnsIpv6Servers struct {
	Optional     types.Bool   `tfsdk:"optional"`
	Address      types.String `tfsdk:"address"`
	Role         types.String `tfsdk:"role"`
	RoleVariable types.String `tfsdk:"role_variable"`
}

type CiscoVPNDnsHosts struct {
	Optional         types.Bool   `tfsdk:"optional"`
	Hostname         types.String `tfsdk:"hostname"`
	HostnameVariable types.String `tfsdk:"hostname_variable"`
	Ip               types.Set    `tfsdk:"ip"`
	IpVariable       types.String `tfsdk:"ip_variable"`
}

type CiscoVPNServices struct {
	Optional            types.Bool   `tfsdk:"optional"`
	ServiceTypes        types.String `tfsdk:"service_types"`
	Address             types.Set    `tfsdk:"address"`
	AddressVariable     types.String `tfsdk:"address_variable"`
	Interface           types.String `tfsdk:"interface"`
	InterfaceVariable   types.String `tfsdk:"interface_variable"`
	TrackEnable         types.Bool   `tfsdk:"track_enable"`
	TrackEnableVariable types.String `tfsdk:"track_enable_variable"`
}

type CiscoVPNIpv4StaticServiceRoutes struct {
	Optional       types.Bool   `tfsdk:"optional"`
	Prefix         types.String `tfsdk:"prefix"`
	PrefixVariable types.String `tfsdk:"prefix_variable"`
	VpnId          types.Int64  `tfsdk:"vpn_id"`
	Service        types.String `tfsdk:"service"`
}

type CiscoVPNIpv4StaticRoutes struct {
	Optional         types.Bool                              `tfsdk:"optional"`
	Prefix           types.String                            `tfsdk:"prefix"`
	PrefixVariable   types.String                            `tfsdk:"prefix_variable"`
	Null0            types.Bool                              `tfsdk:"null0"`
	Null0Variable    types.String                            `tfsdk:"null0_variable"`
	Distance         types.Int64                             `tfsdk:"distance"`
	DistanceVariable types.String                            `tfsdk:"distance_variable"`
	VpnId            types.Int64                             `tfsdk:"vpn_id"`
	VpnIdVariable    types.String                            `tfsdk:"vpn_id_variable"`
	Dhcp             types.Bool                              `tfsdk:"dhcp"`
	DhcpVariable     types.String                            `tfsdk:"dhcp_variable"`
	NextHops         []CiscoVPNIpv4StaticRoutesNextHops      `tfsdk:"next_hops"`
	TrackNextHops    []CiscoVPNIpv4StaticRoutesTrackNextHops `tfsdk:"track_next_hops"`
}

type CiscoVPNIpv6StaticRoutes struct {
	Optional       types.Bool                         `tfsdk:"optional"`
	Prefix         types.String                       `tfsdk:"prefix"`
	PrefixVariable types.String                       `tfsdk:"prefix_variable"`
	Null0          types.Bool                         `tfsdk:"null0"`
	Null0Variable  types.String                       `tfsdk:"null0_variable"`
	VpnId          types.Int64                        `tfsdk:"vpn_id"`
	VpnIdVariable  types.String                       `tfsdk:"vpn_id_variable"`
	Nat            types.String                       `tfsdk:"nat"`
	NatVariable    types.String                       `tfsdk:"nat_variable"`
	NextHops       []CiscoVPNIpv6StaticRoutesNextHops `tfsdk:"next_hops"`
}

type CiscoVPNIpv4StaticGreRoutes struct {
	Optional          types.Bool   `tfsdk:"optional"`
	Prefix            types.String `tfsdk:"prefix"`
	PrefixVariable    types.String `tfsdk:"prefix_variable"`
	VpnId             types.Int64  `tfsdk:"vpn_id"`
	Interface         types.Set    `tfsdk:"interface"`
	InterfaceVariable types.String `tfsdk:"interface_variable"`
}

type CiscoVPNIpv4StaticIpsecRoutes struct {
	Optional          types.Bool   `tfsdk:"optional"`
	Prefix            types.String `tfsdk:"prefix"`
	PrefixVariable    types.String `tfsdk:"prefix_variable"`
	VpnId             types.Int64  `tfsdk:"vpn_id"`
	Interface         types.Set    `tfsdk:"interface"`
	InterfaceVariable types.String `tfsdk:"interface_variable"`
}

type CiscoVPNOmpAdvertiseIpv4Routes struct {
	Optional                types.Bool                               `tfsdk:"optional"`
	Protocol                types.String                             `tfsdk:"protocol"`
	ProtocolVariable        types.String                             `tfsdk:"protocol_variable"`
	RoutePolicy             types.String                             `tfsdk:"route_policy"`
	RoutePolicyVariable     types.String                             `tfsdk:"route_policy_variable"`
	ProtocolSubType         types.Set                                `tfsdk:"protocol_sub_type"`
	ProtocolSubTypeVariable types.String                             `tfsdk:"protocol_sub_type_variable"`
	Prefixes                []CiscoVPNOmpAdvertiseIpv4RoutesPrefixes `tfsdk:"prefixes"`
}

type CiscoVPNOmpAdvertiseIpv6Routes struct {
	Optional                types.Bool                               `tfsdk:"optional"`
	Protocol                types.String                             `tfsdk:"protocol"`
	ProtocolVariable        types.String                             `tfsdk:"protocol_variable"`
	RoutePolicy             types.String                             `tfsdk:"route_policy"`
	RoutePolicyVariable     types.String                             `tfsdk:"route_policy_variable"`
	ProtocolSubType         types.Set                                `tfsdk:"protocol_sub_type"`
	ProtocolSubTypeVariable types.String                             `tfsdk:"protocol_sub_type_variable"`
	Prefixes                []CiscoVPNOmpAdvertiseIpv6RoutesPrefixes `tfsdk:"prefixes"`
}

type CiscoVPNNat64Pools struct {
	Optional               types.Bool   `tfsdk:"optional"`
	Name                   types.String `tfsdk:"name"`
	StartAddress           types.String `tfsdk:"start_address"`
	StartAddressVariable   types.String `tfsdk:"start_address_variable"`
	EndAddress             types.String `tfsdk:"end_address"`
	EndAddressVariable     types.String `tfsdk:"end_address_variable"`
	Overload               types.Bool   `tfsdk:"overload"`
	OverloadVariable       types.String `tfsdk:"overload_variable"`
	LeakFromGlobal         types.Bool   `tfsdk:"leak_from_global"`
	LeakFromGlobalProtocol types.String `tfsdk:"leak_from_global_protocol"`
	LeakToGlobal           types.Bool   `tfsdk:"leak_to_global"`
}

type CiscoVPNNatPools struct {
	Optional             types.Bool   `tfsdk:"optional"`
	Name                 types.Int64  `tfsdk:"name"`
	NameVariable         types.String `tfsdk:"name_variable"`
	PrefixLength         types.Int64  `tfsdk:"prefix_length"`
	PrefixLengthVariable types.String `tfsdk:"prefix_length_variable"`
	RangeStart           types.String `tfsdk:"range_start"`
	RangeStartVariable   types.String `tfsdk:"range_start_variable"`
	RangeEnd             types.String `tfsdk:"range_end"`
	RangeEndVariable     types.String `tfsdk:"range_end_variable"`
	Overload             types.Bool   `tfsdk:"overload"`
	OverloadVariable     types.String `tfsdk:"overload_variable"`
	Direction            types.String `tfsdk:"direction"`
	DirectionVariable    types.String `tfsdk:"direction_variable"`
	TrackerId            types.Int64  `tfsdk:"tracker_id"`
	TrackerIdVariable    types.String `tfsdk:"tracker_id_variable"`
}

type CiscoVPNStaticNatRules struct {
	Optional                   types.Bool   `tfsdk:"optional"`
	PoolName                   types.Int64  `tfsdk:"pool_name"`
	PoolNameVariable           types.String `tfsdk:"pool_name_variable"`
	SourceIp                   types.String `tfsdk:"source_ip"`
	SourceIpVariable           types.String `tfsdk:"source_ip_variable"`
	TranslateIp                types.String `tfsdk:"translate_ip"`
	TranslateIpVariable        types.String `tfsdk:"translate_ip_variable"`
	StaticNatDirection         types.String `tfsdk:"static_nat_direction"`
	StaticNatDirectionVariable types.String `tfsdk:"static_nat_direction_variable"`
	TrackerId                  types.Int64  `tfsdk:"tracker_id"`
	TrackerIdVariable          types.String `tfsdk:"tracker_id_variable"`
}

type CiscoVPNStaticNatSubnetRules struct {
	Optional                   types.Bool   `tfsdk:"optional"`
	SourceIpSubnet             types.String `tfsdk:"source_ip_subnet"`
	SourceIpSubnetVariable     types.String `tfsdk:"source_ip_subnet_variable"`
	TranslateIpSubnet          types.String `tfsdk:"translate_ip_subnet"`
	TranslateIpSubnetVariable  types.String `tfsdk:"translate_ip_subnet_variable"`
	PrefixLength               types.Int64  `tfsdk:"prefix_length"`
	PrefixLengthVariable       types.String `tfsdk:"prefix_length_variable"`
	StaticNatDirection         types.String `tfsdk:"static_nat_direction"`
	StaticNatDirectionVariable types.String `tfsdk:"static_nat_direction_variable"`
	TrackerId                  types.Int64  `tfsdk:"tracker_id"`
	TrackerIdVariable          types.String `tfsdk:"tracker_id_variable"`
}

type CiscoVPNPortForwardRules struct {
	Optional              types.Bool   `tfsdk:"optional"`
	PoolName              types.Int64  `tfsdk:"pool_name"`
	PoolNameVariable      types.String `tfsdk:"pool_name_variable"`
	SourcePort            types.Int64  `tfsdk:"source_port"`
	SourcePortVariable    types.String `tfsdk:"source_port_variable"`
	TranslatePort         types.Int64  `tfsdk:"translate_port"`
	TranslatePortVariable types.String `tfsdk:"translate_port_variable"`
	SourceIp              types.String `tfsdk:"source_ip"`
	SourceIpVariable      types.String `tfsdk:"source_ip_variable"`
	TranslateIp           types.String `tfsdk:"translate_ip"`
	TranslateIpVariable   types.String `tfsdk:"translate_ip_variable"`
	Protocol              types.String `tfsdk:"protocol"`
	ProtocolVariable      types.String `tfsdk:"protocol_variable"`
}

type CiscoVPNRouteGlobalImports struct {
	Optional                types.Bool                                `tfsdk:"optional"`
	Protocol                types.String                              `tfsdk:"protocol"`
	ProtocolVariable        types.String                              `tfsdk:"protocol_variable"`
	ProtocolSubType         types.Set                                 `tfsdk:"protocol_sub_type"`
	ProtocolSubTypeVariable types.String                              `tfsdk:"protocol_sub_type_variable"`
	RoutePolicy             types.String                              `tfsdk:"route_policy"`
	Redistributes           []CiscoVPNRouteGlobalImportsRedistributes `tfsdk:"redistributes"`
}

type CiscoVPNRouteVpnImports struct {
	Optional                types.Bool                             `tfsdk:"optional"`
	SourceVpnId             types.Int64                            `tfsdk:"source_vpn_id"`
	SourceVpnIdVariable     types.String                           `tfsdk:"source_vpn_id_variable"`
	Protocol                types.String                           `tfsdk:"protocol"`
	ProtocolVariable        types.String                           `tfsdk:"protocol_variable"`
	ProtocolSubType         types.Set                              `tfsdk:"protocol_sub_type"`
	ProtocolSubTypeVariable types.String                           `tfsdk:"protocol_sub_type_variable"`
	RoutePolicy             types.String                           `tfsdk:"route_policy"`
	RoutePolicyVariable     types.String                           `tfsdk:"route_policy_variable"`
	Redistributes           []CiscoVPNRouteVpnImportsRedistributes `tfsdk:"redistributes"`
}

type CiscoVPNRouteGlobalExports struct {
	Optional                types.Bool                                `tfsdk:"optional"`
	Protocol                types.String                              `tfsdk:"protocol"`
	ProtocolVariable        types.String                              `tfsdk:"protocol_variable"`
	ProtocolSubType         types.Set                                 `tfsdk:"protocol_sub_type"`
	ProtocolSubTypeVariable types.String                              `tfsdk:"protocol_sub_type_variable"`
	RoutePolicy             types.String                              `tfsdk:"route_policy"`
	Redistributes           []CiscoVPNRouteGlobalExportsRedistributes `tfsdk:"redistributes"`
}

type CiscoVPNIpv4StaticRoutesNextHops struct {
	Optional         types.Bool   `tfsdk:"optional"`
	Address          types.String `tfsdk:"address"`
	AddressVariable  types.String `tfsdk:"address_variable"`
	Distance         types.Int64  `tfsdk:"distance"`
	DistanceVariable types.String `tfsdk:"distance_variable"`
}
type CiscoVPNIpv4StaticRoutesTrackNextHops struct {
	Optional         types.Bool   `tfsdk:"optional"`
	Address          types.String `tfsdk:"address"`
	AddressVariable  types.String `tfsdk:"address_variable"`
	Distance         types.Int64  `tfsdk:"distance"`
	DistanceVariable types.String `tfsdk:"distance_variable"`
	Tracker          types.String `tfsdk:"tracker"`
	TrackerVariable  types.String `tfsdk:"tracker_variable"`
}

type CiscoVPNIpv6StaticRoutesNextHops struct {
	Optional         types.Bool   `tfsdk:"optional"`
	Address          types.String `tfsdk:"address"`
	AddressVariable  types.String `tfsdk:"address_variable"`
	Distance         types.Int64  `tfsdk:"distance"`
	DistanceVariable types.String `tfsdk:"distance_variable"`
}

type CiscoVPNOmpAdvertiseIpv4RoutesPrefixes struct {
	Optional              types.Bool   `tfsdk:"optional"`
	PrefixEntry           types.String `tfsdk:"prefix_entry"`
	PrefixEntryVariable   types.String `tfsdk:"prefix_entry_variable"`
	AggregateOnly         types.Bool   `tfsdk:"aggregate_only"`
	AggregateOnlyVariable types.String `tfsdk:"aggregate_only_variable"`
}

type CiscoVPNOmpAdvertiseIpv6RoutesPrefixes struct {
	Optional              types.Bool   `tfsdk:"optional"`
	PrefixEntry           types.String `tfsdk:"prefix_entry"`
	PrefixEntryVariable   types.String `tfsdk:"prefix_entry_variable"`
	AggregateOnly         types.Bool   `tfsdk:"aggregate_only"`
	AggregateOnlyVariable types.String `tfsdk:"aggregate_only_variable"`
}

type CiscoVPNRouteGlobalImportsRedistributes struct {
	Optional         types.Bool   `tfsdk:"optional"`
	Protocol         types.String `tfsdk:"protocol"`
	ProtocolVariable types.String `tfsdk:"protocol_variable"`
	RoutePolicy      types.String `tfsdk:"route_policy"`
}

type CiscoVPNRouteVpnImportsRedistributes struct {
	Optional            types.Bool   `tfsdk:"optional"`
	Protocol            types.String `tfsdk:"protocol"`
	ProtocolVariable    types.String `tfsdk:"protocol_variable"`
	RoutePolicy         types.String `tfsdk:"route_policy"`
	RoutePolicyVariable types.String `tfsdk:"route_policy_variable"`
}

type CiscoVPNRouteGlobalExportsRedistributes struct {
	Optional         types.Bool   `tfsdk:"optional"`
	Protocol         types.String `tfsdk:"protocol"`
	ProtocolVariable types.String `tfsdk:"protocol_variable"`
	RoutePolicy      types.String `tfsdk:"route_policy"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data CiscoVPN) getModel() string {
	return "cisco_vpn"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data CiscoVPN) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "cisco_vpn")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."
	if data.VpnId.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"vpn-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"vpn-id."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"vpn-id."+"vipValue", data.VpnId.ValueInt64())
	}

	if !data.VpnNameVariable.IsNull() {
		body, _ = sjson.Set(body, path+"name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"name."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"name."+"vipVariableName", data.VpnNameVariable.ValueString())
	} else if data.VpnName.IsNull() {
		body, _ = sjson.Set(body, path+"name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"name."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"name."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"name."+"vipValue", data.VpnName.ValueString())
	}
	if data.TenantVpnId.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tenant-vpn-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tenant-vpn-id."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tenant-vpn-id."+"vipValue", data.TenantVpnId.ValueInt64())
	}
	if data.OrganizationName.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"org-name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"org-name."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"org-name."+"vipValue", data.OrganizationName.ValueString())
	}

	if !data.OmpAdminDistanceIpv4Variable.IsNull() {
		body, _ = sjson.Set(body, path+"omp-admin-distance-ipv4."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"omp-admin-distance-ipv4."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"omp-admin-distance-ipv4."+"vipVariableName", data.OmpAdminDistanceIpv4Variable.ValueString())
	} else if data.OmpAdminDistanceIpv4.IsNull() {
		body, _ = sjson.Set(body, path+"omp-admin-distance-ipv4."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"omp-admin-distance-ipv4."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"omp-admin-distance-ipv4."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"omp-admin-distance-ipv4."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"omp-admin-distance-ipv4."+"vipValue", data.OmpAdminDistanceIpv4.ValueInt64())
	}

	if !data.OmpAdminDistanceIpv6Variable.IsNull() {
		body, _ = sjson.Set(body, path+"omp-admin-distance-ipv6."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"omp-admin-distance-ipv6."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"omp-admin-distance-ipv6."+"vipVariableName", data.OmpAdminDistanceIpv6Variable.ValueString())
	} else if data.OmpAdminDistanceIpv6.IsNull() {
		body, _ = sjson.Set(body, path+"omp-admin-distance-ipv6."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"omp-admin-distance-ipv6."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"omp-admin-distance-ipv6."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"omp-admin-distance-ipv6."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"omp-admin-distance-ipv6."+"vipValue", data.OmpAdminDistanceIpv6.ValueInt64())
	}

	if !data.EnhanceEcmpKeyingVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ecmp-hash-key.layer4."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ecmp-hash-key.layer4."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ecmp-hash-key.layer4."+"vipVariableName", data.EnhanceEcmpKeyingVariable.ValueString())
	} else if data.EnhanceEcmpKeying.IsNull() {
		body, _ = sjson.Set(body, path+"ecmp-hash-key.layer4."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ecmp-hash-key.layer4."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ecmp-hash-key.layer4."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ecmp-hash-key.layer4."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ecmp-hash-key.layer4."+"vipValue", strconv.FormatBool(data.EnhanceEcmpKeying.ValueBool()))
	}
	if len(data.DnsIpv4Servers) > 0 {
		body, _ = sjson.Set(body, path+"dns."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"dns."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"dns."+"vipPrimaryKey", []string{"dns-addr"})
		body, _ = sjson.Set(body, path+"dns."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"dns."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"dns."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"dns."+"vipPrimaryKey", []string{"dns-addr"})
		body, _ = sjson.Set(body, path+"dns."+"vipValue", []interface{}{})
	}
	for _, item := range data.DnsIpv4Servers {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "dns-addr")

		if !item.AddressVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dns-addr."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dns-addr."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "dns-addr."+"vipVariableName", item.AddressVariable.ValueString())
		} else if item.Address.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "dns-addr."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dns-addr."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "dns-addr."+"vipValue", item.Address.ValueString())
		}
		itemAttributes = append(itemAttributes, "role")

		if !item.RoleVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "role."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "role."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "role."+"vipVariableName", item.RoleVariable.ValueString())
		} else if item.Role.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "role."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "role."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "role."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "role."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "role."+"vipValue", item.Role.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"dns."+"vipValue.-1", itemBody)
	}
	if len(data.DnsIpv6Servers) > 0 {
		body, _ = sjson.Set(body, path+"dns-ipv6."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"dns-ipv6."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"dns-ipv6."+"vipPrimaryKey", []string{"dns-addr"})
		body, _ = sjson.Set(body, path+"dns-ipv6."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"dns-ipv6."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"dns-ipv6."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"dns-ipv6."+"vipPrimaryKey", []string{"dns-addr"})
		body, _ = sjson.Set(body, path+"dns-ipv6."+"vipValue", []interface{}{})
	}
	for _, item := range data.DnsIpv6Servers {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "dns-addr")
		if item.Address.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "dns-addr."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dns-addr."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "dns-addr."+"vipValue", item.Address.ValueString())
		}
		itemAttributes = append(itemAttributes, "role")

		if !item.RoleVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "role."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "role."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "role."+"vipVariableName", item.RoleVariable.ValueString())
		} else if item.Role.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "role."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "role."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "role."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "role."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "role."+"vipValue", item.Role.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"dns-ipv6."+"vipValue.-1", itemBody)
	}
	if len(data.DnsHosts) > 0 {
		body, _ = sjson.Set(body, path+"host."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"host."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"host."+"vipPrimaryKey", []string{"hostname"})
		body, _ = sjson.Set(body, path+"host."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"host."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"host."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"host."+"vipPrimaryKey", []string{"hostname"})
		body, _ = sjson.Set(body, path+"host."+"vipValue", []interface{}{})
	}
	for _, item := range data.DnsHosts {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "hostname")

		if !item.HostnameVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "hostname."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "hostname."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "hostname."+"vipVariableName", item.HostnameVariable.ValueString())
		} else if item.Hostname.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "hostname."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "hostname."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "hostname."+"vipValue", item.Hostname.ValueString())
		}
		itemAttributes = append(itemAttributes, "ip")

		if !item.IpVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ip."+"vipObjectType", "list")
			itemBody, _ = sjson.Set(itemBody, "ip."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "ip."+"vipVariableName", item.IpVariable.ValueString())
		} else if item.Ip.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "ip."+"vipObjectType", "list")
			itemBody, _ = sjson.Set(itemBody, "ip."+"vipType", "constant")
			var values []string
			item.Ip.ElementsAs(ctx, &values, false)
			itemBody, _ = sjson.Set(itemBody, "ip."+"vipValue", values)
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"host."+"vipValue.-1", itemBody)
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
		if item.ServiceTypes.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "svc-type."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "svc-type."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "svc-type."+"vipValue", item.ServiceTypes.ValueString())
		}
		itemAttributes = append(itemAttributes, "address")

		if !item.AddressVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "address."+"vipObjectType", "list")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipVariableName", item.AddressVariable.ValueString())
		} else if item.Address.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "address."+"vipObjectType", "list")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipType", "constant")
			var values []string
			item.Address.ElementsAs(ctx, &values, false)
			itemBody, _ = sjson.Set(itemBody, "address."+"vipValue", values)
		}
		itemAttributes = append(itemAttributes, "interface")

		if !item.InterfaceVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipVariableName", item.InterfaceVariable.ValueString())
		} else if item.Interface.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipValue", item.Interface.ValueString())
		}
		itemAttributes = append(itemAttributes, "track-enable")

		if !item.TrackEnableVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "track-enable."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "track-enable."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "track-enable."+"vipVariableName", item.TrackEnableVariable.ValueString())
		} else if item.TrackEnable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "track-enable."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "track-enable."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "track-enable."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "track-enable."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "track-enable."+"vipValue", strconv.FormatBool(item.TrackEnable.ValueBool()))
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"service."+"vipValue.-1", itemBody)
	}
	if len(data.Ipv4StaticServiceRoutes) > 0 {
		body, _ = sjson.Set(body, path+"ip.service-route."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ip.service-route."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ip.service-route."+"vipPrimaryKey", []string{"prefix"})
		body, _ = sjson.Set(body, path+"ip.service-route."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"ip.service-route."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ip.service-route."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"ip.service-route."+"vipPrimaryKey", []string{"prefix"})
		body, _ = sjson.Set(body, path+"ip.service-route."+"vipValue", []interface{}{})
	}
	for _, item := range data.Ipv4StaticServiceRoutes {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "prefix")

		if !item.PrefixVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "prefix."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "prefix."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "prefix."+"vipVariableName", item.PrefixVariable.ValueString())
		} else if item.Prefix.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "prefix."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "prefix."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "prefix."+"vipValue", item.Prefix.ValueString())
		}
		itemAttributes = append(itemAttributes, "vpn")
		if item.VpnId.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipValue", item.VpnId.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "service")
		if item.Service.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "service."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "service."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "service."+"vipValue", item.Service.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"ip.service-route."+"vipValue.-1", itemBody)
	}
	if len(data.Ipv4StaticRoutes) > 0 {
		body, _ = sjson.Set(body, path+"ip.route."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ip.route."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ip.route."+"vipPrimaryKey", []string{"prefix"})
		body, _ = sjson.Set(body, path+"ip.route."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"ip.route."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ip.route."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"ip.route."+"vipPrimaryKey", []string{"prefix"})
		body, _ = sjson.Set(body, path+"ip.route."+"vipValue", []interface{}{})
	}
	for _, item := range data.Ipv4StaticRoutes {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "prefix")

		if !item.PrefixVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "prefix."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "prefix."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "prefix."+"vipVariableName", item.PrefixVariable.ValueString())
		} else if item.Prefix.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "prefix."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "prefix."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "prefix."+"vipValue", item.Prefix.ValueString())
		}
		itemAttributes = append(itemAttributes, "null0")

		if !item.Null0Variable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "null0."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "null0."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "null0."+"vipVariableName", item.Null0Variable.ValueString())
		} else if item.Null0.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "null0."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "null0."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "null0."+"vipValue", strconv.FormatBool(item.Null0.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "distance")

		if !item.DistanceVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "distance."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "distance."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "distance."+"vipVariableName", item.DistanceVariable.ValueString())
		} else if item.Distance.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "distance."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "distance."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "distance."+"vipValue", item.Distance.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "vpn")

		if !item.VpnIdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipVariableName", item.VpnIdVariable.ValueString())
		} else if item.VpnId.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipValue", item.VpnId.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "dhcp")

		if !item.DhcpVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dhcp."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "dhcp."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "dhcp."+"vipVariableName", item.DhcpVariable.ValueString())
		} else if item.Dhcp.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "dhcp."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "dhcp."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "dhcp."+"vipValue", strconv.FormatBool(item.Dhcp.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "next-hop")
		if len(item.NextHops) > 0 {
			itemBody, _ = sjson.Set(itemBody, "next-hop."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "next-hop."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "next-hop."+"vipPrimaryKey", []string{"address"})
			itemBody, _ = sjson.Set(itemBody, "next-hop."+"vipValue", []interface{}{})
		} else {
		}
		for _, childItem := range item.NextHops {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "address")

			if !childItem.AddressVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipVariableName", childItem.AddressVariable.ValueString())
			} else if childItem.Address.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipValue", childItem.Address.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "distance")

			if !childItem.DistanceVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "distance."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "distance."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "distance."+"vipVariableName", childItem.DistanceVariable.ValueString())
			} else if childItem.Distance.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "distance."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "distance."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "distance."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "distance."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "distance."+"vipValue", childItem.Distance.ValueInt64())
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "next-hop."+"vipValue.-1", itemChildBody)
		}
		itemAttributes = append(itemAttributes, "next-hop-with-track")
		if len(item.TrackNextHops) > 0 {
			itemBody, _ = sjson.Set(itemBody, "next-hop-with-track."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "next-hop-with-track."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "next-hop-with-track."+"vipPrimaryKey", []string{"address"})
			itemBody, _ = sjson.Set(itemBody, "next-hop-with-track."+"vipValue", []interface{}{})
		} else {
		}
		for _, childItem := range item.TrackNextHops {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "address")

			if !childItem.AddressVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipVariableName", childItem.AddressVariable.ValueString())
			} else if childItem.Address.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipValue", childItem.Address.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "distance")

			if !childItem.DistanceVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "distance."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "distance."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "distance."+"vipVariableName", childItem.DistanceVariable.ValueString())
			} else if childItem.Distance.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "distance."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "distance."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "distance."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "distance."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "distance."+"vipValue", childItem.Distance.ValueInt64())
			}
			itemChildAttributes = append(itemChildAttributes, "tracker")

			if !childItem.TrackerVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "tracker."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "tracker."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "tracker."+"vipVariableName", childItem.TrackerVariable.ValueString())
			} else if childItem.Tracker.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "tracker."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "tracker."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "tracker."+"vipValue", childItem.Tracker.ValueString())
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "next-hop-with-track."+"vipValue.-1", itemChildBody)
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"ip.route."+"vipValue.-1", itemBody)
	}
	if len(data.Ipv6StaticRoutes) > 0 {
		body, _ = sjson.Set(body, path+"ipv6.route."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ipv6.route."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ipv6.route."+"vipPrimaryKey", []string{"prefix"})
		body, _ = sjson.Set(body, path+"ipv6.route."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"ipv6.route."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ipv6.route."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"ipv6.route."+"vipPrimaryKey", []string{"prefix"})
		body, _ = sjson.Set(body, path+"ipv6.route."+"vipValue", []interface{}{})
	}
	for _, item := range data.Ipv6StaticRoutes {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "prefix")

		if !item.PrefixVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "prefix."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "prefix."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "prefix."+"vipVariableName", item.PrefixVariable.ValueString())
		} else if item.Prefix.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "prefix."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "prefix."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "prefix."+"vipValue", item.Prefix.ValueString())
		}
		itemAttributes = append(itemAttributes, "null0")

		if !item.Null0Variable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "null0."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "null0."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "null0."+"vipVariableName", item.Null0Variable.ValueString())
		} else if item.Null0.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "null0."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "null0."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "null0."+"vipValue", strconv.FormatBool(item.Null0.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "vpn")

		if !item.VpnIdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipVariableName", item.VpnIdVariable.ValueString())
		} else if item.VpnId.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipValue", item.VpnId.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "nat")

		if !item.NatVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "nat."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "nat."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "nat."+"vipVariableName", item.NatVariable.ValueString())
		} else if item.Nat.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "nat."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "nat."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "nat."+"vipValue", item.Nat.ValueString())
		}
		itemAttributes = append(itemAttributes, "next-hop")
		if len(item.NextHops) > 0 {
			itemBody, _ = sjson.Set(itemBody, "next-hop."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "next-hop."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "next-hop."+"vipPrimaryKey", []string{"address"})
			itemBody, _ = sjson.Set(itemBody, "next-hop."+"vipValue", []interface{}{})
		} else {
		}
		for _, childItem := range item.NextHops {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "address")

			if !childItem.AddressVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipVariableName", childItem.AddressVariable.ValueString())
			} else if childItem.Address.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipValue", childItem.Address.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "distance")

			if !childItem.DistanceVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "distance."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "distance."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "distance."+"vipVariableName", childItem.DistanceVariable.ValueString())
			} else if childItem.Distance.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "distance."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "distance."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "distance."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "distance."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "distance."+"vipValue", childItem.Distance.ValueInt64())
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "next-hop."+"vipValue.-1", itemChildBody)
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"ipv6.route."+"vipValue.-1", itemBody)
	}
	if len(data.Ipv4StaticGreRoutes) > 0 {
		body, _ = sjson.Set(body, path+"ip.gre-route."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ip.gre-route."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ip.gre-route."+"vipPrimaryKey", []string{"prefix"})
		body, _ = sjson.Set(body, path+"ip.gre-route."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"ip.gre-route."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ip.gre-route."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"ip.gre-route."+"vipPrimaryKey", []string{"prefix"})
		body, _ = sjson.Set(body, path+"ip.gre-route."+"vipValue", []interface{}{})
	}
	for _, item := range data.Ipv4StaticGreRoutes {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "prefix")

		if !item.PrefixVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "prefix."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "prefix."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "prefix."+"vipVariableName", item.PrefixVariable.ValueString())
		} else if item.Prefix.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "prefix."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "prefix."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "prefix."+"vipValue", item.Prefix.ValueString())
		}
		itemAttributes = append(itemAttributes, "vpn")
		if item.VpnId.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipValue", item.VpnId.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "interface")

		if !item.InterfaceVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipObjectType", "list")
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipVariableName", item.InterfaceVariable.ValueString())
		} else if item.Interface.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipObjectType", "list")
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipObjectType", "list")
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipType", "constant")
			var values []string
			item.Interface.ElementsAs(ctx, &values, false)
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipValue", values)
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"ip.gre-route."+"vipValue.-1", itemBody)
	}
	if len(data.Ipv4StaticIpsecRoutes) > 0 {
		body, _ = sjson.Set(body, path+"ip.ipsec-route."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ip.ipsec-route."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ip.ipsec-route."+"vipPrimaryKey", []string{"prefix"})
		body, _ = sjson.Set(body, path+"ip.ipsec-route."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"ip.ipsec-route."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ip.ipsec-route."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"ip.ipsec-route."+"vipPrimaryKey", []string{"prefix"})
		body, _ = sjson.Set(body, path+"ip.ipsec-route."+"vipValue", []interface{}{})
	}
	for _, item := range data.Ipv4StaticIpsecRoutes {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "prefix")

		if !item.PrefixVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "prefix."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "prefix."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "prefix."+"vipVariableName", item.PrefixVariable.ValueString())
		} else if item.Prefix.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "prefix."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "prefix."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "prefix."+"vipValue", item.Prefix.ValueString())
		}
		itemAttributes = append(itemAttributes, "vpn")
		if item.VpnId.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipValue", item.VpnId.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "interface")

		if !item.InterfaceVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipObjectType", "list")
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipVariableName", item.InterfaceVariable.ValueString())
		} else if item.Interface.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipObjectType", "list")
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipObjectType", "list")
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipType", "constant")
			var values []string
			item.Interface.ElementsAs(ctx, &values, false)
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipValue", values)
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"ip.ipsec-route."+"vipValue.-1", itemBody)
	}
	if len(data.OmpAdvertiseIpv4Routes) > 0 {
		body, _ = sjson.Set(body, path+"omp.advertise."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"omp.advertise."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"omp.advertise."+"vipPrimaryKey", []string{"protocol"})
		body, _ = sjson.Set(body, path+"omp.advertise."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"omp.advertise."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"omp.advertise."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"omp.advertise."+"vipPrimaryKey", []string{"protocol"})
		body, _ = sjson.Set(body, path+"omp.advertise."+"vipValue", []interface{}{})
	}
	for _, item := range data.OmpAdvertiseIpv4Routes {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "protocol")

		if !item.ProtocolVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipVariableName", item.ProtocolVariable.ValueString())
		} else if item.Protocol.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipValue", item.Protocol.ValueString())
		}
		itemAttributes = append(itemAttributes, "route-policy")

		if !item.RoutePolicyVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipVariableName", item.RoutePolicyVariable.ValueString())
		} else if item.RoutePolicy.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipValue", item.RoutePolicy.ValueString())
		}
		itemAttributes = append(itemAttributes, "protocol-sub-type")

		if !item.ProtocolSubTypeVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "protocol-sub-type."+"vipObjectType", "list")
			itemBody, _ = sjson.Set(itemBody, "protocol-sub-type."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "protocol-sub-type."+"vipVariableName", item.ProtocolSubTypeVariable.ValueString())
		} else if item.ProtocolSubType.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "protocol-sub-type."+"vipObjectType", "list")
			itemBody, _ = sjson.Set(itemBody, "protocol-sub-type."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "protocol-sub-type."+"vipObjectType", "list")
			itemBody, _ = sjson.Set(itemBody, "protocol-sub-type."+"vipType", "constant")
			var values []string
			item.ProtocolSubType.ElementsAs(ctx, &values, false)
			itemBody, _ = sjson.Set(itemBody, "protocol-sub-type."+"vipValue", values)
		}
		itemAttributes = append(itemAttributes, "prefix-list")
		if len(item.Prefixes) > 0 {
			itemBody, _ = sjson.Set(itemBody, "prefix-list."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "prefix-list."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "prefix-list."+"vipPrimaryKey", []string{"prefix-entry"})
			itemBody, _ = sjson.Set(itemBody, "prefix-list."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "prefix-list."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "prefix-list."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "prefix-list."+"vipPrimaryKey", []string{"prefix-entry"})
			itemBody, _ = sjson.Set(itemBody, "prefix-list."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.Prefixes {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "prefix-entry")

			if !childItem.PrefixEntryVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix-entry."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix-entry."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix-entry."+"vipVariableName", childItem.PrefixEntryVariable.ValueString())
			} else if childItem.PrefixEntry.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix-entry."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix-entry."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix-entry."+"vipValue", childItem.PrefixEntry.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "aggregate-only")

			if !childItem.AggregateOnlyVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "aggregate-only."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "aggregate-only."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "aggregate-only."+"vipVariableName", childItem.AggregateOnlyVariable.ValueString())
			} else if childItem.AggregateOnly.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "aggregate-only."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "aggregate-only."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "aggregate-only."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "aggregate-only."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "aggregate-only."+"vipValue", strconv.FormatBool(childItem.AggregateOnly.ValueBool()))
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "prefix-list."+"vipValue.-1", itemChildBody)
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"omp.advertise."+"vipValue.-1", itemBody)
	}
	if len(data.OmpAdvertiseIpv6Routes) > 0 {
		body, _ = sjson.Set(body, path+"omp.ipv6-advertise."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"omp.ipv6-advertise."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"omp.ipv6-advertise."+"vipPrimaryKey", []string{"protocol"})
		body, _ = sjson.Set(body, path+"omp.ipv6-advertise."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"omp.ipv6-advertise."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"omp.ipv6-advertise."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"omp.ipv6-advertise."+"vipPrimaryKey", []string{"protocol"})
		body, _ = sjson.Set(body, path+"omp.ipv6-advertise."+"vipValue", []interface{}{})
	}
	for _, item := range data.OmpAdvertiseIpv6Routes {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "protocol")

		if !item.ProtocolVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipVariableName", item.ProtocolVariable.ValueString())
		} else if item.Protocol.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipValue", item.Protocol.ValueString())
		}
		itemAttributes = append(itemAttributes, "route-policy")

		if !item.RoutePolicyVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipVariableName", item.RoutePolicyVariable.ValueString())
		} else if item.RoutePolicy.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipValue", item.RoutePolicy.ValueString())
		}
		itemAttributes = append(itemAttributes, "protocol-sub-type")

		if !item.ProtocolSubTypeVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "protocol-sub-type."+"vipObjectType", "list")
			itemBody, _ = sjson.Set(itemBody, "protocol-sub-type."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "protocol-sub-type."+"vipVariableName", item.ProtocolSubTypeVariable.ValueString())
		} else if item.ProtocolSubType.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "protocol-sub-type."+"vipObjectType", "list")
			itemBody, _ = sjson.Set(itemBody, "protocol-sub-type."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "protocol-sub-type."+"vipObjectType", "list")
			itemBody, _ = sjson.Set(itemBody, "protocol-sub-type."+"vipType", "constant")
			var values []string
			item.ProtocolSubType.ElementsAs(ctx, &values, false)
			itemBody, _ = sjson.Set(itemBody, "protocol-sub-type."+"vipValue", values)
		}
		itemAttributes = append(itemAttributes, "prefix-list")
		if len(item.Prefixes) > 0 {
			itemBody, _ = sjson.Set(itemBody, "prefix-list."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "prefix-list."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "prefix-list."+"vipPrimaryKey", []string{"prefix-entry"})
			itemBody, _ = sjson.Set(itemBody, "prefix-list."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "prefix-list."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "prefix-list."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "prefix-list."+"vipPrimaryKey", []string{"prefix-entry"})
			itemBody, _ = sjson.Set(itemBody, "prefix-list."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.Prefixes {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "prefix-entry")

			if !childItem.PrefixEntryVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix-entry."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix-entry."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix-entry."+"vipVariableName", childItem.PrefixEntryVariable.ValueString())
			} else if childItem.PrefixEntry.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix-entry."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix-entry."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix-entry."+"vipValue", childItem.PrefixEntry.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "aggregate-only")

			if !childItem.AggregateOnlyVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "aggregate-only."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "aggregate-only."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "aggregate-only."+"vipVariableName", childItem.AggregateOnlyVariable.ValueString())
			} else if childItem.AggregateOnly.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "aggregate-only."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "aggregate-only."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "aggregate-only."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "aggregate-only."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "aggregate-only."+"vipValue", strconv.FormatBool(childItem.AggregateOnly.ValueBool()))
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "prefix-list."+"vipValue.-1", itemChildBody)
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"omp.ipv6-advertise."+"vipValue.-1", itemBody)
	}
	if len(data.Nat64Pools) > 0 {
		body, _ = sjson.Set(body, path+"nat64.v4.pool."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"nat64.v4.pool."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"nat64.v4.pool."+"vipPrimaryKey", []string{"name"})
		body, _ = sjson.Set(body, path+"nat64.v4.pool."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"nat64.v4.pool."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"nat64.v4.pool."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"nat64.v4.pool."+"vipPrimaryKey", []string{"name"})
		body, _ = sjson.Set(body, path+"nat64.v4.pool."+"vipValue", []interface{}{})
	}
	for _, item := range data.Nat64Pools {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "name")
		if item.Name.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipValue", item.Name.ValueString())
		}
		itemAttributes = append(itemAttributes, "start-address")

		if !item.StartAddressVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "start-address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "start-address."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "start-address."+"vipVariableName", item.StartAddressVariable.ValueString())
		} else if item.StartAddress.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "start-address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "start-address."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "start-address."+"vipValue", item.StartAddress.ValueString())
		}
		itemAttributes = append(itemAttributes, "end-address")

		if !item.EndAddressVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "end-address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "end-address."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "end-address."+"vipVariableName", item.EndAddressVariable.ValueString())
		} else if item.EndAddress.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "end-address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "end-address."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "end-address."+"vipValue", item.EndAddress.ValueString())
		}
		itemAttributes = append(itemAttributes, "overload")

		if !item.OverloadVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "overload."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "overload."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "overload."+"vipVariableName", item.OverloadVariable.ValueString())
		} else if item.Overload.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "overload."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "overload."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "overload."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "overload."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "overload."+"vipValue", strconv.FormatBool(item.Overload.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "leak_from_global")
		if item.LeakFromGlobal.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "leak_from_global."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "leak_from_global."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "leak_from_global."+"vipValue", strconv.FormatBool(item.LeakFromGlobal.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "leak_from_global_protocol")
		if item.LeakFromGlobalProtocol.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "leak_from_global_protocol."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "leak_from_global_protocol."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "leak_from_global_protocol."+"vipValue", item.LeakFromGlobalProtocol.ValueString())
		}
		itemAttributes = append(itemAttributes, "leak_to_global")
		if item.LeakToGlobal.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "leak_to_global."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "leak_to_global."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "leak_to_global."+"vipValue", strconv.FormatBool(item.LeakToGlobal.ValueBool()))
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"nat64.v4.pool."+"vipValue.-1", itemBody)
	}
	if len(data.NatPools) > 0 {
		body, _ = sjson.Set(body, path+"nat.natpool."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"nat.natpool."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"nat.natpool."+"vipPrimaryKey", []string{"name"})
		body, _ = sjson.Set(body, path+"nat.natpool."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"nat.natpool."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"nat.natpool."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"nat.natpool."+"vipPrimaryKey", []string{"name"})
		body, _ = sjson.Set(body, path+"nat.natpool."+"vipValue", []interface{}{})
	}
	for _, item := range data.NatPools {
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
			itemBody, _ = sjson.Set(itemBody, "name."+"vipValue", item.Name.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "prefix-length")

		if !item.PrefixLengthVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "prefix-length."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "prefix-length."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "prefix-length."+"vipVariableName", item.PrefixLengthVariable.ValueString())
		} else if item.PrefixLength.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "prefix-length."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "prefix-length."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "prefix-length."+"vipValue", item.PrefixLength.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "range-start")

		if !item.RangeStartVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "range-start."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "range-start."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "range-start."+"vipVariableName", item.RangeStartVariable.ValueString())
		} else if item.RangeStart.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "range-start."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "range-start."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "range-start."+"vipValue", item.RangeStart.ValueString())
		}
		itemAttributes = append(itemAttributes, "range-end")

		if !item.RangeEndVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "range-end."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "range-end."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "range-end."+"vipVariableName", item.RangeEndVariable.ValueString())
		} else if item.RangeEnd.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "range-end."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "range-end."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "range-end."+"vipValue", item.RangeEnd.ValueString())
		}
		itemAttributes = append(itemAttributes, "overload")

		if !item.OverloadVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "overload."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "overload."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "overload."+"vipVariableName", item.OverloadVariable.ValueString())
		} else if item.Overload.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "overload."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "overload."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "overload."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "overload."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "overload."+"vipValue", strconv.FormatBool(item.Overload.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "direction")

		if !item.DirectionVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipVariableName", item.DirectionVariable.ValueString())
		} else if item.Direction.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipValue", item.Direction.ValueString())
		}
		itemAttributes = append(itemAttributes, "tracker-id")

		if !item.TrackerIdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tracker-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tracker-id."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "tracker-id."+"vipVariableName", item.TrackerIdVariable.ValueString())
		} else if item.TrackerId.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tracker-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tracker-id."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "tracker-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tracker-id."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "tracker-id."+"vipValue", item.TrackerId.ValueInt64())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"nat.natpool."+"vipValue.-1", itemBody)
	}
	if len(data.StaticNatRules) > 0 {
		body, _ = sjson.Set(body, path+"nat.static."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"nat.static."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"nat.static."+"vipPrimaryKey", []string{"source-ip", "translate-ip"})
		body, _ = sjson.Set(body, path+"nat.static."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"nat.static."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"nat.static."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"nat.static."+"vipPrimaryKey", []string{"source-ip", "translate-ip"})
		body, _ = sjson.Set(body, path+"nat.static."+"vipValue", []interface{}{})
	}
	for _, item := range data.StaticNatRules {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "pool-name")

		if !item.PoolNameVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "pool-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "pool-name."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "pool-name."+"vipVariableName", item.PoolNameVariable.ValueString())
		} else if item.PoolName.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "pool-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "pool-name."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "pool-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "pool-name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "pool-name."+"vipValue", item.PoolName.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "source-ip")

		if !item.SourceIpVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "source-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-ip."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "source-ip."+"vipVariableName", item.SourceIpVariable.ValueString())
		} else if item.SourceIp.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "source-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-ip."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "source-ip."+"vipValue", item.SourceIp.ValueString())
		}
		itemAttributes = append(itemAttributes, "translate-ip")

		if !item.TranslateIpVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "translate-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "translate-ip."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "translate-ip."+"vipVariableName", item.TranslateIpVariable.ValueString())
		} else if item.TranslateIp.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "translate-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "translate-ip."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "translate-ip."+"vipValue", item.TranslateIp.ValueString())
		}
		itemAttributes = append(itemAttributes, "static-nat-direction")

		if !item.StaticNatDirectionVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "static-nat-direction."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "static-nat-direction."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "static-nat-direction."+"vipVariableName", item.StaticNatDirectionVariable.ValueString())
		} else if item.StaticNatDirection.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "static-nat-direction."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "static-nat-direction."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "static-nat-direction."+"vipValue", item.StaticNatDirection.ValueString())
		}
		itemAttributes = append(itemAttributes, "tracker-id")

		if !item.TrackerIdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tracker-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tracker-id."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "tracker-id."+"vipVariableName", item.TrackerIdVariable.ValueString())
		} else if item.TrackerId.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tracker-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tracker-id."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "tracker-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tracker-id."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "tracker-id."+"vipValue", item.TrackerId.ValueInt64())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"nat.static."+"vipValue.-1", itemBody)
	}
	if len(data.StaticNatSubnetRules) > 0 {
		body, _ = sjson.Set(body, path+"nat.subnet-static."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"nat.subnet-static."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"nat.subnet-static."+"vipPrimaryKey", []string{"source-ip-subnet", "translate-ip-subnet"})
		body, _ = sjson.Set(body, path+"nat.subnet-static."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"nat.subnet-static."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"nat.subnet-static."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"nat.subnet-static."+"vipPrimaryKey", []string{"source-ip-subnet", "translate-ip-subnet"})
		body, _ = sjson.Set(body, path+"nat.subnet-static."+"vipValue", []interface{}{})
	}
	for _, item := range data.StaticNatSubnetRules {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "source-ip-subnet")

		if !item.SourceIpSubnetVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "source-ip-subnet."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-ip-subnet."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "source-ip-subnet."+"vipVariableName", item.SourceIpSubnetVariable.ValueString())
		} else if item.SourceIpSubnet.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "source-ip-subnet."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-ip-subnet."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "source-ip-subnet."+"vipValue", item.SourceIpSubnet.ValueString())
		}
		itemAttributes = append(itemAttributes, "translate-ip-subnet")

		if !item.TranslateIpSubnetVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "translate-ip-subnet."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "translate-ip-subnet."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "translate-ip-subnet."+"vipVariableName", item.TranslateIpSubnetVariable.ValueString())
		} else if item.TranslateIpSubnet.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "translate-ip-subnet."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "translate-ip-subnet."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "translate-ip-subnet."+"vipValue", item.TranslateIpSubnet.ValueString())
		}
		itemAttributes = append(itemAttributes, "prefix-length")

		if !item.PrefixLengthVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "prefix-length."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "prefix-length."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "prefix-length."+"vipVariableName", item.PrefixLengthVariable.ValueString())
		} else if item.PrefixLength.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "prefix-length."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "prefix-length."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "prefix-length."+"vipValue", item.PrefixLength.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "static-nat-direction")

		if !item.StaticNatDirectionVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "static-nat-direction."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "static-nat-direction."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "static-nat-direction."+"vipVariableName", item.StaticNatDirectionVariable.ValueString())
		} else if item.StaticNatDirection.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "static-nat-direction."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "static-nat-direction."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "static-nat-direction."+"vipValue", item.StaticNatDirection.ValueString())
		}
		itemAttributes = append(itemAttributes, "tracker-id")

		if !item.TrackerIdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tracker-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tracker-id."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "tracker-id."+"vipVariableName", item.TrackerIdVariable.ValueString())
		} else if item.TrackerId.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tracker-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tracker-id."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "tracker-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tracker-id."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "tracker-id."+"vipValue", item.TrackerId.ValueInt64())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"nat.subnet-static."+"vipValue.-1", itemBody)
	}
	if len(data.PortForwardRules) > 0 {
		body, _ = sjson.Set(body, path+"nat.port-forward."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"nat.port-forward."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"nat.port-forward."+"vipPrimaryKey", []string{"source-port", "translate-port", "source-ip", "translate-ip", "proto"})
		body, _ = sjson.Set(body, path+"nat.port-forward."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"nat.port-forward."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"nat.port-forward."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"nat.port-forward."+"vipPrimaryKey", []string{"source-port", "translate-port", "source-ip", "translate-ip", "proto"})
		body, _ = sjson.Set(body, path+"nat.port-forward."+"vipValue", []interface{}{})
	}
	for _, item := range data.PortForwardRules {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "pool-name")

		if !item.PoolNameVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "pool-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "pool-name."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "pool-name."+"vipVariableName", item.PoolNameVariable.ValueString())
		} else if item.PoolName.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "pool-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "pool-name."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "pool-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "pool-name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "pool-name."+"vipValue", item.PoolName.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "source-port")

		if !item.SourcePortVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "source-port."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-port."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "source-port."+"vipVariableName", item.SourcePortVariable.ValueString())
		} else if item.SourcePort.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "source-port."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-port."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "source-port."+"vipValue", item.SourcePort.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "translate-port")

		if !item.TranslatePortVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "translate-port."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "translate-port."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "translate-port."+"vipVariableName", item.TranslatePortVariable.ValueString())
		} else if item.TranslatePort.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "translate-port."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "translate-port."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "translate-port."+"vipValue", item.TranslatePort.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "source-ip")

		if !item.SourceIpVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "source-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-ip."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "source-ip."+"vipVariableName", item.SourceIpVariable.ValueString())
		} else if item.SourceIp.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "source-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-ip."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "source-ip."+"vipValue", item.SourceIp.ValueString())
		}
		itemAttributes = append(itemAttributes, "translate-ip")

		if !item.TranslateIpVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "translate-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "translate-ip."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "translate-ip."+"vipVariableName", item.TranslateIpVariable.ValueString())
		} else if item.TranslateIp.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "translate-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "translate-ip."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "translate-ip."+"vipValue", item.TranslateIp.ValueString())
		}
		itemAttributes = append(itemAttributes, "proto")

		if !item.ProtocolVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "proto."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "proto."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "proto."+"vipVariableName", item.ProtocolVariable.ValueString())
		} else if item.Protocol.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "proto."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "proto."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "proto."+"vipValue", item.Protocol.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"nat.port-forward."+"vipValue.-1", itemBody)
	}
	if len(data.RouteGlobalImports) > 0 {
		body, _ = sjson.Set(body, path+"route-import."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"route-import."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"route-import."+"vipPrimaryKey", []string{"protocol"})
		body, _ = sjson.Set(body, path+"route-import."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"route-import."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"route-import."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"route-import."+"vipPrimaryKey", []string{"protocol"})
		body, _ = sjson.Set(body, path+"route-import."+"vipValue", []interface{}{})
	}
	for _, item := range data.RouteGlobalImports {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "protocol")

		if !item.ProtocolVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipVariableName", item.ProtocolVariable.ValueString())
		} else if item.Protocol.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipValue", item.Protocol.ValueString())
		}
		itemAttributes = append(itemAttributes, "protocol-sub-type")

		if !item.ProtocolSubTypeVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "protocol-sub-type."+"vipObjectType", "list")
			itemBody, _ = sjson.Set(itemBody, "protocol-sub-type."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "protocol-sub-type."+"vipVariableName", item.ProtocolSubTypeVariable.ValueString())
		} else if item.ProtocolSubType.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "protocol-sub-type."+"vipObjectType", "list")
			itemBody, _ = sjson.Set(itemBody, "protocol-sub-type."+"vipType", "constant")
			var values []string
			item.ProtocolSubType.ElementsAs(ctx, &values, false)
			itemBody, _ = sjson.Set(itemBody, "protocol-sub-type."+"vipValue", values)
		}
		itemAttributes = append(itemAttributes, "route-policy")
		if item.RoutePolicy.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipValue", item.RoutePolicy.ValueString())
		}
		itemAttributes = append(itemAttributes, "redistribute")
		if len(item.Redistributes) > 0 {
			itemBody, _ = sjson.Set(itemBody, "redistribute."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "redistribute."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "redistribute."+"vipPrimaryKey", []string{"protocol"})
			itemBody, _ = sjson.Set(itemBody, "redistribute."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "redistribute."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "redistribute."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "redistribute."+"vipPrimaryKey", []string{"protocol"})
			itemBody, _ = sjson.Set(itemBody, "redistribute."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.Redistributes {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "protocol")

			if !childItem.ProtocolVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol."+"vipVariableName", childItem.ProtocolVariable.ValueString())
			} else if childItem.Protocol.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol."+"vipValue", childItem.Protocol.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "route-policy")
			if childItem.RoutePolicy.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipValue", childItem.RoutePolicy.ValueString())
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "redistribute."+"vipValue.-1", itemChildBody)
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"route-import."+"vipValue.-1", itemBody)
	}
	if len(data.RouteVpnImports) > 0 {
		body, _ = sjson.Set(body, path+"route-import-from."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"route-import-from."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"route-import-from."+"vipPrimaryKey", []string{"source-vpn", "protocol"})
		body, _ = sjson.Set(body, path+"route-import-from."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"route-import-from."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"route-import-from."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"route-import-from."+"vipPrimaryKey", []string{"source-vpn", "protocol"})
		body, _ = sjson.Set(body, path+"route-import-from."+"vipValue", []interface{}{})
	}
	for _, item := range data.RouteVpnImports {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "source-vpn")

		if !item.SourceVpnIdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "source-vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-vpn."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "source-vpn."+"vipVariableName", item.SourceVpnIdVariable.ValueString())
		} else if item.SourceVpnId.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "source-vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-vpn."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "source-vpn."+"vipValue", item.SourceVpnId.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "protocol")

		if !item.ProtocolVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipVariableName", item.ProtocolVariable.ValueString())
		} else if item.Protocol.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipValue", item.Protocol.ValueString())
		}
		itemAttributes = append(itemAttributes, "protocol-sub-type")

		if !item.ProtocolSubTypeVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "protocol-sub-type."+"vipObjectType", "list")
			itemBody, _ = sjson.Set(itemBody, "protocol-sub-type."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "protocol-sub-type."+"vipVariableName", item.ProtocolSubTypeVariable.ValueString())
		} else if item.ProtocolSubType.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "protocol-sub-type."+"vipObjectType", "list")
			itemBody, _ = sjson.Set(itemBody, "protocol-sub-type."+"vipType", "constant")
			var values []string
			item.ProtocolSubType.ElementsAs(ctx, &values, false)
			itemBody, _ = sjson.Set(itemBody, "protocol-sub-type."+"vipValue", values)
		}
		itemAttributes = append(itemAttributes, "route-policy")

		if !item.RoutePolicyVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipVariableName", item.RoutePolicyVariable.ValueString())
		} else if item.RoutePolicy.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipValue", item.RoutePolicy.ValueString())
		}
		itemAttributes = append(itemAttributes, "redistribute")
		if len(item.Redistributes) > 0 {
			itemBody, _ = sjson.Set(itemBody, "redistribute."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "redistribute."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "redistribute."+"vipPrimaryKey", []string{"protocol"})
			itemBody, _ = sjson.Set(itemBody, "redistribute."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "redistribute."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "redistribute."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "redistribute."+"vipPrimaryKey", []string{"protocol"})
			itemBody, _ = sjson.Set(itemBody, "redistribute."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.Redistributes {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "protocol")

			if !childItem.ProtocolVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol."+"vipVariableName", childItem.ProtocolVariable.ValueString())
			} else if childItem.Protocol.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol."+"vipValue", childItem.Protocol.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "route-policy")

			if !childItem.RoutePolicyVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipVariableName", childItem.RoutePolicyVariable.ValueString())
			} else if childItem.RoutePolicy.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipValue", childItem.RoutePolicy.ValueString())
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "redistribute."+"vipValue.-1", itemChildBody)
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"route-import-from."+"vipValue.-1", itemBody)
	}
	if len(data.RouteGlobalExports) > 0 {
		body, _ = sjson.Set(body, path+"route-export."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"route-export."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"route-export."+"vipPrimaryKey", []string{"protocol"})
		body, _ = sjson.Set(body, path+"route-export."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"route-export."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"route-export."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"route-export."+"vipPrimaryKey", []string{"protocol"})
		body, _ = sjson.Set(body, path+"route-export."+"vipValue", []interface{}{})
	}
	for _, item := range data.RouteGlobalExports {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "protocol")

		if !item.ProtocolVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipVariableName", item.ProtocolVariable.ValueString())
		} else if item.Protocol.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipValue", item.Protocol.ValueString())
		}
		itemAttributes = append(itemAttributes, "protocol-sub-type")

		if !item.ProtocolSubTypeVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "protocol-sub-type."+"vipObjectType", "list")
			itemBody, _ = sjson.Set(itemBody, "protocol-sub-type."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "protocol-sub-type."+"vipVariableName", item.ProtocolSubTypeVariable.ValueString())
		} else if item.ProtocolSubType.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "protocol-sub-type."+"vipObjectType", "list")
			itemBody, _ = sjson.Set(itemBody, "protocol-sub-type."+"vipType", "constant")
			var values []string
			item.ProtocolSubType.ElementsAs(ctx, &values, false)
			itemBody, _ = sjson.Set(itemBody, "protocol-sub-type."+"vipValue", values)
		}
		itemAttributes = append(itemAttributes, "route-policy")
		if item.RoutePolicy.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipValue", item.RoutePolicy.ValueString())
		}
		itemAttributes = append(itemAttributes, "redistribute")
		if len(item.Redistributes) > 0 {
			itemBody, _ = sjson.Set(itemBody, "redistribute."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "redistribute."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "redistribute."+"vipPrimaryKey", []string{"protocol"})
			itemBody, _ = sjson.Set(itemBody, "redistribute."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "redistribute."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "redistribute."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "redistribute."+"vipPrimaryKey", []string{"protocol"})
			itemBody, _ = sjson.Set(itemBody, "redistribute."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.Redistributes {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "protocol")

			if !childItem.ProtocolVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol."+"vipVariableName", childItem.ProtocolVariable.ValueString())
			} else if childItem.Protocol.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol."+"vipValue", childItem.Protocol.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "route-policy")
			if childItem.RoutePolicy.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipValue", childItem.RoutePolicy.ValueString())
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "redistribute."+"vipValue.-1", itemChildBody)
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"route-export."+"vipValue.-1", itemBody)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *CiscoVPN) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(path + "name.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.VpnName = types.StringNull()

			v := res.Get(path + "name.vipVariableName")
			data.VpnNameVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.VpnName = types.StringNull()
			data.VpnNameVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "name.vipValue")
			data.VpnName = types.StringValue(v.String())
			data.VpnNameVariable = types.StringNull()
		}
	} else {
		data.VpnName = types.StringNull()
		data.VpnNameVariable = types.StringNull()
	}
	if value := res.Get(path + "tenant-vpn-id.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TenantVpnId = types.Int64Null()

		} else if value.String() == "ignore" {
			data.TenantVpnId = types.Int64Null()

		} else if value.String() == "constant" {
			v := res.Get(path + "tenant-vpn-id.vipValue")
			data.TenantVpnId = types.Int64Value(v.Int())

		}
	} else {
		data.TenantVpnId = types.Int64Null()

	}
	if value := res.Get(path + "org-name.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.OrganizationName = types.StringNull()

		} else if value.String() == "ignore" {
			data.OrganizationName = types.StringNull()

		} else if value.String() == "constant" {
			v := res.Get(path + "org-name.vipValue")
			data.OrganizationName = types.StringValue(v.String())

		}
	} else {
		data.OrganizationName = types.StringNull()

	}
	if value := res.Get(path + "omp-admin-distance-ipv4.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.OmpAdminDistanceIpv4 = types.Int64Null()

			v := res.Get(path + "omp-admin-distance-ipv4.vipVariableName")
			data.OmpAdminDistanceIpv4Variable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.OmpAdminDistanceIpv4 = types.Int64Null()
			data.OmpAdminDistanceIpv4Variable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "omp-admin-distance-ipv4.vipValue")
			data.OmpAdminDistanceIpv4 = types.Int64Value(v.Int())
			data.OmpAdminDistanceIpv4Variable = types.StringNull()
		}
	} else {
		data.OmpAdminDistanceIpv4 = types.Int64Null()
		data.OmpAdminDistanceIpv4Variable = types.StringNull()
	}
	if value := res.Get(path + "omp-admin-distance-ipv6.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.OmpAdminDistanceIpv6 = types.Int64Null()

			v := res.Get(path + "omp-admin-distance-ipv6.vipVariableName")
			data.OmpAdminDistanceIpv6Variable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.OmpAdminDistanceIpv6 = types.Int64Null()
			data.OmpAdminDistanceIpv6Variable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "omp-admin-distance-ipv6.vipValue")
			data.OmpAdminDistanceIpv6 = types.Int64Value(v.Int())
			data.OmpAdminDistanceIpv6Variable = types.StringNull()
		}
	} else {
		data.OmpAdminDistanceIpv6 = types.Int64Null()
		data.OmpAdminDistanceIpv6Variable = types.StringNull()
	}
	if value := res.Get(path + "ecmp-hash-key.layer4.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.EnhanceEcmpKeying = types.BoolNull()

			v := res.Get(path + "ecmp-hash-key.layer4.vipVariableName")
			data.EnhanceEcmpKeyingVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.EnhanceEcmpKeying = types.BoolNull()
			data.EnhanceEcmpKeyingVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ecmp-hash-key.layer4.vipValue")
			data.EnhanceEcmpKeying = types.BoolValue(v.Bool())
			data.EnhanceEcmpKeyingVariable = types.StringNull()
		}
	} else {
		data.EnhanceEcmpKeying = types.BoolNull()
		data.EnhanceEcmpKeyingVariable = types.StringNull()
	}
	if value := res.Get(path + "dns.vipValue"); len(value.Array()) > 0 {
		data.DnsIpv4Servers = make([]CiscoVPNDnsIpv4Servers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoVPNDnsIpv4Servers{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("dns-addr.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Address = types.StringNull()

					cv := v.Get("dns-addr.vipVariableName")
					item.AddressVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Address = types.StringNull()
					item.AddressVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("dns-addr.vipValue")
					item.Address = types.StringValue(cv.String())
					item.AddressVariable = types.StringNull()
				}
			} else {
				item.Address = types.StringNull()
				item.AddressVariable = types.StringNull()
			}
			if cValue := v.Get("role.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Role = types.StringNull()

					cv := v.Get("role.vipVariableName")
					item.RoleVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Role = types.StringNull()
					item.RoleVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("role.vipValue")
					item.Role = types.StringValue(cv.String())
					item.RoleVariable = types.StringNull()
				}
			} else {
				item.Role = types.StringNull()
				item.RoleVariable = types.StringNull()
			}
			data.DnsIpv4Servers = append(data.DnsIpv4Servers, item)
			return true
		})
	} else {
		if len(data.DnsIpv4Servers) > 0 {
			data.DnsIpv4Servers = []CiscoVPNDnsIpv4Servers{}
		}
	}
	if value := res.Get(path + "dns-ipv6.vipValue"); len(value.Array()) > 0 {
		data.DnsIpv6Servers = make([]CiscoVPNDnsIpv6Servers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoVPNDnsIpv6Servers{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("dns-addr.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Address = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.Address = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("dns-addr.vipValue")
					item.Address = types.StringValue(cv.String())

				}
			} else {
				item.Address = types.StringNull()

			}
			if cValue := v.Get("role.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Role = types.StringNull()

					cv := v.Get("role.vipVariableName")
					item.RoleVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Role = types.StringNull()
					item.RoleVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("role.vipValue")
					item.Role = types.StringValue(cv.String())
					item.RoleVariable = types.StringNull()
				}
			} else {
				item.Role = types.StringNull()
				item.RoleVariable = types.StringNull()
			}
			data.DnsIpv6Servers = append(data.DnsIpv6Servers, item)
			return true
		})
	} else {
		if len(data.DnsIpv6Servers) > 0 {
			data.DnsIpv6Servers = []CiscoVPNDnsIpv6Servers{}
		}
	}
	if value := res.Get(path + "host.vipValue"); len(value.Array()) > 0 {
		data.DnsHosts = make([]CiscoVPNDnsHosts, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoVPNDnsHosts{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("hostname.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Hostname = types.StringNull()

					cv := v.Get("hostname.vipVariableName")
					item.HostnameVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Hostname = types.StringNull()
					item.HostnameVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("hostname.vipValue")
					item.Hostname = types.StringValue(cv.String())
					item.HostnameVariable = types.StringNull()
				}
			} else {
				item.Hostname = types.StringNull()
				item.HostnameVariable = types.StringNull()
			}
			if cValue := v.Get("ip.vipType"); len(cValue.Array()) > 0 {
				if cValue.String() == "variableName" {
					item.Ip = types.SetNull(types.StringType)

					cv := v.Get("ip.vipVariableName")
					item.IpVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Ip = types.SetNull(types.StringType)
					item.IpVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("ip.vipValue")
					item.Ip = helpers.GetStringSet(cv.Array())
					item.IpVariable = types.StringNull()
				}
			} else {
				item.Ip = types.SetNull(types.StringType)
				item.IpVariable = types.StringNull()
			}
			data.DnsHosts = append(data.DnsHosts, item)
			return true
		})
	} else {
		if len(data.DnsHosts) > 0 {
			data.DnsHosts = []CiscoVPNDnsHosts{}
		}
	}
	if value := res.Get(path + "service.vipValue"); len(value.Array()) > 0 {
		data.Services = make([]CiscoVPNServices, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoVPNServices{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("svc-type.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ServiceTypes = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.ServiceTypes = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("svc-type.vipValue")
					item.ServiceTypes = types.StringValue(cv.String())

				}
			} else {
				item.ServiceTypes = types.StringNull()

			}
			if cValue := v.Get("address.vipType"); len(cValue.Array()) > 0 {
				if cValue.String() == "variableName" {
					item.Address = types.SetNull(types.StringType)

					cv := v.Get("address.vipVariableName")
					item.AddressVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Address = types.SetNull(types.StringType)
					item.AddressVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("address.vipValue")
					item.Address = helpers.GetStringSet(cv.Array())
					item.AddressVariable = types.StringNull()
				}
			} else {
				item.Address = types.SetNull(types.StringType)
				item.AddressVariable = types.StringNull()
			}
			if cValue := v.Get("interface.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Interface = types.StringNull()

					cv := v.Get("interface.vipVariableName")
					item.InterfaceVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Interface = types.StringNull()
					item.InterfaceVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("interface.vipValue")
					item.Interface = types.StringValue(cv.String())
					item.InterfaceVariable = types.StringNull()
				}
			} else {
				item.Interface = types.StringNull()
				item.InterfaceVariable = types.StringNull()
			}
			if cValue := v.Get("track-enable.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TrackEnable = types.BoolNull()

					cv := v.Get("track-enable.vipVariableName")
					item.TrackEnableVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TrackEnable = types.BoolNull()
					item.TrackEnableVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("track-enable.vipValue")
					item.TrackEnable = types.BoolValue(cv.Bool())
					item.TrackEnableVariable = types.StringNull()
				}
			} else {
				item.TrackEnable = types.BoolNull()
				item.TrackEnableVariable = types.StringNull()
			}
			data.Services = append(data.Services, item)
			return true
		})
	} else {
		if len(data.Services) > 0 {
			data.Services = []CiscoVPNServices{}
		}
	}
	if value := res.Get(path + "ip.service-route.vipValue"); len(value.Array()) > 0 {
		data.Ipv4StaticServiceRoutes = make([]CiscoVPNIpv4StaticServiceRoutes, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoVPNIpv4StaticServiceRoutes{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("prefix.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Prefix = types.StringNull()

					cv := v.Get("prefix.vipVariableName")
					item.PrefixVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Prefix = types.StringNull()
					item.PrefixVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("prefix.vipValue")
					item.Prefix = types.StringValue(cv.String())
					item.PrefixVariable = types.StringNull()
				}
			} else {
				item.Prefix = types.StringNull()
				item.PrefixVariable = types.StringNull()
			}
			if cValue := v.Get("vpn.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.VpnId = types.Int64Null()

				} else if cValue.String() == "ignore" {
					item.VpnId = types.Int64Null()

				} else if cValue.String() == "constant" {
					cv := v.Get("vpn.vipValue")
					item.VpnId = types.Int64Value(cv.Int())

				}
			} else {
				item.VpnId = types.Int64Null()

			}
			if cValue := v.Get("service.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Service = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.Service = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("service.vipValue")
					item.Service = types.StringValue(cv.String())

				}
			} else {
				item.Service = types.StringNull()

			}
			data.Ipv4StaticServiceRoutes = append(data.Ipv4StaticServiceRoutes, item)
			return true
		})
	} else {
		if len(data.Ipv4StaticServiceRoutes) > 0 {
			data.Ipv4StaticServiceRoutes = []CiscoVPNIpv4StaticServiceRoutes{}
		}
	}
	if value := res.Get(path + "ip.route.vipValue"); len(value.Array()) > 0 {
		data.Ipv4StaticRoutes = make([]CiscoVPNIpv4StaticRoutes, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoVPNIpv4StaticRoutes{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("prefix.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Prefix = types.StringNull()

					cv := v.Get("prefix.vipVariableName")
					item.PrefixVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Prefix = types.StringNull()
					item.PrefixVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("prefix.vipValue")
					item.Prefix = types.StringValue(cv.String())
					item.PrefixVariable = types.StringNull()
				}
			} else {
				item.Prefix = types.StringNull()
				item.PrefixVariable = types.StringNull()
			}
			if cValue := v.Get("null0.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Null0 = types.BoolNull()

					cv := v.Get("null0.vipVariableName")
					item.Null0Variable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Null0 = types.BoolNull()
					item.Null0Variable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("null0.vipValue")
					item.Null0 = types.BoolValue(cv.Bool())
					item.Null0Variable = types.StringNull()
				}
			} else {
				item.Null0 = types.BoolNull()
				item.Null0Variable = types.StringNull()
			}
			if cValue := v.Get("distance.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Distance = types.Int64Null()

					cv := v.Get("distance.vipVariableName")
					item.DistanceVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Distance = types.Int64Null()
					item.DistanceVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("distance.vipValue")
					item.Distance = types.Int64Value(cv.Int())
					item.DistanceVariable = types.StringNull()
				}
			} else {
				item.Distance = types.Int64Null()
				item.DistanceVariable = types.StringNull()
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
			if cValue := v.Get("dhcp.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Dhcp = types.BoolNull()

					cv := v.Get("dhcp.vipVariableName")
					item.DhcpVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Dhcp = types.BoolNull()
					item.DhcpVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("dhcp.vipValue")
					item.Dhcp = types.BoolValue(cv.Bool())
					item.DhcpVariable = types.StringNull()
				}
			} else {
				item.Dhcp = types.BoolNull()
				item.DhcpVariable = types.StringNull()
			}
			if cValue := v.Get("next-hop.vipValue"); len(cValue.Array()) > 0 {
				item.NextHops = make([]CiscoVPNIpv4StaticRoutesNextHops, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CiscoVPNIpv4StaticRoutesNextHops{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("address.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Address = types.StringNull()

							ccv := cv.Get("address.vipVariableName")
							cItem.AddressVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Address = types.StringNull()
							cItem.AddressVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("address.vipValue")
							cItem.Address = types.StringValue(ccv.String())
							cItem.AddressVariable = types.StringNull()
						}
					} else {
						cItem.Address = types.StringNull()
						cItem.AddressVariable = types.StringNull()
					}
					if ccValue := cv.Get("distance.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Distance = types.Int64Null()

							ccv := cv.Get("distance.vipVariableName")
							cItem.DistanceVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Distance = types.Int64Null()
							cItem.DistanceVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("distance.vipValue")
							cItem.Distance = types.Int64Value(ccv.Int())
							cItem.DistanceVariable = types.StringNull()
						}
					} else {
						cItem.Distance = types.Int64Null()
						cItem.DistanceVariable = types.StringNull()
					}
					item.NextHops = append(item.NextHops, cItem)
					return true
				})
			} else {
				if len(item.NextHops) > 0 {
					item.NextHops = []CiscoVPNIpv4StaticRoutesNextHops{}
				}
			}
			if cValue := v.Get("next-hop-with-track.vipValue"); len(cValue.Array()) > 0 {
				item.TrackNextHops = make([]CiscoVPNIpv4StaticRoutesTrackNextHops, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CiscoVPNIpv4StaticRoutesTrackNextHops{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("address.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Address = types.StringNull()

							ccv := cv.Get("address.vipVariableName")
							cItem.AddressVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Address = types.StringNull()
							cItem.AddressVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("address.vipValue")
							cItem.Address = types.StringValue(ccv.String())
							cItem.AddressVariable = types.StringNull()
						}
					} else {
						cItem.Address = types.StringNull()
						cItem.AddressVariable = types.StringNull()
					}
					if ccValue := cv.Get("distance.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Distance = types.Int64Null()

							ccv := cv.Get("distance.vipVariableName")
							cItem.DistanceVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Distance = types.Int64Null()
							cItem.DistanceVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("distance.vipValue")
							cItem.Distance = types.Int64Value(ccv.Int())
							cItem.DistanceVariable = types.StringNull()
						}
					} else {
						cItem.Distance = types.Int64Null()
						cItem.DistanceVariable = types.StringNull()
					}
					if ccValue := cv.Get("tracker.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Tracker = types.StringNull()

							ccv := cv.Get("tracker.vipVariableName")
							cItem.TrackerVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Tracker = types.StringNull()
							cItem.TrackerVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("tracker.vipValue")
							cItem.Tracker = types.StringValue(ccv.String())
							cItem.TrackerVariable = types.StringNull()
						}
					} else {
						cItem.Tracker = types.StringNull()
						cItem.TrackerVariable = types.StringNull()
					}
					item.TrackNextHops = append(item.TrackNextHops, cItem)
					return true
				})
			} else {
				if len(item.TrackNextHops) > 0 {
					item.TrackNextHops = []CiscoVPNIpv4StaticRoutesTrackNextHops{}
				}
			}
			data.Ipv4StaticRoutes = append(data.Ipv4StaticRoutes, item)
			return true
		})
	} else {
		if len(data.Ipv4StaticRoutes) > 0 {
			data.Ipv4StaticRoutes = []CiscoVPNIpv4StaticRoutes{}
		}
	}
	if value := res.Get(path + "ipv6.route.vipValue"); len(value.Array()) > 0 {
		data.Ipv6StaticRoutes = make([]CiscoVPNIpv6StaticRoutes, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoVPNIpv6StaticRoutes{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("prefix.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Prefix = types.StringNull()

					cv := v.Get("prefix.vipVariableName")
					item.PrefixVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Prefix = types.StringNull()
					item.PrefixVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("prefix.vipValue")
					item.Prefix = types.StringValue(cv.String())
					item.PrefixVariable = types.StringNull()
				}
			} else {
				item.Prefix = types.StringNull()
				item.PrefixVariable = types.StringNull()
			}
			if cValue := v.Get("null0.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Null0 = types.BoolNull()

					cv := v.Get("null0.vipVariableName")
					item.Null0Variable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Null0 = types.BoolNull()
					item.Null0Variable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("null0.vipValue")
					item.Null0 = types.BoolValue(cv.Bool())
					item.Null0Variable = types.StringNull()
				}
			} else {
				item.Null0 = types.BoolNull()
				item.Null0Variable = types.StringNull()
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
			if cValue := v.Get("nat.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Nat = types.StringNull()

					cv := v.Get("nat.vipVariableName")
					item.NatVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Nat = types.StringNull()
					item.NatVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("nat.vipValue")
					item.Nat = types.StringValue(cv.String())
					item.NatVariable = types.StringNull()
				}
			} else {
				item.Nat = types.StringNull()
				item.NatVariable = types.StringNull()
			}
			if cValue := v.Get("next-hop.vipValue"); len(cValue.Array()) > 0 {
				item.NextHops = make([]CiscoVPNIpv6StaticRoutesNextHops, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CiscoVPNIpv6StaticRoutesNextHops{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("address.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Address = types.StringNull()

							ccv := cv.Get("address.vipVariableName")
							cItem.AddressVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Address = types.StringNull()
							cItem.AddressVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("address.vipValue")
							cItem.Address = types.StringValue(ccv.String())
							cItem.AddressVariable = types.StringNull()
						}
					} else {
						cItem.Address = types.StringNull()
						cItem.AddressVariable = types.StringNull()
					}
					if ccValue := cv.Get("distance.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Distance = types.Int64Null()

							ccv := cv.Get("distance.vipVariableName")
							cItem.DistanceVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Distance = types.Int64Null()
							cItem.DistanceVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("distance.vipValue")
							cItem.Distance = types.Int64Value(ccv.Int())
							cItem.DistanceVariable = types.StringNull()
						}
					} else {
						cItem.Distance = types.Int64Null()
						cItem.DistanceVariable = types.StringNull()
					}
					item.NextHops = append(item.NextHops, cItem)
					return true
				})
			} else {
				if len(item.NextHops) > 0 {
					item.NextHops = []CiscoVPNIpv6StaticRoutesNextHops{}
				}
			}
			data.Ipv6StaticRoutes = append(data.Ipv6StaticRoutes, item)
			return true
		})
	} else {
		if len(data.Ipv6StaticRoutes) > 0 {
			data.Ipv6StaticRoutes = []CiscoVPNIpv6StaticRoutes{}
		}
	}
	if value := res.Get(path + "ip.gre-route.vipValue"); len(value.Array()) > 0 {
		data.Ipv4StaticGreRoutes = make([]CiscoVPNIpv4StaticGreRoutes, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoVPNIpv4StaticGreRoutes{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("prefix.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Prefix = types.StringNull()

					cv := v.Get("prefix.vipVariableName")
					item.PrefixVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Prefix = types.StringNull()
					item.PrefixVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("prefix.vipValue")
					item.Prefix = types.StringValue(cv.String())
					item.PrefixVariable = types.StringNull()
				}
			} else {
				item.Prefix = types.StringNull()
				item.PrefixVariable = types.StringNull()
			}
			if cValue := v.Get("vpn.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.VpnId = types.Int64Null()

				} else if cValue.String() == "ignore" {
					item.VpnId = types.Int64Null()

				} else if cValue.String() == "constant" {
					cv := v.Get("vpn.vipValue")
					item.VpnId = types.Int64Value(cv.Int())

				}
			} else {
				item.VpnId = types.Int64Null()

			}
			if cValue := v.Get("interface.vipType"); len(cValue.Array()) > 0 {
				if cValue.String() == "variableName" {
					item.Interface = types.SetNull(types.StringType)

					cv := v.Get("interface.vipVariableName")
					item.InterfaceVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Interface = types.SetNull(types.StringType)
					item.InterfaceVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("interface.vipValue")
					item.Interface = helpers.GetStringSet(cv.Array())
					item.InterfaceVariable = types.StringNull()
				}
			} else {
				item.Interface = types.SetNull(types.StringType)
				item.InterfaceVariable = types.StringNull()
			}
			data.Ipv4StaticGreRoutes = append(data.Ipv4StaticGreRoutes, item)
			return true
		})
	} else {
		if len(data.Ipv4StaticGreRoutes) > 0 {
			data.Ipv4StaticGreRoutes = []CiscoVPNIpv4StaticGreRoutes{}
		}
	}
	if value := res.Get(path + "ip.ipsec-route.vipValue"); len(value.Array()) > 0 {
		data.Ipv4StaticIpsecRoutes = make([]CiscoVPNIpv4StaticIpsecRoutes, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoVPNIpv4StaticIpsecRoutes{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("prefix.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Prefix = types.StringNull()

					cv := v.Get("prefix.vipVariableName")
					item.PrefixVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Prefix = types.StringNull()
					item.PrefixVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("prefix.vipValue")
					item.Prefix = types.StringValue(cv.String())
					item.PrefixVariable = types.StringNull()
				}
			} else {
				item.Prefix = types.StringNull()
				item.PrefixVariable = types.StringNull()
			}
			if cValue := v.Get("vpn.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.VpnId = types.Int64Null()

				} else if cValue.String() == "ignore" {
					item.VpnId = types.Int64Null()

				} else if cValue.String() == "constant" {
					cv := v.Get("vpn.vipValue")
					item.VpnId = types.Int64Value(cv.Int())

				}
			} else {
				item.VpnId = types.Int64Null()

			}
			if cValue := v.Get("interface.vipType"); len(cValue.Array()) > 0 {
				if cValue.String() == "variableName" {
					item.Interface = types.SetNull(types.StringType)

					cv := v.Get("interface.vipVariableName")
					item.InterfaceVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Interface = types.SetNull(types.StringType)
					item.InterfaceVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("interface.vipValue")
					item.Interface = helpers.GetStringSet(cv.Array())
					item.InterfaceVariable = types.StringNull()
				}
			} else {
				item.Interface = types.SetNull(types.StringType)
				item.InterfaceVariable = types.StringNull()
			}
			data.Ipv4StaticIpsecRoutes = append(data.Ipv4StaticIpsecRoutes, item)
			return true
		})
	} else {
		if len(data.Ipv4StaticIpsecRoutes) > 0 {
			data.Ipv4StaticIpsecRoutes = []CiscoVPNIpv4StaticIpsecRoutes{}
		}
	}
	if value := res.Get(path + "omp.advertise.vipValue"); len(value.Array()) > 0 {
		data.OmpAdvertiseIpv4Routes = make([]CiscoVPNOmpAdvertiseIpv4Routes, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoVPNOmpAdvertiseIpv4Routes{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("protocol.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Protocol = types.StringNull()

					cv := v.Get("protocol.vipVariableName")
					item.ProtocolVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Protocol = types.StringNull()
					item.ProtocolVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("protocol.vipValue")
					item.Protocol = types.StringValue(cv.String())
					item.ProtocolVariable = types.StringNull()
				}
			} else {
				item.Protocol = types.StringNull()
				item.ProtocolVariable = types.StringNull()
			}
			if cValue := v.Get("route-policy.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.RoutePolicy = types.StringNull()

					cv := v.Get("route-policy.vipVariableName")
					item.RoutePolicyVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.RoutePolicy = types.StringNull()
					item.RoutePolicyVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("route-policy.vipValue")
					item.RoutePolicy = types.StringValue(cv.String())
					item.RoutePolicyVariable = types.StringNull()
				}
			} else {
				item.RoutePolicy = types.StringNull()
				item.RoutePolicyVariable = types.StringNull()
			}
			if cValue := v.Get("protocol-sub-type.vipType"); len(cValue.Array()) > 0 {
				if cValue.String() == "variableName" {
					item.ProtocolSubType = types.SetNull(types.StringType)

					cv := v.Get("protocol-sub-type.vipVariableName")
					item.ProtocolSubTypeVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.ProtocolSubType = types.SetNull(types.StringType)
					item.ProtocolSubTypeVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("protocol-sub-type.vipValue")
					item.ProtocolSubType = helpers.GetStringSet(cv.Array())
					item.ProtocolSubTypeVariable = types.StringNull()
				}
			} else {
				item.ProtocolSubType = types.SetNull(types.StringType)
				item.ProtocolSubTypeVariable = types.StringNull()
			}
			if cValue := v.Get("prefix-list.vipValue"); len(cValue.Array()) > 0 {
				item.Prefixes = make([]CiscoVPNOmpAdvertiseIpv4RoutesPrefixes, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CiscoVPNOmpAdvertiseIpv4RoutesPrefixes{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("prefix-entry.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.PrefixEntry = types.StringNull()

							ccv := cv.Get("prefix-entry.vipVariableName")
							cItem.PrefixEntryVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.PrefixEntry = types.StringNull()
							cItem.PrefixEntryVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("prefix-entry.vipValue")
							cItem.PrefixEntry = types.StringValue(ccv.String())
							cItem.PrefixEntryVariable = types.StringNull()
						}
					} else {
						cItem.PrefixEntry = types.StringNull()
						cItem.PrefixEntryVariable = types.StringNull()
					}
					if ccValue := cv.Get("aggregate-only.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.AggregateOnly = types.BoolNull()

							ccv := cv.Get("aggregate-only.vipVariableName")
							cItem.AggregateOnlyVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.AggregateOnly = types.BoolNull()
							cItem.AggregateOnlyVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("aggregate-only.vipValue")
							cItem.AggregateOnly = types.BoolValue(ccv.Bool())
							cItem.AggregateOnlyVariable = types.StringNull()
						}
					} else {
						cItem.AggregateOnly = types.BoolNull()
						cItem.AggregateOnlyVariable = types.StringNull()
					}
					item.Prefixes = append(item.Prefixes, cItem)
					return true
				})
			} else {
				if len(item.Prefixes) > 0 {
					item.Prefixes = []CiscoVPNOmpAdvertiseIpv4RoutesPrefixes{}
				}
			}
			data.OmpAdvertiseIpv4Routes = append(data.OmpAdvertiseIpv4Routes, item)
			return true
		})
	} else {
		if len(data.OmpAdvertiseIpv4Routes) > 0 {
			data.OmpAdvertiseIpv4Routes = []CiscoVPNOmpAdvertiseIpv4Routes{}
		}
	}
	if value := res.Get(path + "omp.ipv6-advertise.vipValue"); len(value.Array()) > 0 {
		data.OmpAdvertiseIpv6Routes = make([]CiscoVPNOmpAdvertiseIpv6Routes, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoVPNOmpAdvertiseIpv6Routes{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("protocol.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Protocol = types.StringNull()

					cv := v.Get("protocol.vipVariableName")
					item.ProtocolVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Protocol = types.StringNull()
					item.ProtocolVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("protocol.vipValue")
					item.Protocol = types.StringValue(cv.String())
					item.ProtocolVariable = types.StringNull()
				}
			} else {
				item.Protocol = types.StringNull()
				item.ProtocolVariable = types.StringNull()
			}
			if cValue := v.Get("route-policy.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.RoutePolicy = types.StringNull()

					cv := v.Get("route-policy.vipVariableName")
					item.RoutePolicyVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.RoutePolicy = types.StringNull()
					item.RoutePolicyVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("route-policy.vipValue")
					item.RoutePolicy = types.StringValue(cv.String())
					item.RoutePolicyVariable = types.StringNull()
				}
			} else {
				item.RoutePolicy = types.StringNull()
				item.RoutePolicyVariable = types.StringNull()
			}
			if cValue := v.Get("protocol-sub-type.vipType"); len(cValue.Array()) > 0 {
				if cValue.String() == "variableName" {
					item.ProtocolSubType = types.SetNull(types.StringType)

					cv := v.Get("protocol-sub-type.vipVariableName")
					item.ProtocolSubTypeVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.ProtocolSubType = types.SetNull(types.StringType)
					item.ProtocolSubTypeVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("protocol-sub-type.vipValue")
					item.ProtocolSubType = helpers.GetStringSet(cv.Array())
					item.ProtocolSubTypeVariable = types.StringNull()
				}
			} else {
				item.ProtocolSubType = types.SetNull(types.StringType)
				item.ProtocolSubTypeVariable = types.StringNull()
			}
			if cValue := v.Get("prefix-list.vipValue"); len(cValue.Array()) > 0 {
				item.Prefixes = make([]CiscoVPNOmpAdvertiseIpv6RoutesPrefixes, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CiscoVPNOmpAdvertiseIpv6RoutesPrefixes{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("prefix-entry.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.PrefixEntry = types.StringNull()

							ccv := cv.Get("prefix-entry.vipVariableName")
							cItem.PrefixEntryVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.PrefixEntry = types.StringNull()
							cItem.PrefixEntryVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("prefix-entry.vipValue")
							cItem.PrefixEntry = types.StringValue(ccv.String())
							cItem.PrefixEntryVariable = types.StringNull()
						}
					} else {
						cItem.PrefixEntry = types.StringNull()
						cItem.PrefixEntryVariable = types.StringNull()
					}
					if ccValue := cv.Get("aggregate-only.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.AggregateOnly = types.BoolNull()

							ccv := cv.Get("aggregate-only.vipVariableName")
							cItem.AggregateOnlyVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.AggregateOnly = types.BoolNull()
							cItem.AggregateOnlyVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("aggregate-only.vipValue")
							cItem.AggregateOnly = types.BoolValue(ccv.Bool())
							cItem.AggregateOnlyVariable = types.StringNull()
						}
					} else {
						cItem.AggregateOnly = types.BoolNull()
						cItem.AggregateOnlyVariable = types.StringNull()
					}
					item.Prefixes = append(item.Prefixes, cItem)
					return true
				})
			} else {
				if len(item.Prefixes) > 0 {
					item.Prefixes = []CiscoVPNOmpAdvertiseIpv6RoutesPrefixes{}
				}
			}
			data.OmpAdvertiseIpv6Routes = append(data.OmpAdvertiseIpv6Routes, item)
			return true
		})
	} else {
		if len(data.OmpAdvertiseIpv6Routes) > 0 {
			data.OmpAdvertiseIpv6Routes = []CiscoVPNOmpAdvertiseIpv6Routes{}
		}
	}
	if value := res.Get(path + "nat64.v4.pool.vipValue"); len(value.Array()) > 0 {
		data.Nat64Pools = make([]CiscoVPNNat64Pools, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoVPNNat64Pools{}
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
			if cValue := v.Get("start-address.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.StartAddress = types.StringNull()

					cv := v.Get("start-address.vipVariableName")
					item.StartAddressVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.StartAddress = types.StringNull()
					item.StartAddressVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("start-address.vipValue")
					item.StartAddress = types.StringValue(cv.String())
					item.StartAddressVariable = types.StringNull()
				}
			} else {
				item.StartAddress = types.StringNull()
				item.StartAddressVariable = types.StringNull()
			}
			if cValue := v.Get("end-address.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.EndAddress = types.StringNull()

					cv := v.Get("end-address.vipVariableName")
					item.EndAddressVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.EndAddress = types.StringNull()
					item.EndAddressVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("end-address.vipValue")
					item.EndAddress = types.StringValue(cv.String())
					item.EndAddressVariable = types.StringNull()
				}
			} else {
				item.EndAddress = types.StringNull()
				item.EndAddressVariable = types.StringNull()
			}
			if cValue := v.Get("overload.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Overload = types.BoolNull()

					cv := v.Get("overload.vipVariableName")
					item.OverloadVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Overload = types.BoolNull()
					item.OverloadVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("overload.vipValue")
					item.Overload = types.BoolValue(cv.Bool())
					item.OverloadVariable = types.StringNull()
				}
			} else {
				item.Overload = types.BoolNull()
				item.OverloadVariable = types.StringNull()
			}
			if cValue := v.Get("leak_from_global.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.LeakFromGlobal = types.BoolNull()

				} else if cValue.String() == "ignore" {
					item.LeakFromGlobal = types.BoolNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("leak_from_global.vipValue")
					item.LeakFromGlobal = types.BoolValue(cv.Bool())

				}
			} else {
				item.LeakFromGlobal = types.BoolNull()

			}
			if cValue := v.Get("leak_from_global_protocol.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.LeakFromGlobalProtocol = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.LeakFromGlobalProtocol = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("leak_from_global_protocol.vipValue")
					item.LeakFromGlobalProtocol = types.StringValue(cv.String())

				}
			} else {
				item.LeakFromGlobalProtocol = types.StringNull()

			}
			if cValue := v.Get("leak_to_global.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.LeakToGlobal = types.BoolNull()

				} else if cValue.String() == "ignore" {
					item.LeakToGlobal = types.BoolNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("leak_to_global.vipValue")
					item.LeakToGlobal = types.BoolValue(cv.Bool())

				}
			} else {
				item.LeakToGlobal = types.BoolNull()

			}
			data.Nat64Pools = append(data.Nat64Pools, item)
			return true
		})
	} else {
		if len(data.Nat64Pools) > 0 {
			data.Nat64Pools = []CiscoVPNNat64Pools{}
		}
	}
	if value := res.Get(path + "nat.natpool.vipValue"); len(value.Array()) > 0 {
		data.NatPools = make([]CiscoVPNNatPools, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoVPNNatPools{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Name = types.Int64Null()

					cv := v.Get("name.vipVariableName")
					item.NameVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Name = types.Int64Null()
					item.NameVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("name.vipValue")
					item.Name = types.Int64Value(cv.Int())
					item.NameVariable = types.StringNull()
				}
			} else {
				item.Name = types.Int64Null()
				item.NameVariable = types.StringNull()
			}
			if cValue := v.Get("prefix-length.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.PrefixLength = types.Int64Null()

					cv := v.Get("prefix-length.vipVariableName")
					item.PrefixLengthVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.PrefixLength = types.Int64Null()
					item.PrefixLengthVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("prefix-length.vipValue")
					item.PrefixLength = types.Int64Value(cv.Int())
					item.PrefixLengthVariable = types.StringNull()
				}
			} else {
				item.PrefixLength = types.Int64Null()
				item.PrefixLengthVariable = types.StringNull()
			}
			if cValue := v.Get("range-start.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.RangeStart = types.StringNull()

					cv := v.Get("range-start.vipVariableName")
					item.RangeStartVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.RangeStart = types.StringNull()
					item.RangeStartVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("range-start.vipValue")
					item.RangeStart = types.StringValue(cv.String())
					item.RangeStartVariable = types.StringNull()
				}
			} else {
				item.RangeStart = types.StringNull()
				item.RangeStartVariable = types.StringNull()
			}
			if cValue := v.Get("range-end.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.RangeEnd = types.StringNull()

					cv := v.Get("range-end.vipVariableName")
					item.RangeEndVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.RangeEnd = types.StringNull()
					item.RangeEndVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("range-end.vipValue")
					item.RangeEnd = types.StringValue(cv.String())
					item.RangeEndVariable = types.StringNull()
				}
			} else {
				item.RangeEnd = types.StringNull()
				item.RangeEndVariable = types.StringNull()
			}
			if cValue := v.Get("overload.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Overload = types.BoolNull()

					cv := v.Get("overload.vipVariableName")
					item.OverloadVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Overload = types.BoolNull()
					item.OverloadVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("overload.vipValue")
					item.Overload = types.BoolValue(cv.Bool())
					item.OverloadVariable = types.StringNull()
				}
			} else {
				item.Overload = types.BoolNull()
				item.OverloadVariable = types.StringNull()
			}
			if cValue := v.Get("direction.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Direction = types.StringNull()

					cv := v.Get("direction.vipVariableName")
					item.DirectionVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Direction = types.StringNull()
					item.DirectionVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("direction.vipValue")
					item.Direction = types.StringValue(cv.String())
					item.DirectionVariable = types.StringNull()
				}
			} else {
				item.Direction = types.StringNull()
				item.DirectionVariable = types.StringNull()
			}
			if cValue := v.Get("tracker-id.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TrackerId = types.Int64Null()

					cv := v.Get("tracker-id.vipVariableName")
					item.TrackerIdVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TrackerId = types.Int64Null()
					item.TrackerIdVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("tracker-id.vipValue")
					item.TrackerId = types.Int64Value(cv.Int())
					item.TrackerIdVariable = types.StringNull()
				}
			} else {
				item.TrackerId = types.Int64Null()
				item.TrackerIdVariable = types.StringNull()
			}
			data.NatPools = append(data.NatPools, item)
			return true
		})
	} else {
		if len(data.NatPools) > 0 {
			data.NatPools = []CiscoVPNNatPools{}
		}
	}
	if value := res.Get(path + "nat.static.vipValue"); len(value.Array()) > 0 {
		data.StaticNatRules = make([]CiscoVPNStaticNatRules, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoVPNStaticNatRules{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("pool-name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.PoolName = types.Int64Null()

					cv := v.Get("pool-name.vipVariableName")
					item.PoolNameVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.PoolName = types.Int64Null()
					item.PoolNameVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("pool-name.vipValue")
					item.PoolName = types.Int64Value(cv.Int())
					item.PoolNameVariable = types.StringNull()
				}
			} else {
				item.PoolName = types.Int64Null()
				item.PoolNameVariable = types.StringNull()
			}
			if cValue := v.Get("source-ip.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SourceIp = types.StringNull()

					cv := v.Get("source-ip.vipVariableName")
					item.SourceIpVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SourceIp = types.StringNull()
					item.SourceIpVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("source-ip.vipValue")
					item.SourceIp = types.StringValue(cv.String())
					item.SourceIpVariable = types.StringNull()
				}
			} else {
				item.SourceIp = types.StringNull()
				item.SourceIpVariable = types.StringNull()
			}
			if cValue := v.Get("translate-ip.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TranslateIp = types.StringNull()

					cv := v.Get("translate-ip.vipVariableName")
					item.TranslateIpVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TranslateIp = types.StringNull()
					item.TranslateIpVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("translate-ip.vipValue")
					item.TranslateIp = types.StringValue(cv.String())
					item.TranslateIpVariable = types.StringNull()
				}
			} else {
				item.TranslateIp = types.StringNull()
				item.TranslateIpVariable = types.StringNull()
			}
			if cValue := v.Get("static-nat-direction.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.StaticNatDirection = types.StringNull()

					cv := v.Get("static-nat-direction.vipVariableName")
					item.StaticNatDirectionVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.StaticNatDirection = types.StringNull()
					item.StaticNatDirectionVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("static-nat-direction.vipValue")
					item.StaticNatDirection = types.StringValue(cv.String())
					item.StaticNatDirectionVariable = types.StringNull()
				}
			} else {
				item.StaticNatDirection = types.StringNull()
				item.StaticNatDirectionVariable = types.StringNull()
			}
			if cValue := v.Get("tracker-id.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TrackerId = types.Int64Null()

					cv := v.Get("tracker-id.vipVariableName")
					item.TrackerIdVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TrackerId = types.Int64Null()
					item.TrackerIdVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("tracker-id.vipValue")
					item.TrackerId = types.Int64Value(cv.Int())
					item.TrackerIdVariable = types.StringNull()
				}
			} else {
				item.TrackerId = types.Int64Null()
				item.TrackerIdVariable = types.StringNull()
			}
			data.StaticNatRules = append(data.StaticNatRules, item)
			return true
		})
	} else {
		if len(data.StaticNatRules) > 0 {
			data.StaticNatRules = []CiscoVPNStaticNatRules{}
		}
	}
	if value := res.Get(path + "nat.subnet-static.vipValue"); len(value.Array()) > 0 {
		data.StaticNatSubnetRules = make([]CiscoVPNStaticNatSubnetRules, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoVPNStaticNatSubnetRules{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("source-ip-subnet.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SourceIpSubnet = types.StringNull()

					cv := v.Get("source-ip-subnet.vipVariableName")
					item.SourceIpSubnetVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SourceIpSubnet = types.StringNull()
					item.SourceIpSubnetVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("source-ip-subnet.vipValue")
					item.SourceIpSubnet = types.StringValue(cv.String())
					item.SourceIpSubnetVariable = types.StringNull()
				}
			} else {
				item.SourceIpSubnet = types.StringNull()
				item.SourceIpSubnetVariable = types.StringNull()
			}
			if cValue := v.Get("translate-ip-subnet.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TranslateIpSubnet = types.StringNull()

					cv := v.Get("translate-ip-subnet.vipVariableName")
					item.TranslateIpSubnetVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TranslateIpSubnet = types.StringNull()
					item.TranslateIpSubnetVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("translate-ip-subnet.vipValue")
					item.TranslateIpSubnet = types.StringValue(cv.String())
					item.TranslateIpSubnetVariable = types.StringNull()
				}
			} else {
				item.TranslateIpSubnet = types.StringNull()
				item.TranslateIpSubnetVariable = types.StringNull()
			}
			if cValue := v.Get("prefix-length.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.PrefixLength = types.Int64Null()

					cv := v.Get("prefix-length.vipVariableName")
					item.PrefixLengthVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.PrefixLength = types.Int64Null()
					item.PrefixLengthVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("prefix-length.vipValue")
					item.PrefixLength = types.Int64Value(cv.Int())
					item.PrefixLengthVariable = types.StringNull()
				}
			} else {
				item.PrefixLength = types.Int64Null()
				item.PrefixLengthVariable = types.StringNull()
			}
			if cValue := v.Get("static-nat-direction.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.StaticNatDirection = types.StringNull()

					cv := v.Get("static-nat-direction.vipVariableName")
					item.StaticNatDirectionVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.StaticNatDirection = types.StringNull()
					item.StaticNatDirectionVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("static-nat-direction.vipValue")
					item.StaticNatDirection = types.StringValue(cv.String())
					item.StaticNatDirectionVariable = types.StringNull()
				}
			} else {
				item.StaticNatDirection = types.StringNull()
				item.StaticNatDirectionVariable = types.StringNull()
			}
			if cValue := v.Get("tracker-id.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TrackerId = types.Int64Null()

					cv := v.Get("tracker-id.vipVariableName")
					item.TrackerIdVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TrackerId = types.Int64Null()
					item.TrackerIdVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("tracker-id.vipValue")
					item.TrackerId = types.Int64Value(cv.Int())
					item.TrackerIdVariable = types.StringNull()
				}
			} else {
				item.TrackerId = types.Int64Null()
				item.TrackerIdVariable = types.StringNull()
			}
			data.StaticNatSubnetRules = append(data.StaticNatSubnetRules, item)
			return true
		})
	} else {
		if len(data.StaticNatSubnetRules) > 0 {
			data.StaticNatSubnetRules = []CiscoVPNStaticNatSubnetRules{}
		}
	}
	if value := res.Get(path + "nat.port-forward.vipValue"); len(value.Array()) > 0 {
		data.PortForwardRules = make([]CiscoVPNPortForwardRules, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoVPNPortForwardRules{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("pool-name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.PoolName = types.Int64Null()

					cv := v.Get("pool-name.vipVariableName")
					item.PoolNameVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.PoolName = types.Int64Null()
					item.PoolNameVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("pool-name.vipValue")
					item.PoolName = types.Int64Value(cv.Int())
					item.PoolNameVariable = types.StringNull()
				}
			} else {
				item.PoolName = types.Int64Null()
				item.PoolNameVariable = types.StringNull()
			}
			if cValue := v.Get("source-port.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SourcePort = types.Int64Null()

					cv := v.Get("source-port.vipVariableName")
					item.SourcePortVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SourcePort = types.Int64Null()
					item.SourcePortVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("source-port.vipValue")
					item.SourcePort = types.Int64Value(cv.Int())
					item.SourcePortVariable = types.StringNull()
				}
			} else {
				item.SourcePort = types.Int64Null()
				item.SourcePortVariable = types.StringNull()
			}
			if cValue := v.Get("translate-port.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TranslatePort = types.Int64Null()

					cv := v.Get("translate-port.vipVariableName")
					item.TranslatePortVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TranslatePort = types.Int64Null()
					item.TranslatePortVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("translate-port.vipValue")
					item.TranslatePort = types.Int64Value(cv.Int())
					item.TranslatePortVariable = types.StringNull()
				}
			} else {
				item.TranslatePort = types.Int64Null()
				item.TranslatePortVariable = types.StringNull()
			}
			if cValue := v.Get("source-ip.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SourceIp = types.StringNull()

					cv := v.Get("source-ip.vipVariableName")
					item.SourceIpVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SourceIp = types.StringNull()
					item.SourceIpVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("source-ip.vipValue")
					item.SourceIp = types.StringValue(cv.String())
					item.SourceIpVariable = types.StringNull()
				}
			} else {
				item.SourceIp = types.StringNull()
				item.SourceIpVariable = types.StringNull()
			}
			if cValue := v.Get("translate-ip.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TranslateIp = types.StringNull()

					cv := v.Get("translate-ip.vipVariableName")
					item.TranslateIpVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TranslateIp = types.StringNull()
					item.TranslateIpVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("translate-ip.vipValue")
					item.TranslateIp = types.StringValue(cv.String())
					item.TranslateIpVariable = types.StringNull()
				}
			} else {
				item.TranslateIp = types.StringNull()
				item.TranslateIpVariable = types.StringNull()
			}
			if cValue := v.Get("proto.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Protocol = types.StringNull()

					cv := v.Get("proto.vipVariableName")
					item.ProtocolVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Protocol = types.StringNull()
					item.ProtocolVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("proto.vipValue")
					item.Protocol = types.StringValue(cv.String())
					item.ProtocolVariable = types.StringNull()
				}
			} else {
				item.Protocol = types.StringNull()
				item.ProtocolVariable = types.StringNull()
			}
			data.PortForwardRules = append(data.PortForwardRules, item)
			return true
		})
	} else {
		if len(data.PortForwardRules) > 0 {
			data.PortForwardRules = []CiscoVPNPortForwardRules{}
		}
	}
	if value := res.Get(path + "route-import.vipValue"); len(value.Array()) > 0 {
		data.RouteGlobalImports = make([]CiscoVPNRouteGlobalImports, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoVPNRouteGlobalImports{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("protocol.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Protocol = types.StringNull()

					cv := v.Get("protocol.vipVariableName")
					item.ProtocolVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Protocol = types.StringNull()
					item.ProtocolVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("protocol.vipValue")
					item.Protocol = types.StringValue(cv.String())
					item.ProtocolVariable = types.StringNull()
				}
			} else {
				item.Protocol = types.StringNull()
				item.ProtocolVariable = types.StringNull()
			}
			if cValue := v.Get("protocol-sub-type.vipType"); len(cValue.Array()) > 0 {
				if cValue.String() == "variableName" {
					item.ProtocolSubType = types.SetNull(types.StringType)

					cv := v.Get("protocol-sub-type.vipVariableName")
					item.ProtocolSubTypeVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.ProtocolSubType = types.SetNull(types.StringType)
					item.ProtocolSubTypeVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("protocol-sub-type.vipValue")
					item.ProtocolSubType = helpers.GetStringSet(cv.Array())
					item.ProtocolSubTypeVariable = types.StringNull()
				}
			} else {
				item.ProtocolSubType = types.SetNull(types.StringType)
				item.ProtocolSubTypeVariable = types.StringNull()
			}
			if cValue := v.Get("route-policy.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.RoutePolicy = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.RoutePolicy = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("route-policy.vipValue")
					item.RoutePolicy = types.StringValue(cv.String())

				}
			} else {
				item.RoutePolicy = types.StringNull()

			}
			if cValue := v.Get("redistribute.vipValue"); len(cValue.Array()) > 0 {
				item.Redistributes = make([]CiscoVPNRouteGlobalImportsRedistributes, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CiscoVPNRouteGlobalImportsRedistributes{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("protocol.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Protocol = types.StringNull()

							ccv := cv.Get("protocol.vipVariableName")
							cItem.ProtocolVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Protocol = types.StringNull()
							cItem.ProtocolVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("protocol.vipValue")
							cItem.Protocol = types.StringValue(ccv.String())
							cItem.ProtocolVariable = types.StringNull()
						}
					} else {
						cItem.Protocol = types.StringNull()
						cItem.ProtocolVariable = types.StringNull()
					}
					if ccValue := cv.Get("route-policy.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.RoutePolicy = types.StringNull()

						} else if ccValue.String() == "ignore" {
							cItem.RoutePolicy = types.StringNull()

						} else if ccValue.String() == "constant" {
							ccv := cv.Get("route-policy.vipValue")
							cItem.RoutePolicy = types.StringValue(ccv.String())

						}
					} else {
						cItem.RoutePolicy = types.StringNull()

					}
					item.Redistributes = append(item.Redistributes, cItem)
					return true
				})
			} else {
				if len(item.Redistributes) > 0 {
					item.Redistributes = []CiscoVPNRouteGlobalImportsRedistributes{}
				}
			}
			data.RouteGlobalImports = append(data.RouteGlobalImports, item)
			return true
		})
	} else {
		if len(data.RouteGlobalImports) > 0 {
			data.RouteGlobalImports = []CiscoVPNRouteGlobalImports{}
		}
	}
	if value := res.Get(path + "route-import-from.vipValue"); len(value.Array()) > 0 {
		data.RouteVpnImports = make([]CiscoVPNRouteVpnImports, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoVPNRouteVpnImports{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("source-vpn.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SourceVpnId = types.Int64Null()

					cv := v.Get("source-vpn.vipVariableName")
					item.SourceVpnIdVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SourceVpnId = types.Int64Null()
					item.SourceVpnIdVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("source-vpn.vipValue")
					item.SourceVpnId = types.Int64Value(cv.Int())
					item.SourceVpnIdVariable = types.StringNull()
				}
			} else {
				item.SourceVpnId = types.Int64Null()
				item.SourceVpnIdVariable = types.StringNull()
			}
			if cValue := v.Get("protocol.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Protocol = types.StringNull()

					cv := v.Get("protocol.vipVariableName")
					item.ProtocolVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Protocol = types.StringNull()
					item.ProtocolVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("protocol.vipValue")
					item.Protocol = types.StringValue(cv.String())
					item.ProtocolVariable = types.StringNull()
				}
			} else {
				item.Protocol = types.StringNull()
				item.ProtocolVariable = types.StringNull()
			}
			if cValue := v.Get("protocol-sub-type.vipType"); len(cValue.Array()) > 0 {
				if cValue.String() == "variableName" {
					item.ProtocolSubType = types.SetNull(types.StringType)

					cv := v.Get("protocol-sub-type.vipVariableName")
					item.ProtocolSubTypeVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.ProtocolSubType = types.SetNull(types.StringType)
					item.ProtocolSubTypeVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("protocol-sub-type.vipValue")
					item.ProtocolSubType = helpers.GetStringSet(cv.Array())
					item.ProtocolSubTypeVariable = types.StringNull()
				}
			} else {
				item.ProtocolSubType = types.SetNull(types.StringType)
				item.ProtocolSubTypeVariable = types.StringNull()
			}
			if cValue := v.Get("route-policy.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.RoutePolicy = types.StringNull()

					cv := v.Get("route-policy.vipVariableName")
					item.RoutePolicyVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.RoutePolicy = types.StringNull()
					item.RoutePolicyVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("route-policy.vipValue")
					item.RoutePolicy = types.StringValue(cv.String())
					item.RoutePolicyVariable = types.StringNull()
				}
			} else {
				item.RoutePolicy = types.StringNull()
				item.RoutePolicyVariable = types.StringNull()
			}
			if cValue := v.Get("redistribute.vipValue"); len(cValue.Array()) > 0 {
				item.Redistributes = make([]CiscoVPNRouteVpnImportsRedistributes, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CiscoVPNRouteVpnImportsRedistributes{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("protocol.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Protocol = types.StringNull()

							ccv := cv.Get("protocol.vipVariableName")
							cItem.ProtocolVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Protocol = types.StringNull()
							cItem.ProtocolVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("protocol.vipValue")
							cItem.Protocol = types.StringValue(ccv.String())
							cItem.ProtocolVariable = types.StringNull()
						}
					} else {
						cItem.Protocol = types.StringNull()
						cItem.ProtocolVariable = types.StringNull()
					}
					if ccValue := cv.Get("route-policy.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.RoutePolicy = types.StringNull()

							ccv := cv.Get("route-policy.vipVariableName")
							cItem.RoutePolicyVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.RoutePolicy = types.StringNull()
							cItem.RoutePolicyVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("route-policy.vipValue")
							cItem.RoutePolicy = types.StringValue(ccv.String())
							cItem.RoutePolicyVariable = types.StringNull()
						}
					} else {
						cItem.RoutePolicy = types.StringNull()
						cItem.RoutePolicyVariable = types.StringNull()
					}
					item.Redistributes = append(item.Redistributes, cItem)
					return true
				})
			} else {
				if len(item.Redistributes) > 0 {
					item.Redistributes = []CiscoVPNRouteVpnImportsRedistributes{}
				}
			}
			data.RouteVpnImports = append(data.RouteVpnImports, item)
			return true
		})
	} else {
		if len(data.RouteVpnImports) > 0 {
			data.RouteVpnImports = []CiscoVPNRouteVpnImports{}
		}
	}
	if value := res.Get(path + "route-export.vipValue"); len(value.Array()) > 0 {
		data.RouteGlobalExports = make([]CiscoVPNRouteGlobalExports, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoVPNRouteGlobalExports{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("protocol.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Protocol = types.StringNull()

					cv := v.Get("protocol.vipVariableName")
					item.ProtocolVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Protocol = types.StringNull()
					item.ProtocolVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("protocol.vipValue")
					item.Protocol = types.StringValue(cv.String())
					item.ProtocolVariable = types.StringNull()
				}
			} else {
				item.Protocol = types.StringNull()
				item.ProtocolVariable = types.StringNull()
			}
			if cValue := v.Get("protocol-sub-type.vipType"); len(cValue.Array()) > 0 {
				if cValue.String() == "variableName" {
					item.ProtocolSubType = types.SetNull(types.StringType)

					cv := v.Get("protocol-sub-type.vipVariableName")
					item.ProtocolSubTypeVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.ProtocolSubType = types.SetNull(types.StringType)
					item.ProtocolSubTypeVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("protocol-sub-type.vipValue")
					item.ProtocolSubType = helpers.GetStringSet(cv.Array())
					item.ProtocolSubTypeVariable = types.StringNull()
				}
			} else {
				item.ProtocolSubType = types.SetNull(types.StringType)
				item.ProtocolSubTypeVariable = types.StringNull()
			}
			if cValue := v.Get("route-policy.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.RoutePolicy = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.RoutePolicy = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("route-policy.vipValue")
					item.RoutePolicy = types.StringValue(cv.String())

				}
			} else {
				item.RoutePolicy = types.StringNull()

			}
			if cValue := v.Get("redistribute.vipValue"); len(cValue.Array()) > 0 {
				item.Redistributes = make([]CiscoVPNRouteGlobalExportsRedistributes, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CiscoVPNRouteGlobalExportsRedistributes{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("protocol.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Protocol = types.StringNull()

							ccv := cv.Get("protocol.vipVariableName")
							cItem.ProtocolVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Protocol = types.StringNull()
							cItem.ProtocolVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("protocol.vipValue")
							cItem.Protocol = types.StringValue(ccv.String())
							cItem.ProtocolVariable = types.StringNull()
						}
					} else {
						cItem.Protocol = types.StringNull()
						cItem.ProtocolVariable = types.StringNull()
					}
					if ccValue := cv.Get("route-policy.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.RoutePolicy = types.StringNull()

						} else if ccValue.String() == "ignore" {
							cItem.RoutePolicy = types.StringNull()

						} else if ccValue.String() == "constant" {
							ccv := cv.Get("route-policy.vipValue")
							cItem.RoutePolicy = types.StringValue(ccv.String())

						}
					} else {
						cItem.RoutePolicy = types.StringNull()

					}
					item.Redistributes = append(item.Redistributes, cItem)
					return true
				})
			} else {
				if len(item.Redistributes) > 0 {
					item.Redistributes = []CiscoVPNRouteGlobalExportsRedistributes{}
				}
			}
			data.RouteGlobalExports = append(data.RouteGlobalExports, item)
			return true
		})
	} else {
		if len(data.RouteGlobalExports) > 0 {
			data.RouteGlobalExports = []CiscoVPNRouteGlobalExports{}
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *CiscoVPN) hasChanges(ctx context.Context, state *CiscoVPN) bool {
	hasChanges := false
	if !data.VpnId.Equal(state.VpnId) {
		hasChanges = true
	}
	if !data.VpnName.Equal(state.VpnName) {
		hasChanges = true
	}
	if !data.TenantVpnId.Equal(state.TenantVpnId) {
		hasChanges = true
	}
	if !data.OrganizationName.Equal(state.OrganizationName) {
		hasChanges = true
	}
	if !data.OmpAdminDistanceIpv4.Equal(state.OmpAdminDistanceIpv4) {
		hasChanges = true
	}
	if !data.OmpAdminDistanceIpv6.Equal(state.OmpAdminDistanceIpv6) {
		hasChanges = true
	}
	if !data.EnhanceEcmpKeying.Equal(state.EnhanceEcmpKeying) {
		hasChanges = true
	}
	if len(data.DnsIpv4Servers) != len(state.DnsIpv4Servers) {
		hasChanges = true
	} else {
		for i := range data.DnsIpv4Servers {
			if !data.DnsIpv4Servers[i].Address.Equal(state.DnsIpv4Servers[i].Address) {
				hasChanges = true
			}
			if !data.DnsIpv4Servers[i].Role.Equal(state.DnsIpv4Servers[i].Role) {
				hasChanges = true
			}
		}
	}
	if len(data.DnsIpv6Servers) != len(state.DnsIpv6Servers) {
		hasChanges = true
	} else {
		for i := range data.DnsIpv6Servers {
			if !data.DnsIpv6Servers[i].Address.Equal(state.DnsIpv6Servers[i].Address) {
				hasChanges = true
			}
			if !data.DnsIpv6Servers[i].Role.Equal(state.DnsIpv6Servers[i].Role) {
				hasChanges = true
			}
		}
	}
	if len(data.DnsHosts) != len(state.DnsHosts) {
		hasChanges = true
	} else {
		for i := range data.DnsHosts {
			if !data.DnsHosts[i].Hostname.Equal(state.DnsHosts[i].Hostname) {
				hasChanges = true
			}
			if !data.DnsHosts[i].Ip.Equal(state.DnsHosts[i].Ip) {
				hasChanges = true
			}
		}
	}
	if len(data.Services) != len(state.Services) {
		hasChanges = true
	} else {
		for i := range data.Services {
			if !data.Services[i].ServiceTypes.Equal(state.Services[i].ServiceTypes) {
				hasChanges = true
			}
			if !data.Services[i].Address.Equal(state.Services[i].Address) {
				hasChanges = true
			}
			if !data.Services[i].Interface.Equal(state.Services[i].Interface) {
				hasChanges = true
			}
			if !data.Services[i].TrackEnable.Equal(state.Services[i].TrackEnable) {
				hasChanges = true
			}
		}
	}
	if len(data.Ipv4StaticServiceRoutes) != len(state.Ipv4StaticServiceRoutes) {
		hasChanges = true
	} else {
		for i := range data.Ipv4StaticServiceRoutes {
			if !data.Ipv4StaticServiceRoutes[i].Prefix.Equal(state.Ipv4StaticServiceRoutes[i].Prefix) {
				hasChanges = true
			}
			if !data.Ipv4StaticServiceRoutes[i].VpnId.Equal(state.Ipv4StaticServiceRoutes[i].VpnId) {
				hasChanges = true
			}
			if !data.Ipv4StaticServiceRoutes[i].Service.Equal(state.Ipv4StaticServiceRoutes[i].Service) {
				hasChanges = true
			}
		}
	}
	if len(data.Ipv4StaticRoutes) != len(state.Ipv4StaticRoutes) {
		hasChanges = true
	} else {
		for i := range data.Ipv4StaticRoutes {
			if !data.Ipv4StaticRoutes[i].Prefix.Equal(state.Ipv4StaticRoutes[i].Prefix) {
				hasChanges = true
			}
			if !data.Ipv4StaticRoutes[i].Null0.Equal(state.Ipv4StaticRoutes[i].Null0) {
				hasChanges = true
			}
			if !data.Ipv4StaticRoutes[i].Distance.Equal(state.Ipv4StaticRoutes[i].Distance) {
				hasChanges = true
			}
			if !data.Ipv4StaticRoutes[i].VpnId.Equal(state.Ipv4StaticRoutes[i].VpnId) {
				hasChanges = true
			}
			if !data.Ipv4StaticRoutes[i].Dhcp.Equal(state.Ipv4StaticRoutes[i].Dhcp) {
				hasChanges = true
			}
			if len(data.Ipv4StaticRoutes[i].NextHops) != len(state.Ipv4StaticRoutes[i].NextHops) {
				hasChanges = true
			} else {
				for ii := range data.Ipv4StaticRoutes[i].NextHops {
					if !data.Ipv4StaticRoutes[i].NextHops[ii].Address.Equal(state.Ipv4StaticRoutes[i].NextHops[ii].Address) {
						hasChanges = true
					}
					if !data.Ipv4StaticRoutes[i].NextHops[ii].Distance.Equal(state.Ipv4StaticRoutes[i].NextHops[ii].Distance) {
						hasChanges = true
					}
				}
			}
			if len(data.Ipv4StaticRoutes[i].TrackNextHops) != len(state.Ipv4StaticRoutes[i].TrackNextHops) {
				hasChanges = true
			} else {
				for ii := range data.Ipv4StaticRoutes[i].TrackNextHops {
					if !data.Ipv4StaticRoutes[i].TrackNextHops[ii].Address.Equal(state.Ipv4StaticRoutes[i].TrackNextHops[ii].Address) {
						hasChanges = true
					}
					if !data.Ipv4StaticRoutes[i].TrackNextHops[ii].Distance.Equal(state.Ipv4StaticRoutes[i].TrackNextHops[ii].Distance) {
						hasChanges = true
					}
					if !data.Ipv4StaticRoutes[i].TrackNextHops[ii].Tracker.Equal(state.Ipv4StaticRoutes[i].TrackNextHops[ii].Tracker) {
						hasChanges = true
					}
				}
			}
		}
	}
	if len(data.Ipv6StaticRoutes) != len(state.Ipv6StaticRoutes) {
		hasChanges = true
	} else {
		for i := range data.Ipv6StaticRoutes {
			if !data.Ipv6StaticRoutes[i].Prefix.Equal(state.Ipv6StaticRoutes[i].Prefix) {
				hasChanges = true
			}
			if !data.Ipv6StaticRoutes[i].Null0.Equal(state.Ipv6StaticRoutes[i].Null0) {
				hasChanges = true
			}
			if !data.Ipv6StaticRoutes[i].VpnId.Equal(state.Ipv6StaticRoutes[i].VpnId) {
				hasChanges = true
			}
			if !data.Ipv6StaticRoutes[i].Nat.Equal(state.Ipv6StaticRoutes[i].Nat) {
				hasChanges = true
			}
			if len(data.Ipv6StaticRoutes[i].NextHops) != len(state.Ipv6StaticRoutes[i].NextHops) {
				hasChanges = true
			} else {
				for ii := range data.Ipv6StaticRoutes[i].NextHops {
					if !data.Ipv6StaticRoutes[i].NextHops[ii].Address.Equal(state.Ipv6StaticRoutes[i].NextHops[ii].Address) {
						hasChanges = true
					}
					if !data.Ipv6StaticRoutes[i].NextHops[ii].Distance.Equal(state.Ipv6StaticRoutes[i].NextHops[ii].Distance) {
						hasChanges = true
					}
				}
			}
		}
	}
	if len(data.Ipv4StaticGreRoutes) != len(state.Ipv4StaticGreRoutes) {
		hasChanges = true
	} else {
		for i := range data.Ipv4StaticGreRoutes {
			if !data.Ipv4StaticGreRoutes[i].Prefix.Equal(state.Ipv4StaticGreRoutes[i].Prefix) {
				hasChanges = true
			}
			if !data.Ipv4StaticGreRoutes[i].VpnId.Equal(state.Ipv4StaticGreRoutes[i].VpnId) {
				hasChanges = true
			}
			if !data.Ipv4StaticGreRoutes[i].Interface.Equal(state.Ipv4StaticGreRoutes[i].Interface) {
				hasChanges = true
			}
		}
	}
	if len(data.Ipv4StaticIpsecRoutes) != len(state.Ipv4StaticIpsecRoutes) {
		hasChanges = true
	} else {
		for i := range data.Ipv4StaticIpsecRoutes {
			if !data.Ipv4StaticIpsecRoutes[i].Prefix.Equal(state.Ipv4StaticIpsecRoutes[i].Prefix) {
				hasChanges = true
			}
			if !data.Ipv4StaticIpsecRoutes[i].VpnId.Equal(state.Ipv4StaticIpsecRoutes[i].VpnId) {
				hasChanges = true
			}
			if !data.Ipv4StaticIpsecRoutes[i].Interface.Equal(state.Ipv4StaticIpsecRoutes[i].Interface) {
				hasChanges = true
			}
		}
	}
	if len(data.OmpAdvertiseIpv4Routes) != len(state.OmpAdvertiseIpv4Routes) {
		hasChanges = true
	} else {
		for i := range data.OmpAdvertiseIpv4Routes {
			if !data.OmpAdvertiseIpv4Routes[i].Protocol.Equal(state.OmpAdvertiseIpv4Routes[i].Protocol) {
				hasChanges = true
			}
			if !data.OmpAdvertiseIpv4Routes[i].RoutePolicy.Equal(state.OmpAdvertiseIpv4Routes[i].RoutePolicy) {
				hasChanges = true
			}
			if !data.OmpAdvertiseIpv4Routes[i].ProtocolSubType.Equal(state.OmpAdvertiseIpv4Routes[i].ProtocolSubType) {
				hasChanges = true
			}
			if len(data.OmpAdvertiseIpv4Routes[i].Prefixes) != len(state.OmpAdvertiseIpv4Routes[i].Prefixes) {
				hasChanges = true
			} else {
				for ii := range data.OmpAdvertiseIpv4Routes[i].Prefixes {
					if !data.OmpAdvertiseIpv4Routes[i].Prefixes[ii].PrefixEntry.Equal(state.OmpAdvertiseIpv4Routes[i].Prefixes[ii].PrefixEntry) {
						hasChanges = true
					}
					if !data.OmpAdvertiseIpv4Routes[i].Prefixes[ii].AggregateOnly.Equal(state.OmpAdvertiseIpv4Routes[i].Prefixes[ii].AggregateOnly) {
						hasChanges = true
					}
				}
			}
		}
	}
	if len(data.OmpAdvertiseIpv6Routes) != len(state.OmpAdvertiseIpv6Routes) {
		hasChanges = true
	} else {
		for i := range data.OmpAdvertiseIpv6Routes {
			if !data.OmpAdvertiseIpv6Routes[i].Protocol.Equal(state.OmpAdvertiseIpv6Routes[i].Protocol) {
				hasChanges = true
			}
			if !data.OmpAdvertiseIpv6Routes[i].RoutePolicy.Equal(state.OmpAdvertiseIpv6Routes[i].RoutePolicy) {
				hasChanges = true
			}
			if !data.OmpAdvertiseIpv6Routes[i].ProtocolSubType.Equal(state.OmpAdvertiseIpv6Routes[i].ProtocolSubType) {
				hasChanges = true
			}
			if len(data.OmpAdvertiseIpv6Routes[i].Prefixes) != len(state.OmpAdvertiseIpv6Routes[i].Prefixes) {
				hasChanges = true
			} else {
				for ii := range data.OmpAdvertiseIpv6Routes[i].Prefixes {
					if !data.OmpAdvertiseIpv6Routes[i].Prefixes[ii].PrefixEntry.Equal(state.OmpAdvertiseIpv6Routes[i].Prefixes[ii].PrefixEntry) {
						hasChanges = true
					}
					if !data.OmpAdvertiseIpv6Routes[i].Prefixes[ii].AggregateOnly.Equal(state.OmpAdvertiseIpv6Routes[i].Prefixes[ii].AggregateOnly) {
						hasChanges = true
					}
				}
			}
		}
	}
	if len(data.Nat64Pools) != len(state.Nat64Pools) {
		hasChanges = true
	} else {
		for i := range data.Nat64Pools {
			if !data.Nat64Pools[i].Name.Equal(state.Nat64Pools[i].Name) {
				hasChanges = true
			}
			if !data.Nat64Pools[i].StartAddress.Equal(state.Nat64Pools[i].StartAddress) {
				hasChanges = true
			}
			if !data.Nat64Pools[i].EndAddress.Equal(state.Nat64Pools[i].EndAddress) {
				hasChanges = true
			}
			if !data.Nat64Pools[i].Overload.Equal(state.Nat64Pools[i].Overload) {
				hasChanges = true
			}
			if !data.Nat64Pools[i].LeakFromGlobal.Equal(state.Nat64Pools[i].LeakFromGlobal) {
				hasChanges = true
			}
			if !data.Nat64Pools[i].LeakFromGlobalProtocol.Equal(state.Nat64Pools[i].LeakFromGlobalProtocol) {
				hasChanges = true
			}
			if !data.Nat64Pools[i].LeakToGlobal.Equal(state.Nat64Pools[i].LeakToGlobal) {
				hasChanges = true
			}
		}
	}
	if len(data.NatPools) != len(state.NatPools) {
		hasChanges = true
	} else {
		for i := range data.NatPools {
			if !data.NatPools[i].Name.Equal(state.NatPools[i].Name) {
				hasChanges = true
			}
			if !data.NatPools[i].PrefixLength.Equal(state.NatPools[i].PrefixLength) {
				hasChanges = true
			}
			if !data.NatPools[i].RangeStart.Equal(state.NatPools[i].RangeStart) {
				hasChanges = true
			}
			if !data.NatPools[i].RangeEnd.Equal(state.NatPools[i].RangeEnd) {
				hasChanges = true
			}
			if !data.NatPools[i].Overload.Equal(state.NatPools[i].Overload) {
				hasChanges = true
			}
			if !data.NatPools[i].Direction.Equal(state.NatPools[i].Direction) {
				hasChanges = true
			}
			if !data.NatPools[i].TrackerId.Equal(state.NatPools[i].TrackerId) {
				hasChanges = true
			}
		}
	}
	if len(data.StaticNatRules) != len(state.StaticNatRules) {
		hasChanges = true
	} else {
		for i := range data.StaticNatRules {
			if !data.StaticNatRules[i].PoolName.Equal(state.StaticNatRules[i].PoolName) {
				hasChanges = true
			}
			if !data.StaticNatRules[i].SourceIp.Equal(state.StaticNatRules[i].SourceIp) {
				hasChanges = true
			}
			if !data.StaticNatRules[i].TranslateIp.Equal(state.StaticNatRules[i].TranslateIp) {
				hasChanges = true
			}
			if !data.StaticNatRules[i].StaticNatDirection.Equal(state.StaticNatRules[i].StaticNatDirection) {
				hasChanges = true
			}
			if !data.StaticNatRules[i].TrackerId.Equal(state.StaticNatRules[i].TrackerId) {
				hasChanges = true
			}
		}
	}
	if len(data.StaticNatSubnetRules) != len(state.StaticNatSubnetRules) {
		hasChanges = true
	} else {
		for i := range data.StaticNatSubnetRules {
			if !data.StaticNatSubnetRules[i].SourceIpSubnet.Equal(state.StaticNatSubnetRules[i].SourceIpSubnet) {
				hasChanges = true
			}
			if !data.StaticNatSubnetRules[i].TranslateIpSubnet.Equal(state.StaticNatSubnetRules[i].TranslateIpSubnet) {
				hasChanges = true
			}
			if !data.StaticNatSubnetRules[i].PrefixLength.Equal(state.StaticNatSubnetRules[i].PrefixLength) {
				hasChanges = true
			}
			if !data.StaticNatSubnetRules[i].StaticNatDirection.Equal(state.StaticNatSubnetRules[i].StaticNatDirection) {
				hasChanges = true
			}
			if !data.StaticNatSubnetRules[i].TrackerId.Equal(state.StaticNatSubnetRules[i].TrackerId) {
				hasChanges = true
			}
		}
	}
	if len(data.PortForwardRules) != len(state.PortForwardRules) {
		hasChanges = true
	} else {
		for i := range data.PortForwardRules {
			if !data.PortForwardRules[i].PoolName.Equal(state.PortForwardRules[i].PoolName) {
				hasChanges = true
			}
			if !data.PortForwardRules[i].SourcePort.Equal(state.PortForwardRules[i].SourcePort) {
				hasChanges = true
			}
			if !data.PortForwardRules[i].TranslatePort.Equal(state.PortForwardRules[i].TranslatePort) {
				hasChanges = true
			}
			if !data.PortForwardRules[i].SourceIp.Equal(state.PortForwardRules[i].SourceIp) {
				hasChanges = true
			}
			if !data.PortForwardRules[i].TranslateIp.Equal(state.PortForwardRules[i].TranslateIp) {
				hasChanges = true
			}
			if !data.PortForwardRules[i].Protocol.Equal(state.PortForwardRules[i].Protocol) {
				hasChanges = true
			}
		}
	}
	if len(data.RouteGlobalImports) != len(state.RouteGlobalImports) {
		hasChanges = true
	} else {
		for i := range data.RouteGlobalImports {
			if !data.RouteGlobalImports[i].Protocol.Equal(state.RouteGlobalImports[i].Protocol) {
				hasChanges = true
			}
			if !data.RouteGlobalImports[i].ProtocolSubType.Equal(state.RouteGlobalImports[i].ProtocolSubType) {
				hasChanges = true
			}
			if !data.RouteGlobalImports[i].RoutePolicy.Equal(state.RouteGlobalImports[i].RoutePolicy) {
				hasChanges = true
			}
			if len(data.RouteGlobalImports[i].Redistributes) != len(state.RouteGlobalImports[i].Redistributes) {
				hasChanges = true
			} else {
				for ii := range data.RouteGlobalImports[i].Redistributes {
					if !data.RouteGlobalImports[i].Redistributes[ii].Protocol.Equal(state.RouteGlobalImports[i].Redistributes[ii].Protocol) {
						hasChanges = true
					}
					if !data.RouteGlobalImports[i].Redistributes[ii].RoutePolicy.Equal(state.RouteGlobalImports[i].Redistributes[ii].RoutePolicy) {
						hasChanges = true
					}
				}
			}
		}
	}
	if len(data.RouteVpnImports) != len(state.RouteVpnImports) {
		hasChanges = true
	} else {
		for i := range data.RouteVpnImports {
			if !data.RouteVpnImports[i].SourceVpnId.Equal(state.RouteVpnImports[i].SourceVpnId) {
				hasChanges = true
			}
			if !data.RouteVpnImports[i].Protocol.Equal(state.RouteVpnImports[i].Protocol) {
				hasChanges = true
			}
			if !data.RouteVpnImports[i].ProtocolSubType.Equal(state.RouteVpnImports[i].ProtocolSubType) {
				hasChanges = true
			}
			if !data.RouteVpnImports[i].RoutePolicy.Equal(state.RouteVpnImports[i].RoutePolicy) {
				hasChanges = true
			}
			if len(data.RouteVpnImports[i].Redistributes) != len(state.RouteVpnImports[i].Redistributes) {
				hasChanges = true
			} else {
				for ii := range data.RouteVpnImports[i].Redistributes {
					if !data.RouteVpnImports[i].Redistributes[ii].Protocol.Equal(state.RouteVpnImports[i].Redistributes[ii].Protocol) {
						hasChanges = true
					}
					if !data.RouteVpnImports[i].Redistributes[ii].RoutePolicy.Equal(state.RouteVpnImports[i].Redistributes[ii].RoutePolicy) {
						hasChanges = true
					}
				}
			}
		}
	}
	if len(data.RouteGlobalExports) != len(state.RouteGlobalExports) {
		hasChanges = true
	} else {
		for i := range data.RouteGlobalExports {
			if !data.RouteGlobalExports[i].Protocol.Equal(state.RouteGlobalExports[i].Protocol) {
				hasChanges = true
			}
			if !data.RouteGlobalExports[i].ProtocolSubType.Equal(state.RouteGlobalExports[i].ProtocolSubType) {
				hasChanges = true
			}
			if !data.RouteGlobalExports[i].RoutePolicy.Equal(state.RouteGlobalExports[i].RoutePolicy) {
				hasChanges = true
			}
			if len(data.RouteGlobalExports[i].Redistributes) != len(state.RouteGlobalExports[i].Redistributes) {
				hasChanges = true
			} else {
				for ii := range data.RouteGlobalExports[i].Redistributes {
					if !data.RouteGlobalExports[i].Redistributes[ii].Protocol.Equal(state.RouteGlobalExports[i].Redistributes[ii].Protocol) {
						hasChanges = true
					}
					if !data.RouteGlobalExports[i].Redistributes[ii].RoutePolicy.Equal(state.RouteGlobalExports[i].Redistributes[ii].RoutePolicy) {
						hasChanges = true
					}
				}
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges
