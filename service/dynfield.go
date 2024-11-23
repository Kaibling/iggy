package service

import (
	"context"

	"github.com/kaibling/apiforge/lib/utils"
	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/pkg/config"
)

type dynFieldRepo interface {
	IDQuery(query string) ([]string, error)
	AddFields(newVars []entity.NewDynamicField) ([]entity.DynamicField, error)
	RemoveDynField(varIDs []string) error
	FetchDynamicFields(varIDs []string) ([]entity.DynamicField, error)
	FetchVarsByTabID(dynTabID string) ([]entity.DynamicField, error)
	FetchVarsByDynFieldName(dynTabName string) ([]entity.DynamicField, error)
}

type DynFieldService struct {
	ctx           context.Context
	repo          dynFieldRepo
	schemaService *DynSchemaService
	cfg           config.Configuration
}

func NewDynFieldService(ctx context.Context, dfs dynFieldRepo, dss *DynSchemaService, cfg config.Configuration) *DynFieldService {
	return &DynFieldService{ctx: ctx, repo: dfs, cfg: cfg, schemaService: dss}
}

func (rls *DynFieldService) AddFields(newVars []entity.NewDynamicField) error {
	for i := range newVars {
		newVars[i].ID = utils.NewULID().String()
	}

	newFields, err := rls.repo.AddFields(newVars)
	if err != nil {
		return err
	}

	return rls.schemaService.AddFields(suffix, newFields)
}

func (rls *DynFieldService) RemoveVars(dynTabName string, fieldIDs []string) error {
	fields, err := rls.FetchFieldsByIDs(fieldIDs)
	if err != nil {
		return err
	}

	// remove fields from db
	if err := rls.repo.RemoveDynField(fieldIDs); err != nil {
		return err
	}
	// remove fields from schema
	return rls.schemaService.DeleteFields(suffix, dynTabName, fields)
}

func (rls *DynFieldService) FetchVars(dynTabID string) ([]entity.DynamicField, error) {
	return rls.repo.FetchVarsByTabID(dynTabID)
}

func (rls *DynFieldService) FetchFieldsByIDs(fieldIDs []string) ([]entity.DynamicField, error) {
	return rls.repo.FetchDynamicFields(fieldIDs)
}

func (rls *DynFieldService) FetchVarsByName(dynTabName string) ([]entity.DynamicField, error) {
	return rls.repo.FetchVarsByDynFieldName(dynTabName)
}

// func (rls *DynFieldService) CreateFieldTable(suffix string, dynTab entity.DynamicTable, vars []entity.DynamicField) error {
// 	return rls.repo.CreateTable(suffix, dynTab, vars)
// }
