package system

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/juliengk/go-utils/filedir"
	"github.com/kassisol/metadata/cli/command"
	"github.com/kassisol/metadata/storage"
	"github.com/kassisol/metadata/version"
	"github.com/spf13/cobra"
)

func NewInfoCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info",
		Short: "Display information about Metadata",
		Long:  infoDescription,
		Run:   runInfo,
	}

	return cmd
}

func runInfo(cmd *cobra.Command, args []string) {
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

	fmt.Println("Datas:", s.CountData())
	fmt.Println("Profiles:", s.CountProfile())
	fmt.Println("IPs:", s.CountIP())
	fmt.Println("Interfaces:", s.CountInterface())
	fmt.Println("Hosts:", s.CountHost())

	fmt.Println("API:")
	fmt.Println(" FQDN:", s.GetConfig("api_fqdn")[0].Value)
	fmt.Println(" Bind Address:", s.GetConfig("api_bind")[0].Value)
	fmt.Println(" Bind Port:", s.GetConfig("api_port")[0].Value)

	fmt.Println("Server Version:", version.Version)
	fmt.Println("Storage Driver: sqlite")
	fmt.Println("Metadata Root Dir:", command.AppPath)
}

var infoDescription = `
Display information about Metadata

`
