package sights

import (
	"database/sql"
	"errors"
)

type UserParks struct {
	Id int  `json:"id" db:"id"`
	UserId string  `json:"user_id" db:"user_id"`
	ParkId string  `json:"park_id" db:"park_id"`
}

type ParkForDB struct {
	Id int `json:"id" db:"id"`
	Title string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Coords []float32 `json:"coords" db:"coords"`
	Address sql.NullString `json:"address" db:"address"`
}

type Park struct {
	Id int `json:"id" db:"id"`
	Title string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Coords []float32 `json:"coords" db:"coords"`
	Address sql.NullString `json:"address" db:"address"`
}

type GetParkOutput struct {
	Id int `json:"id" db:"id"`
	Title *string `json:"title"`
	Description *string `json:"description"`
	Coords []float32 `json:"coords"`
	Address *string `json:"address"`

	
}
// * - если не строка то вернет null 
type UpdateParkInput struct {
	Title *string `json:"title"`
	Description *string `json:"description"`
	Coords *[]float32 `json:"coords"`
	Address *string `json:"address"`
}

func (i UpdateParkInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Address == nil && i.Coords == nil {
		return errors.New("Update structure has no values")
	}

	return nil
}
