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
    r.Post("/", CreateUser)
    // r.Put("/update/:id", UpdateBook)
    // r.Delete("/delete/:id", DeleteBook)
  })


  defer GetDB().Close()

  m.Run()
}
