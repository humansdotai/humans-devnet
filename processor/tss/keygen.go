package tss

import (
	"fmt"
	"net/http"
	"time"

	"github.com/blang/semver"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gitlab.com/thorchain/tss/go-tss/keygen"
	"gitlab.com/thorchain/tss/go-tss/tss"

	"github.com/humansdotai/humans/common"
	"github.com/humansdotai/humans/constants"
	"github.com/humansdotai/humans/processor/humanclient"
	"github.com/humansdotai/humans/x/humans/types"
)

// KeyGen is
type KeyGen struct {
	keys           *humanclient.Keys
	logger         zerolog.Logger
	client         *http.Client
	server         *tss.TssServer
	bridge         *humanclient.HumanChainBridge
	currentVersion semver.Version
	lastCheck      time.Time
}

// NewTssKeyGen create a new instance of TssKeyGen which will look after TSS key stuff
func NewTssKeyGen(keys *humanclient.Keys, server *tss.TssServer, bridge *humanclient.HumanChainBridge) (*KeyGen, error) {
	if keys == nil {
		return nil, fmt.Errorf("keys is nil")
	}
	return &KeyGen{
		keys:   keys,
		logger: log.With().Str("module", "tss_keygen").Logger(),
		client: &http.Client{
			Timeout: time.Second * 130,
		},
		server: server,
		bridge: bridge,
	}, nil
}

type QueryVersion struct {
	Current semver.Version `json:"current"`
	Next    semver.Version `json:"next"`
}

func (kg *KeyGen) getVersion() semver.Version {
	requestTime := time.Now()
	if !kg.currentVersion.Equals(semver.Version{}) && requestTime.Sub(kg.lastCheck).Seconds() < constants.ThorchainBlockTime.Seconds() {
		return kg.currentVersion
	}
	version := QueryVersion{}
	kg.currentVersion = version
	kg.lastCheck = requestTime
	return kg.currentVersion
}

func (kg *KeyGen) GenerateNewKey(pKeys common.PubKeys) (common.PubKeySet, types.Blame, error) {
	// No need to do key gen
	if len(pKeys) == 0 {
		return common.EmptyPubKeySet, types.Blame{}, nil
	}
	var keys []string
	for _, item := range pKeys {
		keys = append(keys, item.String())
	}
	keyGenReq := keygen.Request{
		Keys: keys,
	}
	currentVersion := kg.getVersion()
	// get current THORChain block height
	blockHeight, err := kg.bridge.GetBlockHeight()
	if err != nil {
		return common.EmptyPubKeySet, types.Blame{}, fmt.Errorf("fail to get current thorchain block height: %w", err)
	}

	// this is just round the block height to the nearest 10
	keyGenReq.BlockHeight = blockHeight / 10 * 10
	keyGenReq.Version = currentVersion.String()

	ch := make(chan bool, 1)
	defer close(ch)
	timer := time.NewTimer(30 * time.Minute)
	defer timer.Stop()

	var resp keygen.Response
	go func() {
		resp, err = kg.server.Keygen(keyGenReq)
		ch <- true
	}()

	select {
	case <-ch:
		// do nothing
	case <-timer.C:
		panic("tss keygen timeout")
	}

	// copy blame to our own struct
	blame := types.Blame{
		FailReason: resp.Blame.FailReason,
		IsUnicast:  resp.Blame.IsUnicast,
		BlameNodes: make([]types.Node, len(resp.Blame.BlameNodes)),
	}
	for i, n := range resp.Blame.BlameNodes {
		blame.BlameNodes[i].Pubkey = n.Pubkey
		blame.BlameNodes[i].BlameData = n.BlameData
		blame.BlameNodes[i].BlameSignature = n.BlameSignature
	}

	if err != nil {
		// the resp from kg.server.Keygen will not be nil
		if blame.IsEmpty() {
			blame.FailReason = err.Error()
		}
		return common.EmptyPubKeySet, blame, fmt.Errorf("fail to keygen,err:%w", err)
	}

	cpk, err := common.NewPubKey(resp.PubKey)
	if err != nil {
		return common.EmptyPubKeySet, blame, fmt.Errorf("fail to create common.PubKey,%w", err)
	}

	// TODO later on THORNode need to have both secp256k1 key and ed25519
	return common.NewPubKeySet(cpk, cpk), blame, nil
}
