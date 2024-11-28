package repo

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/persistence/sqlcrepo"
)

type DynSchemaRepo struct {
	ctx      context.Context
	q        *sqlcrepo.Queries
	db       *pgxpool.Pool
	username string
}

func NewDynSchemaRepo(ctx context.Context, username string, dbPool *pgxpool.Pool) *DynSchemaRepo {
	return &DynSchemaRepo{
		ctx:      ctx,
		q:        sqlcrepo.New(dbPool),
		db:       dbPool,
		username: username,
	}
}

func (r *DynSchemaRepo) IDQuery(idQuery string) ([]string, error) {
	rows, err := r.db.Query(context.Background(), idQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var ids []string
	// Process the results
	for rows.Next() {
		var id string
		if err := rows.Scan(
			&id,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		ids = append(ids, id)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	defer rows.Close()

	return ids, nil
}

// func (r *DynSchemaRepo) RemoveSchemas(schemaIDs []string) error {
// 	return r.q.DeleteDynamicSchemas(r.ctx, schemaIDs)
// }

func (r *DynSchemaRepo) CreateTable(suffix string, dynSchema entity.DynamicTable, vars []entity.DynamicField) error {
	variables, err := variablesToSQL(vars)
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`CREATE TABLE "%s_%s" ( %s );`, suffix, dynSchema.Name, variables)

	rows, err := r.db.Query(context.Background(), q)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	defer rows.Close()

	return nil
}

func (r *DynSchemaRepo) Query(idQuery string) error {
	rows, err := r.db.Query(context.Background(), idQuery)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	defer rows.Close()

	return nil
}

func (r *DynSchemaRepo) AddFieldsToSchema(suffix string, vars []entity.DynamicField) error {
	tables := map[string][]entity.DynamicField{}
	for _, v := range vars {
		tables[v.DynamicTable.Name] = append(tables[v.DynamicTable.Name], v)
	}

	for k, vars := range tables {
		addVars := []string{}
		for _, v := range vars {
			addVars = append(addVars, v.Name+" "+v.VariableType)
		}

		q := fmt.Sprintf("ALTER TABLE %s_%s ADD %s;", suffix, k, strings.Join(addVars, ","))

		// todo sql errors are always null
		rows, err := r.db.Query(context.Background(), q)
		if err != nil {
			return fmt.Errorf("failed to execute query: %w", err)
		}

		defer rows.Close()
	}

	return nil
}

func (r *DynSchemaRepo) RemoveFieldsFromSchema(suffix, dynTabName string, vars []entity.DynamicField) error {
	addVars := []string{}
	for _, v := range vars {
		addVars = append(addVars, v.Name)
	}

	q := fmt.Sprintf("ALTER TABLE %s_%s DROP COLUMN  %s;", suffix, dynTabName, strings.Join(addVars, ","))

	// todo sql errors are always null
	rows, err := r.db.Query(context.Background(), q)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	defer rows.Close()

	return nil
}

func variablesToSQL(vars []entity.DynamicField) (string, error) {
	varQ := []string{}

	for _, v := range vars {
		varType, err := toSQLType(v.VariableType)
		if err != nil {
			return "", err
		}

		varQ = append(varQ, v.Name+" "+varType)
	}

	return strings.Join(varQ, ","), nil
}

func toSQLType(s string) (string, error) {
	switch s {
	case "integer":
		return "INT", nil
	case "text":
		return "TEXT", nil
	default:
		return "", fmt.Errorf("variable_type '%s' not valid", s) //nolint: err113
	}
}

// func removeSchemaField() error {
// ALTER TABLE table_name DROP COLUMN column_name;
// }

// func updateSchemaField() error {
// ALTER TABLE table_name
// RENAME COLUMN old_name to new_name;
// }

// func updateTable() error {

// }

// func deleteTable() error {
// format!("DROP TABLE {};", dyntable_name(name))
// }
