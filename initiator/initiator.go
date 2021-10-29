package initiator

import (
	"log"
	"net/http"
	"os"
	"ride_plus/initiator/domain"
	"ride_plus/internal/constant"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	authModel = "config/rbac_model.conf"
)

func Initialize() {
	DATABASE_URL, err := constant.DbConnectionString()
	if err != nil {
		log.Fatal("database connection failed!")
	}

	common, err := GetUtils(DATABASE_URL, authModel)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	// router.Use(corsMW())

	v1 := router.Group("/v1")

	// initialize domains
	domain.AuthInit(common, v1)
	domain.InitNotification(common, v1)
	domain.CompUserInit(common, v1)
	router.Run(":" + os.Getenv("SERVER_PORT"))

	logrus.WithFields(logrus.Fields{
		"host": os.Getenv("DB_HOST"),
		"port": ":" + os.Getenv("SERVER_PORT"),
	}).Info("Starts Serving on HTTP")

	log.Fatal(http.ListenAndServe(os.Getenv("DB_HOST")+":"+os.Getenv("SERVER_PORT"), router))
}
