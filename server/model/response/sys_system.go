package response

import "gin-vue-devops/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
