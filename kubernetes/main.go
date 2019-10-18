package main

import (
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type k8s struct {
	clientset kubernetes.Interface
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

func (o *k8s) getVersion() (string, error) {
	version, err := o.clientset.Discovery().ServerVersion()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", version), nil
}

func (o *k8s) getNamespaces() {
	namespaces, err := o.clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error listing namespaces: %v", err)
		os.Exit(1)
	}

	// Print type
	//fmt.Printf("%T\n", namespaces)

	for _, namespace := range namespaces.Items {
		fmt.Println(namespace.Name)
	}

	//return fmt.Sprintf("%s", namespaces), nil
}

func (o *k8s) getPods() {
	pods, err := o.clientset.CoreV1().Pods("security").List(metav1.ListOptions{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error listing pods: %v", err)
		os.Exit(1)
	}

	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
	}

	//return fmt.Sprintf("%s", pods), nil
}

func (o *k8s) getNodes() {
	nodes, err := o.clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error listing nodes: %v", err)
		os.Exit(1)
	}

	for _, node := range nodes.Items {
		fmt.Printf("Node: %s\n", node.Name)
	}

	//return fmt.Sprintf("%s", nodes), nil
}


// Logs
type PodExpansion interface {
    Bind(binding *v1.Binding) error
    Evict(eviction *policy.Eviction) error
    GetLogs(name string, opts *v1.PodLogOptions) *restclient.Request
}

func (o *k8s, pod corev1.Pod) getLogs() {
	podLogOpts := corev1.PodLogOptions{}
	pod.Namespace := "security"
	pod.Name := "workermailer-20191014084950-64b5c84cf8-zg4gv"
	//podLogOpts := "{Follow: true, Container: 'workermailer'}"
	logs, err := o.clientset.CoreV1().Pods(pod.Namespace).GetLogs(pod.Name, &podLogOpts)
	//logs, err := o.clientset.Core
	if err != nil {
		fmt.Fprintf(os.Stderr, "error getting logs: %v", err)
		os.Exit(1)
	}

	fmt.Println(logs)
	//return fmt.Sprintf("%s", logs), nil
	/*for _, node := range nodes.Items {
		fmt.Printf("Node: %s\n", node.Name)
	}*/
}

func main() {
	// init
	k8s, err := connect()
	if err != nil {
		fmt.Println(err)
		return
	}

	// getVersion
	/*v, err := k8s.getVersion()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)*/

	// getNodes
	//k8s.getNodes()

	// getPods
	//k8s.getPods()

	// getNamespaces
	//k8s.getNamespaces()

	// getLogs
	k8s.getLogs()
}
