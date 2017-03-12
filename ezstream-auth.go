package main

import (
    "fmt"
    "net/http"
    "strings"
)

func on_publish(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    fmt.Printf("Publish Started: %s\n", r.FormValue("name"))
    secretkey := r.FormValue("swfurl")[strings.Index(r.FormValue("swfurl"), "?")+1:]
    fmt.Printf("secret: %s\n", secretkey)
    if(r.FormValue("name") == "scalesstreamkey" && secretkey == "123456"){
        w.WriteHeader(http.StatusOK)
    } else {
        w.WriteHeader(http.StatusUnauthorized)
    }
}

func on_publish_done(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    fmt.Printf("Publish Stopped: %s\n", r.FormValue("name"))
}

func on_play(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    fmt.Printf("asdf")
    //fmt.Printf("Stream Playing: %s\n", r.FormValue("name"))
}

func main() {
    http.HandleFunc("/on_publish", on_publish)
    http.HandleFunc("/on_publish_done", on_publish_done)
    http.HandleFunc("/on_play", on_play)
    http.ListenAndServe(":8080", nil)
}
