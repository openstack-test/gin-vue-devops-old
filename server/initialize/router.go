package initialize

import (
	_ "gin-vue-devops/docs"
	"gin-vue-devops/global"
	"gin-vue-devops/middleware"
	"gin-vue-devops/router"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.StaticFS(global.GVA_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.GVA_LOG.Info("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors())
	global.GVA_LOG.Info("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group("")
	{
		router.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	}
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		router.InitApiRouter(PrivateGroup)                   // 注册功能api路由
		router.InitJwtRouter(PrivateGroup)                   // jwt相关路由
		router.InitUserRouter(PrivateGroup)                  // 注册用户路由
		router.InitMenuRouter(PrivateGroup)                  // 注册menu路由
		router.InitEmailRouter(PrivateGroup)                 // 邮件相关路由
		router.InitSystemRouter(PrivateGroup)                // system相关路由
		router.InitCasbinRouter(PrivateGroup)                // 权限相关路由
		router.InitAutoCodeRouter(PrivateGroup)              // 创建自动化代码
		router.InitAuthorityRouter(PrivateGroup)             // 注册角色路由
<<<<<<< HEAD
		router.InitSysDictionaryRouter(PrivateGroup)         // 字典管理
		router.InitSysDictionaryDetailRouter(PrivateGroup)   // 字典详情管理
		router.InitFileUploadAndDownloadRouter(PrivateGroup) // 文件上传下载功能路由
=======
		router.InitFileUploadAndDownloadRouter(PrivateGroup) // 文件上传下载功能路由
<<<<<<< HEAD
		router.InitK8sNamespacesRouter(PrivateGroup)         // k8s命名空间路由
		router.InitK8sDeploymentRouter(PrivateGroup)         // k8s应用路由
		router.InitK8sPodsRouter(PrivateGroup)               // k8s pod路由
>>>>>>> develop
=======
		router.InitK8sNamespacesRouter(PrivateGroup)         // k8s命名空间管理路由
		router.InitK8sDeploymentRouter(PrivateGroup)         // k8s应用管理路由
		router.InitK8sPodsRouter(PrivateGroup)               // k8s容器管理路由
		router.InitK8sClusterRouter(PrivateGroup)            // k8s集群管理路由
>>>>>>> develop
	}
	global.GVA_LOG.Info("router register success")
	return Router
}
