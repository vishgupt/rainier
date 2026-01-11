package api

import (
	"context"

	"github.com/vishgupt/rainier/src/internal/common"
	"github.com/vishgupt/rainier/src/internal/core/collection"
	pb "github.com/vishgupt/rainier/src/internal/pb"
)

// CollectionServer implements the CollectionService gRPC service
type CollectionServer struct {
	pb.UnimplementedCollectionServiceServer
	manager collection.Manager
}

// NewCollectionServer creates a new collection server
func NewCollectionServer(manager collection.Manager) *CollectionServer {
	return &CollectionServer{
		manager: manager,
	}
}

// ListCollections lists all collections in a database
func (s *CollectionServer) ListCollections(ctx context.Context, req *pb.ListCollectionsRequest) (*pb.ListCollectionsResponse, error) {
	logger := common.GetLogger()
	logger.Infow("ListCollections request", "database_name", req.DatabaseName, "page", req.Page, "limit", req.Limit)

	if req.DatabaseName == "" {
		return nil, convertError(common.NewValidationError("database_name is required"))
	}

	page := req.Page
	if page < 1 {
		page = 1
	}
	limit := req.Limit
	if limit <= 0 {
		limit = 10
	}

	collections, total, err := s.manager.ListCollections(req.DatabaseName, int(page), int(limit))
	if err != nil {
		logger.Errorw("Failed to list collections", "error", err)
		return nil, convertError(err)
	}

	pbCollections := make([]*pb.Collection, len(collections))
	for i, col := range collections {
		pbCollections[i] = &pb.Collection{
			Name:            col.Name,
			DatabaseName:    col.DatabaseName,
			VectorDimension: col.Dimension,
			CreatedAt:       col.CreatedAt.Unix(),
		}
	}

	return &pb.ListCollectionsResponse{
		Collections: pbCollections,
		Total:       int32(total),
	}, nil
}

// GetCollection retrieves a specific collection
func (s *CollectionServer) GetCollection(ctx context.Context, req *pb.GetCollectionRequest) (*pb.GetCollectionResponse, error) {
	logger := common.GetLogger()
	logger.Infow("GetCollection request", "database_name", req.DatabaseName, "collection_name", req.CollectionName)

	if req.DatabaseName == "" {
		return nil, convertError(common.NewValidationError("database_name is required"))
	}
	if req.CollectionName == "" {
		return nil, convertError(common.NewValidationError("collection_name is required"))
	}

	col, err := s.manager.GetCollection(req.DatabaseName, req.CollectionName)
	if err != nil {
		logger.Errorw("Failed to get collection", "error", err)
		return nil, convertError(err)
	}

	return &pb.GetCollectionResponse{
		Collection: &pb.Collection{
			Name:            col.Name,
			DatabaseName:    col.DatabaseName,
			VectorDimension: col.Dimension,
			CreatedAt:       col.CreatedAt.Unix(),
		},
	}, nil
}
