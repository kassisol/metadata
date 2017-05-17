package metadata

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/kassisol/metadata/cli/command"
	"github.com/kassisol/metadata/storage"
	"github.com/labstack/echo"
)

func NetworkInterfacesIndexHandle(c echo.Context) error {
	var result []string

	serverid := c.Get("SERVERID").(int)

	s, err := storage.NewDriver("sqlite", command.DBFilePath)
	if err != nil {
		return err
	}
	defer s.End()

	r := s.GetInterfaces(serverid)
	for _, i := range r {
		result = append(result, fmt.Sprintf("%s/", i))
	}

	return c.String(http.StatusOK, strings.Join(result, "\n"))
}

func NetworkInterfaceTypeIndexHandle(c echo.Context) error {
	var result []string

	itype := c.Param("type")
	serverid := c.Get("SERVERID").(int)

	s, err := storage.NewDriver("sqlite", command.DBFilePath)
	if err != nil {
		return err
	}
	defer s.End()

	r := s.GetInterfacesType(serverid, itype)
	for _, i := range r {
		result = append(result, fmt.Sprintf("%d/", i))
	}

	return c.String(http.StatusOK, strings.Join(result, "\n"))
}

func NetworkEnumeratedInterfaceIndexHandle(c echo.Context) error {
	itype := c.Param("type")
	index := c.Param("num")
	serverid := c.Get("SERVERID").(int)

	result := []string{
		"mac",
		"type",
	}

	s, err := storage.NewDriver("sqlite", command.DBFilePath)
	if err != nil {
		return err
	}
	defer s.End()

	i, _ := strconv.Atoi(index)

	r := s.GetEnumeratedInterface(serverid, itype, i)
	for _, i := range r {
		result = append(result, fmt.Sprintf("%s/", i))
	}

	return c.String(http.StatusOK, strings.Join(result, "\n"))
}

func NetworkInterfaceMACAddressHandle(c echo.Context) error {
	itype := c.Param("type")
	index := c.Param("num")
	serverid := c.Get("SERVERID").(int)

	s, err := storage.NewDriver("sqlite", command.DBFilePath)
	if err != nil {
		return err
	}
	defer s.End()

	i, _ := strconv.Atoi(index)

	result := s.GetInterfaceMACAddress(serverid, itype, i)

	return c.String(http.StatusOK, result)
}

func NetworkInterfaceTypeHandle(c echo.Context) error {
	itype := c.Param("type")
	index := c.Param("num")
	serverid := c.Get("SERVERID").(int)

	s, err := storage.NewDriver("sqlite", command.DBFilePath)
	if err != nil {
		return err
	}
	defer s.End()

	i, _ := strconv.Atoi(index)

	result := s.GetInterfaceType(serverid, itype, i)

	return c.String(http.StatusOK, result)
}

func NetworkInterfaceIPv4IndexHandle(c echo.Context) error {
	result := []string{
		"address",
		"netmask",
		"gateway",
	}

	return c.String(http.StatusOK, strings.Join(result, "\n"))
}

func NetworkInterfaceIPv4AddressHandle(c echo.Context) error {
	itype := c.Param("type")
	index := c.Param("num")
	serverid := c.Get("SERVERID").(int)

	s, err := storage.NewDriver("sqlite", command.DBFilePath)
	if err != nil {
		return err
	}
	defer s.End()

	i, _ := strconv.Atoi(index)

	result := s.GetInterfaceIPv4Address(serverid, itype, i)

	return c.String(http.StatusOK, result)
}

func NetworkInterfaceIPv4NetmaskHandle(c echo.Context) error {
	itype := c.Param("type")
	index := c.Param("num")
	serverid := c.Get("SERVERID").(int)

	s, err := storage.NewDriver("sqlite", command.DBFilePath)
	if err != nil {
		return err
	}
	defer s.End()

	i, _ := strconv.Atoi(index)

	result := s.GetInterfaceIPv4Netmask(serverid, itype, i)

	return c.String(http.StatusOK, result)
}

func NetworkInterfaceIPv4GatewayHandle(c echo.Context) error {
	itype := c.Param("type")
	index := c.Param("num")
	serverid := c.Get("SERVERID").(int)

	s, err := storage.NewDriver("sqlite", command.DBFilePath)
	if err != nil {
		return err
	}
	defer s.End()

	i, _ := strconv.Atoi(index)

	result := s.GetInterfaceIPv4Gateway(serverid, itype, i)

	return c.String(http.StatusOK, result)
}
