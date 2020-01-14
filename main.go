package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/korovkin/gotils"
)

func registerJSON() {
	// a JSON endpoint:
	http.HandleFunc("/json",
		func(w http.ResponseWriter, r *http.Request) {
			log.Println("=> request", r.RequestURI)

			res := map[string]interface{}{}
			resJSON := ""
			res["ok"] = true
			res["time"] = time.Now()
			resJSON = gotils.ToJSONString(res)

			defer func() {
				log.Println("=> response: json:", r.RequestURI, resJSON)

				// content-type: json
				w.Header().Set("Content-Type", "application/json")

				// write response
				_, err := io.WriteString(w, resJSON+"\n")
				gotils.CheckFatal(err)
			}()
		})
}

func registerXML() {
	// an XML Endpoint:
	http.HandleFunc("/xml",
		func(w http.ResponseWriter, r *http.Request) {
			log.Println("=> request", r.RequestURI)

			res := struct {
				XMLName xml.Name `xml:"Response"`
				OK      bool     `xml:"ok"`
				Time    int64    `xml:"time"`
			}{}
			res.OK = true
			res.Time = time.Now().Unix()
			resXML := ""
			resXML = gotils.ToXMLString(&res)

			defer func() {
				log.Println("=> response: xml:", r.RequestURI, resXML)

				// content-type: json
				w.Header().Set("Content-Type", "application/xhtml+xml")

				// write response
				_, err := io.WriteString(w, resXML+"\n")
				gotils.CheckFatal(err)
			}()
		})
}

func registerPANIC() {
	// handle a panic gracefully:
	http.HandleFunc("/panic",
		func(w http.ResponseWriter, r *http.Request) {

			defer func() {
				err := recover()
				if err == nil {
					return
				}
				log.Println("RECOVER: ERR:", err)

				io.WriteString(
					w,
					fmt.Sprintf("PANIC RECOVER: %s\n", err.(error).Error()),
				)
			}()

			log.Println("=> request", r.RequestURI)
			var p *int
			log.Println("=> this should panic", *p)
		},
	)
}

func registerERROR() {
	// HTTP Error:
	http.HandleFunc("/error",
		func(w http.ResponseWriter, r *http.Request) {
			log.Println("=> request", r.RequestURI)
			http.Error(w,
				fmt.Sprintf("nope: %d", http.StatusNotFound),
				http.StatusNotFound)
		},
	)
}

func main() {
	host := flag.String("host", "0.0.0.0", "server address to bind to")
	port := flag.Int("port", 9001, "server port to run on")

	address := fmt.Sprintf("%s:%d", *host, *port)
	log.Println(fmt.Sprintf(" => Running on http://%s", address))

	// a plain txt endpoint:
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			log.Println("=> request", r.RequestURI)
			log.Println("=> response", r.RequestURI)
			_, err := io.WriteString(w, "ok\n")
			gotils.CheckFatal(err)
		},
	)

	log.Println("ENDPOINTS: start")
	registerJSON()
	registerXML()
	registerPANIC()
	registerERROR()
	log.Println("ENDPOINTS: end")

	// run the server forever:
	// err := http.ListenAndServe(address, nil)
	// gotils.CheckFatal(err)

	err := http.ListenAndServeTLS(address,
		"certs/localhost.crt",
		"certs/localhost.key",
		nil)
	gotils.CheckFatal(err)
}
