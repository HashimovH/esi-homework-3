package http

import (
	"encoding/json"
	"github.com/HashimovH/homework3/pkg/domain"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type plantService interface {
	Create(plant *domain.Plant) (*domain.Plant, error)
	GetAll() ([]*domain.Plant, error)
	// GetOne() (*domain.Plant, error)
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