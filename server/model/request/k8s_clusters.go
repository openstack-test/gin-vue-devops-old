package request

import "gin-vue-devops/model"

type K8sClusterSearch struct {
	model.K8sCluster
	PageInfo
}
