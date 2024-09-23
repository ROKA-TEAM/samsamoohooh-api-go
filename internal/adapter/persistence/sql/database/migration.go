package database

import "samsamoohooh-go-api/internal/core/domain"

func AutoMigrate(d *Database) error {
	return d.AutoMigrate(&domain.User{}, &domain.Group{}, &domain.Post{}, &domain.Comment{})
}
