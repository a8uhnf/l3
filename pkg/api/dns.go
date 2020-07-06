package api

type DNSRequest struct {
	X   float64 `json:"x,string" validate:"required"`
	Y   float64 `json:"y,string" validate:"required"`
	Z   float64 `json:"z,string" validate:"required"`
	Vel float64 `json:"vel,string" validate:"required"`
}

type DNSResponse struct {
	Loc float64 `json:"loc"`
}
