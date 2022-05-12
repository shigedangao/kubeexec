package kube

import (
	"os"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
)

// Execute a command in a targeted pod within a targeted namespace
func (c *Client) ExecCommand(pod_name string, namespace string, cmd []string) error {
	req := c.clientset.CoreV1().RESTClient().Post().Resource("pods").Name(pod_name).Namespace(namespace).SubResource("exec")
	opt := &v1.PodExecOptions{
		Command: cmd,
		Stdin:   true,
		Stdout:  true,
		Stderr:  true,
		TTY:     true,
	}

	req.VersionedParams(
		opt,
		scheme.ParameterCodec,
	)

	exec, err := remotecommand.NewSPDYExecutor(c.clientconfig, "POST", req.URL())
	if err != nil {
		return err
	}

	err = exec.Stream(remotecommand.StreamOptions{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	})

	if err != nil {
		return err
	}

	return nil
}
