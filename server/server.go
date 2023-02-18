package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"marunk20/rs/hello"
	"net/http"
)

func homePage(httpWriter http.ResponseWriter, httpRequest *http.Request) {
	fmt.Fprint(httpWriter, "Hello. See hhttps://github.com/arundotin/go-resourceServer for more details")
}

func RegisterRoutesInServer() *mux.Router {
	router := mux.NewRouter()
	router.Use(httpHeaderMiddleware)

	jwtMiddleware := ADFSJWTTokenValidationMiddleware()

	router.HandleFunc("/", homePage)
	router.Handle("/api/v1/hello", jwtMiddleware.Handler(
		http.HandlerFunc(hello.SayHello))).Methods("GET")

	return router

}

func httpHeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
