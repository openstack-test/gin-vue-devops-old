package service

import (
	"context"
	"fmt"
	"gin-vue-devops/global"
	"gin-vue-devops/model"
	"gin-vue-devops/model/request"
	"gin-vue-devops/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

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
	err = global.GVA_DB.Delete(&[]model.K8sPods{}, "id in ?", ids.Ids).Error
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
//@return: err error, list []*model.K8sPods, total int64

func GetK8sPodsInfoList(k8sConfig string) (err error, list []*model.K8sPods, total int64) {
	// 初始化k8s客户端
	clientSet, _ := utils.GetK8sClient(k8sConfig)

	namespace := "app"
	podList, err := clientSet.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}

	for key, pod := range podList.Items {
		var startTime string
		var restartCount int32
		// 必须判断值是否为nil,否则会导致panic(client-go库bug)
		if pod.Status.StartTime != nil {
			startTime = pod.Status.StartTime.Format("2006-01-02 15:04:05")
		}
		if pod.Status.ContainerStatuses != nil {
			restartCount = pod.Status.ContainerStatuses[0].RestartCount
		}

		res := &model.K8sPods{
			ID:        key,
			PodName:   pod.Name,
			PodIP:     pod.Status.PodIP,
			HostIP:    pod.Status.HostIP,
			Status:    string(pod.Status.Phase),
			StartTime: startTime,
			RestartCount: restartCount,
		}

/*		for k, v := range pod.Status.ContainerStatuses {
			fmt.Printf("index: %#v\n", k)
			fmt.Println("restart count:", v.RestartCount)
		}*/

		// 当Pod没有startTime值,默认为空时(如Pending状态)返回“no time”
		if startTime == "" {
			res.StartTime = "no time"
		}

		list = append(list, res)
	}
	total = int64(len(list))
	return err, list, total
}
