package main

import (
	"fmt"
	"github.com/vinicius-maker/multithreading/api"
	"github.com/vinicius-maker/multithreading/types"
	"reflect"
	"time"
)

func main() {
	var cep string
	fmt.Print("Por favor, insira o CEP: ")
	if _, err := fmt.Scanf("%s", &cep); err != nil {
		return
	}

	canal := make(chan types.Resultado)

	go api.BuscaViaCepApi(cep, canal)
	go api.BuscaBrasilCepApi(cep, canal)

	select {
	case resultado := <-canal:
		fmt.Printf("Origem: %s\n", resultado.Origem)

		valoresRetorno := reflect.ValueOf(resultado.Retorno)
		tiposRetorno := valoresRetorno.Type()

		for i := 0; i < valoresRetorno.NumField(); i++ {
			fmt.Printf("%s: %v\n", tiposRetorno.Field(i).Name, valoresRetorno.Field(i).Interface())
		}

	case <-time.After(1 * time.Second):
		fmt.Println("Timeout na requisição excedido.")
	}
}
