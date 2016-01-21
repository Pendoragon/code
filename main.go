package main

import (
	"net/http"
	"fmt"
	"github.com/Pendoragon/go-misc/sni"
)

func main() {
	httpsServer := &http.Server{
		Addr:      ":8001",
	}

	var certs []sni.Certificates
	certs = append(certs, sni.Certificates{
		CertFile: "path-to-cert",
		KeyFile: "path-to-key",
	})
	certs = append(certs, sni.Certificates{
		CertFile: "path-to-cert",
		KeyFile: "path-to-key",
	})

	fmt.Println("Listening on port 8001...")
	err := sni.ListenAndServeTLSSNI(httpsServer, certs)
	if err != nil {
		fmt.Printf("err: %s", err)
	}
}
