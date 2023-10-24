package storage

import (
	"database/sql"
	"log"
	"project-a/util"

	_ "github.com/lib/pq"
)

func NewDatabase() *sql.DB {
	connectionstring, ok := util.GetPostresqlConnectionString()

	if !ok {
		log.Fatal("falha ao abrir conexão com banco de dados")
	}

	db, err := sql.Open("postgres", connectionstring)

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	return db
}
