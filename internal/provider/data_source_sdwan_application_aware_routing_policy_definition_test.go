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

func TestAccDataSourceSdwanApplicationAwareRoutingPolicyDefinition(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanApplicationAwareRoutingPolicyDefinitionConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdwan_application_aware_routing_policy_definition.test", "sequences.0.id", "1"),
					resource.TestCheckResourceAttr("data.sdwan_application_aware_routing_policy_definition.test", "sequences.0.name", "Region1"),
					resource.TestCheckResourceAttr("data.sdwan_application_aware_routing_policy_definition.test", "sequences.0.ip_type", "ipv4"),
					resource.TestCheckResourceAttr("data.sdwan_application_aware_routing_policy_definition.test", "sequences.0.match_entries.0.type", "appList"),
					resource.TestCheckResourceAttr("data.sdwan_application_aware_routing_policy_definition.test", "sequences.0.match_entries.0.application_list_id", "e3aad846-abb9-425f-aaa8-9ed17b9c8d7c"),
					resource.TestCheckResourceAttr("data.sdwan_application_aware_routing_policy_definition.test", "sequences.0.action_entries.0.type", "backupSlaPreferredColor"),
					resource.TestCheckResourceAttr("data.sdwan_application_aware_routing_policy_definition.test", "sequences.0.action_entries.0.backup_sla_preferred_color", "bronze"),
				),
			},
		},
	})
}

const testAccDataSourceSdwanApplicationAwareRoutingPolicyDefinitionConfig = `

resource "sdwan_application_aware_routing_policy_definition" "test" {
  name = "TF_TEST_MIN"
  description = "Terraform integration test"
  sequences = [{
    id = 1
    name = "Region1"
    ip_type = "ipv4"
	match_entries = [{
		type = "appList"
		application_list_id = "e3aad846-abb9-425f-aaa8-9ed17b9c8d7c"
	}]
	action_entries = [{
		type = "backupSlaPreferredColor"
		backup_sla_preferred_color = "bronze"
	}]
  }]
}

data "sdwan_application_aware_routing_policy_definition" "test" {
  id = sdwan_application_aware_routing_policy_definition.test.id
}
`
