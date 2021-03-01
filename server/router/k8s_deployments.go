package router

import (
	v1 "gin-vue-devops/api/v1"
	"gin-vue-devops/middleware"

	"github.com/gin-gonic/gin"
)

func InitK8sDeploymentRouter(Router *gin.RouterGroup) {
	K8sDeploymentRouter := Router.Group("k8sDeployments").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		K8sDeploymentRouter.POST("createK8sDeployment", v1.CreateK8sDeployment)             // 新建K8sDeployment
		K8sDeploymentRouter.DELETE("deleteK8sDeployment", v1.DeleteK8sDeployment)           // 删除K8sDeployment
		K8sDeploymentRouter.DELETE("deleteK8sDeploymentByIds", v1.DeleteK8sDeploymentByIds) // 批量删除K8sDeployment
		K8sDeploymentRouter.PUT("updateK8sDeployment", v1.UpdateK8sDeployment)              // 更新K8sDeployment
		K8sDeploymentRouter.GET("findK8sDeployment", v1.FindK8sDeployment)                  // 根据ID获取K8sDeployment
		K8sDeploymentRouter.GET("getK8sDeploymentList", v1.GetK8sDeploymentList)            // 获取K8sDeployment列表
	}
}
