package external_api

import (
	"Client-Server-API/server/entities"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

const url = "https://economia.awesomeapi.com.br/json/last/USD-BRL"

func GetCambio(ctx context.Context) (*entities.Cotacao, error) {
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("Erro na requisição à API de Câmbio")
	}

	// Leitura e unmarshal do corpo da resposta
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Println(body)
		return nil, errors.New(http.StatusText(resp.StatusCode))
	}

	cotacao := entities.Cotacao{}
	err = json.Unmarshal(body, &cotacao)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &cotacao, nil
}
