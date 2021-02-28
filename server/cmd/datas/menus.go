package datas

import (
	"gin-vue-admin/global"
	"os"
	"time"

	"github.com/gookit/color"

	"gin-vue-admin/model"

	"gorm.io/gorm"
)

var BaseMenus = []model.SysBaseMenu{
	{GVA_MODEL: global.GVA_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "dashboard", Name: "dashboard", Hidden: false, Component: "view/dashboard/index.vue", Sort: 1, Meta: model.Meta{Title: "首页", Icon: "s-home"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "kubernetes", Name: "kubernetes", Component: "view/kubernetes/index.vue", Sort: 2, Meta: model.Meta{Title: "Kubernetes管理", Icon: "cloudy"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "2", Path: "k8sCluster", Name: "k8sCluster", Component: "view/kubernetes/clusters/index.vue", Sort: 1, Meta: model.Meta{Title: "集群管理", Icon: "menu"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "2", Path: "k8sNamespaces", Name: "k8sNamespaces", Component: "view/kubernetes/namespaces/k8sNamespaces.vue", Sort: 2, Meta: model.Meta{Title: "命名空间管理", Icon: "menu"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "2", Path: "k8sDeployments", Name: "k8sDeployments", Component: "view/kubernetes/deployments/k8sDeployments.vue", Sort: 3, Meta: model.Meta{Title: "应用管理", Icon: "s-grid"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "2", Path: "k8sPods", Name: "k8sPods", Component: "view/kubernetes/pods/index.vue", Sort: 4, Meta: model.Meta{Title: "容器管理", Icon: "s-grid"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "admin", Name: "superAdmin", Component: "view/superAdmin/index.vue", Sort: 3, Meta: model.Meta{Title: "超级管理员", Icon: "user-solid"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "authority", Name: "authority", Component: "view/superAdmin/authority/authority.vue", Sort: 1, Meta: model.Meta{Title: "角色管理", Icon: "s-custom"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "menu", Name: "menu", Component: "view/superAdmin/menu/menu.vue", Sort: 2, Meta: model.Meta{Title: "菜单管理", Icon: "s-order", KeepAlive: true}},
	{GVA_MODEL: global.GVA_MODEL{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "api", Name: "api", Component: "view/superAdmin/api/api.vue", Sort: 3, Meta: model.Meta{Title: "api管理", Icon: "s-platform", KeepAlive: true}},
	{GVA_MODEL: global.GVA_MODEL{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "user", Name: "user", Component: "view/superAdmin/user/user.vue", Sort: 4, Meta: model.Meta{Title: "用户管理", Icon: "coordinate"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "system", Name: "system", Component: "view/superAdmin/system/system.vue", Sort: 5, Meta: model.Meta{Title: "系统配置", Icon: "s-operation"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: true, ParentId: "0", Path: "person", Name: "person", Component: "view/person/person.vue", Sort: 4, Meta: model.Meta{Title: "个人信息", Icon: "message-solid"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "systemTools", Name: "systemTools", Component: "view/systemTools/index.vue", Sort: 5, Meta: model.Meta{Title: "系统工具", Icon: "s-cooperation"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "10", Path: "autoCode", Name: "autoCode", Component: "view/systemTools/autoCode/index.vue", Sort: 1, Meta: model.Meta{Title: "代码生成器", Icon: "cpu", KeepAlive: true}},
	{GVA_MODEL: global.GVA_MODEL{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "10", Path: "formCreate", Name: "formCreate", Component: "view/systemTools/formCreate/index.vue", Sort: 2, Meta: model.Meta{Title: "表单生成器", Icon: "magic-stick", KeepAlive: true}},
	{GVA_MODEL: global.GVA_MODEL{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "10", Path: "upload", Name: "upload", Component: "view/systemTools/upload/upload.vue", Sort: 3, Meta: model.Meta{Title: "上传下载", Icon: "upload"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "state", Name: "state", Hidden: false, Component: "view/system/state.vue", Sort: 6, Meta: model.Meta{Title: "服务器状态", Icon: "s-opportunity"}},
}

func InitSysBaseMenus(db *gorm.DB) {
	if err := db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 23}).Find(&[]model.SysBaseMenu{}).RowsAffected == 2 {
			color.Danger.Println("sys_base_menus表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&BaseMenus).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	}); err != nil {
		color.Warn.Printf("[Mysql]--> sys_base_menus 表的初始数据失败,err: %v\n", err)
		os.Exit(0)
	}
}
