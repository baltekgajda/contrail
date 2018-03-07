// nolint
package db

import (
	"context"
	"github.com/satori/go.uuid"
	"testing"
	"time"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/pkg/errors"
)

//For skip import error.
var _ = errors.New("")

func TestContrailWebuiNode(t *testing.T) {
	t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	model := models.MakeContrailWebuiNode()
	model.UUID = uuid.NewV4().String()
	model.FQName = []string{"default", "default-domain", model.UUID}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var NodeCreateRef []*models.ContrailWebuiNodeNodeRef
	var NodeRefModel *models.Node

	NodeRefUUID := uuid.NewV4().String()
	NodeRefUUID1 := uuid.NewV4().String()
	NodeRefUUID2 := uuid.NewV4().String()

	NodeRefModel = models.MakeNode()
	NodeRefModel.UUID = NodeRefUUID
	NodeRefModel.FQName = []string{"test", NodeRefUUID}
	_, err = db.CreateNode(ctx, &models.CreateNodeRequest{
		Node: NodeRefModel,
	})
	NodeRefModel.UUID = NodeRefUUID1
	NodeRefModel.FQName = []string{"test", NodeRefUUID1}
	_, err = db.CreateNode(ctx, &models.CreateNodeRequest{
		Node: NodeRefModel,
	})
	NodeRefModel.UUID = NodeRefUUID2
	NodeRefModel.FQName = []string{"test", NodeRefUUID2}
	_, err = db.CreateNode(ctx, &models.CreateNodeRequest{
		Node: NodeRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	NodeCreateRef = append(NodeCreateRef,
		&models.ContrailWebuiNodeNodeRef{UUID: NodeRefUUID, To: []string{"test", NodeRefUUID}})
	NodeCreateRef = append(NodeCreateRef,
		&models.ContrailWebuiNodeNodeRef{UUID: NodeRefUUID2, To: []string{"test", NodeRefUUID2}})
	model.NodeRefs = NodeCreateRef

	//create project to which resource is shared
	projectModel := models.MakeProject()

	projectModel.UUID = uuid.NewV4().String()
	projectModel.FQName = []string{"default-domain-test", projectModel.UUID}
	projectModel.Perms2.Owner = "admin"

	var createShare []*models.ShareType
	createShare = append(createShare, &models.ShareType{Tenant: "default-domain-test:" + projectModel.UUID, TenantAccess: 7})
	model.Perms2.Share = createShare

	_, err = db.CreateProject(ctx, &models.CreateProjectRequest{
		Project: projectModel,
	})
	if err != nil {
		t.Fatal("project create failed", err)
	}

	_, err = db.CreateContrailWebuiNode(ctx,
		&models.CreateContrailWebuiNodeRequest{
			ContrailWebuiNode: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	response, err := db.ListContrailWebuiNode(ctx, &models.ListContrailWebuiNodeRequest{
		Spec: &models.ListSpec{Limit: 1,
			Filters: []*models.Filter{
				&models.Filter{
					Key:    "uuid",
					Values: []string{model.UUID},
				},
			},
		}})
	if err != nil {
		t.Fatal("list failed", err)
	}
	if len(response.ContrailWebuiNodes) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteContrailWebuiNode(ctxDemo,
		&models.DeleteContrailWebuiNodeRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateContrailWebuiNode(ctx,
		&models.CreateContrailWebuiNodeRequest{
			ContrailWebuiNode: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteContrailWebuiNode(ctx,
		&models.DeleteContrailWebuiNodeRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	_, err = db.GetContrailWebuiNode(ctx, &models.GetContrailWebuiNodeRequest{
		ID: model.UUID})
	if err == nil {
		t.Fatal("expected not found error")
	}

	//Delete the project created for sharing
	_, err = db.DeleteProject(ctx, &models.DeleteProjectRequest{
		ID: projectModel.UUID})
	if err != nil {
		t.Fatal("delete project failed", err)
	}
	return
}
