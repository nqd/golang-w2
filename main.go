/*
 * Secret Server
 *
 * This is an API of a secret service. You can save your secret by using the API. You can restrict the access of a secret after the certen number of views or after a certen period of time.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"log"
	"net/http"

	"github.com/nqd/golang-w2/handlers"
)

func main() {
	log.Printf("Server started")

	router := handlers.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
