package metrics

import (
	"testing"

	metricspkg "github.com/a8uhnf/l3/pkg/metrics"
)

func TestNew(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *Stats
	}{
		{
			name: "basic",
			args: args{
				name: "l3",
			},
			want: &Stats{
				prometheusHistogram: metricspkg.RegisterMetrics("l3"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.name); got == tt.want {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
