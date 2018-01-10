package models

// DatabaseNode

import "encoding/json"

// DatabaseNode
type DatabaseNode struct {
	DatabaseNodeIPAddress IpAddressType  `json:"database_node_ip_address"`
	DisplayName           string         `json:"display_name"`
	Annotations           *KeyValuePairs `json:"annotations"`
	Perms2                *PermType2     `json:"perms2"`
	UUID                  string         `json:"uuid"`
	IDPerms               *IdPermsType   `json:"id_perms"`
	ParentUUID            string         `json:"parent_uuid"`
	ParentType            string         `json:"parent_type"`
	FQName                []string       `json:"fq_name"`
}

// String returns json representation of the object
func (model *DatabaseNode) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeDatabaseNode makes DatabaseNode
func MakeDatabaseNode() *DatabaseNode {
	return &DatabaseNode{
		//TODO(nati): Apply default
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		ParentUUID:  "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		UUID:        "",
		DatabaseNodeIPAddress: MakeIpAddressType(),
		DisplayName:           "",
	}
}

// InterfaceToDatabaseNode makes DatabaseNode from interface
func InterfaceToDatabaseNode(iData interface{}) *DatabaseNode {
	data := iData.(map[string]interface{})
	return &DatabaseNode{
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		DatabaseNodeIPAddress: InterfaceToIpAddressType(data["database_node_ip_address"]),

		//{"description":"Ip address of the database node, set while provisioning.","type":"string"}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}

	}
}

// InterfaceToDatabaseNodeSlice makes a slice of DatabaseNode from interface
func InterfaceToDatabaseNodeSlice(data interface{}) []*DatabaseNode {
	list := data.([]interface{})
	result := MakeDatabaseNodeSlice()
	for _, item := range list {
		result = append(result, InterfaceToDatabaseNode(item))
	}
	return result
}

// MakeDatabaseNodeSlice() makes a slice of DatabaseNode
func MakeDatabaseNodeSlice() []*DatabaseNode {
	return []*DatabaseNode{}
}