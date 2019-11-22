package main

import (
	"fmt"
	"log"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type k8s struct {
	clientset kubernetes.Interface
}

func main() {
	// init
	k8s, err := connect()
	if err != nil {
		fmt.Println(err)
		return
	}

	// get version
	k8s.version()

	// get nodes
	//k8s.nodes()

	// get pods
	//k8s.pods()

	// get namespaces
	//k8s.namespaces()
}

func connect() (*k8s, error) {
	path := os.Getenv("HOME") + "/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", path)
	if err != nil {
		return nil, err
	}
	client := k8s{}
	client.clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (o *k8s) version() {
	version, err := o.clientset.Discovery().ServerVersion()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version)
}

func (o *k8s) namespaces() {
	namespaces, err := o.clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, namespace := range namespaces.Items {
		fmt.Println(namespace.Name)
	}
}

func (o *k8s) pods() {
	pods, err := o.clientset.CoreV1().Pods("security").List(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
	}
}

func (o *k8s) nodes() {
	nodes, err := o.clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, node := range nodes.Items {
		fmt.Printf("Node: %s\n", node.Name)
	}
}
