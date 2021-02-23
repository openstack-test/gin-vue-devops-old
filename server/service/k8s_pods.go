package service

import (
	"context"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/utils"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"time"
)

// 定义结构体绑定数据
type PodResult struct {
	ID           int               `json:"id"`
	PodName      string            `json:"podName"`
	PodIP        string            `json:"podIP"`
	Status       v1.PodPhase      `json:"status"`
	StartTime    time.Time         `json:"startTime"`
	RestartCount int32             `json:"restartCount"`
}

//@function: CreateK8sPods
//@description: 创建K8sPods记录
//@param: k8sPods model.K8sPods
//@return: err error

func CreateK8sPods(k8sPods model.K8sPods) (err error) {
	err = global.GVA_DB.Create(&k8sPods).Error
	return err
}

//@function: DeleteK8sPods
//@description: 删除K8sPods记录
//@param: k8sPods model.K8sPods
//@return: err error

func DeleteK8sPods(k8sPods model.K8sPods) (err error) {
	err = global.GVA_DB.Delete(k8sPods).Error
	return err
}

//@function: DeleteK8sPodsByIds
//@description: 批量删除K8sPods记录
//@param: ids request.IdsReq
//@return: err error

func DeleteK8sPodsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.K8sPods{},"id in ?",ids.Ids).Error
	return err
}

//@function: UpdateK8sPods
//@description: 更新K8sPods记录
//@param: k8sPods *model.K8sPods
//@return: err error

func UpdateK8sPods(k8sPods model.K8sPods) (err error) {
	err = global.GVA_DB.Save(&k8sPods).Error
	return err
}

//@function: GetK8sPods
//@description: 根据id获取K8sPods记录
//@param: id uint
//@return: err error, k8sPods model.K8sPods

func GetK8sPods(id uint) (err error, k8sPods model.K8sPods) {
	err = global.GVA_DB.Where("id = ?", id).First(&k8sPods).Error
	return
}

//@function: GetK8sPodsInfoList
//@description: 分页获取K8sPods记录
//@param: info request.K8sPodsSearch
//@return: err error, list []*PodResult, total int64
func GetK8sPodsInfoList(namespace string, info request.K8sPodsSearch) (err error, list []*PodResult, total int64) {
	// 初始化k8s客户端
	clientset, err := utils.InitClient()
	if err != nil {
		log.Fatalln(err)
	}
    namespace = "app"
	podList, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for key, pods := range podList.Items {
		res := &PodResult{
			ID: key,
			PodName: pods.ObjectMeta.Name,
			PodIP: pods.Status.PodIP,
			Status: pods.Status.Phase,
			StartTime: pods.Status.StartTime.Time,
			RestartCount: pods.Status.ContainerStatuses[0].RestartCount,
		}
		list = append(list, res)
	}
	total = int64(len(list))
	return err, list, total
}