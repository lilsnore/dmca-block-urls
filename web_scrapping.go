package main

import (
  "os"
  "fmt"
  "log"
  "net/http"

  "github.com/PuerkitoBio/goquery"
)

// This will get called for each HTML element found
func process_element(index int, element *goquery.Selection) {
    // See if the href attribute exists on the element
    band := element.Text()
    fmt.Println(band)
}


func main(){

  query_string := os.Args[1]
  base_search_url := "https://www.google.com/search?q="
  get_url := base_search_url + query_string

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

  doc.Find("cite").Each(process_element)

}
