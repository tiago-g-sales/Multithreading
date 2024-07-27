package model

type EnderecoCEPBrasil struct{

	Cep 			string `json:"cep"`
	State      		string `json:"state"`
	City 			string `json:"city"`
	Neighborhood 	string `json:"neighborhood"`
	Street        	string `json:"street"`
	Service  		string `json:"service"`
}


type EnderecoViaCEP struct {
	
	Cep 					string `json:"cep"`
	Logradouro      		string `json:"logradouro"`
	Complemento 			string `json:"complemento"`
	Unidade 				string `json:"unidade"`
	Bairro        			string `json:"bairro"`
	Localidade  			string `json:"localidade"`
	Uf 						string `json:"uf"`
	Ibge		      		string `json:"ibge"`
	Gia 					string `json:"gia"`
	Ddd 					string `json:"ddd"`
	Siafi        			string `json:"siafi"`
}