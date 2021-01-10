package model

import "gin-vue-admin/global"

type K8sNamespaces struct {
	global.GVA_MODEL
	Namespace string `json:"namespace" gorm:"comment:命名空间"`
	Status string `json:"status" gorm:"comment:状态"`
    Age string `json:"age" gorm:"comment:时间"`
}
