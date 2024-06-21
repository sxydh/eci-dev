package service

import (
	"context"
	"eci-res/model"
	"eci-res/util"
	"github.com/gin-gonic/gin"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"strconv"
)

type EciService struct {
}

func (s *EciService) QueryEciList() ([]model.Eci, error) {
	client := util.GetK8sClient()
	podList, err := client.CoreV1().Pods("namespace-demo").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var eciList = make([]model.Eci, 0)
	for _, pod := range podList.Items {
		eci := model.Eci{
			Id:          string(pod.UID),
			Name:        pod.ObjectMeta.Labels["app"],
			ReplicaName: pod.Name,
			Ready:       util.GetPodReady(&pod),
			Status:      string(pod.Status.Phase),
			Restarts:    strconv.FormatInt(int64(util.GetPodRestarts(pod)), 10),
			Age:         util.GetAge(pod.CreationTimestamp.Time),
			Ip:          pod.Status.PodIP,
			Node:        pod.Spec.NodeName,
			Containers:  util.GetPodContainers(&pod),
		}
		eciList = append(eciList, eci)
	}
	return eciList, nil
}

func (s *EciService) AddEci(c *gin.Context) error {
	type reqBodyStruct struct {
		Name   string `json:"name"`
		Image  string `json:"image"`
		Cpu    string `json:"cpu"`
		Memory string `json:"memory"`
	}
	var reqBody reqBodyStruct
	err := c.BindJSON(&reqBody)
	if err != nil {
		log.Printf("BindJSON to request body when AddEci error: err=%v", err)
		return err
	}

	int32Ptr := func(i int32) *int32 {
		return &i
	}

	client := util.GetK8sClient()
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      reqBody.Name,
			Namespace: "namespace-demo",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": reqBody.Name,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": reqBody.Name,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  reqBody.Name,
							Image: reqBody.Image,
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: 8080,
								},
							},
							Resources: corev1.ResourceRequirements{
								Requests: corev1.ResourceList{
									// 1 核等于 1000 豪核
									corev1.ResourceCPU:    resource.MustParse(reqBody.Cpu),
									corev1.ResourceMemory: resource.MustParse(reqBody.Memory),
								},
								Limits: corev1.ResourceList{
									corev1.ResourceCPU:    resource.MustParse(reqBody.Cpu),
									corev1.ResourceMemory: resource.MustParse(reqBody.Memory),
								},
							},
						},
					},
				},
			},
		},
	}
	_, err = client.AppsV1().Deployments("namespace-demo").Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (s *EciService) DeleteEci(c *gin.Context) error {
	name := c.Query("name")

	client := util.GetK8sClient()
	deletePropagation := metav1.DeletePropagationForeground
	deleteOptions := metav1.DeleteOptions{
		PropagationPolicy: &deletePropagation,
	}
	err := client.AppsV1().Deployments("namespace-demo").Delete(context.TODO(), name, deleteOptions)
	if err != nil {
		log.Printf("Delete deployment error: name=%v, err=%v", name, err)
		return err
	}
	return nil
}
