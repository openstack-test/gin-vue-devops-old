package core

import (
	"fmt"
	"gin-vue-devops/global"
	"gin-vue-devops/initialize"
	"time"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
	}
	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	欢迎使用 Gin-Vue-DevOps
	默认文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端地址:http://127.0.0.1:8080
`, address)
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
