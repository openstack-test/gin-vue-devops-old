package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@function: CreateK8sCluster
//@description: 创建K8sCluster记录
//@param: K8sCluster model.K8sCluster
//@return: err error

func CreateK8sCluster(K8sCluster model.K8sCluster) (err error) {
	err = global.GVA_DB.Create(&K8sCluster).Error
	return err
}

//@function: DeleteK8sCluster
//@description: 删除K8sCluster记录
//@param: K8sCluster model.K8sCluster
//@return: err error

func DeleteK8sCluster(K8sCluster model.K8sCluster) (err error) {
	err = global.GVA_DB.Delete(K8sCluster).Error
	return err
}

//@function: DeleteK8sClusterByIds
//@description: 批量删除K8sCluster记录
//@param: ids request.IdsReq
//@return: err error

func DeleteK8sClusterByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.K8sCluster{}, "id in ?", ids.Ids).Error
	return err
}

//@function: UpdateK8sCluster
//@description: 更新K8sCluster记录
//@param: K8sCluster *model.K8sCluster
//@return: err error

func UpdateK8sCluster(K8sCluster model.K8sCluster) (err error) {
	err = global.GVA_DB.Save(&K8sCluster).Error
	return err
}

//@function: GetK8sCluster
//@description: 根据id获取K8sCluster记录
//@param: id uint
//@return: err error, K8sCluster model.K8sCluster

func GetK8sCluster(id uint) (err error, K8sCluster model.K8sCluster) {
	err = global.GVA_DB.Where("id = ?", id).First(&K8sCluster).Error
	return
}

//@function: GetK8sClusterInfoList
//@description: 分页获取K8sCluster记录
//@param: info request.K8sClusterSearch
//@return: err error, list interface{}, total int64

func GetK8sClusterInfoList(info request.K8sClusterSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.K8sCluster{})
	var K8sClusters []model.K8sCluster
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&K8sClusters).Error
	return err, K8sClusters, total
}
