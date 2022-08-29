package observer

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	etherTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/humansdotai/humans/x/humans/types"
	token "github.com/humansdotai/humans/x/humans/types/erc20"
	"github.com/humansdotai/humans/x/humans/types/ethPool"
)

// LogTransfer ..
type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}

// MainMetaData contains all meta data concerning the Main contract.
var MainMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MainABI is the input ABI used to generate the binding from.
// Deprecated: Use MainMetaData.ABI instead.
var MainABI = MainMetaData.ABI

// MainMetaData contains all meta data concerning the Main contract.
var HumanMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"name\":\"SignatureVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenWithdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"balanceOfToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_e\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_m\",\"type\":\"bytes\"}],\"name\":\"setPublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_msg\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_s\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MainABI is the input ABI used to generate the binding from.
// Deprecated: Use MainMetaData.ABI instead.
var HumanABI = HumanMetaData.ABI

// Fetches the USDC balance of Ethereum pool account
func (o *Observer) FetchBalanceOfEtherPool() bool {
	client, err := ethclient.Dial(o.config.URL_Ethereum_RPC_Node_Provider)
	if err != nil {
		return false
	}

	// Golem (GNT) Address
	tokenAddress := common.HexToAddress(o.config.Ethereum_USDK_Token_Address)
	instance, err := token.NewMain(tokenAddress, client)
	if err != nil {
		return false
	}

	address := common.HexToAddress(o.config.Ethereum_Pool_Address)
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		return false
	}

	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		return false
	}

	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))
	_, voter := o.HumanChainBridge.GetVoterInfo()

	msg := types.NewMsgUpdateBalance(voter, types.CHAIN_ETHEREUM, fmt.Sprintf("%f", value), fmt.Sprintf("%v", decimals))
	o.ArrMsgUpdateBalance = append(o.ArrMsgUpdateBalance, msg)

	return true
}

// Transfer token on Ethereum
func (o *Observer) EthereumTransferTokenToTarget(txdata *types.TransactionData, signature string, transMsg string, moniker string) bool {
	if moniker != types.MAIN_VALIDATOR_MONIKER {
		return true
	}

	client, err := ethclient.Dial(o.config.URL_Ethereum_RPC_Node_Provider)

	if err != nil {
		log.Fatalln(err)
		return false
	}

	privateKey, err := crypto.HexToECDSA(o.config.Ethereum_Owner_Account_Private_Key)
	if err != nil {
		log.Fatalln(err)
		return false
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalln(err)
		return false
	}

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

	tokenAddress := common.HexToAddress(o.config.Ethereum_USDK_Token_Address)
	toAddress := common.HexToAddress(txdata.TargetAddress)

	// -----

	amt, _ := strconv.ParseFloat(txdata.Amount, 64)

	// -----
	amtFee, err := strconv.ParseFloat(txdata.Fee, 64)
	amt -= amtFee

	samt := fmt.Sprintf("%f", amt*1e18) /// consider the denomination number to 18

	amount := new(big.Int)
	amount.SetString(samt, 10) // sets the value to 1 USDC, in the token denomination

	// 0x68656c6c6f20776f726c64
	hexMsg := fmt.Sprintf("0x%s", hex.EncodeToString(([]byte)(transMsg)))

	msg, _ := hexutil.Decode(hexMsg)
	sig, _ := hexutil.Decode(signature)

	transaction, err := instance.Withdraw(transactOpts, tokenAddress, toAddress, amount, msg, sig)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Transaction Hash:", transaction.Hash().Hex())

	return true
}

// Keep listening to WSS and fetch transaction deposited to the pool
func (o *Observer) ProcessTxInsEthExternal() {
	client, err := ethclient.Dial(o.config.URL_Ethereum_RPC_Node_Provider_WSS)
	if err != nil {
		return
	}

	contractAddress := common.HexToAddress(o.config.Ethereum_USDK_Token_Address)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan etherTypes.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		return
	}

	for {
		select {
		case <-o.stopChan:
			return
		case <-sub.Err():
			time.Sleep(time.Second * 5)
			o.EthSocketErr <- true
			return
		case vLog := <-logs:
			o.EthereumParseLog(vLog)
		}
	}
}

func (o *Observer) EthereumParseLog(vLog etherTypes.Log) {
	if o.continsHash(o.EthTxHasVoted, vLog.TxHash.String()) {
		return
	}

	contractAbi, err := abi.JSON(strings.NewReader(MainABI))
	if err != nil {
		return
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)

	switch vLog.Topics[0].Hex() {
	case logTransferSigHash.Hex():
		var transferEvent LogTransfer

		i, err := contractAbi.Unpack("Transfer", vLog.Data)
		if err != nil {
			return
		}

		amt := i[0].(*big.Int)
		famt := new(big.Float).SetInt(amt)
		ff, _ := famt.Float64()
		tokenAmount := ff / 1e18

		transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
		transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

		if transferEvent.From.String() == o.config.Ethereum_Pool_Address {
			// Send true to SolPoolchange channel
			o.EthPoolChanged <- true
			return
		}

		if transferEvent.To.String() != o.config.Ethereum_Pool_Address || tokenAmount == 0.0 {
			return
		}

		if transferEvent.From.String() == o.config.Ethereum_USDK_Token_Address {
			return
		}

		o.EthTxHasVoted = append(o.EthTxHasVoted, vLog.TxHash.String())

		_, voter := o.HumanChainBridge.GetVoterInfo()
		msg := types.NewMsgObservationVote(voter, vLog.TxHash.String(), types.CHAIN_ETHEREUM, transferEvent.From.Hex(), transferEvent.To.Hex(), fmt.Sprintf("%f", tokenAmount))
		o.ArrMsgObservationVote = append(o.ArrMsgObservationVote, msg)

		// Send true to EthPoolchange channel
		o.EthPoolChanged <- true
	}
}
