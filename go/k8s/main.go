package main

import (
	"flag"

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

	var test_service *k8s_api.Service
	// var test_pod *k8s_api.Pod

	test_service.Labels["name"] = "test"
	test_service.Name = "test"
	test_service.Spec.Selector = make(map[string]string)

	if _, err := c.Services("default").Create(test_service); err != nil {
		glog.Errorf("Failed to create service due to: %v", err)
	}

	for {}
}
