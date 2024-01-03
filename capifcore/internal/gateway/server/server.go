// server.go
package main

import (
	"log"
	"net/http"
	"strings"

	"oransc.org/nonrtric/capifcore/internal/gateway/conf"
	gw "oransc.org/nonrtric/capifcore/internal/gateway/route"
)

func StartServer() {
	config, err := conf.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalln("Failed to load config:", err)
	}

	for _, route := range config.Gateway.Routes {
		for _, predicate := range route.Predicates {
			if strings.HasPrefix(predicate, "Path=") {
				pathPrefix := strings.TrimPrefix(predicate, "Path=")
				http.HandleFunc(pathPrefix, gw.ProxyHandler(route.URI, pathPrefix))
			}
		}
	}

	log.Println("Starting server on :30092")
	if err := http.ListenAndServe(":30092", nil); err != nil {
		log.Fatalln("ListenAndServe failed:", err)
	}
}
