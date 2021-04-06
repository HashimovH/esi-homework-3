package http

import (
	"encoding/json"
	"github.com/HashimovH/homework3/pkg/domain"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type OrderService interface {
	GetStatus(i string, s string, e string) (int, error)
}

type orderHandler struct {
	OrderService OrderService
}

func OrderStatusHandler(pS OrderService) *orderHandler {
	return &orderHandler{
		OrderService: pS,
	}
}

func (h *orderHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/status", h.GetStatus).Methods(http.MethodPost)
}

func (h *orderHandler) GetStatus(w http.ResponseWriter, r *http.Request) {
	
	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	i := data["ident"].(string)
	s := data["start"].(string)
	e := data["end"].(string)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	status, err := h.OrderService.GetStatus(i, s, e)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if status == 0{
		response := &domain.Check{Availability: "not available"}
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(&response)
		if err != nil {
			log.Errorf("Could not encode json, err %v", err)
		}
	}else{
		response := &domain.Check{Availability: "available"}
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(&response)
		if err != nil {
			log.Errorf("Could not encode json, err %v", err)
		}
	}
	
}


