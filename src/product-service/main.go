package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/flohansen/hsfl-master-ai-cloud-engineering/lib/database"
	"github.com/flohansen/hsfl-master-ai-cloud-engineering/product-service/api/router"
	"github.com/flohansen/hsfl-master-ai-cloud-engineering/product-service/products"
	"gopkg.in/yaml.v3"
)

type ApplicationConfig struct {
	Database database.PsqlConfig `yaml:"database"`
}

func LoadConfigFromFile(path string) (*ApplicationConfig, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var config ApplicationConfig
	if err := yaml.NewDecoder(f).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func main() {
	configPath := flag.String("config", "config.yml", "The path to the configuration file")
	flag.Parse()

	config, err := LoadConfigFromFile(*configPath)
	if err != nil {
		log.Fatalf("could not load application configuration: %s", err.Error())
	}

	productRepository, err := products.NewPsqlRepository(config.Database)
	productsController := products.NewDefaultController(productRepository)
	handler := router.New(productsController)

	if err := productRepository.Migrate(); err != nil {
		log.Fatalf("could not migrate: %s", err.Error())
	}

	if err := http.ListenAndServe(":3000", handler); err != nil {
		log.Fatalf("error while listen and serve: %s", err.Error())
	}
}
