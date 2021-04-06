package service

import (
	// "encoding/json"
	// "fmt"
	"github.com/HashimovH/homework3/pkg/domain"
	// "io/ioutil"
	// "log"
	// "net/http"
	// "strconv"
)

type plantmRepository interface {
	Create(s *domain.Plantm) (*domain.Plantm, error)
	GetAll() ([]*domain.Plantm, error)
	GetOne() (*domain.Plantm, error)
}

type PlantmService struct{
	plantmRepository plantmRepository
}

func NewPlantmService(pR plantmRepository) *PlantmService{
	return &PlantmService{
		plantmRepository: pR,
	}
}

func (s *PlantmService) Create(plantm *domain.Plantm) (*domain.Plantm, error){
	return s.plantmRepository.Create(plantm)
}

func (s *PlantmService) GetAll() ([]*domain.Plantm, error){
	plants, err := s.plantmRepository.GetAll()
	if err != nil{
		return nil, err
	}
	return plants, err
}

func (s *PlantmService) GetOne() (*domain.Plantm, error){
	plant, err := s.plantmRepository.GetOne()
	if err != nil{
		return nil, err
	}
	return plant, err
}