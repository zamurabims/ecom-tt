package storage

import (
	"ecom-tt/internal/models"
	"sync"
)

type TodoStorageStruct struct {
	mu     sync.Mutex
	todos  map[int]models.Todo
	nextId int
}

func NewStorage() *TodoStorageStruct {
	return &TodoStorageStruct{
		todos:  make(map[int]models.Todo),
		nextId: 1,
	}
}
func (s *TodoStorageStruct) GetAlTasks() []models.Todo {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := make([]models.Todo, 0, len(s.todos))
	for _, todo := range s.todos {
		result = append(result, todo)
	}
	return result
}

func (s *TodoStorageStruct) GetByTaskID(id int) (models.Todo, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo, ok := s.todos[id]
	return todo, ok
}

func (s *TodoStorageStruct) CreateTask(todo models.Todo) models.Todo {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo.ID = s.nextId
	s.nextId++

	s.todos[todo.ID] = todo
	return todo
}

func (s *TodoStorageStruct) UpdateTask(id int, todo models.Todo) (models.Todo, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.todos[id]; !ok {
		return models.Todo{}, false
	}

	todo.ID = id
	s.todos[id] = todo
	return todo, true
}

func (s *TodoStorageStruct) DeleteTask(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.todos[id]; !ok {
		return false
	}

	delete(s.todos, id)
	return true
}
