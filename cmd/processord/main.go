package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/humansdotai/humans/cmd"
	config "github.com/humansdotai/humans/processor/config"
	"github.com/humansdotai/humans/processor/humanclient"
	"github.com/humansdotai/humans/processor/humanclient/cosmos"
	"github.com/humansdotai/humans/processor/observer"
	signature "github.com/humansdotai/humans/processor/signature"
	"github.com/humansdotai/humans/processor/wasmclient"
)

func initPrefix() {
	cosmosSDKConfg := cosmos.GetConfig()
	cosmosSDKConfg.SetBech32PrefixForAccount(cmd.Bech32PrefixAccAddr, cmd.Bech32PrefixAccPub)
	cosmosSDKConfg.Seal()
}

func main() {
	// Init prefix
	initPrefix()

	args := os.Args
	if len(args) < 1 {
		return
	}

	signer := args[2]
	password := "password"

	//-----------------------------------
	// Load app configuration
	config, err := config.NewCredentialConfig()
	err = config.LoadConfig()
	if err != nil {
		fmt.Println("fail to load config")
		return
	}

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

	//------------------------------------
	// -------Wasm Tx Bridge configure----
	wasm_signer := config.Humanchain_Pool_Owner_Signer_KeyName
	kWasm := humanclient.NewKeysWithKeybase(kb, wasm_signer, password)
	kbWasm, _, err := humanclient.GetKeyringKeybase("", wasm_signer, password)
	if err != nil {
		fmt.Println("fail to get keyring keybase")
		return
	}

	infoWasm, err := kbWasm.Key(wasm_signer)
	pubKeyWasm := infoWasm.GetPubKey().Address().String()
	addrWasm := infoWasm.GetAddress().String()

	WasmTxBridge, err := wasmclient.NewWasmTxBridge(kWasm, cfg, wasm_signer, pubKeyWasm, addrWasm)
	if err != nil {
		fmt.Println("fail to create wasm bridge config")
		return
	}
	//----------------------------

	HumanChainBridge, err := humanclient.NewHumanChainBridge(k, cfg, signer, pubKey, addr)
	if err != nil {
		fmt.Println("fail to create human bridge config")
		return
	}

	// Initialize key sign module
	tss, err := signature.NewSigGen(config)
	if err != nil {
		fmt.Println("fail to load signature module")
		return
	}

	// Private key generation
	err = tss.GeneratePrivateKey(1024)
	if err != nil {
		fmt.Println("fail to create private key")
		return
	}

	// Key generation module initialization
	if !tss.Start() {
		fmt.Println("fail to start TSS")
		return
	}

	obs_storage := ""
	obs, err := observer.NewObserver(HumanChainBridge, WasmTxBridge, obs_storage, config, tss)
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

}
