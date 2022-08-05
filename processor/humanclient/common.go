package humanclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/humansdotai/humans/app"
	"github.com/humansdotai/humans/common"
	"github.com/humansdotai/humans/processor/config"
	"github.com/humansdotai/humans/processor/humanclient/cosmos"
	stypes "github.com/humansdotai/humans/x/humans/types"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
)

type (
	// TxID is a string that can uniquely represent a transaction on different
	// block chain
	TxID string
	// TxIDs is a slice of TxID
	TxIDs []TxID
)

// BlankTxID represent blank
var BlankTxID = TxID("0000000000000000000000000000000000000000000000000000000000000000")

// NewTxID parse the input hash as TxID
func NewTxID(hash string) (TxID, error) {
	switch len(hash) {
	case 64:
		// do nothing
	case 66: // ETH check
		if !strings.HasPrefix(hash, "0x") {
			err := fmt.Errorf("txid error: must be 66 characters (got %d)", len(hash))
			return TxID(""), err
		}
	default:
		err := fmt.Errorf("txid error: must be 64 characters (got %d)", len(hash))
		return TxID(""), err
	}

	return TxID(strings.ToUpper(hash)), nil
}

// PubKeyContractAddressPair is an entry to map pubkey and contract addresses
type PubKeyContractAddressPair struct {
	PubKey    common.PubKey
	Contracts map[common.Chain]common.Address
}

type BridgeConfig struct {
	ChainId         string
	ChainHost       string
	ChainRPC        string
	ChainHomeFolder string
	TSS             config.TSSConfiguration
}

// HumanChainBridge will be used to send tx to DIVERSIChain
type HumanChainBridge struct {
	keys          *Keys
	cfg           BridgeConfig
	blockHeight   uint64
	accountNumber uint64
	seqNumber     uint64
	httpClient    *retryablehttp.Client
	broadcastLock *sync.RWMutex

	signerName                  string
	lastBlockHeightCheck        time.Time
	lastDiversichainBlockHeight uint64
	pubKey                      string
	voterAddress                string
}

// NewHumanChainBridge create a new instance of HumanChainBridge
func NewHumanChainBridge(k *Keys, cfg *BridgeConfig, signer string, pubKey string, voter string) (*HumanChainBridge, error) {
	httpClient := retryablehttp.NewClient()
	httpClient.Logger = nil

	return &HumanChainBridge{
		keys:          k,
		httpClient:    httpClient,
		signerName:    signer,
		broadcastLock: &sync.RWMutex{},
		pubKey:        pubKey,
		voterAddress:  voter,
		cfg:           *cfg,
	}, nil
}

// GetContext return a valid context with all relevant values set
func (b *HumanChainBridge) GetContext() client.Context {
	ctx := client.Context{}
	ctx = ctx.WithKeyring(b.keys.GetKeybase())
	ctx = ctx.WithChainID(b.cfg.ChainId)
	ctx = ctx.WithHomeDir(b.cfg.ChainHomeFolder)
	ctx = ctx.WithFromName(b.signerName)
	ctx = ctx.WithFromAddress(b.keys.GetSignerInfo().GetAddress())
	ctx = ctx.WithBroadcastMode("sync")

	encodingConfig := app.MakeEncodingConfig()
	ctx = ctx.WithCodec(encodingConfig.Marshaler)
	ctx = ctx.WithInterfaceRegistry(encodingConfig.InterfaceRegistry)
	ctx = ctx.WithTxConfig(encodingConfig.TxConfig)
	ctx = ctx.WithLegacyAmino(encodingConfig.Amino)
	ctx = ctx.WithAccountRetriever(authtypes.AccountRetriever{})

	remote := b.cfg.ChainRPC
	if !strings.HasSuffix(b.cfg.ChainHost, "http") {
		remote = fmt.Sprintf("tcp://%s", remote)
	}
	ctx = ctx.WithNodeURI(remote)
	client, err := rpchttp.New(remote, "/websocket")
	if err != nil {
		panic(err)
	}
	ctx = ctx.WithClient(client)
	return ctx
}

func (b *HumanChainBridge) getWithPath(path string) ([]byte, int, error) {
	return b.get(b.getHumanChainURL(path))
}

// getThorChainURL with the given path
func (b *HumanChainBridge) getHumanChainURL(path string) string {
	uri := url.URL{
		Scheme: "http",
		Host:   b.cfg.ChainHost,
		Path:   path,
	}
	return uri.String()
}

// get handle all the low level http GET calls using retryablehttp.ThorchainBridge
func (b *HumanChainBridge) get(url string) ([]byte, int, error) {
	resp, err := b.httpClient.Get(url)
	if err != nil {
		fmt.Println("ffailed to GET from humanschain")
		return nil, http.StatusNotFound, fmt.Errorf("failed to GET from humanschain: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Println("failed to close response body")
		}
	}()

	buf, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return buf, resp.StatusCode, errors.New("Status code: " + resp.Status + " returned")
	}
	if err != nil {
		fmt.Println("fail_read_humanshain_resp", "")
		return nil, resp.StatusCode, fmt.Errorf("failed to read response body: %w", err)
	}
	return buf, resp.StatusCode, nil
}

//
func (b *HumanChainBridge) GetVoterInfo() (string, string) {
	return b.pubKey, b.voterAddress
}

//
func (b *HumanChainBridge) GetMonikerName() string {
	return b.signerName
}

// getAccountNumberAndSequenceNumber returns account and Sequence number required to post into thorchain
func (b *HumanChainBridge) getAccountNumberAndSequenceNumber() (uint64, uint64, error) {
	path := fmt.Sprintf("%s/%s", "/cosmos/auth/v1beta1/accounts", b.keys.GetSignerInfo().GetAddress())

	body, _, err := b.getWithPath(path)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to get auth accounts: %w", err)
	}

	var resp AccountResp
	if err := json.Unmarshal(body, &resp); err != nil {
		return 0, 0, fmt.Errorf("failed to unmarshal account resp: %w", err)
	}

	acc := resp.Account
	return acc.AccountNumber, acc.Sequence, nil
}

// HumansBlockTime Block time of HUMANSChain
var HumansBlockTime = 5 * time.Second

// GetBlockHeight returns the current height for humans blocks
func (b *HumanChainBridge) GetBlockHeight() (uint64, error) {
	if time.Since(b.lastBlockHeightCheck) < HumansBlockTime && b.lastDiversichainBlockHeight > 0 {
		return b.lastDiversichainBlockHeight, nil
	}

	latestBlock, err := b.GetLastBlock("")
	if err != nil {
		return 0, fmt.Errorf("failed to GetDiverchainHeight: %w", err)
	}

	b.lastBlockHeightCheck = time.Now()
	h, _ := strconv.ParseUint(latestBlock.Block.Header.Height, 10, 64)

	return h, nil
}

// getLastBlock calls the /lastblock/{chain} endpoint and Unmarshal's into the QueryResLastBlockHeights type
func (b *HumanChainBridge) GetLastBlock(chain string) (QueryResLastBlockHeights, error) {
	path := "/cosmos/base/tendermint/v1beta1/blocks/latest"
	if chain != "" {
		path = fmt.Sprintf("%s/%s", path, chain)
	}
	buf, _, err := b.getWithPath(path)
	if err != nil {
		return QueryResLastBlockHeights{}, fmt.Errorf("failed to get lastblock: %w", err)
	}

	var lastBlock QueryResLastBlockHeights
	if err := json.Unmarshal(buf, &lastBlock); err != nil {
		return QueryResLastBlockHeights{}, fmt.Errorf("failed to unmarshal last block: %w", err)
	}
	return lastBlock, nil
}

func (b *HumanChainBridge) GetBalance(addr string) (AccountBalance, error) {
	path := "/cosmos/bank/v1beta1/balances/" + addr
	buf, _, err := b.getWithPath(path)
	if err != nil {
		return AccountBalance{}, fmt.Errorf("failed to get account balance: %w", err)
	}

	var accBalance AccountBalance
	if err := json.Unmarshal(buf, &accBalance); err != nil {
		return AccountBalance{}, fmt.Errorf("failed to unmarshal account balance: %w", err)
	}
	return accBalance, nil
}

// Get Transaction Data List
func (b *HumanChainBridge) GetTxDataList(chain string) (QueryTransactionDataList, error) {
	path := "/VigorousDeveloper/humans/humans/transaction_data"
	if chain != "" {
		path = fmt.Sprintf("%s/%s", path, chain)
	}
	buf, _, err := b.getWithPath(path)
	if err != nil {
		return QueryTransactionDataList{}, fmt.Errorf("failed to get lastblock: %w", err)
	}

	var txDataList QueryTransactionDataList
	if err := json.Unmarshal(buf, &txDataList); err != nil {
		return QueryTransactionDataList{}, fmt.Errorf("failed to unmarshal last block: %w", err)
	}
	return txDataList, nil
}

// Endpoint urls
const (
	PubKeysEndpoint = "/humanchain/vaults/pubkeys"
	AsgardVault     = "/thorchain/vaults/asgard"
)

type ChainContract struct {
	Chain  string `protobuf:"bytes,1,opt,name=chain,proto3,casttype=gitlab.com/thorchain/thornode/common.Chain" json:"chain,omitempty"`
	Router string `protobuf:"bytes,2,opt,name=router,proto3,casttype=gitlab.com/thorchain/thornode/common.Address" json:"router,omitempty"`
}

// QueryVaultPubKeyContract is a type to combine PubKey and it's related contract
type QueryVaultPubKeyContract struct {
	PubKey  common.PubKey   `json:"pub_key"`
	Routers []ChainContract `json:"routers"`
}

// QueryVaultsPubKeys represent the result for query vaults pubkeys
type QueryVaultsPubKeys struct {
	Asgard    []QueryVaultPubKeyContract `json:"asgard"`
	Yggdrasil []QueryVaultPubKeyContract `json:"yggdrasil"`
}

// GetPubKeys retrieve asgard vaults and yggdrasil vaults , and it's relevant smart contracts
func (b *HumanChainBridge) GetPubKeys() ([]PubKeyContractAddressPair, error) {
	buf, s, err := b.getWithPath(PubKeysEndpoint)
	if err != nil {
		return nil, fmt.Errorf("fail to get asgard vaults: %w", err)
	}
	if s != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d", s)
	}
	var result QueryVaultsPubKeys
	if err := json.Unmarshal(buf, &result); err != nil {
		return nil, fmt.Errorf("fail to unmarshal pubkeys: %w", err)
	}
	var addressPairs []PubKeyContractAddressPair
	for _, v := range append(result.Asgard, result.Yggdrasil...) {
		kp := PubKeyContractAddressPair{
			PubKey:    v.PubKey,
			Contracts: make(map[common.Chain]common.Address),
		}
		// for _, item := range v.Routers {
		// 	kp.Contracts[item.Chain] = item.Router
		// }

		addressPairs = append(addressPairs, kp)
	}
	return addressPairs, nil
}

// TODO replace to thorNode's code once endpoint is build.
// Vaults a list of vault
type VaultStatus int32
type VaultType int32
type Vault struct {
	BlockHeight           int64           `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	PubKey                common.PubKey   `protobuf:"bytes,2,opt,name=pub_key,json=pubKey,proto3,casttype=gitlab.com/thorchain/thornode/common.PubKey" json:"pub_key,omitempty"`
	Coins                 common.Coins    `protobuf:"bytes,3,rep,name=coins,proto3,castrepeated=gitlab.com/thorchain/thornode/common.Coins" json:"coins"`
	Type                  VaultType       `protobuf:"varint,4,opt,name=type,proto3,enum=types.VaultType" json:"type,omitempty"`
	Status                VaultStatus     `protobuf:"varint,5,opt,name=status,proto3,enum=types.VaultStatus" json:"status,omitempty"`
	StatusSince           int64           `protobuf:"varint,6,opt,name=status_since,json=statusSince,proto3" json:"status_since,omitempty"`
	Membership            []string        `protobuf:"bytes,7,rep,name=membership,proto3" json:"membership,omitempty"`
	Chains                []string        `protobuf:"bytes,8,rep,name=chains,proto3" json:"chains,omitempty"`
	InboundTxCount        int64           `protobuf:"varint,9,opt,name=inbound_tx_count,json=inboundTxCount,proto3" json:"inbound_tx_count,omitempty"`
	OutboundTxCount       int64           `protobuf:"varint,10,opt,name=outbound_tx_count,json=outboundTxCount,proto3" json:"outbound_tx_count,omitempty"`
	PendingTxBlockHeights []int64         `protobuf:"varint,11,rep,packed,name=pending_tx_block_heights,json=pendingTxBlockHeights,proto3" json:"pending_tx_block_heights,omitempty"`
	Routers               []ChainContract `protobuf:"bytes,22,rep,name=routers,proto3" json:"routers"`
}
type Vaults []Vault

// GetAsgards retrieve all the asgard vaults from thorchain
func (b *HumanChainBridge) GetAsgards() (Vaults, error) {
	buf, s, err := b.getWithPath(AsgardVault)
	if err != nil {
		return Vaults{}, fmt.Errorf("fail to get asgard vaults: %w", err)
	}
	if s != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d", s)
	}
	var vaults Vaults
	if err := json.Unmarshal(buf, &vaults); err != nil {
		return nil, fmt.Errorf("fail to unmarshal asgard vaults from json: %w", err)
	}
	return vaults, nil
}

// MakeLegacyCodec creates codec
func MakeLegacyCodec() *codec.LegacyAmino {
	cdc := codec.NewLegacyAmino()
	banktypes.RegisterLegacyAminoCodec(cdc)
	authtypes.RegisterLegacyAminoCodec(cdc)
	cosmos.RegisterCodec(cdc)
	stypes.RegisterCodec(cdc)
	return cdc
}
