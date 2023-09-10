package main

import (
	"backend-master-class/api"
	db "backend-master-class/db/sqlc"
	"backend-master-class/util"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot open db", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("server start", err)
	}

	if err = server.Start(config.ServerAddress); err != nil {
		log.Fatal("server", err)
	}
}
