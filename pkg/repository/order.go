package repository

import (
	"database/sql"
	"fmt"
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