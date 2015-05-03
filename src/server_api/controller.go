package server_api

import (
  "fmt"

  _ "github.com/mattn/go-sqlite3"
  "github.com/kisielk/sqlstruct"

  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"

  "structures"
)



func GetUser(r render.Render, params martini.Params) {//, db *mgo.Database
  q := fmt.Sprintf("SELECT %s FROM user WHERE insta_user_id = %s", sqlstruct.Columns(structures.User{}), params["insta_user_id"])
  rows, err := GetDB().Query(q)
  checkErr(err)

  var u structures.User
  for rows.Next() {
    err = sqlstruct.Scan(&u, rows)
    fmt.Println(u, err)
  }
  err = rows.Err()
  fmt.Println(err)

  fmt.Println(params["insta_user_id"])

  r.JSON(200, u);
}
