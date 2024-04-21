package controller

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
)

func New(r *gin.Engine) {
	server := &server{}
	swagger, err := GetSwagger()
	if err != nil {
		panic(err)
	}
	swagger.Servers = nil
	swagger.Security = nil
	r.Use(middleware.OapiRequestValidator(swagger))
	RegisterHandlers(r, server)
}
