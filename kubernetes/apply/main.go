package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	log.SetFlags(0)

	// variables
	// fill your values here
	file := "configmap.yaml"
	context := "YOUR_KUBECONFIG_CONTEXT"
	namespace := "NAMESPACE"

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

	// dynamic client
	client, err := dynamic.NewForConfig(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// create resource
	err = apply(file, cfg, client, namespace)
	if err != nil {
		fmt.Println(err)
	}
}

func apply(file string, cfg *rest.Config, client dynamic.Interface, namespace string) error {
	// get file content
	f, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	// decode
	decode := scheme.Codecs.UniversalDeserializer().Decode
	o, _, err := decode(f, nil, nil)
	if err != nil {
		return err
	}

	// find object kind
	gvk := o.GetObjectKind().GroupVersionKind()
	gk := schema.GroupKind{Group: gvk.Group, Kind: gvk.Kind}
	discover, _ := discovery.NewDiscoveryClientForConfig(cfg)
	groupResources, _ := restmapper.GetAPIGroupResources(discover)
	rm := restmapper.NewDiscoveryRESTMapper(groupResources)
	mapping, err := rm.RESTMapping(gk, gvk.Version)
	if err != nil {
		log.Fatal(err)
	}

	// convert the runtime.Object to unstructured.Unstructured
	u, err := runtime.DefaultUnstructuredConverter.ToUnstructured(o)
	if err != nil {
		return err
	}

	// create resource
	fmt.Println("Creating", gk)
	_, err = client.Resource(mapping.Resource).Namespace(namespace).Create(&unstructured.Unstructured{
		Object: u,
	}, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}
