package request

import "gin-vue-admin/model"

type K8sPodsSearch struct{
    model.K8sPods
    PageInfo
}