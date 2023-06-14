package assettypes

import "github.com/goledgerdev/cc-tools/assets"

// Description of a baggage
var Baggage = assets.AssetType{
	Tag:         "baggage",
	Label:       "Bagagem",
	Description: "Baggage",

	Props: []assets.AssetProp{
		{
			// Composite Key
			Required: true,
			IsKey:    true,
			Tag:      "id",
			Label:    "ID da bagagem",
			DataType: "string",
			Writers:  []string{`org1MSP`, "orgMSP"}, // This means only org1 can create the asset (others can edit)
		},
		{
			/// Reference to another asset
			Tag:      "passenger_id",
			Label:    "ID (CPF) do Passageiro",
			DataType: "->person.id",
		},
		{
			Tag:      "color",
			Label:    "Cor da bagagem",
			DataType: "string",
		},
		{
			Tag:      "weight",
			Label:    "Peso da bagagem",
			DataType: "number",
		},
		{
			Tag:      "size",
			Label:    "Tamanho da bagagem",
			DataType: "string",
		},
	},
}
