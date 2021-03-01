package request

import "gin-vue-devops/model"

type {{.StructName}}Search struct{
    model.{{.StructName}}
    PageInfo
}