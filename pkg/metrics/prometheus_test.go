package metrics

import (
	"fmt"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
)

func TestRegisterMetrics(t *testing.T) {
	name := "l3"
	type args struct {
		name string
	}
	prom := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: name,
		Help: fmt.Sprintf("this histogram is for %v api", name),
	}, []string{"method", "endpoint"})
	tests := []struct {
		name string
		args args
		want *prometheus.HistogramVec
	}{
		{
			name: "basic test",
			args: args{
				name: name,
			},
			want: prom,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RegisterMetrics(tt.args.name); got == tt.want {
				t.Errorf("RegisterMetrics() = %v, want %v", got, tt.want)
			}
		})
	}
}
