package usecase

import (
	"errors"
	"github.com/reaperhero/k8s-operations-golang/model"
	"github.com/reaperhero/k8s-operations-golang/utils"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *client) ListDeployment() []model.Deployment {
	listdeployment, _ := c.clientSet.AppsV1().Deployments("default").List(v1.ListOptions{})
	if len(listdeployment.Items) == 0 {
		return nil
	}
	deployments := []model.Deployment{}
	for _, item := range listdeployment.Items {
		envmap := make(map[string]string)
		for _, envvalue := range item.Spec.Template.Spec.Containers[0].Env {
			envmap[envvalue.Name] = envvalue.Value
		}
		deployments = append(deployments, model.Deployment{
			Name:      item.ObjectMeta.Name,
			Namespace: item.ObjectMeta.Namespace,
			Image:     item.Spec.Template.Spec.Containers[0].Image,
			Replicas:  int(*item.Spec.Replicas),
			Env:       envmap,
		})
	}
	return deployments
}

func (c *client) CreateDeployment(request model.Deployment) error {

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: request.Name,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: utils.Int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": request.Name,
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": request.Name,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  request.Name,
							Image: request.Image,
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}
	deploymentsClient := c.clientSet.AppsV1().Deployments(request.Namespace)
	_, err := deploymentsClient.Create(deployment)
	return err
}

func (c *client) DeleteDeployment(name string) error {
	deletePolicy := metav1.DeletePropagationForeground
	deploymentsClient := c.clientSet.AppsV1().Deployments("default")
	err := deploymentsClient.Delete("web", &metav1.DeleteOptions{
		TypeMeta:          v1.TypeMeta{},
		PropagationPolicy: &deletePolicy,
	})
	return err
}

func (c *client) UpdateDeployment(request model.Deployment) error {
	if !c.checkExistdeployment(request.Name) {
		return errors.New("deployment is not exist")
	}

	deploymentsClient := c.clientSet.AppsV1().Deployments("default")
	deploymentinfo, err := c.getDeploymentInfo(request.Name)
	if err != nil {
		return err
	}
	deploymentinfo.Spec.Template.Spec.Containers[0].Image = request.Image
	_, err = deploymentsClient.Update(deploymentinfo)
	return err
}

func (c *client) checkExistdeployment(name string) bool {
	list := c.ListDeployment()
	for _, deployment := range list {
		if deployment.Name == name {
			return true
		}
	}
	return false
}

func (c *client) getDeploymentInfo(name string) (*appsv1.Deployment, error) {
	deploymentsClient := c.clientSet.AppsV1().Deployments("default")
	result, err := deploymentsClient.Get(name, metav1.GetOptions{})
	return result, err
}
