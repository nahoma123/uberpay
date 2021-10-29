package tests

import (
	"fmt"
	"log"
	"os"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	authModel = "rbac_model.conf"
)

// func TestCasbinEnforcer() *casbin.Enforcer {
// 	initiator.NewEnforcer()
// }

// func TestUtils() {
// 	dbUrl := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", "localhost", "admin", "", "ride_plus", 26257)
// 	initiator.GetUtils(dbUrl, authModel)
// }

func TestDb() *gorm.DB {
	dbUrl := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", "localhost", "admin", "", "ride_plus", 26257)
	fmt.Println("DbUrl", dbUrl)
	conn, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{
		SkipDefaultTransaction: true, //30% performance increases
	})
	if err != nil {
		log.Printf("Error when Opening database connection: %v", err)
		os.Exit(1)
	}

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path) // for example /home/user

	return conn
}

// NewEnforcer creates an enforcer via file or DB.
func TestEnforcer(conn *gorm.DB) *casbin.Enforcer {
	adapter, err := gormadapter.NewAdapterByDB(conn)
	if err != nil {
		log.Fatal("error ", err)
	}

	enforcer, err := casbin.NewEnforcer(authModel, adapter)
	if err != nil {
		log.Fatal("error ", err)
	}

	enforcer.EnableAutoSave(true)
	enforcer.LoadPolicy()
	return enforcer
}
