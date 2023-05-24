# Peer To Peer

This project implements a peer-to-peer communication service using gRPC and protocol buffers. It is simply a proof of concept.

## Project Structure

The project is structured as follows:

- `api`: Contains the protocol buffer definition for the PeerToPeer service.
- `cmd`: Contains the main application code.
- `internal/config`: Contains the configuration files for the application.
- `internal/event`: Contains the implementation of the event store.
- `internal/network`: Contains the implementation of the network layer.
- `internal/server`: Contains the implementation of the gRPC server.

## Usage

The PeerToPeer service provides the following RPC methods:

- `Connect`: Connects to a peer at the specified address.
- `Disconnect`: Disconnects from a peer at the specified address.
- `ProduceEvent`: Produces an event for the specified device.
- `ProduceEventStream`: Produces a stream of events for the specified device.
- `RelayEvent`: Relays an event to all connected peers.
- `RelayEventStream`: Relays a stream of events to all connected peers.
- `Health`: Returns the health status of the service.

To use the service, you can use the provided gRPC client.
