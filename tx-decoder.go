package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/calvinlauyh/cosmosutils"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/tendermint/tendermint/crypto"
	marketTypes "github.com/terra-money/core/x/market/types"
	oracleTypes "github.com/terra-money/core/x/oracle/types"

	// legacyTypes "github.com/terra-money/core/x/msgauth/legacy/v04"
	//	treasuryTypes "github.com/terra-money/core/x/treasury/types"
	vestingTypes "github.com/terra-money/core/x/vesting/types"
	wasmTypes "github.com/terra-money/core/x/wasm/types"
)

func main() {
	b64Tx := os.Args[1]

	decoder := cosmosutils.DefaultDecoder
	decoder.RegisterInterfaces(RegisterDefaultInterfaces)
	data, _ := base64.StdEncoding.DecodeString(b64Tx)
	fmt.Println("hash", hex.EncodeToString(crypto.Sha256(data)))
	tx, err := decoder.Decode(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Handle the error and work with tx
	txJson, err := tx.MarshalToJSON()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(txJson))
}

func RegisterDefaultInterfaces(interfaceRegistry types.InterfaceRegistry) {
	std.RegisterInterfaces(interfaceRegistry)
	oracleTypes.RegisterInterfaces(interfaceRegistry)
	marketTypes.RegisterInterfaces(interfaceRegistry)
	// legacyTypes.RegisterInterfaces(interfaceRegistry)
	//	treasuryTypes.RegisterInterfaces(interfaceRegistry)
	vestingTypes.RegisterInterfaces(interfaceRegistry)
	wasmTypes.RegisterInterfaces(interfaceRegistry)

}
