package main

import (
	"database/sql"
	"github.com/cheojeg/simplebank/api"
	db "github.com/cheojeg/simplebank/db/sqlc"
	"github.com/cheojeg/simplebank/db/util"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)

	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
