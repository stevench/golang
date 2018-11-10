/*
created on 2018-11-10 晚 by steven
*/

package main

import (
   "fmt"
   "net/http"
)

func init() {
    fmt.Println("web service listen at 10000!")
}

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World!\n")
}

func main() {
    http.HandleFunc("/", hello)
    err := http.ListenAndServe(":10000", nil)
    if err != nil {
        fmt.Println(err)
	return
    }
}
