package host

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/juliengk/go-utils"
	"github.com/kassisol/metadata/cli/command"
	"github.com/kassisol/metadata/storage"
	"github.com/spf13/cobra"
)

var (
	hostAddEnable    bool
	hostAddFQDN      string
	hostAddUUID      string
	hostAddProfile   string
	hostAddInterface []string
)

func newAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [name]",
		Short: "Add host",
		Long:  addDescription,
		Run:   runAdd,
	}

	flags := cmd.Flags()
	flags.BoolVarP(&hostAddEnable, "enable", "e", true, "Enable host entry")
	flags.StringVarP(&hostAddFQDN, "fqdn", "f", "", "Set FQDN")
	flags.StringVarP(&hostAddUUID, "uuid", "u", "", "Set UUID")
	flags.StringVarP(&hostAddProfile, "profile", "p", "", "Set profile")
	flags.StringSliceVarP(&hostAddInterface, "interface", "i", []string{}, "Set interfaces")

	return cmd
}

func runAdd(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", command.DBFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer s.End()

	if len(args) < 1 || len(args) > 1 {
		cmd.Usage()
		os.Exit(-1)
	}

	if err := s.AddHost(hostAddEnable, args[0], hostAddFQDN, hostAddUUID, hostAddProfile, hostAddInterface); err != nil {
		log.Fatal(err)
	}
}

var addDescription = `
Add host

`
