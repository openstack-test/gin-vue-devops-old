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

// @Tags K8sCluster
// @Summary 创建K8sCluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sCluster true "创建K8sCluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /K8sCluster/createK8sCluster [post]
func CreateK8sCluster(c *gin.Context) {
	var K8sCluster model.K8sCluster
	_ = c.ShouldBindJSON(&K8sCluster)
	if err := service.CreateK8sCluster(K8sCluster); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags K8sCluster
// @Summary 删除K8sCluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sCluster true "删除K8sCluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /K8sCluster/deleteK8sCluster [delete]
func DeleteK8sCluster(c *gin.Context) {
	var K8sCluster model.K8sCluster
	_ = c.ShouldBindJSON(&K8sCluster)
	if err := service.DeleteK8sCluster(K8sCluster); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags K8sCluster
// @Summary 批量删除K8sCluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除K8sCluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /K8sCluster/deleteK8sClusterByIds [delete]
func DeleteK8sClusterByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteK8sClusterByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags K8sCluster
// @Summary 更新K8sCluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sCluster true "更新K8sCluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /K8sCluster/updateK8sCluster [put]
func UpdateK8sCluster(c *gin.Context) {
	var K8sCluster model.K8sCluster
	_ = c.ShouldBindJSON(&K8sCluster)
	if err := service.UpdateK8sCluster(K8sCluster); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags K8sCluster
// @Summary 用id查询K8sCluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sCluster true "用id查询K8sCluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /K8sCluster/findK8sCluster [get]
func FindK8sCluster(c *gin.Context) {
	var K8sCluster model.K8sCluster
	_ = c.ShouldBindQuery(&K8sCluster)
	if err, reK8sCluster := service.GetK8sCluster(K8sCluster.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reK8sCluster": reK8sCluster}, c)
	}
}

// @Tags K8sCluster
// @Summary 分页获取K8sCluster列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.K8sClusterSearch true "分页获取K8sCluster列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /K8sCluster/getK8sClusterList [get]
func GetK8sClusterList(c *gin.Context) {
	var pageInfo request.K8sClusterSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetK8sClusterInfoList(pageInfo); err != nil {
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
