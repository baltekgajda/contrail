package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertGlobalVrouterConfigQuery = "insert into `global_vrouter_config` (`vxlan_network_identifier_mode`,`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`linklocal_service_entry`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`forwarding_mode`,`flow_export_rate`,`flow_aging_timeout`,`encapsulation`,`enable_security_logging`,`source_port`,`source_ip`,`ip_protocol`,`hashing_configured`,`destination_port`,`destination_ip`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateGlobalVrouterConfigQuery = "update `global_vrouter_config` set `vxlan_network_identifier_mode` = ?,`uuid` = ?,`share` = ?,`owner_access` = ?,`owner` = ?,`global_access` = ?,`parent_uuid` = ?,`parent_type` = ?,`linklocal_service_entry` = ?,`user_visible` = ?,`permissions_owner_access` = ?,`permissions_owner` = ?,`other_access` = ?,`group_access` = ?,`group` = ?,`last_modified` = ?,`enable` = ?,`description` = ?,`creator` = ?,`created` = ?,`fq_name` = ?,`forwarding_mode` = ?,`flow_export_rate` = ?,`flow_aging_timeout` = ?,`encapsulation` = ?,`enable_security_logging` = ?,`source_port` = ?,`source_ip` = ?,`ip_protocol` = ?,`hashing_configured` = ?,`destination_port` = ?,`destination_ip` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteGlobalVrouterConfigQuery = "delete from `global_vrouter_config` where uuid = ?"

// GlobalVrouterConfigFields is db columns for GlobalVrouterConfig
var GlobalVrouterConfigFields = []string{
	"vxlan_network_identifier_mode",
	"uuid",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"linklocal_service_entry",
	"user_visible",
	"permissions_owner_access",
	"permissions_owner",
	"other_access",
	"group_access",
	"group",
	"last_modified",
	"enable",
	"description",
	"creator",
	"created",
	"fq_name",
	"forwarding_mode",
	"flow_export_rate",
	"flow_aging_timeout",
	"encapsulation",
	"enable_security_logging",
	"source_port",
	"source_ip",
	"ip_protocol",
	"hashing_configured",
	"destination_port",
	"destination_ip",
	"display_name",
	"key_value_pair",
}

// GlobalVrouterConfigRefFields is db reference fields for GlobalVrouterConfig
var GlobalVrouterConfigRefFields = map[string][]string{}

// GlobalVrouterConfigBackRefFields is db back reference fields for GlobalVrouterConfig
var GlobalVrouterConfigBackRefFields = map[string][]string{

	"security_logging_object": {
		"uuid",
		"rule",
		"security_logging_object_rate",
		"share",
		"owner_access",
		"owner",
		"global_access",
		"parent_uuid",
		"parent_type",
		"user_visible",
		"permissions_owner_access",
		"permissions_owner",
		"other_access",
		"group_access",
		"group",
		"last_modified",
		"enable",
		"description",
		"creator",
		"created",
		"fq_name",
		"display_name",
		"key_value_pair",
	},
}

// CreateGlobalVrouterConfig inserts GlobalVrouterConfig to DB
func CreateGlobalVrouterConfig(tx *sql.Tx, model *models.GlobalVrouterConfig) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertGlobalVrouterConfigQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertGlobalVrouterConfigQuery,
	}).Debug("create query")
	_, err = stmt.Exec(string(model.VxlanNetworkIdentifierMode),
		string(model.UUID),
		common.MustJSON(model.Perms2.Share),
		int(model.Perms2.OwnerAccess),
		string(model.Perms2.Owner),
		int(model.Perms2.GlobalAccess),
		string(model.ParentUUID),
		string(model.ParentType),
		common.MustJSON(model.LinklocalServices.LinklocalServiceEntry),
		bool(model.IDPerms.UserVisible),
		int(model.IDPerms.Permissions.OwnerAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OtherAccess),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Group),
		string(model.IDPerms.LastModified),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Creator),
		string(model.IDPerms.Created),
		common.MustJSON(model.FQName),
		string(model.ForwardingMode),
		int(model.FlowExportRate),
		common.MustJSON(model.FlowAgingTimeoutList.FlowAgingTimeout),
		common.MustJSON(model.EncapsulationPriorities.Encapsulation),
		bool(model.EnableSecurityLogging),
		bool(model.EcmpHashingIncludeFields.SourcePort),
		bool(model.EcmpHashingIncludeFields.SourceIP),
		bool(model.EcmpHashingIncludeFields.IPProtocol),
		bool(model.EcmpHashingIncludeFields.HashingConfigured),
		bool(model.EcmpHashingIncludeFields.DestinationPort),
		bool(model.EcmpHashingIncludeFields.DestinationIP),
		string(model.DisplayName),
		common.MustJSON(model.Annotations.KeyValuePair))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanGlobalVrouterConfig(values map[string]interface{}) (*models.GlobalVrouterConfig, error) {
	m := models.MakeGlobalVrouterConfig()

	if value, ok := values["vxlan_network_identifier_mode"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VxlanNetworkIdentifierMode = models.VxlanNetworkIdentifierModeType(castedValue)

	}

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["global_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["parent_uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ParentUUID = castedValue

	}

	if value, ok := values["parent_type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ParentType = castedValue

	}

	if value, ok := values["linklocal_service_entry"]; ok {

		json.Unmarshal(value.([]byte), &m.LinklocalServices.LinklocalServiceEntry)

	}

	if value, ok := values["user_visible"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.UserVisible = castedValue

	}

	if value, ok := values["permissions_owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["permissions_owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["other_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Group = castedValue

	}

	if value, ok := values["last_modified"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["description"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Description = castedValue

	}

	if value, ok := values["creator"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Creator = castedValue

	}

	if value, ok := values["created"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Created = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["forwarding_mode"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ForwardingMode = models.ForwardingModeType(castedValue)

	}

	if value, ok := values["flow_export_rate"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.FlowExportRate = castedValue

	}

	if value, ok := values["flow_aging_timeout"]; ok {

		json.Unmarshal(value.([]byte), &m.FlowAgingTimeoutList.FlowAgingTimeout)

	}

	if value, ok := values["encapsulation"]; ok {

		json.Unmarshal(value.([]byte), &m.EncapsulationPriorities.Encapsulation)

	}

	if value, ok := values["enable_security_logging"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.EnableSecurityLogging = castedValue

	}

	if value, ok := values["source_port"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.SourcePort = castedValue

	}

	if value, ok := values["source_ip"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.SourceIP = castedValue

	}

	if value, ok := values["ip_protocol"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.IPProtocol = castedValue

	}

	if value, ok := values["hashing_configured"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.HashingConfigured = castedValue

	}

	if value, ok := values["destination_port"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.DestinationPort = castedValue

	}

	if value, ok := values["destination_ip"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.DestinationIP = castedValue

	}

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["backref_security_logging_object"]; ok {
		var childResources []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &childResources)
		for _, childResource := range childResources {
			childResourceMap, ok := childResource.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(childResourceMap["uuid"])
			if uuid == "" {
				continue
			}
			childModel := models.MakeSecurityLoggingObject()
			m.SecurityLoggingObjects = append(m.SecurityLoggingObjects, childModel)

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.UUID = castedValue

			}

			if propertyValue, ok := childResourceMap["rule"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.SecurityLoggingObjectRules.Rule)

			}

			if propertyValue, ok := childResourceMap["security_logging_object_rate"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.SecurityLoggingObjectRate = castedValue

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Perms2.OwnerAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.Perms2.Owner = castedValue

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Perms2.GlobalAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.ParentUUID = castedValue

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.ParentType = castedValue

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.IDPerms.UserVisible = castedValue

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Permissions.Owner = castedValue

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Permissions.Group = castedValue

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.LastModified = castedValue

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.IDPerms.Enable = castedValue

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Description = castedValue

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Creator = castedValue

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Created = castedValue

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.DisplayName = castedValue

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

		}
	}

	return m, nil
}

// ListGlobalVrouterConfig lists GlobalVrouterConfig with list spec.
func ListGlobalVrouterConfig(tx *sql.Tx, spec *common.ListSpec) ([]*models.GlobalVrouterConfig, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "global_vrouter_config"
	if spec.Fields == nil {
		spec.Fields = GlobalVrouterConfigFields
	}
	spec.RefFields = GlobalVrouterConfigRefFields
	spec.BackRefFields = GlobalVrouterConfigBackRefFields
	result := models.MakeGlobalVrouterConfigSlice()
	query, columns, values := common.BuildListQuery(spec)
	log.WithFields(log.Fields{
		"listSpec": spec,
		"query":    query,
	}).Debug("select query")
	rows, err = tx.Query(query, values...)
	if err != nil {
		return nil, errors.Wrap(err, "select query failed")
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "row error")
	}
	for rows.Next() {
		valuesMap := map[string]interface{}{}
		values := make([]interface{}, len(columns))
		valuesPointers := make([]interface{}, len(columns))
		for _, index := range columns {
			valuesPointers[index] = &values[index]
		}
		if err := rows.Scan(valuesPointers...); err != nil {
			return nil, errors.Wrap(err, "scan failed")
		}
		for column, index := range columns {
			val := valuesPointers[index].(*interface{})
			valuesMap[column] = *val
		}
		m, err := scanGlobalVrouterConfig(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// UpdateGlobalVrouterConfig updates a resource
func UpdateGlobalVrouterConfig(tx *sql.Tx, uuid string, model *models.GlobalVrouterConfig) error {
	//TODO(nati) support update
	return nil
}

// DeleteGlobalVrouterConfig deletes a resource
func DeleteGlobalVrouterConfig(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
	query := deleteGlobalVrouterConfigQuery
	var err error

	if auth.IsAdmin() {
		_, err = tx.Exec(query, uuid)
	} else {
		query += " and owner = ?"
		_, err = tx.Exec(query, uuid, auth.ProjectID())
	}

	if err != nil {
		return errors.Wrap(err, "delete failed")
	}

	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("deleted")
	return nil
}