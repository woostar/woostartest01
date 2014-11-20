package main
 
import (
    "net/http"
    "text/template"
    "html"
)
 
var mainTemplate, _ = template.ParseFiles("topics.html")
 
func mainPage(w http.ResponseWriter, r *http.Request) {
    mainTemplate.Execute(w, nil)
}
 
func main() {
    http.HandleFunc("/", mainPage)
    http.ListenAndServe(":8080", nil)
}