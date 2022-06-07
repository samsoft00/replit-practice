package main

import (
	"fmt"
  "group_title"
)

func main() {
  searchTitle := NewGroupTitle()
  result :=searchTitle.Search("car")
  
  fmt.Println(result)
}
