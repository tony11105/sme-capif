// route.go
package route

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"oransc.org/nonrtric/capifcore/internal/publishservice"

	"oransc.org/nonrtric/capifcore/internal/eventservice"
)

func ProxyHandler(ps *publishservice.PublishService, es *eventservice.EventService, uri string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("r.URL.Path: ", r.URL.Path)
		remote, err := url.Parse(uri)

		if err != nil {
			log.Printf("Error parsing URI: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		log.Println("Request Host: ", r.Host)
		if r.Host != "backend" {
			token := r.URL.Query().Get("token")
			apiId := r.URL.Query().Get("apiId")
			if token == "" {
				http.Error(w, "Missing token parameter", http.StatusBadRequest)
				return
			}
			if apiId == "" {
				http.Error(w, "Missing apiId parameter", http.StatusBadRequest)
				return
			}

			validateRequest := eventservice.ValidateAPIRequest(token, apiId)
			if validateRequest == false {
				http.Error(w, "ValidateAPIRequest Error", http.StatusInternalServerError)
				return
			}

			validateAPIendpoint := publishservice.ValidateApiEndpoint(ps, r.URL.Path)
			if validateAPIendpoint == false {
				http.Error(w, "ValidateApiEndpoint", http.StatusInternalServerError)
				return
			}
		}

		proxy := httputil.NewSingleHostReverseProxy(remote)
		if remote.Scheme == "https" {
			proxy.Transport = &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
		}
		proxy.ServeHTTP(w, r)
	}
}
