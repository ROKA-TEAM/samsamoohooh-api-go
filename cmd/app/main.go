package main

import (
	"fmt"
	"log"
	"samsamoohooh-go-api/internal/infra/config"
)

func main() {
	cfg, err := config.NewConfig(".toml")
	if err != nil {
		log.Panicf("failed to load config: %v\n", err)
	}

	fmt.Printf("cfg: %+v\n", cfg)
}
