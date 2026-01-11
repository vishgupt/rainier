package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	"github.com/vishgupt/rainier/src/internal/api"
	"github.com/vishgupt/rainier/src/internal/common"
	"github.com/vishgupt/rainier/src/internal/core/collection"
	"github.com/vishgupt/rainier/src/internal/core/database"
	"github.com/vishgupt/rainier/src/internal/core/point"
	pb "github.com/vishgupt/rainier/src/internal/pb"
)

func main() {
	// Initialize logging
	logger := common.InitLogger()
	defer logger.Sync()

	port := ":50051"
	logger.Infow("Starting Rainier Vector Database server",
		"version", "0.1.0",
		"port", port,
	)

	// Initialize managers
	dbManager := database.NewInMemoryManager()
	collManager := collection.NewInMemoryManager()
	pointManager := point.NewInMemoryManager()

	// Create gRPC server
	grpcServer := grpc.NewServer()

	// Register services
	pb.RegisterDatabaseServiceServer(grpcServer, api.NewDatabaseServer(dbManager))
	pb.RegisterCollectionServiceServer(grpcServer, api.NewCollectionServer(collManager))
	pb.RegisterPointServiceServer(grpcServer, api.NewPointServer(pointManager))

	// Listen on port
	listener, err := net.Listen("tcp", port)
	if err != nil {
		logger.Fatalw("Failed to listen", "error", err)
	}

	logger.Infow("gRPC server listening", "address", listener.Addr())

	// Handle graceful shutdown
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
		logger.Info("Shutting down server...")
		grpcServer.GracefulStop()
	}()

	// Start server
	if err := grpcServer.Serve(listener); err != nil {
		logger.Fatalw("Server error", "error", err)
	}

	logger.Info("Server shutdown successfully")
}
