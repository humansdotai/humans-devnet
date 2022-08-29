package signature

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	config "github.com/humansdotai/humans/processor/config"
	"github.com/humansdotai/humans/x/humans/types"
	"github.com/humansdotai/humans/x/humans/types/ethPool"
)

// RSASignature
type RSASignature struct {
	privkey *rsa.PrivateKey
	pubkey  *rsa.PublicKey

	config *config.CredentialConfiguration
}

// Generates instances for
func NewSigGen(config *config.CredentialConfiguration) (*RSASignature, error) {
	return &RSASignature{
		privkey: nil,
		pubkey:  nil,
		config:  config,
	}, nil
}

// Generates RSA private & pubKey
func (o *RSASignature) GeneratePrivateKey(bits int) error {
	privkey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}

	// Generate priv & pub keys
	o.privkey = privkey
	o.pubkey = &privkey.PublicKey

	return nil
}

// Set publickey to ethereum smart contract
func (o *RSASignature) InitializePubKeyOnEtherem(signer string) bool {
	if signer != types.MAIN_VALIDATOR_MONIKER {
		return true
	}

	client, err := ethclient.Dial(o.config.URL_Ethereum_RPC_Node_Provider)

	if err != nil {
		log.Fatalln(err)
		return false
	}

	privateKey, err := ethCrypto.HexToECDSA(o.config.Ethereum_Owner_Account_Private_Key)
	if err != nil {
		log.Fatalln(err)
		return false
	}

	chainID := new(big.Int)
	chainID.SetInt64(4)

	transactOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalln(err)
		return false
	}

	poolAddress := common.HexToAddress(o.config.Ethereum_Pool_Address)
	instance, err := ethPool.NewHumansPool(poolAddress, client)

	if err != nil {
		log.Fatalln(err)
		return false
	}

	//--------------------
	m := o.privkey.N.Bytes()
	M := fmt.Sprintf("0x%x", m)

	e := strconv.FormatInt(int64(o.pubkey.E), 16)
	E := fmt.Sprintf("0x%0256s", e)

	// --- setting public key---
	modulus, _ := hexutil.Decode(M)
	exponent, _ := hexutil.Decode(E)

	transaction, err := instance.SetPublicKey(transactOpts, exponent, modulus)

	if err != nil {
		log.Fatalln(err)
		return false
	}

	fmt.Println("Transaction Hash:", transaction.Hash().Hex())

	return true
}

// Set Pub Key to contract
func (o *RSASignature) Start(signer string) bool {
	if o.privkey == nil || o.pubkey == nil {
		log.Fatalln("Keys are not yet created!")

		return false
	}

	// Initialize Ethereum Pool contract
	if !o.InitializePubKeyOnEtherem(signer) {
		log.Fatalln("Can't initialize pool contract on Ethereum!")

		return false
	}

	return true
}

// Sign a message and verify signature with go using PKCS1.
func (o *RSASignature) GenerateSignature(msg string) (string, error) {
	// Generate msg from string
	data := []byte(msg)

	// Sum it up with sha256
	digest := sha256.Sum256(data)

	// Generates signature
	signature, signErr := rsa.SignPKCS1v15(rand.Reader, o.privkey, crypto.SHA256, digest[:])

	if signErr != nil {
		fmt.Println("Could not sign message:", signErr.Error())

		return "", signErr
	}

	sig := fmt.Sprintf("0x%x", signature)

	return sig, nil
}
