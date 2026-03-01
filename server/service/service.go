package service

import (
	api "Client-Server-API/server/external-api"
	"Client-Server-API/server/sql"
	"context"
	"log"
)

func ProcessarRequest(ctx context.Context) (string, error) {
	// Chamando a API de cambio
	cotacao, err := api.GetCambio(ctx)
	if err != nil {
		log.Println("Erro ao processar o request:", err)
		return "", err
	}

	log.Println("Cotacao: ", cotacao)

	// Salvando no banco de dados
	err = sql.Create(ctx, *cotacao)
	if err != nil {
		log.Println("Erro ao salvar a cotacao no banco de dados:", err)
		return "", err
	}

	return cotacao.USDBRL.Compra, nil
}
