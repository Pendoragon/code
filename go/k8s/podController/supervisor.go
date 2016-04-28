package main

import (
	"fmt"
	"time"
	"reflect"

	"github.com/golang/glog"

	"k8s.io/kubernetes/pkg/client/cache"
	"k8s.io/kubernetes/pkg/controller/framework"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/watch"
	k8s_api "k8s.io/kubernetes/pkg/api"
	k8s_client "k8s.io/kubernetes/pkg/client/unversioned"
)

type supervisor struct {
	client         *k8s_client.Client
	podController  *framework.Controller
	podLister      cache.StoreToPodLister
	podManager     *PodManager

	stopCh   chan struct{}

}

func NewSupervisor(kubeClient *k8s_client.Client, resyncPeriod time.Duration, namespace string) (*supervisor, error){
	c := supervisor{
		client: kubeClient,
		stopCh: make(chan struct{}),
		podManager: NewPodManager(kubeClient, 1),
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

	c.podLister.Store, c.podController = framework.NewInformer(
		&cache.ListWatch{
			ListFunc:  podListFunc(c.client, namespace),
			WatchFunc: podWatchFunc(c.client, namespace),
		},
		&k8s_api.Pod{}, resyncPeriod, podEventHandler)

	return &c, nil
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

func (s *supervisor)StartPodManager() {
	go s.podManager.SpawnPod("default")
}

// func serviceListFunc(c *k8s_client.Client, ns string) func(k8s_api.ListOptions) (runtime.Object, error) {
// 	return func(opts k8s_api.ListOptions) (runtime.Object, error) {
// 		return c.Services(ns).List(opts)
// 	}
// }

// func serviceWatchFunc(c *k8s_client.Client, ns string) func(options k8s_api.ListOptions) (watch.Interface, error) {
// 	return func(options k8s_api.ListOptions) (watch.Interface, error) {
// 		return c.Services(ns).Watch(options)
// 	}
// }

// func (c *controller) sync(key string) {
// 	if !lbc.controllersInSync() {
// 		lbc.syncQueue.requeue(key, fmt.Errorf("deferring sync till endpoints controller has synced"))
// 		return
// 	}

// 	ings := lbc.ingLister.Store.List()
// 	upstreams, servers := lbc.getUpstreamServers(ings)

// 	var cfg *k8s_api.ConfigMap

// 	ns, name, _ := parseNsName(lbc.nxgConfigMap)
// 	cfg, err := lbc.getConfigMap(ns, name)
// 	if err != nil {
// 		cfg = &k8s_api.ConfigMap{}
// 	}

// 	ngxConfig := lbc.nginx.ReadConfig(cfg)
// 	lbc.nginx.CheckAndReload(ngxConfig, nginx.IngressConfig{
// 		Upstreams:    upstreams,
// 		Servers:      servers,
// 		TCPUpstreams: lbc.getTCPServices(),
// 		UDPUpstreams: lbc.getUDPServices(),
// 	})
// }
