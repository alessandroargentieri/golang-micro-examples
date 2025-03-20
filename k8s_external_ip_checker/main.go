package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// getKubeClient creates and returns a Kubernetes clientset
func getKubeClient() (*kubernetes.Clientset, error) {
	var kubeconfig string
	
	// Try to use default kubeconfig location
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config_myclusterkubeconfig.yaml")
	} else {
		// Otherwise use the KUBECONFIG environment variable
		kubeconfig = os.Getenv("KUBECONFIG")
	}

	// Build config from the kubeconfig file
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("error building kubeconfig: %v", err)
	}

	// Create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("error creating Kubernetes client: %v", err)
	}

	return clientset, nil
}

// getNodes retrieves all nodes in the cluster
func getNodes(clientset *kubernetes.Clientset) ([]corev1.Node, error) {
	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("error listing nodes: %v", err)
	}
	return nodes.Items, nil
}

// getNodeExternalIP creates a pod on the specified node, runs curl ifconfig.me,
// and returns the external IP address
func getNodeExternalIP(clientset *kubernetes.Clientset, nodeName string) (string, error) {
	// Generate a unique pod name using node name and timestamp
	podName := fmt.Sprintf("ip-checker-%s-%d", strings.ToLower(strings.Replace(nodeName, ".", "-", -1)), time.Now().Unix())
	namespace := "default"

	fmt.Printf("Creating pod %s on node %s...\n", podName, nodeName)
	
	// Create a pod specification that will run on the specific node
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      podName,
			Namespace: namespace,
		},
		Spec: corev1.PodSpec{
			NodeName: nodeName,
			Containers: []corev1.Container{
				{
					Name:    "curl",
					Image:   "curlimages/curl:latest",
					Command: []string{"sh", "-c", "curl -s ifconfig.me"},
				},
			},
			RestartPolicy: corev1.RestartPolicyNever,
		},
	}

	// Create the pod
	_, err := clientset.CoreV1().Pods(namespace).Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to create pod: %v", err)
	}

	// Set up pod deletion when function exits
	defer func() {
		err := clientset.CoreV1().Pods(namespace).Delete(context.TODO(), podName, metav1.DeleteOptions{})
		if err != nil {
			fmt.Printf("Warning: failed to delete pod %s: %v\n", podName, err)
		} else {
			fmt.Printf("Deleted pod %s\n", podName)
		}
	}()

	// Wait for pod to complete execution
	fmt.Printf("Waiting for pod %s to complete...\n", podName)
	var podCompleted bool
	for i := 0; i < 60; i++ { // Wait up to 2 minutes (60 * 2 seconds)
		podStatus, err := clientset.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
		if err != nil {
			return "", fmt.Errorf("failed to get pod status: %v", err)
		}
		
		if podStatus.Status.Phase == corev1.PodSucceeded {
			podCompleted = true
			break
		}
		
		if podStatus.Status.Phase == corev1.PodFailed {
			return "", fmt.Errorf("pod failed to run")
		}
		
		time.Sleep(2 * time.Second)
	}

	if !podCompleted {
		return "", fmt.Errorf("pod did not complete within timeout")
	}

	// Get pod logs to retrieve the external IP
	fmt.Printf("Getting logs from pod %s...\n", podName)
	req := clientset.CoreV1().Pods(namespace).GetLogs(podName, &corev1.PodLogOptions{})
	podLogs, err := req.Stream(context.TODO())
	if err != nil {
		return "", fmt.Errorf("failed to get pod logs: %v", err)
	}
	defer podLogs.Close()

	buf := new(strings.Builder)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		return "", fmt.Errorf("failed to copy pod logs: %v", err)
	}

	externalIP := strings.TrimSpace(buf.String())
	if externalIP == "" {
		return "", fmt.Errorf("pod returned empty IP")
	}
	
	return externalIP, nil
}

func main() {
	// Get the Kubernetes client
	fmt.Println("Connecting to Kubernetes cluster...")
	clientset, err := getKubeClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully connected to Kubernetes cluster")

	// Get all nodes
	fmt.Println("Getting nodes...")
	nodes, err := getNodes(clientset)
	if err != nil {
		fmt.Printf("Error getting nodes: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Found %d nodes in the cluster\n", len(nodes))

	// Map to store node names and their external IPs
	nodeExternalIPs := make(map[string]string)
	var mutex sync.Mutex
	var wg sync.WaitGroup

	// Process each node in parallel
	for _, node := range nodes {
		wg.Add(1)
		go func(nodeName string) {
			defer wg.Done()

			fmt.Printf("Processing node %s...\n", nodeName)
			externalIP, err := getNodeExternalIP(clientset, nodeName)
			
			mutex.Lock()
			defer mutex.Unlock()
			
			if err != nil {
				fmt.Printf("Error getting external IP for node %s: %v\n", nodeName, err)
				nodeExternalIPs[nodeName] = "Error: " + err.Error()
			} else {
				nodeExternalIPs[nodeName] = externalIP
				fmt.Printf("Node %s external IP: %s\n", nodeName, externalIP)
			}
		}(node.Name)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Print the final map of nodes and their external IPs
	fmt.Println("\nNode External IPs:")
	fmt.Println("====================")
	for nodeName, externalIP := range nodeExternalIPs {
		fmt.Printf("%s: %s\n", nodeName, externalIP)
	}
}

