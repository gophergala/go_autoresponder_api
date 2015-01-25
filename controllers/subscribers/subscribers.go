package subscribers

import (
  "net/http"
  "github.com/jinzhu/gorm"
  "github.com/heridev/go_autoresponder_api/models/subscriber"
  "github.com/heridev/go_autoresponder_api/utils"
  "gopkg.in/unrolled/render.v1"
)

func IndexHandler(w http.ResponseWriter, req *http.Request) {
  // connect to db using standard Go database/sql API
  db, err := gorm.Open("postgres", "dbname=hmail sslmode=disable")
  utils.PanicIf(err)

  defer db.Close()

  r := render.New(render.Options{})

  var subscribers []subscriber.Subscriber
  db.Find(&subscribers)
  if subscribers == nil {
    r.JSON(w, http.StatusOK, nil) // If we have no subscribers, just return an empty array, instead of null.
  } else {
    r.JSON(w, http.StatusOK, &subscribers)
  }
}
