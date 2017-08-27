package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

func indexHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	log.Println("starting")
    
    certManager := autocert.Manager{
        Prompt:     autocert.AcceptTOS,
        HostPolicy: autocert.HostWhitelist("shankarvellal.com"),
        Cache:      autocert.DirCache("certs"), //folder for storing certificates
    }

    server := &http.Server{
        Addr: ":443",
        TLSConfig: &tls.Config{
            GetCertificate: certManager.GetCertificate,
        },
    }

	http.HandleFunc("/", indexHandler)
	if err := server.ListenAndServeTLS("", ""); err != nil {
        log.Fatalf(err.Error())
    }
}
