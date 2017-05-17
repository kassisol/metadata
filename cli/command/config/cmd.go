package config

import (
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "admin",
		Short: "Manage Metadata config",
		Long:  adminDescription,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	cmd.AddCommand(
		newBindCommand(),
	)

	return cmd
}

var adminDescription = `
The **metadata config** command has subcommands for managing Metadata config.

To see help for a subcommand, use:

    metadata config [command] --help

For full details on using metadata config visit Metadata's github repository.

`
