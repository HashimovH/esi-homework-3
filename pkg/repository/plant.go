package repository

import (
	"github.com/HashimovH/esi-homework-3/pkg/domain"
	// "context"
	// "fmt"
	// "database/sql"
	"gorm.io/gorm"
)

type PlantRepository struct {
	db *gorm.DB
}

func NewPlantRepository(db *gorm.DB) *PlantRepository {
	return &PlantRepository{
		db: db,
	}
}

func (r *PlantRepository) Create(plant *domain.Plant) (*domain.Plant, error) {
	// query := "INSERT INTO plant (ident, status, name, price) values($1, $2, $3, $4) RETURNING id, created_at"

	// row := r.db.QueryRowContext(context.Background(), query, plant.Ident, plant.Status, plant.Name, plant.Price)
	// if row == nil {
	// 	return nil, fmt.Errorf("Error while inserting plant %v", plant)
	// }

	// err := row.Scan(&plant.ID, &plant.CreatedAt)
	// if err != nil {
	// 	return nil, fmt.Errorf("Error while inserting plant %v", plant)
	// }

	r.db.AutoMigrate(&domain.Plant{})
	r.db.Create(plant)
	return plant, nil
}


func (r *PlantRepository) GetAll() ([]*domain.Plant, error) {
	// query := "SELECT id, ident, status, name, price, created_at from plant"
	// rows, err := r.db.QueryContext(context.Background(), query)
	// if err != nil {
	// 	return nil, fmt.Errorf("error querying plants, err: %v", err)
	// }

	// plants := []*domain.Plant{}
	// for rows.Next() {
	// 	b := &domain.Plant{}
	// 	err := rows.Scan(&b.ID, &b.Ident, &b.Status, &b.Name, &b.Price, &b.CreatedAt)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("error scaning query, err: %v", err)
	// 	}
	// 	plants = append(plants, b)
	// }
	// // close rows to avoid memory leak
	// err = rows.Close()
	// if err != nil {
	// 	return nil, fmt.Errorf("could not close rows, err %v", err)
	// }
	r.db.AutoMigrate(&domain.Plant{})
	var plants = []*domain.Plant{}
	r.db.Find(&plants)
	return plants, nil
}

func (r *PlantRepository) GetOne(ident string) (float64, error) {
	// query := "SELECT price from plant WHERE ident=$1"
	// row := r.db.QueryRow(query, ident)

	// var price string
	// err := row.Scan(&price)
	// if err != nil {
	// 	return "d", fmt.Errorf("could not close rows, err %v", err)
	// }
	plant := domain.Plant{}
	r.db.First(&plant, "ID = ?", ident)

	return plant.Price, nil
}