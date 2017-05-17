package server

import (
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Starts a metadat server",
		Long:  serverDescription,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	cmd.AddCommand(
		newStartCommand(),
	)

	return cmd
}

var serverDescription = `
The **metadata server** command has subcommands for starting a metadata server.

To see help for a subcommand, use:

    metadata server [command] --help

For full details on using metadata server visit metadata's github repository.

`
