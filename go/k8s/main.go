package main

import (
	"flag"

	"github.com/golang/glog"
	k8s_core_api "k8s.io/kubernetes/pkg/api/v1"
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

	var test_service *k8s_core_api.Service
	var test_pod *k8s_core_api.Pod

	test_service.Labels["name"] = "test"
	test_service.Name = "test"
	test_service.Spec.Selector = make(map[string]string)

	if _, err := k8s_client.Services("default").Create(test_service); err != nil {
		glog.Errorf("Failed to create service due to: %v", err)
	}

	for {}
}
