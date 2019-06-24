package main

import (
  "os"
  "net/http"
  "github.com/gorilla/mux"
  "html/template"
  "log"
  "encoding/json"
  "fmt"
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

// message struct model
type PostMsg struct {
  UserID string `json:"UserID"`
  Msg string `json:"Msg"`
  Timestamp string `json:"Timestamp"`
}
// initialize message struct
var postMsgs []PostMsg

func init()  {
  // intialize the template
  templates = template.Must(template.ParseFiles("template/index.html", "template/message.html"))
  //mock data
  postMsgs = append(postMsgs, PostMsg{
    UserID : "mchan",
    Msg : "hello, how r u?",
    Timestamp : "10:00:00 jun",
  })

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

func main()  {

  // Router
  r := mux.NewRouter()
  fs := http.FileServer(http.Dir("./static/"))

  // api's
  r.HandleFunc(conf.API_MSGS, getMessages).Methods("GET")
  r.HandleFunc(conf.API_MSGS, postMessages).Methods("POST")

  r.HandleFunc(conf.SHUTDOWN, shutdown)
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

// Get all getMessages
func getMessages(w http.ResponseWriter, r *http.Request)  {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(postMsgs)
}

// Post Messages
func postMessages(w http.ResponseWriter, r *http.Request)  {
  w.Header().Set("Content-Type", "application/json")
  var postedMsg PostMsg
  _=json.NewDecoder(r.Body).Decode(&postedMsg)
  postMsgs = append(postMsgs, postedMsg)
  json.NewEncoder(w).Encode(postedMsg)
}

func shutdown(w http.ResponseWriter, r *http.Request) {
  os.Exit(0)
}
