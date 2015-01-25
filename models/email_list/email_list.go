package email_list

import (
  "time"
)

type EmailList struct {
  Title        string     `db:"title" json:"title"`
  Content      string     `sql:"size:0" db:"content" json:"content"`
  CreatedAt  time.Time    `db:"created" json:"created_at"`
  UpdatedAt  time.Time    `db:"updated" json:"updated_at"`
}
