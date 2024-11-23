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

// func (r *DynTabRepo) AddDynSchema(newVars []entity.NewDynamicField) ([]entity.DynamicField, error) {
// 	newVarIDs := []string{}

// 	for _, v := range newVars {
// 		id, err := r.q.CreateDynamicSchema(r.ctx, sqlcrepo.CreateDynamicSchemaParams{
// 			ID:             v.ID,
// 			Name:           v.Name,
// 			VariableType:   v.VariableType,
// 			DynamicTableID: v.DynamicTableID,
// 			CreatedAt: pgtype.Timestamp{
// 				Time:  time.Now(),
// 				Valid: true,
// 			},
// 			CreatedBy: r.username,
// 			ModifiedAt: pgtype.Timestamp{
// 				Time:  time.Now(),
// 				Valid: true,
// 			},
// 			ModifiedBy: r.username,
// 		})
// 		if err != nil {
// 			return nil, err
// 		}

// 		newVarIDs = append(newVarIDs, id)
// 	}

// 	return r.FetchVarsByIDs(newVarIDs)
// }
