package models

// AddressGroup

import "encoding/json"

// AddressGroup
type AddressGroup struct {
	Perms2             *PermType2      `json:"perms2"`
	UUID               string          `json:"uuid"`
	ParentType         string          `json:"parent_type"`
	Annotations        *KeyValuePairs  `json:"annotations"`
	AddressGroupPrefix *SubnetListType `json:"address_group_prefix"`
	FQName             []string        `json:"fq_name"`
	IDPerms            *IdPermsType    `json:"id_perms"`
	DisplayName        string          `json:"display_name"`
	ParentUUID         string          `json:"parent_uuid"`
}

// String returns json representation of the object
func (model *AddressGroup) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeAddressGroup makes AddressGroup
func MakeAddressGroup() *AddressGroup {
	return &AddressGroup{
		//TODO(nati): Apply default
		FQName:             []string{},
		IDPerms:            MakeIdPermsType(),
		DisplayName:        "",
		ParentUUID:         "",
		Perms2:             MakePermType2(),
		UUID:               "",
		ParentType:         "",
		Annotations:        MakeKeyValuePairs(),
		AddressGroupPrefix: MakeSubnetListType(),
	}
}

// InterfaceToAddressGroup makes AddressGroup from interface
func InterfaceToAddressGroup(iData interface{}) *AddressGroup {
	data := iData.(map[string]interface{})
	return &AddressGroup{
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		AddressGroupPrefix: InterfaceToSubnetListType(data["address_group_prefix"]),

		//{"description":"List of IP prefix","type":"object","properties":{"subnet":{"type":"array","item":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}

	}
}

// InterfaceToAddressGroupSlice makes a slice of AddressGroup from interface
func InterfaceToAddressGroupSlice(data interface{}) []*AddressGroup {
	list := data.([]interface{})
	result := MakeAddressGroupSlice()
	for _, item := range list {
		result = append(result, InterfaceToAddressGroup(item))
	}
	return result
}

// MakeAddressGroupSlice() makes a slice of AddressGroup
func MakeAddressGroupSlice() []*AddressGroup {
	return []*AddressGroup{}
}