package wasmclient

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
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/humansdotai/humans/app"
	"github.com/humansdotai/humans/processor/humanclient"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
)

// WasmTxBridge will be used to send tx to HumansChain
type WasmTxBridge struct {
	keys          *humanclient.Keys
	cfg           humanclient.BridgeConfig
	blockHeight   uint64
	accountNumber uint64
	seqNumber     uint64
	httpClient    *retryablehttp.Client
	broadcastLock *sync.RWMutex

	signerName                 string
	lastBlockHeightCheck       time.Time
	lastHumanschainBlockHeight uint64
	pubKey                     string
	voterAddress               string
}

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

// NewWasmTxBridge create a new instance of WasmTxBridge
func NewWasmTxBridge(k *humanclient.Keys, cfg *humanclient.BridgeConfig, signer string, pubKey string, voter string) (*WasmTxBridge, error) {
	httpClient := retryablehttp.NewClient()
	httpClient.Logger = nil

	return &WasmTxBridge{
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
func (b *WasmTxBridge) GetContext() client.Context {
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

func (b *WasmTxBridge) getWithPath(path string) ([]byte, int, error) {
	return b.get(b.getHumanChainURL(path))
}

// getThorChainURL with the given path
func (b *WasmTxBridge) getHumanChainURL(path string) string {
	uri := url.URL{
		Scheme: "http",
		Host:   b.cfg.ChainHost,
		Path:   path,
	}
	return uri.String()
}

// get handle all the low level http GET calls using retryablehttp.ThorchainBridge
func (b *WasmTxBridge) get(url string) ([]byte, int, error) {
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
func (b *WasmTxBridge) GetVoterInfo() (string, string) {
	return b.pubKey, b.voterAddress
}

//
func (b *WasmTxBridge) GetMonikerName() string {
	return b.signerName
}

// getAccountNumberAndSequenceNumber returns account and Sequence number required to post into thorchain
func (b *WasmTxBridge) getAccountNumberAndSequenceNumber() (uint64, uint64, error) {
	path := fmt.Sprintf("%s/%s", "/cosmos/auth/v1beta1/accounts", b.keys.GetSignerInfo().GetAddress())

	body, _, err := b.getWithPath(path)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to get auth accounts: %w", err)
	}

	var resp humanclient.AccountResp
	if err := json.Unmarshal(body, &resp); err != nil {
		return 0, 0, fmt.Errorf("failed to unmarshal account resp: %w", err)
	}

	acc := resp.Account
	return acc.AccountNumber, acc.Sequence, nil
}

// HumansBlockTime Block time of HUMANSChain
var HumansBlockTime = 5 * time.Second

// GetBlockHeight returns the current height for humans blocks
func (b *WasmTxBridge) GetBlockHeight() (uint64, error) {
	if time.Since(b.lastBlockHeightCheck) < HumansBlockTime && b.lastHumanschainBlockHeight > 0 {
		return b.lastHumanschainBlockHeight, nil
	}

	latestBlock, err := b.GetLastBlock("")
	if err != nil {
		return 0, fmt.Errorf("failed to HumanchainHeight: %w", err)
	}

	b.lastBlockHeightCheck = time.Now()
	h, _ := strconv.ParseUint(latestBlock.Block.Header.Height, 10, 64)

	return h, nil
}

// getLastBlock calls the /lastblock/{chain} endpoint and Unmarshal's into the QueryResLastBlockHeights type
func (b *WasmTxBridge) GetLastBlock(chain string) (humanclient.QueryResLastBlockHeights, error) {
	path := "/cosmos/base/tendermint/v1beta1/blocks/latest"
	if chain != "" {
		path = fmt.Sprintf("%s/%s", path, chain)
	}
	buf, _, err := b.getWithPath(path)
	if err != nil {
		return humanclient.QueryResLastBlockHeights{}, fmt.Errorf("failed to get lastblock: %w", err)
	}

	var lastBlock humanclient.QueryResLastBlockHeights
	if err := json.Unmarshal(buf, &lastBlock); err != nil {
		return humanclient.QueryResLastBlockHeights{}, fmt.Errorf("failed to unmarshal last block: %w", err)
	}
	return lastBlock, nil
}
