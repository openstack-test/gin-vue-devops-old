import service from '@/utils/request'

// @Tags K8sDeployment
// @Summary 创建K8sDeployment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sDeployment true "创建K8sDeployment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sDeployments/createK8sDeployment [post]
export const createK8sDeployment = (data) => {
     return service({
         url: "/k8sDeployments/createK8sDeployment",
         method: 'post',
         data
     })
 }


// @Tags K8sDeployment
// @Summary 删除K8sDeployment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sDeployment true "删除K8sDeployment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sDeployments/deleteK8sDeployment [delete]
 export const deleteK8sDeployment = (data) => {
     return service({
         url: "/k8sDeployments/deleteK8sDeployment",
         method: 'delete',
         data
     })
 }

// @Tags K8sDeployment
// @Summary 删除K8sDeployment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除K8sDeployment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sDeployments/deleteK8sDeployment [delete]
 export const deleteK8sDeploymentByIds = (data) => {
     return service({
         url: "/k8sDeployments/deleteK8sDeploymentByIds",
         method: 'delete',
         data
     })
 }

// @Tags K8sDeployment
// @Summary 更新K8sDeployment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sDeployment true "更新K8sDeployment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /k8sDeployments/updateK8sDeployment [put]
 export const updateK8sDeployment = (data) => {
     return service({
         url: "/k8sDeployments/updateK8sDeployment",
         method: 'put',
         data
     })
 }


// @Tags K8sDeployment
// @Summary 用id查询K8sDeployment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sDeployment true "用id查询K8sDeployment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /k8sDeployments/findK8sDeployment [get]
 export const findK8sDeployment = (params) => {
     return service({
         url: "/k8sDeployments/findK8sDeployment",
         method: 'get',
         params
     })
 }


// @Tags K8sDeployment
// @Summary 分页获取K8sDeployment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取K8sDeployment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sDeployments/getK8sDeploymentList [get]
 export const getK8sDeploymentList = (params) => {
     return service({
         url: "/k8sDeployments/getK8sDeploymentList",
         method: 'get',
         params
     })
 }