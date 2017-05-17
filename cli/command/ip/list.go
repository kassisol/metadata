package ip

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

var ipListFilter []string

func newListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ls",
		Aliases: []string{"list"},
		Short:   "List IP addresses",
		Long:    listDescription,
		Run:     runList,
	}

	flags := cmd.Flags()
	flags.StringSliceVarP(&ipListFilter, "filter", "f", []string{}, "Filter output based on conditions provided")

	return cmd
}

func runList(cmd *cobra.Command, args []string) {
	//	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", command.DBFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer s.End()

	filters := utils.ConvertSliceToMap("=", ipListFilter)

	ips := s.ListIP(filters)

	if len(ips) > 0 {
		w := tabwriter.NewWriter(os.Stdout, 20, 1, 2, ' ', 0)
		fmt.Fprintln(w, "IP ADDRESS\tNETMASK\tGATEWAY")

		for _, ip := range ips {
			fmt.Fprintf(w, "%s\t%s\t%s\n", ip.IPAddress, ip.Netmask, ip.Gateway)
		}

		w.Flush()
	}
}

var listDescription = `
List IP addresses

`
