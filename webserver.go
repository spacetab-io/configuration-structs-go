package cfgstructs

import (
	"fmt"
	"time"

	"github.com/spacetab-io/configuration-structs-go/v2/contracts"
)

type CORSCfg struct {
	Enabled          bool          `yaml:"enabled"`
	AllowAllOrigins  bool          `yaml:"allow_all_origins"`
	AllowCredentials bool          `yaml:"allow_credentials"`
	Lifetime         time.Duration `yaml:"lifetime"` // In hours! Ensure to enter in config value as string in format <val>h
	AllowMethods     []string      `yaml:"allow_methods"`
	AllowHeaders     []string      `yaml:"allow_headers"`
}

type HTTPTimeouts struct {
	ReadHeader    time.Duration `yaml:"read_header" valid:"required"`
	ReadRequest   time.Duration `yaml:"read_request" valid:"required"`
	WriteResponse time.Duration `yaml:"write_response" valid:"required"`
	Idle          time.Duration `yaml:"idle" valid:"required"`
	Shutdown      time.Duration `yaml:"shutdown" valid:"required"`
}

type RACCfg struct {
	MaxConnsPerIP      int `yaml:"max_conn_per_ip" valid:"required"`
	MaxRequestsPerConn int `yaml:"max_req_per_conn" valid:"required"`
}

type WebServer struct {
	Host                   string       `yaml:"host" valid:"required"`
	Port                   int          `yaml:"port" valid:"required"`
	CORS                   CORSCfg      `yaml:"cors" valid:"required"`
	Timeouts               HTTPTimeouts `yaml:"timeouts" valid:"required"`
	RequestsAndConnections RACCfg       `yaml:"requests_and_connections" valid:"required"`
	Compress               bool         `yaml:"compress"`
	Debug                  bool         `yaml:"debug"`
}

func (w WebServer) String() string {
	return "web server config"
}

func (w WebServer) Validate() (bool, error) {
	return contracts.ConfigValidate(w)
}

func (w WebServer) GetReadRequestTimeout() time.Duration {
	return w.Timeouts.ReadRequest
}

func (w WebServer) GetWriteResponseTimeout() time.Duration {
	return w.Timeouts.WriteResponse
}

func (w WebServer) GetIdleTimeout() time.Duration {
	return w.Timeouts.Idle
}

func (w WebServer) GetShutdownTimeout() time.Duration {
	return w.Timeouts.Shutdown
}

func (w WebServer) GetMaxConnsPerIP() int {
	return w.RequestsAndConnections.MaxConnsPerIP
}

func (w WebServer) GetMaxRequestsPerConn() int {
	return w.RequestsAndConnections.MaxRequestsPerConn
}

func (w WebServer) UseCompression() bool {
	return w.Compress
}

func (w WebServer) CORSEnabled() bool {
	return w.CORS.Enabled
}

func (w WebServer) GetListenAddress() string {
	return fmt.Sprintf("%v:%d", w.Host, w.Port)
}
