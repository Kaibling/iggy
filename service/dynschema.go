package service

import (
	"context"

	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/pkg/config"
)

type dynSchemaRepo interface {
	IDQuery(query string) ([]string, error)
	AddFieldsToSchema(suffix string, vars []entity.DynamicField) error
	RemoveFieldsFromSchema(suffix, dynTabName string, fields []entity.DynamicField) error
	CreateTable(suffix string, dynTab entity.DynamicTable, vars []entity.DynamicField) error
}

type DynSchemaService struct {
	ctx  context.Context
	repo dynSchemaRepo
	cfg  config.Configuration
}

func NewDynSchemaService(ctx context.Context, dss dynSchemaRepo, cfg config.Configuration) *DynSchemaService {
	return &DynSchemaService{ctx: ctx, repo: dss, cfg: cfg}
}

func (rls *DynSchemaService) AddFields(suffix string, fields []entity.DynamicField) error {
	return rls.repo.AddFieldsToSchema(suffix, fields)
}

func (rls *DynSchemaService) DeleteFields(suffix, dynTabName string, fields []entity.DynamicField) error {
	return rls.repo.RemoveFieldsFromSchema(suffix, dynTabName, fields)
}

func (rls *DynSchemaService) CreateSchemaTable(suffix string, dynTab entity.DynamicTable, vars []entity.DynamicField) error { //nolint:lll
	return rls.repo.CreateTable(suffix, dynTab, vars)
}
