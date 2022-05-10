package main

import (
	"flag"
	"path/filepath"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func loadKubeConfig() {
	kubeconfig_path := flag.String(
		"kubeconfig",
		filepath.Join(homedir.HomeDir(), ".kube", "config"),
		"Absolute path of kubeconfig",
	)

	flag.Parse()
	// load the kubeconfig
	config := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: *kubeconfig_path},
		&clientcmd.ConfigOverrides{
			CurrentContext: "",
		},
	)

	context, err := config.RawConfig()
	if err != nil {
		panic(err)
	}

	// do something with the contexts
	context.Contexts
}
