package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var wg sync.WaitGroup

func main() {
	log.SetFlags(0)

	// variables
	// fill your values here
	context := "YOUR_KUBECONFIG_CONTEXT"
	namespace := "NAMESPACE"
	label := "POD_LABEL"

	// kubeconfig
	kubeconfig := os.Getenv("HOME") + "/.kube/config"
	cfg, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	// switching context
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	cfgOverrides := &clientcmd.ConfigOverrides{CurrentContext: context}
	override := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, cfgOverrides)
	cfg, err = override.ClientConfig()

	// clientset
	clientset, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// list pods = podLabel
	podLabel := "app=" + label

	pods, err := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{LabelSelector: podLabel})
	if err != nil {
		log.Fatal(err)
	}

	// save list of pods
	if len(pods.Items) == 0 {
		fmt.Println("No pods found")
		os.Exit(0)
	}

	podList := []string{}
	for _, pod := range pods.Items {
		podList = append(podList, pod.Name)
	}

	// get logs
	tail := int64(10)
	podLogOpts := v1.PodLogOptions{
		Follow:    true,
		Container: label,
		TailLines: &tail,
	}

	var logs *rest.Request
	for _, pod := range podList {
		logs = clientset.CoreV1().Pods(namespace).GetLogs(pod, &podLogOpts)
		wg.Add(1)
		// print pod logs in parallel with goroutines
		go print(logs)
	}
	wg.Wait()
}

func print(logs *rest.Request) {
	defer wg.Done()
	readCloser, err := logs.Stream()
	if err != nil {
		fmt.Println(err)
	} else {
		buf := make([]byte, 4096)

		for {
			n, err := readCloser.Read(buf)
			if err == io.EOF {
				break
			}
			if n != 0 {
				fmt.Print(string(buf[:n]))
			}
		}
	}
}
