package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nonchan7720/user-flex-feature/pkg/interfaces/api/gateway"
	middleware "github.com/oapi-codegen/gin-middleware"
	"github.com/samber/do"
)

func ProvideServer(i *do.Injector) (ServerInterface, error) {
	api := do.MustInvoke[API](i)
	userFlexFeatureGateway := do.MustInvokeNamed[*gateway.Gateway](i, "user-flex-feature-gateway")
	engine := do.MustInvoke[*gin.Engine](i)
	return newServer(engine, api, userFlexFeatureGateway), nil
}

func newServer(r *gin.Engine, api API, userFlexFeatureGateway *gateway.Gateway) *server {
	server := &server{
		api:                    api,
		userFlexFeatureGateway: userFlexFeatureGateway,
	}
	swagger, err := GetSwagger()
	if err != nil {
		panic(err)
	}
	swagger.Servers = nil
	swagger.Security = nil
	r.Use(middleware.OapiRequestValidator(swagger))
	RegisterHandlers(r, server)
	return server
}
