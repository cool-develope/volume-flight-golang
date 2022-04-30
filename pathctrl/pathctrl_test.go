package pathctrl

import (
	"testing"

	"github.com/cool-develope/volume-flight-golang/pathctrl/pb"
	"github.com/stretchr/testify/require"
)

type FlightVectors struct {
	flights []*pb.Flight
	sorted  *pb.Flight
}

func TestFlight(t *testing.T) {
	invalidFlight := &pb.Flight{
		Source:      "SFO",
		Destination: "SFO",
	}

	validFlight := &pb.Flight{
		Source:      "SFO",
		Destination: "EWR",
	}

	t.Run("Invalid Flight", func(t *testing.T) {
		res := checkFlight(invalidFlight)
		require.Equal(t, res, false)
	})

	t.Run("Valid Flight", func(t *testing.T) {
		res := checkFlight(validFlight)
		require.Equal(t, res, true)
	})
}

func TestPathCtrl(t *testing.T) {
	validTestVectors := []FlightVectors{
		{
			flights: []*pb.Flight{
				{
					Source:      "SFO",
					Destination: "EWR",
				},
			},
			sorted: &pb.Flight{
				Source:      "SFO",
				Destination: "EWR",
			},
		},
		{
			flights: []*pb.Flight{
				{
					Source:      "ATL",
					Destination: "EWR",
				},
				{
					Source:      "SFO",
					Destination: "ATL",
				},
			},
			sorted: &pb.Flight{
				Source:      "SFO",
				Destination: "EWR",
			},
		},
		{
			flights: []*pb.Flight{
				{
					Source:      "IND",
					Destination: "EWR",
				},
				{
					Source:      "SFO",
					Destination: "ATL",
				},
				{
					Source:      "GSO",
					Destination: "IND",
				},
				{
					Source:      "ATL",
					Destination: "GSO",
				},
			},
			sorted: &pb.Flight{
				Source:      "SFO",
				Destination: "EWR",
			},
		},
	}

	invalidTestVectors := []FlightVectors{
		{
			flights: []*pb.Flight{
				{
					Source:      "ATL",
					Destination: "EWR",
				},
				{
					Source:      "EWR",
					Destination: "ATL",
				},
			},
			sorted: &pb.Flight{
				Source:      "SFO",
				Destination: "EWR",
			},
		},
		{
			flights: []*pb.Flight{
				{
					Source:      "IND",
					Destination: "EWR",
				},
				{
					Source:      "SFO",
					Destination: "ATL",
				},
				{
					Source:      "GSO",
					Destination: "SFO",
				},
				{
					Source:      "ATL",
					Destination: "GSO",
				},
			},
			sorted: &pb.Flight{
				Source:      "SFO",
				Destination: "EWR",
			},
		},
	}

	t.Run("Test valid flights", func(t *testing.T) {
		for _, testVector := range validTestVectors {
			res, err := getSortedFlight(testVector.flights)
			require.NoError(t, err)
			require.Equal(t, testVector.sorted, res)
		}
	})

	t.Run("Test invalid flights", func(t *testing.T) {
		for _, testVector := range invalidTestVectors {
			_, err := getSortedFlight(testVector.flights)
			require.EqualError(t, err, ErrMessyPath.Error())
		}
	})
}
