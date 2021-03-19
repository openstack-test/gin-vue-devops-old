package model

type K8sPods struct {
	ID           int         `json:"id"`
	PodName      string      `json:"podName"`
	PodIP        string      `json:"podIP"`
	HostIP       string      `json:"hostIP"`
	Status       string      `json:"status"`
	StartTime    string      `json:"startTime"`
	RestartCount int32       `json:"restartCount"`
}
