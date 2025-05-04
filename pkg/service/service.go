package service

import (
	sights "github.com/safwood/go-server"
	"github.com/safwood/go-server/pkg/repository"
)

type Authorization interface {
	CreateUser(user sights.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Park interface {
	CreatePark(userId int, park sights.Park) (int, error)
	GetAllParks(userId int) ([]sights.Park, error)
	GetParkById(userId, parkId int) (sights.GetParkOutput, error)
	DeletePark(userId, parkId int) (error)
	UpdatePark(userId, parkId int, input sights.UpdateParkInput) error
}

type Service struct {
	Authorization
	Park
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: newAuthService(repos.Authorization),
		Park: newParkService(repos.Park),
	}
}