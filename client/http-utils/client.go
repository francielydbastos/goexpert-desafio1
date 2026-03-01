package http_utils

import (
	file_utils "Client-Server-API/client/file-utils"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

const url = "http://localhost:8080/cotacao"

type Cotacao struct {
	Bid string `json:"bid"`
}

func ChamarEndpoint(ctx context.Context) error {
	// Chamando o endpoint de cotacao
	ctx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	// Leitura e unmarshal do corpo da resposta
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	if resp.StatusCode != http.StatusOK {
		log.Println(body)
		return errors.New(http.StatusText(resp.StatusCode))
	}

	var cotacao Cotacao
	err = json.Unmarshal(body, &cotacao)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	// Escrita dos dados no arquivo
	err = file_utils.EscreverArquivo(cotacao.Bid)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
