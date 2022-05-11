package kube

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) ListPodForNamespace(namespace string) ([]string, error) {
	ctx := context.Background()
	pods, err := c.clientset.CoreV1().Pods(namespace).List(ctx, v1.ListOptions{})
	if err != nil {
		return nil, err
	}

	pod_list := make([]string, 1)
	for _, pod := range pods.Items {
		pod_list = append(pod_list, pod.ObjectMeta.Name)
	}

	return pod_list, nil
}
