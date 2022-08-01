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
	ltcchaincfg "github.com/ltcsuite/ltcd/chaincfg"
	"github.com/ltcsuite/ltcutil"
	"github.com/humansdotai/humans/common/cosmos"
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

// ConvertToNewBCHAddressFormat convert the given BCH to new address format
func ConvertToNewBCHAddressFormat(addr Address) (Address, error) {
	if !addr.IsChain(BCHChain) {
		return NoAddress, fmt.Errorf("address(%s) is not BCH chain", addr)
	}
	network := GetCurrentChainNetwork()
	var param *bchchaincfg.Params
	switch network {
	case MockNet:
		param = &bchchaincfg.RegressionNetParams
	case TestNet:
		param = &bchchaincfg.TestNet3Params
	case MainNet:
		param = &bchchaincfg.MainNetParams
	case StageNet:
		param = &bchchaincfg.MainNetParams
	}
	bchAddr, err := bchutil.DecodeAddress(addr.String(), param)
	if err != nil {
		return NoAddress, fmt.Errorf("fail to decode address(%s), %w", addr, err)
	}
	return getBCHAddress(bchAddr, param)
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

// ConvertToNewBCHAddressFormatV83 convert the given BCH to new address format
func ConvertToNewBCHAddressFormatV83(addr Address) (Address, error) {
	if !addr.IsChain(BCHChain) {
		return NoAddress, fmt.Errorf("address(%s) is not BCH chain", addr)
	}
	network := GetCurrentChainNetwork()
	var param *bchchaincfg.Params
	switch network {
	case MockNet:
		param = &bchchaincfg.RegressionNetParams
	case TestNet:
		param = &bchchaincfg.TestNet3Params
	case MainNet:
		param = &bchchaincfg.MainNetParams
	case StageNet:
		param = &bchchaincfg.MainNetParams
	}
	bchAddr, err := bchutil.DecodeAddress(addr.String(), param)
	if err != nil {
		return NoAddress, fmt.Errorf("fail to decode address(%s), %w", addr, err)
	}
	return getBCHAddressV83(bchAddr, param)
}

func getBCHAddressV83(address bchutil.Address, cfg *bchchaincfg.Params) (Address, error) {
	switch address.(type) {
	case *bchutil.LegacyAddressPubKeyHash, *bchutil.AddressPubKeyHash:
		h, err := bchutil.NewAddressPubKeyHash(address.ScriptAddress(), cfg)
		if err != nil {
			return NoAddress, fmt.Errorf("fail to convert to new pubkey hash address: %w", err)
		}
		return NewAddress(h.String())
	case *bchutil.LegacyAddressScriptHash, *bchutil.AddressScriptHash:
		h, err := bchutil.NewAddressScriptHashFromHash(address.ScriptAddress(), cfg)
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
	case BNBChain:
		prefix, _, _ := bech32.Decode(addr.String())
		return prefix == "bnb" || prefix == "tbnb"
	case TERRAChain:
		// Note: Terra does not use a special prefix for testnet
		prefix, _, _ := bech32.Decode(addr.String())
		return prefix == "terra"
	case HumansChain:
		prefix, _, _ := bech32.Decode(addr.String())
		return prefix == "human" || prefix == "thuman" || prefix == "shuman"
	case BTCChain:
		prefix, _, err := bech32.Decode(addr.String())
		if err == nil && (prefix == "bc" || prefix == "tb") {
			return true
		}
		// Check mainnet other formats
		_, err = btcutil.DecodeAddress(addr.String(), &chaincfg.MainNetParams)
		if err == nil {
			return true
		}
		// Check testnet other formats
		_, err = btcutil.DecodeAddress(addr.String(), &chaincfg.TestNet3Params)
		if err == nil {
			return true
		}
		return false
	case LTCChain:
		prefix, _, err := bech32.Decode(addr.String())
		if err == nil && (prefix == "ltc" || prefix == "tltc" || prefix == "rltc") {
			return true
		}
		// Check mainnet other formats
		_, err = ltcutil.DecodeAddress(addr.String(), &ltcchaincfg.MainNetParams)
		if err == nil {
			return true
		}
		// Check testnet other formats
		_, err = ltcutil.DecodeAddress(addr.String(), &ltcchaincfg.TestNet4Params)
		if err == nil {
			return true
		}
		return false
	case BCHChain:
		// Check mainnet other formats
		_, err := bchutil.DecodeAddress(addr.String(), &bchchaincfg.MainNetParams)
		if err == nil {
			return true
		}
		// Check testnet other formats
		_, err = bchutil.DecodeAddress(addr.String(), &bchchaincfg.TestNet3Params)
		if err == nil {
			return true
		}
		// Check mocknet / regression other formats
		_, err = bchutil.DecodeAddress(addr.String(), &bchchaincfg.RegressionNetParams)
		if err == nil {
			return true
		}
		return false
	case DOGEChain:
		// Check mainnet other formats
		_, err := dogutil.DecodeAddress(addr.String(), &dogchaincfg.MainNetParams)
		if err == nil {
			return true
		}
		// Check testnet other formats
		_, err = dogutil.DecodeAddress(addr.String(), &dogchaincfg.TestNet3Params)
		if err == nil {
			return true
		}
		// Check mocknet / regression other formats
		_, err = dogutil.DecodeAddress(addr.String(), &dogchaincfg.RegressionNetParams)
		if err == nil {
			return true
		}
		return false
	default:
		return true // if HumansNode don't specifically check a chain yet, assume its ok.
	}
}

func (addr Address) GetChain() Chain {
	for _, chain := range []Chain{ETHChain, BNBChain, HumansChain, BTCChain, LTCChain, BCHChain, DOGEChain, TERRAChain} {
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
	case BNBChain:
		prefix, _, _ := bech32.Decode(addr.String())
		if strings.EqualFold(prefix, "bnb") {
			return mainNetPredicate()
		}
		if strings.EqualFold(prefix, "tbnb") {
			return TestNet
		}
	case TERRAChain:
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
	case BTCChain:
		prefix, _, _ := bech32.Decode(addr.String())
		switch prefix {
		case "bc":
			return mainNetPredicate()
		case "tb":
			return TestNet
		case "bcrt":
			return MockNet
		default:
			_, err := btcutil.DecodeAddress(addr.String(), &chaincfg.MainNetParams)
			if err == nil {
				return mainNetPredicate()
			}
			_, err = btcutil.DecodeAddress(addr.String(), &chaincfg.TestNet3Params)
			if err == nil {
				return TestNet
			}
			_, err = btcutil.DecodeAddress(addr.String(), &chaincfg.RegressionNetParams)
			if err == nil {
				return MockNet
			}
		}
	case LTCChain:
		prefix, _, _ := bech32.Decode(addr.String())
		switch prefix {
		case "ltc":
			return mainNetPredicate()
		case "tltc":
			return TestNet
		case "rltc":
			return MockNet
		default:
			_, err := ltcutil.DecodeAddress(addr.String(), &ltcchaincfg.MainNetParams)
			if err == nil {
				return mainNetPredicate()
			}
			_, err = ltcutil.DecodeAddress(addr.String(), &ltcchaincfg.TestNet4Params)
			if err == nil {
				return TestNet
			}
			_, err = ltcutil.DecodeAddress(addr.String(), &ltcchaincfg.RegressionNetParams)
			if err == nil {
				return MockNet
			}
		}
	case BCHChain:
		// Check mainnet other formats
		_, err := bchutil.DecodeAddress(addr.String(), &bchchaincfg.MainNetParams)
		if err == nil {
			return mainNetPredicate()
		}
		// Check testnet other formats
		_, err = bchutil.DecodeAddress(addr.String(), &bchchaincfg.TestNet3Params)
		if err == nil {
			return TestNet
		}
		// Check mocknet / regression other formats
		_, err = bchutil.DecodeAddress(addr.String(), &bchchaincfg.RegressionNetParams)
		if err == nil {
			return MockNet
		}
	case DOGEChain:
		// Check mainnet other formats
		_, err := dogutil.DecodeAddress(addr.String(), &dogchaincfg.MainNetParams)
		if err == nil {
			return mainNetPredicate()
		}
		// Check testnet other formats
		_, err = dogutil.DecodeAddress(addr.String(), &dogchaincfg.TestNet3Params)
		if err == nil {
			return TestNet
		}
		// Check mocknet / regression other formats
		_, err = dogutil.DecodeAddress(addr.String(), &dogchaincfg.RegressionNetParams)
		if err == nil {
			return MockNet
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
