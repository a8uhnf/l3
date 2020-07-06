package serve

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStart(t *testing.T) {
	tests := []struct {
		name string
	}{
		// {
		// 	"basic test",
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Start()
			err := got.Execute()
			require.Nil(t, err, "Start() commsand execute should produce nil")
		})
	}
}
