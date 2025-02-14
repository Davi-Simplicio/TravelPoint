package utils

import (
	"TravelPointbackend/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetInfoByCep(cep string) (models.Address, error) {
	var address models.Address

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
		return models.Address{}, fmt.Errorf("Invalid CEP")
	}

	fmt.Println(string(body))
	if err := json.Unmarshal(body, &address); err != nil {
		fmt.Println(err)
	}

	return address, nil
	
}