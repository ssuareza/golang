package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	log.SetFlags(0)

	// variables
	// fill your values here
	file := "configmap.yaml"
	context := "security"
	namespace := "security"

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
	client, err := dynamic.NewForConfig(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// create resource
	err = create(file, cfg, client, namespace)
	if err != nil {
		log.Fatal(err)
	}
}

func create(file string, cfg *rest.Config, client dynamic.Interface, namespace string) error {
	// get file content
	f, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	// and decode
	decode := scheme.Codecs.UniversalDeserializer().Decode
	obj, _, err := decode(f, nil, nil)
	if err != nil {
		return err
	}

	// actions to take based in object kind
	switch o := obj.(type) {
	case *v1.Pod:
		fmt.Printf("%s\n", o.GetObjectKind())
	case *v1.ConfigMap:
		fmt.Printf("%s\n", o.GetObjectKind())
		fmt.Println(o.GetObjectKind().GroupVersionKind())
	case *v1.Secret:
		fmt.Printf("%s\n", o.GetObjectKind())
	case *v1.Service:
		fmt.Printf("%s\n", o.GetObjectKind())
	case *v1.ServiceAccount:
		fmt.Printf("%s\n", o.GetObjectKind())
	case *appsv1.Deployment:
		fmt.Printf("%s\n", o.GetObjectKind())
	case *batchv1.Job:
		fmt.Printf("%s\n", o.GetObjectKind())
	default:
		fmt.Printf("%s\n", o.GetObjectKind())
	}

	return nil
}
