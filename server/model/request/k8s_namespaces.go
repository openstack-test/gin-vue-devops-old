package request

import "gin-vue-admin/model"

type K8sNamespacesSearch struct{
    model.K8sNamespaces
    PageInfo
}