package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gocraft/web"
)

type Context struct {
	ResponseJSON interface{}
}

func (c *Context) GeneratePresigned(rw web.ResponseWriter, req *web.Request) {
	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rw.WriteHeader(http.StatusOK)
	presignedURL := GetS3Presigned("contribute-crossy-io", "octoblu/meshblu/v1.0.0/iamruinous/meshblu-blah.tar.gz", 30)
	if err := json.NewEncoder(rw).Encode(presignedURL); err != nil {
		panic(err)
	}
}

func main() {
	router := web.New(Context{}). // Create your router
					Middleware(web.LoggerMiddleware).     // Use some included middleware
					Middleware(web.ShowErrorsMiddleware). // ...
		// Middleware((*Context).SetHelloCount). // Your own middleware!
		Get("/", (*Context).GeneratePresigned) // Add a route
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	http.ListenAndServe("0.0.0.0:"+port, router) // Start the server!
}
