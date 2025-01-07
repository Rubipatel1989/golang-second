package engine

import (
	"context"
	"database/sql"
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
	return models.Engine{}, nil
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
