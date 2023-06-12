package main

import (
	txdefs "github.com/lucasm-paulo/bagagem-cc/chaincode/txdefs"

	tx "github.com/goledgerdev/cc-tools/transactions"
)

var txList = []tx.Transaction{
	tx.CreateAsset,
	tx.UpdateAsset,
	tx.DeleteAsset,
	txdefs.LinkBaggageToPassenger,
}
