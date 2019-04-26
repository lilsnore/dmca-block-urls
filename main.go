package main

import (
  "os"
  "./web_scrapping"
)

func main() {

  query := os.Args[1]
  doc := web_scrapping.Search_google(query)
  web_scrapping.Print_each_element(doc, "cite")

}
