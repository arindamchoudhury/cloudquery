// Code generated by codegen; DO NOT EDIT.

package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
  {{- range .}}
  "github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/{{.Service}}"
  {{- end}}
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "development"
)

func Plugin() *plugins.SourcePlugin {
	return plugins.NewSourcePlugin(
		"k8s",
		Version,
		[]*schema.Table{
      {{- range .}}
      {{.Service}}.{{.SubService | ToCamel}}(),
      {{- end}}
		},
		client.Configure,
	)
}