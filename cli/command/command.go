package command

import (
	"path"
)

var AppPath = "/var/lib/metadata"

var DBFilePath = path.Join(AppPath, "data.db")
