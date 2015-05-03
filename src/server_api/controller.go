package server_api

import (
  "fmt"
  "strconv"

  _ "github.com/mattn/go-sqlite3"
  "github.com/kisielk/sqlstruct"

  "net/http"
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

  r.JSON(200, u)
}

func CreateUser(resp render.Render, req *http.Request) {
  u := getPostUser(req)

  //fmt.Println(sqlstruct.Columns(structures.User{}))
  //=> age, first_name, insta_user_id, insta_user_name, last_name, study_faculty, study_year

  q := fmt.Sprintf("INSERT INTO user(%s) values(%d, '%s', %d, '%s', '%s', '%s', %d)", sqlstruct.Columns(structures.User{}),
                   u.Age, u.FirstName, u.InstaUserID, u.InstaUserName, u.LastName, u.StudyFaculty, u.StudyYear)
  _, err := db.Exec(q)
  checkErr(err)

  resp.JSON(201, u)
}


func getPostUser(r *http.Request) structures.User {
  id, err := strconv.ParseInt(r.PostFormValue("id"), 10, 0)
  checkErr(err)
  study_year, err := strconv.ParseInt(r.PostFormValue("study_year"), 10, 0)
  checkErr(err)
  age, err := strconv.ParseInt(r.PostFormValue("age"), 10, 0)
  checkErr(err)


  u := structures.User{
    InstaUserID: int(id),
    InstaUserName: r.PostFormValue("username"),
    FirstName: r.PostFormValue("first_name"),
    LastName: r.PostFormValue("last_name"),
    StudyYear: int(study_year),
    StudyFaculty: r.PostFormValue("study_faculty"),
    Age: int(age)}

  return u
}
