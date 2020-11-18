package httputil

import (
	"fmt"
	"net/http"
)

func APIHttpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("calling")
	_, _ = w.Write([]byte("I am here."))
}

func apiMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("here start")
	next(rw, r)
	fmt.Println("here end")
}
/*
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", APIHttpHandler)
	n := negroni.New()
	// 中间件的使用必须在 use handler 之前
	n.Use(negroni.HandlerFunc(apiMiddleware))
	n.UseHandler(router)
	_ = http.ListenAndServe(":3001", n)
}
*/
