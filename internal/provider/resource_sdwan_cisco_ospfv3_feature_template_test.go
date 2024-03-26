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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSdwanCiscoOSPFv3FeatureTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanCiscoOSPFv3FeatureTemplateConfig_minimum(),
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: testAccSdwanCiscoOSPFv3FeatureTemplateConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "router_id_ipv4", "1.2.3.4"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "auto_cost_reference_bandwidth_ipv4", "100000"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "compatible_rfc1583_ipv4", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "default_information_originate_ipv4", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "default_information_originate_always_ipv4", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "default_information_originate_metric_ipv4", "100"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "default_information_originate_metric_type_ipv4", "type1"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "distance_external_ipv4", "111"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "distance_inter_area_ipv4", "111"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "distance_intra_area_ipv4", "112"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "timers_spf_delay_ipv4", "300"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "timers_spf_initial_hold_ipv4", "2000"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "timers_spf_max_hold_ipv4", "20000"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "distance_ipv4", "110"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "policy_name_ipv4", "example"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "filter_ipv4", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "redistribute_ipv4.0.protocol", "static"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "redistribute_ipv4.0.route_policy", "RP1"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "redistribute_ipv4.0.nat_dia", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "max_metric_router_lsa_ipv4.0.ad_type", "on-startup"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "max_metric_router_lsa_ipv4.0.time", "100"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv4.0.area_number", "1"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv4.0.stub", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv4.0.stub_no_summary", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv4.0.nssa", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv4.0.nssa_no_summary", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv4.0.translate", "always"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv4.0.normal", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv4.0.interfaces.0.name", "e1"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv4.0.interfaces.0.hello_interval", "20"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv4.0.interfaces.0.dead_interval", "60"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv4.0.interfaces.0.retransmit_interval", "10"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv4.0.interfaces.0.cost", "100"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv4.0.interfaces.0.network", "point-to-point"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv4.0.interfaces.0.passive_interface", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv4.0.interfaces.0.authentication_type", "message-digest"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv4.0.interfaces.0.authentication_key", "authenticationKey"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv4.0.interfaces.0.ipsec_spi", "256"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv4.0.ranges.0.address", "1.1.1.0/24"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv4.0.ranges.0.cost", "100"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv4.0.ranges.0.no_advertise", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "router_id_ipv6", "1.2.3.4"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "auto_cost_reference_bandwidth_ipv6", "100000"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "compatible_rfc1583_ipv6", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "default_information_originate_ipv6", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "default_information_originate_always_ipv6", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "default_information_originate_metric_ipv6", "100"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "default_information_originate_metric_type_ipv6", "type1"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "distance_external_ipv6", "111"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "distance_inter_area_ipv6", "111"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "distance_intra_area_ipv6", "112"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "timers_spf_delay_ipv6", "300"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "timers_spf_initial_hold_ipv6", "2000"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "timers_spf_max_hold_ipv6", "20000"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "distance_ipv6", "110"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "policy_name_ipv6", "example"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "filter_ipv6", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "redistribute_ipv6.0.protocol", "static"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "redistribute_ipv6.0.route_policy", "RP1"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "max_metric_router_lsa_ipv6.0.ad_type", "on-startup"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "max_metric_router_lsa_ipv6.0.time", "100"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv6.0.area_number", "1"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv6.0.stub", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv6.0.stub_no_summary", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv6.0.nssa", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv6.0.nssa_no_summary", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv6.0.translate", "always"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv6.0.normal", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv6.0.interfaces.0.name", "e1"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv6.0.interfaces.0.hello_interval", "20"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv6.0.interfaces.0.dead_interval", "60"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv6.0.interfaces.0.retransmit_interval", "10"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv6.0.interfaces.0.cost", "100"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv6.0.interfaces.0.network", "point-to-point"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv6.0.interfaces.0.passive_interface", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv6.0.interfaces.0.authentication_type", "message-digest"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv6.0.interfaces.0.authentication_key", "authenticationKey"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv6.0.interfaces.0.ipsec_spi", "256"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv6.0.ranges.0.address", "1.1.1.0/24"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv6.0.ranges.0.cost", "100"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "areas_ipv6.0.ranges.0.no_advertise", "true"),
				),
			},
		},
	})
}

func testAccSdwanCiscoOSPFv3FeatureTemplateConfig_minimum() string {
	return `
	resource "sdwan_cisco_ospfv3_feature_template" "test" {
		name = "TF_TEST_MIN"
		description = "Terraform integration test"
		device_types = ["vedge-C8000V"]
	}
	`
}

func testAccSdwanCiscoOSPFv3FeatureTemplateConfig_all() string {
	return `
	resource "sdwan_cisco_ospfv3_feature_template" "test" {
		name = "TF_TEST_ALL"
		description = "Terraform integration test"
		device_types = ["vedge-C8000V"]
		router_id_ipv4 = "1.2.3.4"
		auto_cost_reference_bandwidth_ipv4 = 100000
		compatible_rfc1583_ipv4 = true
		default_information_originate_ipv4 = true
		default_information_originate_always_ipv4 = true
		default_information_originate_metric_ipv4 = 100
		default_information_originate_metric_type_ipv4 = "type1"
		distance_external_ipv4 = 111
		distance_inter_area_ipv4 = 111
		distance_intra_area_ipv4 = 112
		timers_spf_delay_ipv4 = 300
		timers_spf_initial_hold_ipv4 = 2000
		timers_spf_max_hold_ipv4 = 20000
		distance_ipv4 = 110
		policy_name_ipv4 = "example"
		filter_ipv4 = false
		redistribute_ipv4 = [{
			protocol = "static"
			route_policy = "RP1"
			nat_dia = true
		}]
		max_metric_router_lsa_ipv4 = [{
			ad_type = "on-startup"
			time = 100
		}]
		areas_ipv4 = [{
			area_number = 1
			stub = false
			stub_no_summary = false
			nssa = false
			nssa_no_summary = true
			translate = "always"
			normal = false
			interfaces = [{
				name = "e1"
				hello_interval = 20
				dead_interval = 60
				retransmit_interval = 10
				cost = 100
				network = "point-to-point"
				passive_interface = true
				authentication_type = "message-digest"
				authentication_key = "authenticationKey"
				ipsec_spi = 256
			}]
			ranges = [{
				address = "1.1.1.0/24"
				cost = 100
				no_advertise = true
			}]
		}]
		router_id_ipv6 = "1.2.3.4"
		auto_cost_reference_bandwidth_ipv6 = 100000
		compatible_rfc1583_ipv6 = true
		default_information_originate_ipv6 = true
		default_information_originate_always_ipv6 = true
		default_information_originate_metric_ipv6 = 100
		default_information_originate_metric_type_ipv6 = "type1"
		distance_external_ipv6 = 111
		distance_inter_area_ipv6 = 111
		distance_intra_area_ipv6 = 112
		timers_spf_delay_ipv6 = 300
		timers_spf_initial_hold_ipv6 = 2000
		timers_spf_max_hold_ipv6 = 20000
		distance_ipv6 = 110
		policy_name_ipv6 = "example"
		filter_ipv6 = false
		redistribute_ipv6 = [{
			protocol = "static"
			route_policy = "RP1"
		}]
		max_metric_router_lsa_ipv6 = [{
			ad_type = "on-startup"
			time = 100
		}]
		areas_ipv6 = [{
			area_number = 1
			stub = false
			stub_no_summary = false
			nssa = false
			nssa_no_summary = true
			translate = "always"
			normal = false
			interfaces = [{
				name = "e1"
				hello_interval = 20
				dead_interval = 60
				retransmit_interval = 10
				cost = 100
				network = "point-to-point"
				passive_interface = true
				authentication_type = "message-digest"
				authentication_key = "authenticationKey"
				ipsec_spi = 256
			}]
			ranges = [{
				address = "1.1.1.0/24"
				cost = 100
				no_advertise = true
			}]
		}]
	}
	`
}
