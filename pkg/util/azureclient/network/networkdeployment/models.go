package networkdeployment

// TODO: Needs to be replaced with a module genered from the swagger specs

import (
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-07-01/network"
	"github.com/Azure/go-autorest/autorest"
)

// VirtualNetwork virtual Network resource.
type VirtualNetwork struct {
	autorest.Response `json:"-"`
	// VirtualNetworkPropertiesFormat - Properties of the virtual network.
	*network.VirtualNetworkPropertiesFormat `json:"properties,omitempty"`
	// Etag - Gets a unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`
	// ID - Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type.
	Type *string `json:"type,omitempty"`
	// Location - Resource location.
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for VirtualNetwork.
func (vn VirtualNetwork) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if vn.VirtualNetworkPropertiesFormat != nil {
		objectMap["properties"] = vn.VirtualNetworkPropertiesFormat
	}
	if vn.Etag != nil {
		objectMap["etag"] = vn.Etag
	}
	if vn.ID != nil {
		objectMap["id"] = vn.ID
	}
	if vn.Name != nil {
		objectMap["name"] = vn.Name
	}
	if vn.Type != nil {
		objectMap["type"] = vn.Type
	}
	if vn.Location != nil {
		objectMap["location"] = vn.Location
	}
	if vn.Tags != nil {
		objectMap["tags"] = vn.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for VirtualNetwork struct.
func (vn *VirtualNetwork) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var virtualNetworkPropertiesFormat network.VirtualNetworkPropertiesFormat
				err = json.Unmarshal(*v, &virtualNetworkPropertiesFormat)
				if err != nil {
					return err
				}
				vn.VirtualNetworkPropertiesFormat = &virtualNetworkPropertiesFormat
			}
		case "etag":
			if v != nil {
				var etag string
				err = json.Unmarshal(*v, &etag)
				if err != nil {
					return err
				}
				vn.Etag = &etag
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				vn.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				vn.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				vn.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				vn.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				vn.Tags = tags
			}
		}
	}

	return nil
}

// PublicIPAddress public IP address resource.
type PublicIPAddress struct {
	autorest.Response `json:"-"`
	// Sku - The public IP address SKU.
	Sku *network.PublicIPAddressSku `json:"sku,omitempty"`
	// PublicIPAddressPropertiesFormat - Public IP address properties.
	*network.PublicIPAddressPropertiesFormat `json:"properties,omitempty"`
	// Etag - A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`
	// Zones - A list of availability zones denoting the IP allocated for the resource needs to come from.
	Zones *[]string `json:"zones,omitempty"`
	// ID - Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type.
	Type *string `json:"type,omitempty"`
	// Location - Resource location.
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for PublicIPAddress.
func (pia PublicIPAddress) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if pia.Sku != nil {
		objectMap["sku"] = pia.Sku
	}
	if pia.PublicIPAddressPropertiesFormat != nil {
		objectMap["properties"] = pia.PublicIPAddressPropertiesFormat
	}
	if pia.Etag != nil {
		objectMap["etag"] = pia.Etag
	}
	if pia.Zones != nil {
		objectMap["zones"] = pia.Zones
	}
	if pia.ID != nil {
		objectMap["id"] = pia.ID
	}
	if pia.Name != nil {
		objectMap["name"] = pia.Name
	}
	if pia.Type != nil {
		objectMap["type"] = pia.Type
	}
	if pia.Location != nil {
		objectMap["location"] = pia.Location
	}
	if pia.Tags != nil {
		objectMap["tags"] = pia.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for PublicIPAddress struct.
func (pia *PublicIPAddress) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "sku":
			if v != nil {
				var sku network.PublicIPAddressSku
				err = json.Unmarshal(*v, &sku)
				if err != nil {
					return err
				}
				pia.Sku = &sku
			}
		case "properties":
			if v != nil {
				var publicIPAddressPropertiesFormat network.PublicIPAddressPropertiesFormat
				err = json.Unmarshal(*v, &publicIPAddressPropertiesFormat)
				if err != nil {
					return err
				}
				pia.PublicIPAddressPropertiesFormat = &publicIPAddressPropertiesFormat
			}
		case "etag":
			if v != nil {
				var etag string
				err = json.Unmarshal(*v, &etag)
				if err != nil {
					return err
				}
				pia.Etag = &etag
			}
		case "zones":
			if v != nil {
				var zones []string
				err = json.Unmarshal(*v, &zones)
				if err != nil {
					return err
				}
				pia.Zones = &zones
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				pia.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				pia.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				pia.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				pia.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				pia.Tags = tags
			}
		}
	}

	return nil
}

// LoadBalancer loadBalancer resource
type LoadBalancer struct {
	autorest.Response `json:"-"`
	// Sku - The load balancer SKU.
	Sku *network.LoadBalancerSku `json:"sku,omitempty"`
	// LoadBalancerPropertiesFormat - Properties of load balancer.
	*network.LoadBalancerPropertiesFormat `json:"properties,omitempty"`
	// Etag - A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`
	// ID - Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type.
	Type *string `json:"type,omitempty"`
	// Location - Resource location.
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for LoadBalancer.
func (lb LoadBalancer) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if lb.Sku != nil {
		objectMap["sku"] = lb.Sku
	}
	if lb.LoadBalancerPropertiesFormat != nil {
		objectMap["properties"] = lb.LoadBalancerPropertiesFormat
	}
	if lb.Etag != nil {
		objectMap["etag"] = lb.Etag
	}
	if lb.ID != nil {
		objectMap["id"] = lb.ID
	}
	if lb.Name != nil {
		objectMap["name"] = lb.Name
	}
	if lb.Type != nil {
		objectMap["type"] = lb.Type
	}
	if lb.Location != nil {
		objectMap["location"] = lb.Location
	}
	if lb.Tags != nil {
		objectMap["tags"] = lb.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for LoadBalancer struct.
func (lb *LoadBalancer) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "sku":
			if v != nil {
				var sku network.LoadBalancerSku
				err = json.Unmarshal(*v, &sku)
				if err != nil {
					return err
				}
				lb.Sku = &sku
			}
		case "properties":
			if v != nil {
				var loadBalancerPropertiesFormat network.LoadBalancerPropertiesFormat
				err = json.Unmarshal(*v, &loadBalancerPropertiesFormat)
				if err != nil {
					return err
				}
				lb.LoadBalancerPropertiesFormat = &loadBalancerPropertiesFormat
			}
		case "etag":
			if v != nil {
				var etag string
				err = json.Unmarshal(*v, &etag)
				if err != nil {
					return err
				}
				lb.Etag = &etag
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				lb.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				lb.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				lb.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				lb.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				lb.Tags = tags
			}
		}
	}

	return nil
}

// SecurityGroup networkSecurityGroup resource.
type SecurityGroup struct {
	autorest.Response `json:"-"`
	// SecurityGroupPropertiesFormat - Properties of the network security group
	*network.SecurityGroupPropertiesFormat `json:"properties,omitempty"`
	// Etag - A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`
	// ID - Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type.
	Type *string `json:"type,omitempty"`
	// Location - Resource location.
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for SecurityGroup.
func (sg SecurityGroup) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if sg.SecurityGroupPropertiesFormat != nil {
		objectMap["properties"] = sg.SecurityGroupPropertiesFormat
	}
	if sg.Etag != nil {
		objectMap["etag"] = sg.Etag
	}
	if sg.ID != nil {
		objectMap["id"] = sg.ID
	}
	if sg.Name != nil {
		objectMap["name"] = sg.Name
	}
	if sg.Type != nil {
		objectMap["type"] = sg.Type
	}
	if sg.Location != nil {
		objectMap["location"] = sg.Location
	}
	if sg.Tags != nil {
		objectMap["tags"] = sg.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for SecurityGroup struct.
func (sg *SecurityGroup) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var securityGroupPropertiesFormat network.SecurityGroupPropertiesFormat
				err = json.Unmarshal(*v, &securityGroupPropertiesFormat)
				if err != nil {
					return err
				}
				sg.SecurityGroupPropertiesFormat = &securityGroupPropertiesFormat
			}
		case "etag":
			if v != nil {
				var etag string
				err = json.Unmarshal(*v, &etag)
				if err != nil {
					return err
				}
				sg.Etag = &etag
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				sg.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				sg.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				sg.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				sg.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				sg.Tags = tags
			}
		}
	}

	return nil
}

// Interface a network interface in a resource group.
type Interface struct {
	autorest.Response `json:"-"`
	// InterfacePropertiesFormat - Properties of the network interface.
	*network.InterfacePropertiesFormat `json:"properties,omitempty"`
	// Etag - A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`
	// ID - Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type.
	Type *string `json:"type,omitempty"`
	// Location - Resource location.
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Interface.
func (i Interface) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if i.InterfacePropertiesFormat != nil {
		objectMap["properties"] = i.InterfacePropertiesFormat
	}
	if i.Etag != nil {
		objectMap["etag"] = i.Etag
	}
	if i.ID != nil {
		objectMap["id"] = i.ID
	}
	if i.Name != nil {
		objectMap["name"] = i.Name
	}
	if i.Type != nil {
		objectMap["type"] = i.Type
	}
	if i.Location != nil {
		objectMap["location"] = i.Location
	}
	if i.Tags != nil {
		objectMap["tags"] = i.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Interface struct.
func (i *Interface) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var interfacePropertiesFormat network.InterfacePropertiesFormat
				err = json.Unmarshal(*v, &interfacePropertiesFormat)
				if err != nil {
					return err
				}
				i.InterfacePropertiesFormat = &interfacePropertiesFormat
			}
		case "etag":
			if v != nil {
				var etag string
				err = json.Unmarshal(*v, &etag)
				if err != nil {
					return err
				}
				i.Etag = &etag
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				i.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				i.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				i.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				i.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				i.Tags = tags
			}
		}
	}

	return nil
}
