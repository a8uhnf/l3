package tracing

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_nextRequestId(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "basic test",
			want: "string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := nextRequestId()
			require.Equal(t, reflect.TypeOf(got), reflect.TypeOf(tt.want), "should be equal")
		})
	}
}
