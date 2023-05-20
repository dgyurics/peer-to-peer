package server

import (
	"context"
	"log"

	"github.com/dgyurics/p2p/api"
	pb "github.com/dgyurics/p2p/api"
	ev "github.com/dgyurics/p2p/internal/event"
	nwk "github.com/dgyurics/p2p/internal/network"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewGRPCServer() *grpc.Server {
	server := grpc.NewServer(
		grpc.UnaryInterceptor(LoggingInterceptor),
		// grpc.StreamInterceptor(LoggingStreamInterceptor),
	)
	pb.RegisterPeerToPeerServer(server, &peerToPeerServer{
		store:   ev.NewStore(),
		network: nwk.NewNetwork(),
	})
	return server
}

type peerToPeerServer struct {
	store   ev.Store
	network nwk.Network
	pb.UnimplementedPeerToPeerServer
}

func (s *peerToPeerServer) Connect(ctx context.Context, in *pb.ConnectRequest) (*pb.ConnectResponse, error) {
	if in.Address == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Address is required")
	}
	if err := s.network.Connect(ctx, nwk.Node{Address: in.Address}); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to connect to node: %v", err)
	}
	return &pb.ConnectResponse{Message: "Connection successful"}, nil
}

func (s *peerToPeerServer) Disconnect(ctx context.Context, in *pb.DisconnectRequest) (*pb.DisconnectResponse, error) {
	if in.Address == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Address is required")
	}
	if err := s.network.Disconnect(ctx, nwk.Node{Address: in.Address}); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to connect to node: %v", err)
	}
	return &pb.DisconnectResponse{Message: "Disconnection successful"}, nil
}

func (s *peerToPeerServer) Event(ctx context.Context, in *pb.EventRequest) (*pb.EventResponse, error) {
	if in.DeviceId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "DeviceId is required")
	}
	if in.Event == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Event is required")
	}
	event, err := s.store.Persist(in.DeviceId, in.Event)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to publish event: %v", err)
	}
	return &pb.EventResponse{
		EventId: event.ID.String(),
	}, nil
}

// EventStream implements bidirectional streaming RPC
func (s *peerToPeerServer) EventStream(stream api.PeerToPeer_EventStreamServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}

		// Process the received request
		log.Printf("Received request: %v", req)

		res, err := s.Event(stream.Context(), req)
		if err != nil {
			return err
		}

		if err = stream.Send(res); err != nil {
			return err
		}
	}
}

func (s *peerToPeerServer) Health(ctx context.Context, in *pb.Empty) (*pb.HealthResponse, error) {
	return &pb.HealthResponse{
		Status: pb.HealthStatus_HEALTH_STATUS_OK,
	}, nil
}
