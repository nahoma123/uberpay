package initiator

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	compHandler "template/internal/adapter/http/rest/server/company"
	"template/internal/adapter/http/rest/server/image"
	usrHandler "template/internal/adapter/http/rest/server/user"
	"template/internal/adapter/repository"
	"template/internal/adapter/storage/persistence/company"
	"template/internal/adapter/storage/persistence/user"
	routing2 "template/internal/glue/routing"
	usrUsecase "template/internal/module/user"

	compUsecase "template/internal/module/company"

	utils "template/internal/constant/model/init"

	"github.com/gin-gonic/gin"
)

func CompUserInit(utils utils.Utils, router *gin.RouterGroup) {
	usrPersistence := user.UserInit(utils)
	usrRepo := repository.UserInit()

	usrUsecase := usrUsecase.Initialize(usrRepo, usrPersistence, utils)

	usrHandler := usrHandler.UserInit(usrUsecase)

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
