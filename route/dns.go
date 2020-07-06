package route

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"

	data_types "github.com/a8uhnf/l3/pkg/api"
	"github.com/a8uhnf/l3/pkg/location"
)

type DNS struct {
	SectorID float64
}

// NewDNS initialize
func NewDNS(sectorID float64) *DNS {
	return &DNS{
		SectorID: sectorID,
	}
}

// dnsHandler handle location calculate request
// and send the calculated response to drones
func (d *DNS) dnsHandler(ctx echo.Context) error {
	req := &data_types.DNSRequest{}
	err := json.NewDecoder(ctx.Request().Body).Decode(req)
	if err != nil {
		return ctx.NoContent(http.StatusUnprocessableEntity)
	}
	v := validator.New()
	err = v.Struct(req)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	loc := location.New()
	loc.SectorID = d.SectorID
	position := loc.CalculateLocation(*req)

	resp := data_types.DNSResponse{
		Loc: position,
	}
	ctx.Logger().Debug("got location %v", position)
	return ctx.JSON(http.StatusOK, resp)
}

// AddDNS http handler for drone navigation system
func AddDNS(echo *echo.Echo, sectorId float64) {
	dns := NewDNS(sectorId)
	g := echo.Group("/api/v1")

	g.POST("/location", dns.dnsHandler)
}
