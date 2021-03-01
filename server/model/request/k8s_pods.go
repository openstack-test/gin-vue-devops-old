package request

import "gin-vue-devops/model"

type K8sPodsSearch struct {
	model.K8sPods
	PageInfo
}
