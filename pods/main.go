package main

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {

	pod := v1.Pod{}
	fmt.Println(IsHidden(pod.GetObjectMeta()))
	fmt.Println(IsHidden2(pod.ObjectMeta))
}

func IsHidden(o metav1.Object) bool {
	//var mapp map[string]string = nil
	//return mapp["hello"] == "hi"
	return o.GetLabels()["hidden"] == "true"
}

func IsHidden2(o metav1.ObjectMeta) bool {
	return o.Labels["hidden"] == "true"
}
