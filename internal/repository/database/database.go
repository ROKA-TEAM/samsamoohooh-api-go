package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"samsamoohooh-go-api/internal/domain"
	"samsamoohooh-go-api/internal/infra/config"
)

type Database struct {
	*gorm.DB
}

func NewDatabase(config *config.Config) (*Database, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Database{db}, nil
}

func (d Database) AutoMigration() error {
	return d.AutoMigrate(&domain.User{}, &domain.Group{})
}

func (d Database) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
