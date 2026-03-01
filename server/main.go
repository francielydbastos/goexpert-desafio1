package main

import (
	http_utils "Client-Server-API/server/http-utils"
	"Client-Server-API/server/sql"
	"log"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

func main() {
	db, err := gorm.Open(sqlite.New(sqlite.Config{
		DriverName: "sqlite",
		DSN:        "goexpert.db",
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao estabelecer conexão com sqlite: ", err)
	}

	db = db.Debug()

	err = sql.Init(db)
	if err != nil {
		log.Fatal("Falha ao estabelecer conexão com sqlite: ", err)
	}

	http.HandleFunc("/cotacao", http_utils.CotacaoHandler)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Falha ao iniciar server: ", err)
	}
}
