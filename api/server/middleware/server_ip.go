package api

import (
	"github.com/kassisol/metadata/api/storage"
	"github.com/kassisol/metadata/pkg/adf"
	"github.com/labstack/echo"
)

func ServerIP() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cfg := adf.NewDaemon()
			if err := cfg.Init(); err != nil {
				return err
			}

			s, err := storage.NewDriver("sqlite", cfg.App.Dir.Root)
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
