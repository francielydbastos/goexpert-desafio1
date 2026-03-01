package sql

import (
	"Client-Server-API/server/entities"
)

type CotacaoModel struct {
	ID              int    `gorm:"primaryKey;column:id"`
	Code            string `gorm:"column:code"`
	Codein          string `gorm:"column:codein"`
	Nome            string `gorm:"column:name"`
	Maximo          string `gorm:"column:high"`
	Minimo          string `gorm:"column:low"`
	Variacao        string `gorm:"column:varBid"`
	PorcentVariacao string `gorm:"column:pctChange"`
	Compra          string `gorm:"column:bid"`
	Venda           string `gorm:"column:ask"`
	Timestamp       string `gorm:"column:timestamp"`
	DataCriacao     string `gorm:"column:create_date"`
}

func FromEntity(cotacao entities.Cotacao) *CotacaoModel {
	return &CotacaoModel{
		Code:            cotacao.USDBRL.Code,
		Codein:          cotacao.USDBRL.Codein,
		Nome:            cotacao.USDBRL.Nome,
		Maximo:          cotacao.USDBRL.Maximo,
		Minimo:          cotacao.USDBRL.Minimo,
		Variacao:        cotacao.USDBRL.Variacao,
		PorcentVariacao: cotacao.USDBRL.PorcentVariacao,
		Compra:          cotacao.USDBRL.Compra,
		Venda:           cotacao.USDBRL.Venda,
		Timestamp:       cotacao.USDBRL.Timestamp,
		DataCriacao:     cotacao.USDBRL.DataCriacao,
	}
}
