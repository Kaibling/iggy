package service

import (
	"context"

	"github.com/kaibling/apiforge/lib/utils"
	"github.com/kaibling/apiforge/params"
	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/pkg/config"
)

const suffix = "custom"

type dynTabRepo interface {
	FetchDynamicTables(ids []string) ([]entity.DynamicTable, error)
	CreateDynamicTables(newModels []entity.NewDynamicTable) ([]entity.DynamicTable, error)
	IDQuery(query string) ([]string, error)
	AddDynTabVars(newVars []entity.NewDynamicTableVariable) ([]entity.DynamicTableVariable, error)
	RemoveDynTabVars(varIDs []string) error
	FetchVarsByIDs(varIDs []string) ([]entity.DynamicTableVariable, error)
	FetchVarsByTab(dynTabID string) ([]entity.DynamicTableVariable, error)

	CreateTable(suffix string, dynTab entity.DynamicTable, vars []entity.DynamicTableVariable) error
	AddSchemaField(suffix string, vars []entity.DynamicTableVariable) error
}

type DynTabService struct {
	ctx  context.Context
	repo dynTabRepo
	cfg  config.Configuration
}

func NewDynTabService(ctx context.Context, u dynTabRepo, cfg config.Configuration) *DynTabService {
	return &DynTabService{ctx: ctx, repo: u, cfg: cfg}
}

func (rls *DynTabService) FetchDynamicTables(ids []string) ([]entity.DynamicTable, error) {
	return rls.repo.FetchDynamicTables(ids)
}

func (rls *DynTabService) CreateDynamicTables(newEntities []entity.NewDynamicTable) ([]entity.DynamicTable, error) {
	for i := range newEntities {
		newEntities[i].ID = utils.NewULID().String()
	}

	ndts, err := rls.repo.CreateDynamicTables(newEntities)
	if err != nil {
		return nil, err
	}

	for _, ndt := range ndts {
		rls.repo.CreateTable(suffix, ndt, []entity.DynamicTableVariable{})
	}
	return ndts, nil
}

func (rls *DynTabService) FetchByPagination(qp params.Pagination) ([]entity.DynamicTable, params.Pagination, error) {
	p := NewPagination(qp, "dynamic_tables")

	idQuery := p.GetCursorSQL()

	ids, err := rls.repo.IDQuery(idQuery)
	if err != nil {
		return nil, params.Pagination{}, err
	}

	ids, pag := p.FinishPagination(ids)

	loadedRuns, err := rls.FetchDynamicTables(ids)
	if err != nil {
		return nil, params.Pagination{}, err
	}

	return loadedRuns, pag, nil
}

func (rls *DynTabService) AddVars(newVars []entity.NewDynamicTableVariable) error {
	for i := range newVars {
		newVars[i].ID = utils.NewULID().String()
	}

	newVariables, err := rls.repo.AddDynTabVars(newVars)
	if err != nil {
		return err
	}

	return rls.repo.AddSchemaField(suffix, newVariables)
}

func (rls *DynTabService) RemoveVars(varIDs []string) error {
	return rls.repo.RemoveDynTabVars(varIDs)
}

func (rls *DynTabService) FetchVars(dynTabID string) ([]entity.DynamicTableVariable, error) {
	return rls.repo.FetchVarsByTab(dynTabID)
}
