package nic

import (
	"fmt"
	"os"
	"text/tabwriter"

	log "github.com/Sirupsen/logrus"
	"github.com/juliengk/go-utils"
	"github.com/kassisol/metadata/cli/command"
	"github.com/kassisol/metadata/storage"
	"github.com/spf13/cobra"
)

var nicListFilter []string

func newListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ls",
		Aliases: []string{"list"},
		Short:   "List interfaces",
		Long:    listDescription,
		Run:     runList,
	}

	flags := cmd.Flags()
	flags.StringSliceVarP(&nicListFilter, "filter", "f", []string{}, "Filter output based on conditions provided")

	return cmd
}

func runList(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", command.DBFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer s.End()

	filters := utils.ConvertSliceToMap("=", nicListFilter)

	interfaces := s.ListInterface(filters)

	if len(interfaces) > 0 {
		w := tabwriter.NewWriter(os.Stdout, 20, 1, 2, ' ', 0)
		fmt.Fprintln(w, "MAC ADDRESS\tIP ADDRESS\tNETMASK\tGATEWAY")

		for _, inf := range interfaces {
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", inf.MACAddress, inf.IP.IPAddress, inf.IP.Netmask, inf.IP.Gateway)
		}

		w.Flush()
	}
}

var listDescription = `
List interfaces

`
