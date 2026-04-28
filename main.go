package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Endereco struct {
	Cep        string
	Logradouro string
	Bairro     string
	Uf         string
	Estado     string
	Cidade     string
}

func main() {
	cep := "01000-000" //Coloque aqui seu CEP
	rch1 := make(chan Endereco)
	rch2 := make(chan Endereco)

	go GetViaCEP(cep, rch1)
	go GetBrasilAPI(cep, rch2)

	select {
	case endereco := <-rch1:
		fmt.Printf("Endereço encontrado API: ViaCEP: %+v\n", endereco)
	case endereco := <-rch2:
		fmt.Printf("Endereço encontrado API: BrasilAPI: %+v\n", endereco)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout: Nenhum endereço encontrado")
	}
}

type ViaCEPResponse struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Uf         string `json:"uf"`
	Cidade     string `json:"localidade"`
}

func GetViaCEP(cep string, rch chan<- Endereco) {
	uri := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: received status code %d\n", resp.StatusCode)
		return
	}

	var viaCEPResponse ViaCEPResponse
	err = json.NewDecoder(resp.Body).Decode(&viaCEPResponse)
	if err != nil {
		fmt.Printf("Error decoding response: %v\n", err)
		return
	}

	endereco := Endereco{
		Cep:        viaCEPResponse.Cep,
		Logradouro: viaCEPResponse.Logradouro,
		Bairro:     viaCEPResponse.Bairro,
		Estado:     viaCEPResponse.Uf,
		Cidade:     viaCEPResponse.Cidade,
	}

	rch <- endereco
}

type BrasilAPIResponse struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"street"`
	Bairro     string `json:"neighborhood"`
	Estado     string `json:"state"`
	Cidade     string `json:"city"`
}

func GetBrasilAPI(cep string, rch chan<- Endereco) {
	uri := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: received status code %d\n", resp.StatusCode)
		return
	}

	var brasilAPIResponse BrasilAPIResponse
	err = json.NewDecoder(resp.Body).Decode(&brasilAPIResponse)
	if err != nil {
		fmt.Printf("Error decoding response: %v\n", err)
		return
	}

	endereco := Endereco{
		Cep:        brasilAPIResponse.Cep,
		Logradouro: brasilAPIResponse.Logradouro,
		Bairro:     brasilAPIResponse.Bairro,
		Estado:     brasilAPIResponse.Estado,
		Cidade:     brasilAPIResponse.Cidade,
	}

	rch <- endereco
}
