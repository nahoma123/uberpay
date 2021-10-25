package initiator

import (
	"log"
	"os"
	"template/internal/constant"
	"template/internal/constant/model"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DbInit() *gorm.DB {
	DATABASE_URL, err := constant.DbConnectionString()
	if err != nil {
		log.Fatal("database connection failed!")
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)
	conn, err := gorm.Open(postgres.Open(DATABASE_URL), &gorm.Config{
		SkipDefaultTransaction: true, //30% performance increases
		Logger:                 newLogger,
	})
	if err != nil {
		log.Printf("Error when Opening database connection: %v", err)
		os.Exit(1)
	}
	conn.AutoMigrate(
		&model.User{},
		&model.CasbinRule{},
		&model.PushedNotification{},
		&model.EmailNotification{},
		&model.SMS{},
		&model.Company{},
		&model.CompanyUser{},
		&model.Image{},
		&model.ImageFormat{},
	)
	return conn
}
