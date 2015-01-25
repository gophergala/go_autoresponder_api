package database

import (
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
  "github.com/heridev/go_autoresponder_api/utils"
  "github.com/heridev/go_autoresponder_api/models/subscriber"
  "github.com/heridev/go_autoresponder_api/models/email_list"
)

var DbInstance gorm.DB

func InitDb() {
  // connect to db using standard Go database/sql API
  var err error
  DbInstance, err = gorm.Open("postgres", "dbname=hmail sslmode=disable")
  utils.PanicIf(err)

  DbInstance.AutoMigrate(&subscriber.Subscriber{},
                 &email_list.EmailList{})
}
