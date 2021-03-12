package model

import (
	v1 "k8s.io/api/core/v1"
)

type K8sNamespaces struct {
	ID        int              `json:"id" gorm:"primarykey"`
	Namespace string            `json:"namespace" gorm:"comment:命名空间"`
	Status    v1.NamespacePhase `json:"status" gorm:"comment:状态"`
	CreateTime      string            `json:"createTime" gorm:"comment:创建时间"`
}
