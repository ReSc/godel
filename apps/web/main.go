package main

import (
	"encoding/xml"
	"github.com/ReSc/godel/core/rest"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	xml.NewEncoder(os.Stdout).Encode(bootstrapNodes())
	http.ListenAndServe(":3000", router())
}

func authenticate() {
	ratelimit()
	// get credentials
	// if creds ok {
	//   set credential context
	// } else {
	//	 deny request and redirect to login
	// }
}

func authorize() {
	ratelimit()
	// get credentials from context
	// if no credentials deny request
	// check url acl,accept or deny request
}

func ratelimit() {

}

func pipeline() http.Handler {
	return nil
}

func router() *mux.Router {
	r := mux.NewRouter()
	serveFiles(r)
	serveApi(r)
	return r
}

func serveFiles(r *mux.Router) {
	dir := http.Dir("./static")
	fileServer := http.FileServer(dir)
	r.PathPrefix("/app/").Handler(http.StripPrefix("/app/", fileServer))
}

func serveApi(r *mux.Router) {
	r = r.PathPrefix("/api/v1/").Subrouter()
	rest.Register(r, newNodeResource(), "id:[0-9]+")
}
