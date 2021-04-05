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

type plantRepository interface {
	// plantmRepository *repository.PlantmRepository
	// plantRepository  *repository.PlantRepository
	// cacheRepository  *repository.CacheRepository
	Create(s *domain.Plant) (*domain.Plant, error)
	GetAll() ([]*domain.Plant, error)
	// GetOne() (*domain.Plant, error)
}

type PlantService struct{
	plantRepository plantRepository
}

func NewPlantService(pR plantRepository) *PlantService{
	return &PlantService{
		plantRepository: pR,
	}
}

func (s *PlantService) Create(plant *domain.Plant) (*domain.Plant, error){
	return s.plantRepository.Create(plant)
}

func (s *PlantService) GetAll() ([]*domain.Plant, error){
	plants, err := s.plantRepository.GetAll()
	if err != nil{
		return nil, err
	}
	return plants, err
}

// func (s *PlantService) GetOne() (*domain.Plant, error){
// 	plant, err := s.plantRepository.GetOne()
// 	if err != nil{
// 		return nil, err
// 	}
// 	return plant, err
// }