package contracts

import (
	"fmt"
	"strings"
)

type HostCfg struct {
	Host string `yaml:"host" valid:"required,host"`
	Port uint   `yaml:"port" valid:"required,port"`
}

// nolint: gomnd
func (hc HostCfg) String() string {
	h := strings.Split(hc.Host, ":")

	switch h[0] {
	case "http":
		if hc.Port == 80 {
			return hc.Host
		}
	case "https":
		if hc.Port == 443 {
			return hc.Host
		}
	case "ssl":
		if hc.Port == 21 {
			return hc.Host
		}
	}

	if hc.Port == 0 {
		return hc.Host
	}

	return fmt.Sprintf("%s:%d", hc.Host, hc.Port)
}

func (hc *HostCfg) GetHost() string {
	return hc.Host
}

func (hc *HostCfg) GetPort() uint {
	return hc.Port
}

func (hc *HostCfg) IsEmpty() bool {
	return hc == nil || hc.Host == ""
}

type HostsCfg []HostCfg

func (c HostsCfg) GetHosts() []string {
	hosts := make([]string, 0, len(c))

	for _, host := range c {
		hosts = append(hosts, host.GetHost())
	}

	return hosts
}

func (c HostsCfg) GetPorts() []uint {
	ports := make([]uint, 0, len(c))

	for _, host := range c {
		ports = append(ports, host.GetPort())
	}

	return ports
}

func (c HostsCfg) GetHostPortPairs() map[string]string {
	hp := make(map[string]string)

	for i, host := range c {
		hp[fmt.Sprint(i)] = host.String()
	}

	return hp
}
