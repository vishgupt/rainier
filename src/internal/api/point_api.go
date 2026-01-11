package api

import (
	"context"

	"github.com/vishgupt/rainier/src/internal/common"
	"github.com/vishgupt/rainier/src/internal/core/point"
	pb "github.com/vishgupt/rainier/src/internal/pb"
)

// PointServer implements the PointService gRPC service
type PointServer struct {
	pb.UnimplementedPointServiceServer
	manager point.Manager
}

// NewPointServer creates a new point server
func NewPointServer(manager point.Manager) *PointServer {
	return &PointServer{
		manager: manager,
	}
}

// GetPoint retrieves a specific point by ID
func (s *PointServer) GetPoint(ctx context.Context, req *pb.GetPointRequest) (*pb.GetPointResponse, error) {
	logger := common.GetLogger()
	logger.Infow("GetPoint request", "database_name", req.DatabaseName, "collection_name", req.CollectionName, "point_id", req.PointId)

	if req.DatabaseName == "" {
		return nil, convertError(common.NewValidationError("database_name is required"))
	}
	if req.CollectionName == "" {
		return nil, convertError(common.NewValidationError("collection_name is required"))
	}
	if req.PointId == "" {
		return nil, convertError(common.NewValidationError("point_id is required"))
	}

	points, err := s.manager.GetPoints(req.CollectionName, []string{req.PointId})
	if err != nil {
		logger.Errorw("Failed to get point", "error", err)
		return nil, convertError(err)
	}

	if len(points) == 0 {
		return nil, convertError(common.NewNotFoundError("point not found"))
	}

	p := points[0]
	return &pb.GetPointResponse{
		Point: &pb.Point{
			Id:             p.ID,
			CollectionName: req.CollectionName,
			DatabaseName:   req.DatabaseName,
			Vector:         p.Values,
			CreatedAt:      p.CreatedAt.Unix(),
		},
	}, nil
}

// SearchNearest searches for nearest neighbor points
func (s *PointServer) SearchNearest(ctx context.Context, req *pb.SearchNearestRequest) (*pb.SearchNearestResponse, error) {
	logger := common.GetLogger()
	logger.Infow("SearchNearest request", "database_name", req.DatabaseName, "collection_name", req.CollectionName, "limit", req.Limit)

	if req.DatabaseName == "" {
		return nil, convertError(common.NewValidationError("database_name is required"))
	}
	if req.CollectionName == "" {
		return nil, convertError(common.NewValidationError("collection_name is required"))
	}
	if len(req.QueryVector) == 0 {
		return nil, convertError(common.NewValidationError("query_vector is required"))
	}

	limit := req.Limit
	if limit <= 0 {
		limit = 10
	}

	points, err := s.manager.SearchPoints(req.CollectionName, req.QueryVector, int(limit))
	if err != nil {
		logger.Errorw("Failed to search points", "error", err)
		return nil, convertError(err)
	}

	results := make([]*pb.NearestPoint, len(points))
	for i, p := range points {
		results[i] = &pb.NearestPoint{
			Point: &pb.Point{
				Id:             p.ID,
				CollectionName: req.CollectionName,
				DatabaseName:   req.DatabaseName,
				Vector:         p.Values,
				CreatedAt:      p.CreatedAt.Unix(),
			},
			Distance: 0.0,
		}
	}

	return &pb.SearchNearestResponse{
		Results:        results,
		DistanceMetric: req.DistanceMetric,
		QueryVector:    req.QueryVector,
		TotalMatches:   int32(len(points)),
	}, nil
}
