package model

import "gin-vue-admin/global"

type K8sCluster struct {
	global.GVA_MODEL
	ClusterName string `json:"clusterName" gorm:"comment:集群名称"`
	KubeConfig string `json:"kubeConfig" gorm:"comment:Kube配置"`
	ClusterVersion float32 `json:"clusterVersion" gorm:"comment:集群版本"`
}
