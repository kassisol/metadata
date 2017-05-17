package system

import (
	"os"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/juliengk/go-utils/filedir"
	"github.com/juliengk/go-utils/readinput"
	"github.com/juliengk/go-utils/validation"
	"github.com/kassisol/metadata/cli/command"
	"github.com/kassisol/metadata/storage"
	"github.com/spf13/cobra"
)

var (
	serverAPIFQDN string
	serverAPIBind string
	serverAPIPort string
)

func NewInitCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize config",
		Long:  initDescription,
		Run:   runInit,
	}

	flags := cmd.Flags()

	flags.StringVarP(&serverAPIFQDN, "api-fqdn", "", "", "API FQDN")
	flags.StringVarP(&serverAPIBind, "api-bind", "", "0.0.0.0", "API Bind Interface")
	flags.StringVarP(&serverAPIPort, "api-port", "", "80", "API Port")

	return cmd
}

func runInit(cmd *cobra.Command, args []string) {
	var apifqdn string
	var apibind string
	var apiport string

	if len(args) > 0 {
		cmd.Usage()
		os.Exit(-1)
	}

	if filedir.FileExists(command.DBFilePath) {
		log.Info("Initialization already done")
		os.Exit(0)
	}

	if err := filedir.CreateDirIfNotExist(command.AppPath, false, 0700); err != nil {
		log.Fatal(err)
	}

	if len(serverAPIFQDN) <= 0 {
		apifqdn = readinput.ReadInput("API FQDN")
		if len(apifqdn) <= 0 {
			af, err := os.Hostname()
			if err != nil {
				log.Fatal(err)
			}
			apifqdn = af
		}
	} else {
		apifqdn = serverAPIFQDN
	}

	if len(serverAPIBind) <= 0 {
		apibind = readinput.ReadInput("API Bind")
	} else {
		apibind = serverAPIBind
	}

	if len(serverAPIPort) <= 0 {
		apiport = readinput.ReadInput("API Port")
	} else {
		apiport = serverAPIPort
	}

	// DB
	s, err := storage.NewDriver("sqlite", command.DBFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer s.End()

	// Input validations
	// IV - API FQDN
	if err = validation.IsValidFQDN(apifqdn); err != nil {
		log.Fatal(err)
	}

	// IV - API Bind
	if err = validation.IsValidIP(apibind); err != nil {
		log.Fatal(err)
	}

	// IV - API Port
	port, err := strconv.Atoi(apiport)
	if err != nil {
		log.Fatal(err)
	}
	if err = validation.IsValidPort(port); err != nil {
		log.Fatal(err)
	}

	s.AddConfig("api_fqdn", apifqdn)
	s.AddConfig("api_bind", apibind)
	s.AddConfig("api_port", apiport)
}

var initDescription = `
Initialize config

`
