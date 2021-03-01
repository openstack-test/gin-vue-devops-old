package service

import (
	"gin-vue-devops/config"
	"gin-vue-devops/global"
	"gin-vue-devops/model"
	"gin-vue-devops/utils"

	"go.uber.org/zap"
)

//@function: GetSystemConfig
//@description: 读取配置文件
//@return: err error, conf config.Server

func GetSystemConfig() (err error, conf config.Server) {
	return nil, global.GVA_CONFIG
}

// @description   set system config,
//@function: SetSystemConfig
//@description: 设置配置文件
//@param: system model.System
//@return: err error

func SetSystemConfig(system model.System) (err error) {
	cs := utils.StructToMap(system.Config)
	for k, v := range cs {
		global.GVA_VP.Set(k, v)
	}
	err = global.GVA_VP.WriteConfig()
	return err
}

//@function: GetServerInfo
//@description: 获取服务器信息
//@return: server *utils.Server, err error

func GetServerInfo() (server *utils.Server, err error) {
	var s utils.Server
	s.Os = utils.InitOS()
	if s.Cpu, err = utils.InitCPU(); err != nil {
		global.GVA_LOG.Error("func utils.InitCPU() Failed!", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Rrm, err = utils.InitRAM(); err != nil {
		global.GVA_LOG.Error("func utils.InitRAM() Failed!", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Disk, err = utils.InitDisk(); err != nil {
		global.GVA_LOG.Error("func utils.InitDisk() Failed!", zap.String("err", err.Error()))
		return &s, err
	}

	return &s, nil
}
