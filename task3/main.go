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
	
	var portVar int
	var port string

	flag.IntVar(&portVar,"port",80,"http server listen port")
	flag.Parse()
	if portVar < 0 ||  portVar > 65535{
		glog.Error("http server port ",portVar," is illegal")
		return 
	}
	port = fmt.Sprintf(":%d",portVar)

	defer glog.Flush()

	glog.Info("http server will work")

	mux := http.NewServeMux()
	mux.HandleFunc("/", rootResponseHandler)
	mux.HandleFunc("/healthz", healthzResponseHandler)

	err := http.ListenAndServe(port, mux)

	if err != nil {
		glog.Error("[http server error]---", err)
	}

}
