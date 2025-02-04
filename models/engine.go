package models

import (
	"errors"

	"github.com/google/uuid"
)

type Engine struct {
	EngineID      uuid.UUID `json:"engine_id"`
	Displacement  int64     `json:"displacement"`
	NoOfCylinders int64     `json:"noOfCylinders"`
	CarRange      int64     `json:"carRange"`
}
type EngineRequest struct {
	Displacement  int64 `json:"displacement"`
	NoOfCylinders int64 `json:"noOfCylinders"`
	CarRange      int64 `json:"carRange"`
}

func ValidateEngineRequest(EngineReq EngineRequest) error {
	if err := ValidateDisplacement(EngineReq.Displacement); err != nil {
		return err
	}
	if err := ValidateNoOfCylinders(EngineReq.NoOfCylinders); err != nil {
		return err
	}
	if err := ValidateCarRange(EngineReq.CarRange); err != nil {
		return err
	}
	return nil
}

func ValidateDisplacement(displacement int64) error {
	if displacement <= 0 {
		return errors.New("displacement must be greater than 0")
	}
	return nil
}

func ValidateNoOfCylinders(noOfCylinders int64) error {
	if noOfCylinders <= 0 {
		return errors.New("no of cylinders must be greater than 0")
	}
	return nil
}
func ValidateCarRange(carRange int64) error {
	if carRange <= 0 {
		return errors.New("car range must be greater than 0")
	}
	return nil
}
