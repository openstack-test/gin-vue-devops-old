package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateK8sNamespaces
//@description: 创建K8sNamespaces记录
//@param: k8sNamespaces model.K8sNamespaces
//@return: err error

func CreateK8sNamespaces(k8sNamespaces model.K8sNamespaces) (err error) {
	err = global.GVA_DB.Create(&k8sNamespaces).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteK8sNamespaces
//@description: 删除K8sNamespaces记录
//@param: k8sNamespaces model.K8sNamespaces
//@return: err error

func DeleteK8sNamespaces(k8sNamespaces model.K8sNamespaces) (err error) {
	err = global.GVA_DB.Delete(k8sNamespaces).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteK8sNamespacesByIds
//@description: 批量删除K8sNamespaces记录
//@param: ids request.IdsReq
//@return: err error

func DeleteK8sNamespacesByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.K8sNamespaces{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateK8sNamespaces
//@description: 更新K8sNamespaces记录
//@param: k8sNamespaces *model.K8sNamespaces
//@return: err error

func UpdateK8sNamespaces(k8sNamespaces model.K8sNamespaces) (err error) {
	err = global.GVA_DB.Save(&k8sNamespaces).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetK8sNamespaces
//@description: 根据id获取K8sNamespaces记录
//@param: id uint
//@return: err error, k8sNamespaces model.K8sNamespaces

func GetK8sNamespaces(id uint) (err error, k8sNamespaces model.K8sNamespaces) {
	err = global.GVA_DB.Where("id = ?", id).First(&k8sNamespaces).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetK8sNamespacesInfoList
//@description: 分页获取K8sNamespaces记录
//@param: info request.K8sNamespacesSearch
//@return: err error, list interface{}, total int64

func GetK8sNamespacesInfoList(info request.K8sNamespacesSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.K8sNamespaces{})
    var k8sNamespacess []model.K8sNamespaces
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&k8sNamespacess).Error
	return err, k8sNamespacess, total
}