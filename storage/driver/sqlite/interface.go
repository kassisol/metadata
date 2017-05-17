package sqlite

import (
	"fmt"

	"github.com/kassisol/metadata/storage/driver"
)

func (c *Config) AddInterface(index int, mac, ip string) error {
	i := IP{}

	c.DB.Where("ip_address = ?", ip).First(&i)

	c.DB.Create(&Interface{
		Index:      index,
		MACAddress: mac,
		IPID:       i.ID,
	})

	return nil
}

func (c *Config) RemoveInterface(mac string) error {
	if c.memberOfHost(mac) {
		return fmt.Errorf("interface \"%s\" cannot be removed. It is being used by an host", mac)
	}

	c.DB.Where("mac_address = ?", mac).Delete(Interface{})

	return nil
}

func (c *Config) ListInterface(filter map[string]string) []driver.InterfaceResult {
	var result []driver.InterfaceResult

	sql := c.DB.Table("interfaces").Select("interfaces.mac_address, ips.id, ips.ip_address, ips.netmask, ips.gateway").Joins("JOIN ips ON ips.id = interfaces.ip_id")

	if v, ok := filter["ip"]; ok {
		sql = sql.Where("ips.ip_address = ?", v)
	}

	if v, ok := filter["netmask"]; ok {
		sql = sql.Where("ips.netmask = ?", v)
	}

	if v, ok := filter["gateway"]; ok {
		sql = sql.Where("ips.gateway = ?", v)
	}

	rows, _ := sql.Rows()
	defer rows.Close()

	for rows.Next() {
		var d_mac string
		var d_id uint
		var d_ip string
		var d_netmask string
		var d_gateway string

		rows.Scan(&d_mac, &d_id, &d_ip, &d_netmask, &d_gateway)

		ir := driver.IPResult{ID: d_id, IPAddress: d_ip, Netmask: d_netmask, Gateway: d_gateway}

		result = append(result, driver.InterfaceResult{MACAddress: d_mac, IP: ir})
	}

	return result
}

func (c *Config) CountInterface() int {
	var count int64

	c.DB.Model(&Interface{}).Count(&count)

	return int(count)
}

func (c *Config) memberOfHost(mac string) bool {
	var count int64

	c.DB.Table("hosts").Joins("JOIN host_interfaces ON host_interfaces.host_id = hosts.id").Joins("JOIN interfaces ON interfaces.id = host_interfaces.host_id").Where("interfaces.mac_address = ?", mac).Count(&count)

	if count > 0 {
		return true
	}

	return false
}
