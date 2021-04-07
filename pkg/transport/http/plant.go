package http

import (
	"encoding/json"
	"github.com/HashimovH/esi-homework-3/pkg/domain"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	//"strconv"
)

type plantService interface {
	Create(plant *domain.Plant) (*domain.Plant, error)
	GetAll() ([]*domain.Plant, error)
	GetOne(ident string) (string, error)
}

type plantHandler struct {
	plantService plantService
}

func NewPlantHandler(pS plantService) *plantHandler {
	return &plantHandler{
		plantService: pS,
	}
}

func (h *plantHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/plant", h.Create).Methods(http.MethodPost)
	router.HandleFunc("/plant", h.GetAll).Methods(http.MethodGet)
	
	//router.HandleFunc("/price", h.GetPrice).Methods(http.MethodPost)
}

func (h *plantHandler) Create(w http.ResponseWriter, r *http.Request) {

	plant := &domain.Plant{}
	err := json.NewDecoder(r.Body).Decode(plant)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// close body to avoid memory leak
	err = r.Body.Close()
	if err != nil {
		log.Errorf("Could not close request body, err %v", err)
	}

	createdPlant, err := h.plantService.Create(plant)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// write success response
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(&createdPlant)
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}

func (h *plantHandler) GetAll(w http.ResponseWriter, _ *http.Request) {
	plants, err := h.plantService.GetAll()
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// write success response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&plants)
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}

func (h *plantHandler) GetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	plants, err := h.plantService.GetOne(key)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// write success response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&plants)
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}


// func (h *orderHandler) GetStatus(w http.ResponseWriter, r *http.Request) {
	
// 	var data map[string]interface{}
// 	err := json.NewDecoder(r.Body).Decode(&data)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	i := data["ident"].(string)
// 	s := data["start"].(string)
// 	e := data["end"].(string)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	status, err := h.OrderService.GetStatus(i, s, e)
// 	if err != nil {
// 		log.Error(err.Error())
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	if status == 0{
// 		response := &domain.Check{Availability: "not available"}
// 		w.WriteHeader(http.StatusOK)
// 		err = json.NewEncoder(w).Encode(&response)
// 		if err != nil {
// 			log.Errorf("Could not encode json, err %v", err)
// 		}
// 	}else{
// 		response := &domain.Check{Availability: "available"}
// 		w.WriteHeader(http.StatusOK)
// 		err = json.NewEncoder(w).Encode(&response)
// 		if err != nil {
// 			log.Errorf("Could not encode json, err %v", err)
// 		}
// 	}
	
// }
