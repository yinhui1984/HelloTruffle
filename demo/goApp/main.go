package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"goApp/contracts"
	"math/big"
)

//https://goethereumbook.org/client-setup/

func main() {

	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		fmt.Println("could not connect to local node, err:", err)
		return
	}

	contractAddress := common.HexToAddress("0x6B62e4E253823FBC65E0B93d63ee149350158a18")
	calculator, err := contracts.NewCalculator(contractAddress, client)
	if err != nil {
		fmt.Println("could not instantiate contract, err:", err)
		return
	}

	callOpts := bind.CallOpts{
		Context: nil,
		Pending: false,
	}
	a := big.NewInt(1)
	b := big.NewInt(2)
	result, err := calculator.Add(&callOpts, a, b)
	if err != nil {
		fmt.Println("could not call contract, err:", err)
		return
	}
	fmt.Println("add result:", result)

	result, err = calculator.Subtract(&callOpts, a, b)
	if err != nil {
		fmt.Println("could not call contract, err:", err)
		return
	}
	fmt.Println("add result:", result)
}
