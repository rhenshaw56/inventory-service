package product

import (
	"strconv"
	"strings"
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/rhenshaw56/inventory-service/cors"
)

const productBasePath = "products"

// SetupRoutes  should Set up Routes
func SetupRoutes(apiBasePath string) {
	handleProducts := http.HandlerFunc(productsHandler)
	handleProduct := http.HandlerFunc(productHandler)

	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, productBasePath), cors.Middleware(handleProducts))
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, productBasePath), cors.Middleware(handleProduct))
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		productList := getProductList()
		productsJSON, err := json.Marshal(productList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(productsJSON)

	case http.MethodPost:
		var newProduct Product
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		err = json.Unmarshal(bodyBytes, &newProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		if newProduct.ProductID != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		_, err = addOrUpdateProduct(newProduct)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// newProduct.ProductID = getNextID()
		// productList = append(productList, newProduct)

		w.WriteHeader(http.StatusCreated)
		return

	case http.MethodOptions:
		return

	}
	// w.Write([]byte("product hannlder called called"))
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, "products/")
	fmt.Println(urlPathSegments)

	fmt.Println(len(urlPathSegments))
	
	productID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	product := getProduct(productID)

	if product == nil {
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, fmt.Sprintf("no product with id %d", productID), http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		productJSON, err := json.Marshal(product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(productJSON)

	case http.MethodPut:
		var updatedProduct Product
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		err = json.Unmarshal(bodyBytes, &updatedProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if updatedProduct.ProductID != productID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		addOrUpdateProduct(updatedProduct)
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodDelete:
		removeProduct(productID)
		return

	case http.MethodOptions:
		return
	
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}



}