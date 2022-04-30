package pathctrl

import (
	"context"

	"github.com/cool-develope/volume-flight-golang/pathctrl/pb"
)

type flightService struct {
	pb.UnimplementedFlightServiceServer
}

// NewFlightService returns a new FlightServiceServer
func NewFlightService() pb.FlightServiceServer {
	return flightService{}
}

// GetSortedFlight returns the sorted result
func (flightService) GetSortedFlight(ctx context.Context, req *pb.GetSortedFlightRequest) (*pb.GetSortedFlightResponse, error) {
	flight, err := getSortedFlight(req.Flights)

	return &pb.GetSortedFlightResponse{
		Sorted: flight,
	}, err
}
