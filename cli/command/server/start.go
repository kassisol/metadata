package server

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/juliengk/go-utils/filedir"
	"github.com/kassisol/metadata/api"
	"github.com/kassisol/metadata/cli/command"
	"github.com/kassisol/metadata/storage"
	"github.com/spf13/cobra"
)

var (
	serverBindAddress string
	serverBindPort    string
)

func newStartCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Starts a metadata server",
		Long:  startDescription,
		Run:   runStart,
	}

	flags := cmd.Flags()
	flags.StringVarP(&serverBindAddress, "bind-address", "a", "", "Bind Address")
	flags.StringVarP(&serverBindPort, "bind-port", "p", "", "Bind Port")

	return cmd
}

func runStart(cmd *cobra.Command, args []string) {
	var bindAddress string
	var bindPort string

	if len(args) > 0 {
		cmd.Usage()
		os.Exit(-1)
	}

	if !filedir.FileExists(command.DBFilePath) {
		log.Fatal("Initialization needs to be done first")
	}

	s, err := storage.NewDriver("sqlite", command.DBFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer s.End()

	if serverBindAddress != "" {
		bindAddress = serverBindAddress
	} else {
		bindAddress = s.GetConfig("api_bind")[0].Value
	}

	if serverBindPort != "" {
		bindPort = serverBindPort
	} else {
		bindPort = s.GetConfig("api_port")[0].Value
	}

	addr := fmt.Sprintf("%s:%s", bindAddress, bindPort)

	api.API(addr)
}

var startDescription = `
Starts a metadata server

`
