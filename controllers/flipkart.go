package controllers
import (
    "net/http"
    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"
    "github.com/amardaxini/martini_goquery/models"   

    // "strconv"
)

type FlipkartResponse struct {
    Name      int    `json:"string"`
    Description string `json:"description"`

}
func FlipkartItems(params martini.Params, render render.Render, request *http.Request) {
  query := request.URL.Query()
  url:= query.Get("url")

  flipkart_item := models.FlipkartParser(url)
  render.JSON(200, flipkart_item)
}
