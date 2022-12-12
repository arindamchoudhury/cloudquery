// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb"

func Armmariadb() []*Table {
	tables := []*Table{
		{
			NewFunc:        armmariadb.NewServersClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.DBforMariaDB/servers",
			Namespace:      "Microsoft.DBforMariaDB",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_DBforMariaDB)`,
			Pager:          `NewListPager`,
			ResponseStruct: "ServersClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armmariadb())
}