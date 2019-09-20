package storagedeployment

import (
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2018-02-01/storage"
	"github.com/Azure/go-autorest/autorest"
)

// TODO: Needs to be replaced with a module genered from the swagger specs

// Account the storage account.
type Account struct {
	autorest.Response `json:"-"`
	// Sku - READ-ONLY; Gets the SKU.
	Sku *storage.Sku `json:"sku,omitempty"`
	// Kind - READ-ONLY; Gets the Kind. Possible values include: 'Storage', 'StorageV2', 'BlobStorage'
	Kind storage.Kind `json:"kind,omitempty"`
	// Identity - The identity of the resource.
	Identity *storage.Identity `json:"identity,omitempty"`
	// AccountProperties - Properties of the storage account.
	*storage.AccountProperties `json:"properties,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Account.
func (a Account) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if a.Sku != nil {
		objectMap["sku"] = a.Sku
	}
	if a.Kind != "" {
		objectMap["kind"] = a.Kind
	}
	if a.Identity != nil {
		objectMap["identity"] = a.Identity
	}
	if a.AccountProperties != nil {
		objectMap["properties"] = a.AccountProperties
	}
	if a.Tags != nil {
		objectMap["tags"] = a.Tags
	}
	if a.Location != nil {
		objectMap["location"] = a.Location
	}
	if a.ID != nil {
		objectMap["id"] = a.ID
	}
	if a.Name != nil {
		objectMap["name"] = a.Name
	}
	if a.Type != nil {
		objectMap["type"] = a.Type
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Account struct.
func (a *Account) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "sku":
			if v != nil {
				var sku storage.Sku
				err = json.Unmarshal(*v, &sku)
				if err != nil {
					return err
				}
				a.Sku = &sku
			}
		case "kind":
			if v != nil {
				var kind storage.Kind
				err = json.Unmarshal(*v, &kind)
				if err != nil {
					return err
				}
				a.Kind = kind
			}
		case "identity":
			if v != nil {
				var identity storage.Identity
				err = json.Unmarshal(*v, &identity)
				if err != nil {
					return err
				}
				a.Identity = &identity
			}
		case "properties":
			if v != nil {
				var accountProperties storage.AccountProperties
				err = json.Unmarshal(*v, &accountProperties)
				if err != nil {
					return err
				}
				a.AccountProperties = &accountProperties
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				a.Tags = tags
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				a.Location = &location
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				a.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				a.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				a.Type = &typeVar
			}
		}
	}

	return nil
}
