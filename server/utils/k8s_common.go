package utils

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func GetK8sClient(k8sConf string) (*kubernetes.Clientset, error) {
	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(k8sConf))
	if err != nil {
		return nil, err
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientSet, nil
}

/*
func IsK8sHealth(k8sConf string) error {
	clientSet, err := GetK8sClient(k8sConf)
	if err != nil {
		return err
	}
	_, err = clientSet.CoreV1().Nodes().List(context.TODO(),v1.ListOptions{})
	if err != nil {
		return err
	}
	return nil
}
 */
