package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type IndexPageViewModel struct {
	Products []Product
}

type Product struct {
	Name     string
	Retailer string
	Price    float32
}

func requestProducts() ([]Product, error) {
	endpoint := fmt.Sprintf("http://%s/api/v1/products", os.Getenv("PRODUCTS_ENDPOINT"))

	req, _ := http.NewRequest("GET", endpoint, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var products []Product
	if err := json.NewDecoder(res.Body).Decode(&products); err != nil {
		return nil, err
	}

	return products, nil
}

func indexHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := requestProducts()
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		viewModel := IndexPageViewModel{
			Products: products,
		}

		tmpl.ExecuteTemplate(w, "index", viewModel)
	}
}

func main() {
	tmpl := template.Must(template.ParseGlob("templates/*.gohtml"))

	router := http.NewServeMux()
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))
	router.HandleFunc("/", indexHandler(tmpl))

	log.Fatal(http.ListenAndServe("0.0.0.0:3000", router))
}
