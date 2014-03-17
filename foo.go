package main

import "encoding/json"
import "fmt"

type Trinity int
const (
  Rock Trinity = iota
  Paper
  Scissors
)

type leader_card struct {
  Id int
  Name string
  Chrisma int
  Hp int
  Ap int
  Trin Trinity 
}

type support_card struct {
  id int
  name string
  cost int
  Trin Trinity
}

type hand struct {
  leader leader_card
  support [3]support_card
}

func main() {
  card := leader_card{1, "Some Leader", 10, 10, 10, Rock}
  js, err := json.Marshal(card)
  fmt.Println(err)
  fmt.Println(string(js))
}
