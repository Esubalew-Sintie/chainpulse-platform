package dto

import (
	"encoding/json"
	"net/http"

	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Clean write for success
func WriteSuccessResponse(w http.ResponseWriter, data interface{}, message string) {
	w.Header().Set("Content-Type", "application/json")
	
	w.WriteHeader(http.StatusOK)

	response := APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	}

	_ = json.NewEncoder(w).Encode(response)
}

// Generic error response
func WriteErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := APIResponse{
		Success: false,
		Message: message,
	}

	_ = json.NewEncoder(w).Encode(response)
}

// New unified gRPC error writer
func WriteGrpcErrorResponse(w http.ResponseWriter, err error) {
	st, ok := status.FromError(err)
	if !ok {
		WriteErrorResponse(w, http.StatusInternalServerError, "Unexpected error occurred")
		return
	}
	WriteErrorResponse(w, grpcCodeToHTTPStatus(st.Code()), st.Message())
}

func grpcCodeToHTTPStatus(code codes.Code) int {
	switch code {
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	default:
		return http.StatusInternalServerError
	}
}
