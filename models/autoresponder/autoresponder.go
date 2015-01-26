package autoresponder

import (
  "github.com/heridev/go_autoresponder_api/models/email_list"
)

type Autoresponder struct {
  Id           int64
  Title        string `db:"title" json:"title"`
  Description  string `db:"description" json:"description"`
  Lists        []email_list.EmailList `json:"lists"`
}
