package main

import (
	"github.com/prometheus/common/log"
	"github.com/spf13/cobra"

	"github.com/a8uhnf/l3/cmd/serve"
)

var rootCommand = &cobra.Command{
	Use:   "l3",
	Short: "main command for drone navigation system",
}

func main() {
	rootCommand.AddCommand(serve.Start())
	if err := rootCommand.Execute(); err != nil {
		log.Error(err)
	}
}
