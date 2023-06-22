package datatypes

import (
	"strings"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
)

var cpf = assets.DataType{
	AcceptedFormats: []string{"string"},
	Parse: func(data interface{}) (string, interface{}, errors.ICCError) {
		cpf, ok := data.(string)
		if !ok {
			return "", nil, errors.NewCCError("?a propriedade deve ser do tipo 'string'", 400)
		}

		cpf = strings.ReplaceAll(cpf, ".", "")
		cpf = strings.ReplaceAll(cpf, "-", "")

		if len(cpf) != 11 {
			return "", nil, errors.NewCCError("?O CPF deve ter 11 dígitos", 400)
		}

		var vd0 int
		for i, d := range cpf {
			if i >= 9 {
				break
			}
			dnum := int(d) - '0'
			vd0 += (10 - i) * dnum
		}
		vd0 = 11 - vd0%11
		if vd0 > 9 {
			vd0 = 0
		}
		if int(cpf[9])-'0' != vd0 {
			return "", nil, errors.NewCCError("?CPF inválido", 400)
		}

		var vd1 int
		for i, d := range cpf {
			if i >= 10 {
				break
			}
			dnum := int(d) - '0'
			vd1 += (11 - i) * dnum
		}
		vd1 = 11 - vd1%11
		if vd1 > 9 {
			vd1 = 0
		}
		if int(cpf[10])-'0' != vd1 {
			return "", nil, errors.NewCCError("?CPF inválido", 400)
		}

		return cpf, cpf, nil
	},
}
