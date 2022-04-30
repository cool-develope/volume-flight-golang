package pathctrl

import (
	"errors"

	"github.com/cool-develope/volume-flight-golang/pathctrl/pb"
)

var (
	// ErrMessyPath error
	ErrMessyPath = errors.New("messy path")
	// ErrZeroPath error
	ErrZeroPath = errors.New("zero length path")
	// ErrInvalidFlight error
	ErrInvalidFlight = errors.New("invalid flight")
)

func checkFlight(flight *pb.Flight) bool {
	if flight.Destination == flight.Source {
		return false
	}

	// TODO
	// check if the airport code is valid

	return true
}

func getSortedFlight(flights []*pb.Flight) (*pb.Flight, error) {
	if len(flights) == 0 {
		return nil, ErrZeroPath
	}
	for _, f := range flights {
		if !checkFlight(f) {
			return nil, ErrInvalidFlight
		}
	}

	var (
		edges   = make(map[string]string)
		visited = make(map[string]bool)
	)

	for _, f := range flights {
		// check if exist multiple sources
		if _, ok := edges[f.Source]; ok {
			return nil, ErrMessyPath
		}
		edges[f.Source] = f.Destination

		// check if exist multiple destinations
		if visited[f.Destination] {
			return nil, ErrMessyPath
		}
		visited[f.Destination] = true
	}

	for _, f := range flights {
		// check the initial source
		if !visited[f.Source] {
			source := f.Source
			destination := f.Source
			// check if the path is a directed line
			for i := 0; i < len(flights); i++ {
				if dest, ok := edges[destination]; ok {
					destination = dest
				} else {
					return nil, ErrMessyPath
				}
			}
			return &pb.Flight{
				Source:      source,
				Destination: destination,
			}, nil
		}
	}

	return nil, ErrMessyPath
}
