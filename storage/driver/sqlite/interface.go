package sqlite

import (
	"fmt"

	"github.com/kassisol/metadata/storage/driver"
)

func (c *Config) AddInterface(index int, mac, ip, floatingIP string) error {
	i := IP{}
	c.DB.Where("ip_address = ?", ip).First(&i)

	intf := Interface{
		Index:      index,
		MACAddress: mac,
		IPID:       i.ID,
	}

	if len(floatingIP) > 0 {
		f := IP{}
		c.DB.Where("ip_address = ?", floatingIP).First(&f)

		intf.FloatingIPID = f.ID
	}

	c.DB.Create(&intf)

	return nil
}

func (c *Config) UpdateInterface(mac, itype, value string) error {
	r := c.ListIP(map[string]string{"ip": value})
	if len(r) == 0 {
		return fmt.Errorf("IP address '%s' does not exist", value)
	}

	i := IP{}
	c.DB.Where("ip_address = ?", value).First(&i)

	intf := Interface{}
	c.DB.Where("mac_address= ?", mac).First(&intf)

	if itype == "ip" {
		intf.IPID = i.ID
	}
	if itype == "floatingip" {
		intf.FloatingIPID = i.ID
	}

	c.DB.Save(&intf)

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

	sql := c.DB.Table("interfaces").Select("`index`, mac_address, ip_id, floating_ip_id")

	rows, _ := sql.Rows()
	defer rows.Close()

	for rows.Next() {
		var d_index int
		var d_mac string
		var d_ip string
		var d_floating_ip string

		rows.Scan(&d_index, &d_mac, &d_ip, &d_floating_ip)

		ir := c.ListIP(map[string]string{"id": d_ip})[0]

		fir := driver.IPResult{}
		if d_floating_ip != "0" {
			fir = c.ListIP(map[string]string{"id": d_floating_ip})[0]
		}

		result = append(result, driver.InterfaceResult{Index: d_index, MACAddress: d_mac, IP: ir, FloatingIP: fir})
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
