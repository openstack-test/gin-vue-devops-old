package model

import "gin-vue-devops/global"

type K8sCluster struct {
	global.GVA_MODEL
	ID             uint    `json:"id" gorm:"primarykey"`
	ClusterName    string  `json:"clusterName" gorm:"comment:集群名称"`
	KubeConfig     string  `json:"kubeConfig" gorm:"comment:Kubeconfig内容;type:varchar(12800)"`
	ClusterVersion float32 `json:"clusterVersion" gorm:"comment:集群版本"`
}
