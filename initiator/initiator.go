package initiator

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	utils "template/internal/constant/model/init"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func Initialize() {
	// load .env file
	err := godotenv.Load("./../../.env")
	fmt.Println("err ", err, "os host ", os.Getenv("DB_USER"))
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	trans, validate, err := GetValidation()
	if err != nil {
		log.Fatal("error ", err)
	}

	dbConn := DbInit()

	duration, _ := strconv.Atoi(os.Getenv("timeout"))
	timeoutContext := time.Duration(duration) * time.Second

	common := utils.Utils{
		Conn:        dbConn,
		GoValidator: validate,
		Translator:  trans,
		Timeout:     timeoutContext,
	}

	router := gin.Default()
	router.Use(corsMW())

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
