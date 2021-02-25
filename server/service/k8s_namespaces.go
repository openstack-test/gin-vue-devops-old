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
)

// 定义结构体绑定数据
type Result struct {
	ID        int               `json:"id"`
	Namespace string            `json:"namespace"`
	Status    v1.NamespacePhase `json:"status"`
	CreateTime    string       `json:"createTime"`
}

//@function: GetK8sNamespaces
//@description: 根据id获取K8sNamespaces记录
//@param: id uint
//@return: err error, k8sNamespaces model.K8sNamespaces
func FindK8sNamespaces(id uint) (err error, k8sNamespaces model.K8sNamespaces) {
	err = global.GVA_DB.Where("id = ?", id).First(&k8sNamespaces).Error
	return
}

//@function: GetK8sNamespacesList
//@description: 分页获取K8sNamespaces记录
//@param: info request.K8sNamespacesSearch
//@return: err error, list []*Result, total int64
func GetK8sNamespacesList(info request.K8sNamespacesSearch) (err error, list []*Result, total int64) {
	// 初始化k8s客户端
	clientset, err := utils.InitClient()
	if err != nil {
		log.Fatalln(err)
	}
	// 获取所有Namespaces
	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	for key, ns := range namespaces.Items {
		creatTime := ns.CreationTimestamp.Time
		// 将time.Time类型转成指定格式字符串
		formatTime := creatTime.Format("2006-01-02 15:04:05")

		res := &Result{
			ID:        key,
			Namespace: ns.Name,
			Status:    ns.Status.Phase,
			CreateTime:      formatTime,
		}
		list = append(list, res)
	}
	total = int64(len(list))
	return err, list, total
}
