// Command sched starts the chainr scheduler.
//
// The scheduler allows to run jobs on the Kubernetes cluster it is deployed on.
// It manages the whole dependency tree, and checks all preconditions are met
// before running a job.
// It exposes an API allowing to run a pipeline, and get its execution status.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Tyrame/chainr/sched/internal/config"
	"github.com/Tyrame/chainr/sched/internal/k8s"
)

var configFile string = "config.yaml"

func init() {
	val, ok := os.LookupEnv("CONFIG_FILE")
	if ok {
		configFile = val
	}
}

func main() {
	log.Println("Starting chainr scheduler")
	cfg, err := config.Load(configFile)
	if err != nil {
		log.Println("Configuration loading failed:", err.Error())
		log.Println("Using default configuration")
	}

	k8s, err := k8s.New(cfg.Kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Println("Listening on", addr)
	log.Fatal(http.ListenAndServe(addr, NewHandler(k8s)))
}
