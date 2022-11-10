package cluster

import (
	"fmt"
	"net/http"

	"g.hz.netease.com/horizon/core/common"
	"g.hz.netease.com/horizon/pkg/server/route"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes register routes
func RegisterRoutes(engine *gin.Engine, api *API) {
	apiGroup := engine.Group("/apis/core/v1")
	var routes = route.Routes{
		{
			Method:      http.MethodPost,
			Pattern:     fmt.Sprintf("/applications/:%v/clusters", common.ParamApplicationID),
			HandlerFunc: api.Create,
		}, {
			Method:      http.MethodGet,
			Pattern:     fmt.Sprintf("/applications/:%v/clusters", common.ParamApplicationID),
			HandlerFunc: api.List,
		}, {
			Method:      http.MethodPut,
			Pattern:     fmt.Sprintf("/clusters/:%v", common.ParamClusterID),
			HandlerFunc: api.Update,
		}, {
			Method:      http.MethodGet,
			Pattern:     fmt.Sprintf("/clusters/:%v", common.ParamClusterID),
			HandlerFunc: api.Get,
		}, {
			Method:      http.MethodDelete,
			Pattern:     fmt.Sprintf("/clusters/:%v", common.ParamClusterID),
			HandlerFunc: api.Delete,
		}, {
			Method:      http.MethodPost,
			Pattern:     fmt.Sprintf("/clusters/:%v/builddeploy", common.ParamClusterID),
			HandlerFunc: api.BuildDeploy,
		}, {
			Method:      http.MethodGet,
			Pattern:     fmt.Sprintf("/clusters/:%v/diffs", common.ParamClusterID),
			HandlerFunc: api.GetDiff,
		}, {
			Method:      http.MethodGet,
			Pattern:     fmt.Sprintf("/clusters/:%v/status", common.ParamClusterID),
			HandlerFunc: api.ClusterStatus,
		}, {
			Method:      http.MethodPost,
			Pattern:     fmt.Sprintf("/clusters/:%v/restart", common.ParamClusterID),
			HandlerFunc: api.Restart,
		}, {
			Method:      http.MethodPost,
			Pattern:     fmt.Sprintf("/clusters/:%v/deploy", common.ParamClusterID),
			HandlerFunc: api.Deploy,
		}, {
			Method:      http.MethodPost,
			Pattern:     fmt.Sprintf("/clusters/:%v/rollback", common.ParamClusterID),
			HandlerFunc: api.Rollback,
		}, {
			Method:      http.MethodPost,
			Pattern:     fmt.Sprintf("/clusters/:%v/next", common.ParamClusterID),
			HandlerFunc: api.Next,
		}, {
			Method:      http.MethodGet,
			Pattern:     fmt.Sprintf("/clusters/:%v/containerlog", common.ParamClusterID),
			HandlerFunc: api.GetContainerLog,
		}, {
			Method:      http.MethodPost,
			Pattern:     fmt.Sprintf("/clusters/:%v/online", common.ParamClusterID),
			HandlerFunc: api.Online,
		}, {
			Method:      http.MethodPost,
			Pattern:     fmt.Sprintf("/clusters/:%v/offline", common.ParamClusterID),
			HandlerFunc: api.Offline,
		}, {
			Method:      http.MethodGet,
			Pattern:     fmt.Sprintf("/clusters/:%v/dashboards", common.ParamClusterID),
			HandlerFunc: api.GetGrafanaDashBoard,
		}, {
			Method:      http.MethodGet,
			Pattern:     fmt.Sprintf("/clusters/:%v/pod", common.ParamClusterID),
			HandlerFunc: api.GetClusterPod,
		}, {
			Method:      http.MethodDelete,
			Pattern:     fmt.Sprintf("/clusters/:%v/pods", common.ParamClusterID),
			HandlerFunc: api.DeleteClusterPods,
		}, {
			Method:      http.MethodPost,
			Pattern:     fmt.Sprintf("/clusters/:%v/free", common.ParamClusterID),
			HandlerFunc: api.Free,
		}, {
			Method:      http.MethodGet,
			Pattern:     fmt.Sprintf("/clusters/:%v/events", common.ParamClusterID),
			HandlerFunc: api.PodEvents,
		}, {
			Method:      http.MethodGet,
			Pattern:     fmt.Sprintf("/clusters/:%v/outputs", common.ParamClusterID),
			HandlerFunc: api.GetOutput,
		}, {
			Method:      http.MethodPost,
			Pattern:     fmt.Sprintf("/clusters/:%v/promote", common.ParamClusterID),
			HandlerFunc: api.Promote,
		}, {
			Method:      http.MethodPost,
			Pattern:     fmt.Sprintf("/clusters/:%v/pause", common.ParamClusterID),
			HandlerFunc: api.Pause,
		}, {
			Method:      http.MethodPost,
			Pattern:     fmt.Sprintf("/clusters/:%v/resume", common.ParamClusterID),
			HandlerFunc: api.Resume,
		}, {
			Method:      http.MethodGet,
			Pattern:     fmt.Sprintf("/clusters/:%v/containers", common.ParamClusterID),
			HandlerFunc: api.GetContainers,
		},
	}

	frontGroup := engine.Group("/apis/front/v1/clusters")
	var frontRoutes = route.Routes{
		{
			Method:      http.MethodGet,
			Pattern:     "/searchclusters",
			HandlerFunc: api.ListByNameFuzzily,
		},
		{
			Method:      http.MethodGet,
			Pattern:     "/searchmyclusters",
			HandlerFunc: api.ListUserClusterByNameFuzzily,
		},
		{
			Method:      http.MethodGet,
			Pattern:     fmt.Sprintf("/:%v", common.ParamClusterName),
			HandlerFunc: api.GetByName,
		},
	}

	internalGroup := engine.Group("/apis/internal/v1/clusters")
	var internalRoutes = route.Routes{
		{
			Method:      http.MethodPost,
			Pattern:     fmt.Sprintf("/:%v/deploy", common.ParamClusterID),
			HandlerFunc: api.InternalDeploy,
		},
		{
			Method:      http.MethodPost,
			Pattern:     fmt.Sprintf("/:%v/deploy/:%v", common.ParamClusterID, common.ParamPipelinerunID),
			HandlerFunc: api.InternalDeploy,
		},
	}
	// TODO use middleware to auth token
	internalGroup.Use()

	route.RegisterRoutes(apiGroup, routes)
	route.RegisterRoutes(frontGroup, frontRoutes)
	route.RegisterRoutes(internalGroup, internalRoutes)
}
