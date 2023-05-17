package server

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

// LoggingInterceptor is a gRPC interceptor that logs incoming requests and outgoing responses.
func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf(info.FullMethod)

	// Invoke the gRPC handler to process the request
	resp, err := handler(ctx, req)

	// if err != nil {
	// 	log.Printf("Error: %v", err)
	// }
	// } else {
	// 	log.Printf("Response: %+v", resp)
	// }

	return resp, err
}
