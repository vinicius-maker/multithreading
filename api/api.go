package api

import (
	"encoding/json"
	"fmt"
	"github.com/vinicius-maker/multithreading/types"
	"github.com/vinicius-maker/multithreading/utils"
)

func BuscaViaCepApi(cep string, canal chan types.Resultado) {
	brasilApiUrl := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	origem := "viacep"

	body, err := utils.RequestApi(brasilApiUrl, origem)
	if err != nil {
		canal <- types.Resultado{Retorno: err, Origem: origem}
	}

	var viaCepApi types.ViaCepApi
	if err := json.Unmarshal(body, &viaCepApi); err != nil {
		canal <- types.Resultado{Retorno: err, Origem: origem}
	}

	canal <- types.Resultado{Retorno: viaCepApi, Origem: origem}
}

func BuscaBrasilCepApi(cep string, canal chan types.Resultado) {
	viaCepUrl := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	origem := "brasilapi"

	body, err := utils.RequestApi(viaCepUrl, origem)
	if err != nil {
		canal <- types.Resultado{Retorno: err, Origem: origem}
	}

	var brasilCepApi types.BrasilCepApi
	if err := json.Unmarshal(body, &brasilCepApi); err != nil {
		canal <- types.Resultado{Retorno: err, Origem: origem}
	}

	canal <- types.Resultado{Retorno: brasilCepApi, Origem: origem}
}
