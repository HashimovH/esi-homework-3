package http

import (
	"encoding/json"
	"github.com/HashimovH/esi-homework-3/pkg/domain"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type OrderService interface {
	GetStatus(i string, s string, e string) (int, error)
	CreateOrder(*domain.Order) (*domain.Order, error)
	ListOrder() ([]*domain.Order, error)
	CancelOrder(id int) (*domain.Order, error)
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
	router.HandleFunc("/order", h.CreateOrder).Methods(http.MethodPost) // Create Order
	router.HandleFunc("/order", h.ListOrder).Methods(http.MethodGet) // Create Order
	// router.HandleFunc("/order-delivery", h.OrderDelivery).Methods(http.MethodPost) // Find orders for given data
	router.HandleFunc("/cancel/:id", h.CancelOrder).Methods(http.MethodGet) // Cancel order if it is not delivered yet
}

func (h *orderHandler) CreateOrder(w http.ResponseWriter, r *http.Request){
	order := &domain.Order{}
	err := json.NewDecoder(r.Body).Decode(order)
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

	createdOrder, err := h.orderService.CreateOrder(order)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// write success response
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(&createdOrder)
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}

func (h *orderHandler) ListOrder(w http.ResponseWriter, r *http.Request){
	orders, err := h.orderService.ListOrder()
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// write success response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&orders)
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}

func (h *orderHandler) CancelOrder(w http.ResponseWriter, r *http.Request){
	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	i := data["id"].(int)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	order, err = h.OrderService.CancelOrder(i)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
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


