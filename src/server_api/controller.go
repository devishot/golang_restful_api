package server_api

import (
  "fmt"

  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"

  "database/sql"
  "database/sql/driver"
  _ "github.com/mattn/go-sqlite3"
)



func GetUser(r render.Render, params martini.Params) {//, db *mgo.Database
  rows, err := db.Query("SELECT * FROM user WHERE insta_user_id = ? ", params["insta_user_id"])
  checkErr(err)

  for rows.Next() {
    //u := structures.User{}
    col := rows.Columns()
    vals := make([]interface{}, col)
    rows.Scan(vals...)
    fmt.Println(vals)
    // err = rows.Scan(&u., &username, &department, &created)
    // checkErr(err)
    // fmt.Println(uid)
    // fmt.Println(username)
    // fmt.Println(department)
    // fmt.Println(created)
  }

  fmt.Println(params["insta_user_id"])

  //r.JSON(200, );
}
