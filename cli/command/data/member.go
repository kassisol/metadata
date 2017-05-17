package data

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/juliengk/go-utils"
	"github.com/kassisol/metadata/cli/command"
	"github.com/kassisol/metadata/storage"
	"github.com/spf13/cobra"
)

var (
	dataMemberAdd    bool
	dataMemberRemove bool
)

func newMemberCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "member [profile] [data]",
		Short: "Manage data membership to profile",
		Long:  memberDescription,
		Run:   runMember,
	}

	flags := cmd.Flags()
	flags.BoolVarP(&dataMemberAdd, "add", "a", false, "Add data to profile")
	flags.BoolVarP(&dataMemberRemove, "remove", "r", false, "Remove data from profile")

	return cmd
}

func runMember(cmd *cobra.Command, args []string) {
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

	/*
		if !s.FindProfile(args[0]) {
			log.Fatalf("%s does not exist", args[0])
		}

		if !s.FindData(args[1]) {
			log.Fatalf("%s does not exist", args[1])
		}
	*/

	if dataMemberAdd {
		s.AddDataToProfile(args[0], args[1])
	}
	if dataMemberRemove {
		s.RemoveDataFromProfile(args[0], args[1])
	}
}

var memberDescription = `
Manage data membership to profile

`
