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
	_, err := kube.CreateConfigFromCustomContext(ctx, kubeconfig_path)
	if err != nil {
		return err
	}

	fmt.Println("ok")
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
