package email_lists

import (
  "net/http"
  "github.com/jinzhu/gorm"
  "github.com/heridev/go_autoresponder_api/utils"
  "github.com/heridev/go_autoresponder_api/models/email_list"
  "gopkg.in/unrolled/render.v1"
)


func IndexHandler(w http.ResponseWriter, req *http.Request) {
  // connect to db using standard Go database/sql API
  db, err := gorm.Open("postgres", "dbname=hmail sslmode=disable")
  utils.PanicIf(err)

  r := render.New(render.Options{})

  var lists []email_list.EmailList
  db.Find(&lists)
  if lists == nil {
    r.JSON(w, http.StatusOK, nil) // If we have no subscribers, just return an empty array, instead of null.
  } else {
    r.JSON(w, http.StatusOK, &lists)
  }
}

//func GetAll(dbgorm *gorm.DB, w http.ResponseWriter, r *http.Request) {
  //subscribers, err := subscriber.GetAll(dbmap)
  //if err != nil {
    //fmt.Printf("ERROR: %s\n", err.Error())
    //w.WriteHeader(http.StatusInternalServerError)
    //return
  //}

  //response, err := json.Marshal(subscribers)
  //if err != nil {
    //fmt.Printf("ERROR: %s\n", err.Error())
    //w.WriteHeader(http.StatusInternalServerError)
    //return
  //}

  //utils.WriteOkResponse(w, http.StatusOK, response)
//}

//func WriteInternalError(err error, w http.ResponseWriter) {
  //utils.WriteError(err, w, http.StatusInternalServerError)
//}

//func WriteNotFound(err error, w http.ResponseWriter) {
  //utils.WriteError(err, w, http.StatusNotFound)
//}

