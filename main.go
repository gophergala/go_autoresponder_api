package main

import (
  "net/http"
  //"os"
  "github.com/codegangsta/negroni"
  //"github.com/heridev/go_autoresponder_api/controllers/subscribers"
  "github.com/heridev/go_autoresponder_api/models/subscriber"
  "github.com/gorilla/mux"
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
  "gopkg.in/unrolled/render.v1"
)

var dberr error

type DBHandler struct {
  db *gorm.DB
  r  *render.Render
}

func InitDb() gorm.DB {
  // connect to db using standard Go database/sql API
  db, err := gorm.Open("postgres", "dbname=hmail sslmode=disable")
  PanicIf(err)

  db.CreateTable(subscriber.Subscriber{})
  return db
}

func main() {
  db := InitDb()
  defer db.Close()

  // using the gopkg.in/unrolled/render.v1 library
  r := render.New(render.Options{})
  h := DBHandler{db: &db, r: r}

  router := CreateRoutes(&h)

  n := negroni.New()
  n.UseHandler(router)
  n.Run(":" + "3000")
}

func CreateRoutes(h *DBHandler) *mux.Router {
  r := mux.NewRouter()
  r.HandleFunc("/subscribers",  h.subscribersIndexHandler).Methods("GET")
  return r
}

func (h *DBHandler) subscribersIndexHandler(w http.ResponseWriter, req *http.Request) {
  var subscribers []subscriber.Subscriber
  h.db.Find(&subscribers)
  if subscribers == nil {
    h.r.JSON(w, http.StatusOK, nil) // If we have no subscribers, just return an empty array, instead of null.
  } else {
    h.r.JSON(w, http.StatusOK, &subscribers)
  }
}

func PanicIf(err error) {
  if err != nil {
    panic(err)
  }
}
