package utils

import (
  "database/sql"
  "github.com/coopernurse/gorp"
  "fmt"
  "net/http"
)

func CreateTableWithID(tableName string, entity interface{}, incrementable bool) (*gorp.DbMap, error) {
  db, err := sql.Open("postgres", "dbname=hmail sslmode=disable")
  if err != nil {
    return nil, err
  }

  // construct a gorp DbMap
  dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

  // add a table, setting the table name to 'posts' and
  // specifying that the Id property is an auto incrementing PK
  dbmap.AddTableWithName(entity, tableName).SetKeys(incrementable, "id")

  // create the table. in a production system you'd generally
  // use a migration tool, or create the tables via scripts
  err = dbmap.CreateTablesIfNotExists()
  if err != nil {
    return nil, err
  }

  return dbmap, nil
}

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
