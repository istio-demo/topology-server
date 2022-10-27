package main

import (
	"context"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	podName := os.Getenv("POD_NAME")
	nodeName := os.Getenv("NODE_NAME")

	if podName == "" || nodeName == "" {
		panic("need POD_NAME and NODE_NAME environments!")
	}

	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	log.Println("get node info for", nodeName)
	node, err := clientset.CoreV1().Nodes().Get(ctx, nodeName, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}
	labels := node.GetObjectMeta().GetLabels()
	if labels == nil {
		panic("no label found for node " + nodeName)
	}
	region := labels["topology.kubernetes.io/region"]
	zone := labels["topology.kubernetes.io/zone"]
	if region == "" || zone == "" {
		panic("region or zone not found for node " + nodeName)
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "%s/%s - %s", region, zone, podName)
	})
	log.Fatal(r.Run("0.0.0.0:80"))
}
