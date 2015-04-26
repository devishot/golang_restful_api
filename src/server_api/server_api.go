package server_api

import (
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
  "structures"
)



type Server *martini.ClassicMartini


func RunServer() Server {//session *DatabaseSession
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
  r.Get(`/albums`, func(r render.Render) {//, db *mgo.Database
      al := structures.User{FirstName: "Write presentation", InstaUserID: 1234567890}
      r.JSON(200, al)
    })
/*  r.Get(`/albums/:id`, GetAlbum)
  r.Post(`/albums`, AddAlbum)
  r.Put(`/albums/:id`, UpdateAlbum)
  r.Delete(`/albums/:id`, DeleteAlbum)
*/

  // Add the router action
  m.Action(r.Handle)

  m.Run()

  return m
}
