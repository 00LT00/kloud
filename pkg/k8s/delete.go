package k8s

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kloud/pkg/conf"
	"log"
)

func DeleteApp(AppID string) (err error) {
	namespace := conf.GetConf().K8s.Namespace
	err = client.CoreV1().Pods(namespace).Delete(context.TODO(), AppID, metav1.DeleteOptions{})
	if err != nil {
		log.Fatal(err)
	}
	return
}
