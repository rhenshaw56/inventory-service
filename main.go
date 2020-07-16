package main

import (
	"net/http"
	"github.com/rhenshaw56/inventory-service/product"
)


// var productList []Product

// func init() {
// 	productsJSON := `[
// 		{
// 			"productId": 1,
// 			"manufacturer": "Big Box Company 1",
// 			"sku": "6861qhJK",
// 			"upc": "12345679987",
// 			"pricePerUnit": "12.99",
// 			"quantityOnHand": 28,
// 			"productName": "Gizmo"
// 		},
// 		{
// 			"productId": 2,
// 			"manufacturer": "Big Box Company 2",
// 			"sku": "531qhJK",
// 			"upc": "12345679987",
// 			"pricePerUnit": "12.99",
// 			"quantityOnHand": 28,
// 			"productName": "Gizmo"
// 		},
// 		{
// 			"productId": 3,
// 			"manufacturer": "Big Box Company 3",
// 			"sku": "901qhJK",
// 			"upc": "12345679987",
// 			"pricePerUnit": "12.99",
// 			"quantityOnHand": 28,
// 			"productName": "Gizmo"
// 		}
// // 	]`


// 	err := json.Unmarshal([]byte(productsJSON), &productList)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func getNextID() int {
// 	highestID := -1
// 	for _, p := range productList {
// 		if highestID < p.ProductID {
// 			highestID = p.ProductID
// 		}
// 	}
// 	return highestID + 1
// }

// func findProductByID(productID int) (*Product, int) {
// 	for i, p := range productList {
// 		if productID == p.ProductID {
// 			return &p, i
// 		}
// 	}

// 	return nil, 0
// }



// type fooHandler struct {
// 	Message string
// }

// func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte(f.Message))
// }


// func barHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("bar called"))
// }

// func middlewareHandler(handler http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("before handler; middleware start")
// 		start := time.Now()
// 		handler.ServeHTTP(w, r)
// 		fmt.Printf("middleware finished; %s", time.Since(start))
// 	})
// }

const apiBasePath = "/api"

func main()  {

	product.SetupRoutes(apiBasePath)
	// productListHandler := http.HandlerFunc(productsHandler)
	// productItemHandler := http.HandlerFunc(productHandler)

	// fmt.Println("test")
	// http.Handle("/foo", &fooHandler{Message: "foo called"})
	// http.Handle("/products", middlewareHandler(productListHandler))
	// http.Handle("/products/", middlewareHandler(productItemHandler))
	// http.HandleFunc("/bar", barHandler)
	http.ListenAndServe(":5000", nil)
}