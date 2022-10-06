package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"github.com/golang/glog"
)

func addSystemVariableToHeader(w http.ResponseWriter, key string) {
	value := os.Getenv(key)

	if len(value) == 0 {
		glog.Warning("The system variable ", key, " is NULL")
	}

	w.Header().Set(key, value)
}

func addRequestHeaderToResponseHeader(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		w.Header().Set(k, v[0])
	}
}

func logHttpClient(r *http.Request) {

	index := strings.LastIndex(r.RemoteAddr,":")

	remote_addr := r.RemoteAddr[:index] 

	fmt.Println(remote_addr," access http server")
	glog.Info(remote_addr, " access http server")
}

func rootResponseHandler(w http.ResponseWriter, r *http.Request) {

	addSystemVariableToHeader(w, "VERSION")

	addRequestHeaderToResponseHeader(w, r)

	io.WriteString(w, "hello world!\n")

	logHttpClient(r)
	//defer logHttpClient(r)
}

func healthzResponseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "the server is working well\n")
}

func main() {
	flag.Parse()
	defer glog.Flush()

	glog.Info("http server will work")

	mux := http.NewServeMux()
	mux.HandleFunc("/", rootResponseHandler)
	mux.HandleFunc("/healthz", healthzResponseHandler)

	err := http.ListenAndServe(":12345", mux)

	if err != nil {
		glog.Error("[http server error]---", err)
	}

}
