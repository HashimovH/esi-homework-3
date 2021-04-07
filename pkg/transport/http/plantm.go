package http

import (
	"encoding/json"
	"github.com/HashimovH/esi-homework-3/pkg/domain"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type plantmService interface {
	Create(plant *domain.Plantm) (*domain.Plantm, error)
	GetAll() ([]*domain.Plantm, error)
	// GetOnePlant() (*domain.Plantm, error)
}

type plantmHandler struct {
	plantmService plantmService
}

func NewPlantmHandler(pS plantmService) *plantmHandler {
	return &plantmHandler{
		plantmService: pS,
	}
}

func (h *plantmHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/plantm", h.Create).Methods(http.MethodPost)
	router.HandleFunc("/plantm", h.GetAll).Methods(http.MethodGet)
}

func (h *plantmHandler) Create(w http.ResponseWriter, r *http.Request) {

	plantm := &domain.Plantm{}
	err := json.NewDecoder(r.Body).Decode(plantm)
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

	createdPlant, err := h.plantmService.Create(plantm)
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

func (h *plantmHandler) GetAll(w http.ResponseWriter, _ *http.Request) {
	plants, err := h.plantmService.GetAll()
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