package http

import (
	"encoding/json"
	"github.com/HashimovH/homework3/pkg/domain"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type NewOrderService interface {
	GetAll() (*domain.Order, error)
}

type orderHandler struct {
	NewOrderService NewOrderService
}

func OrderStatusHandler(pS NewOrderService) *orderHandler {
	return &orderHandler{
		OrderService: pS,
	}
}

func (h *orderHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/status", h.GetStatus).Methods(http.MethodPost)
}

func (h *orderHandler) GetStatus(w http.ResponseWriter, _ *http.Request) {
	status, err := h.OrderService.GetStatus()
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// write success response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&status)
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}


