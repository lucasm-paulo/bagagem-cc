package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/lucasm-paulo/bagagem-cc/chaincode/assettypes"
	"github.com/lucasm-paulo/bagagem-cc/chaincode/datatypes"
	"github.com/lucasm-paulo/bagagem-cc/chaincode/header"
	"github.com/goledgerdev/cc-tools/assets"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

func SetupCC() error {
	tx.InitHeader(tx.Header{
		Name:    header.Name,
		Version: header.Version,
		Colors:  header.Colors,
		Title:   header.Title,
	})

	assets.InitDynamicAssetTypeConfig(assettypes.DynamicAssetTypes)

	tx.InitTxList(txList)

	err := assets.CustomDataTypes(datatypes.CustomDataTypes)
	if err != nil {
		fmt.Printf("Erro ao injetar tipo de dados personalizado: %s", err)
		return err
	}
	assets.InitAssetList(append(assetTypeList, assettypes.CustomAssets...))
	return nil
}

// main function starts up the chaincode in the container during instantiate
func main() {
	log.Printf("Iniciando chaincode %s versão %s\n", header.Name, header.Version)

	err := SetupCC()
	if err != nil {
		return
	}
	if err = shim.Start(new(CCDemo)); err != nil {
		fmt.Printf("Erro ao iniciar chaincode: %s", err)
	}
}

// CCDemo implements the shim.Chaincode interface
type CCDemo struct{}

// Init is called during chaincode instantiation to initialize any
// data. Note that chaincode upgrade also calls this function to reset
// or to migrate data.
func (t *CCDemo) Init(stub shim.ChaincodeStubInterface) (response pb.Response) {
	// Defer logging function
	defer logTx(stub, time.Now(), &response)

	if assettypes.DynamicAssetTypes.Enabled {
		sw := &sw.StubWrapper{
			Stub: stub,
		}
		err := assets.RestoreAssetList(sw, true)
		if err != nil {
			response = err.GetErrorResponse()
			return
		}
	}

	err := assets.StartupCheck()
	if err != nil {
		response = err.GetErrorResponse()
		return
	}

	err = tx.StartupCheck()
	if err != nil {
		response = err.GetErrorResponse()
		return
	}

	// Get the args from the transaction proposal
	args := stub.GetStringArgs()

	// Test if argument list is empty
	if len(args) != 1 {
		response = shim.Error("o método init espera 1 argumento e recebeu: " + strings.Join(args, ", "))
		response.Status = 400
		return
	}

	// Test if argument is "init" or "upgrade". Fails otherwise.
	if args[0] != "init" && args[0] != "upgrade" {
		response = shim.Error("os argumentos deveriam ser 'init' ou 'upgrade' (como definido pelo SDK do Node.js)")
		response.Status = 400
		return
	}

	response = shim.Success(nil)
	return
}

// Invoke is called per transaction on the chaincode.
func (t *CCDemo) Invoke(stub shim.ChaincodeStubInterface) (response pb.Response) {
	// Defer logging function
	defer logTx(stub, time.Now(), &response)

	var result []byte

	result, err := tx.Run(stub)

	if err != nil {
		response = err.GetErrorResponse()
		return
	}
	response = shim.Success([]byte(result))
	return
}

func logTx(stub shim.ChaincodeStubInterface, beginTime time.Time, response *pb.Response) {
	fn, _ := stub.GetFunctionAndParameters()
	log.Printf("%d %s %s %s\n", response.Status, fn, time.Since(beginTime), response.Message)
}
