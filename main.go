package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
	"unsafe"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

var addr = flag.String("listen-address", ":9100", "The address to listen on for HTTP requests.")

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

	go allocateMemory()

	flag.Parse()
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func allocateMemory() {

	var buffer [100 * 1024 * 1024]string
	fmt.Printf("The size of the buffer is: %d bytes\n", unsafe.Sizeof(buffer))

	for e, _ := range buffer {
		time.Sleep(100 * time.Nanosecond)
		buffer[e] = "string"
	}

}
