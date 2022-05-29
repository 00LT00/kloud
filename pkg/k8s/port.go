package k8s

import (
	"errors"
	"fmt"
	"k8s.io/client-go/tools/portforward"
	"k8s.io/client-go/transport/spdy"
	"log"
	"net/http"
)

// 创建端口映射 resource指的是kubernetes对象类型，name是指要操作的对象
func ForwardPort(resource, name string, port, target int) error {
	//判断端口范围是否合法
	if port < 1 || port > 65535 {
		return errors.New("port must be in range [1, 65535]")
	}
	if _, ok := portmap[port]; ok {
		return errors.New("port has been used")
	}
	//判断目标端口范围是否合法
	if target < 1 || target > 65535 {
		return errors.New("target port must be in range [1, 65535]")
	}

	req := client.CoreV1().RESTClient().Post().Namespace("default").Resource(resource).Name(name).SubResource("portforward")

	StopChannel := make(chan struct{}, 1)
	ReadyChannel := make(chan struct{})
	transport, upgrader, err := spdy.RoundTripperFor(config)
	if err != nil {
		return err
	}
	// 不限制地址
	address := []string{"0.0.0.0"}
	// 目标端口
	ports := []string{fmt.Sprintf("%d:%d", port, target)}
	// 创建转发对象
	dialer := spdy.NewDialer(upgrader, &http.Client{Transport: transport}, "POST", req.URL())
	fw, err := portforward.NewOnAddresses(dialer, address, ports, StopChannel, ReadyChannel, nil, nil)
	if err != nil {
		return err
	}
	portmap[port] = StopChannel
	// 此时认为已经创建成功了
	go func() {
		//无论是否出现错误，都不应该对portmap进行修改
		if err := fw.ForwardPorts(); err != nil {
			log.Println(err)
		}
	}()
	return nil
}

func StopPort(port int) error {
	if _, ok := portmap[port]; !ok {
		return errors.New("port not found")
	}
	portmap[port] <- struct{}{}
	close(portmap[port])
	delete(portmap, port)
	return nil
}
