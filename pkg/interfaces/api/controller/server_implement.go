package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/nonchan7720/user-flex-feature/pkg/container"
	"github.com/samber/do"
)

func init() {
	do.Provide(container.Injector, ProvideServer)
}

type server struct {
	api           API
	gatewayServer *runtime.ServeMux
}

var (
	_ ServerInterface = (*server)(nil)
)

func ProvideServer(i *do.Injector) (ServerInterface, error) {
	api := do.MustInvoke[API](i)
	gatewayServer := do.MustInvoke[*runtime.ServeMux](i)
	return newServer(api, gatewayServer), nil
}

func newServer(api API, gatewayServer *runtime.ServeMux) *server {
	return &server{
		api:           api,
		gatewayServer: gatewayServer,
	}
}

func (srv *server) GetOfrepV1Configuration(c *gin.Context, params GetOfrepV1ConfigurationParams) {
	resp := srv.api.GetOfrepV1Configuration(c, params)
	var (
		obj    any
		status int
	)
	if resp.JSON200 != nil {
		status, obj = http.StatusOK, resp.JSON200
	} else if resp.JSON500 != nil {
		status, obj = http.StatusInternalServerError, resp.JSON500
	}
	c.JSON(status, obj)
}

func (srv *server) PostOfrepV1EvaluateFlags(c *gin.Context, params PostOfrepV1EvaluateFlagsParams) {
	resp := srv.api.PostOfrepV1EvaluateFlags(c, params)
	var (
		obj    any
		status int
	)
	if resp.JSON200 != nil {
		status, obj = http.StatusOK, resp.JSON200
	} else if resp.JSON400 != nil {
		status, obj = http.StatusBadRequest, resp.JSON400
	} else if resp.JSON500 != nil {
		status, obj = http.StatusInternalServerError, resp.JSON500
	}
	c.JSON(status, obj)
}

func (srv *server) PostOfrepV1EvaluateFlagsKey(c *gin.Context, key string) {
	resp := srv.api.PostOfrepV1EvaluateFlagsKey(c, key)
	var (
		obj    any
		status int
	)
	if resp.JSON200 != nil {
		status, obj = http.StatusOK, resp.JSON200
	} else if resp.JSON400 != nil {
		status, obj = http.StatusBadRequest, resp.JSON400
	} else if resp.JSON404 != nil {
		status, obj = http.StatusNotFound, resp.JSON404
	} else if resp.JSON500 != nil {
		status, obj = http.StatusInternalServerError, resp.JSON500
	}
	c.JSON(status, obj)
}

func (srv *server) UserFlexFeatureServiceRuleUpdate(c *gin.Context, key string) {
	srv.gatewayServer.ServeHTTP(c.Writer, c.Request)
}
