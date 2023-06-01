package main

import (
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

// NewFakeClient returns a fake client for mock testing
func TestGetPods(t *testing.T) {
	// fake client
	clientset := fake.NewSimpleClientset()

	// create fake pod
	_, err := clientset.CoreV1().Pods("testing").Create(
		&v1.Pod{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Pod",
				APIVersion: "v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "fake_pod",
				Labels: map[string]string{
					"tag": "",
				},
			},
		})

	pods, err := getPods(clientset, "testing")
	if err != nil || len(pods.Items) != 1 {
		t.Fail()
	}
}
