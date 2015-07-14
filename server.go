package main

import (
    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"
    "./controllers"  
    // "strconv"
)






func main() {

    m := martini.Classic()
    m.Use(render.Renderer())
    m.Use(martini.Recovery())

  
  
    m.Group("/api/v1",func(router martini.Router) {
      router.Get("/flipkart", controllers.FlipkartItems)
    })
   

    m.Run()
    // go func() {
    //   if err := http.ListenAndServe(":8000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    //     http.Error(w, "https scheme is required", http.StatusBadRequest)
    //   }))
    //   err != nil {
    //     log.Fatal(err)
    //   }()

}