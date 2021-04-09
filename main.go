package main

import (
	// "context"
	// "database/sql"
	"fmt"
	// "github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"github.com/HashimovH/esi-homework-3/pkg/repository"
	"github.com/HashimovH/esi-homework-3/pkg/service"
	http2 "github.com/HashimovH/esi-homework-3/pkg/transport/http"
	// "github.com/graphql-go/graphql"
	// "github.com/HashimovH/homework3/pkg/handler/graphql/schema"
	// SQL driver
	// https://www.calhoun.io/why-we-import-sql-drivers-with-the-blank-identifier/
	// The sql package must be used in conjunction with a database driver. In this case PostgreSQL driver.
	// See https://golang.org/s/sqldrivers for a list of drivers.
	_ "github.com/lib/pq"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	// "encoding/json"
)

const (
	httpServicePort    = 8080
	postgresConnection = "dbname=postgres host=postgres password=postgres user=postgres sslmode=disable port=5432"
)



func main() {
	log.Println("Start plant server")

	router := Router()
	// setup http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", httpServicePort),
		Handler: router,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Could not start server")
	}

	log.Println("Stop plant server")
}

func Router() *mux.Router {
	dbConn, err := gorm.Open(postgres.Open(postgresConnection), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	plantRepository := repository.NewPlantRepository(dbConn)
	plantService := service.NewPlantService(plantRepository)
	plantHTTPHandler := http2.NewPlantHandler(plantService)

	orderRepository := repository.NewOrderRepository(dbConn)
	orderService := service.NewOrderService(orderRepository)
	orderHTTPHandler := http2.OrderStatusHandler(orderService)


	router := mux.NewRouter()
	plantHTTPHandler.RegisterRoutes(router)
	orderHTTPHandler.RegisterRoutes(router)

	return router
}