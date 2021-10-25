package initiator

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	compHandler "ride_plus/internal/adapter/http/rest/server/company"
	"ride_plus/internal/adapter/http/rest/server/image"
	usrHandler "ride_plus/internal/adapter/http/rest/server/user"
	"ride_plus/internal/adapter/repository"
	"ride_plus/internal/adapter/storage/persistence/company"
	"ride_plus/internal/adapter/storage/persistence/user"
	routing2 "ride_plus/internal/glue/routing"
	usrUsecase "ride_plus/internal/module/user"

	compUsecase "ride_plus/internal/module/company"

	utils "ride_plus/internal/constant/model/init"

	"github.com/gin-gonic/gin"
)

func CompUserInit(utils utils.Utils, router *gin.RouterGroup) {
	usrPersistence := user.UserInit(utils)
	usrRepo := repository.UserInit()

	usrUsecase := usrUsecase.Initialize(usrRepo, usrPersistence, utils)

	usrHandler := usrHandler.UserInit(usrUsecase, utils)

	compPersistence := company.CompanyInit(utils)
	compUsecase := compUsecase.Initialize(compPersistence, utils)
	basePath, err := os.Getwd()

	if err != nil {
		log.Fatalf("cannot get base path: %v", err)
	}

	path := filepath.Dir(basePath)
	path = filepath.Dir(path)
	fmt.Println("path ", path)
	store, err := image.NewStorage(path)
	if err != nil {
		log.Fatalf("cannot create storage: %v", err)
	}
	compHandler := compHandler.CompanyInit(compUsecase, *store)

	routing2.CompanyRoutes(router, compHandler)
	routing2.UserRoutes(router, usrHandler)
}
