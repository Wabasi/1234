package main

import (
    "fmt"
    "net/http"
    "strings"
    "log"
)

/**
 * @api {get} /hi Request User information
 * @apiName Get
 *
 */
func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()  // parse arguments, you have to call this by yourself
    fmt.Println(r.Form)  // print form information in server side
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Yeet some Wheat") // send data to client side

}
/**
 * @api {get} /swag Request User information
 * @apiName Get
 *
 */
func sayswag(w http.ResponseWriter, r *http.Request) {

    fmt.Fprintf(w, "What's good family")

}

func main() {
    http.HandleFunc("/hi", sayhelloName,)
    http.HandleFunc("/swag", sayswag,)// set router
    fmt.Println("Let's get some bread")
    err := http.ListenAndServe(":80", nil) // set listen port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
