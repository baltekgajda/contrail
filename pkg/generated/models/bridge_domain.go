package models

// BridgeDomain

import "encoding/json"

// BridgeDomain
type BridgeDomain struct {
	MacMoveControl     *MACMoveLimitControlType `json:"mac_move_control"`
	Annotations        *KeyValuePairs           `json:"annotations"`
	Perms2             *PermType2               `json:"perms2"`
	UUID               string                   `json:"uuid"`
	ParentType         string                   `json:"parent_type"`
	FQName             []string                 `json:"fq_name"`
	IDPerms            *IdPermsType             `json:"id_perms"`
	MacAgingTime       MACAgingTime             `json:"mac_aging_time"`
	DisplayName        string                   `json:"display_name"`
	MacLearningEnabled bool                     `json:"mac_learning_enabled"`
	MacLimitControl    *MACLimitControlType     `json:"mac_limit_control"`
	ParentUUID         string                   `json:"parent_uuid"`
	Isid               IsidType                 `json:"isid"`
}

// String returns json representation of the object
func (model *BridgeDomain) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeBridgeDomain makes BridgeDomain
func MakeBridgeDomain() *BridgeDomain {
	return &BridgeDomain{
		//TODO(nati): Apply default
		MacMoveControl:     MakeMACMoveLimitControlType(),
		Annotations:        MakeKeyValuePairs(),
		Perms2:             MakePermType2(),
		UUID:               "",
		ParentType:         "",
		FQName:             []string{},
		IDPerms:            MakeIdPermsType(),
		MacAgingTime:       MakeMACAgingTime(),
		DisplayName:        "",
		MacLearningEnabled: false,
		MacLimitControl:    MakeMACLimitControlType(),
		ParentUUID:         "",
		Isid:               MakeIsidType(),
	}
}

// InterfaceToBridgeDomain makes BridgeDomain from interface
func InterfaceToBridgeDomain(iData interface{}) *BridgeDomain {
	data := iData.(map[string]interface{})
	return &BridgeDomain{
		MacLearningEnabled: data["mac_learning_enabled"].(bool),

		//{"description":"Enable MAC learning on the network","default":false,"type":"boolean"}
		MacLimitControl: InterfaceToMACLimitControlType(data["mac_limit_control"]),

		//{"description":"MAC limit control on the network","type":"object","properties":{"mac_limit":{"type":"integer"},"mac_limit_action":{"type":"string","enum":["log","alarm","shutdown","drop"]}}}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		Isid: InterfaceToIsidType(data["isid"]),

		//{"description":"i-sid value","type":"integer","minimum":1,"maximum":16777215}
		MacMoveControl: InterfaceToMACMoveLimitControlType(data["mac_move_control"]),

		//{"description":"MAC move control on the network","type":"object","properties":{"mac_move_limit":{"type":"integer"},"mac_move_limit_action":{"type":"string","enum":["log","alarm","shutdown","drop"]},"mac_move_time_window":{"type":"integer","minimum":1,"maximum":60}}}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		MacAgingTime: InterfaceToMACAgingTime(data["mac_aging_time"]),

		//{"description":"MAC aging time on the network","default":"300","type":"integer","minimum":0,"maximum":86400}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}

	}
}

// InterfaceToBridgeDomainSlice makes a slice of BridgeDomain from interface
func InterfaceToBridgeDomainSlice(data interface{}) []*BridgeDomain {
	list := data.([]interface{})
	result := MakeBridgeDomainSlice()
	for _, item := range list {
		result = append(result, InterfaceToBridgeDomain(item))
	}
	return result
}

// MakeBridgeDomainSlice() makes a slice of BridgeDomain
func MakeBridgeDomainSlice() []*BridgeDomain {
	return []*BridgeDomain{}
}