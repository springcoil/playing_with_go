package main

import (
    "encoding/json"	
    "fmt"
    "net/http"
)    

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, %s!", r.URL.Path[1:])
}

func main() {
     http.HandleFunc("/", handler)
     http.HandleFunc("/about/", about)
     http.ListenAndServe(":8081", nil)
}     

type Message struct {
    Text string
}

func about (w http.ResponseWriter, r *http.Request) {
	
	m := Message{"Welcome to the Peadar Coyle API, build v 0.1, 20/9/16"}
    b, err := json.Marshal(m)

    if err != nil {
	panic(err)
    }

    w.Write(b)
}    
