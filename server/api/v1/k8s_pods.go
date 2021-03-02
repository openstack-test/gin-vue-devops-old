package v1

import (
	"gin-vue-devops/global"
	"gin-vue-devops/model"
	"gin-vue-devops/model/request"
	"gin-vue-devops/model/response"
	"gin-vue-devops/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
    _ "k8s.io/api/core/v1"
)

// @Tags K8sPods
// @Summary 创建K8sPods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sPods true "创建K8sPods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sPods/createK8sPods [post]
func CreateK8sPods(c *gin.Context) {
	var k8sPods model.K8sPods
	_ = c.ShouldBindJSON(&k8sPods)
	if err := service.CreateK8sPods(k8sPods); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags K8sPods
// @Summary 删除K8sPods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sPods true "删除K8sPods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sPods/deleteK8sPods [delete]
func DeleteK8sPods(c *gin.Context) {
	var k8sPods model.K8sPods
	_ = c.ShouldBindJSON(&k8sPods)
	if err := service.DeleteK8sPods(k8sPods); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags K8sPods
// @Summary 批量删除K8sPods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除K8sPods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /k8sPods/deleteK8sPodsByIds [delete]
func DeleteK8sPodsByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteK8sPodsByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags K8sPods
// @Summary 更新K8sPods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sPods true "更新K8sPods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /k8sPods/updateK8sPods [put]
func UpdateK8sPods(c *gin.Context) {
	var k8sPods model.K8sPods
	_ = c.ShouldBindJSON(&k8sPods)
	if err := service.UpdateK8sPods(k8sPods); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags K8sPods
// @Summary 用id查询K8sPods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sPods true "用id查询K8sPods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /k8sPods/findK8sPods [get]
func FindK8sPods(c *gin.Context) {
	var k8sPods model.K8sPods
	_ = c.ShouldBindQuery(&k8sPods)
	if err, rek8sPods := service.GetK8sPods(uint(k8sPods.ID)); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rek8sPods": rek8sPods}, c)
	}
}

// @Tags K8sPods
// @Summary 分页获取K8sPods列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.K8sPodsSearch true "分页获取K8sPods列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sPods/getK8sPodsList [get]
func GetK8sPodsList(c *gin.Context) {
	var pageInfo request.K8sPodsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	namespace := c.Query("namespace")
	//deployment := c.Query("deployment")
	if err, list, total := service.GetK8sPodsInfoList(namespace, pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
