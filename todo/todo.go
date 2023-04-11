package todo

import "sync"

var ToDoTask = NewTodos()

type Todo struct {
	Todo string `json:"todo"`
}

type Todos struct {
	sync.RWMutex
	Map map[string][]Todo
}

func NewTodos() *Todos {
	return &Todos{Map: make(map[string][]Todo)}
}

func (t *Todos) Add(username string, todo Todo) {
	t.Lock()
	defer t.Unlock()

	if _, ok := t.Map[username]; !ok {
		t.Map[username] = []Todo{}
	}
	t.Map[username] = append(t.Map[username], todo)
}

func (t *Todos) Get(username string) []Todo {
	t.RLock()
	defer t.Unlock()
	if todos, ok := t.Map[username]; ok {
		return todos
	}
	return []Todo{}
}

func (t *Todos) Delete(username string, index int) {
	t.Lock()
	defer t.Unlock()
	if todos, ok := t.Map[username]; ok && index >= 0 && index < len(todos) {
		t.Map[username] = append(todos[:index], todos[index+1:]...)
	}
}
