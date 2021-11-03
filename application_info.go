package cfgstructs

import (
	"fmt"
)

type ApplicationInfo struct {
	ID        int64  `json:"id,omitempty"`
	Alias     string `json:"alias"`
	Name      string `json:"name"`
	About     string `json:"about"`
	Version   string `json:"version"`
	Docs      string `json:"docs"`
	Contacts  string `json:"contacts"`
	Copyright string `json:"copyright"`
}

func (i ApplicationInfo) GetString() string {
	return fmt.Sprintf("%s %s [%s] %s", i.Name, i.Version, i.Alias, i.Copyright)
}

func (i ApplicationInfo) GetAlias() string {
	return i.Alias
}

func (i ApplicationInfo) GetVersion() string {
	return i.Version
}

func (i ApplicationInfo) Summary() string {
	return fmt.Sprintf("%s (%s) %s", i.Name, i.Version, i.Copyright)
}
