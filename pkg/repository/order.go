package repository

import (
	// "database/sql"
	// "fmt"
	"github.com/HashimovH/esi-homework-3/pkg/domain"
	"gorm.io/gorm"
	"time"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{
		db: db,
	}
}

func (r *orderRepository) GetStatus(ident string, start string, end string) (int64, error) {
	// query := "SELECT COUNT(*) FROM porder WHERE ident_order=$1 AND ((start_order>$2 AND start_order<$3) OR (end_order>$2 AND end_order<$3))"
	// row := r.db.QueryRow(query, ident, start, end)
	
	// num := 0
	
	// err := row.Scan(&num)
	// fmt.Println(num)
	// if err != nil {
	// 	return 0, fmt.Errorf("could not close rows, err %v", err)
	// }

	var order []domain.Order
    var num int64
    r.db.Model(&order).Where("ID=?", ident).Where(r.db.Where("start_order>(?)", start).Where("start_order<(?)", end).Or(r.db.Where("end_order>(?)", start).Where("end_order<(?)", end))).Count(&num)
	return num, nil
}

func (r *orderRepository) CreateOrder(order *domain.Order) (*domain.Order, error){
	r.db.AutoMigrate(&domain.Order{})
	r.db.Create(order)
	return order, nil
}

func (r *orderRepository) ListOrder(due time.Time) ([]*domain.Order, error){
	// If http parameter is there, we should filter for given date orders
	// Query: Select all
	// return orders, nil
	r.db.AutoMigrate(&domain.Order{})
	var orders = []*domain.Order{}
	r.db.Where("start_order = ?", due).Find(&orders)
	return orders, nil
}

func (r *orderRepository) CancelOrder(id int) (string, error){
	// Query to mark order cancelled
	// return response
	r.db.Model(&domain.Order{}).Where("ID = ?", id).Update("status", "cancelled")
	return "Order has been cancelled", nil
}


func (r *orderRepository) UpdateOrder(id int, start time.Time, end time.Time) (*domain.Order, error){
	order := &domain.Order{}
	r.db.First(&order, id)
	order.Start = start
	order.End = end
	err := r.db.Save(&order).Error; 
	if err != nil {
		return order, err
	}
	return order, nil
}