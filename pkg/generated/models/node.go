package models

// Node

import "encoding/json"

// Node
type Node struct {
	Username                       string         `json:"username"`
	AwsAmi                         string         `json:"aws_ami"`
	GCPImage                       string         `json:"gcp_image"`
	PrivatePowerManagementUsername string         `json:"private_power_management_username"`
	UUID                           string         `json:"uuid"`
	Hostname                       string         `json:"hostname"`
	MacAddress                     string         `json:"mac_address"`
	PrivatePowerManagementIP       string         `json:"private_power_management_ip"`
	DisplayName                    string         `json:"display_name"`
	Perms2                         *PermType2     `json:"perms2"`
	ParentType                     string         `json:"parent_type"`
	Type                           string         `json:"type"`
	GCPMachineType                 string         `json:"gcp_machine_type"`
	Annotations                    *KeyValuePairs `json:"annotations"`
	FQName                         []string       `json:"fq_name"`
	IPAddress                      string         `json:"ip_address"`
	PrivatePowerManagementPassword string         `json:"private_power_management_password"`
	AwsInstanceType                string         `json:"aws_instance_type"`
	PrivateMachineProperties       string         `json:"private_machine_properties"`
	PrivateMachineState            string         `json:"private_machine_state"`
	IDPerms                        *IdPermsType   `json:"id_perms"`
	ParentUUID                     string         `json:"parent_uuid"`
	Password                       string         `json:"password"`
	SSHKey                         string         `json:"ssh_key"`
}

// String returns json representation of the object
func (model *Node) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeNode makes Node
func MakeNode() *Node {
	return &Node{
		//TODO(nati): Apply default
		ParentType:                     "",
		Type:                           "",
		GCPMachineType:                 "",
		PrivatePowerManagementIP:       "",
		DisplayName:                    "",
		Perms2:                         MakePermType2(),
		IPAddress:                      "",
		PrivatePowerManagementPassword: "",
		Annotations:                    MakeKeyValuePairs(),
		FQName:                         []string{},
		IDPerms:                        MakeIdPermsType(),
		ParentUUID:                     "",
		Password:                       "",
		SSHKey:                         "",
		AwsInstanceType:                "",
		PrivateMachineProperties:       "",
		PrivateMachineState:            "",
		PrivatePowerManagementUsername: "",
		UUID:       "",
		Hostname:   "",
		MacAddress: "",
		Username:   "",
		AwsAmi:     "",
		GCPImage:   "",
	}
}

// InterfaceToNode makes Node from interface
func InterfaceToNode(iData interface{}) *Node {
	data := iData.(map[string]interface{})
	return &Node{
		Type: data["type"].(string),

		//{"title":"Machine Type","description":"Type of machine resource","default":"private","type":"string","permission":["create","update"],"enum":["private","virtual","aws","container","gcp"]}
		GCPMachineType: data["gcp_machine_type"].(string),

		//{"title":"Machine Type","default":"n1-standard-1","type":"string","permission":["create","update"]}
		PrivatePowerManagementIP: data["private_power_management_ip"].(string),

		//{"title":"Power Management IP","description":"IP address used for power management (IPMI)","default":"","type":"string","permission":["create","update"]}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		IPAddress: data["ip_address"].(string),

		//{"title":"IP Address","description":"IP Address","default":"","type":"string","permission":["create","update"]}
		PrivatePowerManagementPassword: data["private_power_management_password"].(string),

		//{"title":"Power Management UserPassword","description":"UserPassword for PowerManagement","default":"ADMIN","type":"string","permission":["create","update"]}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		Password: data["password"].(string),

		//{"title":"UserPassword","description":"UserPassword","default":"ADMIN","type":"string","permission":["create","update"]}
		SSHKey: data["ssh_key"].(string),

		//{"title":"SSH public key","description":"SSH Public Key","type":"string","permission":["create","update"]}
		AwsInstanceType: data["aws_instance_type"].(string),

		//{"title":"Instance Type","default":"t2.micro","type":"string","permission":["create","update"]}
		PrivateMachineProperties: data["private_machine_properties"].(string),

		//{"title":"Machine Properties","description":"Machine Properties from ironic","default":"","type":"string","permission":["create","update"]}
		PrivateMachineState: data["private_machine_state"].(string),

		//{"title":"Machine State","description":"Machine State","default":"enroll","type":"string","permission":["create","update"],"enum":["enroll","manageable","available","assigned"]}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		Hostname: data["hostname"].(string),

		//{"title":"Hostname","description":"Fully qualified host name","default":"","type":"string","permission":["create","update"]}
		MacAddress: data["mac_address"].(string),

		//{"title":"Interface MAC Address","description":"Provisioning Interface MAC Address","default":"","type":"string","permission":["create","update"]}
		Username: data["username"].(string),

		//{"title":"User Name","description":"User Name","default":"ADMIN","type":"string","permission":["create","update"]}
		AwsAmi: data["aws_ami"].(string),

		//{"title":"AMI","default":"ami-73f7da13","type":"string","permission":["create","update"]}
		GCPImage: data["gcp_image"].(string),

		//{"title":"Image","default":"ubuntu-os-cloud/ubuntu-1604-lts","type":"string","permission":["create","update"]}
		PrivatePowerManagementUsername: data["private_power_management_username"].(string),

		//{"title":"Power Management User Name","description":"User Name for PowerManagement","default":"ADMIN","type":"string","permission":["create","update"]}
		UUID: data["uuid"].(string),

		//{"type":"string"}

	}
}

// InterfaceToNodeSlice makes a slice of Node from interface
func InterfaceToNodeSlice(data interface{}) []*Node {
	list := data.([]interface{})
	result := MakeNodeSlice()
	for _, item := range list {
		result = append(result, InterfaceToNode(item))
	}
	return result
}

// MakeNodeSlice() makes a slice of Node
func MakeNodeSlice() []*Node {
	return []*Node{}
}