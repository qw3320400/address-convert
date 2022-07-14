package main

import (
	"flag"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

/*
./main --address 0xc8d5d73eb919b37d131b815ec9f1dd915f3eafd3 --prefix cro
./main --address cro1er2aw04erxeh6ycms90vnuwaj90nat7nsen9cc --prefix cro
./main --address 0xc8d5d73eb919b37d131b815ec9f1dd915f3eafd3 --prefix cosmos
*/

var (
	address = flag.String("address", "", "address")
	prefix  = flag.String("prefix", "", "prefix")
)

func main() {
	flag.Parse()

	if strings.HasPrefix(*address, "0x") {
		sdk.GetConfig().SetBech32PrefixForAccount(*prefix, *prefix+"pub")
		addr := common.FromHex(*address)
		fmt.Println(sdk.AccAddress(addr).String())
	} else {
		sdk.GetConfig().SetBech32PrefixForAccount(*prefix, *prefix+"pub")
		addr, err := sdk.AccAddressFromBech32(*address)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(common.BytesToAddress(addr.Bytes()).Hex())
	}
}
