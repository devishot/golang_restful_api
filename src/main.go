package main

// "server_api"
import (
  "fmt"
  "log"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"

  "structures"
)



/*
  sqlStmt := `
  CREATE TABLE user (first_name text, last_name text, study_year int, study_faculty text, age int, insta_user_id int not null primary key, insta_user_name text);
  delete from user;
  `
  _, err = db.Exec(sqlStmt)
  if err != nil {
    log.Printf("%q: %s\n", err, sqlStmt)
    return
  }
*/

func checkErr(err error) {
  if err != nil {
    log.Fatal(err)
  }
}


func main() {
  db, err := sql.Open("sqlite3", "file:dbfile.db")
  checkErr(err)
  defer db.Close()


  m := martini.Classic()

  // Setup middleware
  /*
  m.Use(martini.Recovery())
  m.Use(martini.Logger())
  m.Use(auth.Basic(AuthToken, ""))
  */
  m.Use(render.Renderer(render.Options{
        IndentJSON: true,
    }))

  // Setup routes
  r := martini.NewRouter()
  r.Get(`/users/:insta_user_id`, func(r render.Render, params martini.Params) {//, db *mgo.Database
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
  })
  r.Get(`/testadd`, func(r render.Render, params martini.Params) {//, db *mgo.Database
    s1 := structures.User{FirstName: "Write", LastName: "presentation", InstaUserID: 1234567890}

    stmt, err := db.Prepare("INSERT INTO user(first_name, last_name, insta_user_id) values(?,?,?)")
    checkErr(err)
    res, err := stmt.Exec(s1.FirstName, s1.LastName, s1.InstaUserID)
    checkErr(err)

    fmt.Println(res)

    r.JSON(200, s1)
  })

  // Add the router action
  m.Action(r.Handle)

  m.Run()
}


/*import "github.com/go-martini/martini"

func main() {
  m := martini.Classic()
  m.Get("/", func() string {
    return "Hello world!"
  })
  m.Run()
}
*/
