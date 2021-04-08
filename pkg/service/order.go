package service

// import (
// 	// "encoding/json"
// 	// "fmt"
// 	// "github.com/HashimovH/homework3/pkg/domain"
// 	// "io/ioutil"
// 	// "log"
// 	// "net/http"
// 	// "strconv"
// )

type orderRepository interface {
	GetStatus(ident string, start string, end string) (int, error)
	CreateOrder(*domain.Order) (*domain.Order, error)
	ListOrder() ([]*domain.Order, error)
	CancelOrder(id int) (*domain.Order, error)
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
	return s.orderRepository.Create(order)
}

func (s *OrderService) ListOrder() ([]*domain.Order, error){
	orders, err := s.orderRepository.ListOrder()
	if err != nil{
		return nil, err
	}
	return orders, err
}

func (s *OrderService) CancelOrder(id int) (string, error){
	return s.orderRepository.CancelOrder(id)
}


func (h *OrderService) GetStatus(i string, s string, e string) (int, error) {
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