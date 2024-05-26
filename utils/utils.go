package utils

import (
	"fmt"
	"io"
	"net/http"
)

func RequestApi(url, origem string) ([]byte, error) {
	req, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Erro ao buscar o CEP na API %s: %v\n", origem, err)
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("Erro ao ler o corpo da resposta da API %s: %v\n", origem, err)
	}

	return body, nil
}
