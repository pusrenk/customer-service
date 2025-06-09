package main

import (
	"fmt"
	"net"

	"github.com/pusrenk/customer-service/configs"
	"github.com/pusrenk/customer-service/database"
	"github.com/pusrenk/customer-service/internal/customers/handlers/rpc"
	pb "github.com/pusrenk/customer-service/internal/protobuf"
	"github.com/pusrenk/customer-service/internal/customers/repositories"
	"github.com/pusrenk/customer-service/internal/customers/services"
	"github.com/pusrenk/customer-service/log"
	"google.golang.org/grpc"
)

func main() {
	// Init config
	configs.InitConfig()
	cfg := configs.GetConfig()

	// Init database
	db, err := database.InitDatabase(cfg)
	if err != nil {
		log.Panicf("Failed to init database: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Panicf("Failed to get SQL DB: %v", err)
	}
	defer func() {
		err := sqlDB.Close()
		if err != nil {
			log.Errorf("Error closing connection: %v", err)
		}
	}()

	// inti grpc server
	grpcServer := grpc.NewServer()

	// init customers
	customerRepository := repositories.NewCustomerRepository(db)
	customerService := services.NewCustomerService(customerRepository)
	customerServer := rpc.NewCustomerServiceServer(customerService)

	// register customers
	pb.RegisterCustomerServiceServer(grpcServer, customerServer)

	// start server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.App.GRPCPort))
	log.Infof("server is running on port %d", cfg.App.GRPCPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
