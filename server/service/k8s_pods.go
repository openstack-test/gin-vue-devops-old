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
func GetK8sPodsInfoList(namespace string, info request.K8sPodsSearch) (err error, list []*model.K8sPods, total int64) {
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

	for key, pod := range podList.Items {
		startTime := pod.Status.StartTime.Time
		// 将time.Time类型转成指定格式字符串
		formatTime := startTime.Format("2006-01-02 15:04:05")

		res := &model.K8sPods{
			ID:           key,
			PodName:      pod.ObjectMeta.Name,
			PodIP:        pod.Status.PodIP,
			HostIP:       pod.Status.HostIP,
			Status:       string(pod.Status.Phase),
			StartTime:    formatTime,
			RestartCount: pod.Status.ContainerStatuses[0].RestartCount,
		}
		list = append(list, res)
	}
	total = int64(len(list))
	return err, list, total
}
