package k8s

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"kloud/pkg/conf"
)

var client *kubernetes.Clientset
var config *rest.Config
var portmap map[int]chan struct{}

func init() {
	var err error
	config, err = clientcmd.BuildConfigFromFlags("", conf.GetConf().K8s.ConfigPath)
	if err != nil {
		panic(err.Error())
	}

	// create the client
	client, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	portmap = make(map[int]chan struct{})
}
