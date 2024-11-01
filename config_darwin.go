package swifttunnel

import (
	"github.com/XenonCommunity/swifttunnel/swiftypes"
	"net"
)

type DriverType int

const (
	DriverTypeTunTapOSX DriverType = iota
	DriverTypeSystem
)

type Config struct {
	AdapterName string
	AdapterType swiftypes.AdapterType
	DriverType  DriverType

	MTU       int
	UnicastIP net.IPNet
}

func NewDefaultConfig() Config {
	return Config{
		AdapterName: "SwiftTunnel VPN",
		AdapterType: swiftypes.AdapterTypeTUN,
		MTU:         1500,
		UnicastIP: net.IPNet{
			IP:   net.IPv4(10, 18, 21, 1),
			Mask: net.IPv4Mask(255, 255, 255, 0),
		},
	}
}