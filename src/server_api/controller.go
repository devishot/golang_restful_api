package server_api

import (
  "fmt"
  "strconv"
  "encoding/json"

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

  if u.InstaUserID != 0 {
    r.JSON(200, u)
    fmt.Println("here")
  } else {
    fmt.Println("now there")
    errorResp, err := json.Marshal(map[string]string{"error": "user not found"})
    checkErr(err)
    r.JSON(404, errorResp)
  }
}

func GetUsers(r render.Render, params martini.Params) {//, db *mgo.Database
  q := fmt.Sprintf("SELECT %s FROM user", sqlstruct.Columns(structures.User{}))
  rows, err := GetDB().Query(q)
  checkErr(err)

  var users []structures.User
  var u structures.User
  for rows.Next() {
    err = sqlstruct.Scan(&u, rows)
    checkErr(err)
    fmt.Println(u, err)

    users = append(users, u)
  }
  err = rows.Err()
  checkErr(err)

  fmt.Println(users)

  r.JSON(200, users)
}

func CreateUser(resp render.Render, req *http.Request) {
  u := getPostUser(req)

  //fmt.Println(sqlstruct.Columns(structures.User{}))
  //=> full_name, insta_user_id, insta_user_name, study_faculty, study_year

  q := fmt.Sprintf("INSERT INTO user(%s) values('%s', %d, '%s', '%s', %d)", sqlstruct.Columns(structures.User{}),
                   u.FullName, u.InstaUserID, u.InstaUserName, u.StudyFaculty, u.StudyYear)
  _, err := db.Exec(q)
  checkErr(err)

  resp.JSON(201, u)
}


func getPostUser(r *http.Request) structures.User {
  id, err := strconv.ParseInt(r.PostFormValue("insta_user_id"), 10, 0)
  checkErr(err)
  study_year, err := strconv.ParseInt(r.PostFormValue("study_year"), 10, 0)
  checkErr(err)


  u := structures.User{
    InstaUserID: int(id),
    InstaUserName: r.PostFormValue("insta_user_name"),
    FullName: r.PostFormValue("full_name"),
    StudyYear: int(study_year),
    StudyFaculty: r.PostFormValue("study_faculty")}

  return u
}
