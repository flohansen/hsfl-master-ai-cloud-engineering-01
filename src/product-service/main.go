package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/flohansen/hsfl-master-ai-cloud-engineering/lib/database"
	"github.com/flohansen/hsfl-master-ai-cloud-engineering/product-service/api/router"
	"github.com/flohansen/hsfl-master-ai-cloud-engineering/product-service/products"
)

func GetenvInt(key string) int {
	value := os.Getenv(key)
	valueInt, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}

	return valueInt
}

func main() {
	config := database.PsqlConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     GetenvInt("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Database: os.Getenv("DB_NAME"),
	}

	productRepository, err := products.NewPsqlRepository(config)
	if err != nil {
		log.Fatalf("could not create product repo: %s", err.Error())
	}

	productsController := products.NewDefaultController(productRepository)
	handler := router.New(productsController)

	if err := productRepository.Migrate(); err != nil {
		log.Fatalf("could not migrate: %s", err.Error())
	}

	if err := http.ListenAndServe(":3000", handler); err != nil {
		log.Fatalf("error while listen and serve: %s", err.Error())
	}
}
