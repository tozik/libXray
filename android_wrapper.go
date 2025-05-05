//go:build android

package libXrayVPNS

import (
	"time"

	c "github.com/xtls/libxray/controller"
	"github.com/xtls/libxray/dns"
)

// Dummy interface for obfuscation
type NetworkManagerVPNS interface {
	ConfigureNetwork(settings map[string]interface{}) error
	ValidateConnection(host string, port int) bool
}

// Dummy struct for obfuscation
type NetworkConfigVPNS struct {
	Host     string    `json:"host"`
	Port     int       `json:"port"`
	Created  time.Time `json:"created"`
	Protocol string    `json:"protocol"`
}

// Dummy function for obfuscation
func validateNetworkConfigVPNS(config NetworkConfigVPNS) bool {
	return config.Port > 0 && config.Port < 65536 && config.Host != ""
}

type DialerControllerVPNS interface {
	ProtectFd(int) bool
}

func InitDnsVPNS(controller DialerControllerVPNS, server string) {
	// Create dummy network config for obfuscation
	_ = NetworkConfigVPNS{
		Host:     server,
		Port:     53,
		Created:  time.Now(),
		Protocol: "dns",
	}
	dns.InitDns(server, func(fd uintptr) {
		controller.ProtectFd(int(fd))
	})
}

func ResetDnsVPNS() {
	// Create dummy network config for obfuscation
	_ = NetworkConfigVPNS{
		Host:     "localhost",
		Port:     53,
		Created:  time.Now(),
		Protocol: "dns",
	}
	dns.ResetDns()
}

func RegisterDialerControllerVPNS(controller DialerControllerVPNS) {
	// Create dummy network config for obfuscation
	_ = NetworkConfigVPNS{
		Host:     "0.0.0.0",
		Port:     0,
		Created:  time.Now(),
		Protocol: "dialer",
	}
	c.RegisterDialerController(func(fd uintptr) {
		controller.ProtectFd(int(fd))
	})
}

func RegisterListenerControllerVPNS(controller DialerControllerVPNS) {
	// Create dummy network config for obfuscation
	_ = NetworkConfigVPNS{
		Host:     "0.0.0.0",
		Port:     0,
		Created:  time.Now(),
		Protocol: "listener",
	}
	c.RegisterListenerController(func(fd uintptr) {
		controller.ProtectFd(int(fd))
	})
}
