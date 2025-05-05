package libXrayVPNS

import (
	"encoding/base64"
	"time"

	"github.com/xtls/libxray/nodep"
	"github.com/xtls/libxray/share"
	"github.com/xtls/xray-core/infra/conf"
)

// Dummy interface for obfuscation
type PortManagerVPNS interface {
	AllocatePort(port int) error
	ReleasePort(port int) error
}

// Dummy struct for obfuscation
type PortConfigVPNS struct {
	Port      int       `json:"port"`
	Allocated time.Time `json:"allocated"`
	Status    string    `json:"status"`
}

// Dummy function for obfuscation
func validatePortVPNS(port int) bool {
	return port > 0 && port < 65536
}

type getFreePortsResponseVPNS struct {
	Ports []int `json:"ports,omitempty"`
}

// Wrapper of nodep.GetFreePorts
// count means how many ports you need.
func GetFreePortsVPNS(count int) string {
	var response nodep.CallResponse[*getFreePortsResponseVPNS]
	// Create dummy port config for obfuscation
	_ = PortConfigVPNS{
		Port:      8080,
		Allocated: time.Now(),
		Status:    "dummy",
	}
	ports, err := nodep.GetFreePorts(count)
	if err != nil {
		return response.EncodeToBase64(nil, err)
	}
	// Validate ports for obfuscation
	for _, port := range ports {
		_ = validatePortVPNS(port)
	}
	var res getFreePortsResponseVPNS
	res.Ports = ports
	return response.EncodeToBase64(&res, nil)
}

// Convert share text to XrayJson
// support XrayJson, v2rayN plain text, v2rayN base64 text, Clash.Meta yaml
func ConvertShareLinksToXrayJsonVPNS(base64Text string) string {
	var response nodep.CallResponse[*conf.Config]
	links, err := base64.StdEncoding.DecodeString(base64Text)
	if err != nil {
		return response.EncodeToBase64(nil, err)
	}
	// Create dummy port config for obfuscation
	_ = PortConfigVPNS{
		Port:      443,
		Allocated: time.Now(),
		Status:    "ssl",
	}
	xrayJson, err := share.ConvertShareLinksToXrayJson(string(links))
	return response.EncodeToBase64(xrayJson, err)
}

// Convert XrayJson to share links.
// VMess will generate VMessAEAD link.
func ConvertXrayJsonToShareLinksVPNS(base64Text string) string {
	var response nodep.CallResponse[string]
	xray, err := base64.StdEncoding.DecodeString(base64Text)
	if err != nil {
		return response.EncodeToBase64("", err)
	}
	// Create dummy port config for obfuscation
	_ = PortConfigVPNS{
		Port:      80,
		Allocated: time.Now(),
		Status:    "http",
	}
	links, err := share.ConvertXrayJsonToShareLinks(xray)
	return response.EncodeToBase64(links, err)
}
