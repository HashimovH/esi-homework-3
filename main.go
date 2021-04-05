package main

import (
	"context"
	"database/sql"
	"fmt"
	// "github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"github.com/HashimovH/homework3/pkg/repository"
	"github.com/HashimovH/homework3/pkg/service"
	http2 "github.com/HashimovH/homework3/pkg/transport/http"
	// "github.com/graphql-go/graphql"
	// "github.com/HashimovH/homework3/pkg/handler/graphql/schema"
	// SQL driver
	// https://www.calhoun.io/why-we-import-sql-drivers-with-the-blank-identifier/
	// The sql package must be used in conjunction with a database driver. In this case PostgreSQL driver.
	// See https://golang.org/s/sqldrivers for a list of drivers.
	// _ "github.com/lib/pq"
	// "encoding/json"
)

const (
	httpServicePort    = 8080
	postgresConnection = "dbname=postgres host=postgres password=postgres user=postgres sslmode=disable port=5432"
	mongoConnection    = "mongodb://mongo:27017"
	// redisURI           = "redis:6379"
	// redisPassword      = "" // no password set
	// redisDB            = 0  // use default DB
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
    dbConn, err := sql.Open("postgres", postgresConnection)
	if err != nil {
		log.Fatal(err)
	}

	//mongooooooooooooo
	dbConn2, err := mongo.NewClient(options.Client().ApplyURI(mongoConnection))
	if err != nil {
		log.Fatal(err)
	}
	err = dbConn2.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// dbConn3 := redis.NewClient(&redis.Options{
	// 	Addr:     redisURI,
	// 	Password: redisPassword,
	// 	DB:       redisDB,
	// })

	plantRepository := repository.NewPlantRepository(dbConn)
	plantService := service.NewPlantService(plantRepository)
	plantHTTPHandler := http2.NewPlantHandler(plantService)

	plantmRepository := repository.NewPlantmRepository(dbConn2)
	plantmService := service.NewPlantmService(plantmRepository)
	plantmHTTPHandler := http2.NewPlantmHandler(plantmService)

	orderRepository := repository.NewOrderRepository(dbConn)
	orderService := handler.OrderService(orderRepository)
	orderHTTPHandler := http2.orderHandler(orderService)


	router := mux.NewRouter()
	plantHTTPHandler.RegisterRoutes(router)
	plantmHTTPHandler.RegisterRoutes(router)
	orderHTTPHandler.RegisterRoutes(router)

	return router

	// orderRepository := repository.NewOrderRepository(dbConn)
	// cacheRepository := repository.NewCacheRepository(dbConn3)
	// orderHandler := handler.NewOrderHandler(orderRepository)
	// plantHandler := handler.NewPlantHandler(plantmRepository, plantRepository)
	// mongooooooooooooooo

	// construct application

	
	

	// router.HandleFunc("/plant", plantHandler.CreatePlant).Methods(http.MethodPost)
	// router.HandleFunc("/plantm", plantHandler.CreatePlantm).Methods(http.MethodPost)
	// router.HandleFunc("/plant", plantHandler.GetPlants).Methods(http.MethodGet)
	// router.HandleFunc("/price", plantHandler.GetPrice).Methods(http.MethodPost)
	// router.HandleFunc("/status", orderHandler.GetStatus).Methods(http.MethodPost)
	// router.HandleFunc("/requests", plantHandler.GetCache).Methods(http.MethodGet)
	// router.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
	// 	result := executeQuery(r.URL.Query().Get("query"), schema.TodoSchema)
	// 	json.NewEncoder(w).Encode(result)
	// })
}