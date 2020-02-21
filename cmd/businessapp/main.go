package main

import (
	"fmt"
	"log/syslog"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	lSyslog "github.com/sirupsen/logrus/hooks/syslog"
)

func main() {

	log := logrus.New()
	hook, err := lSyslog.NewSyslogHook("tcp", "127.0.0.1:8086", syslog.LOG_INFO, "")

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
		// go logAsync("got some business to do %s", "yeah", log)
		log.Infof("got some business to do %s", "yeah")
		fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	})
}

func logAsync(message, arg string, log *logrus.Logger) {
	log.Infof(message, arg)
}
