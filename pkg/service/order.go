package service

import (
	"encoding/json"
	// "fmt"
	"github.com/HashimovH/homework3/pkg/domain"
	// "io/ioutil"
	// "log"
	"net/http"
	// "strconv"
)

type OrderRepository interface {
	GetStatus(s *domain.Order) (*domain.Order, error)
}

type OrderService struct{
	OrderRepository OrderRepository
}

func NewOrderService(pR OrderRepository) *OrderService{
	return &OrderService{
		OrderRepository: pR,
	}
}


func (h *OrderService) GetStatus(w http.ResponseWriter, r *http.Request) {
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

	num, err := h.OrderRepository.GetNumIntersects(i, s, e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

	if num == "0" {
		w.Write([]byte("available"))
		response := &domain.Check{Availability: "available"}
		err = json.NewEncoder(w).Encode(&response)
	} else {
		w.Write([]byte("Not available"))
		response := &domain.Check{Availability: "Not available"}
		err = json.NewEncoder(w).Encode(&response)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}