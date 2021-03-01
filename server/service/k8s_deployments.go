package service

import (
	"context"
	"gin-vue-devops/global"
	"gin-vue-devops/model"
	"gin-vue-devops/model/request"
	"gin-vue-devops/utils"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 定义结构体绑定数据
type DeploymentResult struct {
	ID         int    `json:"id"`
	Namespace  string `json:"namespace"`
	Deployment string `json:"deployment"`
	Replicas   int32  `json:"replicas"`
	CreateTime string `json:"createTime"`
}

//@function: CreateK8sDeployment
//@description: 创建K8sDeployment记录
//@param: k8sDeployments model.K8sDeployment
//@return: err error

func CreateK8sDeployment(k8sDeployments model.K8sDeployment) (err error) {
	err = global.GVA_DB.Create(&k8sDeployments).Error
	return err
}

//@function: DeleteK8sDeployment
//@description: 删除K8sDeployment记录
//@param: k8sDeployments model.K8sDeployment
//@return: err error

func DeleteK8sDeployment(k8sDeployments model.K8sDeployment) (err error) {
	err = global.GVA_DB.Delete(k8sDeployments).Error
	return err
}

//@function: DeleteK8sDeploymentByIds
//@description: 批量删除K8sDeployment记录
//@param: ids request.IdsReq
//@return: err error

func DeleteK8sDeploymentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.K8sDeployment{}, "id in ?", ids.Ids).Error
	return err
}

//@function: UpdateK8sDeployment
//@description: 更新K8sDeployment记录
//@param: k8sDeployments *model.K8sDeployment
//@return: err error

func UpdateK8sDeployment(k8sDeployments model.K8sDeployment) (err error) {
	err = global.GVA_DB.Save(&k8sDeployments).Error
	return err
}

//@function: GetK8sDeployment
//@description: 根据id获取K8sDeployment记录
//@param: id uint
//@return: err error, k8sDeployments model.K8sDeployment

func GetK8sDeployment(id uint) (err error, k8sDeployments model.K8sDeployment) {
	err = global.GVA_DB.Where("id = ?", id).First(&k8sDeployments).Error
	return
}

//@function: GetK8sDeploymentInfoList
//@description: 分页获取K8sDeployment记录
//@param: info request.K8sDeploymentSearch
//@return: err error, list []*DeploymentResult, total int64

func GetK8sDeploymentInfoList(namespace string, info request.K8sDeploymentSearch) (err error, list []*DeploymentResult, total int64) {
	//初始化k8s客户端
	clientset, err := utils.InitClient()
	if err != nil {
		log.Fatalln(err)
	}

	//获取指定namespace的deployments
	deployments, err := clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	for key, deployment := range deployments.Items {
		createTime := deployment.ObjectMeta.CreationTimestamp.Time
		// 将time.Time类型转成指定格式字符串
		formatTime := createTime.Format("2006-01-02 15:04:05")

		res := &DeploymentResult{
			ID:         key,
			Namespace:  namespace,
			Deployment: deployment.ObjectMeta.Name,
			Replicas:   deployment.Status.ReadyReplicas,
			CreateTime: formatTime,
		}
		list = append(list, res)
	}
	total = int64(len(list))
	return err, list, total
}
