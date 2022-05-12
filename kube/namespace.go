package kube

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// List the namespace and return the name of the namespaces
func (c *Client) ListNamespaces() ([]string, error) {
	ctx := context.Background()
	namespaces, err := c.clientset.CoreV1().Namespaces().List(ctx, v1.ListOptions{})
	if err != nil {
		return nil, err
	}

	namespace_name := make([]string, 1)
	for _, ns := range namespaces.Items {
		namespace_name = append(namespace_name, ns.ObjectMeta.Name)
	}

	return namespace_name, nil
}
