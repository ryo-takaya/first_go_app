package main

import (
	"os"
	"fmt"
)

type Todo string

type TodoRepository interface {
  add(Todo)
  getList() []Todo
}

type MemoryDb struct {
    store []Todo
}

func newMemoryDb () *MemoryDb {
  db := new(MemoryDb)
  return db
}

func ( m *MemoryDb) add (todo Todo) {
    fmt.Println(m.store)
    m.store = append(m.store, todo)
    return
}

func ( m *MemoryDb) getList () []Todo {
    return m.store
}

func getTodoStore () {
    // 環境変数で使い分ける
    environment := os.Getenv("GO_TEST_ENV")
    switch environment {
        case "local":
            newMemoryDb()
    }
}