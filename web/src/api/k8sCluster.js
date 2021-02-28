import service from '@/utils/request'

// @Tags K8sCluster
// @Summary 创建K8sCluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sCluster true "创建K8sCluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sCluster/createK8sCluster [post]
export const createK8sCluster = (data) => {
     return service({
         url: "/k8sCluster/createK8sCluster",
         method: 'post',
         data
     })
 }


// @Tags K8sCluster
// @Summary 删除K8sCluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sCluster true "删除K8sCluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sCluster/deleteK8sCluster [delete]
 export const deleteK8sCluster = (data) => {
     return service({
         url: "/k8sCluster/deleteK8sCluster",
         method: 'delete',
         data
     })
 }

// @Tags K8sCluster
// @Summary 删除K8sCluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除K8sCluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sCluster/deleteK8sCluster [delete]
 export const deleteK8sClusterByIds = (data) => {
     return service({
         url: "/k8sCluster/deleteK8sClusterByIds",
         method: 'delete',
         data
     })
 }

// @Tags K8sCluster
// @Summary 更新K8sCluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sCluster true "更新K8sCluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /k8sCluster/updateK8sCluster [put]
 export const updateK8sCluster = (data) => {
     return service({
         url: "/k8sCluster/updateK8sCluster",
         method: 'put',
         data
     })
 }


// @Tags K8sCluster
// @Summary 用id查询K8sCluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sCluster true "用id查询K8sCluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /k8sCluster/findK8sCluster [get]
 export const findK8sCluster = (params) => {
     return service({
         url: "/k8sCluster/findK8sCluster",
         method: 'get',
         params
     })
 }


// @Tags K8sCluster
// @Summary 分页获取K8sCluster列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取K8sCluster列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sCluster/getK8sClusterList [get]
 export const getK8sClusterList = (params) => {
     return service({
         url: "/k8sCluster/getK8sClusterList",
         method: 'get',
         params
     })
 }