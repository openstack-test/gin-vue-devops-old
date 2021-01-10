import service from '@/utils/request'

// @Tags K8sNamespaces
// @Summary 创建K8sNamespaces
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sNamespaces true "创建K8sNamespaces"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sNamespaces/createK8sNamespaces [post]
export const createK8sNamespaces = (data) => {
     return service({
         url: "/k8sNamespaces/createK8sNamespaces",
         method: 'post',
         data
     })
 }


// @Tags K8sNamespaces
// @Summary 删除K8sNamespaces
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sNamespaces true "删除K8sNamespaces"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sNamespaces/deleteK8sNamespaces [delete]
 export const deleteK8sNamespaces = (data) => {
     return service({
         url: "/k8sNamespaces/deleteK8sNamespaces",
         method: 'delete',
         data
     })
 }

// @Tags K8sNamespaces
// @Summary 删除K8sNamespaces
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除K8sNamespaces"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sNamespaces/deleteK8sNamespaces [delete]
 export const deleteK8sNamespacesByIds = (data) => {
     return service({
         url: "/k8sNamespaces/deleteK8sNamespacesByIds",
         method: 'delete',
         data
     })
 }

// @Tags K8sNamespaces
// @Summary 更新K8sNamespaces
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sNamespaces true "更新K8sNamespaces"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /k8sNamespaces/updateK8sNamespaces [put]
 export const updateK8sNamespaces = (data) => {
     return service({
         url: "/k8sNamespaces/updateK8sNamespaces",
         method: 'put',
         data
     })
 }


// @Tags K8sNamespaces
// @Summary 用id查询K8sNamespaces
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sNamespaces true "用id查询K8sNamespaces"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /k8sNamespaces/findK8sNamespaces [get]
 export const findK8sNamespaces = (params) => {
     return service({
         url: "/k8sNamespaces/findK8sNamespaces",
         method: 'get',
         params
     })
 }


// @Tags K8sNamespaces
// @Summary 分页获取K8sNamespaces列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取K8sNamespaces列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sNamespaces/getK8sNamespacesList [get]
 export const getK8sNamespacesList = (params) => {
     return service({
         url: "/k8sNamespaces/getK8sNamespacesList",
         method: 'get',
         params
     })
 }