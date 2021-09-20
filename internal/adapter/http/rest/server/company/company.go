package company

import (
	"errors"
	"net/http"
	appErr "template/internal/constant/errors"
	"template/internal/constant/model"
	"template/internal/module/company"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	uuid "github.com/satori/go.uuid"
)

type CompanyHandler interface {
	Companies(c *gin.Context)
	CreateCompany(c *gin.Context)
	GetCompanyById(c *gin.Context)
	DeleteCompany(c *gin.Context)
}

type companyHandler struct {
	compUsecase company.Usecase
	trans       ut.Translator
}

func CompanyInit(compUsecase company.Usecase, trans ut.Translator) CompanyHandler {
	return &companyHandler{
		compUsecase,
		trans,
	}
}

func (ch companyHandler) Companies(c *gin.Context) {
  companies, err := ch.compUsecase.Companies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": appErr.NewErrorResponse(err)})
		return
	}

	c.JSON(http.StatusOK,  companies)
}
func (ch companyHandler) CreateCompany(c *gin.Context) {
	var insertCompany model.Company
	if err := c.ShouldBind(&insertCompany); err != nil {
		// var verr validator.ValidationErrors

		// if errors.As(err, &verr) {
		// 	c.JSON(http.StatusBadRequest, gin.H{"errors": verr.Translate(ch.trans)})
		// 	return
		// }
		c.JSON(http.StatusBadRequest, gin.H{"errors": appErr.NewErrorResponse(appErr.ErrUnknown)})
		return

	}
	company, err := ch.compUsecase.CreateCompany(&insertCompany)
	if err != nil {
		if errors.As(err, &appErr.ValErr{}) {
			c.JSON(http.StatusBadRequest, gin.H{"errors": err})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": appErr.NewErrorResponse(appErr.ErrUnknown)})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"company": company})

}

func (ch companyHandler) GetCompanyById(c *gin.Context) {
	ID := c.Param("id")

	id, err := uuid.FromString(ID)
	if err != nil {
		// TODO: this error is not in errors package
		c.JSON(http.StatusBadRequest, gin.H{"errors": appErr.NewErrorResponse(err)})
		return
	}

	company, err := ch.compUsecase.GetCompanyById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": appErr.NewErrorResponse(err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"company": company})

}

func (ch companyHandler) DeleteCompany(c *gin.Context) {}
