package main

import (
  "os"
  "net/http"
  "github.com/gorilla/mux"
  "html/template"
  "log"
  "encoding/json"
  "fmt"
  api "github.com/ManojChandran/webapp/api"
)

var templates *template.Template

type configuration struct {
  PORT            string
  ROOT            string
  SHUTDOWN        string
  STATIC          string
  API_MSGS        string
}

var conf configuration

type Page struct{
  Title string
}

func init()  {
  // intialize the template
  templates = template.Must(template.ParseFiles("template/index.html", "template/message.html"))
  readConfFile()
}

func main()  {

  // Router
  r := mux.NewRouter()
  fs := http.FileServer(http.Dir("./static/"))

  // api's
  r.HandleFunc(conf.API_MSGS, api.GetMessages).Methods("GET")
  r.HandleFunc(conf.API_MSGS, api.PostMessages).Methods("POST")

  r.HandleFunc(conf.SHUTDOWN, api.Shutdown)
  r.HandleFunc(conf.ROOT, indexHandler)
  r.PathPrefix(conf.STATIC).Handler(http.StripPrefix(conf.STATIC, fs))
  http.Handle(conf.ROOT, r)
  log.Fatal(http.ListenAndServe(conf.PORT, nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  data := &Page{
    Title : "TITLE",
  }
  templates.ExecuteTemplate(w, "index.html", data)
}

func readConfFile()  {
    // reading api paths
    file,_ := os.Open("conf.json")
    defer file.Close()
    decoder := json.NewDecoder(file)
    //  conf := configuration{}
    err := decoder.Decode(&conf)
    if err != nil {
      fmt.Println("path file not found", err)
    }
}
