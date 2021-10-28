package initiator

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Initialize() {

	common, err := GetUtils()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	// router.Use(corsMW())

	v1 := router.Group("/v1")

	// initialize domains
	AuthInit(common, v1)
	InitNotification(common, v1)
	CompUserInit(common, v1)
	router.Run(":" + os.Getenv("SERVER_PORT"))

	logrus.WithFields(logrus.Fields{
		"host": os.Getenv("DB_HOST"),
		"port": ":" + os.Getenv("SERVER_PORT"),
	}).Info("Starts Serving on HTTP")

	log.Fatal(http.ListenAndServe(os.Getenv("DB_HOST")+":"+os.Getenv("SERVER_PORT"), router))
}
