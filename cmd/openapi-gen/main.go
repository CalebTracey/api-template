package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/calebtracey/api-template/internal/routes"
	"github.com/ghodss/yaml"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
)

func main() {
	var output string
	flag.StringVar(&output, "path", "", "Path to use for generating OpenAPI 3 files")
	flag.Parse()
	if output == "" {
		log.Fatalln("path is required")
	}
	swagger := routes.NewOpenAPI3()

	// openapi3.json
	data, err := json.Marshal(&swagger)
	if err != nil {
		log.Fatalf("Couldn't marshal json: %s", err)
	}
	if err := os.WriteFile(path.Join(output, "openapi3.json"), data, 0644); err != nil {
		log.Fatalf("Couldn't write json: %s", err)
	}

	// openapi3.yaml
	data, err = yaml.Marshal(&swagger)
	if err != nil {
		log.Fatalf("Couldn't marshal json: %s", err)
	}
	if err := os.WriteFile(path.Join(output, "openapi3.yaml"), data, 0644); err != nil {
		log.Fatalf("Couldn't write json: %s", err)
	}

	fmt.Println("all generated")
}
