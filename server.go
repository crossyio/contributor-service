package main

import (
	"encoding/json"
	"fmt"
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
	path := req.PathParams["organization"] + "/" + req.PathParams["project"] + "/" + req.PathParams["packager"] + "/" + req.PathParams["version"] + "/" + req.PathParams["platform"] + "/" + req.PathParams["arch"] + "/" + req.PathParams["user"] + "/" + req.PathParams["file"]
	presignedURL := GetS3Presigned("contribute-crossy-io", path, 30)
	if err := json.NewEncoder(rw).Encode(presignedURL); err != nil {
		panic(err)
	}
}

func (c *Context) Healthcheck(rw web.ResponseWriter, req *web.Request) {
	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, "{\"online\": true}")
}

func main() {
	router := web.New(Context{}). // Create your router
					Middleware(web.LoggerMiddleware).     // Use some included middleware
					Middleware(web.ShowErrorsMiddleware). // ...
		// Middleware((*Context).SetHelloCount). // Your own middleware!
		Get("/healthcheck", (*Context).Healthcheck).
		Post("/api/v1/:organization/:project/:packager/:version/:platform/:arch/:user/:file", (*Context).GeneratePresigned) // Add a route
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	http.ListenAndServe("0.0.0.0:"+port, router) // Start the server!
}
