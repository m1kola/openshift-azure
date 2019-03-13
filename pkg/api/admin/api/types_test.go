package api

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-test/deep"

	"github.com/openshift/openshift-azure/pkg/util/structtags"
	"github.com/openshift/openshift-azure/test/util/populate"
)

var marshalled = []byte(`{
	"plan": {
		"name": "Plan.Name",
		"product": "Plan.Product",
		"promotionCode": "Plan.PromotionCode",
		"publisher": "Plan.Publisher"
	},
	"properties": {
		"provisioningState": "Properties.ProvisioningState",
		"openShiftVersion": "Properties.OpenShiftVersion",
		"clusterVersion": "Properties.ClusterVersion",
		"publicHostname": "Properties.PublicHostname",
		"fqdn": "Properties.FQDN",
		"networkProfile": {
			"vnetCidr": "Properties.NetworkProfile.VnetCIDR",
			"vnetId": "Properties.NetworkProfile.VnetID",
			"peerVnetId": "Properties.NetworkProfile.PeerVnetID"
		},
		"routerProfiles": [
			{
				"name": "Properties.RouterProfiles[0].Name",
				"publicSubdomain": "Properties.RouterProfiles[0].PublicSubdomain",
				"fqdn": "Properties.RouterProfiles[0].FQDN"
			}
		],
		"masterPoolProfile": {
			"count": 1,
			"vmSize": "Properties.MasterPoolProfile.VMSize",
			"subnetCidr": "Properties.MasterPoolProfile.SubnetCIDR"
		},
		"agentPoolProfiles": [
			{
				"name": "Properties.AgentPoolProfiles[0].Name",
				"count": 1,
				"vmSize": "Properties.AgentPoolProfiles[0].VMSize",
				"subnetCidr": "Properties.AgentPoolProfiles[0].SubnetCIDR",
				"osType": "Properties.AgentPoolProfiles[0].OSType",
				"role": "Properties.AgentPoolProfiles[0].Role"
			}
		],
		"authProfile": {
			"identityProviders": [
				{
					"name": "Properties.AuthProfile.IdentityProviders[0].Name",
					"provider": {
						"kind": "AADIdentityProvider",
						"clientId": "Properties.AuthProfile.IdentityProviders[0].Provider.ClientID",
						"tenantId": "Properties.AuthProfile.IdentityProviders[0].Provider.TenantID",
						"customerAdminGroupId": "Properties.AuthProfile.IdentityProviders[0].Provider.CustomerAdminGroupID"
					}
				}
			]
		}
	},
	"id": "ID",
	"name": "Name",
	"type": "Type",
	"location": "Location",
	"tags": {
		"Tags.key": "Tags.val"
	},
	"config": {
		"pluginVersion": "Config.PluginVersion",
		"componentLogLevel": {
			"apiServer": 1,
			"controllerManager": 1,
			"node": 1
		},
		"imageOffer": "Config.ImageOffer",
		"imagePublisher": "Config.ImagePublisher",
		"imageSku": "Config.ImageSKU",
		"imageVersion": "Config.ImageVersion",
		"sshSourceAddressPrefixes": [
			"Config.SSHSourceAddressPrefixes[0]"
		],
		"configStorageAccount": "Config.ConfigStorageAccount",
		"registryStorageAccount": "Config.RegistryStorageAccount",
		"azureFileStorageAccount": "Config.AzureFileStorageAccount",
		"certificates": {
			"etcdCa": {
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2akNDQVpTZ0F3SUJBZ0lKQVBaMC8ydDJqYXJ5TUEwR0NTcUdTSWIzRFFFQkN3VUFNQlV4RXpBUkJnTlYKQkFNVENtdDFZbVZ5Ym1WMFpYTXdIaGNOTVRneE1URXlNVGN5TWpJeFdoY05NakF4TVRFeU1UY3lNakl4V2pBVgpNUk13RVFZRFZRUURFd3ByZFdKbGNtNWxkR1Z6TUZ3d0RRWUpLb1pJaHZjTkFRRUJCUUFEU3dBd1NBSkJBTko2CnFoY1ppQStLbFFERE5WakE2NE1USW0rd1hhVlJ2ekNmcG5vcmtGNDlSRzFWL1ZuZlZCTmVPNTBvb3E1ZGNpcEkKTjNvOG5lcGNPUEFOWG5Ob1ZFTUNBd0VBQWFPQnhqQ0J3ekFPQmdOVkhROEJBZjhFQkFNQ0JhQXdEQVlEVlIwVApBUUgvQkFJd0FEQ0JvZ1lEVlIwUkJJR2FNSUdYZ2dwcmRXSmxjbTVsZEdWemdnMXRZWE4wWlhJdE1EQXdNREF3CmdnMXRZWE4wWlhJdE1EQXdNREF4Z2cxdFlYTjBaWEl0TURBd01EQXlnZ3ByZFdKbGNtNWxkR1Z6Z2hKcmRXSmwKY201bGRHVnpMbVJsWm1GMWJIU0NGbXQxWW1WeWJtVjBaWE11WkdWbVlYVnNkQzV6ZG1PQ0pHdDFZbVZ5Ym1WMApaWE11WkdWbVlYVnNkQzV6ZG1NdVkyeDFjM1JsY2k1c2IyTmhiREFOQmdrcWhraUc5dzBCQVFzRkFBTkJBTkF2CjUwaXJlRHRqUkRiMVZydjJEcmRGZXhrT2hhZzNJM3dEVWlPb0loYjVuTkNNRjdnMEF1VEpOWkhFUXFOWDNYb1gKYkZpRkExdUxLZEZMc1B1T1dTUT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
			},
			"ca": {
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2akNDQVpTZ0F3SUJBZ0lKQVBaMC8ydDJqYXJ5TUEwR0NTcUdTSWIzRFFFQkN3VUFNQlV4RXpBUkJnTlYKQkFNVENtdDFZbVZ5Ym1WMFpYTXdIaGNOTVRneE1URXlNVGN5TWpJeFdoY05NakF4TVRFeU1UY3lNakl4V2pBVgpNUk13RVFZRFZRUURFd3ByZFdKbGNtNWxkR1Z6TUZ3d0RRWUpLb1pJaHZjTkFRRUJCUUFEU3dBd1NBSkJBTko2CnFoY1ppQStLbFFERE5WakE2NE1USW0rd1hhVlJ2ekNmcG5vcmtGNDlSRzFWL1ZuZlZCTmVPNTBvb3E1ZGNpcEkKTjNvOG5lcGNPUEFOWG5Ob1ZFTUNBd0VBQWFPQnhqQ0J3ekFPQmdOVkhROEJBZjhFQkFNQ0JhQXdEQVlEVlIwVApBUUgvQkFJd0FEQ0JvZ1lEVlIwUkJJR2FNSUdYZ2dwcmRXSmxjbTVsZEdWemdnMXRZWE4wWlhJdE1EQXdNREF3CmdnMXRZWE4wWlhJdE1EQXdNREF4Z2cxdFlYTjBaWEl0TURBd01EQXlnZ3ByZFdKbGNtNWxkR1Z6Z2hKcmRXSmwKY201bGRHVnpMbVJsWm1GMWJIU0NGbXQxWW1WeWJtVjBaWE11WkdWbVlYVnNkQzV6ZG1PQ0pHdDFZbVZ5Ym1WMApaWE11WkdWbVlYVnNkQzV6ZG1NdVkyeDFjM1JsY2k1c2IyTmhiREFOQmdrcWhraUc5dzBCQVFzRkFBTkJBTkF2CjUwaXJlRHRqUkRiMVZydjJEcmRGZXhrT2hhZzNJM3dEVWlPb0loYjVuTkNNRjdnMEF1VEpOWkhFUXFOWDNYb1gKYkZpRkExdUxLZEZMc1B1T1dTUT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
			},
			"frontProxyCa": {
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2akNDQVpTZ0F3SUJBZ0lKQVBaMC8ydDJqYXJ5TUEwR0NTcUdTSWIzRFFFQkN3VUFNQlV4RXpBUkJnTlYKQkFNVENtdDFZbVZ5Ym1WMFpYTXdIaGNOTVRneE1URXlNVGN5TWpJeFdoY05NakF4TVRFeU1UY3lNakl4V2pBVgpNUk13RVFZRFZRUURFd3ByZFdKbGNtNWxkR1Z6TUZ3d0RRWUpLb1pJaHZjTkFRRUJCUUFEU3dBd1NBSkJBTko2CnFoY1ppQStLbFFERE5WakE2NE1USW0rd1hhVlJ2ekNmcG5vcmtGNDlSRzFWL1ZuZlZCTmVPNTBvb3E1ZGNpcEkKTjNvOG5lcGNPUEFOWG5Ob1ZFTUNBd0VBQWFPQnhqQ0J3ekFPQmdOVkhROEJBZjhFQkFNQ0JhQXdEQVlEVlIwVApBUUgvQkFJd0FEQ0JvZ1lEVlIwUkJJR2FNSUdYZ2dwcmRXSmxjbTVsZEdWemdnMXRZWE4wWlhJdE1EQXdNREF3CmdnMXRZWE4wWlhJdE1EQXdNREF4Z2cxdFlYTjBaWEl0TURBd01EQXlnZ3ByZFdKbGNtNWxkR1Z6Z2hKcmRXSmwKY201bGRHVnpMbVJsWm1GMWJIU0NGbXQxWW1WeWJtVjBaWE11WkdWbVlYVnNkQzV6ZG1PQ0pHdDFZbVZ5Ym1WMApaWE11WkdWbVlYVnNkQzV6ZG1NdVkyeDFjM1JsY2k1c2IyTmhiREFOQmdrcWhraUc5dzBCQVFzRkFBTkJBTkF2CjUwaXJlRHRqUkRiMVZydjJEcmRGZXhrT2hhZzNJM3dEVWlPb0loYjVuTkNNRjdnMEF1VEpOWkhFUXFOWDNYb1gKYkZpRkExdUxLZEZMc1B1T1dTUT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
			},
			"serviceSigningCa": {
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2akNDQVpTZ0F3SUJBZ0lKQVBaMC8ydDJqYXJ5TUEwR0NTcUdTSWIzRFFFQkN3VUFNQlV4RXpBUkJnTlYKQkFNVENtdDFZbVZ5Ym1WMFpYTXdIaGNOTVRneE1URXlNVGN5TWpJeFdoY05NakF4TVRFeU1UY3lNakl4V2pBVgpNUk13RVFZRFZRUURFd3ByZFdKbGNtNWxkR1Z6TUZ3d0RRWUpLb1pJaHZjTkFRRUJCUUFEU3dBd1NBSkJBTko2CnFoY1ppQStLbFFERE5WakE2NE1USW0rd1hhVlJ2ekNmcG5vcmtGNDlSRzFWL1ZuZlZCTmVPNTBvb3E1ZGNpcEkKTjNvOG5lcGNPUEFOWG5Ob1ZFTUNBd0VBQWFPQnhqQ0J3ekFPQmdOVkhROEJBZjhFQkFNQ0JhQXdEQVlEVlIwVApBUUgvQkFJd0FEQ0JvZ1lEVlIwUkJJR2FNSUdYZ2dwcmRXSmxjbTVsZEdWemdnMXRZWE4wWlhJdE1EQXdNREF3CmdnMXRZWE4wWlhJdE1EQXdNREF4Z2cxdFlYTjBaWEl0TURBd01EQXlnZ3ByZFdKbGNtNWxkR1Z6Z2hKcmRXSmwKY201bGRHVnpMbVJsWm1GMWJIU0NGbXQxWW1WeWJtVjBaWE11WkdWbVlYVnNkQzV6ZG1PQ0pHdDFZbVZ5Ym1WMApaWE11WkdWbVlYVnNkQzV6ZG1NdVkyeDFjM1JsY2k1c2IyTmhiREFOQmdrcWhraUc5dzBCQVFzRkFBTkJBTkF2CjUwaXJlRHRqUkRiMVZydjJEcmRGZXhrT2hhZzNJM3dEVWlPb0loYjVuTkNNRjdnMEF1VEpOWkhFUXFOWDNYb1gKYkZpRkExdUxLZEZMc1B1T1dTUT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
			},
			"serviceCatalogCa": {
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2akNDQVpTZ0F3SUJBZ0lKQVBaMC8ydDJqYXJ5TUEwR0NTcUdTSWIzRFFFQkN3VUFNQlV4RXpBUkJnTlYKQkFNVENtdDFZbVZ5Ym1WMFpYTXdIaGNOTVRneE1URXlNVGN5TWpJeFdoY05NakF4TVRFeU1UY3lNakl4V2pBVgpNUk13RVFZRFZRUURFd3ByZFdKbGNtNWxkR1Z6TUZ3d0RRWUpLb1pJaHZjTkFRRUJCUUFEU3dBd1NBSkJBTko2CnFoY1ppQStLbFFERE5WakE2NE1USW0rd1hhVlJ2ekNmcG5vcmtGNDlSRzFWL1ZuZlZCTmVPNTBvb3E1ZGNpcEkKTjNvOG5lcGNPUEFOWG5Ob1ZFTUNBd0VBQWFPQnhqQ0J3ekFPQmdOVkhROEJBZjhFQkFNQ0JhQXdEQVlEVlIwVApBUUgvQkFJd0FEQ0JvZ1lEVlIwUkJJR2FNSUdYZ2dwcmRXSmxjbTVsZEdWemdnMXRZWE4wWlhJdE1EQXdNREF3CmdnMXRZWE4wWlhJdE1EQXdNREF4Z2cxdFlYTjBaWEl0TURBd01EQXlnZ3ByZFdKbGNtNWxkR1Z6Z2hKcmRXSmwKY201bGRHVnpMbVJsWm1GMWJIU0NGbXQxWW1WeWJtVjBaWE11WkdWbVlYVnNkQzV6ZG1PQ0pHdDFZbVZ5Ym1WMApaWE11WkdWbVlYVnNkQzV6ZG1NdVkyeDFjM1JsY2k1c2IyTmhiREFOQmdrcWhraUc5dzBCQVFzRkFBTkJBTkF2CjUwaXJlRHRqUkRiMVZydjJEcmRGZXhrT2hhZzNJM3dEVWlPb0loYjVuTkNNRjdnMEF1VEpOWkhFUXFOWDNYb1gKYkZpRkExdUxLZEZMc1B1T1dTUT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
			},
			"etcdServer": {
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2akNDQVpTZ0F3SUJBZ0lKQVBaMC8ydDJqYXJ5TUEwR0NTcUdTSWIzRFFFQkN3VUFNQlV4RXpBUkJnTlYKQkFNVENtdDFZbVZ5Ym1WMFpYTXdIaGNOTVRneE1URXlNVGN5TWpJeFdoY05NakF4TVRFeU1UY3lNakl4V2pBVgpNUk13RVFZRFZRUURFd3ByZFdKbGNtNWxkR1Z6TUZ3d0RRWUpLb1pJaHZjTkFRRUJCUUFEU3dBd1NBSkJBTko2CnFoY1ppQStLbFFERE5WakE2NE1USW0rd1hhVlJ2ekNmcG5vcmtGNDlSRzFWL1ZuZlZCTmVPNTBvb3E1ZGNpcEkKTjNvOG5lcGNPUEFOWG5Ob1ZFTUNBd0VBQWFPQnhqQ0J3ekFPQmdOVkhROEJBZjhFQkFNQ0JhQXdEQVlEVlIwVApBUUgvQkFJd0FEQ0JvZ1lEVlIwUkJJR2FNSUdYZ2dwcmRXSmxjbTVsZEdWemdnMXRZWE4wWlhJdE1EQXdNREF3CmdnMXRZWE4wWlhJdE1EQXdNREF4Z2cxdFlYTjBaWEl0TURBd01EQXlnZ3ByZFdKbGNtNWxkR1Z6Z2hKcmRXSmwKY201bGRHVnpMbVJsWm1GMWJIU0NGbXQxWW1WeWJtVjBaWE11WkdWbVlYVnNkQzV6ZG1PQ0pHdDFZbVZ5Ym1WMApaWE11WkdWbVlYVnNkQzV6ZG1NdVkyeDFjM1JsY2k1c2IyTmhiREFOQmdrcWhraUc5dzBCQVFzRkFBTkJBTkF2CjUwaXJlRHRqUkRiMVZydjJEcmRGZXhrT2hhZzNJM3dEVWlPb0loYjVuTkNNRjdnMEF1VEpOWkhFUXFOWDNYb1gKYkZpRkExdUxLZEZMc1B1T1dTUT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
			},
			"etcdPeer": {
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2akNDQVpTZ0F3SUJBZ0lKQVBaMC8ydDJqYXJ5TUEwR0NTcUdTSWIzRFFFQkN3VUFNQlV4RXpBUkJnTlYKQkFNVENtdDFZbVZ5Ym1WMFpYTXdIaGNOTVRneE1URXlNVGN5TWpJeFdoY05NakF4TVRFeU1UY3lNakl4V2pBVgpNUk13RVFZRFZRUURFd3ByZFdKbGNtNWxkR1Z6TUZ3d0RRWUpLb1pJaHZjTkFRRUJCUUFEU3dBd1NBSkJBTko2CnFoY1ppQStLbFFERE5WakE2NE1USW0rd1hhVlJ2ekNmcG5vcmtGNDlSRzFWL1ZuZlZCTmVPNTBvb3E1ZGNpcEkKTjNvOG5lcGNPUEFOWG5Ob1ZFTUNBd0VBQWFPQnhqQ0J3ekFPQmdOVkhROEJBZjhFQkFNQ0JhQXdEQVlEVlIwVApBUUgvQkFJd0FEQ0JvZ1lEVlIwUkJJR2FNSUdYZ2dwcmRXSmxjbTVsZEdWemdnMXRZWE4wWlhJdE1EQXdNREF3CmdnMXRZWE4wWlhJdE1EQXdNREF4Z2cxdFlYTjBaWEl0TURBd01EQXlnZ3ByZFdKbGNtNWxkR1Z6Z2hKcmRXSmwKY201bGRHVnpMbVJsWm1GMWJIU0NGbXQxWW1WeWJtVjBaWE11WkdWbVlYVnNkQzV6ZG1PQ0pHdDFZbVZ5Ym1WMApaWE11WkdWbVlYVnNkQzV6ZG1NdVkyeDFjM1JsY2k1c2IyTmhiREFOQmdrcWhraUc5dzBCQVFzRkFBTkJBTkF2CjUwaXJlRHRqUkRiMVZydjJEcmRGZXhrT2hhZzNJM3dEVWlPb0loYjVuTkNNRjdnMEF1VEpOWkhFUXFOWDNYb1gKYkZpRkExdUxLZEZMc1B1T1dTUT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
			},
			"etcdClient": {
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2akNDQVpTZ0F3SUJBZ0lKQVBaMC8ydDJqYXJ5TUEwR0NTcUdTSWIzRFFFQkN3VUFNQlV4RXpBUkJnTlYKQkFNVENtdDFZbVZ5Ym1WMFpYTXdIaGNOTVRneE1URXlNVGN5TWpJeFdoY05NakF4TVRFeU1UY3lNakl4V2pBVgpNUk13RVFZRFZRUURFd3ByZFdKbGNtNWxkR1Z6TUZ3d0RRWUpLb1pJaHZjTkFRRUJCUUFEU3dBd1NBSkJBTko2CnFoY1ppQStLbFFERE5WakE2NE1USW0rd1hhVlJ2ekNmcG5vcmtGNDlSRzFWL1ZuZlZCTmVPNTBvb3E1ZGNpcEkKTjNvOG5lcGNPUEFOWG5Ob1ZFTUNBd0VBQWFPQnhqQ0J3ekFPQmdOVkhROEJBZjhFQkFNQ0JhQXdEQVlEVlIwVApBUUgvQkFJd0FEQ0JvZ1lEVlIwUkJJR2FNSUdYZ2dwcmRXSmxjbTVsZEdWemdnMXRZWE4wWlhJdE1EQXdNREF3CmdnMXRZWE4wWlhJdE1EQXdNREF4Z2cxdFlYTjBaWEl0TURBd01EQXlnZ3ByZFdKbGNtNWxkR1Z6Z2hKcmRXSmwKY201bGRHVnpMbVJsWm1GMWJIU0NGbXQxWW1WeWJtVjBaWE11WkdWbVlYVnNkQzV6ZG1PQ0pHdDFZbVZ5Ym1WMApaWE11WkdWbVlYVnNkQzV6ZG1NdVkyeDFjM1JsY2k1c2IyTmhiREFOQmdrcWhraUc5dzBCQVFzRkFBTkJBTkF2CjUwaXJlRHRqUkRiMVZydjJEcmRGZXhrT2hhZzNJM3dEVWlPb0loYjVuTkNNRjdnMEF1VEpOWkhFUXFOWDNYb1gKYkZpRkExdUxLZEZMc1B1T1dTUT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
			},
			"masterServer": {
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2akNDQVpTZ0F3SUJBZ0lKQVBaMC8ydDJqYXJ5TUEwR0NTcUdTSWIzRFFFQkN3VUFNQlV4RXpBUkJnTlYKQkFNVENtdDFZbVZ5Ym1WMFpYTXdIaGNOTVRneE1URXlNVGN5TWpJeFdoY05NakF4TVRFeU1UY3lNakl4V2pBVgpNUk13RVFZRFZRUURFd3ByZFdKbGNtNWxkR1Z6TUZ3d0RRWUpLb1pJaHZjTkFRRUJCUUFEU3dBd1NBSkJBTko2CnFoY1ppQStLbFFERE5WakE2NE1USW0rd1hhVlJ2ekNmcG5vcmtGNDlSRzFWL1ZuZlZCTmVPNTBvb3E1ZGNpcEkKTjNvOG5lcGNPUEFOWG5Ob1ZFTUNBd0VBQWFPQnhqQ0J3ekFPQmdOVkhROEJBZjhFQkFNQ0JhQXdEQVlEVlIwVApBUUgvQkFJd0FEQ0JvZ1lEVlIwUkJJR2FNSUdYZ2dwcmRXSmxjbTVsZEdWemdnMXRZWE4wWlhJdE1EQXdNREF3CmdnMXRZWE4wWlhJdE1EQXdNREF4Z2cxdFlYTjBaWEl0TURBd01EQXlnZ3ByZFdKbGNtNWxkR1Z6Z2hKcmRXSmwKY201bGRHVnpMbVJsWm1GMWJIU0NGbXQxWW1WeWJtVjBaWE11WkdWbVlYVnNkQzV6ZG1PQ0pHdDFZbVZ5Ym1WMApaWE11WkdWbVlYVnNkQzV6ZG1NdVkyeDFjM1JsY2k1c2IyTmhiREFOQmdrcWhraUc5dzBCQVFzRkFBTkJBTkF2CjUwaXJlRHRqUkRiMVZydjJEcmRGZXhrT2hhZzNJM3dEVWlPb0loYjVuTkNNRjdnMEF1VEpOWkhFUXFOWDNYb1gKYkZpRkExdUxLZEZMc1B1T1dTUT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
			},
			"openShiftConsole": {
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2akNDQVpTZ0F3SUJBZ0lKQVBaMC8ydDJqYXJ5TUEwR0NTcUdTSWIzRFFFQkN3VUFNQlV4RXpBUkJnTlYKQkFNVENtdDFZbVZ5Ym1WMFpYTXdIaGNOTVRneE1URXlNVGN5TWpJeFdoY05NakF4TVRFeU1UY3lNakl4V2pBVgpNUk13RVFZRFZRUURFd3ByZFdKbGNtNWxkR1Z6TUZ3d0RRWUpLb1pJaHZjTkFRRUJCUUFEU3dBd1NBSkJBTko2CnFoY1ppQStLbFFERE5WakE2NE1USW0rd1hhVlJ2ekNmcG5vcmtGNDlSRzFWL1ZuZlZCTmVPNTBvb3E1ZGNpcEkKTjNvOG5lcGNPUEFOWG5Ob1ZFTUNBd0VBQWFPQnhqQ0J3ekFPQmdOVkhROEJBZjhFQkFNQ0JhQXdEQVlEVlIwVApBUUgvQkFJd0FEQ0JvZ1lEVlIwUkJJR2FNSUdYZ2dwcmRXSmxjbTVsZEdWemdnMXRZWE4wWlhJdE1EQXdNREF3CmdnMXRZWE4wWlhJdE1EQXdNREF4Z2cxdFlYTjBaWEl0TURBd01EQXlnZ3ByZFdKbGNtNWxkR1Z6Z2hKcmRXSmwKY201bGRHVnpMbVJsWm1GMWJIU0NGbXQxWW1WeWJtVjBaWE11WkdWbVlYVnNkQzV6ZG1PQ0pHdDFZbVZ5Ym1WMApaWE11WkdWbVlYVnNkQzV6ZG1NdVkyeDFjM1JsY2k1c2IyTmhiREFOQmdrcWhraUc5dzBCQVFzRkFBTkJBTkF2CjUwaXJlRHRqUkRiMVZydjJEcmRGZXhrT2hhZzNJM3dEVWlPb0loYjVuTkNNRjdnMEF1VEpOWkhFUXFOWDNYb1gKYkZpRkExdUxLZEZMc1B1T1dTUT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
			},
			"admin": {
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2akNDQVpTZ0F3SUJBZ0lKQVBaMC8ydDJqYXJ5TUEwR0NTcUdTSWIzRFFFQkN3VUFNQlV4RXpBUkJnTlYKQkFNVENtdDFZbVZ5Ym1WMFpYTXdIaGNOTVRneE1URXlNVGN5TWpJeFdoY05NakF4TVRFeU1UY3lNakl4V2pBVgpNUk13RVFZRFZRUURFd3ByZFdKbGNtNWxkR1Z6TUZ3d0RRWUpLb1pJaHZjTkFRRUJCUUFEU3dBd1NBSkJBTko2CnFoY1ppQStLbFFERE5WakE2NE1USW0rd1hhVlJ2ekNmcG5vcmtGNDlSRzFWL1ZuZlZCTmVPNTBvb3E1ZGNpcEkKTjNvOG5lcGNPUEFOWG5Ob1ZFTUNBd0VBQWFPQnhqQ0J3ekFPQmdOVkhROEJBZjhFQkFNQ0JhQXdEQVlEVlIwVApBUUgvQkFJd0FEQ0JvZ1lEVlIwUkJJR2FNSUdYZ2dwcmRXSmxjbTVsZEdWemdnMXRZWE4wWlhJdE1EQXdNREF3CmdnMXRZWE4wWlhJdE1EQXdNREF4Z2cxdFlYTjBaWEl0TURBd01EQXlnZ3ByZFdKbGNtNWxkR1Z6Z2hKcmRXSmwKY201bGRHVnpMbVJsWm1GMWJIU0NGbXQxWW1WeWJtVjBaWE11WkdWbVlYVnNkQzV6ZG1PQ0pHdDFZbVZ5Ym1WMApaWE11WkdWbVlYVnNkQzV6ZG1NdVkyeDFjM1JsY2k1c2IyTmhiREFOQmdrcWhraUc5dzBCQVFzRkFBTkJBTkF2CjUwaXJlRHRqUkRiMVZydjJEcmRGZXhrT2hhZzNJM3dEVWlPb0loYjVuTkNNRjdnMEF1VEpOWkhFUXFOWDNYb1gKYkZpRkExdUxLZEZMc1B1T1dTUT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
			},
			"aggregatorFrontProxy": {
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2akNDQVpTZ0F3SUJBZ0lKQVBaMC8ydDJqYXJ5TUEwR0NTcUdTSWIzRFFFQkN3VUFNQlV4RXpBUkJnTlYKQkFNVENtdDFZbVZ5Ym1WMFpYTXdIaGNOTVRneE1URXlNVGN5TWpJeFdoY05NakF4TVRFeU1UY3lNakl4V2pBVgpNUk13RVFZRFZRUURFd3ByZFdKbGNtNWxkR1Z6TUZ3d0RRWUpLb1pJaHZjTkFRRUJCUUFEU3dBd1NBSkJBTko2CnFoY1ppQStLbFFERE5WakE2NE1USW0rd1hhVlJ2ekNmcG5vcmtGNDlSRzFWL1ZuZlZCTmVPNTBvb3E1ZGNpcEkKTjNvOG5lcGNPUEFOWG5Ob1ZFTUNBd0VBQWFPQnhqQ0J3ekFPQmdOVkhROEJBZjhFQkFNQ0JhQXdEQVlEVlIwVApBUUgvQkFJd0FEQ0JvZ1lEVlIwUkJJR2FNSUdYZ2dwcmRXSmxjbTVsZEdWemdnMXRZWE4wWlhJdE1EQXdNREF3CmdnMXRZWE4wWlhJdE1EQXdNREF4Z2cxdFlYTjBaWEl0TURBd01EQXlnZ3ByZFdKbGNtNWxkR1Z6Z2hKcmRXSmwKY201bGRHVnpMbVJsWm1GMWJIU0NGbXQxWW1WeWJtVjBaWE11WkdWbVlYVnNkQzV6ZG1PQ0pHdDFZbVZ5Ym1WMApaWE11WkdWbVlYVnNkQzV6ZG1NdVkyeDFjM1JsY2k1c2IyTmhiREFOQmdrcWhraUc5dzBCQVFzRkFBTkJBTkF2CjUwaXJlRHRqUkRiMVZydjJEcmRGZXhrT2hhZzNJM3dEVWlPb0loYjVuTkNNRjdnMEF1VEpOWkhFUXFOWDNYb1gKYkZpRkExdUxLZEZMc1B1T1dTUT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
			},
			"masterKubeletClient": {
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2akNDQVpTZ0F3SUJBZ0lKQVBaMC8ydDJqYXJ5TUEwR0NTcUdTSWIzRFFFQkN3VUFNQlV4RXpBUkJnTlYKQkFNVENtdDFZbVZ5Ym1WMFpYTXdIaGNOTVRneE1URXlNVGN5TWpJeFdoY05NakF4TVRFeU1UY3lNakl4V2pBVgpNUk13RVFZRFZRUURFd3ByZFdKbGNtNWxkR1Z6TUZ3d0RRWUpLb1pJaHZjTkFRRUJCUUFEU3dBd1NBSkJBTko2CnFoY1ppQStLbFFERE5WakE2NE1USW0rd1hhVlJ2ekNmcG5vcmtGNDlSRzFWL1ZuZlZCTmVPNTBvb3E1ZGNpcEkKTjNvOG5lcGNPUEFOWG5Ob1ZFTUNBd0VBQWFPQnhqQ0J3ekFPQmdOVkhROEJBZjhFQkFNQ0JhQXdEQVlEVlIwVApBUUgvQkFJd0FEQ0JvZ1lEVlIwUkJJR2FNSUdYZ2dwcmRXSmxjbTVsZEdWemdnMXRZWE4wWlhJdE1EQXdNREF3CmdnMXRZWE4wWlhJdE1EQXdNREF4Z2cxdFlYTjBaWEl0TURBd01EQXlnZ3ByZFdKbGNtNWxkR1Z6Z2hKcmRXSmwKY201bGRHVnpMbVJsWm1GMWJIU0NGbXQxWW1WeWJtVjBaWE11WkdWbVlYVnNkQzV6ZG1PQ0pHdDFZbVZ5Ym1WMApaWE11WkdWbVlYVnNkQzV6ZG1NdVkyeDFjM1JsY2k1c2IyTmhiREFOQmdrcWhraUc5dzBCQVFzRkFBTkJBTkF2CjUwaXJlRHRqUkRiMVZydjJEcmRGZXhrT2hhZzNJM3dEVWlPb0loYjVuTkNNRjdnMEF1VEpOWkhFUXFOWDNYb1gKYkZpRkExdUxLZEZMc1B1T1dTUT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
			},
			"masterProxyClient": {
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2akNDQVpTZ0F3SUJBZ0lKQVBaMC8ydDJqYXJ5TUEwR0NTcUdTSWIzRFFFQkN3VUFNQlV4RXpBUkJnTlYKQkFNVENtdDFZbVZ5Ym1WMFpYTXdIaGNOTVRneE1URXlNVGN5TWpJeFdoY05NakF4TVRFeU1UY3lNakl4V2pBVgpNUk13RVFZRFZRUURFd3ByZFdKbGNtNWxkR1Z6TUZ3d0RRWUpLb1pJaHZjTkFRRUJCUUFEU3dBd1NBSkJBTko2CnFoY1ppQStLbFFERE5WakE2NE1USW0rd1hhVlJ2ekNmcG5vcmtGNDlSRzFWL1ZuZlZCTmVPNTBvb3E1ZGNpcEkKTjNvOG5lcGNPUEFOWG5Ob1ZFTUNBd0VBQWFPQnhqQ0J3ekFPQmdOVkhROEJBZjhFQkFNQ0JhQXdEQVlEVlIwVApBUUgvQkFJd0FEQ0JvZ1lEVlIwUkJJR2FNSUdYZ2dwcmRXSmxjbTVsZEdWemdnMXRZWE4wWlhJdE1EQXdNREF3CmdnMXRZWE4wWlhJdE1EQXdNREF4Z2cxdFlYTjBaWEl0TURBd01EQXlnZ3ByZFdKbGNtNWxkR1Z6Z2hKcmRXSmwKY201bGRHVnpMbVJsWm1GMWJIU0NGbXQxWW1WeWJtVjBaWE11WkdWbVlYVnNkQzV6ZG1PQ0pHdDFZbVZ5Ym1WMApaWE11WkdWbVlYVnNkQzV6ZG1NdVkyeDFjM1JsY2k1c2IyTmhiREFOQmdrcWhraUc5dzBCQVFzRkFBTkJBTkF2CjUwaXJlRHRqUkRiMVZydjJEcmRGZXhrT2hhZzNJM3dEVWlPb0loYjVuTkNNRjdnMEF1VEpOWkhFUXFOWDNYb1gKYkZpRkExdUxLZEZMc1B1T1dTUT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
			},
			"openShiftMaster": {
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2akNDQVpTZ0F3SUJBZ0lKQVBaMC8ydDJqYXJ5TUEwR0NTcUdTSWIzRFFFQkN3VUFNQlV4RXpBUkJnTlYKQkFNVENtdDFZbVZ5Ym1WMFpYTXdIaGNOTVRneE1URXlNVGN5TWpJeFdoY05NakF4TVRFeU1UY3lNakl4V2pBVgpNUk13RVFZRFZRUURFd3ByZFdKbGNtNWxkR1Z6TUZ3d0RRWUpLb1pJaHZjTkFRRUJCUUFEU3dBd1NBSkJBTko2CnFoY1ppQStLbFFERE5WakE2NE1USW0rd1hhVlJ2ekNmcG5vcmtGNDlSRzFWL1ZuZlZCTmVPNTBvb3E1ZGNpcEkKTjNvOG5lcGNPUEFOWG5Ob1ZFTUNBd0VBQWFPQnhqQ0J3ekFPQmdOVkhROEJBZjhFQkFNQ0JhQXdEQVlEVlIwVApBUUgvQkFJd0FEQ0JvZ1lEVlIwUkJJR2FNSUdYZ2dwcmRXSmxjbTVsZEdWemdnMXRZWE4wWlhJdE1EQXdNREF3CmdnMXRZWE4wWlhJdE1EQXdNREF4Z2cxdFlYTjBaWEl0TURBd01EQXlnZ3ByZFdKbGNtNWxkR1Z6Z2hKcmRXSmwKY201bGRHVnpMbVJsWm1GMWJIU0NGbXQxWW1WeWJtVjBaWE11WkdWbVlYVnNkQzV6ZG1PQ0pHdDFZbVZ5Ym1WMApaWE11WkdWbVlYVnNkQzV6ZG1NdVkyeDFjM1JsY2k1c2IyTmhiREFOQmdrcWhraUc5dzBCQVFzRkFBTkJBTkF2CjUwaXJlRHRqUkRiMVZydjJEcmRGZXhrT2hhZzNJM3dEVWlPb0loYjVuTkNNRjdnMEF1VEpOWkhFUXFOWDNYb1gKYkZpRkExdUxLZEZMc1B1T1dTUT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
			},
			"nodeBootstrap": {
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2akNDQVpTZ0F3SUJBZ0lKQVBaMC8ydDJqYXJ5TUEwR0NTcUdTSWIzRFFFQkN3VUFNQlV4RXpBUkJnTlYKQkFNVENtdDFZbVZ5Ym1WMFpYTXdIaGNOTVRneE1URXlNVGN5TWpJeFdoY05NakF4TVRFeU1UY3lNakl4V2pBVgpNUk13RVFZRFZRUURFd3ByZFdKbGNtNWxkR1Z6TUZ3d0RRWUpLb1pJaHZjTkFRRUJCUUFEU3dBd1NBSkJBTko2CnFoY1ppQStLbFFERE5WakE2NE1USW0rd1hhVlJ2ekNmcG5vcmtGNDlSRzFWL1ZuZlZCTmVPNTBvb3E1ZGNpcEkKTjNvOG5lcGNPUEFOWG5Ob1ZFTUNBd0VBQWFPQnhqQ0J3ekFPQmdOVkhROEJBZjhFQkFNQ0JhQXdEQVlEVlIwVApBUUgvQkFJd0FEQ0JvZ1lEVlIwUkJJR2FNSUdYZ2dwcmRXSmxjbTVsZEdWemdnMXRZWE4wWlhJdE1EQXdNREF3CmdnMXRZWE4wWlhJdE1EQXdNREF4Z2cxdFlYTjBaWEl0TURBd01EQXlnZ3ByZFdKbGNtNWxkR1Z6Z2hKcmRXSmwKY201bGRHVnpMbVJsWm1GMWJIU0NGbXQxWW1WeWJtVjBaWE11WkdWbVlYVnNkQzV6ZG1PQ0pHdDFZbVZ5Ym1WMApaWE11WkdWbVlYVnNkQzV6ZG1NdVkyeDFjM1JsY2k1c2IyTmhiREFOQmdrcWhraUc5dzBCQVFzRkFBTkJBTkF2CjUwaXJlRHRqUkRiMVZydjJEcmRGZXhrT2hhZzNJM3dEVWlPb0loYjVuTkNNRjdnMEF1VEpOWkhFUXFOWDNYb1gKYkZpRkExdUxLZEZMc1B1T1dTUT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
			},
			"registry": {
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2akNDQVpTZ0F3SUJBZ0lKQVBaMC8ydDJqYXJ5TUEwR0NTcUdTSWIzRFFFQkN3VUFNQlV4RXpBUkJnTlYKQkFNVENtdDFZbVZ5Ym1WMFpYTXdIaGNOTVRneE1URXlNVGN5TWpJeFdoY05NakF4TVRFeU1UY3lNakl4V2pBVgpNUk13RVFZRFZRUURFd3ByZFdKbGNtNWxkR1Z6TUZ3d0RRWUpLb1pJaHZjTkFRRUJCUUFEU3dBd1NBSkJBTko2CnFoY1ppQStLbFFERE5WakE2NE1USW0rd1hhVlJ2ekNmcG5vcmtGNDlSRzFWL1ZuZlZCTmVPNTBvb3E1ZGNpcEkKTjNvOG5lcGNPUEFOWG5Ob1ZFTUNBd0VBQWFPQnhqQ0J3ekFPQmdOVkhROEJBZjhFQkFNQ0JhQXdEQVlEVlIwVApBUUgvQkFJd0FEQ0JvZ1lEVlIwUkJJR2FNSUdYZ2dwcmRXSmxjbTVsZEdWemdnMXRZWE4wWlhJdE1EQXdNREF3CmdnMXRZWE4wWlhJdE1EQXdNREF4Z2cxdFlYTjBaWEl0TURBd01EQXlnZ3ByZFdKbGNtNWxkR1Z6Z2hKcmRXSmwKY201bGRHVnpMbVJsWm1GMWJIU0NGbXQxWW1WeWJtVjBaWE11WkdWbVlYVnNkQzV6ZG1PQ0pHdDFZbVZ5Ym1WMApaWE11WkdWbVlYVnNkQzV6ZG1NdVkyeDFjM1JsY2k1c2IyTmhiREFOQmdrcWhraUc5dzBCQVFzRkFBTkJBTkF2CjUwaXJlRHRqUkRiMVZydjJEcmRGZXhrT2hhZzNJM3dEVWlPb0loYjVuTkNNRjdnMEF1VEpOWkhFUXFOWDNYb1gKYkZpRkExdUxLZEZMc1B1T1dTUT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
			},
			"registryConsole": {
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2akNDQVpTZ0F3SUJBZ0lKQVBaMC8ydDJqYXJ5TUEwR0NTcUdTSWIzRFFFQkN3VUFNQlV4RXpBUkJnTlYKQkFNVENtdDFZbVZ5Ym1WMFpYTXdIaGNOTVRneE1URXlNVGN5TWpJeFdoY05NakF4TVRFeU1UY3lNakl4V2pBVgpNUk13RVFZRFZRUURFd3ByZFdKbGNtNWxkR1Z6TUZ3d0RRWUpLb1pJaHZjTkFRRUJCUUFEU3dBd1NBSkJBTko2CnFoY1ppQStLbFFERE5WakE2NE1USW0rd1hhVlJ2ekNmcG5vcmtGNDlSRzFWL1ZuZlZCTmVPNTBvb3E1ZGNpcEkKTjNvOG5lcGNPUEFOWG5Ob1ZFTUNBd0VBQWFPQnhqQ0J3ekFPQmdOVkhROEJBZjhFQkFNQ0JhQXdEQVlEVlIwVApBUUgvQkFJd0FEQ0JvZ1lEVlIwUkJJR2FNSUdYZ2dwcmRXSmxjbTVsZEdWemdnMXRZWE4wWlhJdE1EQXdNREF3CmdnMXRZWE4wWlhJdE1EQXdNREF4Z2cxdFlYTjBaWEl0TURBd01EQXlnZ3ByZFdKbGNtNWxkR1Z6Z2hKcmRXSmwKY201bGRHVnpMbVJsWm1GMWJIU0NGbXQxWW1WeWJtVjBaWE11WkdWbVlYVnNkQzV6ZG1PQ0pHdDFZbVZ5Ym1WMApaWE11WkdWbVlYVnNkQzV6ZG1NdVkyeDFjM1JsY2k1c2IyTmhiREFOQmdrcWhraUc5dzBCQVFzRkFBTkJBTkF2CjUwaXJlRHRqUkRiMVZydjJEcmRGZXhrT2hhZzNJM3dEVWlPb0loYjVuTkNNRjdnMEF1VEpOWkhFUXFOWDNYb1gKYkZpRkExdUxLZEZMc1B1T1dTUT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
			},
			"router": {
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2akNDQVpTZ0F3SUJBZ0lKQVBaMC8ydDJqYXJ5TUEwR0NTcUdTSWIzRFFFQkN3VUFNQlV4RXpBUkJnTlYKQkFNVENtdDFZbVZ5Ym1WMFpYTXdIaGNOTVRneE1URXlNVGN5TWpJeFdoY05NakF4TVRFeU1UY3lNakl4V2pBVgpNUk13RVFZRFZRUURFd3ByZFdKbGNtNWxkR1Z6TUZ3d0RRWUpLb1pJaHZjTkFRRUJCUUFEU3dBd1NBSkJBTko2CnFoY1ppQStLbFFERE5WakE2NE1USW0rd1hhVlJ2ekNmcG5vcmtGNDlSRzFWL1ZuZlZCTmVPNTBvb3E1ZGNpcEkKTjNvOG5lcGNPUEFOWG5Ob1ZFTUNBd0VBQWFPQnhqQ0J3ekFPQmdOVkhROEJBZjhFQkFNQ0JhQXdEQVlEVlIwVApBUUgvQkFJd0FEQ0JvZ1lEVlIwUkJJR2FNSUdYZ2dwcmRXSmxjbTVsZEdWemdnMXRZWE4wWlhJdE1EQXdNREF3CmdnMXRZWE4wWlhJdE1EQXdNREF4Z2cxdFlYTjBaWEl0TURBd01EQXlnZ3ByZFdKbGNtNWxkR1Z6Z2hKcmRXSmwKY201bGRHVnpMbVJsWm1GMWJIU0NGbXQxWW1WeWJtVjBaWE11WkdWbVlYVnNkQzV6ZG1PQ0pHdDFZbVZ5Ym1WMApaWE11WkdWbVlYVnNkQzV6ZG1NdVkyeDFjM1JsY2k1c2IyTmhiREFOQmdrcWhraUc5dzBCQVFzRkFBTkJBTkF2CjUwaXJlRHRqUkRiMVZydjJEcmRGZXhrT2hhZzNJM3dEVWlPb0loYjVuTkNNRjdnMEF1VEpOWkhFUXFOWDNYb1gKYkZpRkExdUxLZEZMc1B1T1dTUT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
			},
			"serviceCatalogServer": {
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2akNDQVpTZ0F3SUJBZ0lKQVBaMC8ydDJqYXJ5TUEwR0NTcUdTSWIzRFFFQkN3VUFNQlV4RXpBUkJnTlYKQkFNVENtdDFZbVZ5Ym1WMFpYTXdIaGNOTVRneE1URXlNVGN5TWpJeFdoY05NakF4TVRFeU1UY3lNakl4V2pBVgpNUk13RVFZRFZRUURFd3ByZFdKbGNtNWxkR1Z6TUZ3d0RRWUpLb1pJaHZjTkFRRUJCUUFEU3dBd1NBSkJBTko2CnFoY1ppQStLbFFERE5WakE2NE1USW0rd1hhVlJ2ekNmcG5vcmtGNDlSRzFWL1ZuZlZCTmVPNTBvb3E1ZGNpcEkKTjNvOG5lcGNPUEFOWG5Ob1ZFTUNBd0VBQWFPQnhqQ0J3ekFPQmdOVkhROEJBZjhFQkFNQ0JhQXdEQVlEVlIwVApBUUgvQkFJd0FEQ0JvZ1lEVlIwUkJJR2FNSUdYZ2dwcmRXSmxjbTVsZEdWemdnMXRZWE4wWlhJdE1EQXdNREF3CmdnMXRZWE4wWlhJdE1EQXdNREF4Z2cxdFlYTjBaWEl0TURBd01EQXlnZ3ByZFdKbGNtNWxkR1Z6Z2hKcmRXSmwKY201bGRHVnpMbVJsWm1GMWJIU0NGbXQxWW1WeWJtVjBaWE11WkdWbVlYVnNkQzV6ZG1PQ0pHdDFZbVZ5Ym1WMApaWE11WkdWbVlYVnNkQzV6ZG1NdVkyeDFjM1JsY2k1c2IyTmhiREFOQmdrcWhraUc5dzBCQVFzRkFBTkJBTkF2CjUwaXJlRHRqUkRiMVZydjJEcmRGZXhrT2hhZzNJM3dEVWlPb0loYjVuTkNNRjdnMEF1VEpOWkhFUXFOWDNYb1gKYkZpRkExdUxLZEZMc1B1T1dTUT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
			},
			"blackBoxMonitor": {
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2akNDQVpTZ0F3SUJBZ0lKQVBaMC8ydDJqYXJ5TUEwR0NTcUdTSWIzRFFFQkN3VUFNQlV4RXpBUkJnTlYKQkFNVENtdDFZbVZ5Ym1WMFpYTXdIaGNOTVRneE1URXlNVGN5TWpJeFdoY05NakF4TVRFeU1UY3lNakl4V2pBVgpNUk13RVFZRFZRUURFd3ByZFdKbGNtNWxkR1Z6TUZ3d0RRWUpLb1pJaHZjTkFRRUJCUUFEU3dBd1NBSkJBTko2CnFoY1ppQStLbFFERE5WakE2NE1USW0rd1hhVlJ2ekNmcG5vcmtGNDlSRzFWL1ZuZlZCTmVPNTBvb3E1ZGNpcEkKTjNvOG5lcGNPUEFOWG5Ob1ZFTUNBd0VBQWFPQnhqQ0J3ekFPQmdOVkhROEJBZjhFQkFNQ0JhQXdEQVlEVlIwVApBUUgvQkFJd0FEQ0JvZ1lEVlIwUkJJR2FNSUdYZ2dwcmRXSmxjbTVsZEdWemdnMXRZWE4wWlhJdE1EQXdNREF3CmdnMXRZWE4wWlhJdE1EQXdNREF4Z2cxdFlYTjBaWEl0TURBd01EQXlnZ3ByZFdKbGNtNWxkR1Z6Z2hKcmRXSmwKY201bGRHVnpMbVJsWm1GMWJIU0NGbXQxWW1WeWJtVjBaWE11WkdWbVlYVnNkQzV6ZG1PQ0pHdDFZbVZ5Ym1WMApaWE11WkdWbVlYVnNkQzV6ZG1NdVkyeDFjM1JsY2k1c2IyTmhiREFOQmdrcWhraUc5dzBCQVFzRkFBTkJBTkF2CjUwaXJlRHRqUkRiMVZydjJEcmRGZXhrT2hhZzNJM3dEVWlPb0loYjVuTkNNRjdnMEF1VEpOWkhFUXFOWDNYb1gKYkZpRkExdUxLZEZMc1B1T1dTUT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
			},
			"genevaLogging": {
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2akNDQVpTZ0F3SUJBZ0lKQVBaMC8ydDJqYXJ5TUEwR0NTcUdTSWIzRFFFQkN3VUFNQlV4RXpBUkJnTlYKQkFNVENtdDFZbVZ5Ym1WMFpYTXdIaGNOTVRneE1URXlNVGN5TWpJeFdoY05NakF4TVRFeU1UY3lNakl4V2pBVgpNUk13RVFZRFZRUURFd3ByZFdKbGNtNWxkR1Z6TUZ3d0RRWUpLb1pJaHZjTkFRRUJCUUFEU3dBd1NBSkJBTko2CnFoY1ppQStLbFFERE5WakE2NE1USW0rd1hhVlJ2ekNmcG5vcmtGNDlSRzFWL1ZuZlZCTmVPNTBvb3E1ZGNpcEkKTjNvOG5lcGNPUEFOWG5Ob1ZFTUNBd0VBQWFPQnhqQ0J3ekFPQmdOVkhROEJBZjhFQkFNQ0JhQXdEQVlEVlIwVApBUUgvQkFJd0FEQ0JvZ1lEVlIwUkJJR2FNSUdYZ2dwcmRXSmxjbTVsZEdWemdnMXRZWE4wWlhJdE1EQXdNREF3CmdnMXRZWE4wWlhJdE1EQXdNREF4Z2cxdFlYTjBaWEl0TURBd01EQXlnZ3ByZFdKbGNtNWxkR1Z6Z2hKcmRXSmwKY201bGRHVnpMbVJsWm1GMWJIU0NGbXQxWW1WeWJtVjBaWE11WkdWbVlYVnNkQzV6ZG1PQ0pHdDFZbVZ5Ym1WMApaWE11WkdWbVlYVnNkQzV6ZG1NdVkyeDFjM1JsY2k1c2IyTmhiREFOQmdrcWhraUc5dzBCQVFzRkFBTkJBTkF2CjUwaXJlRHRqUkRiMVZydjJEcmRGZXhrT2hhZzNJM3dEVWlPb0loYjVuTkNNRjdnMEF1VEpOWkhFUXFOWDNYb1gKYkZpRkExdUxLZEZMc1B1T1dTUT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
			},
			"genevaMetrics": {
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2akNDQVpTZ0F3SUJBZ0lKQVBaMC8ydDJqYXJ5TUEwR0NTcUdTSWIzRFFFQkN3VUFNQlV4RXpBUkJnTlYKQkFNVENtdDFZbVZ5Ym1WMFpYTXdIaGNOTVRneE1URXlNVGN5TWpJeFdoY05NakF4TVRFeU1UY3lNakl4V2pBVgpNUk13RVFZRFZRUURFd3ByZFdKbGNtNWxkR1Z6TUZ3d0RRWUpLb1pJaHZjTkFRRUJCUUFEU3dBd1NBSkJBTko2CnFoY1ppQStLbFFERE5WakE2NE1USW0rd1hhVlJ2ekNmcG5vcmtGNDlSRzFWL1ZuZlZCTmVPNTBvb3E1ZGNpcEkKTjNvOG5lcGNPUEFOWG5Ob1ZFTUNBd0VBQWFPQnhqQ0J3ekFPQmdOVkhROEJBZjhFQkFNQ0JhQXdEQVlEVlIwVApBUUgvQkFJd0FEQ0JvZ1lEVlIwUkJJR2FNSUdYZ2dwcmRXSmxjbTVsZEdWemdnMXRZWE4wWlhJdE1EQXdNREF3CmdnMXRZWE4wWlhJdE1EQXdNREF4Z2cxdFlYTjBaWEl0TURBd01EQXlnZ3ByZFdKbGNtNWxkR1Z6Z2hKcmRXSmwKY201bGRHVnpMbVJsWm1GMWJIU0NGbXQxWW1WeWJtVjBaWE11WkdWbVlYVnNkQzV6ZG1PQ0pHdDFZbVZ5Ym1WMApaWE11WkdWbVlYVnNkQzV6ZG1NdVkyeDFjM1JsY2k1c2IyTmhiREFOQmdrcWhraUc5dzBCQVFzRkFBTkJBTkF2CjUwaXJlRHRqUkRiMVZydjJEcmRGZXhrT2hhZzNJM3dEVWlPb0loYjVuTkNNRjdnMEF1VEpOWkhFUXFOWDNYb1gKYkZpRkExdUxLZEZMc1B1T1dTUT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
			}
		},
		"images": {
			"format": "Config.Images.Format",
			"clusterMonitoringOperator": "Config.Images.ClusterMonitoringOperator",
			"azureControllers": "Config.Images.AzureControllers",
			"prometheusOperator": "Config.Images.PrometheusOperator",
			"prometheus": "Config.Images.Prometheus",
			"prometheusConfigReloader": "Config.Images.PrometheusConfigReloader",
			"configReloader": "Config.Images.ConfigReloader",
			"alertManager": "Config.Images.AlertManager",
			"nodeExporter": "Config.Images.NodeExporter",
			"grafana": "Config.Images.Grafana",
			"kubeStateMetrics": "Config.Images.KubeStateMetrics",
			"kubeRbacProxy": "Config.Images.KubeRbacProxy",
			"oAuthProxy": "Config.Images.OAuthProxy",
			"masterEtcd": "Config.Images.MasterEtcd",
			"controlPlane": "Config.Images.ControlPlane",
			"node": "Config.Images.Node",
			"serviceCatalog": "Config.Images.ServiceCatalog",
			"sync": "Config.Images.Sync",
			"startup": "Config.Images.Startup",
			"templateServiceBroker": "Config.Images.TemplateServiceBroker",
			"registry": "Config.Images.Registry",
			"router": "Config.Images.Router",
			"registryConsole": "Config.Images.RegistryConsole",
			"ansibleServiceBroker": "Config.Images.AnsibleServiceBroker",
			"webConsole": "Config.Images.WebConsole",
			"console": "Config.Images.Console",
			"etcdBackup": "Config.Images.EtcdBackup",
			"httpd": "Config.Images.Httpd",
			"genevaLogging": "Config.Images.GenevaLogging",
			"genevaTDAgent": "Config.Images.GenevaTDAgent",
			"genevaStatsd": "Config.Images.GenevaStatsd",
			"metricsBridge": "Config.Images.MetricsBridge"
		},
		"serviceCatalogClusterId": "01010101-0101-0101-0101-010101010101",
		"genevaLoggingSector": "Config.GenevaLoggingSector",
		"genevaLoggingAccount": "Config.GenevaLoggingAccount",
		"genevaLoggingNamespace": "Config.GenevaLoggingNamespace",
		"genevaLoggingControlPlaneAccount": "Config.GenevaLoggingControlPlaneAccount",
		"genevaLoggingControlPlaneEnvironment": "Config.GenevaLoggingControlPlaneEnvironment",
		"genevaLoggingControlPlaneRegion": "Config.GenevaLoggingControlPlaneRegion",
		"genevaMetricsAccount": "Config.GenevaMetricsAccount",
		"genevaMetricsEndpoint": "Config.GenevaMetricsEndpoint"
	}
}`)

func TestMarshal(t *testing.T) {
	prepare := func(v reflect.Value) {
		switch v.Interface().(type) {
		case []IdentityProvider:
			// set the Provider to AADIdentityProvider
			v.Set(reflect.ValueOf([]IdentityProvider{{Provider: &AADIdentityProvider{Kind: to.StringPtr("AADIdentityProvider")}}}))
		}
	}

	populatedOc := OpenShiftManagedCluster{}
	populate.Walk(&populatedOc, prepare)

	b, err := json.MarshalIndent(populatedOc, "", "\t")
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(b, marshalled) {
		t.Errorf("json.MarshalIndent returned unexpected result\n%s\n", string(b))
	}
}

func TestUnmarshal(t *testing.T) {
	prepare := func(v reflect.Value) {
		switch v.Interface().(type) {
		case []IdentityProvider:
			// set the Provider to AADIdentityProvider
			v.Set(reflect.ValueOf([]IdentityProvider{{Provider: &AADIdentityProvider{Kind: to.StringPtr("AADIdentityProvider")}}}))
		}
	}

	populatedOc := OpenShiftManagedCluster{}
	populate.Walk(&populatedOc, prepare)

	var unmarshalledOc OpenShiftManagedCluster
	err := json.Unmarshal(marshalled, &unmarshalledOc)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(populatedOc, unmarshalledOc) {
		t.Errorf("json.Unmarshal returned unexpected result\n%#v\n", unmarshalledOc)
	}
}

// TestJSONTags ensures that all the `json:"..."` struct field tags under
// OpenShiftManagedCluster correspond with their field names
func TestJSONTags(t *testing.T) {
	o := OpenShiftManagedCluster{}
	for _, err := range structtags.CheckJsonTags(o) {
		t.Errorf("mismatch in struct tags for %T: %s", o, err.Error())
	}
}

func TestEnsureMasterProfileExists(t *testing.T) {
	prepare := func(v reflect.Value) {
		switch v.Interface().(type) {
		case []IdentityProvider:
			// set the Provider to AADIdentityProvider
			v.Set(reflect.ValueOf([]IdentityProvider{{Provider: &AADIdentityProvider{Kind: to.StringPtr("AADIdentityProvider")}}}))
		}
	}

	populatedOc := OpenShiftManagedCluster{}
	populate.Walk(&populatedOc, prepare)

	var unmarshalledOc OpenShiftManagedCluster
	err := json.Unmarshal(marshalled, &unmarshalledOc)
	if err != nil {
		t.Fatal(err)
	}

	if unmarshalledOc.Properties.MasterPoolProfile == nil {
		t.Fatalf("expected master pool profile to be set in external struct:\n%#v\n", spew.Sprint(unmarshalledOc))
	}

	if populatedOc.Properties.MasterPoolProfile == nil {
		t.Fatalf("expected master pool profile to be set in external struct:\n%#v\n", spew.Sprint(populatedOc))
	}

	if !reflect.DeepEqual(unmarshalledOc.Properties.MasterPoolProfile, populatedOc.Properties.MasterPoolProfile) {
		t.Errorf("json.Unmarshal returned unexpected result\n%#v\n", deep.Equal(unmarshalledOc.Properties.MasterPoolProfile, populatedOc.Properties.MasterPoolProfile))
	}
}
