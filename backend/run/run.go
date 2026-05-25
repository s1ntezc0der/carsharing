package run

import (
	"log"
	"net/http"

	"github.com/s1ntezc0der/car-rental/backend/internal/config"
	"github.com/s1ntezc0der/car-rental/backend/internal/controller"
	"github.com/s1ntezc0der/car-rental/backend/internal/repository"
	"github.com/s1ntezc0der/car-rental/backend/internal/usecase"
)

func Run() error {
	// Load config
	dbConfig := config.LoadDBConfig()
	serverConfig := config.LoadServerConfig()

	// Initialize DB
	db, err := repository.InitDB(dbConfig)
	if err != nil {
		return err
	}
	defer db.Close()

	// Initialize layers
	orderRepo := repository.NewOrderRepository(db)

	carService := usecase.NewCarService()
	orderService := usecase.NewOrderService(orderRepo)

	carHandler := controller.NewCarHandler(carService)
	orderHandler := controller.NewOrderHandler(orderService)

	// Create router
	router := controller.NewRouter(carHandler, orderHandler)

	// Start server
	log.Printf("Server started on :%s", serverConfig.Port)
	return http.ListenAndServe(":"+serverConfig.Port, router)
}
