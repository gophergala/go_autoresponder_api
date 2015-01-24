package subscriber

import (
  "github.com/coopernurse/gorp"
)

type Subscriber struct {
  ID         int64     `db:"id" json:"id"`
  Created    string    `db:"created" json:"created_at"`
  Updated    string    `db:"updated" json:"updated_at"`
  Name       string    `db:"name" json:"name"`
  Email      string    `db:"email" json:"email"`
}

type OrderVariantData struct {
  ID       string `json:"variant_id"`
  Quantity int    `json:"quantity"`
}

func GetAll(dbmap *gorp.DbMap) ([]Subscriber, error) {
  var subscribers []Subscriber
  _, err := dbmap.Select(&subscribers, "select * from subscribers order by id")
  if err != nil {
    return nil, err
  }
  return subscribers, nil
}
