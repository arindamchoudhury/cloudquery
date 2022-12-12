// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase"

func Armtestbase() []*Table {
	tables := []*Table{
		{
			NewFunc:        armtestbase.NewSKUsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.TestBase/skus",
			Namespace:      "Microsoft.TestBase",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_TestBase)`,
			Pager:          `NewListPager`,
			ResponseStruct: "SKUsClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armtestbase())
}