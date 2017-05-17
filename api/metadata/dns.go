package metadata

import (
	"net/http"
	"strings"

	"github.com/kassisol/metadata/cli/command"
	"github.com/kassisol/metadata/storage"
	"github.com/labstack/echo"
)

func DnsIndexHandle(c echo.Context) error {
	serverid := c.Get("SERVERID").(int)

	s, err := storage.NewDriver("sqlite", command.DBFilePath)
	if err != nil {
		return err
	}
	defer s.End()

	result := s.GetDNSIndex(serverid)

	return c.String(http.StatusOK, strings.Join(result, "\n"))
}

func DnsNameserversHandle(c echo.Context) error {
	serverid := c.Get("SERVERID").(int)

	s, err := storage.NewDriver("sqlite", command.DBFilePath)
	if err != nil {
		return err
	}
	defer s.End()

	result := s.GetDNSNameservers(serverid)

	return c.String(http.StatusOK, strings.Join(result, "\n"))
}

func DnsSearchDomainsHandle(c echo.Context) error {
	serverid := c.Get("SERVERID").(int)

	s, err := storage.NewDriver("sqlite", command.DBFilePath)
	if err != nil {
		return err
	}
	defer s.End()

	result := s.GetDNSSearchDomains(serverid)

	return c.String(http.StatusOK, strings.Join(result, "\n"))
}

func DnsOptionsHandle(c echo.Context) error {
	serverid := c.Get("SERVERID").(int)

	s, err := storage.NewDriver("sqlite", command.DBFilePath)
	if err != nil {
		return err
	}
	defer s.End()

	result := s.GetDNSOptions(serverid)

	return c.String(http.StatusOK, strings.Join(result, "\n"))
}
