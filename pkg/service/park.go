package service

import (
	sights "github.com/safwood/go-server"
	"github.com/safwood/go-server/pkg/repository"
)

type ParkService struct {
	repo repository.Park
}

func newParkService(repo repository.Park) *ParkService {
	return &ParkService{repo}
}

func (s *ParkService) CreatePark(userId int, park sights.Park) (int, error) {
	return s.repo.CreatePark(userId, park)
}

func (s *ParkService) GetAllParks(parkId int) ([]sights.Park, error) {
	return s.repo.GetAllParks(parkId)
}

func (s *ParkService) GetParkById(userId int, parkId int) (sights.GetParkOutput, error) {
	return s.repo.GetParkById(userId, parkId)
}

func (s *ParkService) DeletePark(userId, parkId int) (error) {
	return s.repo.DeletePark(userId, parkId)
}

func (s *ParkService) UpdatePark(userId, parkId int, input sights.UpdateParkInput) (error) {
	return s.repo.UpdatePark(userId, parkId, input)
}