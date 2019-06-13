package main

import (
  "os"
  "net/http"
  "github.com/gorilla/mux"
  "html/template"
)

var templates *template.Template
func init()  {
  templates = template.Must(template.ParseGlob("template/index.html"))
}

func main()  {
  r := mux.NewRouter()
  fs := http.FileServer(http.Dir("./static/"))

  r.HandleFunc("/shutdown", shutdown)
  r.HandleFunc("/", indexHandler)
  r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
  http.Handle("/", r)
  http.ListenAndServe(":5000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  templates.ExecuteTemplate(w, "index.html", nil)
}

func shutdown(w http.ResponseWriter, r *http.Request) {
  os.Exit(0)
}
