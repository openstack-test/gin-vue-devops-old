package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitK8sPodsRouter(Router *gin.RouterGroup) {
	K8sPodsRouter := Router.Group("k8sPods").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		K8sPodsRouter.POST("createK8sPods", v1.CreateK8sPods)   // 新建K8sPods
		K8sPodsRouter.DELETE("deleteK8sPods", v1.DeleteK8sPods) // 删除K8sPods
		K8sPodsRouter.DELETE("deleteK8sPodsByIds", v1.DeleteK8sPodsByIds) // 批量删除K8sPods
		K8sPodsRouter.PUT("updateK8sPods", v1.UpdateK8sPods)    // 更新K8sPods
		K8sPodsRouter.GET("findK8sPods", v1.FindK8sPods)        // 根据ID获取K8sPods
		K8sPodsRouter.GET("getK8sPodsList", v1.GetK8sPodsList)  // 获取K8sPods列表
	}
}
