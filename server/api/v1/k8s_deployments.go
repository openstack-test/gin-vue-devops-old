package v1

import (
	"gin-vue-devops/global"
	"gin-vue-devops/model"
	"gin-vue-devops/model/request"
	"gin-vue-devops/model/response"
	"gin-vue-devops/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags K8sDeployment
// @Summary 创建K8sDeployment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sDeployment true "创建K8sDeployment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sDeployments/createK8sDeployment [post]
func CreateK8sDeployment(c *gin.Context) {
	var k8sDeployments model.K8sDeployment
	_ = c.ShouldBindJSON(&k8sDeployments)
	if err := service.CreateK8sDeployment(k8sDeployments); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags K8sDeployment
// @Summary 删除K8sDeployment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sDeployment true "删除K8sDeployment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sDeployments/deleteK8sDeployment [delete]
func DeleteK8sDeployment(c *gin.Context) {
	var k8sDeployments model.K8sDeployment
	_ = c.ShouldBindJSON(&k8sDeployments)
	if err := service.DeleteK8sDeployment(k8sDeployments); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags K8sDeployment
// @Summary 批量删除K8sDeployment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除K8sDeployment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /k8sDeployments/deleteK8sDeploymentByIds [delete]
func DeleteK8sDeploymentByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteK8sDeploymentByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags K8sDeployment
// @Summary 更新K8sDeployment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sDeployment true "更新K8sDeployment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /k8sDeployments/updateK8sDeployment [put]
func UpdateK8sDeployment(c *gin.Context) {
	var k8sDeployments model.K8sDeployment
	_ = c.ShouldBindJSON(&k8sDeployments)
	if err := service.UpdateK8sDeployment(k8sDeployments); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags K8sDeployment
// @Summary 用id查询K8sDeployment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sDeployment true "用id查询K8sDeployment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /k8sDeployments/findK8sDeployment [get]
func FindK8sDeployment(c *gin.Context) {
	var k8sDeployments model.K8sDeployment
	_ = c.ShouldBindQuery(&k8sDeployments)
	if err, rek8sDeployments := service.GetK8sDeployment(uint(k8sDeployments.ID)); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rek8sDeployments": rek8sDeployments}, c)
	}
}

// @Tags K8sDeployment
// @Summary 分页获取K8sDeployment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.K8sDeploymentSearch true "分页获取K8sDeployment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sDeployments/getK8sDeploymentList [get]
func GetK8sDeploymentList(c *gin.Context) {
	// 指定Cluster ID
	kubeConfig, _ := ClusterID(c)
	var pageInfo request.K8sDeploymentSearch
	_ = c.ShouldBindQuery(&pageInfo)
	namespace := c.Query("namespace")
	if err, list, total := service.GetK8sDeploymentInfoList(kubeConfig, namespace); err != nil {
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
