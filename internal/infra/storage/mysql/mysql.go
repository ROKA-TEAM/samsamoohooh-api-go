package mysql

import (
	"context"
	"fmt"
	"samsamoohooh-go-api/internal/infra/config"
	"samsamoohooh-go-api/internal/infra/storage/mysql/ent"
)

type MySQL struct {
	*ent.Client
}

func NewDatabase(config *config.Config) (*MySQL, error) {
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

	return &MySQL{client}, nil
}

func (d *MySQL) AutoMigration(ctx context.Context) error {
	err := d.Client.Schema.Create(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (d *MySQL) Close() error {
	return d.Client.Close()
}
