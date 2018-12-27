package main

import (
    "fmt"
    "log"
    "net/http"
    "html/template"
    "wiki"
)

func handler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request){
    title := r.URL.Path[len("/view/"):]
    p, _ := wiki.LoadPage(title)
    t, _ := template.ParseFiles("templates/view.html")
    t.Execute(w, p) 
}

func main(){
    http.HandleFunc("/view/", viewHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}