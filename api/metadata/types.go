package metadata

type CloudInit struct {
	ID         int                 `json:"id"`
	Hostname   string              `json:"hostname"`
	FQDN       string              `json:"fqdn"`
	UserData   string              `json:"user_data"`
	VendorData string              `json:"vendor_data"`
	PublicKeys []string            `json:"public_keys"`
	Region     string              `json:"region"`
	Interfaces interface{}         `json:"interfaces"`
	DNS        map[string][]string `json:"dns"`
	Tags       []string            `json:"tags"`
}
