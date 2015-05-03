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
    r.Get("/", GetUsers)
    r.Post("/", CreateUser)
    r.Get("/:insta_user_id/", GetUser)
  })


  defer GetDB().Close()

  m.Run()
}
