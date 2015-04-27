package main

// "server_api"
import (
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"

  "structures"
)


func main() {
  db, err := sql.Open("sqlite3", "dbfile.db")
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
    // rows, err := db.Query("SELECT * FROM user")
    // checkErr(err)

    // for rows.Next() {
    //   var uid int
    //   var username string
    //   var department string
    //   var created string
    //   err = rows.Scan(&uid, &username, &department, &created)
    //   checkErr(err)
    //   fmt.Println(uid)
    //   fmt.Println(username)
    //   fmt.Println(department)
    //   fmt.Println(created)
    // }

    fmt.Println(params['insta_user_id'])

    r.JSON(200)
  })
  r.Get(`/testadd`, func(r render.Render, params martini.Params) {//, db *mgo.Database
    s1 := structures.User{FirstName: "Write", SecondName: "presentation", InstaUserID: 1234567890}

    stmt, err := db.Prepare("INSERT INTO user(first_name, second_name, insta_user_id) values(?,?,?)")
    checkErr(err)
    res, err := stmt.Exec(s1.FirstName, s1.SecondName, s1.InstaUserID)
    checkErr(err)

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
