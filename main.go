package main

import (
	"Multithreading/model"
	"encoding/json"
	"fmt"

	"io"
	"net/http"
	"time"
)

func main(){

	c1 := make(chan int)
	c2 := make(chan int)

	endereco1 := model.EnderecoCEPBrasil{Cep:"01153000" }
	endereco2 := model.EnderecoViaCEP{Cep:"01153000"}


	go ConsultaCEPBrasilAPI(&endereco1, c1)
	go ConsultaCEPViaCEP(&endereco2, c2)

	select {

	case <-c1:
		fmt.Println("Consultando API BrasilAPI")
		fmt.Println(" Endereço:")
		
		fmt.Printf("  CEP: %s\n  State: %s\n  City: %s\n  Neighborhood: %s\n  Street: %s\n  Service: %s\n " , 
		endereco1.Cep, 
		endereco1.State, 
		endereco1.City, 
		endereco1.Neighborhood, 
		endereco1.Street, 
		endereco1.Service)
		
	case <-c2:
		fmt.Println("Consultando API Via CEP")
		fmt.Println(" Endereço:")
	
		fmt.Printf(
			"  CEP: %s\n  Logradouro: %s\n  Complemento: %s\n  Unidade: %s\n  Bairro: %s\n  Localidade: %s\n  Uf: %s\n  Ibge: %s\n  Gia: %s\n  Ddd: %s \n  Siafi: %s\n",
			endereco2.Cep,
			endereco2.Logradouro,
			endereco2.Complemento,
			endereco2.Unidade,
			endereco2.Bairro,
			endereco2.Localidade,
			endereco2.Uf,
			endereco2.Ibge,
			endereco2.Gia,
			endereco2.Ddd,
			endereco2.Siafi,
		)

	case <-time.After(1 * time.Second) :
		fmt.Println("Cancelado por timeout na request da API CEP!")
			
	}

}




func ConsultaCEPBrasilAPI(endereco *model.EnderecoCEPBrasil, c1 chan int){


	req, err := http.NewRequest("GET", "https://brasilapi.com.br/api/cep/v1/" + endereco.Cep, nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()


	body, err := io.ReadAll(resp.Body)
	if err != nil{
		panic(err)
	}

	err = json.Unmarshal(body, &endereco)
	if err != nil{
		panic(err)
	}

	c1 <- 1

}

func ConsultaCEPViaCEP(endereco *model.EnderecoViaCEP, c2 chan int) {


	var url =  "http://viacep.com.br/ws/" + endereco.Cep + "/json/"


	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil{
		panic(err)
	}

	err = json.Unmarshal(body, &endereco)
	if err != nil{
		panic(err)
	}

	c2 <- 1

}
