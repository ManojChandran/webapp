package api

import (
  "os"
  "net/http"
  "encoding/json"
)

// message struct model
type PostMsg struct {
  UserID string `json:"UserID"`
  Msg string `json:"Msg"`
  Timestamp string `json:"Timestamp"`
}
// initialize message struct
var postMsgs []PostMsg

func init(){
  //mock data
  postMsgs = append(postMsgs, PostMsg{
    UserID : "mchan",
    Msg : "hello, how r u?",
    Timestamp : "10:00:00 jun",
  })
}

// Get all getMessages
func GetMessages(w http.ResponseWriter, r *http.Request)  {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(postMsgs)
}

// Post Messages
func PostMessages(w http.ResponseWriter, r *http.Request)  {
  w.Header().Set("Content-Type", "application/json")
  var postedMsg PostMsg
  _=json.NewDecoder(r.Body).Decode(&postedMsg)
  postMsgs = append(postMsgs, postedMsg)
  json.NewEncoder(w).Encode(postedMsg)
}

func Shutdown(w http.ResponseWriter, r *http.Request) {
  os.Exit(0)
}
