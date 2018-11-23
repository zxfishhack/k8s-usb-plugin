// Copyright (c) 2017, NVIDIA CORPORATION. All rights reserved.

package main

import (
	"log"
	"fmt"

	pluginapi "k8s.io/kubernetes/pkg/kubelet/apis/deviceplugin/v1beta1"
)

func check(err error) {
	if err != nil {
		log.Panicln("Fatal:", err)
	}
}

func getDevices() []*pluginapi.Device {
	n := 1024; //TODO:1024个设备应该够用了

	var devs []*pluginapi.Device
	for i := 1; i <= n; i++ {
		devs = append(devs, &pluginapi.Device{
			ID:     fmt.Sprintf("%v", i),
			Health: pluginapi.Healthy,
		})
	}

	return devs
}
