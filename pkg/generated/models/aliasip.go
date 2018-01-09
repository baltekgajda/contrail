package models

// AliasIP

import "encoding/json"

// AliasIP
type AliasIP struct {
	ParentUUID           string              `json:"parent_uuid"`
	ParentType           string              `json:"parent_type"`
	FQName               []string            `json:"fq_name"`
	AliasIPAddress       IpAddressType       `json:"alias_ip_address"`
	DisplayName          string              `json:"display_name"`
	UUID                 string              `json:"uuid"`
	Perms2               *PermType2          `json:"perms2"`
	AliasIPAddressFamily IpAddressFamilyType `json:"alias_ip_address_family"`
	IDPerms              *IdPermsType        `json:"id_perms"`
	Annotations          *KeyValuePairs      `json:"annotations"`

	ProjectRefs                 []*AliasIPProjectRef                 `json:"project_refs"`
	VirtualMachineInterfaceRefs []*AliasIPVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs"`
}

// AliasIPProjectRef references each other
type AliasIPProjectRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// AliasIPVirtualMachineInterfaceRef references each other
type AliasIPVirtualMachineInterfaceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *AliasIP) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeAliasIP makes AliasIP
func MakeAliasIP() *AliasIP {
	return &AliasIP{
		//TODO(nati): Apply default
		AliasIPAddress:       MakeIpAddressType(),
		DisplayName:          "",
		UUID:                 "",
		ParentUUID:           "",
		ParentType:           "",
		FQName:               []string{},
		AliasIPAddressFamily: MakeIpAddressFamilyType(),
		IDPerms:              MakeIdPermsType(),
		Annotations:          MakeKeyValuePairs(),
		Perms2:               MakePermType2(),
	}
}

// InterfaceToAliasIP makes AliasIP from interface
func InterfaceToAliasIP(iData interface{}) *AliasIP {
	data := iData.(map[string]interface{})
	return &AliasIP{
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		AliasIPAddress: InterfaceToIpAddressType(data["alias_ip_address"]),

		//{"description":"Alias ip address.","type":"string"}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		AliasIPAddressFamily: InterfaceToIpAddressFamilyType(data["alias_ip_address_family"]),

		//{"description":"Ip address family of the alias ip, IpV4 or IpV6","type":"string","enum":["v4","v6"]}

	}
}

// InterfaceToAliasIPSlice makes a slice of AliasIP from interface
func InterfaceToAliasIPSlice(data interface{}) []*AliasIP {
	list := data.([]interface{})
	result := MakeAliasIPSlice()
	for _, item := range list {
		result = append(result, InterfaceToAliasIP(item))
	}
	return result
}

// MakeAliasIPSlice() makes a slice of AliasIP
func MakeAliasIPSlice() []*AliasIP {
	return []*AliasIP{}
}
