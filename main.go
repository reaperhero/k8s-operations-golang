package main

import (
	"fmt"
	"github.com/reaperhero/k8s-operations-golang/model"
	"github.com/reaperhero/k8s-operations-golang/model/usecase"
)

func main() {
	usecase := usecase.Newclient("/Users/chenqiangjun/.kube/prometheus")
	fmt.Println(usecase.ListDeployment())
	err := usecase.CreateDeployment(model.Deployment{
		Name:      "web",
		Namespace: "default",
		Image:     "nginx:1.12",
		Replicas:  1,
		Port:      nil,
		Env:       nil,
	})
	if err == nil {
		fmt.Println(usecase.ListDeployment())
	}
	err = usecase.UpdateDeployment(model.Deployment{
		Name:      "web",
		Namespace: "default",
		Image:     "nginx:1.13",
		Replicas:  1,
		Port:      nil,
		Env:       nil,
	})
	if err == nil {
		fmt.Println(usecase.ListDeployment())
	}
	usecase.DeleteDeployment("web")
}
