package main

import (
	"crypto/tls"
	"log"
)

func main() {
	certificate, err := tls.X509KeyPair([]byte(cert), []byte(key))
	if err != nil {
		log.Fatalln("Can't create key/cert")
	}
	pool := x509.NewCertPool()
	if !pool.AppendCertsFromPEM([]byte(ca)) {
		log.Fatalln("Can't append ca")
	}
	conn, err := tls.Dial("tcp", "", &tls.Config{
		Certificates: []tls.Certificate{certificate},
		RootCAs:      pool,
	})
	if err != nil {
		log.Fatalln("Can't dial apiserver")
	}
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/namespaces/default/pods/swagger-editor-847cf74cd6-2g4mf/exec?container=swagger-editor&command=/bin/sh&stdin=true&stderr=true&stdout=true&tty=true", nil)
	req.Header.Set("Connection", "Upgrade")
	req.Header.Set("Upgrade", "xxxx")
	err = req.Write(conn)
	if err != nil {
		log.Fatalln("Can't send req")
	}

	resp, err := http.ReadResponse(bufio.NewReader(conn), req)
	if err != nil {
		log.Fatalln("Can't read resp")
	}

	log.Printf("%d", resp.StatusCode)
	log.Printf("%+v", resp.Header)
	data, _ := ioutil.ReadAll(resp.Body)
	log.Printf("%s", data)

	req, _ = http.NewRequest(http.MethodGet, "/containerLogs/default/swagger-editor-847cf74cd6-2g4mf/swagger-editor", nil)
	err = req.Write(conn)
	if err != nil {
		log.Fatalln("Can't send req")
	}
	resp, err = http.ReadResponse(bufio.NewReader(conn), req)
	if err != nil {
		log.Fatalln("Can't read resp")
	}

	log.Printf("%d", resp.StatusCode)
	log.Printf("%+v", resp.Header)
	data, _ = ioutil.ReadAll(resp.Body)
	log.Printf("%s", data)

}
