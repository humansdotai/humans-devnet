package common

import (
	"fmt"
	"strings"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/bech32"
	dogchaincfg "github.com/eager7/dogd/chaincfg"
	"github.com/eager7/dogutil"
	eth "github.com/ethereum/go-ethereum/common"
	bchchaincfg "github.com/gcash/bchd/chaincfg"
	"github.com/gcash/bchutil"
	"github.com/humansdotai/humans/common/cosmos"
	ltcchaincfg "github.com/ltcsuite/ltcd/chaincfg"
	"github.com/ltcsuite/ltcutil"
)

type Address string

var NoAddress = Address("")

const ETHAddressLen = 42

// NewAddress create a new Address. Supports Binance, Bitcoin, and Ethereum
func NewAddress(address string) (Address, error) {
	if len(address) == 0 {
		return NoAddress, nil
	}

	// Check is eth address
	if eth.IsHexAddress(address) {
		return Address(address), nil
	}

	// Check bech32 addresses, would succeed any string bech32 encoded (e.g. BNB, TERRA, ATOM)
	_, _, err := bech32.Decode(address)
	if err == nil {
		return Address(address), nil
	}

	// Check other BTC address formats with mainnet
	_, err = btcutil.DecodeAddress(address, &chaincfg.MainNetParams)
	if err == nil {
		return Address(address), nil
	}

	// Check BTC address formats with testnet
	_, err = btcutil.DecodeAddress(address, &chaincfg.TestNet3Params)
	if err == nil {
		return Address(address), nil
	}

	// Check other LTC address formats with mainnet
	_, err = ltcutil.DecodeAddress(address, &ltcchaincfg.MainNetParams)
	if err == nil {
		return Address(address), nil
	}

	// Check LTC address formats with testnet
	_, err = ltcutil.DecodeAddress(address, &ltcchaincfg.TestNet4Params)
	if err == nil {
		return Address(address), nil
	}

	// Check BCH address formats with mainnet
	_, err = bchutil.DecodeAddress(address, &bchchaincfg.MainNetParams)
	if err == nil {
		return Address(address), nil
	}

	// Check BCH address formats with testnet
	_, err = bchutil.DecodeAddress(address, &bchchaincfg.TestNet3Params)
	if err == nil {
		return Address(address), nil
	}

	// Check BCH address formats with mocknet
	_, err = bchutil.DecodeAddress(address, &bchchaincfg.RegressionNetParams)
	if err == nil {
		return Address(address), nil
	}

	// Check DOGE address formats with mainnet
	_, err = dogutil.DecodeAddress(address, &dogchaincfg.MainNetParams)
	if err == nil {
		return Address(address), nil
	}

	// Check DOGE address formats with testnet
	_, err = dogutil.DecodeAddress(address, &dogchaincfg.TestNet3Params)
	if err == nil {
		return Address(address), nil
	}

	// Check DOGE address formats with mocknet
	_, err = dogutil.DecodeAddress(address, &dogchaincfg.RegressionNetParams)
	if err == nil {
		return Address(address), nil
	}

	return NoAddress, fmt.Errorf("address format not supported: %s", address)
}

// IsValidBCHAddress determinate whether the address is a valid new BCH address format
func (addr Address) IsValidBCHAddress() bool {
	// Check mainnet other formats
	bchAddr, err := bchutil.DecodeAddress(addr.String(), &bchchaincfg.MainNetParams)
	if err == nil {
		switch bchAddr.(type) {
		case *bchutil.LegacyAddressPubKeyHash, *bchutil.LegacyAddressScriptHash:
			return false
		}
		return true
	}
	bchAddr, err = bchutil.DecodeAddress(addr.String(), &bchchaincfg.TestNet3Params)
	if err == nil {
		switch bchAddr.(type) {
		case *bchutil.LegacyAddressPubKeyHash, *bchutil.LegacyAddressScriptHash:
			return false
		}
		return true
	}
	bchAddr, err = bchutil.DecodeAddress(addr.String(), &bchchaincfg.RegressionNetParams)
	if err == nil {
		switch bchAddr.(type) {
		case *bchutil.LegacyAddressPubKeyHash, *bchutil.LegacyAddressScriptHash:
			return false
		}
		return true
	}
	return false
}

func getBCHAddress(address bchutil.Address, cfg *bchchaincfg.Params) (Address, error) {
	switch address.(type) {
	case *bchutil.LegacyAddressPubKeyHash, *bchutil.AddressPubKeyHash:
		h, err := bchutil.NewAddressPubKeyHash(address.ScriptAddress(), cfg)
		if err != nil {
			return NoAddress, fmt.Errorf("fail to convert to new pubkey hash address: %w", err)
		}
		return NewAddress(h.String())
	case *bchutil.LegacyAddressScriptHash, *bchutil.AddressScriptHash:
		h, err := bchutil.NewAddressScriptHash(address.ScriptAddress(), cfg)
		if err != nil {
			return NoAddress, fmt.Errorf("fail to convert to new address script hash address: %w", err)
		}
		return NewAddress(h.String())
	}
	return NoAddress, fmt.Errorf("invalid address type")
}

func (addr Address) IsChain(chain Chain) bool {
	switch chain {
	case ETHChain:
		return strings.HasPrefix(addr.String(), "0x")
	case HumansChain:
		prefix, _, _ := bech32.Decode(addr.String())
		return prefix == "human" || prefix == "thuman" || prefix == "shuman"
	default:
		return true // if HumansNode don't specifically check a chain yet, assume its ok.
	}
}

func (addr Address) GetChain() Chain {
	for _, chain := range []Chain{ETHChain, HumansChain} {
		if addr.IsChain(chain) {
			return chain
		}
	}
	return EmptyChain
}

func (addr Address) GetNetwork(chain Chain) ChainNetwork {
	currentNetwork := GetCurrentChainNetwork()
	mainNetPredicate := func() ChainNetwork {
		if currentNetwork == StageNet {
			return StageNet
		}
		return MainNet
	}
	switch chain {
	case ETHChain:
		return currentNetwork
	case HumansChain:
		prefix, _, _ := bech32.Decode(addr.String())
		if strings.EqualFold(prefix, "human") {
			return mainNetPredicate()
		}
		if strings.EqualFold(prefix, "thuman") {
			return TestNet
		}
		if strings.EqualFold(prefix, "shuman") {
			return StageNet
		}
	}
	return MockNet
}

func (addr Address) AccAddress() (cosmos.AccAddress, error) {
	return cosmos.AccAddressFromBech32(addr.String())
}

func (addr Address) Equals(addr2 Address) bool {
	return strings.EqualFold(addr.String(), addr2.String())
}

func (addr Address) IsEmpty() bool {
	return strings.TrimSpace(addr.String()) == ""
}

func (addr Address) String() string {
	return string(addr)
}
