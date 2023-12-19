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

func TestAccSdwanIPv6DeviceACLPolicyDefinition(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanIPv6DeviceACLPolicyDefinitionConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_ipv6_device_acl_policy_definition.test", "name", "Example"),
					resource.TestCheckResourceAttr("sdwan_ipv6_device_acl_policy_definition.test", "description", "My description"),
					resource.TestCheckResourceAttr("sdwan_ipv6_device_acl_policy_definition.test", "default_action", "drop"),
					resource.TestCheckResourceAttr("sdwan_ipv6_device_acl_policy_definition.test", "sequences.0.id", "10"),
					resource.TestCheckResourceAttr("sdwan_ipv6_device_acl_policy_definition.test", "sequences.0.name", "Sequence 10"),
					resource.TestCheckResourceAttr("sdwan_ipv6_device_acl_policy_definition.test", "sequences.0.base_action", "accept"),
					resource.TestCheckResourceAttr("sdwan_ipv6_device_acl_policy_definition.test", "sequences.0.match_entries.0.type", "destinationPort"),
					resource.TestCheckResourceAttr("sdwan_ipv6_device_acl_policy_definition.test", "sequences.0.match_entries.0.destination_port", "22"),
					resource.TestCheckResourceAttr("sdwan_ipv6_device_acl_policy_definition.test", "sequences.0.action_entries.0.type", "count"),
					resource.TestCheckResourceAttr("sdwan_ipv6_device_acl_policy_definition.test", "sequences.0.action_entries.0.counter_name", "count1"),
				),
			},
		},
	})
}

const testAccSdwanIPv6DeviceACLPolicyDefinitionConfig = `


resource "sdwan_ipv6_device_acl_policy_definition" "test" {
	name = "Example"
	description = "My description"
	default_action = "drop"
	sequences = [{
		id = 10
		name = "Sequence 10"
		base_action = "accept"
		match_entries = [{
			type = "destinationPort"
			destination_port = 22
		}]
		action_entries = [{
			type = "count"
			counter_name = "count1"
		}]
	}]
}
`
