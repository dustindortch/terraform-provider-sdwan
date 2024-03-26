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

func TestAccDataSourceSdwanEigrpFeatureTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanEigrpFeatureTemplateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdwan_eigrp_feature_template.test", "as_num", "1"),
					resource.TestCheckResourceAttr("data.sdwan_eigrp_feature_template.test", "address_family.0.type", "ipv4"),
					resource.TestCheckResourceAttr("data.sdwan_eigrp_feature_template.test", "address_family.0.redistribute.0.protocol", "bgp"),
					resource.TestCheckResourceAttr("data.sdwan_eigrp_feature_template.test", "address_family.0.redistribute.0.route_policy", "1.2.3.4"),
					resource.TestCheckResourceAttr("data.sdwan_eigrp_feature_template.test", "address_family.0.network.0.prefix", "1.2.3.4/24"),
					resource.TestCheckResourceAttr("data.sdwan_eigrp_feature_template.test", "hello_interval", "5"),
					resource.TestCheckResourceAttr("data.sdwan_eigrp_feature_template.test", "hold_time", "15"),
					resource.TestCheckResourceAttr("data.sdwan_eigrp_feature_template.test", "route_policy_name", "example"),
					resource.TestCheckResourceAttr("data.sdwan_eigrp_feature_template.test", "filter", "false"),
					resource.TestCheckResourceAttr("data.sdwan_eigrp_feature_template.test", "authentication_type", "hmac-sha-256"),
					resource.TestCheckResourceAttr("data.sdwan_eigrp_feature_template.test", "hmac_authentication_key", "myAuthKey"),
					resource.TestCheckResourceAttr("data.sdwan_eigrp_feature_template.test", "interfaces.0.interface_name", "Ethernet"),
					resource.TestCheckResourceAttr("data.sdwan_eigrp_feature_template.test", "interfaces.0.shutdown", "false"),
					resource.TestCheckResourceAttr("data.sdwan_eigrp_feature_template.test", "interfaces.0.summary_address.0.prefix", "1.2.3.4/24"),
				),
			},
		},
	})
}

const testAccDataSourceSdwanEigrpFeatureTemplateConfig = `

resource "sdwan_eigrp_feature_template" "test" {
  name = "TF_TEST_MIN"
  description = "Terraform integration test"
  device_types = ["vedge-C8000V"]
  as_num = 1
  address_family = [{
    type = "ipv4"
	redistribute = [{
		protocol = "bgp"
		route_policy = "1.2.3.4"
	}]
	network = [{
		prefix = "1.2.3.4/24"
	}]
  }]
  hello_interval = 5
  hold_time = 15
  route_policy_name = "example"
  filter = false
  authentication_type = "hmac-sha-256"
  hmac_authentication_key = "myAuthKey"
  interfaces = [{
    interface_name = "Ethernet"
    shutdown = false
	summary_address = [{
		prefix = "1.2.3.4/24"
	}]
  }]
}

data "sdwan_eigrp_feature_template" "test" {
  id = sdwan_eigrp_feature_template.test.id
}
`
