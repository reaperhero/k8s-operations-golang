package usecase

import (
	"github.com/reaperhero/k8s-operations-golang/model"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type ClientUsecase interface {
	ListDeployment() []model.Deployment
	CreateDeployment(request model.Deployment) error
	DeleteDeployment(name string) error
	UpdateDeployment(request model.Deployment) error
}
type client struct {
	clientSet *kubernetes.Clientset
}

func Newclient(kubepath string) ClientUsecase {
	config, err := clientcmd.BuildConfigFromFlags("", kubepath)
	if err != nil {
		return nil
	}
	clientset, err := kubernetes.NewForConfig(config)
	return &client{clientSet: clientset}
}
