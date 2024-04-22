package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nonchan7720/user-flex-feature/pkg/container"
	"github.com/nonchan7720/user-flex-feature/pkg/interfaces/api/gateway"
	"github.com/samber/do"
)

func init() {
	do.Provide(container.Injector, ProvideServer)
}

type server struct {
	api                    API
	userFlexFeatureGateway *gateway.Gateway
}

var (
	_ ServerInterface = (*server)(nil)
)

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
	srv.userFlexFeatureGateway.ServeHTTP(c.Writer, c.Request)
}
