package database

import (
	"github.com/676767ap/project/internal/config"
	"github.com/676767ap/project/internal/entity"
	"github.com/676767ap/project/util/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg *config.Config) *gorm.DB {
	dsn := "host=" + cfg.PostgreSQL.Host + " user=" + cfg.PostgreSQL.User + " password=" + cfg.PostgreSQL.Password + " dbname=" + cfg.PostgreSQL.Database +
		" port=" + cfg.PostgreSQL.Port + " sslmode=disable TimeZone=Europe/Moscow"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Open:", err)
	}
	migrate(db)
	if cfg.DevMode {
		return db.Debug()
	}

	return db
}

func CloseDatabase(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("DB Shutdown:", err)
	}
	sqlDB.Close()
}

func migrate(db *gorm.DB) {
	err := db.AutoMigrate(&entity.Banner{}, &entity.Slot{}, &entity.SocGroup{}, &entity.Stat{})
	if err != nil {
		log.Fatal("DB AutoMigrate:", err)
	}
}
