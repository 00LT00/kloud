package k8s

import (
	"bytes"
	"context"
	"github.com/gofrs/uuid"
	"io/ioutil"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"kloud/model"
	"kloud/pkg/DB"
	"kloud/pkg/conf"
	"log"
	"path"
	"text/template"
)

type Creator struct {
	Resource *model.Resource
}

func NewCreator(resource *model.Resource) Creator {
	return Creator{
		Resource: resource,
	}
}

func (c Creator) Create(a *model.App) (err error) {
	log.Println("create k8s app")
	v4, _ := uuid.NewV4()
	id := v4.String()
	var data []byte
	//获取模板文件
	filename := path.Join(c.Resource.Folder, "template.yaml")
	if data, err = ioutil.ReadFile(filename); err != nil {
		return
	}
	//生成配置文件的[]byte
	config := make(map[string]any)
	config["id"] = id
	for k, v := range a.Config {
		config[k] = v
	}
	if data, err = c.parse(config, string(data)); err != nil {
		return
	}
	//生成pod
	if err = c.createPod(data); err != nil {
		return
	}
	a.AppID = id
	err = DB.GetDB().Create(a).Error
	return
}

func (c Creator) parse(config map[string]any, tpl string) (result []byte, err error) {
	t := template.New("pod")
	t, err = t.Parse(tpl)
	if err != nil {
		return
	}
	var buf bytes.Buffer
	err = t.Execute(&buf, config)
	if err != nil {
		return
	}
	result = buf.Bytes()
	return
}

func (c Creator) createPod(data []byte) (err error) {
	pod := &v1.Pod{}
	err = yaml.Unmarshal(data, pod)
	if err != nil {
		panic(err)
	}
	namespace := conf.GetConf().K8s.Namespace
	pod, err = client.CoreV1().Pods(namespace).Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		return
	}
	return
}
