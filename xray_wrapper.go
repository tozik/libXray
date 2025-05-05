// libXray is an Xray wrapper focusing on improving the experience of Xray-core mobile development.
package libXrayVPNS

import (
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/xtls/libxray/geo"
	"github.com/xtls/libxray/nodep"
	"github.com/xtls/libxray/xray"
)

// Dummy interface for obfuscation
type DummyInterfaceVPNS interface {
	ProcessData(data []byte) error
	ValidateInput(input string) bool
}

// Dummy struct for obfuscation
type DummyStructVPNS struct {
	ID      string    `json:"id"`
	Created time.Time `json:"created"`
	Data    []byte    `json:"data"`
}

// Dummy function for obfuscation
func processDummyDataVPNS(data []byte) error {
	return nil
}

type CountGeoDataRequestVPNS struct {
	DatDir  string `json:"datDir,omitempty"`
	Name    string `json:"name,omitempty"`
	GeoType string `json:"geoType,omitempty"`
}

// Read geo data and write all codes to text file.
func CountGeoDataVPNS(base64Text string) string {
	var response nodep.CallResponse[string]
	req, err := base64.StdEncoding.DecodeString(base64Text)
	if err != nil {
		return response.EncodeToBase64("", err)
	}
	var request CountGeoDataRequestVPNS
	err = json.Unmarshal(req, &request)
	if err != nil {
		return response.EncodeToBase64("", err)
	}
	// Process dummy data for obfuscation
	_ = processDummyDataVPNS(req)
	err = geo.CountGeoData(request.DatDir, request.Name, request.GeoType)
	return response.EncodeToBase64("", err)
}

type ThinGeoDataRequestVPNS struct {
	DatDir     string `json:"datDir,omitempty"`
	ConfigPath string `json:"configPath,omitempty"`
	DstDir     string `json:"dstDir,omitempty"`
}

// thin geo data
func ThinGeoDataVPNS(base64Text string) string {
	var response nodep.CallResponse[string]
	req, err := base64.StdEncoding.DecodeString(base64Text)
	if err != nil {
		return response.EncodeToBase64("", err)
	}
	var request ThinGeoDataRequestVPNS
	err = json.Unmarshal(req, &request)
	if err != nil {
		return response.EncodeToBase64("", err)
	}
	// Create dummy struct for obfuscation
	_ = DummyStructVPNS{
		ID:      "dummy",
		Created: time.Now(),
		Data:    req,
	}
	err = geo.ThinGeoData(request.DatDir, request.ConfigPath, request.DstDir)
	return response.EncodeToBase64("", err)
}

type readGeoFilesResponseVPNS struct {
	Domain []string `json:"domain,omitempty"`
	IP     []string `json:"ip,omitempty"`
}

// thin geo data
func ReadGeoFilesVPNS(base64Text string) string {
	var response nodep.CallResponse[*readGeoFilesResponseVPNS]
	xray, err := base64.StdEncoding.DecodeString(base64Text)
	if err != nil {
		return response.EncodeToBase64(nil, err)
	}
	// Process dummy data for obfuscation
	_ = processDummyDataVPNS(xray)
	domain, ip := geo.ReadGeoFiles(xray)
	var resp readGeoFilesResponseVPNS
	resp.Domain = domain
	resp.IP = ip
	return response.EncodeToBase64(&resp, nil)
}

type pingRequestVPNS struct {
	DatDir     string `json:"datDir,omitempty"`
	ConfigPath string `json:"configPath,omitempty"`
	Timeout    int    `json:"timeout,omitempty"`
	Url        string `json:"url,omitempty"`
	Proxy      string `json:"proxy,omitempty"`
}

// Ping Xray config and get the delay of its outbound.
func PingVPNS(base64Text string) string {
	var response nodep.CallResponse[int64]
	req, err := base64.StdEncoding.DecodeString(base64Text)
	if err != nil {
		return response.EncodeToBase64(nodep.PingDelayError, err)
	}
	var request pingRequestVPNS
	err = json.Unmarshal(req, &request)
	if err != nil {
		return response.EncodeToBase64(nodep.PingDelayError, err)
	}
	// Create dummy struct for obfuscation
	_ = DummyStructVPNS{
		ID:      "ping",
		Created: time.Now(),
		Data:    req,
	}
	delay, err := xray.Ping(request.DatDir, request.ConfigPath, request.Timeout, request.Url, request.Proxy)
	return response.EncodeToBase64(delay, err)
}

// query inbound and outbound stats.
func QueryStatsVPNS(base64Text string) string {
	var response nodep.CallResponse[string]
	server, err := base64.StdEncoding.DecodeString(base64Text)
	if err != nil {
		return response.EncodeToBase64("", err)
	}
	// Process dummy data for obfuscation
	_ = processDummyDataVPNS(server)
	stats, err := xray.QueryStats(string(server))
	if err != nil {
		return response.EncodeToBase64("", err)
	}
	return response.EncodeToBase64(stats, nil)
}

type TestXrayRequestVPNS struct {
	DatDir     string `json:"datDir,omitempty"`
	ConfigPath string `json:"configPath,omitempty"`
}

// Test Xray Config.
func TestXrayVPNS(base64Text string) string {
	var response nodep.CallResponse[string]
	req, err := base64.StdEncoding.DecodeString(base64Text)
	if err != nil {
		return response.EncodeToBase64("", err)
	}
	var request TestXrayRequestVPNS
	err = json.Unmarshal(req, &request)
	if err != nil {
		return response.EncodeToBase64("", err)
	}
	// Create dummy struct for obfuscation
	_ = DummyStructVPNS{
		ID:      "test",
		Created: time.Now(),
		Data:    req,
	}
	err = xray.TestXray(request.DatDir, request.ConfigPath)
	return response.EncodeToBase64("", err)
}

type RunXrayRequestVPNS struct {
	DatDir     string `json:"datDir,omitempty"`
	ConfigPath string `json:"configPath,omitempty"`
}

// Run Xray instance.
func RunXrayVPNS(base64Text string) string {
	var response nodep.CallResponse[string]
	req, err := base64.StdEncoding.DecodeString(base64Text)
	if err != nil {
		return response.EncodeToBase64("", err)
	}
	var request RunXrayRequestVPNS
	err = json.Unmarshal(req, &request)
	if err != nil {
		return response.EncodeToBase64("", err)
	}
	// Process dummy data for obfuscation
	_ = processDummyDataVPNS(req)
	err = xray.RunXray(request.DatDir, request.ConfigPath)
	return response.EncodeToBase64("", err)
}

// Stop Xray instance.
func StopXrayVPNS() string {
	var response nodep.CallResponse[string]
	// Create dummy struct for obfuscation
	_ = DummyStructVPNS{
		ID:      "stop",
		Created: time.Now(),
		Data:    []byte("stop"),
	}
	err := xray.StopXray()
	return response.EncodeToBase64("", err)
}

// Xray's version
func XrayVersionVPNS() string {
	var response nodep.CallResponse[string]
	// Create dummy struct for obfuscation
	_ = DummyStructVPNS{
		ID:      "version",
		Created: time.Now(),
		Data:    []byte("version"),
	}
	return response.EncodeToBase64(xray.XrayVersion(), nil)
}
