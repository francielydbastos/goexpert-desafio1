package file_utils

import (
	"log"
	"os"
)

func EscreverArquivo(bid string) error {
	file, err := os.OpenFile("cotacao.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Erro ao criar o arquivo cotacao.txt")
		return err
	}
	defer file.Close()

	_, err = file.WriteString("Dólar: " + bid + "\n")
	if err != nil {
		log.Println("Erro ao inserir o valor no arquivo")
		return err
	}
	log.Println("Cotação escrita com sucesso no arquivo.")
	return nil
}
