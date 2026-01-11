package api

import (
	"context"

	"github.com/vishgupt/rainier/src/internal/common"
	"github.com/vishgupt/rainier/src/internal/core/database"
	pb "github.com/vishgupt/rainier/src/internal/pb"
)

// DatabaseServer implements the DatabaseService gRPC service
type DatabaseServer struct {
	pb.UnimplementedDatabaseServiceServer
	manager database.Manager
}

// NewDatabaseServer creates a new database server
func NewDatabaseServer(manager database.Manager) *DatabaseServer {
	return &DatabaseServer{
		manager: manager,
	}
}

// ListDatabases lists all databases
func (s *DatabaseServer) ListDatabases(ctx context.Context, req *pb.ListDatabasesRequest) (*pb.ListDatabasesResponse, error) {
	logger := common.GetLogger()
	logger.Infow("ListDatabases request", "page", req.Page, "limit", req.Limit)

	page := req.Page
	if page < 1 {
		page = 1
	}
	limit := req.Limit
	if limit <= 0 {
		limit = 10
	}

	databases, total, err := s.manager.ListDatabases(int(page), int(limit))
	if err != nil {
		logger.Errorw("Failed to list databases", "error", err)
		return nil, convertError(err)
	}

	pbDatabases := make([]*pb.Database, len(databases))
	for i, db := range databases {
		pbDatabases[i] = &pb.Database{
			Name:      db.Name,
			CreatedAt: db.CreatedAt.Unix(),
		}
	}

	return &pb.ListDatabasesResponse{
		Databases: pbDatabases,
		Total:     int32(total),
	}, nil
}

// GetDatabase retrieves a specific database
func (s *DatabaseServer) GetDatabase(ctx context.Context, req *pb.GetDatabaseRequest) (*pb.GetDatabaseResponse, error) {
	logger := common.GetLogger()
	logger.Infow("GetDatabase request", "database_name", req.DatabaseName)

	if req.DatabaseName == "" {
		return nil, convertError(common.NewValidationError("database_name is required"))
	}

	db, err := s.manager.GetDatabase(req.DatabaseName)
	if err != nil {
		logger.Errorw("Failed to get database", "error", err)
		return nil, convertError(err)
	}

	return &pb.GetDatabaseResponse{
		Database: &pb.Database{
			Name:      db.Name,
			CreatedAt: db.CreatedAt.Unix(),
		},
	}, nil
}
