package metadata

import (
	"net/http"
	"strings"

	"github.com/kassisol/metadata/cli/command"
	"github.com/kassisol/metadata/storage"
	"github.com/labstack/echo"
)

func TagsIndexHandle(c echo.Context) error {
	serverid := c.Get("SERVERID").(int)

	s, err := storage.NewDriver("sqlite", command.DBFilePath)
	if err != nil {
		return err
	}
	defer s.End()

	result := s.GetTags(serverid)

	return c.String(http.StatusOK, strings.Join(result, "\n"))
}
