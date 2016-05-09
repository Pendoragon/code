package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/ghodss/yaml"
	k8s_api "k8s.io/kubernetes/pkg/api"
)

func main() {
	filename, _ := filepath.Abs("./pod.yaml")
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var pod k8s_api.Pod
	j, err := yaml.YAMLToJSON(yamlFile)
	err = json.Unmarshal(j, &pod)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Value: %#v\n", pod)
}
