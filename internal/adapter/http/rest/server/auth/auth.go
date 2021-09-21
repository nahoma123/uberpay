package auth

import (
	// "encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
	appErr "template/internal/constant/errors"

	"template/internal/constant/model"
	"template/internal/module/auth"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Authorizer(e *casbin.Enforcer) gin.HandlerFunc
	Login(c *gin.Context)

}

type authHandler struct {
	authUseCase auth.UseCase
}

func NewAuthHandler(authUseCase auth.UseCase) AuthHandler {
	return &authHandler{
		authUseCase: authUseCase,
	}
}
var actions = map[string]string{
	"GET": 		"read",
	"POST": 	"create",
	"PUT":  	"update",
	"DELETE": 	"delete",
	"PATCH":    "update",
}
//Authorizer is a middleware for authorization
func (n *authHandler) Authorizer(e *casbin.Enforcer) gin.HandlerFunc {
	log.Println("authorizer")
	return func(c *gin.Context) {
		role := "anonymous"
		token := ExtractToken(c.Request)
		claims, _ := n.authUseCase.GetClaims(token)
        if e!=nil{
			log.Println("e is differenet from n")
		}else{
			log.Println("e  nill")
			c.AbortWithStatus(http.StatusUnauthorized)

		}
		if claims != nil {
			log.Println("----claim")
			// log.Println(json.MarshalIndent(claims,"","  "))
			role = claims.Role

			c.Set("x-userid", claims.Subject)
			c.Set("x-userrole", role)

			if claims.CompanyID != "" {
				c.Set("x-companyid", claims.CompanyID)
			}
		}
		log.Printf("%v %v %v", role, c.Request.URL.Path, c.Request.Method)

		e.LoadPolicy()
		res, err := e.Enforce(role, c.Request.URL.Path, actions[c.Request.Method])
		if err != nil {
			log.Println("Error enforcing the casbin rules", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if res {
			c.Next()
		} else {
			log.Println("No permission")
			c.AbortWithStatus(http.StatusUnauthorized)

		}

	}
}

func (n authHandler) Login(c *gin.Context) {
	authData := &model.User{}
	err := c.BindJSON(authData)
	if err != nil {
		log.Println("Err binding", err)
		c.JSON(http.StatusBadRequest, gin.H{"errors": appErr.NewErrorResponse(appErr.ErrUnknown)})
		return
	}
	loginResponse, err := n.authUseCase.Login(authData.Phone, authData.Password)

	if err != nil {
		if errors.As(err, &appErr.ValErr{}) {
			c.JSON(http.StatusBadRequest, gin.H{"errors": err})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": appErr.NewErrorResponse(err)})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"user": loginResponse})
}
func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("auth")
	if token != "" {
		return token
	}
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}