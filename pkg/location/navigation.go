package location

import (
	data_types "github.com/a8uhnf/l3/pkg/api"
)

type DNS struct {
	SectorID float64
}

// New initialize drone navigation system struct
func New() *DNS {
	return &DNS{}
}

// CalculateLocation calculate location for a specific sectorIDs
func (d *DNS) CalculateLocation(request data_types.DNSRequest) float64 {
	location := ((request.X + request.Y + request.Z) * d.SectorID) + request.Vel

	return location
}
