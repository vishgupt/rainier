package api

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/vishgupt/rainier/src/internal/common"
)

// convertError converts application errors to gRPC status errors
func convertError(err error) error {
	if err == nil {
		return nil
	}

	if appErr, ok := err.(*common.AppError); ok {
		switch appErr.Type {
		case common.ValidationError:
			return status.Error(codes.InvalidArgument, appErr.Message)
		case common.NotFoundError:
			return status.Error(codes.NotFound, appErr.Message)
		case common.InternalError:
			return status.Error(codes.Internal, appErr.Message)
		}
	}

	return status.Error(codes.Internal, "internal server error")
}
