package service

import (
	"context"

	"github.com/kaibling/apiforge/lib/utils"
	"github.com/kaibling/apiforge/params"
	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/pkg/config"
)

type dynTabRepo interface {
	FetchDynamicTables(ids []string) ([]*entity.DynamicTable, error)
	CreateDynamicTables(newModels []entity.NewDynamicTable) ([]*entity.DynamicTable, error)
	IDQuery(query string) ([]string, error)
	AddDynTabVars(dynTabID string, newVars []entity.NewDynamicTableVariable) error
	RemoveDynTabVars(varIDs []string) error
	FetchVars(dynTabID string) ([]*entity.DynamicTableVariable, error)
}

type DynTabService struct {
	ctx  context.Context
	repo dynTabRepo
	cfg  config.Configuration
}

func NewDynTabService(ctx context.Context, u dynTabRepo, cfg config.Configuration) *DynTabService {
	return &DynTabService{ctx: ctx, repo: u, cfg: cfg}
}

func (rls *DynTabService) FetchDynamicTables(ids []string) ([]*entity.DynamicTable, error) {
	return rls.repo.FetchDynamicTables(ids)
}

func (rls *DynTabService) CreateDynamicTables(newEntities []entity.NewDynamicTable) ([]*entity.DynamicTable, error) {
	return rls.repo.CreateDynamicTables(newEntities)
}

func (rls *DynTabService) FetchByPagination(qp params.Pagination) ([]*entity.DynamicTable, params.Pagination, error) {
	p := NewPagination(qp, "dynamic-tables")

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

func (rls *DynTabService) AddVars(dynTabID string, newVars []entity.NewDynamicTableVariable) error {
	for i := range newVars {
		newVars[i].ID = utils.NewULID().String()
	}

	return rls.repo.AddDynTabVars(dynTabID, newVars)
}

func (rls *DynTabService) RemoveVars(varIDs []string) error {
	return rls.repo.RemoveDynTabVars(varIDs)
}

func (rls *DynTabService) FetchVars(dynTabID string) ([]*entity.DynamicTableVariable, error) {
	return rls.repo.FetchVars(dynTabID)
}
