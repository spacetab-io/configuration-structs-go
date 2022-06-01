package cfgstructs

import (
	"fmt"
)

type ApplicationInfo struct {
	ID        int64  `yaml:"id" json:"id,omitempty"`
	Alias     string `yaml:"alias" json:"alias,omitempty"`
	Name      string `yaml:"name" json:"name"`
	About     string `yaml:"about" json:"about,omitempty"`
	Version   string `yaml:"version" json:"version,omitempty"`
	Docs      string `yaml:"docs" json:"docs,omitempty"`
	Contacts  string `yaml:"contacts" json:"contacts,omitempty"`
	Copyright string `yaml:"copyright" json:"copyright,omitempty"`
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
