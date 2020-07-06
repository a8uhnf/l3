package server

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	type args struct {
		port          string
		sectorID      float64
		serverRunning chan bool
	}
	tests := []struct {
		name string
		args args
		want *Server
	}{
		{
			name: "basic test",
			args: args{
				port:          "8080",
				sectorID:      1.0,
				serverRunning: make(chan bool),
			},
			want: &Server{
				Port:          "8080",
				SectorID:      1.0,
				ServerRunning: make(chan bool),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.port, tt.args.sectorID)
			require.Equalf(t, got.SectorID, tt.want.SectorID, "got %v, but wanted %v", got.SectorID, tt.want.SectorID)
			require.Equalf(t, reflect.TypeOf(got.ServerRunning), reflect.TypeOf(tt.want.ServerRunning), "got %v, but wanted %v", got.ServerRunning, tt.want.ServerRunning)
		})
	}
}

func TestServer_Spawn(t *testing.T) {
	type fields struct {
		Port string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "basic test",
			fields:  struct{ Port string }{Port: string("8080")},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New("8080", 1)
			serverStarted := make(chan bool)
			go func() {
				close(serverStarted)
				if err := s.Spawn(); (err != nil) != tt.wantErr {
					t.Errorf("Server.Spawn() error = %v, wantErr %v", err, tt.wantErr)
				}
			}()
			<-serverStarted
			s.ServerRunning <- true
		})
	}
}
