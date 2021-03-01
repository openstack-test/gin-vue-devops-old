package request

import "gin-vue-devops/model"

type K8sDeploymentSearch struct {
	model.K8sDeployment
	PageInfo
}
