package server_api

import (
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
)



func RunServer() {
  OpenDB()

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

  defer GetDB().Close()

  m.Run()
}
