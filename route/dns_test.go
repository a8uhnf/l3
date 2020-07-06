package route

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"

	data_types "github.com/a8uhnf/l3/pkg/api"
)

func Test_dnsHandler(t *testing.T) {
	url := "/location"
	type fields struct {
		SectorId float64
	}
	type args struct {
		request data_types.DNSRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    error
		wantErr bool
	}{
		{
			name: "basic test",
			fields: fields{
				SectorId: 1.0,
			},
			args: args{
				request: data_types.DNSRequest{
					1.0,
					1.0,
					1.0,
					1.0,
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		e := echo.New()
		buf := new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(tt.args.request)
		require.Nil(t, err, "json decoder's error should be nil")
		req := httptest.NewRequest(http.MethodPost, "/location", buf)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(url)
		dns := NewDNS(tt.fields.SectorId)
		err = dns.dnsHandler(c)
		require.Nil(t, err, "error should be nil")
		require.Equalf(t, 200, rec.Code, "status should be 200, but found %v", rec.Code)
	}
}

func TestNewDNS(t *testing.T) {
	type args struct {
		sectorID float64
	}
	tests := []struct {
		name string
		args args
		want *DNS
	}{
		{
			name: "basic test",
			args: args{
				sectorID: 1.0,
			},
			want: &DNS{
				SectorID: 1.0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDNS(tt.args.sectorID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDNS() = %v, want %v", got, tt.want)
			}
		})
	}
}
