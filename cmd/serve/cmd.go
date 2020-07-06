package serve

import (
	"github.com/spf13/cobra"

	"github.com/a8uhnf/l3/server"
)

// Start spawn http server
func Start() *cobra.Command {
	c := &cobra.Command{
		Use:   "serve",
		Short: "Start drone navigation system web-server",
		RunE: func(cmd *cobra.Command, args []string) error {
			port, err := cmd.Flags().GetString("port")
			if err != nil {
				return err
			}
			sectorID, err := cmd.Flags().GetFloat64("sector-id")
			if err != nil {
				return err
			}
			s := server.New(port, sectorID)
			if err := s.Spawn(); err != nil {
				return err
			}
			return nil
		},
	}
	c.Flags().String("port", "8080", "web server port")
	c.Flags().Float64("sector-id", 1.0, "drone navigation system sector id")

	return c
}
