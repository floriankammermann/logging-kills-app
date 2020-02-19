package main

import (
	"fmt"
	"log/syslog"
	"net/http"
	"time"
	"unsafe"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	lSyslog "github.com/sirupsen/logrus/hooks/syslog"
)

func main() {

	log := logrus.New()
	//hook, err := lSyslog.NewSyslogHook("tcp", "localhost:8086", syslog.LOG_INFO, "")
	hook, err := lSyslog.NewSyslogHook("tcp", "localhost:8086", syslog.LOG_INFO, "")

	if err == nil {
		log.Hooks.Add(hook)
	}

	finish := make(chan bool)

	server9100 := http.NewServeMux()
	server9100.Handle("/metrics", promhttp.Handler())

	server8081 := http.NewServeMux()
	server8081.Handle("/", businessHandler(log))

	go func() {
		http.ListenAndServe(":9100", server9100)
	}()

	go func() {
		http.ListenAndServe(":8081", server8081)
	}()

	<-finish
}

func businessHandler(log *logrus.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof("got some business to do", "yeah")
		fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	})
}

func allocateMemory() {

	var buffer [100 * 1024 * 1024]string
	fmt.Printf("The size of the buffer is: %d bytes\n", unsafe.Sizeof(buffer))

	for e, _ := range buffer {
		time.Sleep(100 * time.Nanosecond)
		buffer[e] = "string"
	}

}
