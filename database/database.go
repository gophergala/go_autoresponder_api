package database

import (
  "github.com/jinzhu/gorm"
  "github.com/heridev/go_autoresponder_api/utils"
)

var Db gorm.DB

func Init() {
  // connect to db using standard Go database/sql API
  db, err := gorm.Open("postgres", "dbname=hmail sslmode=disable")
  utils.PanicIf(err)
  Db = db
}

//func Jojo() gorm.DB{
  //// connect to db using standard Go database/sql API
  //db, err := gorm.Open("postgres", "dbname=hmail sslmode=disable")
  //utils.PanicIf(err)

  //return db
//}
