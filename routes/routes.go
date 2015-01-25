package routes

import(
  "github.com/heridev/go_autoresponder_api/controllers/subscribers"
  "github.com/heridev/go_autoresponder_api/controllers/email_lists"
  "github.com/gorilla/mux"
)

func Create() *mux.Router {
  r := mux.NewRouter()
  r.HandleFunc("/subscribers",  subscribers.IndexHandler ).Methods("GET")
  r.HandleFunc("/lists",  email_lists.IndexHandler).Methods("GET")
  return r
}
