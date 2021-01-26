import service from '@/utils/request'

// @Tags K8sNamespaces
// @Summary 用id查询K8sNamespaces
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sNamespaces true "用id查询K8sNamespaces"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /k8sNamespaces/findK8sNamespaces [get]
 export const findK8sNamespaces = (data) => {
     return service({
         url: "/k8sNamespaces/findK8sNamespaces",
         method: 'get',
         data: data
     })
 }


// @Tags K8sNamespaces
// @Summary 分页获取K8sNamespaces列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取K8sNamespaces列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sNamespaces/getK8sNamespacesList [post]
 export const getK8sNamespacesList = (data) => {
     return service({
         url: "/k8sNamespaces/getK8sNamespacesList",
         method: 'post',
         data: data
     })
 }