package initiator

import (
	"log"
	"os"
	"ride_plus/internal/constant"
	model "ride_plus/internal/constant/model/dbmodels"
	utils "ride_plus/internal/constant/model/init"
	"strconv"
	"time"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/casbin/casbin/v2"
)

func GetUtils() (utils.Utils, error) {
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
		&model.PushedNotification{},
		&model.EmailNotification{},
		&model.SMS{},
		&model.Company{},
		&model.CompanyUser{},
		&model.Image{},
		&model.ImageFormat{},
	)

	trans, validate, err := GetValidation()
	if err != nil {
		log.Fatal("error ", err)
	}

	duration, _ := strconv.Atoi(os.Getenv("timeout"))
	timeoutContext := time.Duration(duration) * time.Second
	enforser := NewEnforcer(conn, authModel)

	return utils.Utils{
		Timeout:     timeoutContext,
		Translator:  trans,
		GoValidator: validate,
		Conn:        conn,
		Enforcer:    enforser,
	}, nil
}

// NewEnforcer creates an enforcer via file or DB.
func NewEnforcer(conn *gorm.DB, model string) *casbin.Enforcer {
	adapter, err := gormadapter.NewAdapterByDB(conn)
	if err != nil {
		log.Fatal("error ", err)
	}

	enforcer, err := casbin.NewEnforcer(model, adapter)
	if err != nil {
		log.Fatal("error ", err)
	}

	enforcer.EnableAutoSave(true)
	enforcer.LoadPolicy()
	return enforcer
}
