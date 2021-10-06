package initiator

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	gomail "gopkg.in/mail.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"os"
	"strconv"
	authHandler "template/internal/adapter/http/rest/server/auth"
	compHandler "template/internal/adapter/http/rest/server/company"
	email3 "template/internal/adapter/http/rest/server/notification/email"
	publisher3 "template/internal/adapter/http/rest/server/notification/publisher"
	sms3 "template/internal/adapter/http/rest/server/notification/sms"
	permHandler "template/internal/adapter/http/rest/server/policy"
	rlHandler "template/internal/adapter/http/rest/server/role"
	usrHandler "template/internal/adapter/http/rest/server/user"
	"template/internal/adapter/repository"
	"template/internal/adapter/storage/persistence/company"
	"template/internal/adapter/storage/persistence/notification/email"
	"template/internal/adapter/storage/persistence/notification/publisher"
	"template/internal/adapter/storage/persistence/notification/sms"
	"template/internal/adapter/storage/persistence/role"
	"template/internal/adapter/storage/persistence/user"
	"template/internal/constant"
	"template/internal/constant/model"
	routing2 "template/internal/glue/routing"
	authUsecase "template/internal/module/auth"
	compUsecase "template/internal/module/company"
	email2 "template/internal/module/notification/email"
	publisher2 "template/internal/module/notification/publisher"
	sms2 "template/internal/module/notification/sms"
	roleUsecase "template/internal/module/role"
	usrUsecase "template/internal/module/user"
	casAuth "template/platform/casbin"
	"time"
)

//global validator instance
var (
	// uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func Initialize() {
	// load .env file
	err := godotenv.Load("./../../.env")
	fmt.Println("err ", err, "os host ", os.Getenv("DB_USER"))
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	en := en.New()
	uni := ut.New(en, en)
	trans, _ = uni.GetTranslator("en")
	validate = validator.New()
	err = en_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		log.Fatal("error ", err)
	}
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
		Logger: newLogger,
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
	)
	a, err := gormadapter.NewAdapterByDBWithCustomTable(conn, &model.CasbinRule{})
	e, err := casbin.NewEnforcer("../../rbac_model.conf", a)
	if err != nil {
		log.Fatal("error ", err)
	}
	duration, _ := strconv.Atoi(os.Getenv("timeout"))
	timeoutContext := time.Duration(duration) * time.Second
	usrPersistence := user.UserInit(conn)
	compPersistence := company.CompanyInit(conn)
	rolePersistent := role.RoleInit(conn)

	//notification persistence
	emailPersistent := email.EmailInit(conn)
	smsPersistent := sms.SmsInit(conn)
	publisherPersistent := publisher.NotificationInit(conn)

	//notification services
	emailUsecase := email2.Initialize(emailPersistent, validate, trans, timeoutContext)
	smsUsecase := sms2.Initialize(smsPersistent, validate, trans, timeoutContext)
	publisherUsecase := publisher2.Initialize(publisherPersistent, validate, trans, timeoutContext)

	//notification handlers
	m := gomail.NewMessage()
	v := validator.New()
	emailHandler := email3.NewEmailHandler(emailUsecase, m)
	smsHandler := sms3.NewSmsHandler(smsUsecase, v)
	publisherHandler := publisher3.NewNotificationHandler(publisherUsecase, v)

	roleUsecase := roleUsecase.RoleInitialize(rolePersistent, validate, trans, timeoutContext)
	roleHandler := rlHandler.NewRoleHandler(roleUsecase, trans, validate)

	usrRepo := repository.UserInit()
	usrUsecase := usrUsecase.Initialize(usrRepo, usrPersistence, validate, trans, timeoutContext)
	usrHandler := usrHandler.UserInit(usrUsecase)

	casbinAuth := casAuth.NewEnforcer(e, validate, trans)
	permHandler := permHandler.PolicyInit(casbinAuth)

	jwtManager := authUsecase.NewJWTManager("secret")
	authUsecases := authUsecase.Initialize(usrPersistence, *jwtManager, timeoutContext)
	authHandlers := authHandler.NewAuthHandler(authUsecases, casbinAuth)

	compUsecase := compUsecase.Initialize(compPersistence, validate, trans, timeoutContext)
	compHandler := compHandler.CompanyInit(compUsecase)

	router := gin.Default()
	router.Use(corsMW())
	router.Use(authHandlers.Authorizer(e))
	v1 := router.Group("/v1")
	routing2.UserRoutes(v1, usrHandler)
	routing2.CompanyRoutes(v1, compHandler)
	routing2.RoleRoutes(v1, roleHandler)
	routing2.PolicyRoutes(v1, permHandler)
	routing2.AuthRoutes(v1, authHandlers)
	//notification
	routing2.EmailRoutes(v1, emailHandler)
	routing2.SmsRoutes(v1, smsHandler)
	routing2.PublisherRoutes(v1, publisherHandler)
	router.Run(":" + os.Getenv("SERVER_PORT"))

	logrus.WithFields(logrus.Fields{
		"host": os.Getenv("DB_HOST"),
		"port": ":" + os.Getenv("SERVER_PORT"),
	}).Info("Starts Serving on HTTP")

	log.Fatal(http.ListenAndServe(os.Getenv("DB_HOST")+":"+os.Getenv("SERVER_PORT"), router))
}
func corsMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, OPTIONS, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}

}
