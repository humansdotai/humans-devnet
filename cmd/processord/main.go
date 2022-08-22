package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/humansdotai/humans/app"
	"github.com/humansdotai/humans/cmd"
	tcommon "github.com/humansdotai/humans/common"
	config "github.com/humansdotai/humans/processor/config"
	humanclient "github.com/humansdotai/humans/processor/humanclient"
	"github.com/humansdotai/humans/processor/humanclient/cosmos"
	"github.com/humansdotai/humans/processor/observer"
	"github.com/humansdotai/humans/processor/pubkeymanager"
	"github.com/rs/zerolog/log"
	"gitlab.com/thorchain/tss/go-tss/common"
	"gitlab.com/thorchain/tss/go-tss/tss"
)

func initPrefix() {
	cosmosSDKConfg := cosmos.GetConfig()
	cosmosSDKConfg.SetBech32PrefixForAccount(cmd.Bech32PrefixAccAddr, cmd.Bech32PrefixAccPub)
	cosmosSDKConfg.Seal()
}

func main() {
	// Init prefix
	initPrefix()

	// args := os.Args
	// if len(args) < 1 {
	// 	return
	// }

	signer := "validator" //args[2]
	password := "password"

	kb, _, err := humanclient.GetKeyringKeybase("", signer, password)
	if err != nil {
		fmt.Println("fail to get keyring keybase")
		return
	}

	info, err := kb.Key(signer)
	pubKey := info.GetPubKey().Address().String()
	addr := info.GetAddress().String()

	k := humanclient.NewKeysWithKeybase(kb, signer, password)

	cfg := &humanclient.BridgeConfig{
		ChainId:         "test",
		ChainHost:       "127.0.0.1:1317",
		ChainRPC:        "127.0.0.1:26657",
		ChainHomeFolder: "~/.humans/",
	}

	HumanChainBridge, err := humanclient.NewHumanChainBridge(k, cfg, signer, pubKey, addr)

	// PubKey Manager
	pubkeyMgr, err := pubkeymanager.NewPubKeyManager(HumanChainBridge)
	if err != nil {
		log.Fatal().Err(err).Msg("fail to create pubkey manager")
	}
	if err := pubkeyMgr.Start(); err != nil {
		log.Fatal().Err(err).Msg("fail to start pubkey manager")
	}

	// setup TSS signing
	priKey, err := k.GetPrivateKey()
	if err != nil {
		fmt.Println("fail to get private key")
		return
	}

	fmt.Println(priKey)

	bootstrapPeers, err := cfg.TSS.GetBootstrapPeers()
	if err != nil {
		log.Fatal().Err(err).Msg("fail to get bootstrap peers")
	}

	tmPrivateKey := tcommon.CosmosPrivateKeyToTMPrivateKey(priKey)
	tssIns, err := tss.NewTss(
		bootstrapPeers,
		cfg.TSS.P2PPort,
		tmPrivateKey,
		cfg.TSS.Rendezvous,
		app.DefaultNodeHome,
		common.TssConfig{
			EnableMonitor:   true,
			KeyGenTimeout:   300 * time.Second, // must be shorter than constants.JailTimeKeygen
			KeySignTimeout:  60 * time.Second,  // must be shorter than constants.JailTimeKeysign
			PartyTimeout:    45 * time.Second,
			PreParamTimeout: 5 * time.Minute,
		},
		nil,
		cfg.TSS.ExternalIP,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("fail to create tss instance")
	}

	if err := tssIns.Start(); err != nil {
		log.Err(err).Msg("fail to start tss instance")
	}

	// Load app configuration
	config, err := config.NewCredentialConfig()
	err = config.LoadConfig()
	if err != nil {
		fmt.Println("fail to load config")
		return
	}

	obs_storage := ""
	obs, err := observer.NewObserver(HumanChainBridge, obs_storage, config)
	if err != nil {
		fmt.Println("fail to create observer")
		return
	}

	if err = obs.Start(); err != nil {
		fmt.Println("fail to start observer")
		return
	}

	// wait....
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("stop signal received")

	// stop observer
	if err := obs.Stop(); err != nil {
		fmt.Println("fail to stop observer")
	}

	// // stop tss
	// if err := tss.Stop(); err != nil {
	// 	fmt.Println("fail to stop tss")
	// }
}
