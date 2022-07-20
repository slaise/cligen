package client

import (
	"fmt"
	"os"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/tools/clientcmd"
    restclient "k8s.io/client-go/rest"

	// utilities for kubernetes integration
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

// ClientConfig
func ClientConfig() (*restclient.Config, error) {
		var kubeconfig *string
    	if home := homedir.HomeDir(); home != "" {
    		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
    	} else {
    		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
    	}
    	flag.Parse()

    	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
    	if err != nil {
    		return nil, nil, err
    	}
    	return config, nil
}

// InitClient - Kubernetes Client
func InitClient() *kubernetes.Clientset {
	cfg := ClientConfig()

	clientset, err := kubernetes.NewForConfig(cfg)
    	if err != nil {
    		panic(err.Error())
    	}

    return clientset
}

func CreateDynamicClient() (*discovery.DiscoveryClient, dynamic.Interface, error) {
    cfg := ClientConfig()
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(cfg)
	if err != nil {
		return nil, nil, err
	}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, nil, err
	}

	return discoveryClient, dynamicClient, err
}