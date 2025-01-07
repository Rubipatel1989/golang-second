package car

import (
	"context"
	"database/sql"
	"golangSecond/models"
)

type Store struct {
	db *sql.DB
}

func new(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}
func (s *Store) GetCarById(ctx context.Context, id string) (models.Car, error) {
	return models.Car{}, nil
}
func (s Store) GetCarByBrand(ctx context.Context, brand string, isEngine bool) {

}

func (s *Store) CreateCar(ctx context.Context, carReq *models.CarRequest) (models.Car, error) {
	return models.Car{}, nil
}
func (s *Store) UpdateCar(ctx context.Context, id string, carReq *models.CarRequest) (models.Car, error) {
	return models.Car{}, nil
}
func (s *Store) DeleteCar(ctx context.Context, id string) (models.Car, error) {
	return models.Car{}, nil
}
