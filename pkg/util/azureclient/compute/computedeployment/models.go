package computedeployment

import (
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"github.com/Azure/go-autorest/autorest"
)

// TODO: Needs to be replaced with a module genered from the swagger specs

// VirtualMachineScaleSet describes a Virtual Machine Scale Set.
type VirtualMachineScaleSet struct {
	autorest.Response `json:"-"`
	// Sku - The virtual machine scale set sku.
	Sku *compute.Sku `json:"sku,omitempty"`
	// Plan - Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
	Plan                                      *compute.Plan `json:"plan,omitempty"`
	*compute.VirtualMachineScaleSetProperties `json:"properties,omitempty"`
	// Identity - The identity of the virtual machine scale set, if configured.
	Identity *compute.VirtualMachineScaleSetIdentity `json:"identity,omitempty"`
	// Zones - The virtual machine scale set zones.
	Zones *[]string `json:"zones,omitempty"`
	// ID - READ-ONLY; Resource Id
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for VirtualMachineScaleSet.
func (vmss VirtualMachineScaleSet) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if vmss.Sku != nil {
		objectMap["sku"] = vmss.Sku
	}
	if vmss.Plan != nil {
		objectMap["plan"] = vmss.Plan
	}
	if vmss.VirtualMachineScaleSetProperties != nil {
		objectMap["properties"] = vmss.VirtualMachineScaleSetProperties
	}
	if vmss.Identity != nil {
		objectMap["identity"] = vmss.Identity
	}
	if vmss.Zones != nil {
		objectMap["zones"] = vmss.Zones
	}
	if vmss.ID != nil {
		objectMap["id"] = vmss.ID
	}
	if vmss.Name != nil {
		objectMap["name"] = vmss.Name
	}
	if vmss.Type != nil {
		objectMap["type"] = vmss.Type
	}
	if vmss.Location != nil {
		objectMap["location"] = vmss.Location
	}
	if vmss.Tags != nil {
		objectMap["tags"] = vmss.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for VirtualMachineScaleSet struct.
func (vmss *VirtualMachineScaleSet) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "sku":
			if v != nil {
				var sku compute.Sku
				err = json.Unmarshal(*v, &sku)
				if err != nil {
					return err
				}
				vmss.Sku = &sku
			}
		case "plan":
			if v != nil {
				var plan compute.Plan
				err = json.Unmarshal(*v, &plan)
				if err != nil {
					return err
				}
				vmss.Plan = &plan
			}
		case "properties":
			if v != nil {
				var virtualMachineScaleSetProperties compute.VirtualMachineScaleSetProperties
				err = json.Unmarshal(*v, &virtualMachineScaleSetProperties)
				if err != nil {
					return err
				}
				vmss.VirtualMachineScaleSetProperties = &virtualMachineScaleSetProperties
			}
		case "identity":
			if v != nil {
				var identity compute.VirtualMachineScaleSetIdentity
				err = json.Unmarshal(*v, &identity)
				if err != nil {
					return err
				}
				vmss.Identity = &identity
			}
		case "zones":
			if v != nil {
				var zones []string
				err = json.Unmarshal(*v, &zones)
				if err != nil {
					return err
				}
				vmss.Zones = &zones
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				vmss.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				vmss.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				vmss.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				vmss.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				vmss.Tags = tags
			}
		}
	}

	return nil
}

// VirtualMachine describes a Virtual Machine.
type VirtualMachine struct {
	autorest.Response `json:"-"`
	// Plan - Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
	Plan                              *compute.Plan `json:"plan,omitempty"`
	*compute.VirtualMachineProperties `json:"properties,omitempty"`
	// Resources - READ-ONLY; The virtual machine child extension resources.
	Resources *[]VirtualMachineExtension `json:"resources,omitempty"`
	// Identity - The identity of the virtual machine, if configured.
	Identity *compute.VirtualMachineIdentity `json:"identity,omitempty"`
	// Zones - The virtual machine zones.
	Zones *[]string `json:"zones,omitempty"`
	// ID - READ-ONLY; Resource Id
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for VirtualMachine.
func (VM VirtualMachine) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if VM.Plan != nil {
		objectMap["plan"] = VM.Plan
	}
	if VM.VirtualMachineProperties != nil {
		objectMap["properties"] = VM.VirtualMachineProperties
	}
	if VM.Resources != nil {
		objectMap["resources"] = VM.Resources
	}
	if VM.Identity != nil {
		objectMap["identity"] = VM.Identity
	}
	if VM.Zones != nil {
		objectMap["zones"] = VM.Zones
	}
	if VM.ID != nil {
		objectMap["id"] = VM.ID
	}
	if VM.Name != nil {
		objectMap["name"] = VM.Name
	}
	if VM.Type != nil {
		objectMap["type"] = VM.Type
	}
	if VM.Location != nil {
		objectMap["location"] = VM.Location
	}
	if VM.Tags != nil {
		objectMap["tags"] = VM.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for VirtualMachine struct.
func (VM *VirtualMachine) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "plan":
			if v != nil {
				var plan compute.Plan
				err = json.Unmarshal(*v, &plan)
				if err != nil {
					return err
				}
				VM.Plan = &plan
			}
		case "properties":
			if v != nil {
				var virtualMachineProperties compute.VirtualMachineProperties
				err = json.Unmarshal(*v, &virtualMachineProperties)
				if err != nil {
					return err
				}
				VM.VirtualMachineProperties = &virtualMachineProperties
			}
		case "resources":
			if v != nil {
				var resources []VirtualMachineExtension
				err = json.Unmarshal(*v, &resources)
				if err != nil {
					return err
				}
				VM.Resources = &resources
			}
		case "identity":
			if v != nil {
				var identity compute.VirtualMachineIdentity
				err = json.Unmarshal(*v, &identity)
				if err != nil {
					return err
				}
				VM.Identity = &identity
			}
		case "zones":
			if v != nil {
				var zones []string
				err = json.Unmarshal(*v, &zones)
				if err != nil {
					return err
				}
				VM.Zones = &zones
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				VM.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				VM.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				VM.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				VM.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				VM.Tags = tags
			}
		}
	}

	return nil
}

// VirtualMachineExtension describes a Virtual Machine Extension.
type VirtualMachineExtension struct {
	autorest.Response                          `json:"-"`
	*compute.VirtualMachineExtensionProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; Resource Id
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for VirtualMachineExtension.
// MarshalJSON is the custom marshaler for VirtualMachineExtension.
func (vme VirtualMachineExtension) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if vme.VirtualMachineExtensionProperties != nil {
		objectMap["properties"] = vme.VirtualMachineExtensionProperties
	}
	if vme.ID != nil {
		objectMap["id"] = vme.ID
	}
	if vme.Name != nil {
		objectMap["name"] = vme.Name
	}
	if vme.Type != nil {
		objectMap["type"] = vme.Type
	}
	if vme.Location != nil {
		objectMap["location"] = vme.Location
	}
	if vme.Tags != nil {
		objectMap["tags"] = vme.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for VirtualMachineExtension struct.
func (vme *VirtualMachineExtension) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var virtualMachineExtensionProperties compute.VirtualMachineExtensionProperties
				err = json.Unmarshal(*v, &virtualMachineExtensionProperties)
				if err != nil {
					return err
				}
				vme.VirtualMachineExtensionProperties = &virtualMachineExtensionProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				vme.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				vme.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				vme.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				vme.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				vme.Tags = tags
			}
		}
	}

	return nil
}
