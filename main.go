package main

import (
  "database/sql"
  "net/http"
  //"os"
  "github.com/codegangsta/negroni"
  "github.com/coopernurse/gorp"
  "github.com/heridev/go_autoresponder_api/controllers/subscribers"
  "github.com/heridev/go_autoresponder_api/models/subscriber"
  _ "github.com/lib/pq"
  "github.com/gorilla/mux"
)

var dbmap *gorp.DbMap
var dberr error

func init() {
  dbmap, dberr = InitDb()
  PanicIf(dberr)
}

func main() {
  r := CreateHandler(CreateDbMapHandlerToHTTPHandler(dbmap))
  n := negroni.New()
  n.UseHandler(r)
  n.Run(":" + "3000")
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request)
type DbMapHandlerFunc func(dbmap *gorp.DbMap, w http.ResponseWriter, r *http.Request)
type DbMapHandlerToHTTPHandlerHOF func(f DbMapHandlerFunc) HandlerFunc

func CreateDbMapHandlerToHTTPHandler(dbmap *gorp.DbMap) DbMapHandlerToHTTPHandlerHOF {
  return func(f DbMapHandlerFunc) HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
      f(dbmap, w, r)
    }
  }
}

func CreateHandler(f DbMapHandlerToHTTPHandlerHOF) *mux.Router {
  r := mux.NewRouter()
  r.HandleFunc("/subscribers", f(subscribers.GetAll)).Methods("GET")
  return r
}

func InitDb() (*gorp.DbMap, error) {
  // connect to db using standard Go database/sql API
  db, err := sql.Open("postgres", "dbname=hmail sslmode=disable")
  PanicIf(err)

  // construct a gorp DbMap
  dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
  dbmap.AddTableWithName(subscriber.Subscriber{}, "subscribers").SetKeys(true, "id")

  err = dbmap.CreateTablesIfNotExists()
  if err != nil {
    return nil, err
  }

  return dbmap, nil
}

func PanicIf(err error) {
  if err != nil {
    panic(err)
  }
}

