package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags K8sNamespaces
// @Summary 用id查询K8sNamespaces
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sNamespaces true "用id查询K8sNamespaces"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /k8sNamespaces/findK8sNamespaces [get]
func FindK8sNamespaces(c *gin.Context) {
	var k8sNamespaces model.K8sNamespaces
	_ = c.ShouldBindQuery(&k8sNamespaces)
	if err, rek8sNamespaces := service.FindK8sNamespaces(k8sNamespaces.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rek8sNamespaces": rek8sNamespaces}, c)
	}
}

// @Tags K8sNamespaces
// @Summary 分页获取K8sNamespaces列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.K8sNamespacesSearch true "分页获取K8sNamespaces列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sNamespaces/getK8sNamespacesList [post]
func GetK8sNamespacesList(c *gin.Context) {
	var pageInfo request.K8sNamespacesSearch
	_ = c.ShouldBindJSON(&pageInfo)
	if err, list, total := service.GetK8sNamespacesList(pageInfo); err != nil {
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
