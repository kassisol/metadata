package api

import (
	"github.com/kassisol/metadata/cli/command"
	"github.com/kassisol/metadata/storage"
	"github.com/labstack/echo"
)

func ServerIP() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			s, err := storage.NewDriver("sqlite", command.DBFilePath)
			if err != nil {
				return err
			}
			defer s.End()

			ip := c.RealIP()

			id := s.GetIDFromIP(ip)
			if id == 0 {
				return echo.ErrUnauthorized
			}

			c.Set("SERVERID", id)

			return next(c)
		}
	}
}
