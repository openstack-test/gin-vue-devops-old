package model

import v1 "k8s.io/api/core/v1"

type K8sPods struct {
	ID           int         `json:"id"`
	PodName      string      `json:"podName"`
	PodIP        string      `json:"podIP"`
	HostIP       string      `json:"hostIP"`
	Status       v1.PodPhase `json:"status"`
	StartTime    string      `json:"startTime"`
	RestartCount int32       `json:"restartCount"`
}
