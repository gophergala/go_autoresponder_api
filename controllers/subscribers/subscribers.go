package subscribers

import (
  "encoding/json"
  "fmt"
  "net/http"
  "github.com/coopernurse/gorp"
  "github.com/heridev/go_autoresponder_api/models/subscriber"
  "github.com/heridev/go_autoresponder_api/utils"
)

func GetAll(dbmap *gorp.DbMap, w http.ResponseWriter, r *http.Request) {
  subscribers, err := subscriber.GetAll(dbmap)
  if err != nil {
    fmt.Printf("ERROR: %s\n", err.Error())
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  response, err := json.Marshal(subscribers)
  if err != nil {
    fmt.Printf("ERROR: %s\n", err.Error())
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  utils.WriteOkResponse(w, http.StatusOK, response)
}

func WriteInternalError(err error, w http.ResponseWriter) {
  utils.WriteError(err, w, http.StatusInternalServerError)
}

func WriteNotFound(err error, w http.ResponseWriter) {
  utils.WriteError(err, w, http.StatusNotFound)
}

