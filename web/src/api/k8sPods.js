import service from '@/utils/request'

// @Tags K8sPods
// @Summary 创建K8sPods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sPods true "创建K8sPods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sPods/createK8sPods [post]
export const createK8sPods = (data) => {
     return service({
         url: "/k8sPods/createK8sPods",
         method: 'post',
         data
     })
 }


// @Tags K8sPods
// @Summary 删除K8sPods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sPods true "删除K8sPods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sPods/deleteK8sPods [delete]
 export const deleteK8sPods = (data) => {
     return service({
         url: "/k8sPods/deleteK8sPods",
         method: 'delete',
         data
     })
 }

// @Tags K8sPods
// @Summary 删除K8sPods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除K8sPods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sPods/deleteK8sPods [delete]
 export const deleteK8sPodsByIds = (data) => {
     return service({
         url: "/k8sPods/deleteK8sPodsByIds",
         method: 'delete',
         data
     })
 }

// @Tags K8sPods
// @Summary 更新K8sPods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sPods true "更新K8sPods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /k8sPods/updateK8sPods [put]
 export const updateK8sPods = (data) => {
     return service({
         url: "/k8sPods/updateK8sPods",
         method: 'put',
         data
     })
 }


// @Tags K8sPods
// @Summary 用id查询K8sPods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sPods true "用id查询K8sPods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /k8sPods/findK8sPods [get]
 export const findK8sPods = (params) => {
     return service({
         url: "/k8sPods/findK8sPods",
         method: 'get',
         params
     })
 }


// @Tags K8sPods
// @Summary 分页获取K8sPods列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取K8sPods列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sPods/getK8sPodsList [get]
 export const getK8sPodsList = (params) => {
     return service({
         url: "/k8sPods/getK8sPodsList",
         method: 'get',
         params
     })
 }