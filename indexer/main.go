package main

import (
	"fmt"
	"log"
	"os"

	"github.com/HusskyAngel/ZincIndexer/pkg/stack"
)

func main() {

  if len(os.Args)<2 {
    log.Fatal("Non folder for load specified")
  }

  s:=stack.Stack{}
  s.Push(1)
  s.Push(2)
  s.Pop()
  fmt.Println(s.GetTop())
  
}
