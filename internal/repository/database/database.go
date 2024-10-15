package database

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"samsamoohooh-go-api/internal/infra/config"
	"samsamoohooh-go-api/internal/repository/database/ent"
)

type Database struct {
	*ent.Client
}

func NewDatabase(config *config.Config) (*Database, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Database,
	)

	client, err := ent.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return &Database{client}, nil
}

func (d Database) AutoMigration(ctx context.Context) error {
	err := d.Client.Schema.Create(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (d Database) Close() error {
	return d.Client.Close()
}
