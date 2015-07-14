package models
import (
  // "fmt"
  "log"
  "github.com/PuerkitoBio/goquery"
  "net/url"
  "strings"
  "regexp"

)
type FlipkartResponse struct {
    Name      string    `json:"name"`
    Description string `json:"description"`
    Price string `json:"price"`
    Discontinued bool `json:"discontinued"`
    Specifications (map[string]string) `json:"specifications"`
    Sku string `json:"sku"`
    Availability  bool `json:"availability"`

   
}
func FlipkartParser(flipkarturl string) FlipkartResponse{
  // doc, err := goquery.NewDocument("http://www.flipkart.com/samsung-galaxy-s3-neo/p/itme3q7q93m77cg6?pid=MOBDUZSYRMG7R3U9") 
  u, err := url.Parse(flipkarturl)
  if err != nil {
    log.Fatal(err)
  }
  q := u.Query()
  
  doc, err := goquery.NewDocument(flipkarturl) 
  if err != nil {
    log.Fatal(err)
  }
  // m := make(map[string]string)
  product_details := doc.Find("div#fk-mainbody-id")
  product_info := product_details.Find("div.product-details")

  title := doc.Find("h1.title").Text()
  price := product_details.Find("div.prices span.selling-price").Text()
  description := product_details.Find("div.description div.description-text").Text()
  availability:= false
  availability_text := product_info.Find("span.subtitle").Text()

  if len(availability_text) != 0{
    availability = true
  }
  // seller_avail_text := product_info.Find("div.see-all-sellers-wrap").Text()
  // seller_avail := false

  // if len(seller_avail_text) != 0{
  //   seller_avail = true
  // }
  discontinued := false
  discontinued_text := product_info.Find("div.see-all-sellers-wrap").Text()
  if len(discontinued_text) !=0 {
    discontinued = true
  }


  sku := q["pid"][0] 

  specification_nodes := product_details.Find("div.productSpecs")
  spec_table := specification_nodes.Find("table.specTable")
  reg, err := regexp.Compile("\\s")
  specifications := make(map[string]string)
  spec_table.Each(func(i int, specs *goquery.Selection) {
    specs.Find("tr").Each(func(i int,spec *goquery.Selection){
        key := spec.Find(".specsKey").Text()
        value := spec.Find(".specsValue").Text()
        if len(key) > 0 {
          specification_key := strings.ToLower(strings.TrimSpace(key))
          specification_key = reg.ReplaceAllString(specification_key, "-")
          specifications[specification_key] = strings.TrimSpace(value)
        } 
      })
      
    })
  // for spec_key := range specifications {
  //     fmt.Println(spec_key,specifications[spec_key])
  //  }



  response := FlipkartResponse{
        Name:      title,
        Description:description,
        Price: price,
        Discontinued: discontinued,
        Specifications: specifications,
        Sku: sku,
        Availability: availability,

    }
  return response
}