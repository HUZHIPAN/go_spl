package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	httpProt := os.Args[1]
	
    server :=&http.Server{
        Addr: fmt.Sprintf(":%d", httpProt),
    }
    http.HandleFunc("/", requestHandle)
    server.ListenAndServe()
}

func requestHandle(w http.ResponseWriter,r *http.Request)  {

    fmt.Fprintf(w,"url：%s", r.URL)
}