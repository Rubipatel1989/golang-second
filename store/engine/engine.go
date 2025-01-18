package engine

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"golangSecond/models"
)

type EngineStore struct {
	db *sql.DB
}

func New(db *sql.DB) *EngineStore {
	return &EngineStore{
		db: db,
	}
}
func (e EngineStore) EngineById(ctx context.Context, id string) (models.Engine, error) {
	var engine models.Engine
	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return engine, err
	}
	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				fmt.Printf("error rolling back transaction: %v", rbErr)
			}
		} else {
			if cmErr := tx.Commit(); cmErr != nil {
				fmt.Printf("error committing transaction: %v", cmErr)
			}
		}
		err = tx.Commit()
	}()
	err = tx.QueryRowContext(ctx, `SELECT id, displacement, no_of_cylinders, car_range FROM engine WHERE id = $1`, id).Scan(
		&engine.EngineID,
		&engine.Displacement,
		&engine.NoOfCylinders,
		&engine.CarRange,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return engine, nil
		}
		return engine, err
	}
	return engine, err
}
func (e EngineStore) CreateEngine(ctx context.Context, engineReq *models.EngineRequest) (models.Engine, error) {
	return models.Engine{}, nil
}
func (e EngineStore) EngineUpdate(ctx context.Context, id string, engineReq *models.EngineRequest) (models.Engine, error) {
	return models.Engine{}, nil
}
func (e EngineStore) EngineDelete(ctx context.Context, id string) (models.Engine, error) {
	return models.Engine{}, nil
}
