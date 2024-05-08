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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc
func TestAccSdwanIPv6DeviceACLPolicyDefinition(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_ipv6_device_acl_policy_definition.test", "name", "Example"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_ipv6_device_acl_policy_definition.test", "description", "My description"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_ipv6_device_acl_policy_definition.test", "default_action", "drop"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_ipv6_device_acl_policy_definition.test", "sequences.0.id", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_ipv6_device_acl_policy_definition.test", "sequences.0.name", "Sequence 10"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_ipv6_device_acl_policy_definition.test", "sequences.0.base_action", "accept"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_ipv6_device_acl_policy_definition.test", "sequences.0.match_entries.0.type", "destinationPort"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_ipv6_device_acl_policy_definition.test", "sequences.0.match_entries.0.destination_port", "22"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_ipv6_device_acl_policy_definition.test", "sequences.0.action_entries.0.type", "count"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_ipv6_device_acl_policy_definition.test", "sequences.0.action_entries.0.counter_name", "count1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanIPv6DeviceACLPolicyDefinitionConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanIPv6DeviceACLPolicyDefinitionConfig_all() string {
	config := `resource "sdwan_ipv6_device_acl_policy_definition" "test" {` + "\n"
	config += `	name = "Example"` + "\n"
	config += `	description = "My description"` + "\n"
	config += `	default_action = "drop"` + "\n"
	config += `	sequences = [{` + "\n"
	config += `	  id = 10` + "\n"
	config += `	  name = "Sequence 10"` + "\n"
	config += `	  base_action = "accept"` + "\n"
	config += `	  match_entries = [{` + "\n"
	config += `		type = "destinationPort"` + "\n"
	config += `		destination_port = 22` + "\n"
	config += `	}]` + "\n"
	config += `	  action_entries = [{` + "\n"
	config += `		type = "count"` + "\n"
	config += `		counter_name = "count1"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
