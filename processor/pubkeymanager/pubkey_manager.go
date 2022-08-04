package pubkeymanager

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/humansdotai/humans/common"
	"github.com/humansdotai/humans/constants"
	"github.com/humansdotai/humans/processor/humanclient"
	"github.com/humansdotai/humans/processor/metrics"
)

// OnNewPubKey is a function that used as a callback , if somehow we need to do additional process when a new pubkey get added
type OnNewPubKey func(pk common.PubKey) error

// PubKeyValidator define the method that can be used to interact with public keys
type PubKeyValidator interface {
	IsValidPoolAddress(addr string, chain common.Chain) (bool, common.ChainPoolInfo)
	HasPubKey(pk common.PubKey) bool
	AddPubKey(pk common.PubKey, _ bool)
	AddNodePubKey(pk common.PubKey)
	RemovePubKey(pk common.PubKey)
	GetSignPubKeys() common.PubKeys
	GetNodePubKey() common.PubKey
	GetPubKeys() common.PubKeys
	RegisterCallback(callback OnNewPubKey)
	GetContracts(chain common.Chain) []common.Address
	GetContract(chain common.Chain, pk common.PubKey) common.Address
}

// pubKeyInfo is a struct to store pubkey information  in memory
type pubKeyInfo struct {
	PubKey      common.PubKey
	Contracts   map[common.Chain]common.Address
	Signer      bool
	NodeAccount bool
}

// PubKeyManager manager an always up to date pubkeys , which implement PubKeyValidator interface
type PubKeyManager struct {
	cdc        *codec.LegacyAmino
	bridge     *humanclient.HumanChainBridge
	pubkeys    []pubKeyInfo
	rwMutex    *sync.RWMutex
	logger     zerolog.Logger
	errCounter *prometheus.CounterVec
	m          *metrics.Metrics
	stopChan   chan struct{}
	callback   []OnNewPubKey
}

// NewPubKeyManager create a new instance of PubKeyManager
func NewPubKeyManager(bridge *humanclient.HumanChainBridge, m *metrics.Metrics) (*PubKeyManager, error) {
	return &PubKeyManager{
		cdc:        humanclient.MakeLegacyCodec(),
		logger:     log.With().Str("module", "public_key_mgr").Logger(),
		bridge:     bridge,
		errCounter: m.GetCounterVec(metrics.PubKeyManagerError),
		m:          m,
		stopChan:   make(chan struct{}),
		rwMutex:    &sync.RWMutex{},
		callback:   []OnNewPubKey{},
	}, nil
}

// Start to poll pubkeys from thorchain
func (pkm *PubKeyManager) Start() error {
	pubkeys, err := pkm.getPubkeys()
	if err != nil {
		return fmt.Errorf("fail to get pubkeys from humanschain: %w", err)
	}
	for _, pk := range pubkeys {
		pkm.AddPubKey(pk.PubKey, false)
	}

	// get smart contract address from HumanNode , and update it's internal
	pkm.updateContractAddresses(pubkeys)
	go pkm.updatePubKeys()
	return nil
}

// Stop pubkey manager
func (pkm *PubKeyManager) Stop() error {
	defer pkm.logger.Info().Msg("pubkey manager stopped")
	close(pkm.stopChan)
	return nil
}

func (pkm *PubKeyManager) updateContractAddresses(pairs []humanclient.PubKeyContractAddressPair) {
	pkm.rwMutex.Lock()
	defer pkm.rwMutex.Unlock()
	for _, pair := range pairs {
		for idx, item := range pkm.pubkeys {
			if item.PubKey == pair.PubKey {
				pkm.pubkeys[idx].Contracts = pair.Contracts
			}
		}
	}
}

// GetPubKeys return all the public keys managed by this PubKeyManager
func (pkm *PubKeyManager) GetPubKeys() common.PubKeys {
	pkm.rwMutex.RLock()
	defer pkm.rwMutex.RUnlock()
	pubkeys := make(common.PubKeys, len(pkm.pubkeys))
	for i, pk := range pkm.pubkeys {
		pubkeys[i] = pk.PubKey
	}
	return pubkeys
}

// GetSignPubKeys get all the public keys that local node is a signer
func (pkm *PubKeyManager) GetSignPubKeys() common.PubKeys {
	pkm.rwMutex.RLock()
	defer pkm.rwMutex.RUnlock()
	pubkeys := make(common.PubKeys, 0)
	for _, pk := range pkm.pubkeys {
		if pk.Signer {
			pubkeys = append(pubkeys, pk.PubKey)
		}
	}
	return pubkeys
}

// GetNodePubKey get node account pub key
func (pkm *PubKeyManager) GetNodePubKey() common.PubKey {
	pkm.rwMutex.RLock()
	defer pkm.rwMutex.RUnlock()
	for _, pk := range pkm.pubkeys {
		if pk.NodeAccount {
			return pk.PubKey
		}
	}
	return common.EmptyPubKey
}

// HasPubKey return true if the given public key exist
func (pkm *PubKeyManager) HasPubKey(pk common.PubKey) bool {
	pkm.rwMutex.RLock()
	defer pkm.rwMutex.RUnlock()
	return pkm.hasPubKeyNoLock(pk)
}

// hasPubKeyNoLock internal used only
func (pkm *PubKeyManager) hasPubKeyNoLock(pk common.PubKey) bool {
	for _, pubkey := range pkm.pubkeys {
		if pk.Equals(pubkey.PubKey) {
			return true
		}
	}
	return false
}

// AddPubKey add the given public key to internal storage
func (pkm *PubKeyManager) AddPubKey(pk common.PubKey, signer bool) {
	pkm.rwMutex.Lock()
	defer pkm.rwMutex.Unlock()

	if pkm.hasPubKeyNoLock(pk) {
		// pubkey already exists, update the signer... but only if signer is true
		if signer {
			for i, pubkey := range pkm.pubkeys {
				if pk.Equals(pubkey.PubKey) {
					pkm.pubkeys[i].Signer = signer
				}
			}
		}
	} else {
		// pubkey doesn't exist yet, append it...
		pkm.pubkeys = append(pkm.pubkeys, pubKeyInfo{
			PubKey:      pk,
			Signer:      signer,
			NodeAccount: false,
			Contracts:   map[common.Chain]common.Address{},
		})
		pkm.fireCallback(pk)
	}
}

// AddNodePubKey add the given public key as a node public key to internal storage
func (pkm *PubKeyManager) AddNodePubKey(pk common.PubKey) {
	pkm.rwMutex.Lock()
	defer pkm.rwMutex.Unlock()

	for i, pubkey := range pkm.pubkeys {
		if pubkey.PubKey.Equals(pk) {
			pkm.pubkeys[i].Signer = true
			pkm.pubkeys[i].NodeAccount = true
			return
		}
	}

	if !pkm.hasPubKeyNoLock(pk) {
		pkm.pubkeys = append(pkm.pubkeys, pubKeyInfo{
			PubKey:      pk,
			Signer:      true,
			NodeAccount: true,
			Contracts:   map[common.Chain]common.Address{},
		})
		// a new pubkey get added , fire callback
		pkm.fireCallback(pk)
	}
}

// RemovePubKey remove the given public key from internal storage
func (pkm *PubKeyManager) RemovePubKey(pk common.PubKey) {
	pkm.rwMutex.Lock()
	defer pkm.rwMutex.Unlock()
	pkm.removePubKeyInternal(pk)
}

// removePubKeyInternal is a func to be used internally , and it doesn't lock the access to pkm.pubkeys
// caller need to lock pkm.pubkeys
func (pkm *PubKeyManager) removePubKeyInternal(pk common.PubKey) {
	for i, pubkey := range pkm.pubkeys {
		if pk.Equals(pubkey.PubKey) {
			pkm.pubkeys[i] = pkm.pubkeys[len(pkm.pubkeys)-1] // Copy last element to index i.
			pkm.pubkeys[len(pkm.pubkeys)-1] = pubKeyInfo{}   // Erase last element (write zero value).
			pkm.pubkeys = pkm.pubkeys[:len(pkm.pubkeys)-1]   // Truncate slice.
			break
		}
	}
}

func (pkm *PubKeyManager) fetchPubKeys() {
	addressPairs, err := pkm.getPubkeys()
	if err != nil {
		pkm.logger.Error().Err(err).Msg("fail to get pubkeys from HumansChain")
		return
	}
	var pubkeys common.PubKeys
	for _, pk := range addressPairs {
		pkm.AddPubKey(pk.PubKey, false)
		pubkeys = append(pubkeys, pk.PubKey)
	}
	pkm.updateContractAddresses(addressPairs)
	vaults, err := pkm.bridge.GetAsgards()
	if err != nil {
		return
	}

	for _, vault := range vaults {
		if vault.GetMembership().Contains(pkm.GetNodePubKey()) {
			pkm.AddPubKey(vault.PubKey, true)
			pubkeys = append(pubkeys, vault.PubKey)
		}
	}
	pkm.rwMutex.Lock()
	defer pkm.rwMutex.Unlock()
	// prune retired addresses
	for i, pk := range pkm.pubkeys {
		if pk.NodeAccount {
			// never remove our own pubkey
			continue
		}
		if i < (len(pkm.pubkeys) - 2) { // don't delete the more recent (last) pubkeys
			if !pubkeys.Contains(pk.PubKey) {
				pkm.removePubKeyInternal(pk.PubKey)
			}
		}
	}
}

func (pkm *PubKeyManager) updatePubKeys() {
	pkm.logger.Info().Msg("start to update pub keys")
	defer pkm.logger.Info().Msg("stop to update pub keys")
	for {
		select {
		case <-pkm.stopChan:
			return
		case <-time.After(constants.HumanchainBlockTime):
			pkm.fetchPubKeys()
		}
	}
}

func matchAddress(addr string, chain common.Chain, key common.PubKey) (bool, common.ChainPoolInfo) {
	cpi, err := common.NewChainPoolInfo(chain, key)
	if err != nil {
		return false, common.EmptyChainPoolInfo
	}
	if strings.EqualFold(cpi.PoolAddress.String(), addr) {
		return true, cpi
	}
	return false, common.EmptyChainPoolInfo
}

// IsValidPoolAddress check whether the given address is a pool addr
func (pkm *PubKeyManager) IsValidPoolAddress(addr string, chain common.Chain) (bool, common.ChainPoolInfo) {
	pkm.rwMutex.RLock()
	defer pkm.rwMutex.RUnlock()

	for _, pk := range pkm.pubkeys {
		ok, cpi := matchAddress(addr, chain, pk.PubKey)
		if ok {
			return ok, cpi
		}
	}
	return false, common.EmptyChainPoolInfo
}

// getPubkeys from HumansChain
func (pkm *PubKeyManager) getPubkeys() ([]humanclient.PubKeyContractAddressPair, error) {
	return pkm.bridge.GetPubKeys()
}

// RegisterCallback register a call back that will be fired when a new key get added into the local memory storage
func (pkm *PubKeyManager) RegisterCallback(callback OnNewPubKey) {
	pkm.callback = append(pkm.callback, callback)
}

func (pkm *PubKeyManager) fireCallback(pk common.PubKey) {
	for _, item := range pkm.callback {
		if err := item(pk); err != nil {
			pkm.logger.Err(err).Msg("fail to call callback")
		}
	}
}

// GetContracts return all the contracts for the requested chain
func (pkm *PubKeyManager) GetContracts(chain common.Chain) []common.Address {
	pkm.rwMutex.RLock()
	defer pkm.rwMutex.RUnlock()
	var result []common.Address
	for _, pk := range pkm.pubkeys {
		if pk.Contracts == nil {
			continue
		}
		if addr, ok := pk.Contracts[chain]; ok {
			result = append(result, addr)
		}
	}
	return result
}

// GetContract return the contract address that match the given chain and pubkey
func (pkm *PubKeyManager) GetContract(chain common.Chain, pubKey common.PubKey) common.Address {
	pkm.rwMutex.RLock()
	defer pkm.rwMutex.RUnlock()
	var result common.Address
	for _, pk := range pkm.pubkeys {
		if !pk.PubKey.Equals(pubKey) {
			continue
		}
		if pk.Contracts == nil {
			continue
		}
		result = pk.Contracts[chain]
	}
	return result
}
