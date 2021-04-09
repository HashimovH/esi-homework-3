package service

import (
// 	// "encoding/json"
// 	// "fmt"
	"github.com/HashimovH/esi-homework-3/pkg/domain"
// 	// "io/ioutil"
// 	// "log"
// 	// "net/http"
// 	// "strconv"
	"time"
)

type orderRepository interface {
	GetStatus(ident string, start string, end string) (int64, error)
	CreateOrder(*domain.Order) (*domain.Order, error)
	UpdateOrder(id int, start time.Time, end time.Time) (*domain.Order, error)
	ListOrder(due time.Time) ([]*domain.Order, error)
	CancelOrder(id int) (string, error)
}

type OrderService struct{
	orderRepository orderRepository
}

func NewOrderService(pR orderRepository) *OrderService{
	return &OrderService{
		orderRepository: pR,
	}
}

func (s *OrderService) CreateOrder(order *domain.Order) (*domain.Order, error){
	return s.orderRepository.CreateOrder(order)
}

func (s *OrderService) UpdateOrder(id int, start time.Time, end time.Time) (*domain.Order, error){
	order, err := s.orderRepository.UpdateOrder(id, start, end)
	if err != nil{
		return nil, err
	}
	return order, err
}

func (s *OrderService) ListOrder(due time.Time) ([]*domain.Order, error){
	orders, err := s.orderRepository.ListOrder(due)
	if err != nil{
		return nil, err
	}
	return orders, err
}

func (s *OrderService) CancelOrder(id int) (string, error){
	return s.orderRepository.CancelOrder(id)
}


func (h *OrderService) GetStatus(i string, s string, e string) (int64, error) {
	// var data map[string]interface{}
	// err := json.NewDecoder(r.Body).Decode(&data)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// i := data["ident"].(string)
	// s := data["start"].(string)
	// e := data["end"].(string)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	num, err := h.orderRepository.GetStatus(i, s, e)
	if err != nil {
		return 0, err
	}
	// w.WriteHeader(http.StatusOK)
	return num, err

	// if num == 0 {
	// 	w.Write([]byte("available"))
	// 	response := &domain.Check{Availability: "available"}
	// 	return response
	// } else {
	// 	w.Write([]byte("Not available"))
	// 	response := &domain.Check{Availability: "Not available"}
	// 	return response
	// }


	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

}