package txdefs

import (
	"encoding/json"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

// Links baggage to a passenger
// POST Method
var LinkBaggageToPassenger = tx.Transaction{
	Tag:         "linkBaggageToPassenger",
	Label:       "Vincular bagagem a um passageiro",
	Description: "Atribui uma bagagem a um passageiro",
	Method:      "PUT", // verificar
	Callers:     []string{`org1MSP`, "orgMSP"}, // // This means only org1 can call this transaction

	Args: []tx.Argument{
		{
			Tag:         "baggage",
			Label:       "Bagagem",
			Description: "Baggage",
			DataType:    "->baggage",
			Required:    true,
		},
		{
			Tag:         "person",
			Label:       "Person",
			Description: "Dados pessoais de uma pessoa",
			DataType:    "->person",
			Required:    true,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		baggageKey, ok := req["baggage"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "O parâmetro 'bagagem' deve ser um ativo")
		}
		passengerKey, ok := req["passenger"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "O parâmetro 'passageiro' deve ser um ativo")
		}

		// Returns Baggage from channel
		baggageAsset, err := baggageKey.Get(stub)
		if err != nil {
			return nil, errors.WrapError(err, "falha ao obter ativo 'bagagem' do ledger")
		}
		baggageMap := (map[string]interface{})(*baggageAsset)

		// Returns person from channel
		passengerAsset, err := passengerKey.Get(stub)
		if err != nil {
			return nil, errors.WrapError(err, "falha ao obter ativo 'passageiro' do ledger")
		}
		passengerMap := (map[string]interface{})(*passengerAsset)

		setPassengerKey := make(map[string]interface{})
		setPassengerKey["@assetType"] = "person"
		setPassengerKey["@key"] = passengerMap["@key"]

		// Set data
		baggageMap["passenger"] = setPassengerKey

		baggageMap, err = baggageAsset.Update(stub, baggageMap)
		if err != nil {
			return nil, errors.WrapError(err, "falha ao vincular bagagem ao passageiro")
		}

		// Marshal asset back to JSON format
		baggageJSON, nerr := json.Marshal(baggageMap)
		if nerr != nil {
			return nil, errors.WrapError(err, "falha ao converter resposta")
		}

		return baggageJSON, nil
	},
}
