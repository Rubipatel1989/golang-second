package car

import (
	"context"
	"database/sql"
	"errors"
	"golangSecond/models"
	"time"

	"github.com/google/uuid"
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
	var car models.Car
	query := `SELECT c.id, c.name, c.year, c.brand, c.fuel_type, c.price, c.engine_id, c.created_at, c.updated_at, e.id, e.displacement, e.no_of_cylinders, e.car_range 
	from cars c 
	left join engines e on c.engine_id = e.id 
	where c.id = $1`

	row := s.db.QueryRowContext(ctx, query, id)
	if err := row.Scan(
		&car.ID,
		&car.Name,
		&car.Year,
		&car.Brand,
		&car.FuelType,
		&car.Price,
		&car.Engine.EngineID,
		&car.CreatedAt,
		&car.UpdatedAt,
		&car.Engine.EngineID,
		&car.Engine.Displacement,
		&car.Engine.NoOfCylinders,
		&car.Engine.CarRange,
	); err != nil {
		if err == sql.ErrNoRows {
			return car, nil
		}
		return car, err
	}

	return car, nil
}
func (s Store) GetCarByBrand(ctx context.Context, brand string, isEngine bool) ([]models.Car, error) {
	var cars []models.Car
	var query string
	if isEngine {
		query = `SELECT c.id, c.name, c.year, c.brand, c.fuel_type, c.price, c.engine_id, c.created_at, c.updated_at, e.id, e.displacement, e.no_of_cylinders, e.car_range 
		from car c 
		left join engines e on c.engine_id = e.id 
		where c.brand = $1`
	} else {
		query = `SELECT id, name, year, brand, fuel_type, price, engine_id, created_at, updated_at
		from car
		where brand = $1`
	}
	rows, err := s.db.QueryContext(ctx, query, brand)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var car models.Car
		if isEngine {
			var engine models.Engine
			err := rows.Scan(
				&car.ID,
				&car.Name,
				&car.Year,
				&car.Brand,
				&car.FuelType,
				&car.Price,
				&car.Engine.EngineID,
				&car.CreatedAt,
				&car.UpdatedAt,
				&engine.EngineID,
				&engine.Displacement,
				&engine.NoOfCylinders,
				&engine.CarRange)
			if err != nil {
				return nil, err
			}
			car.Engine = engine
		} else {
			err := rows.Scan(
				&car.ID,
				&car.Name,
				&car.Year,
				&car.Brand,
				&car.FuelType,
				&car.Price,
				&car.Engine.EngineID,
				&car.CreatedAt,
				&car.UpdatedAt)
			if err != nil {
				return nil, err
			}
		}

		cars = append(cars, car)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return cars, nil

}

func (s *Store) CreateCar(ctx context.Context, carReq *models.CarRequest) (models.Car, error) {
	var createdCar models.Car
	var engineID uuid.UUID

	err := s.db.QueryRowContext(ctx, "SELECT id FROM engines WHERE id = $1", carReq.Engine.EngineID).Scan(&engineID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return createdCar, errors.New("engine id does not found in the engines table")
		}
		return createdCar, err
	}
	carId := uuid.New()
	createdAt := time.Now()
	updatedAt := createdAt
	newCar := models.Car{
		ID:        carId,
		Name:      carReq.Name,
		Year:      carReq.Year,
		Brand:     carReq.Brand,
		FuelType:  carReq.FuelType,
		Engine:    carReq.Engine,
		Price:     carReq.Price,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
	// Begain the Transaction
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return createdCar, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	// Insert the car into the cars table

	query := `
	INSERT INTO cars (id, name, year, brand, fuel_type, engine_id, price, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING id, name, year, brand, fuel_type, engine_id, price, created_at, updated_at;
	`
	err = tx.QueryRowContext(ctx, query,
		newCar.ID,
		newCar.Name,
		newCar.Year,
		newCar.Brand,
		newCar.FuelType,
		newCar.Engine.EngineID,
		newCar.Price,
		newCar.CreatedAt,
		newCar.UpdatedAt).Scan(
		&createdCar.ID,
		&createdCar.Name,
		&createdCar.Year,
		&createdCar.Brand,
		&createdCar.FuelType,
		&createdCar.Engine.EngineID,
		&createdCar.Price,
		&createdCar.CreatedAt,
		&createdCar.UpdatedAt,
	)
	if err != nil {
		return createdCar, err
	}
	

	return createdCar, nil

}
func (s *Store) UpdateCar(ctx context.Context, id string, carReq *models.CarRequest) (models.Car, error) {
	return models.Car{}, nil
}
func (s *Store) DeleteCar(ctx context.Context, id string) (models.Car, error) {
	return models.Car{}, nil
}
