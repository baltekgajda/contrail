package models

// PhysicalRouter

import "encoding/json"

// PhysicalRouter
type PhysicalRouter struct {
	PhysicalRouterSNMP              bool                `json:"physical_router_snmp"`
	PhysicalRouterDataplaneIP       string              `json:"physical_router_dataplane_ip"`
	DisplayName                     string              `json:"display_name"`
	ParentUUID                      string              `json:"parent_uuid"`
	ParentType                      string              `json:"parent_type"`
	PhysicalRouterSNMPCredentials   *SNMPCredentials    `json:"physical_router_snmp_credentials"`
	PhysicalRouterUserCredentials   *UserCredentials    `json:"physical_router_user_credentials"`
	PhysicalRouterLoopbackIP        string              `json:"physical_router_loopback_ip"`
	TelemetryInfo                   *TelemetryStateInfo `json:"telemetry_info"`
	Perms2                          *PermType2          `json:"perms2"`
	UUID                            string              `json:"uuid"`
	PhysicalRouterRole              PhysicalRouterRole  `json:"physical_router_role"`
	PhysicalRouterImageURI          string              `json:"physical_router_image_uri"`
	Annotations                     *KeyValuePairs      `json:"annotations"`
	FQName                          []string            `json:"fq_name"`
	PhysicalRouterVNCManaged        bool                `json:"physical_router_vnc_managed"`
	PhysicalRouterVendorName        string              `json:"physical_router_vendor_name"`
	PhysicalRouterProductName       string              `json:"physical_router_product_name"`
	PhysicalRouterLLDP              bool                `json:"physical_router_lldp"`
	PhysicalRouterJunosServicePorts *JunosServicePorts  `json:"physical_router_junos_service_ports"`
	IDPerms                         *IdPermsType        `json:"id_perms"`
	PhysicalRouterManagementIP      string              `json:"physical_router_management_ip"`

	VirtualNetworkRefs []*PhysicalRouterVirtualNetworkRef `json:"virtual_network_refs"`
	BGPRouterRefs      []*PhysicalRouterBGPRouterRef      `json:"bgp_router_refs"`
	VirtualRouterRefs  []*PhysicalRouterVirtualRouterRef  `json:"virtual_router_refs"`

	LogicalInterfaces  []*LogicalInterface  `json:"logical_interfaces"`
	PhysicalInterfaces []*PhysicalInterface `json:"physical_interfaces"`
}

// PhysicalRouterVirtualNetworkRef references each other
type PhysicalRouterVirtualNetworkRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// PhysicalRouterBGPRouterRef references each other
type PhysicalRouterBGPRouterRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// PhysicalRouterVirtualRouterRef references each other
type PhysicalRouterVirtualRouterRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *PhysicalRouter) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakePhysicalRouter makes PhysicalRouter
func MakePhysicalRouter() *PhysicalRouter {
	return &PhysicalRouter{
		//TODO(nati): Apply default
		PhysicalRouterProductName:       "",
		PhysicalRouterLLDP:              false,
		PhysicalRouterJunosServicePorts: MakeJunosServicePorts(),
		IDPerms:                       MakeIdPermsType(),
		PhysicalRouterManagementIP:    "",
		PhysicalRouterVendorName:      "",
		PhysicalRouterDataplaneIP:     "",
		DisplayName:                   "",
		ParentUUID:                    "",
		ParentType:                    "",
		PhysicalRouterSNMPCredentials: MakeSNMPCredentials(),
		PhysicalRouterSNMP:            false,
		PhysicalRouterLoopbackIP:      "",
		TelemetryInfo:                 MakeTelemetryStateInfo(),
		Perms2:                        MakePermType2(),
		UUID:                          "",
		PhysicalRouterRole:            MakePhysicalRouterRole(),
		PhysicalRouterUserCredentials: MakeUserCredentials(),
		Annotations:                   MakeKeyValuePairs(),
		FQName:                        []string{},
		PhysicalRouterVNCManaged: false,
		PhysicalRouterImageURI:   "",
	}
}

// InterfaceToPhysicalRouter makes PhysicalRouter from interface
func InterfaceToPhysicalRouter(iData interface{}) *PhysicalRouter {
	data := iData.(map[string]interface{})
	return &PhysicalRouter{
		PhysicalRouterSNMP: data["physical_router_snmp"].(bool),

		//{"description":"SNMP support on this router","type":"boolean"}
		PhysicalRouterDataplaneIP: data["physical_router_dataplane_ip"].(string),

		//{"description":"This is ip address in the ip-fabric(underlay) network that can be used in data plane by physical router. Usually it is the VTEP address in VxLAN for the TOR switch.","type":"string"}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		PhysicalRouterSNMPCredentials: InterfaceToSNMPCredentials(data["physical_router_snmp_credentials"]),

		//{"description":"SNMP credentials for the physical router used by SNMP collector.","type":"object","properties":{"local_port":{"type":"integer"},"retries":{"type":"integer"},"timeout":{"type":"integer"},"v2_community":{"type":"string"},"v3_authentication_password":{"type":"string"},"v3_authentication_protocol":{"type":"string"},"v3_context":{"type":"string"},"v3_context_engine_id":{"type":"string"},"v3_engine_boots":{"type":"integer"},"v3_engine_id":{"type":"string"},"v3_engine_time":{"type":"integer"},"v3_privacy_password":{"type":"string"},"v3_privacy_protocol":{"type":"string"},"v3_security_engine_id":{"type":"string"},"v3_security_level":{"type":"string"},"v3_security_name":{"type":"string"},"version":{"type":"integer"}}}
		PhysicalRouterUserCredentials: InterfaceToUserCredentials(data["physical_router_user_credentials"]),

		//{"description":"Username and password for netconf to the physical router by device manager.","type":"object","properties":{"password":{"type":"string"},"username":{"type":"string"}}}
		PhysicalRouterLoopbackIP: data["physical_router_loopback_ip"].(string),

		//{"description":"This is ip address of loopback interface of physical router. Used by the device manager to configure physical router loopback interface.","type":"string"}
		TelemetryInfo: InterfaceToTelemetryStateInfo(data["telemetry_info"]),

		//{"description":"Telemetry info of router. Check TelemetryStateInfo","type":"object","properties":{"resource":{"type":"array","item":{"type":"object","properties":{"name":{"type":"string"},"path":{"type":"string"},"rate":{"type":"string"}}}},"server_ip":{"type":"string"},"server_port":{"type":"integer"}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		PhysicalRouterRole: InterfaceToPhysicalRouterRole(data["physical_router_role"]),

		//{"description":"Physical router role (e.g spine or leaf), used by the device manager to provision physical router, for e.g device manager may choose to configure physical router based on its role.","type":"string","enum":["spine","leaf","e2-access","e2-provider","e2-internet","e2-vrr"]}
		PhysicalRouterImageURI: data["physical_router_image_uri"].(string),

		//{"description":"Physical router OS image uri","type":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		PhysicalRouterVNCManaged: data["physical_router_vnc_managed"].(bool),

		//{"description":"This physical router is enabled to be configured by device manager.","type":"boolean"}
		PhysicalRouterVendorName: data["physical_router_vendor_name"].(string),

		//{"description":"Vendor name of the physical router (e.g juniper). Used by the device manager to select driver.","type":"string"}
		PhysicalRouterProductName: data["physical_router_product_name"].(string),

		//{"description":"Model name of the physical router (e.g juniper). Used by the device manager to select driver.","type":"string"}
		PhysicalRouterLLDP: data["physical_router_lldp"].(bool),

		//{"description":"LLDP support on this router","type":"boolean"}
		PhysicalRouterJunosServicePorts: InterfaceToJunosServicePorts(data["physical_router_junos_service_ports"]),

		//{"description":"Juniper JUNOS specific service interfaces name  to perform services like NAT.","type":"object","properties":{"service_port":{"type":"array","item":{"type":"string"}}}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		PhysicalRouterManagementIP: data["physical_router_management_ip"].(string),

		//{"description":"Management ip for this physical router. It is used by the device manager to perform netconf and by SNMP collector if enabled.","type":"string"}

	}
}

// InterfaceToPhysicalRouterSlice makes a slice of PhysicalRouter from interface
func InterfaceToPhysicalRouterSlice(data interface{}) []*PhysicalRouter {
	list := data.([]interface{})
	result := MakePhysicalRouterSlice()
	for _, item := range list {
		result = append(result, InterfaceToPhysicalRouter(item))
	}
	return result
}

// MakePhysicalRouterSlice() makes a slice of PhysicalRouter
func MakePhysicalRouterSlice() []*PhysicalRouter {
	return []*PhysicalRouter{}
}
