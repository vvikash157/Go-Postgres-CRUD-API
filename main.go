package main

import (
	"fmt"
	"go-postgress/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("starting server on PORT 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
