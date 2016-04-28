package main

import (
	"time"

	"github.com/golang/glog"

	k8s_api "k8s.io/kubernetes/pkg/api"
	k8s_client "k8s.io/kubernetes/pkg/client/unversioned"
)

type PodManager struct {
	client     *k8s_client.Client
	podLimit   int
}

func NewPodManager(client *k8s_client.Client, limit int) *PodManager {
	return &PodManager{
		client: client,
		podLimit: limit,
	}
}

func (m *PodManager) SpawnPod(ns string) {
	// maybe use something from kubectl to convert yaml to api objects
	// test_pod := m.ConvertYamlToPod()
	test_pod := &k8s_api.Pod{}
	test_pod.Name = "test-pod"
	test_pod.Labels = make(map[string]string)
	test_pod.Labels["name"] = "test-pod"
	test_pod.Spec.Containers = []k8s_api.Container{
		k8s_api.Container{
			Name: "test-pod",
			Image: "index.caicloud.io/caicloud/busybox-curl:1.0",
			Command: []string{"echo", "done!"},
		},
	}

	test_pod.Spec.RestartPolicy = k8s_api.RestartPolicyNever
	glog.Infof("Pod template: %+v", *test_pod)
	for i := 0; i < m.podLimit; i++ {
		glog.Infof("Creating pod %v", i)
		_, err := m.client.Pods(ns).Create(test_pod)
		if err != nil {
			glog.Errorf("Error creating pod: %v", err)
		}
		time.Sleep(5 * time.Second)
	}
}
