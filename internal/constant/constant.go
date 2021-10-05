package constant

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"os"
	"strings"
)

const (
	SuperAdmin   = "SUPER-ADMIN"
	CompanyAdmin = "COMPANY-ADMIN"
	CompanyClerk = "COMPANY-CLERK"
)

type SuccessData struct {
	Code int
	Data interface{}
}

func ResponseJson(c *gin.Context, responseData interface{}, statusCode int) {
	c.JSON(statusCode, responseData)
	return
}
func StructValidator(structName interface{}, validate *validator.Validate, trans ut.Translator) error {
	//st:=structName.(model.EmailNotification)
	errV := validate.Struct(structName)
	fmt.Println("v1-err ", errV)
	if errV != nil {
		errs := errV.(validator.ValidationErrors)
		valErr := errs.Translate(trans)
		fmt.Println("v err ", valErr)
		for key, _ := range valErr {
			value := strings.TrimSpace(valErr[key])
			value += " " + os.Getenv("ErrSecretKey")
			fmt.Println("error ", value)
			return errors.New(value)
		}
	}
	return nil
}

func ValidateVariable(parm interface{}, validate *validator.Validate) error {
	errs := validate.Var(parm, "required")
	if errs != nil {
		return errs
	}
	return nil
}
func DbConnectionString() (string, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	addr := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", host, user, password, dbname, port)
	return addr, nil
}
