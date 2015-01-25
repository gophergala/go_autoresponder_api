package utils

import (
  "fmt"
  "net/http"
)

func WriteOkResponse(w http.ResponseWriter, httpStatus int, body []byte) {
  w.WriteHeader(httpStatus)
  w.Header().Set("Content-Type", "application/json")
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Write(body)
}

func WriteError(err error, w http.ResponseWriter, httpStatus int) {
  fmt.Printf("ERROR: %s\n", err.Error())
  w.WriteHeader(httpStatus)
}

func PanicIf(err error) {
  if err != nil {
    panic(err)
  }
}
