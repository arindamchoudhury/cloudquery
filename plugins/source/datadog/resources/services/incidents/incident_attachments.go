// Code generated by codegen; DO NOT EDIT.

package incidents

import (
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func IncidentAttachments() *schema.Table {
	return &schema.Table{
		Name:     "datadog_incident_attachments",
		Resolver: fetchIncidentAttachments,
		Columns: []schema.Column{
			{
				Name:     "account_name",
				Type:     schema.TypeString,
				Resolver: client.ResolveAccountName,
			},
			{
				Name:     "attributes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Attributes"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "relationships",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Relationships"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},
	}
}