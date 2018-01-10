package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertOpenstackClusterQuery = "insert into `openstack_cluster` (`uuid`,`public_ip`,`public_gateway`,`provisioning_state`,`provisioning_start_time`,`provisioning_progress_stage`,`provisioning_progress`,`provisioning_log`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`openstack_webui`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`external_net_cidr`,`external_allocation_pool_start`,`external_allocation_pool_end`,`display_name`,`default_storage_backend_bond_interface_members`,`default_storage_access_bond_interface_members`,`default_performance_drives`,`default_osd_drives`,`default_journal_drives`,`default_capacity_drives`,`contrail_cluster_id`,`key_value_pair`,`admin_password`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateOpenstackClusterQuery = "update `openstack_cluster` set `uuid` = ?,`public_ip` = ?,`public_gateway` = ?,`provisioning_state` = ?,`provisioning_start_time` = ?,`provisioning_progress_stage` = ?,`provisioning_progress` = ?,`provisioning_log` = ?,`share` = ?,`owner_access` = ?,`owner` = ?,`global_access` = ?,`parent_uuid` = ?,`parent_type` = ?,`openstack_webui` = ?,`user_visible` = ?,`permissions_owner_access` = ?,`permissions_owner` = ?,`other_access` = ?,`group_access` = ?,`group` = ?,`last_modified` = ?,`enable` = ?,`description` = ?,`creator` = ?,`created` = ?,`fq_name` = ?,`external_net_cidr` = ?,`external_allocation_pool_start` = ?,`external_allocation_pool_end` = ?,`display_name` = ?,`default_storage_backend_bond_interface_members` = ?,`default_storage_access_bond_interface_members` = ?,`default_performance_drives` = ?,`default_osd_drives` = ?,`default_journal_drives` = ?,`default_capacity_drives` = ?,`contrail_cluster_id` = ?,`key_value_pair` = ?,`admin_password` = ?;"
const deleteOpenstackClusterQuery = "delete from `openstack_cluster` where uuid = ?"

// OpenstackClusterFields is db columns for OpenstackCluster
var OpenstackClusterFields = []string{
	"uuid",
	"public_ip",
	"public_gateway",
	"provisioning_state",
	"provisioning_start_time",
	"provisioning_progress_stage",
	"provisioning_progress",
	"provisioning_log",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"openstack_webui",
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
	"external_net_cidr",
	"external_allocation_pool_start",
	"external_allocation_pool_end",
	"display_name",
	"default_storage_backend_bond_interface_members",
	"default_storage_access_bond_interface_members",
	"default_performance_drives",
	"default_osd_drives",
	"default_journal_drives",
	"default_capacity_drives",
	"contrail_cluster_id",
	"key_value_pair",
	"admin_password",
}

// OpenstackClusterRefFields is db reference fields for OpenstackCluster
var OpenstackClusterRefFields = map[string][]string{}

// OpenstackClusterBackRefFields is db back reference fields for OpenstackCluster
var OpenstackClusterBackRefFields = map[string][]string{}

// CreateOpenstackCluster inserts OpenstackCluster to DB
func CreateOpenstackCluster(tx *sql.Tx, model *models.OpenstackCluster) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertOpenstackClusterQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertOpenstackClusterQuery,
	}).Debug("create query")
	_, err = stmt.Exec(string(model.UUID),
		string(model.PublicIP),
		string(model.PublicGateway),
		string(model.ProvisioningState),
		string(model.ProvisioningStartTime),
		string(model.ProvisioningProgressStage),
		int(model.ProvisioningProgress),
		string(model.ProvisioningLog),
		common.MustJSON(model.Perms2.Share),
		int(model.Perms2.OwnerAccess),
		string(model.Perms2.Owner),
		int(model.Perms2.GlobalAccess),
		string(model.ParentUUID),
		string(model.ParentType),
		string(model.OpenstackWebui),
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
		string(model.ExternalNetCidr),
		string(model.ExternalAllocationPoolStart),
		string(model.ExternalAllocationPoolEnd),
		string(model.DisplayName),
		string(model.DefaultStorageBackendBondInterfaceMembers),
		string(model.DefaultStorageAccessBondInterfaceMembers),
		string(model.DefaultPerformanceDrives),
		string(model.DefaultOsdDrives),
		string(model.DefaultJournalDrives),
		string(model.DefaultCapacityDrives),
		string(model.ContrailClusterID),
		common.MustJSON(model.Annotations.KeyValuePair),
		string(model.AdminPassword))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanOpenstackCluster(values map[string]interface{}) (*models.OpenstackCluster, error) {
	m := models.MakeOpenstackCluster()

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["public_ip"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PublicIP = castedValue

	}

	if value, ok := values["public_gateway"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PublicGateway = castedValue

	}

	if value, ok := values["provisioning_state"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ProvisioningState = castedValue

	}

	if value, ok := values["provisioning_start_time"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ProvisioningStartTime = castedValue

	}

	if value, ok := values["provisioning_progress_stage"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ProvisioningProgressStage = castedValue

	}

	if value, ok := values["provisioning_progress"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.ProvisioningProgress = castedValue

	}

	if value, ok := values["provisioning_log"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ProvisioningLog = castedValue

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

	if value, ok := values["openstack_webui"]; ok {

		castedValue := common.InterfaceToString(value)

		m.OpenstackWebui = castedValue

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

	if value, ok := values["external_net_cidr"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ExternalNetCidr = castedValue

	}

	if value, ok := values["external_allocation_pool_start"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ExternalAllocationPoolStart = castedValue

	}

	if value, ok := values["external_allocation_pool_end"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ExternalAllocationPoolEnd = castedValue

	}

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["default_storage_backend_bond_interface_members"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DefaultStorageBackendBondInterfaceMembers = castedValue

	}

	if value, ok := values["default_storage_access_bond_interface_members"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DefaultStorageAccessBondInterfaceMembers = castedValue

	}

	if value, ok := values["default_performance_drives"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DefaultPerformanceDrives = castedValue

	}

	if value, ok := values["default_osd_drives"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DefaultOsdDrives = castedValue

	}

	if value, ok := values["default_journal_drives"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DefaultJournalDrives = castedValue

	}

	if value, ok := values["default_capacity_drives"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DefaultCapacityDrives = castedValue

	}

	if value, ok := values["contrail_cluster_id"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ContrailClusterID = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["admin_password"]; ok {

		castedValue := common.InterfaceToString(value)

		m.AdminPassword = castedValue

	}

	return m, nil
}

// ListOpenstackCluster lists OpenstackCluster with list spec.
func ListOpenstackCluster(tx *sql.Tx, spec *common.ListSpec) ([]*models.OpenstackCluster, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "openstack_cluster"
	if spec.Fields == nil {
		spec.Fields = OpenstackClusterFields
	}
	spec.RefFields = OpenstackClusterRefFields
	spec.BackRefFields = OpenstackClusterBackRefFields
	result := models.MakeOpenstackClusterSlice()
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
		m, err := scanOpenstackCluster(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// UpdateOpenstackCluster updates a resource
func UpdateOpenstackCluster(tx *sql.Tx, uuid string, model *models.OpenstackCluster) error {
	//TODO(nati) support update
	return nil
}

// DeleteOpenstackCluster deletes a resource
func DeleteOpenstackCluster(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
	query := deleteOpenstackClusterQuery
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