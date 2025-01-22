package repository

import (
	sights "github.com/Safwood/go-server"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user sights.User) (int, error)
	GetUser(username string, password string) (sights.User, error)
}

type Park interface {
	CreatePark(userId int, park sights.Park) (int, error)
	GetAllParks(userId int) ([]sights.Park, error)
	GetParkById(userId, parkId int) (sights.GetParkOutput, error)
	DeletePark(userId, parkId int) (error)
	UpdatePark(userId, parkId int, input sights.UpdateParkInput) error
}

type Repository struct {
	Authorization
	Park
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db), 
		Park: NewParkPostgres(db),
	}
}