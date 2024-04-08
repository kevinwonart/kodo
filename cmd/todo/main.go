package main

import (
  "fmt"
  "flag"
  "os"
  "github.com/kevinwonart/todo"
)

const (
  todoFile = ".todos.json"
)


func main() {

  add := flag.Bool(name:"add", value:false, usage: "add a new todo")

  flag.Parse()

  todos := &todo.Todos{}

  if err := todos.Load(todoFile); err != nil {
    fmt.Fprintln(os.Stderr, err.Error())
    os.Exit(code:1)
  }
  switch{
  case *add:
    todos.Add(task: "Sample todo")
    err := todos.Store(todoFile);
    if err != nil {
      fmt.Fprintln(os.Stderr, err.Error())
      os.Exit(code:1)
    }
  default:
    fmt.Fprintln(os.Stdout, a...:"invalid command")
    os.Exit(code:0)
  }
}

