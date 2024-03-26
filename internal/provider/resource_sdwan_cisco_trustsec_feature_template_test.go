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

func TestAccSdwanCiscoTrustSecFeatureTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanCiscoTrustSecFeatureTemplateConfig_minimum(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "sxp_node_id_type", "interface-name"),
					resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "sxp_node_id", "VirtualPortGroup"),
				),
			},
			{
				Config: testAccSdwanCiscoTrustSecFeatureTemplateConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "device_sgt", "100"),
					resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "credentials_id", "example"),
					resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "credentials_password", "MyPassword"),
					resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "enable_enforcement", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "enable_sxp", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "sxp_source_ip", "1.2.3.4"),
					resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "sxp_default_password", "MyPassword"),
					resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "sxp_key_chain", "example"),
					resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "sxp_log_binding_changes", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "sxp_reconciliation_period", "120"),
					resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "sxp_retry_period", "120"),
					resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "speaker_hold_time", "120"),
					resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "listener_hold_time_min", "90"),
					resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "listener_hold_time_max", "180"),
					resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "sxp_node_id_type", "interface-name"),
					resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "sxp_node_id", "VirtualPortGroup"),
					resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "sxp_connection_list.0.connection_peer_ip", "1.2.3.4"),
					resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "sxp_connection_list.0.connection_source_ip", "2.3.4.5"),
					resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "sxp_connection_list.0.connection_preshared_key", "default"),
					resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "sxp_connection_list.0.connection_mode", "local"),
					resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "sxp_connection_list.0.connection_mode_type", "listener"),
					resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "sxp_connection_list.0.connection_min_hold_time", "0"),
					resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "sxp_connection_list.0.connection_max_hold_time", "0"),
					resource.TestCheckResourceAttr("sdwan_cisco_trustsec_feature_template.test", "sxp_connection_list.0.connection_vpn_id", "0"),
				),
			},
		},
	})
}

func testAccSdwanCiscoTrustSecFeatureTemplateConfig_minimum() string {
	return `
	resource "sdwan_cisco_trustsec_feature_template" "test" {
		name = "TF_TEST_MIN"
		description = "Terraform integration test"
		device_types = ["vedge-C8000V"]
		sxp_node_id_type = "interface-name"
		sxp_node_id = "VirtualPortGroup"
	}
	`
}

func testAccSdwanCiscoTrustSecFeatureTemplateConfig_all() string {
	return `
	resource "sdwan_cisco_trustsec_feature_template" "test" {
		name = "TF_TEST_ALL"
		description = "Terraform integration test"
		device_types = ["vedge-C8000V"]
		device_sgt = 100
		credentials_id = "example"
		credentials_password = "MyPassword"
		enable_enforcement = true
		enable_sxp = true
		sxp_source_ip = "1.2.3.4"
		sxp_default_password = "MyPassword"
		sxp_key_chain = "example"
		sxp_log_binding_changes = false
		sxp_reconciliation_period = 120
		sxp_retry_period = 120
		speaker_hold_time = 120
		listener_hold_time_min = 90
		listener_hold_time_max = 180
		sxp_node_id_type = "interface-name"
		sxp_node_id = "VirtualPortGroup"
		sxp_connection_list = [{
			connection_peer_ip = "1.2.3.4"
			connection_source_ip = "2.3.4.5"
			connection_preshared_key = "default"
			connection_mode = "local"
			connection_mode_type = "listener"
			connection_min_hold_time = 0
			connection_max_hold_time = 0
			connection_vpn_id = 0
		}]
	}
	`
}
