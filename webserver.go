package cfgstructs

import (
	"fmt"
	"time"
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
	ReadHeader    time.Duration `yaml:"read_header"`    // In seconds! Ensure to enter in config value as string in format <val>s
	ReadRequest   time.Duration `yaml:"read_request"`   // In seconds! Ensure to enter in config value as string in format <val>s
	WriteResponse time.Duration `yaml:"write_response"` // In seconds! Ensure to enter in config value as string in format <val>s
	Idle          time.Duration `yaml:"idle"`           // In seconds! Ensure to enter in config value as string in format <val>s
	Shutdown      time.Duration `yaml:"shutdown"`       // In seconds! Ensure to enter in config value as string in format <val>s
}

type RACCfg struct {
	MaxConnsPerIP      int `yaml:"max_conn_per_ip"`
	MaxRequestsPerConn int `yaml:"max_req_per_conn"`
}

type WebServer struct {
	Host                   string       `yaml:"host"`
	Port                   int          `yaml:"port"`
	CORS                   CORSCfg      `yaml:"cors"`
	Timeouts               HTTPTimeouts `yaml:"timeouts"`
	RequestsAndConnections RACCfg       `yaml:"requests_and_connections"`
	Compress               bool         `yaml:"compress"`
	Debug                  bool         `yaml:"debug"`
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
