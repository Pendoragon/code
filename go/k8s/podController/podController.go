package main

import (
	"time"

	"github.com/golang/glog"

	"k8s.io/kubernetes/pkg/client/cache"
	"k8s.io/kubernetes/pkg/controller/framework"
	"k8s.io/kubernetes/pkg/watch"
	k8s_api "k8s.io/kubernetes/pkg/api"
	k8s_client "k8s.io/kubernetes/pkg/client/unversioned"
)

type PodController struct {
	client     *k8s_client.Client
	podController  *framework.Controller
	podLister      cache.StoreToPodLister
	podLimit   int
}

func NewPodController(client *k8s_client.Client, limit int) *PodController {
	pc := &PodController{
		client: client,
		podLimit: limit,
	}

	// event handler to just print out all the new events.
	podEventHandler := framework.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			addPod := obj.(*k8s_api.Pod)
			if addPod.Labels["name"] == "test-pod" {
				glog.Info(fmt.Sprintf("ADD %s/%s", addPod.Namespace, addPod.Name))
				glog.Info(fmt.Sprintf("- Pod in status %s", addPod.Status))
			}
		},
		DeleteFunc: func(obj interface{}) {
			delPod := obj.(*k8s_api.Pod)
			if delPod.Labels["name"] == "test-pod" {
				glog.Info(fmt.Sprintf("DELETE %s/%s", delPod.Namespace, delPod.Name))
				glog.Info(fmt.Sprintf("- Pod in status %s", delPod.Status))
			}
		},
		UpdateFunc: func(old, cur interface{}) {
			if !reflect.DeepEqual(old, cur) {
				upPod := cur.(*k8s_api.Pod)
				if upPod.Labels["name"] == "test-pod" {
					glog.Info(fmt.Sprintf("UPDATE %s/%s", upPod.Namespace, upPod.Name))
					glog.Info(fmt.Sprintf("- Pod in status %s", upPod.Status))
				}
			}
		},
	}

	pc.podLister.Store, pc.podController = framework.NewInformer(
		&cache.ListWatch{
			ListFunc:  podListFunc(c.client, namespace),
			WatchFunc: podWatchFunc(c.client, namespace),
		},
		&k8s_api.Pod{}, resyncPeriod, podEventHandler)

}

func podWatchFunc(c *k8s_client.Client, ns string) func(options k8s_api.ListOptions) (watch.Interface, error) {
	return func(options k8s_api.ListOptions) (watch.Interface, error) {
		return c.Pods(ns).Watch(options)
	}
}

func podListFunc(c *k8s_client.Client, ns string) func(k8s_api.ListOptions) (runtime.Object, error) {
	return func(opts k8s_api.ListOptions) (runtime.Object, error) {
		return c.Pods(ns).List(opts)
	}
}


func (m *PodController) SpawnPod(ns string) {
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
