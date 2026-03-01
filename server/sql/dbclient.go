package sql

import (
	"Client-Server-API/server/entities"
	"context"
	"log"
	"time"

	"gorm.io/gorm"
)

var db *gorm.DB

func Init(database *gorm.DB) error {
	db = database

	return db.AutoMigrate(&CotacaoModel{})
}

func Create(ctx context.Context, cotacao entities.Cotacao) error {
	cotacaoModel := FromEntity(cotacao)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	transacao := db.WithContext(ctx).Create(cotacaoModel)
	if transacao.Error != nil {
		log.Println("Erro ao criar o registro no banco de dados: ", transacao.Error)
		return transacao.Error
	}
	log.Println("Cotacao inserida no banco de dados com o ID: ", cotacaoModel.ID)
	return nil
}
