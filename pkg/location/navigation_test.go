package location

import (
	"testing"

	data_types "github.com/a8uhnf/l3/pkg/api"
)

func TestDNS_CalculateLocation(t *testing.T) {
	type fields struct {
		SectorID float64
	}
	type args struct {
		request data_types.DNSRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr bool
	}{
		{
			name:   "basic test 1",
			fields: struct{ SectorID float64 }{SectorID: 1},
			args: struct{ request data_types.DNSRequest }{
				request: data_types.DNSRequest{
					X: 1.0, Y: 1.0, Z: 1.0, Vel: 1.0,
				},
			},
			want:    4.0,
			wantErr: false,
		},
		{
			name:   "basic test 2",
			fields: struct{ SectorID float64 }{SectorID: 1},
			args: struct{ request data_types.DNSRequest }{
				request: data_types.DNSRequest{
					X: 2.0, Y: 2.0, Z: 2.0, Vel: 2.0,
				},
			},
			want:    8.0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DNS{
				SectorID: tt.fields.SectorID,
			}
			got := d.CalculateLocation(tt.args.request)
			if got != tt.want {
				t.Errorf("DNS.CalculateLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}
