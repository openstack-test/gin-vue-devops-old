package request

import "gin-vue-devops/model"

type K8sNamespacesSearch struct {
	model.K8sNamespaces
	PageInfo
}
