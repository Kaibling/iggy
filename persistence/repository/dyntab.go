package repo

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/persistence/sqlcrepo"
)

type DynTabRepo struct {
	ctx      context.Context
	q        *sqlcrepo.Queries
	db       *pgxpool.Pool
	username string
}

func NewDynTabRepo(ctx context.Context, username string, dbPool *pgxpool.Pool) *DynTabRepo {
	return &DynTabRepo{
		ctx:      ctx,
		q:        sqlcrepo.New(dbPool),
		db:       dbPool,
		username: username,
	}
}
func (r *DynTabRepo) CreateDynamicTables(newModels []entity.NewDynamicTable) ([]entity.DynamicTable, error) {
	newDynamicTableIDs := []string{}

	for _, newModel := range newModels {
		newDynamicTableID, err := r.q.CreateDynamicTable(r.ctx, sqlcrepo.CreateDynamicTableParams{
			ID:        newModel.ID,
			TableName: newModel.Name,
			CreatedAt: pgtype.Timestamp{
				Time:  time.Now(),
				Valid: true,
			},
			CreatedBy: r.username,
			ModifiedAt: pgtype.Timestamp{
				Time:  time.Now(),
				Valid: true,
			},
			ModifiedBy: r.username,
		})
		if err != nil {
			return nil, err
		}

		newDynamicTableIDs = append(newDynamicTableIDs, newDynamicTableID)
	}

	return r.FetchDynamicTables(newDynamicTableIDs)
}

func (r *DynTabRepo) FetchDynamicTables(ids []string) ([]entity.DynamicTable, error) {
	dbdts, err := r.q.FetchDynamicTables(r.ctx, ids)
	if err != nil {
		return nil, err
	}

	fetchedDynamicTables := []entity.DynamicTable{}

	for _, dbDynTable := range dbdts {
		fetchedDynamicTables = append(fetchedDynamicTables, entity.DynamicTable{
			ID:   dbDynTable.ID,
			Name: dbDynTable.TableName,
			Meta: entity.MetaData{
				CreatedAt:  dbDynTable.CreatedAt.Time,
				CreatedBy:  dbDynTable.CreatedBy,
				ModifiedAt: dbDynTable.ModifiedAt.Time,
				ModifiedBy: dbDynTable.ModifiedBy,
			},
		})
	}

	return fetchedDynamicTables, nil
}

func (r *DynTabRepo) IDQuery(idQuery string) ([]string, error) {
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

func (r *DynTabRepo) AddDynTabVars(newVars []entity.NewDynamicTableVariable) ([]entity.DynamicTableVariable, error) {
	newVarIDs := []string{}
	for _, v := range newVars {
		id, err := r.q.CreateDynamicTableVariable(r.ctx, sqlcrepo.CreateDynamicTableVariableParams{
			ID:             v.ID,
			Name:           v.Name,
			VariableType:   v.VariableType,
			DynamicTableID: v.DynamicTableID,
			CreatedAt: pgtype.Timestamp{
				Time:  time.Now(),
				Valid: true,
			},
			CreatedBy: r.username,
			ModifiedAt: pgtype.Timestamp{
				Time:  time.Now(),
				Valid: true,
			},
			ModifiedBy: r.username,
		})
		if err != nil {
			return nil, err
		}
		newVarIDs = append(newVarIDs, id)
	}
	return r.FetchVarsByIDs(newVarIDs)
}

func (r *DynTabRepo) RemoveDynTabVars(varIDs []string) error {
	return r.q.DeleteDynamicTableVariables(r.ctx, varIDs)
}

func (r *DynTabRepo) FetchVarsByIDs(varsIDs []string) ([]entity.DynamicTableVariable, error) {
	fetchedDBVars, err := r.q.FetchDynamicTableVariables(r.ctx, varsIDs)
	if err != nil {
		return nil, err
	}

	fetchedDynamicTableVars := []entity.DynamicTableVariable{}

	for _, dbDynTable := range fetchedDBVars {
		fetchedDynamicTableVars = append(fetchedDynamicTableVars, entity.DynamicTableVariable{
			ID:           dbDynTable.ID,
			Name:         dbDynTable.Name,
			VariableType: dbDynTable.VariableType,
			DynamicTable: entity.Identifier{ID: dbDynTable.DynamicTableID, Name: dbDynTable.DynamicTableName},
			Meta: entity.MetaData{
				CreatedAt:  dbDynTable.CreatedAt.Time,
				CreatedBy:  dbDynTable.CreatedBy,
				ModifiedAt: dbDynTable.ModifiedAt.Time,
				ModifiedBy: dbDynTable.ModifiedBy,
			},
		})
	}

	return fetchedDynamicTableVars, nil
}

func (r *DynTabRepo) FetchVarsByTab(dynTabID string) ([]entity.DynamicTableVariable, error) {
	fetchedDBVars, err := r.q.FetchDynamicTableVariablesByDynamicTable(r.ctx, dynTabID)
	if err != nil {
		return nil, err
	}

	fetchedDynamicTableVars := []entity.DynamicTableVariable{}

	for _, dbDynTable := range fetchedDBVars {
		fetchedDynamicTableVars = append(fetchedDynamicTableVars, entity.DynamicTableVariable{
			ID:   dbDynTable.ID,
			Name: dbDynTable.Name,
			Meta: entity.MetaData{
				CreatedAt:  dbDynTable.CreatedAt.Time,
				CreatedBy:  dbDynTable.CreatedBy,
				ModifiedAt: dbDynTable.ModifiedAt.Time,
				ModifiedBy: dbDynTable.ModifiedBy,
			},
		})
	}

	return fetchedDynamicTableVars, nil

}

func (r *DynTabRepo) CreateTable(suffix string, dynTab entity.DynamicTable, vars []entity.DynamicTableVariable) error {
	variables, err := variablesToSQL(vars)
	if err != nil {
		return err
	}
	q := fmt.Sprintf(`CREATE TABLE "%s_%s" ( %s );`, suffix, dynTab.Name, variables)
	_, err = r.db.Query(context.Background(), q)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	return nil
}

// func updateTable() error {

// }

// func deleteTable() error {
// format!("DROP TABLE {};", dyntable_name(name))
// }

func (r *DynTabRepo) AddSchemaField(suffix string, vars []entity.DynamicTableVariable) error {
	tables := map[string][]entity.DynamicTableVariable{}
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
		_, err := r.db.Query(context.Background(), q)
		if err != nil {
			return fmt.Errorf("failed to execute query: %w", err)
		}
	}

	return nil
}

// func removeSchemaField() error {
//ALTER TABLE table_name DROP COLUMN column_name;
// }

// func updateSchemaField() error {
// ALTER TABLE table_name
// RENAME COLUMN old_name to new_name;
// }

func variablesToSQL(vars []entity.DynamicTableVariable) (string, error) {
	varQ := []string{}
	for _, v := range vars {
		varType, err := to_sql_type(v.VariableType)
		if err != nil {
			return "", nil
		}
		varQ = append(varQ, v.Name+" "+varType)
	}
	return strings.Join(varQ, ","), nil
}

func to_sql_type(s string) (string, error) {
	switch s {
	case "integer":
		return "INT", nil
	case "text":
		return "TEXT", nil
	default:
		return "", fmt.Errorf("variable_type '%s' not valid", s)
	}
}
