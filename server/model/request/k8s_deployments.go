package request

import "gin-vue-admin/model"

type K8sDeploymentSearch struct{
    model.K8sDeployment
    PageInfo
}