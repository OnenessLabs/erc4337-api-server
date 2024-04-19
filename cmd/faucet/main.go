package main

import (
	"fmt"
	"github.com/oneness/erc-4337-api/chain"
	"github.com/oneness/erc-4337-api/crypto"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/abi"
	"github.com/umbracle/ethgo/contract"
	"github.com/umbracle/ethgo/jsonrpc"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

// call a contract
func main() {
	var functions = []string{
		"function requestTokens()",
	}

	abiContract, err := abi.NewABIFromList(functions)
	handleErr(err)

	// Matic token
	addr := ethgo.HexToAddress("0x4a67546101cB9a5c3F6A807229eeBadB3511c989")

	chainRpc, err := jsonrpc.NewClient("https://rpc.devnet.onenesslabs.io/")
	handleErr(err)

	var maybeKey *chain.EcdsaKey
	ChainSKHex := "cec327609d36790502f393464b6cd9de2ba228b295bce9779f8bdffe3f5787d4"
	if len(ChainSKHex) != 0 {
		if sk, err := crypto.SKFromHex(ChainSKHex); err != nil {
			handleErr(err)
		} else {
			maybeKey = &chain.EcdsaKey{SK: sk}
		}
		ChainKeyAddr := maybeKey.Address()
		fmt.Printf(ChainKeyAddr.String())
	}

	opts := []contract.ContractOption{contract.WithJsonRPC(chainRpc.Eth())}
	if maybeKey != nil {
		opts = append(opts, contract.WithSender(maybeKey))
	}
	c := contract.NewContract(addr, abiContract, opts...)

	//res, err := c.Call("requestTokens", ethgo.Latest)
	//handleErr(err)

	txn, err := c.Txn("requestTokens", ethgo.Latest)
	handleErr(err)

	err = txn.Do()
	handleErr(err)

	receipt, err := txn.Wait()
	handleErr(err)

	fmt.Printf("Transaction mined at: %s", receipt.TransactionHash)
	//fmt.Printf("TotalSupply: %s", res["totalSupply"].(*big.Int))
}
