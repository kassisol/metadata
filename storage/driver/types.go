package driver

type ServerConfigResult struct {
	Key   string
	Value string
}

type DataResult struct {
	ID          uint
	Name        string
	Type        string
	Value       string
	Description string
}

type IPResult struct {
	ID        uint
	IPAddress string
	Netmask   string
	Gateway   string
}

type InterfaceResult struct {
	Index      int
	MACAddress string
	IP         IPResult
	FloatingIP IPResult
}

type HostResult struct {
	ID         uint
	Enabled    bool
	Name       string
	FQDN       string
	UUID       string
	Interfaces []string
	Profile    string
}
