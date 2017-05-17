package sqlite

import (
	"github.com/kassisol/metadata/storage/driver"
)

func (c *Config) AddHost(enable bool, name, fqdn, uuid, profile string, interfaces []string) error {
	p := Profile{}

	c.DB.Where("name = ?", profile).First(&p)

	host := Host{
		IsEnabled: enable,
		Name:      name,
		FQDN:      fqdn,
		UUID:      uuid,
		ProfileID: p.ID,
	}

	c.DB.Create(&host)

	for _, i := range interfaces {
		inf := Interface{}

		c.DB.Where("mac_address = ?", i).Find(&inf)

		c.DB.Model(&host).Association("Interfaces").Append(&inf)
	}

	return nil
}

func (c *Config) RemoveHost(name string) error {
	c.DB.Where("name = ?", name).Delete(Host{})

	return nil
}

func (c *Config) ListHost(filter map[string]string) []driver.HostResult {
	var result []driver.HostResult

	sql := c.DB.Table("hosts").Select("hosts.id, hosts.is_enabled, hosts.name, hosts.fqdn, hosts.uuid, profiles.name").Joins("JOIN profiles ON profiles.id = hosts.profile_id")

	if v, ok := filter["host"]; ok {
		sql = sql.Where("hosts.name = ?", v)
	}
	if v, ok := filter["uuid"]; ok {
		sql = sql.Where("hosts.uuid = ?", v)
	}

	rows, _ := sql.Rows()
	defer rows.Close()

	for rows.Next() {
		var id uint
		var enabled bool
		var name string
		var fqdn string
		var uuid string
		var profile string

		rows.Scan(&id, &enabled, &name, &uuid, &profile)

		h := driver.HostResult{
			ID:         id,
			Enabled:    enabled,
			Name:       name,
			FQDN:       fqdn,
			UUID:       uuid,
			Interfaces: []string{},
			Profile:    profile,
		}

		result = append(result, h)
	}

	return result
}

func (c *Config) EnableHost(name string) error {
	host := Host{}

	c.DB.First(&host)
	host.IsEnabled = true

	c.DB.Save(&host)

	return nil
}

func (c *Config) DisableHost(name string) error {
	host := Host{}

	c.DB.First(&host)
	host.IsEnabled = false

	c.DB.Save(&host)

	return nil
}

func (c *Config) CountHost() int {
	var count int64

	c.DB.Model(&Host{}).Count(&count)

	return int(count)
}
