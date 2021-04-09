package service

import (
	// "encoding/json"
	// "fmt"
	"github.com/HashimovH/esi-homework-3/pkg/domain"
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
	GetOne(ident string) (float64, error)
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

func (s *PlantService) GetOne(ident string) (float64, error){
	plant, err := s.plantRepository.GetOne(ident)
	if err != nil{
		return 0, err
	}
	return plant, err
}



// func (h *PlantHandler) GetPrice(w http.ResponseWriter, r *http.Request) {
// 	log.Printf("received request %v", r)
// 	req_cache := r
// 	var data map[string]interface{}
// 	err := json.NewDecoder(r.Body).Decode(&data)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	i := data["ident"].(string)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	start, _ := strconv.Atoi(data["start"].(string))
// 	end, _ := strconv.Atoi(data["end"].(string))
// 	price, err := h.plantRepository.GetOne(i)
// 	p, _ := strconv.Atoi(price)
// 	if err != nil {
// 		price, err := h.plantmRepository.GetOne(i)
// 		p, _ := strconv.Atoi(price)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		res := &domain.Cost{Ident: i, Price: strconv.Itoa(p * (end - start))} //// calculating perdiod mult by price
// 		w.WriteHeader(http.StatusOK)
// 		err = json.NewEncoder(w).Encode(&res)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}
// 		return
// 	}
// 	s := strconv.Itoa(p * (end - start))
// 	res := &domain.Cost{Ident: i, Price: s}
// 	w.WriteHeader(http.StatusOK)
// 	err = json.NewEncoder(w).Encode(&res)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// 	_, err = h.CacheRequest(req_cache)
// 	//w.Write([]byte(answer))
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// }