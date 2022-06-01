package cfgstructs

import (
	"time"

	"github.com/spacetab-io/configuration-structs-go/v2/contracts"
)

type AuthConfig struct {
	Enable bool               `yaml:"enable"`
	Type   contracts.AuthType `yaml:"type"`
	Login  string             `yaml:"login"`
	Secret string             `yaml:"secret"`
}

type RESTService struct {
	Enable         bool          `yaml:"enable"`
	BaseURL        string        `yaml:"base_url"`
	Timeout        time.Duration `yaml:"timeout"`
	GzipContent    bool          `yaml:"gzip_content"`
	DebugEnable    bool          `yaml:"debug"`
	Authentication AuthConfig    `yaml:"authentication"`
}
