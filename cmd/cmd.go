//go:build !testnet && !mocknet && !stagenet
// +build !testnet,!mocknet,!stagenet

package cmd

const (
	Bech32PrefixAccAddr            = "kima"
	Bech32PrefixAccPub             = "kimapub"
	Bech32PrefixValAddr            = "kimavaloper"
	Bech32PrefixValPub             = "kimavaloperpub"
	Bech32PrefixConsAddr           = "kimavalcons"
	Bech32PrefixConsPub            = "kimavalconspub"
	DenomRegex                     = `[a-zA-Z][a-zA-Z0-9:\\/\\\-\\_\\.]{2,127}`
	DIVERSIchainCoinType    uint32 = 931
	DIVERSIchainCoinPurpose uint32 = 44
	DIVERSIchainHDPath      string = `m/44'/931'/0'/0/0`
)
