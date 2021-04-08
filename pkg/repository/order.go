package repository

import (
	"database/sql"
	"fmt"
	"github.com/HashimovH/esi-homework-3/pkg/domain"
)

type orderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *orderRepository {
	return &orderRepository{
		db: db,
	}
}

func (r *orderRepository) GetStatus(ident string, start string, end string) (int, error) {
	query := "SELECT COUNT(*) FROM porder WHERE ident_order=$1 AND ((start_order>$2 AND start_order<$3) OR (end_order>$2 AND end_order<$3))"
	row := r.db.QueryRow(query, ident, start, end)
	
	num := 0
	
	err := row.Scan(&num)
	fmt.Println(num)
	if err != nil {
		return 0, fmt.Errorf("could not close rows, err %v", err)
	}

	return num, nil
}

func (r *orderRepository) CreateOrder(*domain.Order) (*domain.Order, error){
	// Create Order here
	// return order, nil 
}

func (r *orderRepository) ListOrder() ([]*domain.Order, error){
	// If http parameter is there, we should filter for given date orders
	// Query: Select all
	// return orders, nil
	return ""
}

func (r *orderRepository) CancelOrder(id int) (string, error){
	// Query to mark order cancelled
	// return response
}