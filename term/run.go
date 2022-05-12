package term

import (
	"errors"
	"kubeexec/kube"

	survey "github.com/AlecAivazis/survey/v2"
)

// Run the scenario to execute a command in a kubernetes pod
func RunScenario() error {
	kubeconfig_path := kube.GetKubeConfigLocalPath()
	choices := kube.GetKubeConfigCtx(kubeconfig_path)
	// select the context
	ctx := selectContexts(choices)
	// get a list of namespaces
	client, err := kube.CreateConfigFromCustomContext(ctx, kubeconfig_path)
	if err != nil {
		return err
	}

	ns, err := selectNamespace(&client)
	if err != nil {
		return err
	}

	pod, err := selectPod(ns, &client)
	if err != nil {
		return err
	}

	cmd := selectCommand()
	err = client.ExecCommand(pod, ns, cmd)
	if err != nil {
		return err
	}

	return nil
}

// Let the user select the kubernetes cluster
func selectContexts(choices []string) string {
	context := ""
	prompt := &survey.Select{
		Message: "Choose a kubernetes context",
		Options: choices,
	}

	survey.AskOne(prompt, &context)

	return context
}

// Let the user select the namespace
func selectNamespace(client *kube.Client) (string, error) {
	ns := ""
	namespaces, err := client.ListNamespaces()
	if len(namespaces) == 0 {
		return "", errors.New("no namespaces can be founded on the selected context")
	}

	if err != nil {
		return "", nil
	}

	prompt := &survey.Select{
		Message: "Choose a kubernetes namespace",
		Options: namespaces,
	}

	survey.AskOne(prompt, &ns)

	return ns, nil
}

// Select the pod to execute a command
func selectPod(ns string, client *kube.Client) (string, error) {
	pod := ""
	pod_list, err := client.ListPodForNamespace(ns)
	if len(pod_list) == 0 {
		return "", errors.New("no pod can be founded on the selected namespace")
	}

	if err != nil {
		return "", err
	}

	prompt := &survey.Select{
		Message: "Choose a pod",
		Options: pod_list,
	}

	survey.AskOne(prompt, &pod)

	return pod, nil
}

// Select a command (only two option could add more later)
func selectCommand() []string {
	cmd := ""
	prompt := &survey.Select{
		Message: "Choose a kubernetes namespace",
		Options: []string{"/bin/sh", "/bin/bash"},
	}

	survey.AskOne(prompt, &cmd)

	return []string{cmd}
}
