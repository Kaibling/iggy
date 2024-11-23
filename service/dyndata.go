package service

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/pkg/config"
)

type dynDataRepo interface {
	Query(query string) error
}

type DynDataService struct {
	ctx           context.Context
	repo          dynDataRepo
	schemaService *DynSchemaService
	fieldService  *DynFieldService
	cfg           config.Configuration
}

func NewDynDataService(ctx context.Context, repo dynDataRepo, schemaService *DynSchemaService, fieldService *DynFieldService, cfg config.Configuration) *DynDataService {
	return &DynDataService{ctx, repo, schemaService, fieldService, cfg}
}

func (rls *DynDataService) CreateObject(obj map[string]any, dynTabName string) error {
	// get columns of dynTab
	vars, err := rls.fieldService.FetchVarsByName(dynTabName)
	if err != nil {
		return err
	}

	// build insert statement
	query, err := dynTabObjectInsertSQL(obj, dynTabName, vars)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println(query)
	return rls.repo.Query(query)
}

func dynTabObjectInsertSQL(obj map[string]any, dynTabName string, vars []entity.DynamicField) (string, error) {
	if err := validateObj(obj, vars); err != nil {
		return "", err
	}

	types := map[string]string{}
	for _, v := range vars {
		types[v.Name] = v.VariableType
	}

	columns := []string{}
	values := []string{}

	for k, v := range obj {
		columns = append(columns, k)

		if types[k] == "text" {
			values = append(values, fmt.Sprintf(`'%s'`, v))
		} else {
			values = append(values, fmt.Sprintf("%s", v))
		}
	}

	columnStr := strings.Join(columns, `","`)
	valuesStr := strings.Join(values, ",")

	q := fmt.Sprintf(`INSERT INTO "%s" (
		"%s"
	  ) VALUES (
		%s
	  );`, rawDynDataName(dynTabName), columnStr, valuesStr)

	return q, nil
}

func validateObj(obj map[string]any, vars []entity.DynamicField) error {
	missing := []string{}

	for _, v := range vars {
		if _, ok := obj[v.Name]; !ok {
			missing = append(missing, v.Name)
		}
	}

	if len(missing) == 0 {
		return nil
	}

	return errors.New("object validation failed: " + strings.Join(missing, " ,")) //nolint: err113
}

func rawDynDataName(dynTabName string) string {
	return suffix + "_" + dynTabName
}
