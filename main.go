package main

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/chenqiangjun/.kube/prometheus")
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	list,err := clientset.AppsV1().Deployments("default").List(v1.ListOptions{
	})
	if err !=nil{
		log.Println(err)
	}
	log.Println(list.Items[0])
}
