package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateK8sDeployment
//@description: 创建K8sDeployment记录
//@param: k8sDeployments model.K8sDeployment
//@return: err error

func CreateK8sDeployment(k8sDeployments model.K8sDeployment) (err error) {
	err = global.GVA_DB.Create(&k8sDeployments).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteK8sDeployment
//@description: 删除K8sDeployment记录
//@param: k8sDeployments model.K8sDeployment
//@return: err error

func DeleteK8sDeployment(k8sDeployments model.K8sDeployment) (err error) {
	err = global.GVA_DB.Delete(k8sDeployments).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteK8sDeploymentByIds
//@description: 批量删除K8sDeployment记录
//@param: ids request.IdsReq
//@return: err error

func DeleteK8sDeploymentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.K8sDeployment{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateK8sDeployment
//@description: 更新K8sDeployment记录
//@param: k8sDeployments *model.K8sDeployment
//@return: err error

func UpdateK8sDeployment(k8sDeployments model.K8sDeployment) (err error) {
	err = global.GVA_DB.Save(&k8sDeployments).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetK8sDeployment
//@description: 根据id获取K8sDeployment记录
//@param: id uint
//@return: err error, k8sDeployments model.K8sDeployment

func GetK8sDeployment(id uint) (err error, k8sDeployments model.K8sDeployment) {
	err = global.GVA_DB.Where("id = ?", id).First(&k8sDeployments).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetK8sDeploymentInfoList
//@description: 分页获取K8sDeployment记录
//@param: info request.K8sDeploymentSearch
//@return: err error, list interface{}, total int64

func GetK8sDeploymentInfoList(info request.K8sDeploymentSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.K8sDeployment{})
    var k8sDeploymentss []model.K8sDeployment
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&k8sDeploymentss).Error
	return err, k8sDeploymentss, total
}