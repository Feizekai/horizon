package application

import (
	"fmt"
	"net/http"

	"g.hz.netease.com/horizon/core/common"
	"g.hz.netease.com/horizon/pkg/server/route"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes register routes
func RegisterRoutes(engine *gin.Engine, api *API) {
	apiV2Group := engine.Group("/apis/core/v2")
	apiV2Routes := route.Routes{
		{
			Method:      http.MethodPost,
			Pattern:     fmt.Sprintf("/groups/:%v/applications", common.ParamGroupID),
			HandlerFunc: api.Create,
		},
	}
	route.RegisterRoutes(apiV2Group, apiV2Routes)
}
