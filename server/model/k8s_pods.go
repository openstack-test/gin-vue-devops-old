package model

import "gin-vue-admin/global"

type K8sPods struct {
	global.GVA_MODEL
	pod string `json:"pod" gorm:"comment:pod名"`
	Status string `json:"status" gorm:"comment:状态"`
    Time string `json:"time" gorm:"comment:时间"`
}
