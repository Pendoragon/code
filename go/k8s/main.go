package main

import (
	"flag"
	"time"

	"github.com/golang/glog"
	k8s_api "k8s.io/kubernetes/pkg/api"
	k8s_client "k8s.io/kubernetes/pkg/client/unversioned"
)

func main() {
	flag.Parse()
	glog.Info("k8s programmatically create pods test.")

	c, err := k8s_client.NewInCluster()
	if err != nil {
		glog.Fatalf("Failed to make client: %v", err)
	}

	if err != nil {
		glog.Fatalf("Failed to make client: %v", err)
	}

	test_service := &k8s_api.Service{}
	// var test_pod *k8s_api.Pod
	test_service.Labels = make(map[string]string)
	test_service.Labels["name"] = "test"
	test_service.Name = "test"
	test_service.Spec.Selector = make(map[string]string)
	service_port := []k8s_api.ServicePort{
		k8s_api.ServicePort{
			Protocol: k8s_api.ProtocolTCP,
			Port: 8080,
		},
	}

	test_service.Spec.Ports = service_port
	if _, err := c.Services("default").Create(test_service); err != nil {
		glog.Errorf("Failed to create service due to: %v", err)
	}

	s, err := newSupervisor(c, 5 * time.Second, k8s_api.NamespaceAll)
	for {
		s.podLister.Store.List()
		time.Sleep(3 * time.Second)
	}
}
