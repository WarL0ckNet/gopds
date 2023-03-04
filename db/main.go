package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type DB_connect struct {
	Host string
	Name string
	User string
	Pass string
	Port int
}

func ConnectDatabase(cnt *DB_connect) (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Warn,
			Colorful: true,
		},
	)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Novosibirsk", cnt.Host, cnt.User, cnt.Pass, cnt.Name, cnt.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}})
	if err != nil {
		return db, err
	}
	return db, nil
}

func InitDatabase(db *gorm.DB) {
	MigrateAll(db)
	genres := GetGenres()
	var count int64
	for _, genre := range genres {
		db.Model(&Genre{}).Where("genre = ?", genre.Genre).Count(&count)
		if count == 0 {
			res := db.Create(&genre)
			if res.Error != nil {
				log.Fatalf("Create ganre table error: %v", res)
			}
		} else {
			db.Model(&Genre{}).Where("genre = ?", genre.Genre).Updates(&genre)
		}
	}
}

func ClearDatabase(db *gorm.DB) {
	db.Where("1=1").Delete(&Bseries{})
	db.Where("1=1").Delete(&Bauthor{})
	db.Where("1=1").Delete(&Bgenre{})
	db.Where("1=1").Delete(&Book{})
	db.Where("1=1").Delete(&Catalog{})
	db.Where("1=1").Delete(&Author{})
	db.Where("1=1").Delete(&Genre{})
	db.Where("1=1").Delete(&Series{})
}
