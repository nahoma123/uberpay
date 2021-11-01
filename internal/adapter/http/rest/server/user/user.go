package user

import (
	"net/http"
	"os"
	"ride_plus/internal/adapter/http/rest/server"
	custErr "ride_plus/internal/constant/errors"
	model "ride_plus/internal/constant/model/dbmodel"
	utils "ride_plus/internal/constant/model/init"
	"ride_plus/internal/constant/rest"
	"ride_plus/internal/module"
	"strings"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// userHandler defines all the things necessary for users handlers
type userHandler struct {
	userUsecase module.UserUsecase
}

//UserInit initializes a company handler for the domin company
func UserInit(urs module.UserUsecase, utils utils.Utils) server.UserHandler {
	return userHandler{
		userUsecase: urs,
	}
}
func (uh userHandler) UserMiddleWare(c *gin.Context) {
	user := &model.User{}
	err := c.Bind(&user)
	if err != nil {
		rest.ErrorResponseJson(c, custErr.ServiceError(custErr.ErrorUnableToBindJsonToStruct), http.StatusBadRequest)
		return
	}

	users, err := uh.userUsecase.Users(c.Request.Context())
	if err != nil {
		nrr := custErr.ServiceError(err)
		rest.ErrorResponseJson(c, nrr, http.StatusBadRequest)
		return
	}
	if len(users) == 0 && err == nil {
		user.Status = "Active"
		user.RoleName = "SUPER-ADMIN"
	} else {
		user.Status = "Pending"
		user.RoleName = "anonymous"
	}
	c.Set("x-user", *user)
	c.Next()
}
func (uh userHandler) UserByID(c *gin.Context) {
	id, err := uuid.FromString(c.Param("user-id"))
	if err != nil {
		rest.ErrorResponseJson(c, custErr.ConvertionError(), http.StatusBadRequest)
		return
	}
	usr := model.User{ID: id}
	ctx := c.Request.Context()
	successData, err := uh.userUsecase.UserByID(ctx, usr)
	if err != nil {
		nrr := custErr.ServiceError(err)
		rest.ErrorResponseJson(c, nrr, http.StatusBadRequest)
		return
	}
	rest.ErrorResponseJson(c, successData, http.StatusOK)
}

func (uh userHandler) Users(c *gin.Context) {
	ctx := c.Request.Context()
	successData, err := uh.userUsecase.Users(ctx)
	if err != nil {
		nrr := custErr.ServiceError(err)
		rest.ErrorResponseJson(c, nrr, http.StatusBadRequest)
		return
	}
	rest.ErrorResponseJson(c, successData, http.StatusOK)
}

func (uh userHandler) UpdateUser(c *gin.Context) {
	usr := c.MustGet("x-user").(model.User)
	ctx := c.Request.Context()
	successData, err := uh.userUsecase.UpdateUser(ctx, usr)
	if err != nil {
		if strings.Contains(err.Error(), os.Getenv("ErrSecretKey")) {
			e := strings.Replace(err.Error(), os.Getenv("ErrSecretKey"), "", 1)
			errValue := custErr.ErrorModel{
				ErrorCode:        custErr.ErrCodes[custErr.ErrInvalidField],
				ErrorDescription: custErr.Descriptions[custErr.ErrInvalidField],
				ErrorMessage:     e,
			}
			rest.ErrorResponseJson(c, errValue, http.StatusBadRequest)
			return
		}
		rest.ErrorResponseJson(c, custErr.ServiceError(err), custErr.ErrCodes[err])
		return
	}
	rest.ErrorResponseJson(c, *successData, http.StatusOK)
}

func (uh userHandler) StoreUser(c *gin.Context) {
	usr := c.MustGet("x-user").(model.User)
	ctx := c.Request.Context()
	successData, err := uh.userUsecase.StoreUser(ctx, usr)
	if err != nil {
		if strings.Contains(err.Error(), os.Getenv("ErrSecretKey")) {
			e := strings.Replace(err.Error(), os.Getenv("ErrSecretKey"), "", 1)
			errValue := custErr.ErrorModel{
				ErrorCode:        custErr.ErrCodes[custErr.ErrInvalidField],
				ErrorDescription: custErr.Descriptions[custErr.ErrInvalidField],
				ErrorMessage:     e,
			}
			rest.ErrorResponseJson(c, errValue, http.StatusBadRequest)
			return
		}
		rest.ErrorResponseJson(c, custErr.ServiceError(err), http.StatusBadRequest)
		return
	}
	rest.ErrorResponseJson(c, *successData, http.StatusOK)
}

func (uh userHandler) RegisterUser(c *gin.Context) {
	usr := c.MustGet("x-user").(model.User)
	ctx := c.Request.Context()
	successData, err := uh.userUsecase.StoreUser(ctx, usr)
	if err != nil {
		if strings.Contains(err.Error(), os.Getenv("ErrSecretKey")) {
			e := strings.Replace(err.Error(), os.Getenv("ErrSecretKey"), "", 1)
			errValue := custErr.ErrorModel{
				ErrorCode:        custErr.ErrCodes[custErr.ErrInvalidField],
				ErrorDescription: custErr.Descriptions[custErr.ErrInvalidField],
				ErrorMessage:     e,
			}
			rest.ErrorResponseJson(c, errValue, http.StatusBadRequest)
			return
		}
		rest.ErrorResponseJson(c, custErr.ServiceError(err), http.StatusBadRequest)
		return
	}
	rest.ErrorResponseJson(c, *successData, http.StatusOK)
}

func (uh userHandler) AddUserToCompany(c *gin.Context) {
	ctx := c.Request.Context()
	company_user := model.CompanyUser{}
	err := c.Bind(&company_user)
	if err != nil {
		rest.ErrorResponseJson(c, custErr.ServiceError(custErr.ErrorUnableToBindJsonToStruct), http.StatusBadRequest)
		return
	}
	usr := model.User{ID: company_user.UserID}
	isExist, err := uh.userUsecase.UserExists(ctx, usr)
	if err != nil || isExist == false {
		nrr := custErr.ServiceError(err)
		rest.ErrorResponseJson(c, nrr, http.StatusBadRequest)
		return
	}
	err = uh.userUsecase.AddUserToCompany(ctx, company_user)
	if err != nil {
		rest.ErrorResponseJson(c, custErr.ServiceError(err), http.StatusBadRequest)
		return
	}
	rest.ErrorResponseJson(c, "New User added to company", http.StatusOK)
}

func (uh userHandler) RemoveUserFromCompany(c *gin.Context) {
	id, err := uuid.FromString(c.Param("user-id"))
	if err != nil {
		rest.ErrorResponseJson(c, custErr.ConvertionError(), http.StatusBadRequest)
		return
	}
	usr := model.User{ID: id}
	ctx := c.Request.Context()
	err = uh.userUsecase.RemoveUser(ctx, model.CompanyUser{UserID: usr.ID})
	if err != nil {
		nrr := custErr.ServiceError(err)
		rest.ErrorResponseJson(c, nrr, http.StatusBadRequest)
		return
	}
	rest.ErrorResponseJson(c, "User Remove from company database", http.StatusOK)
}

func (uh userHandler) GetCompanyUsers(c *gin.Context) {
	ctx := c.Request.Context()
	companyID, err := uuid.FromString(c.GetString("x-companyid"))
	if err != nil {
		rest.ErrorResponseJson(c, custErr.ConvertionError(), http.StatusBadRequest)
		return
	}
	company_users, err := uh.userUsecase.GetCompanyUsers(ctx, companyID)
	if err != nil {
		nrr := custErr.ServiceError(err)
		rest.ErrorResponseJson(c, nrr, http.StatusBadRequest)
		return
	}
	users := []*model.User{}
	for _, user := range company_users {
		u := model.User{}
		u.ID = user.UserID
		usr, err := uh.userUsecase.UserByID(ctx, u)
		if err != nil {
			nrr := custErr.ServiceError(err)
			rest.ErrorResponseJson(c, nrr, http.StatusBadRequest)
			return
		}
		users = append(users, usr)
	}
	rest.ErrorResponseJson(c, users, http.StatusOK)
}
func (uh userHandler) GetCompanyUserByID(c *gin.Context) {
	id, err := uuid.FromString(c.Param("user-id"))
	if err != nil {
		rest.ErrorResponseJson(c, custErr.ConvertionError(), http.StatusBadRequest)
		return
	}
	ctx := c.Request.Context()
	user, err := uh.userUsecase.GetCompanyUserByID(ctx, id)
	u := model.User{}
	u.ID = user.UserID
	usr, err := uh.userUsecase.UserByID(ctx, u)
	if err != nil {
		nrr := custErr.ServiceError(err)
		rest.ErrorResponseJson(c, nrr, http.StatusBadRequest)
		return
	}

	rest.ErrorResponseJson(c, usr, http.StatusOK)
}

func (uh userHandler) DeleteUser(c *gin.Context) {
	id, err := uuid.FromString(c.Param("user-id"))
	if err != nil {
		rest.ErrorResponseJson(c, custErr.ConvertionError(), http.StatusBadRequest)
		return
	}
	usr := model.User{ID: id}
	ctx := c.Request.Context()
	err = uh.userUsecase.DeleteUser(ctx, usr)
	if err != nil {
		nrr := custErr.ServiceError(err)
		rest.ErrorResponseJson(c, nrr, http.StatusBadRequest)
		return
	}
	rest.ErrorResponseJson(c, "User deleted", http.StatusOK)
}
