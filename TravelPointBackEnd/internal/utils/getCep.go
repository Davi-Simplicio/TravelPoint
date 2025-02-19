package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Coordinates struct {
	LONGITUDE string `json:"longitude"`
	LATITUDE string `json:"latitude"`
}

type Location struct {
	TYPE string `json:"type"`
	COORDINATES Coordinates `json:"coordinates"`
}


type addressModelBrasilApi struct {
	CEP string `json:"cep"`
	STATE string `json:"state"`
	CITY string `json:"city"`
	NEIGHBORHOOD string `json:"neighborhood"`
	STREET string `json:"street"`
	SERVICE string `json:"service"`
	LOCATION Location `json:"location"`
}


func GetInfoByCep(cep string) (addressModelBrasilApi, error) {
	var address addressModelBrasilApi
	fmt.Println("https://brasilapi.com.br/api/cep/v2/"+cep)

	resp,err := http.Get("https://brasilapi.com.br/api/cep/v2/"+cep)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body,err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	if resp.StatusCode != 200 {
		return addressModelBrasilApi{}, fmt.Errorf("Invalid CEP")
	}

	fmt.Println(string(body))
	if err := json.Unmarshal(body, &address); err != nil {
		fmt.Println(err)
	}

	return address, nil
	
}