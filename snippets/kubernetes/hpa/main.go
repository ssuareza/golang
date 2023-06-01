package main

import (
	"fmt"
	"log"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	log.SetFlags(0)

	// variables
	// fill your values here
	namespace := "NAMESPACE"
	label := "POD_LABEL"

	// kubeconfig
	kubeconfig := os.Getenv("HOME") + "/.kube/config"
	cfg, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	// clientset
	clientset, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// get hpa details
	hpa, err := clientset.AutoscalingV2beta2().HorizontalPodAutoscalers(namespace).Get(label, metav1.GetOptions{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hpa.DeepCopy().Status.DesiredReplicas)

}
