package todo

import (
  "encoding/json"
  "errors"
  "fmt"
  //"io"
  "github.com/charmbracelet/lipgloss"
  "github.com/charmbracelet/lipgloss/table"
  "os"
  "time"
)

type item struct {
  Task string
  Done bool
  CreatedAt time.Time
  CompletedAt time.Time
}

type Todos []item

func (t *Todos) Add(task string) {

  todo := item {
    Task: task,
    Done: false,
    CreatedAt: time.Now(),
    CompletedAt: time.Time{},
  }

  *t = append(*t, todo)
}

func (t *Todos) Complete(index int) error {
  ls := *t
  if index <= 0 || index > len(ls) {
    return errors.New("invalid index")
  }

  ls[index-1].CompletedAt = time.Now()
  ls[index-1].Done = true

  return nil
}

func (t *Todos) Delete(index int) error {
  ls := *t
  if index <= 0 || index > len(ls) {
    return errors.New("invalid index")
  }

  *t = append(ls[:index-1], ls[index:]...)

  return nil
}

func (t *Todos) Load(filename string) error {
  file, err := os.ReadFile(filename)
  if err != nil {
    if errors.Is(err, os.ErrNotExist) {
      return nil
    }
    return err
  }

  if len(file) == 0 {
    return err
  }
  err = json.Unmarshal(file, t)
  if err != nil {
    return err
  }

  return nil
}

func (t *Todos) Store(filename string) error {
  data,err := json.Marshal(t)
  if err != nil {
    return err
  }
  return os.WriteFile(filename, data, 0644)
}

func (t *Todos) Print() {

  teaTable := table.New().
      Border(lipgloss.DoubleBorder()).
      Headers("#", "Task", "isDone", "CreatedAt", "CompleteAt").
      Rows()


  for idx, item := range *t {
    idx++
    task := lipgloss.NewStyle().Foreground(lipgloss.Color("4")).Render(item.Task)
    done := lipgloss.NewStyle().Foreground(lipgloss.Color("4")).Render("no")
    if item.Done {
      task = lipgloss.NewStyle().Foreground(lipgloss.Color("2")).Render(fmt.Sprintf("\u2705 %s", item.Task))
      done = lipgloss.NewStyle().Foreground(lipgloss.Color("2")).Render("yes")
    }
      teaTable.Row(
              fmt.Sprintf("%d", idx),
              task,
              done,
              item.CreatedAt.Format(time.RFC822),
              item.CompletedAt.Format(time.RFC822),
          )
      }    

    fmt.Println(teaTable)      
}

func (t *Todos) CountPending() int {
  total := 0
  for _, item := range *t {
    if !item.Done {
      total++
    }
  }
  return total
}
