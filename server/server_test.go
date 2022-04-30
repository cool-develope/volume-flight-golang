package server

import (
	"context"
	"testing"

	"github.com/cool-develope/volume-flight-golang/pathctrl/pb"
	"github.com/stretchr/testify/require"
)

const grpcPort = "9090"

func TestServer(t *testing.T) {
	go func() {
		RunServer(Config{
			GRPCPort: grpcPort,
		})
	}()

	address := "0.0.0.0:" + grpcPort
	w := NewWait()
	err := w.GRPCHealthy(address)
	require.NoError(t, err)

	conn, err := GetgRPCConn(address)
	require.NoError(t, err)
	defer conn.Close()

	client := pb.NewFlightServiceClient(conn)
	res, err := client.GetSortedFlight(context.TODO(),
		&pb.GetSortedFlightRequest{
			Flights: []*pb.Flight{
				{
					Source:      "SFO",
					Destination: "EWR",
				},
			},
		},
	)
	require.NoError(t, err)
	require.Equal(t, res.Sorted.Source, "SFO")
	require.Equal(t, res.Sorted.Destination, "EWR")
}
