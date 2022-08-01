package common

import (
	"os"
	"strings"
)

// ChainNetwork is to indicate which chain environment HumanNode are working with
type ChainNetwork uint8

const (
	// TestNet network for test
	TestNet ChainNetwork = iota
	// MainNet network for main net
	MainNet
	// MockNet network for main net
	MockNet
	// Stagenet network for stage net
	StageNet
)

// GetCurrentChainNetwork determinate what kind of network currently it is working with
func GetCurrentChainNetwork() ChainNetwork {
	if strings.EqualFold(os.Getenv("NET"), "mocknet") {
		return MockNet
	}
	if strings.EqualFold(os.Getenv("NET"), "testnet") {
		return TestNet
	}
	if strings.EqualFold(os.Getenv("NET"), "stagenet") {
		return StageNet
	}
	return MainNet
}

// Soft Equals check is mainnet == mainet, or (testnet/mocknet == testnet/mocknet)
func (net ChainNetwork) SoftEquals(net2 ChainNetwork) bool {
	if net == MainNet && net2 == MainNet {
		return true
	}
	if net != MainNet && net2 != MainNet {
		return true
	}

	return false
}
