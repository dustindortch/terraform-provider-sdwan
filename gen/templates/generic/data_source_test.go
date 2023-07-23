//go:build ignore
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

{{if .ExcludeTest}}//go:build testAll{{end}}

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceSdwan{{camelCase .Name}}(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwan{{camelCase .Name}}Config,
				Check: resource.ComposeTestCheckFunc(
					{{- $name := .Name }}
					{{- range  .Attributes}}
					{{- if and (not .WriteOnly) (not .ExcludeTest) (not .TfOnly) (not .Value)}}
					{{- if eq .Type "List"}}
					{{- $list := .TfName }}
					{{- range  .Attributes}}
					{{- if and (not .WriteOnly) (not .ExcludeTest) (not .TfOnly) (not .Value)}}
					{{- if eq .Type "List"}}
					{{- $clist := .TfName }}
					{{- range  .Attributes}}
					{{- if and (not .WriteOnly) (not .ExcludeTest) (not .TfOnly) (not .Value)}}
					{{- if eq .Type "List"}}
					{{- $cclist := .TfName }}
					{{- range  .Attributes}}
					{{- if and (not .WriteOnly) (not .ExcludeTest) (not .TfOnly) (not .Value) (ne .Type "ListString")}}
					resource.TestCheckResourceAttr("data.sdwan_{{snakeCase $name}}.test", "{{$list}}.0.{{$clist}}.0.{{$cclist}}.0.{{.TfName}}", "{{.Example}}"),
					{{- end}}
					{{- end}}
					{{- else if ne .Type "ListString"}}
					resource.TestCheckResourceAttr("data.sdwan_{{snakeCase $name}}.test", "{{$list}}.0.{{$clist}}.0.{{.TfName}}", "{{.Example}}"),
					{{- end}}
					{{- end}}
					{{- end}}
					{{- else if ne .Type "ListString"}}
					resource.TestCheckResourceAttr("data.sdwan_{{snakeCase $name}}.test", "{{$list}}.0.{{.TfName}}", "{{.Example}}"),
					{{- end}}
					{{- end}}
					{{- end}}
					{{- else if ne .Type "ListString"}}
					resource.TestCheckResourceAttr("data.sdwan_{{snakeCase $name}}.test", "{{.TfName}}", "{{.Example}}"),
					{{- end}}
					{{- end}}
					{{- end}}
				),
			},
		},
	})
}

const testAccDataSourceSdwan{{camelCase .Name}}Config = `

resource "sdwan_{{snakeCase $name}}" "test" {
{{- range  .Attributes}}
{{- if and (not .ExcludeTest) (not .TfOnly) (not .Value)}}
{{- if eq .Type "List"}}
  {{.TfName}} = [{
    {{- range  .Attributes}}
    {{- if and (not .ExcludeTest) (not .TfOnly) (not .Value)}}
	{{- if eq .Type "List"}}
	{{.TfName}} = [{
		{{- range  .Attributes}}
		{{- if and (not .ExcludeTest) (not .TfOnly) (not .Value)}}
		{{- if eq .Type "List"}}
		{{.TfName}} = [{
			{{- range  .Attributes}}
			{{- if and (not .ExcludeTest) (not .TfOnly) (not .Value)}}
			{{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
			{{- end}}
			{{- end}}
		}]
		{{- else}}
		{{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
		{{- end}}
		{{- end}}
		{{- end}}
	}]
	{{- else}}
    {{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
    {{- end}}
	{{- end}}
	{{- end}}
  }]
{{- else}}
  {{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
{{- end}}
{{- end}}
{{- end}}
}

data "sdwan_{{snakeCase .Name}}" "test" {
  id = sdwan_{{snakeCase $name}}.test.id
}
`
