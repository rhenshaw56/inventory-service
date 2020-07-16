package cors

import "net/http"

// Middleware to set cors headers
func Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Add("Access-Control-Allow-Headerss", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		// fmt.Println("before handler; middleware start")
		// start := time.Now()
		handler.ServeHTTP(w, r)
		// fmt.Printf("middleware finished; %s", time.Since(start))
	})
}