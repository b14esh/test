package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func main() {

    req, err := http.Get("https://b14esh.com")
    if err != nil {
        log.Fatal(err)
    }
    defer req.Body.Close()
    data, err := ioutil.ReadAll(req.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(data))

}
