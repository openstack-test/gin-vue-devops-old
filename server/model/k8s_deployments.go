package model

import "gin-vue-admin/global"

type K8sDeployment struct {
	global.GVA_MODEL
	Namespace string `json:"namespace" gorm:"comment:命名空间"`
	Deployment string `json:"deployment" gorm:"comment:kind类型"`
	Status string `json:"status" gorm:"comment:状态"`
    Time string `json:"time" gorm:"comment:时间"`
}
