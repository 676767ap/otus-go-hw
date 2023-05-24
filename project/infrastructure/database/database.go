package database

import (
	"github.com/676767ap/otus-go-hw/project/internal/config"
	"github.com/676767ap/otus-go-hw/project/internal/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg *config.Config) (*gorm.DB, error) {
	dsn := "host=" + cfg.PostgreSQL.Host + " user=" + cfg.PostgreSQL.User + " password=" + cfg.PostgreSQL.Password + " dbname=" + cfg.PostgreSQL.Database +
		" port=" + cfg.PostgreSQL.Port + " sslmode=disable TimeZone=Europe/Moscow"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	migrate(db)
	if cfg.DevMode {
		return db.Debug(), nil
	}

	return db, nil
}

func CloseDatabase(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.Close()
	return nil
}

func migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&entity.Banner{}, &entity.Slot{}, &entity.SocGroup{}, &entity.Stat{})
	if err != nil {
		return err
	}
	return nil
}
