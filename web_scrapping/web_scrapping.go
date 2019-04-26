package web_scrapping

import (
  "fmt"
  "log"
  "net/http"
  "net/url"

  "github.com/PuerkitoBio/goquery"
)

// This will get called for each HTML element found
func process_element(index int, element *goquery.Selection) {
    band := element.Text()
    fmt.Println(band)
}

func Print_each_element(doc *goquery.Document, tag_name string){
  doc.Find(tag_name).Each(process_element)
}


func Search_google(query string) *goquery.Document{

  base_search_url := "https://www.google.com/search?q="
  get_url := base_search_url + url.QueryEscape(query)

  res, err := http.Get(get_url)
  if err != nil {
    log.Fatal(err)
  }

  defer res.Body.Close()

  if res.StatusCode != 200 {
    log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
  }

  // Load the HTML document
  doc, err := goquery.NewDocumentFromReader(res.Body)
  if err != nil {
    log.Fatal(err)
  }

  return doc

}
