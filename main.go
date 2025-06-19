package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"time"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/api/core/v1"
)

const (
	watchDir     = "/config"
	configMapName = "dynamic-properties"
	namespace     = "default"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "/kubeconfig/config")
	if err != nil {
		log.Fatalf("Error loading kubeconfig: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating clientset: %v", err)
	}

	for {
		files, err := filepath.Glob(filepath.Join(watchDir, "*.properties"))
		if err != nil {
			log.Printf("Error scanning directory: %v", err)
			continue
		}

		data := make(map[string]string)
		for _, file := range files {
			content, err := ioutil.ReadFile(file)
			if err != nil {
				log.Printf("Error reading file %s: %v", file, err)
				continue
			}
			filename := filepath.Base(file)
			data[filename] = string(content)
		}

		err = updateConfigMap(clientset, data)
		if err != nil {
			log.Printf("Error updating ConfigMap: %v", err)
		}

		time.Sleep(10 * time.Second)
	}
}

func updateConfigMap(clientset *kubernetes.Clientset, data map[string]string) error {
	cmClient := clientset.CoreV1().ConfigMaps(namespace)

	cm, err := cmClient.Get(context.Background(), configMapName, metav1.GetOptions{})
	if err == nil {
		cm.Data = data
		_, err = cmClient.Update(context.Background(), cm, metav1.UpdateOptions{})
	} else {
		cm = &corev1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Name: configMapName,
			},
			Data: data,
		}
		_, err = cmClient.Create(context.Background(), cm, metav1.CreateOptions{})
	}
	return err
}
