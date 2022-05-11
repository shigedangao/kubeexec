package term

import (
	"fmt"
	"kubeexec/kube"

	survey "github.com/AlecAivazis/survey/v2"
)

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

	fmt.Println(pod)

	return nil
}

func selectContexts(choices []string) string {
	context := ""
	prompt := &survey.Select{
		Message: "Choose a kubernetes context",
		Options: choices,
	}

	survey.AskOne(prompt, &context)

	return context
}

func selectNamespace(client *kube.Client) (string, error) {
	ns := ""
	namespaces, err := client.ListNamespaces()
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

func selectPod(ns string, client *kube.Client) (string, error) {
	pod := ""
	pod_list, err := client.ListPodForNamespace(ns)
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
