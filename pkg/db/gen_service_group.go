// nolint
package db

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

// ServiceGroupFields is db columns for ServiceGroup
var ServiceGroupFields = []string{
	"uuid",
	"firewall_service",
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
	"configuration_version",
	"key_value_pair",
}

// ServiceGroupRefFields is db reference fields for ServiceGroup
var ServiceGroupRefFields = map[string][]string{}

// ServiceGroupBackRefFields is db back reference fields for ServiceGroup
var ServiceGroupBackRefFields = map[string][]string{}

// ServiceGroupParentTypes is possible parents for ServiceGroup
var ServiceGroupParents = []string{

	"project",

	"policy_management",
}

// CreateServiceGroup inserts ServiceGroup to DB
// nolint
func (db *DB) createServiceGroup(
	ctx context.Context,
	request *models.CreateServiceGroupRequest) error {
	qb := db.queryBuilders["service_group"]
	tx := GetTransaction(ctx)
	model := request.ServiceGroup
	_, err := tx.ExecContext(ctx, qb.CreateQuery(), string(model.GetUUID()),
		common.MustJSON(model.GetServiceGroupFirewallServiceList().GetFirewallService()),
		common.MustJSON(model.GetPerms2().GetShare()),
		int(model.GetPerms2().GetOwnerAccess()),
		string(model.GetPerms2().GetOwner()),
		int(model.GetPerms2().GetGlobalAccess()),
		string(model.GetParentUUID()),
		string(model.GetParentType()),
		bool(model.GetIDPerms().GetUserVisible()),
		int(model.GetIDPerms().GetPermissions().GetOwnerAccess()),
		string(model.GetIDPerms().GetPermissions().GetOwner()),
		int(model.GetIDPerms().GetPermissions().GetOtherAccess()),
		int(model.GetIDPerms().GetPermissions().GetGroupAccess()),
		string(model.GetIDPerms().GetPermissions().GetGroup()),
		string(model.GetIDPerms().GetLastModified()),
		bool(model.GetIDPerms().GetEnable()),
		string(model.GetIDPerms().GetDescription()),
		string(model.GetIDPerms().GetCreator()),
		string(model.GetIDPerms().GetCreated()),
		common.MustJSON(model.GetFQName()),
		string(model.GetDisplayName()),
		int(model.GetConfigurationVersion()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
		log.WithFields(log.Fields{
			"model": model,
			"err":   err}).Debug("create failed")
		return errors.Wrap(err, "create failed")
	}

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "service_group",
		FQName: model.FQName,
	}
	err = db.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = db.CreateSharing(tx, "service_group", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanServiceGroup(values map[string]interface{}) (*models.ServiceGroup, error) {
	m := models.MakeServiceGroup()

	if value, ok := values["uuid"]; ok {

		m.UUID = common.InterfaceToString(value)

	}

	if value, ok := values["firewall_service"]; ok {

		json.Unmarshal(value.([]byte), &m.ServiceGroupFirewallServiceList.FirewallService)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["owner_access"]; ok {

		m.Perms2.OwnerAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["owner"]; ok {

		m.Perms2.Owner = common.InterfaceToString(value)

	}

	if value, ok := values["global_access"]; ok {

		m.Perms2.GlobalAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["parent_uuid"]; ok {

		m.ParentUUID = common.InterfaceToString(value)

	}

	if value, ok := values["parent_type"]; ok {

		m.ParentType = common.InterfaceToString(value)

	}

	if value, ok := values["user_visible"]; ok {

		m.IDPerms.UserVisible = common.InterfaceToBool(value)

	}

	if value, ok := values["permissions_owner_access"]; ok {

		m.IDPerms.Permissions.OwnerAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["permissions_owner"]; ok {

		m.IDPerms.Permissions.Owner = common.InterfaceToString(value)

	}

	if value, ok := values["other_access"]; ok {

		m.IDPerms.Permissions.OtherAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["group_access"]; ok {

		m.IDPerms.Permissions.GroupAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["group"]; ok {

		m.IDPerms.Permissions.Group = common.InterfaceToString(value)

	}

	if value, ok := values["last_modified"]; ok {

		m.IDPerms.LastModified = common.InterfaceToString(value)

	}

	if value, ok := values["enable"]; ok {

		m.IDPerms.Enable = common.InterfaceToBool(value)

	}

	if value, ok := values["description"]; ok {

		m.IDPerms.Description = common.InterfaceToString(value)

	}

	if value, ok := values["creator"]; ok {

		m.IDPerms.Creator = common.InterfaceToString(value)

	}

	if value, ok := values["created"]; ok {

		m.IDPerms.Created = common.InterfaceToString(value)

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["display_name"]; ok {

		m.DisplayName = common.InterfaceToString(value)

	}

	if value, ok := values["configuration_version"]; ok {

		m.ConfigurationVersion = common.InterfaceToInt64(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	return m, nil
}

// ListServiceGroup lists ServiceGroup with list spec.
func (db *DB) listServiceGroup(ctx context.Context, request *models.ListServiceGroupRequest) (response *models.ListServiceGroupResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)

	qb := db.queryBuilders["service_group"]

	auth := common.GetAuthCTX(ctx)
	spec := request.Spec
	result := []*models.ServiceGroup{}

	if spec.ParentFQName != nil {
		parentMetaData, err := db.GetMetaData(tx, "", spec.ParentFQName)
		if err != nil {
			return nil, errors.Wrap(err, "can't find parents")
		}
		spec.Filters = models.AppendFilter(spec.Filters, "parent_uuid", parentMetaData.UUID)
	}
	query, columns, values := qb.ListQuery(auth, spec)
	rows, err = tx.QueryContext(ctx, query, values...)
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
		m, err := scanServiceGroup(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListServiceGroupResponse{
		ServiceGroups: result,
	}
	return response, nil
}

// UpdateServiceGroup updates a resource
func (db *DB) updateServiceGroup(
	ctx context.Context,
	request *models.UpdateServiceGroupRequest,
) error {
	//TODO
	return nil
}

// DeleteServiceGroup deletes a resource
func (db *DB) deleteServiceGroup(
	ctx context.Context,
	request *models.DeleteServiceGroupRequest) error {
	qb := db.queryBuilders["service_group"]

	selectQuery := qb.SelectForDeleteQuery()
	deleteQuery := qb.DeleteQuery()

	var err error
	var count int
	uuid := request.ID
	tx := GetTransaction(ctx)
	auth := common.GetAuthCTX(ctx)
	if auth.IsAdmin() {
		row := tx.QueryRowContext(ctx, selectQuery, uuid)
		if err != nil {
			return errors.Wrap(err, "not found")
		}
		row.Scan(&count)
		if count == 0 {
			return errors.New("Not found")
		}
		_, err = tx.ExecContext(ctx, deleteQuery, uuid)
	} else {
		deleteQuery += " and owner = ?"
		selectQuery += " and owner = ?"
		row := tx.QueryRowContext(ctx, selectQuery, uuid, auth.ProjectID())
		if err != nil {
			return errors.Wrap(err, "not found")
		}
		row.Scan(&count)
		if count == 0 {
			return errors.New("Not found")
		}
		_, err = tx.ExecContext(ctx, deleteQuery, uuid, auth.ProjectID())
	}

	if err != nil {
		return errors.Wrap(err, "delete failed")
	}

	err = db.DeleteMetaData(tx, uuid)
	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("deleted")
	return err
}

//CreateServiceGroup handle a Create API
// nolint
func (db *DB) CreateServiceGroup(
	ctx context.Context,
	request *models.CreateServiceGroupRequest) (*models.CreateServiceGroupResponse, error) {
	model := request.ServiceGroup
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createServiceGroup(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_group",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateServiceGroupResponse{
		ServiceGroup: request.ServiceGroup,
	}, nil
}

//UpdateServiceGroup handles a Update request.
func (db *DB) UpdateServiceGroup(
	ctx context.Context,
	request *models.UpdateServiceGroupRequest) (*models.UpdateServiceGroupResponse, error) {
	model := request.ServiceGroup
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateServiceGroup(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_group",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateServiceGroupResponse{
		ServiceGroup: model,
	}, nil
}

//DeleteServiceGroup delete a resource.
func (db *DB) DeleteServiceGroup(ctx context.Context, request *models.DeleteServiceGroupRequest) (*models.DeleteServiceGroupResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteServiceGroup(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteServiceGroupResponse{
		ID: request.ID,
	}, nil
}

//GetServiceGroup a Get request.
func (db *DB) GetServiceGroup(ctx context.Context, request *models.GetServiceGroupRequest) (response *models.GetServiceGroupResponse, err error) {
	spec := &models.ListSpec{
		Limit:  1,
		Detail: true,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListServiceGroupRequest{
		Spec: spec,
	}
	var result *models.ListServiceGroupResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listServiceGroup(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.ServiceGroups) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetServiceGroupResponse{
		ServiceGroup: result.ServiceGroups[0],
	}
	return response, nil
}

//ListServiceGroup handles a List service Request.
// nolint
func (db *DB) ListServiceGroup(
	ctx context.Context,
	request *models.ListServiceGroupRequest) (response *models.ListServiceGroupResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listServiceGroup(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
