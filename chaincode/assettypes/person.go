package assettypes

import (
	"fmt"

	"github.com/goledgerdev/cc-tools/assets"
)

var Person = assets.AssetType{
	Tag:         "person",
	Label:       "Person",
	Description: "Dados pessoais de uma pessoa",

	Props: []assets.AssetProp{
		{
			// Primary key
			Required: true,
			IsKey:    true,
			Tag:      "id",
			Label:    "CPF",
			DataType: "cpf",                         // Datatypes are identified at datatypes folder
			Writers:  []string{`org1MSP`, "orgMSP"}, // This means only org1 can create the asset (others can edit)
		},
		{
			// Mandatory property
			Required: true,
			Tag:      "name",
			Label:    "Nome da pessoa",
			DataType: "string",
			// Validate funcion
			Validate: func(name interface{}) error {
				nameStr := name.(string)
				if nameStr == "" {
					return fmt.Errorf("o nome não deve estar vazio")
				}
				return nil
			},
		},
		{
			Tag:      "address",
			Label:    "Endereço",
			DataType: "string",
		},
		{
			Tag:      "city",
			Label:    "Cidade",
			DataType: "string",
		},
		{
			Tag:      "estate",
			Label:    "Estado",
			DataType: "string",
		},
	},
}
