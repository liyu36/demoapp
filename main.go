package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	readyzStatusCode = http.StatusOK
	readyzStatus     = "Ready"
	livezStatusCode  = http.StatusOK
	livezStatus      = "Live"
	version          = "v1.0"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
}

func readyz(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if strings.ToLower(r.Header.Get("status")) != "ready" {
			readyzStatus = "NotReady"
			readyzStatusCode = http.StatusServiceUnavailable
		} else {
			readyzStatus = "ready"
			readyzStatusCode = http.StatusOK
		}
		logrus.WithFields(logrus.Fields{"client": r.RemoteAddr, "path": r.URL.Path, "code": readyzStatusCode}).Info()
	} else if r.Method == http.MethodGet {
		w.WriteHeader(readyzStatusCode)
		fmt.Fprintln(w, readyzStatus)
	}
}

func livez(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if strings.ToLower(r.Header.Get("status")) != "live" {
			livezStatus = "NotLive"
			livezStatusCode = http.StatusServiceUnavailable
		} else {
			livezStatus = "live"
			livezStatusCode = http.StatusOK
		}
		logrus.WithFields(logrus.Fields{"client": r.RemoteAddr, "path": r.URL.Path, "code": livezStatusCode}).Info()
	} else if r.Method == http.MethodGet {
		w.WriteHeader(livezStatusCode)
		fmt.Fprintln(w, livezStatus)
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hostname: %s Pod: %s ClientIP: %s Version: %s\n", os.Getenv("HOSTNAME"), os.Getenv("POD_NAME"), r.RemoteAddr, version)
	logrus.WithFields(logrus.Fields{"client": r.RemoteAddr, "path": r.URL.Path}).Info()
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/readyz", readyz)
	http.HandleFunc("/livez", livez)
	port := ""
	if port = os.Getenv("PORT"); port == "" {
		port = ":8080"
	} else {
		port = ":" + port
	}
	if err := http.ListenAndServe(port, nil); err != nil {
		logrus.Fatal(err)
	}
}
