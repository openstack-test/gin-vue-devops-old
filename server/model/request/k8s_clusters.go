package request

import "gin-vue-admin/model"

type K8sClusterSearch struct{
    model.K8sCluster
    PageInfo
}