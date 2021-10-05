package user

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"os"
	"strings"
	"template/internal/adapter/http/rest/server"
	"template/internal/constant"
	custErr "template/internal/constant/errors"
	"template/internal/constant/model"
	"template/internal/module"
)


// userHandler defines all the things neccessary for users handlers
type userHandler struct {
	userUsecase  module.UserUsecase
}

//UserInit initializes a company handler for the domin company
func UserInit(urs module.UserUsecase) server.UserHandler {
	return userHandler{
		userUsecase: urs,
	}
}
func (com userHandler) MiddleWareValidateUser(c *gin.Context) {
	user := &model.Company{}
	err := c.Bind(user)
	if err != nil {
		constant.ResponseJson(c, custErr.NewErrorResponse(custErr.ErrorUnableToBindJsonToStruct), http.StatusBadRequest)
		return
	}
	c.Set("x-user", *user)
	c.Next()
}
func (uh userHandler) UserByID(c *gin.Context) {
	id, err := uuid.FromString(c.Param("user-id"))
	if err != nil {
		constant.ResponseJson(c, custErr.ConvertionError(), http.StatusBadRequest)
	}
	usr := model.User{ID: id}
	ctx := c.Request.Context()
	successData, err := uh.userUsecase.UserByID(ctx, usr)
	if err != nil {
		nrr := custErr.NewErrorResponse(err)
		constant.ResponseJson(c, nrr, http.StatusBadRequest)
		return
	}
	constant.ResponseJson(c, successData, http.StatusOK)
}

func (uh userHandler) Users(c *gin.Context) {
	ctx := c.Request.Context()
	successData, err := uh.userUsecase.Users(ctx)
	if err != nil {
		nrr := custErr.NewErrorResponse(err)
		constant.ResponseJson(c, nrr, http.StatusBadRequest)
		return
	}
	constant.ResponseJson(c, successData, http.StatusOK)
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
			constant.ResponseJson(c, errValue, http.StatusBadRequest)
			return
		}
		constant.ResponseJson(c, custErr.NewErrorResponse(err), custErr.ErrCodes[err])
		return
	}
	constant.ResponseJson(c, *successData, http.StatusOK)
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
			constant.ResponseJson(c, errValue, http.StatusBadRequest)
			return
		}
		constant.ResponseJson(c, custErr.NewErrorResponse(err), custErr.ErrCodes[err])
		return
	}
	constant.ResponseJson(c, *successData, http.StatusOK)
}
func (uh userHandler) AddUserToCompany(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := uuid.FromString(c.Param("user-id"))
	if err != nil {
		constant.ResponseJson(c, custErr.ConvertionError(), http.StatusBadRequest)
	}
	usr := model.User{ID: id}
	user, err := uh.userUsecase.UserByID(ctx, usr)
	if err != nil {
		nrr := custErr.NewErrorResponse(err)
		constant.ResponseJson(c, nrr, http.StatusBadRequest)
		return
	}
	companyID, err := uuid.FromString(c.GetString("x-companyid"))
	if err != nil {
		constant.ResponseJson(c, custErr.ConvertionError(), http.StatusBadRequest)
		return
	}
	cmpu := model.CompanyUser{
		UserID:    user.ID,
		CompanyID: companyID,
		Role:      usr.RoleName,
	}
	err = uh.userUsecase.AddUserToCompany(ctx, cmpu)
	if err != nil {
		constant.ResponseJson(c, custErr.NewErrorResponse(err), http.StatusBadRequest)
		return
	}
	constant.ResponseJson(c, "New User added to company", http.StatusOK)
}

func (uh userHandler) RemoveUserFromCompany(c *gin.Context) {
	id, err := uuid.FromString(c.Param("user-id"))
	if err != nil {
		constant.ResponseJson(c, custErr.ConvertionError(), http.StatusBadRequest)
	}
	usr := model.User{ID: id}
	ctx := c.Request.Context()
	err = uh.userUsecase.RemoveUser(ctx, model.CompanyUser{UserID: usr.ID})
	if err != nil {
		nrr := custErr.NewErrorResponse(err)
		constant.ResponseJson(c, nrr, http.StatusBadRequest)
		return
	}
	constant.ResponseJson(c, "User Remove from company database", http.StatusOK)
}

func (uh userHandler) GetCompanyUsers(c *gin.Context) {
	ctx := c.Request.Context()
	companyID, err := uuid.FromString(c.GetString("x-companyid"))
	if err != nil {
		constant.ResponseJson(c, custErr.ConvertionError(), http.StatusBadRequest)
		return
	}
	successData, err := uh.userUsecase.GetCompanyUsers(ctx, companyID)
	if err != nil {
		nrr := custErr.NewErrorResponse(err)
		constant.ResponseJson(c, nrr, http.StatusBadRequest)
		return
	}
	constant.ResponseJson(c, successData, http.StatusOK)
}
func (uh userHandler) GetCompanyUserByID(c *gin.Context) {
	id, err := uuid.FromString(c.Param("user-id"))
	if err != nil {
		constant.ResponseJson(c, custErr.ConvertionError(), http.StatusBadRequest)
	}
	usr := model.User{ID: id}
	ctx := c.Request.Context()
	successData, err := uh.userUsecase.GetCompanyUserByID(ctx, usr.ID)
	if err != nil {
		nrr := custErr.NewErrorResponse(err)
		constant.ResponseJson(c, nrr, http.StatusBadRequest)
		return
	}
	constant.ResponseJson(c, successData, http.StatusOK)
}

func (uh userHandler) DeleteUser(c *gin.Context) {
	id, err := uuid.FromString(c.Param("user-id"))
	if err != nil {
		constant.ResponseJson(c, custErr.ConvertionError(), http.StatusBadRequest)
	}
	usr := model.User{ID: id}
	ctx := c.Request.Context()
	err = uh.userUsecase.DeleteUser(ctx, usr)
	if err != nil {
		nrr := custErr.NewErrorResponse(err)
		constant.ResponseJson(c, nrr, http.StatusBadRequest)
		return
	}
	constant.ResponseJson(c, "User deleted", http.StatusOK)
}
