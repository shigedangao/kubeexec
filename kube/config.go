package kube

import (
	"flag"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type Client struct {
	clientset    *kubernetes.Clientset
	clientconfig *rest.Config
}

// Get the kubeconfig location
func GetKubeConfigLocalPath() string {
	kubeconfig_path := flag.String(
		"kubeconfig",
		filepath.Join(homedir.HomeDir(), ".kube", "config"),
		"Absolute path of kubeconfig",
	)

	flag.Parse()

	return *kubeconfig_path
}

// Read the kubeconfig and return the list of available
// kubeconfig contexts
func GetKubeConfigCtx(kubeconfig_path string) []string {
	// load the kubeconfig
	config := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfig_path},
		&clientcmd.ConfigOverrides{
			CurrentContext: "",
		},
	)

	context, err := config.RawConfig()
	if err != nil {
		panic(err)
	}

	kubeconfig_ctx := make([]string, 1)
	for key := range context.Contexts {
		kubeconfig_ctx = append(kubeconfig_ctx, key)
	}

	return kubeconfig_ctx
}

// Create a kubernetes client wrapper from the selected kubeconfig context
func CreateConfigFromCustomContext(ctx string, kubeconfig_path string) (Client, error) {
	// load the kubeconfig
	config := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfig_path},
		&clientcmd.ConfigOverrides{
			CurrentContext: ctx,
		},
	)

	client_config, err := config.ClientConfig()
	if err != nil {
		return Client{}, err
	}

	clientset, err := kubernetes.NewForConfig(client_config)
	if err != nil {
		return Client{}, err
	}

	return Client{
		clientset:    clientset,
		clientconfig: client_config,
	}, nil
}
