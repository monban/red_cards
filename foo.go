package main

import "fmt"

type Trinity int
const (
  Rock Trinity = iota
  Paper
  Scissors
)

type leader_card struct {
  id int
  name string
  chrisma int
  hp int
  ap int
  trinity Trinity
}

func main() {
  fmt.Printf("hello world!\n")
}
