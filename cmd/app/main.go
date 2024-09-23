package main

import (
	"log"
	"samsamoohooh-go-api/internal/adapter/persistence/sql/database"
	"samsamoohooh-go-api/internal/infra/config"
)

func main() {

	config, err := config.New(".toml")
	if err != nil {
		log.Panicf("config 생성에 실패하였습니다: %v\n", err)
	}

	db, err := database.NewDatabase(config)
	if err != nil {
		log.Panicf("database 생성에 실패하였습니다: %v", err)
	}

	err = database.AutoMigrate(db)
	if err != nil {
		log.Panicf("migrate에 실패하였습니다: %v", err)
	}

}
