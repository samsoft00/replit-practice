package group_title

import (
	"fmt"
	"sort"
	"strconv"
)

type Title struct {
  titles []string
}

func (t *Title) groupTitles() [][]string {
  var output [][]string
  if len(t.titles) == 0 {
    return output
  }

  res := make(map[string][]string)

  for _, s := range t.titles {
    count := make([]int, 26)
    
    for _, c := range s {
      index := c - 'a'
      count[index]++
    }

    key := ""
    for i := 0; i < 26; i++ {
      key += "#"
      key += strconv.Itoa(count[i])
    }

    res[key] = append(res[key], s)
  }

  for _, v := range res {
    output = append(output, v)
  }

  return output  
}

func (t *Title) Search(query string) {
  output := t.groupTitles()

  for _, o := range output {
    sort.Strings(o) // ["duel", "dule", "duel"]
    i := sort.Search(len(o), func (i int) bool { return o[i] >= query })
    
    if i < len(o) && o[i] == query {
      fmt.Println(o)
    }
    
  }
}

func NewGroupTitle() *Title {
  return &Title{
    titles: []string{"duel", "dule", "speed", "spede", "duel", "cars"},
  }
}