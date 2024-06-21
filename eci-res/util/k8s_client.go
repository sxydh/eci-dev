package util

import (
	"eci-res/model"
	"flag"
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var k8sClient *kubernetes.Clientset
var k8sConfig *rest.Config
var err error

func init() {
	/* k8s 客户端初始化 */
	initK8sClient()
}

func initK8sClient() {
	/* 构建配置 */
	var configFlag *string
	if home := homedir.HomeDir(); home != "" {
		configFlag = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "")
	} else {
		configFlag = flag.String("kubeconfig", "C:/Users/Administrator/.kube/config", "")
	}
	flag.Parse()
	k8sConfig, err = clientcmd.BuildConfigFromFlags("", *configFlag)
	if err != nil {
		log.Fatalf("BuildConfigFromFlags error: %v", err)
	}

	/* 构建客户端 */
	k8sClient, err = kubernetes.NewForConfig(k8sConfig)
	if err != nil {
		log.Fatalf("NewForConfig error: %v", err)
	}
}

func GetK8sConfig() *rest.Config {
	return k8sConfig
}
func GetK8sClient() *kubernetes.Clientset {
	return k8sClient
}

func GetPodRestarts(pod v1.Pod) int32 {
	restarts := int32(0)
	for _, containerStatus := range pod.Status.ContainerStatuses {
		restarts += containerStatus.RestartCount
	}
	return restarts
}

func GetPodReady(pod *v1.Pod) string {
	readyContainers := 0
	totalContainers := len(pod.Spec.Containers)
	for _, containerStatus := range pod.Status.ContainerStatuses {
		if containerStatus.Ready {
			readyContainers++
		}
	}
	if totalContainers == 0 {
		return "0/0"
	}
	return fmt.Sprintf("%d/%d", readyContainers, totalContainers)
}

func GetAge(creationTime time.Time) string {
	now := time.Now()
	duration := now.Sub(creationTime)
	dayCount := duration / (24 * time.Hour)
	hourCount := (duration % (24 * time.Hour)) / time.Hour
	return fmt.Sprintf("%dd%dh", dayCount, hourCount)
}

func GetPodContainers(pod *v1.Pod) []model.Container {
	var getContainerStatus = func(containerName string) *v1.ContainerStatus {
		for _, containerStatus := range pod.Status.ContainerStatuses {
			if containerStatus.Name == containerName {
				return &containerStatus
			}
		}
		return nil
	}

	containers := make([]model.Container, len(pod.Spec.Containers))
	for i, container := range pod.Spec.Containers {
		containerStatus := getContainerStatus(container.Name)
		containers[i] = model.Container{
			Name:            container.Name,
			Image:           container.Image,
			ImagePullPolicy: string(container.ImagePullPolicy),
			Command:         strings.Join(container.Command, " "),
			Ready:           strconv.FormatBool(containerStatus.Ready),
			Restarts:        strconv.FormatInt(int64(containerStatus.RestartCount), 10),
			ResourceRequest: model.Resource{
				Cpu:    container.Resources.Requests.Cpu().String(),
				Memory: container.Resources.Requests.Memory().String(),
			},
			ResourceLimit: model.Resource{
				Cpu:    container.Resources.Limits.Cpu().String(),
				Memory: container.Resources.Limits.Memory().String(),
			},
		}
	}
	return containers
}
