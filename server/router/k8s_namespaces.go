package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitK8sNamespacesRouter(Router *gin.RouterGroup) {
	K8sNamespacesRouter := Router.Group("k8sNamespaces").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		K8sNamespacesRouter.GET("findK8sNamespaces", v1.FindK8sNamespaces)        // 根据ID获取K8sNamespaces
		K8sNamespacesRouter.POST("getK8sNamespacesList", v1.GetK8sNamespacesList)  // 获取K8sNamespaces
	}
}
