package metadata

import (
	"net/http"
	"strings"

	"github.com/kassisol/metadata/cli/command"
	"github.com/kassisol/metadata/storage"
	"github.com/labstack/echo"
)

func KeysIndexHandle(c echo.Context) error {
	serverid := c.Get("SERVERID").(int)

	s, err := storage.NewDriver("sqlite", command.DBFilePath)
	if err != nil {
		return err
	}
	defer s.End()

	result := s.GetKeys(serverid)

	return c.String(http.StatusOK, strings.Join(result, "\n"))
}

func KeyNameHandle(c echo.Context) error {
	serverid := c.Get("SERVERID").(int)
	key := c.Param("name")

	s, err := storage.NewDriver("sqlite", command.DBFilePath)
	if err != nil {
		return err
	}
	defer s.End()

	result := s.GetKey(serverid, key)

	return c.String(http.StatusOK, result)
}
