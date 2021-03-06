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

type ServiceAccountResource struct {
	Client   v1Typed.ServiceAccountInterface
	Resource *v1.ServiceAccount
	Error    error
}

func (r *ServiceAccountResource) GetObject() runtime.Object {
	//fmt.Printf("%#v\n\n", r.Resource)
	return r.Resource
}

func (r *ServiceAccountResource) GetError() error {
	return r.Error
}

func (r *ServiceAccountResource) GetResourceName() string {
	return r.Resource.Name
}

func (r *ServiceAccountResource) GetResourceKind() string {
	kind := strings.Split(fmt.Sprintf("%T", r.Resource), ".")
	return kind[len(kind)-1 : len(kind)][0]
}

func (r *ServiceAccountResource) IsReady() bool {
	if r.Resource.CreationTimestamp.IsZero() {
		return false
	}
	return true
}

func (r *ServiceAccountResource) GetClient(namespace string) {
	r.Client = clientset.CoreV1().ServiceAccounts(namespace)
}

func (r *ServiceAccountResource) Get() {
	resource, err := r.Client.Get(context.TODO(), r.Resource.Name, metav1.GetOptions{})
	if getHandler(r.Resource.Kind, r.Resource.Name, err) {
		r.Resource = resource
		return
	}
	r.Error = err
}
func (r *ServiceAccountResource) Create() {
	result, err := r.Client.Create(context.TODO(), r.Resource, metav1.CreateOptions{})
	r.Error = err
	r.Resource = result
}
func (r *ServiceAccountResource) Update() {
}
func (r *ServiceAccountResource) Delete() {
}
