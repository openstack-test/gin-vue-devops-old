package router

import (
	v1 "gin-vue-devops/api/v1"
	"gin-vue-devops/middleware"

	"github.com/gin-gonic/gin"
)

func InitK8sClusterRouter(Router *gin.RouterGroup) {
	K8sClusterRouter := Router.Group("k8sCluster").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		K8sClusterRouter.POST("createK8sCluster", v1.CreateK8sCluster)             // 新建K8sCluster
		K8sClusterRouter.DELETE("deleteK8sCluster", v1.DeleteK8sCluster)           // 删除K8sCluster
		K8sClusterRouter.DELETE("deleteK8sClusterByIds", v1.DeleteK8sClusterByIds) // 批量删除K8sCluster
		K8sClusterRouter.PUT("updateK8sCluster", v1.UpdateK8sCluster)              // 更新K8sCluster
		K8sClusterRouter.GET("findK8sCluster", v1.FindK8sCluster)                  // 根据ID获取K8sCluster
		K8sClusterRouter.GET("getK8sClusterList", v1.GetK8sClusterList)            // 获取K8sCluster列表
	}
}
