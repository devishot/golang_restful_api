package main

// "server_api"
import (
  "fmt"
  "log"

  "database/sql"
  _ "github.com/mattn/go-sqlite3"
  "github.com/kisielk/sqlstruct"

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


var db *sql.DB
func getDB() (*sql.DB) {
  fmt.Println(db)
  return db
}

func openDB() {
  fmt.Println(db)
  _db, err := sql.Open("sqlite3", "file:dbfile.db")
  checkErr(err)
  db = _db
}


func GetUser(r render.Render, params martini.Params) {//, db *mgo.Database
  q := fmt.Sprintf("SELECT %s FROM user WHERE insta_user_id = %s", sqlstruct.Columns(structures.User{}), params["insta_user_id"])
  rows, err := getDB().Query(q)
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


func main() {
  openDB()


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


  m.Group("/users", func(r martini.Router) {
    r.Get("/:insta_user_id", GetUser)
    // r.Post("/", NewBook)
    // r.Put("/update/:id", UpdateBook)
    // r.Delete("/delete/:id", DeleteBook)
  })

  // r := martini.NewRouter()
  // r.Get(`/testadd`, func(r render.Render, params martini.Params) {//, db *mgo.Database
  //   s1 := structures.User{FirstName: "Write", LastName: "presentation", InstaUserID: 1234567890}

  //   stmt, err := getDB().Prepare("INSERT INTO user(first_name, last_name, insta_user_id) values(?,?,?)")
  //   checkErr(err)
  //   res, err := stmt.Exec(s1.FirstName, s1.LastName, s1.InstaUserID)
  //   checkErr(err)

  //   fmt.Println(res)

  //   r.JSON(200, s1)
  // })
  // // Add the router action
  // m.Action(r.Handle)

  defer getDB().Close()

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
