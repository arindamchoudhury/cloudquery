// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"

func init() {
	tables := []Table{
		{
			Service:        "armappservice",
			Name:           "certificate_orders",
			Struct:         &armappservice.CertificateOrder{},
			ResponseStruct: &armappservice.CertificateOrdersClientListResponse{},
			Client:         &armappservice.CertificateOrdersClient{},
			ListFunc:       (&armappservice.CertificateOrdersClient{}).NewListPager,
			NewFunc:        armappservice.NewCertificateOrdersClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.CertificateRegistration/certificateOrders",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_CertificateRegistration)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armappservice",
			Name:           "certificates",
			Struct:         &armappservice.AppCertificate{},
			ResponseStruct: &armappservice.CertificatesClientListResponse{},
			Client:         &armappservice.CertificatesClient{},
			ListFunc:       (&armappservice.CertificatesClient{}).NewListPager,
			NewFunc:        armappservice.NewCertificatesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Web/certificates",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Web)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armappservice",
			Name:           "deleted_web_apps",
			Struct:         &armappservice.DeletedSite{},
			ResponseStruct: &armappservice.DeletedWebAppsClientListResponse{},
			Client:         &armappservice.DeletedWebAppsClient{},
			ListFunc:       (&armappservice.DeletedWebAppsClient{}).NewListPager,
			NewFunc:        armappservice.NewDeletedWebAppsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Web/deletedSites",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Web)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armappservice",
			Name:           "domains",
			Struct:         &armappservice.Domain{},
			ResponseStruct: &armappservice.DomainsClientListResponse{},
			Client:         &armappservice.DomainsClient{},
			ListFunc:       (&armappservice.DomainsClient{}).NewListPager,
			NewFunc:        armappservice.NewDomainsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.DomainRegistration/domains",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_DomainRegistration)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armappservice",
			Name:           "environments",
			Struct:         &armappservice.EnvironmentResource{},
			ResponseStruct: &armappservice.EnvironmentsClientListResponse{},
			Client:         &armappservice.EnvironmentsClient{},
			ListFunc:       (&armappservice.EnvironmentsClient{}).NewListPager,
			NewFunc:        armappservice.NewEnvironmentsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Web/hostingEnvironments",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Web)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armappservice",
			Name:           "plans",
			Struct:         &armappservice.Plan{},
			ResponseStruct: &armappservice.PlansClientListResponse{},
			Client:         &armappservice.PlansClient{},
			ListFunc:       (&armappservice.PlansClient{}).NewListPager,
			NewFunc:        armappservice.NewPlansClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Web/serverfarms",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Web)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armappservice",
			Name:           "recommendations",
			Struct:         &armappservice.Recommendation{},
			ResponseStruct: &armappservice.RecommendationsClientListResponse{},
			Client:         &armappservice.RecommendationsClient{},
			ListFunc:       (&armappservice.RecommendationsClient{}).NewListPager,
			NewFunc:        armappservice.NewRecommendationsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Web/recommendations",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Web)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armappservice",
			Name:           "resource_health_metadata",
			Struct:         &armappservice.ResourceHealthMetadata{},
			ResponseStruct: &armappservice.ResourceHealthMetadataClientListResponse{},
			Client:         &armappservice.ResourceHealthMetadataClient{},
			ListFunc:       (&armappservice.ResourceHealthMetadataClient{}).NewListPager,
			NewFunc:        armappservice.NewResourceHealthMetadataClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Web/resourceHealthMetadata",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Web)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armappservice",
			Name:           "static_sites",
			Struct:         &armappservice.StaticSiteARMResource{},
			ResponseStruct: &armappservice.StaticSitesClientListResponse{},
			Client:         &armappservice.StaticSitesClient{},
			ListFunc:       (&armappservice.StaticSitesClient{}).NewListPager,
			NewFunc:        armappservice.NewStaticSitesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Web/staticSites",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Web)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armappservice",
			Name:           "top_level_domains",
			Struct:         &armappservice.TopLevelDomain{},
			ResponseStruct: &armappservice.TopLevelDomainsClientListResponse{},
			Client:         &armappservice.TopLevelDomainsClient{},
			ListFunc:       (&armappservice.TopLevelDomainsClient{}).NewListPager,
			NewFunc:        armappservice.NewTopLevelDomainsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.DomainRegistration/topLevelDomains",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_DomainRegistration)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armappservice",
			Name:           "web_apps",
			Struct:         &armappservice.Site{},
			ResponseStruct: &armappservice.WebAppsClientListResponse{},
			Client:         &armappservice.WebAppsClient{},
			ListFunc:       (&armappservice.WebAppsClient{}).NewListPager,
			NewFunc:        armappservice.NewWebAppsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Web/sites",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Web)`,
			ExtraColumns:   DefaultExtraColumns,
		},
	}
	Tables = append(Tables, tables...)
}