package deployment

import (
	"context"
	"time"

	models "count/models"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	// "fmt"
	// "reflect"
)

func Deployments() []models.DeploymentStatus {
	config, err := clientcmd.BuildConfigFromFlags("", "/home/apton/.kube/config")
	if err != nil {
		panic(err)
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	// create a context
	ctx := context.TODO()

	// List all namespaces
	namespaces, err := client.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	var deploymentStatuses []models.DeploymentStatus
	for _, ns := range namespaces.Items {
		// List all deployments for each namespace
		deployments, err := client.AppsV1().Deployments(ns.Name).List(ctx, metav1.ListOptions{})
		
		if err != nil {
			continue
		}
		for _, deployment := range deployments.Items {
			// fmt.Println("1111111111111111111111",reflect.ValueOf(deployment).Kind())
			if (deployment.Status.ReadyReplicas == 0) &&  (deployment.Status.AvailableReplicas == 0){

				if deployment.Name == "coredns" {
					continue
				} else {
					status := models.DeploymentStatus{
						Name:      deployment.Name,
						Namespace: deployment.Namespace,
						Status:    "",
						Timestamp: time.Now(),
					}
					deploymentStatuses = append(deploymentStatuses, status)
					
				}
		}
	}
	}

	return deploymentStatuses
}