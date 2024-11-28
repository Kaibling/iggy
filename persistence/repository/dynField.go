package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/persistence/sqlcrepo"
)

type DynFieldRepo struct {
	ctx      context.Context
	q        *sqlcrepo.Queries
	db       *pgxpool.Pool
	username string
}

func NewDynFieldRepo(ctx context.Context, username string, dbPool *pgxpool.Pool) *DynFieldRepo {
	return &DynFieldRepo{
		ctx:      ctx,
		q:        sqlcrepo.New(dbPool),
		db:       dbPool,
		username: username,
	}
}

func (r *DynFieldRepo) AddFields(newModels []entity.NewDynamicField) ([]entity.DynamicField, error) {
	newDynamicFieldIDs := []string{}

	for _, newModel := range newModels {
		newDynamicFieldID, err := r.q.CreateDynamicField(r.ctx, sqlcrepo.CreateDynamicFieldParams{
			ID:             newModel.ID,
			Name:           newModel.Name,
			DynamicTableID: newModel.DynamicTableID,
			VariableType:   newModel.VariableType,
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

		newDynamicFieldIDs = append(newDynamicFieldIDs, newDynamicFieldID)
	}

	return r.FetchDynamicFields(newDynamicFieldIDs)
}

func (r *DynFieldRepo) FetchDynamicFields(dynFieldIDs []string) ([]entity.DynamicField, error) {
	fetchedDBVars, err := r.q.FetchDynamicFields(r.ctx, dynFieldIDs)
	if err != nil {
		return nil, err
	}

	fetchedDynamicTableVars := []entity.DynamicField{}

	for _, dbDynTable := range fetchedDBVars {
		fetchedDynamicTableVars = append(fetchedDynamicTableVars, entity.DynamicField{
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

func (r *DynFieldRepo) IDQuery(idQuery string) ([]string, error) {
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

func (r *DynFieldRepo) RemoveDynField(varIDs []string) error {
	return r.q.DeleteDynamicFields(r.ctx, varIDs)
}

func (r *DynFieldRepo) FetchVarsByTabID(dynTabID string) ([]entity.DynamicField, error) {
	fetchedDBVars, err := r.q.FetchDynamicFieldsByDynamicTable(r.ctx, dynTabID)
	if err != nil {
		return nil, err
	}

	fetchedDynamicTableVars := []entity.DynamicField{}

	for _, dbDynTable := range fetchedDBVars {
		fetchedDynamicTableVars = append(fetchedDynamicTableVars, entity.DynamicField{
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

func (r *DynFieldRepo) FetchVarsByDynFieldName(dynFieldName string) ([]entity.DynamicField, error) {
	fetchedDBVars, err := r.q.FetchDynamicFieldsByDynamicTableName(r.ctx, dynFieldName)
	if err != nil {
		return nil, err
	}

	fetchedDynamicTableVars := []entity.DynamicField{}

	for _, dbDynTable := range fetchedDBVars {
		fetchedDynamicTableVars = append(fetchedDynamicTableVars, entity.DynamicField{
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
