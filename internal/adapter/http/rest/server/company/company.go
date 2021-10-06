package company

import (
	"fmt"
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



// companyHandler defines all the things necessary for company handlers
type companyHandler struct {
	companyUsecase module.CompanyUsecase
}

//CompanyInit initializes a company handler for the domin company
func CompanyInit(cmp module.CompanyUsecase) server.CompanyHandler {
	return companyHandler{
		companyUsecase: cmp,
	}
}
func (com companyHandler) CompanyMiddleWare(c *gin.Context) {
	compX := &model.Company{}
	err := c.Bind(compX)
	fmt.Println("error bind ",err)
	if err != nil {
		constant.ResponseJson(c, custErr.NewErrorResponse(custErr.ErrorUnableToBindJsonToStruct), http.StatusBadRequest)
		return
	}
	c.Set("x-company", *compX)
	c.Next()
}
func (com companyHandler) CompanyByID(c *gin.Context) {
	id, err := uuid.FromString(c.Param("company-id"))
	if err != nil {
		constant.ResponseJson(c, custErr.ConvertionError(), http.StatusBadRequest)
	}
	comp := model.Company{ID: id}
	ctx := c.Request.Context()
	successData, err := com.companyUsecase.CompanyByID(ctx, comp)
	fmt.Println("error ",err)
	if err != nil {
		nrr := custErr.NewErrorResponse(err)
		constant.ResponseJson(c, nrr, http.StatusBadRequest)
		return
	}
	constant.ResponseJson(c, successData, http.StatusOK)
}

func (com companyHandler) Companies(c *gin.Context) {
	ctx := c.Request.Context()
	successData, err := com.companyUsecase.Companies(ctx)
	if err != nil {
		nrr := custErr.NewErrorResponse(err)
		constant.ResponseJson(c, nrr, http.StatusBadRequest)
		return
	}
	constant.ResponseJson(c, successData, http.StatusOK)
}
func (com companyHandler) StoreCompany(c *gin.Context) {
	comp := c.MustGet("x-company").(model.Company)
	ctx := c.Request.Context()
	successData, err := com.companyUsecase.StoreCompany(ctx, comp)
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

func (com companyHandler) UpdateCompany(c *gin.Context) {
	id, err := uuid.FromString(c.Param("company-id"))
	if err != nil {
		constant.ResponseJson(c, custErr.ConvertionError(), http.StatusBadRequest)
	}
	comp := model.Company{ID: id}
	ctx := c.Request.Context()
	successData, err := com.companyUsecase.UpdateCompany(ctx, comp)
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

func (com companyHandler) DeleteCompany(c *gin.Context) {
	id, err := uuid.FromString(c.Param("company-id"))
	if err != nil {
		constant.ResponseJson(c, custErr.ConvertionError(), http.StatusBadRequest)
	}
	comp := model.Company{ID: id}
	ctx := c.Request.Context()
	err = com.companyUsecase.DeleteCompany(ctx, comp)

	if err != nil {
		nrr := custErr.NewErrorResponse(err)
		constant.ResponseJson(c, nrr, http.StatusBadRequest)
		return
	}
	constant.ResponseJson(c, "company Deleted", http.StatusOK)
}
