package trip

import (
	"context"
	trippb "coolcar/server/proto/gen/go"
)

// Service implements trip service.
type Service struct{}

func (*Service) GetTrip(c context.Context, req *trippb.GetTripRequest) (*trippb.GettripResponse, error) {
	return &trippb.GettripResponse{
		Id: req.Id,
		Trip: &trippb.Trip{
			Start:       "aba",
			End:         "def",
			DurationSec: 3600,
			FeeCent:     10000,
			StartPos: &trippb.Location{
				Latitude:  30,
				Longitude: 120,
			},
			EndPos: &trippb.Location{
				Latitude:  35,
				Longitude: 115,
			},
			PathLocations: []*trippb.Location{
				{
					Latitude:  31,
					Longitude: 119,
				},
				{
					Latitude:  32,
					Longitude: 118,
				},
			},
		},
	}, nil
}
