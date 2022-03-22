package main

import (
	"Centralized_transaction/auto"
	"Centralized_transaction/config"
	"Centralized_transaction/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()
	auto.Load()
	fmt.Printf("\n\tListening [::]:%d\n", config.PORT)
	listen(config.PORT)
}
func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
