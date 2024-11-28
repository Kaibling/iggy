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
	// AddDynSchema(newVars []entity.NewDynamicSchema) ([]entity.DynamicField, error)
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

func (rls *DynTabService) CreateDynamicTables(newEntities []entity.NewDynamicTable, dss *DynSchemaService) ([]entity.DynamicTable, error) { //nolint:lll
	for i := range newEntities {
		newEntities[i].ID = utils.NewULID().String()
	}

	ndts, err := rls.repo.CreateDynamicTables(newEntities)
	if err != nil {
		return nil, err
	}

	for _, ndt := range ndts {
		if err := dss.CreateSchemaTable(suffix, ndt, []entity.DynamicField{}); err != nil {
			return nil, err
		}
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

// func (rls *DynTabService) FetchDynFields(dynTabID string) ([]entity.DynamicField, error) {
// 	return rls.repo.FetchVarsByTab(dynTabID)
// }

// func (rls *DynTabService) FetchDynFieldsByTabName(dynTabName string) ([]entity.DynamicField, error) {
// 	return rls.repo.FetchVarsByDynSchemaName(dynTabName)
// }

// func (rls *DynTabService) AddVars(newVars []entity.NewDynamicSchema) error {
// 	for i := range newVars {
// 		newVars[i].ID = utils.NewULID().String()
// 	}

// 	newVariables, err := rls.repo.AddDynSchema(newVars)
// 	if err != nil {
// 		return err
// 	}

// 	return rls.repo.AddSchemaField(suffix, newVariables)
// }

// func (rls *DynTabService) RemoveVars(varIDs []string) error {
// 	return rls.repo.RemoveDynSchema(varIDs)
// }

// func (rls *DynTabService) FetchVars(dynTabID string) ([]entity.DynamicField, error) {
// 	return rls.repo.FetchVarsByTab(dynTabID)
// }

// func (rls *DynTabService) CreateObject(obj map[string]any, dynTabName string) error {
// 	// get columns of dynTab
// 	vars, err := rls.repo.FetchVarsByDynTabName(dynTabName)
// 	if err != nil {
// 		return err
// 	}

// 	// build insert statement
// 	query, err := dynTabObjectInsertSQL(obj, dynTabName, vars)
// 	if err != nil {
// 		return err
// 	}

// 	return rls.repo.Query(query)
// }

// func dynTabObjectInsertSQL(obj map[string]any, dynTabName string, vars []entity.DynamicField) (string, error) {
// 	if err := validateObj(obj, vars); err != nil {
// 		return "", err
// 	}

// 	types := map[string]string{}
// 	for _, v := range vars {
// 		types[v.Name] = v.VariableType
// 	}

// 	columns := []string{}
// 	values := []string{}

// 	for k, v := range obj {
// 		columns = append(columns, k)

// 		if types[k] == "text" {
// 			values = append(values, fmt.Sprintf(`'%s'`, v))
// 		} else {
// 			values = append(values, fmt.Sprintf("%s", v))
// 		}
// 	}

// 	columnStr := strings.Join(columns, `","`)
// 	valuesStr := strings.Join(values, ",")

// 	q := fmt.Sprintf(`INSERT INTO "%s" (
// 		"%s"
// 	  ) VALUES (
// 		%s
// 	  );`, rawDynTabName(dynTabName), columnStr, valuesStr)

// 	return q, nil
// }

// func validateObj(obj map[string]any, vars []entity.DynamicField) error {
// 	missing := []string{}

// 	for _, v := range vars {
// 		if _, ok := obj[v.Name]; !ok {
// 			missing = append(missing, v.Name)
// 		}
// 	}

// 	if len(missing) == 0 {
// 		return nil
// 	}

// 	return errors.New(strings.Join(missing, " ,")) //nolint: err113
// }

// func rawDynTabName(dynTabName string) string {
// 	return suffix + "_" + dynTabName
// }
