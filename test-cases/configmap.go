package test_cases

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	v1Typed "k8s.io/client-go/kubernetes/typed/core/v1"
	"strings"
)

type ConfigMapResource struct {
	Client   v1Typed.ConfigMapInterface
	Resource *v1.ConfigMap
	Error    error
}

func (r *ConfigMapResource) GetObject() runtime.Object {
	//fmt.Printf("%#v\n\n", r.Resource)
	return r.Resource
}

func (r *ConfigMapResource) GetError() error {
	return r.Error
}

func (r *ConfigMapResource) GetResourceName() string {
	return r.Resource.Name
}

func (r *ConfigMapResource) GetResourceKind() string {
	kind := strings.Split(fmt.Sprintf("%T", r.Resource), ".")
	return kind[len(kind)-1 : len(kind)][0]
}

func (r *ConfigMapResource) IsReady() bool {
	r.Get()
	if r.Resource.CreationTimestamp.IsZero() {
		return false
	}
	return true
}

func (r *ConfigMapResource) GetClient(namespace string) {
	r.Client = clientset.CoreV1().ConfigMaps(namespace)
}

func (r *ConfigMapResource) Get() {
	resource, err := r.Client.Get(context.TODO(), r.Resource.Name, metav1.GetOptions{})
	if getHandler(r.Resource.Kind, r.Resource.Name, err) {
		r.Resource = resource
		return
	}
	r.Error = err
}
func (r *ConfigMapResource) Create() {
	result, err := r.Client.Create(context.TODO(), r.Resource, metav1.CreateOptions{})
	r.Error = err
	r.Resource = result
}
func (r *ConfigMapResource) Update() {

}
func (r *ConfigMapResource) Delete() {
}

func rawConfigMap() {

}
