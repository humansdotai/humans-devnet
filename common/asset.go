package common

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/jsonpb"
)

var (
	// EmptyAsset empty asset, not valid
	EmptyAsset = Asset{Chain: EmptyChain, Symbol: "", Ticker: "", Synth: false}
	// ETHAsset ETH
	ETHAsset = Asset{Chain: ETHChain, Symbol: "ETH", Ticker: "ETH", Synth: false}
	// HEART on humans
	HeartNative = Asset{Chain: HumansChain, Symbol: "HEART", Ticker: "HEART", Synth: false}
)

// NewAsset parse the given input into Asset object
func NewAsset(input string) (Asset, error) {
	var err error
	var asset Asset
	var sym string
	var parts []string
	if strings.Count(input, "/") > 0 {
		parts = strings.SplitN(input, "/", 2)
		asset.Synth = true
	} else {
		parts = strings.SplitN(input, ".", 2)
		asset.Synth = false
	}
	if len(parts) == 1 {
		asset.Chain = HumansChain
		sym = parts[0]
	} else {
		asset.Chain, err = NewChain(parts[0])
		if err != nil {
			return EmptyAsset, err
		}
		sym = parts[1]
	}

	asset.Symbol, err = NewSymbol(sym)
	if err != nil {
		return EmptyAsset, err
	}

	parts = strings.SplitN(sym, "-", 2)
	asset.Ticker, err = NewTicker(parts[0])
	if err != nil {
		return EmptyAsset, err
	}

	return asset, nil
}

// Equals determinate whether two assets are equivalent
func (a Asset) Equals(a2 Asset) bool {
	return a.Chain.Equals(a2.Chain) && a.Symbol.Equals(a2.Symbol) && a.Ticker.Equals(a2.Ticker) && a.Synth == a2.Synth
}

func (a Asset) GetChain() Chain {
	if a.Synth {
		return HumansChain
	}
	return a.Chain
}

// Get layer1 asset version
func (a Asset) GetLayer1Asset() Asset {
	if !a.IsSyntheticAsset() {
		return a
	}
	return Asset{
		Chain:  a.Chain,
		Symbol: a.Symbol,
		Ticker: a.Ticker,
		Synth:  false,
	}
}

// Get synthetic asset of asset
func (a Asset) GetSyntheticAsset() Asset {
	if a.IsSyntheticAsset() {
		return a
	}
	return Asset{
		Chain:  a.Chain,
		Symbol: a.Symbol,
		Ticker: a.Ticker,
		Synth:  true,
	}
}

// Check if asset is a pegged asset
func (a Asset) IsSyntheticAsset() bool {
	return a.Synth
}

// Native return native asset, only relevant on HumansChain
func (a Asset) Native() string {
	if a.IsHeart() {
		return "heart"
	}
	return strings.ToLower(a.String())
}

// IsEmpty will be true when any of the field is empty, chain,symbol or ticker
func (a Asset) IsEmpty() bool {
	return a.Chain.IsEmpty() || a.Symbol.IsEmpty() || a.Ticker.IsEmpty()
}

// String implement fmt.Stringer , return the string representation of Asset
func (a Asset) String() string {
	div := "."
	if a.Synth {
		div = "/"
	}
	return fmt.Sprintf("%s%s%s", a.Chain.String(), div, a.Symbol.String())
}

// IsGasAsset check whether asset is base asset used to pay for gas
func (a Asset) IsGasAsset() bool {
	gasAsset := a.GetChain().GetGasAsset()
	if gasAsset.IsEmpty() {
		return false
	}
	return a.Equals(gasAsset)
}

// IsRune is a helper function ,return true only when the asset represent RUNE
func (a Asset) IsHeart() bool {
	return a.Equals(HeartNative)
}

// IsNativeRune is a helper function, return true only when the asset represent NATIVE RUNE
func (a Asset) IsNativeHeart() bool {
	return a.IsHeart() && a.Chain.IsHumanChain()
}

// MarshalJSON implement Marshaler interface
func (a Asset) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.String())
}

// UnmarshalJSON implement Unmarshaler interface
func (a *Asset) UnmarshalJSON(data []byte) error {
	var err error
	var assetStr string
	if err := json.Unmarshal(data, &assetStr); err != nil {
		return err
	}
	*a, err = NewAsset(assetStr)
	return err
}

// MarshalJSONPB implement jsonpb.Marshaler
func (a Asset) MarshalJSONPB(*jsonpb.Marshaler) ([]byte, error) {
	return a.MarshalJSON()
}

// UnmarshalJSONPB implement jsonpb.Unmarshaler
func (a *Asset) UnmarshalJSONPB(unmarshal *jsonpb.Unmarshaler, content []byte) error {
	return a.UnmarshalJSON(content)
}

// RuneAsset return RUNE Asset depends on different environment
func RuneAsset() Asset {
	return HeartNative
}
