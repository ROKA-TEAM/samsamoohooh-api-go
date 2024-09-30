package provider

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"samsamoohooh-go-api/internal/adapter/persistence/sql/database"
	"sync"
)

type Provider struct {
	*database.Database
}

var instance *Provider
var once sync.Once

func Provide() *Provider {
	once.Do(func() {
		// open mysql
		gormDB, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			panic(err)
		}

		// warp database.Database
		db := &database.Database{DB: gormDB}

		// auto migration
		err = database.AutoMigrate(db)
		if err != nil {
			panic(err)
		}

		instance = &Provider{
			Database: db,
		}
	})

	return instance
}

func (p *Provider) Reset() error {
	if instance == nil {
		return nil
	}

	// gorm으로 모든 테이블 다 삭제
	err := database.DropAllTable(instance.Database)
	if err != nil {
		return err
	}

	// gorm으로 다시 생성
	err = database.AutoMigrate(instance.Database)
	if err != nil {
		return err
	}

	return nil
}
