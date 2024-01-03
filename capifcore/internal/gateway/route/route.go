// route.go
package route

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ProxyHandler(uri string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// log.Println("r.URL.Path: ", r.URL.Path)
		remote, err := url.Parse(uri)
		log.Println("print remote: ", remote)
		if err != nil {
			log.Printf("Error parsing URI: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
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
