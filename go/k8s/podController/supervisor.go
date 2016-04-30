package main

import (
	"fmt"
	"time"
	"reflect"

	"github.com/golang/glog"

	"k8s.io/kubernetes/pkg/runtime"
	k8s_api "k8s.io/kubernetes/pkg/api"
	k8s_client "k8s.io/kubernetes/pkg/client/unversioned"
)

type Manager struct {
	PodController  PodController

	stopCh   chan struct{}

}

func NewManager(kubeClient *k8s_client.Client, resyncPeriod time.Duration, namespace string) (*supervisor, error){
	c := Manager{
		stopCh: make(chan struct{}),
		PodController: NewPodController(kubeClient, 1),
	}


	return &c, nil
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
