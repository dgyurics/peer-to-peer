package network

import (
	"context"
	"errors"
	"sync"

	pb "github.com/dgyurics/p2p/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Network represents a set of functions responsible
// for communicating with other nodes in the network.
type Network interface {
	Connect(ctx context.Context, node Node) error
	Disconnect(ctx context.Context, node Node) error
	Publish(ctx context.Context, event *pb.Event) error
}

func NewNetwork() Network {
	return &network{
		mu: sync.Mutex{},
	}
}

type network struct {
	mu    sync.Mutex
	nodes []Node
}

func (n *network) Connect(ctx context.Context, node Node) error {
	n.mu.Lock()
	defer n.mu.Unlock()

	// Verify node not already connected
	for _, n := range n.nodes {
		if n.Address == node.Address {
			return errors.New("node already connected")
		}
	}
	if err := n.ping(ctx, node); err != nil {
		return err
	}
	n.nodes = append(n.nodes, node)
	return nil
}

func (n *network) Disconnect(ctx context.Context, node Node) error {
	n.mu.Lock()
	defer n.mu.Unlock()
	for i, nd := range n.nodes {
		if nd.Address == node.Address {
			n.nodes = append(n.nodes[:i], n.nodes[i+1:]...)
			return nil
		}
	}
	return errors.New("node not found")
}

// this will publish events one by one...
// but if we want to stream events to neighbors...?
func (n *network) Publish(ctx context.Context, event *pb.Event) error {
	n.mu.Lock()
	defer n.mu.Unlock()
	for _, node := range n.nodes {
		// Set up a connection to the server.
		conn, err := grpc.Dial(node.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return err
		}
		defer conn.Close()
		c := pb.NewPeerToPeerClient(conn)

		_, err = c.RelayEvent(ctx, event)
		if err != nil {
			return err
		}
	}
	return nil
}

func (n *network) ping(ctx context.Context, node Node) error {
	addr := node.Address
	// Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()
	c := pb.NewPeerToPeerClient(conn)

	res, err := c.Health(ctx, &pb.Empty{})
	if err != nil {
		return errors.New("failed to ping node")
	}
	if res.Status != pb.HealthStatus_HEALTH_STATUS_OK {
		return errors.New("node is not healthy")
	}
	return nil
}
