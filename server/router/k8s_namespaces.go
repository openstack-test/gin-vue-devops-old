package router

import (
	"gin-vue-devops/api/v1"
	"gin-vue-devops/middleware"
	"github.com/gin-gonic/gin"
)

func InitK8sNamespacesRouter(Router *gin.RouterGroup) {
	K8sNamespacesRouter := Router.Group("k8sNamespaces").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		K8sNamespacesRouter.POST("createK8sNamespaces", v1.CreateK8sNamespaces)   // 新建K8sNamespaces
		K8sNamespacesRouter.DELETE("deleteK8sNamespaces", v1.DeleteK8sNamespaces) // 删除K8sNamespaces
		K8sNamespacesRouter.DELETE("deleteK8sNamespacesByIds", v1.DeleteK8sNamespacesByIds) // 批量删除K8sNamespaces
		K8sNamespacesRouter.PUT("updateK8sNamespaces", v1.UpdateK8sNamespaces)    // 更新K8sNamespaces
		K8sNamespacesRouter.GET("findK8sNamespaces", v1.FindK8sNamespaces)        // 根据ID获取K8sNamespaces
		K8sNamespacesRouter.GET("getK8sNamespacesList", v1.GetK8sNamespacesList)  // 获取K8sNamespaces列表
	}
}
