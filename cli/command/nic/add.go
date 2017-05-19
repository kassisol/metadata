package nic

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/juliengk/go-utils"
	"github.com/kassisol/metadata/cli/command"
	"github.com/kassisol/metadata/storage"
	"github.com/spf13/cobra"
)

var nicAddFloatingIP string
var nicAddIndex int

func newAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [mac] [ip]",
		Short: "Add interface",
		Long:  addDescription,
		Run:   runAdd,
	}

	flags := cmd.Flags()
	flags.StringVarP(&nicAddFloatingIP, "floating-ip", "f", "", "Floating IP address")
	flags.IntVarP(&nicAddIndex, "index", "i", 0, "NIC index")

	return cmd
}

func runAdd(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", command.DBFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer s.End()

	if len(args) < 2 || len(args) > 2 {
		cmd.Usage()
		os.Exit(-1)
	}

	s.AddInterface(nicAddIndex, args[0], args[1], nicAddFloatingIP)
}

var addDescription = `
Add interface

`
