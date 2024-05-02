package main

type MockStore struct {}

func (s *MockStore) CreateUser() error {
	return nil
}
func (s *MockStore) CreateTask(t *Task) (*Task, error) {
	return &Task{}, nil
}

func (s *MockStore) GetTask(id string) (*Task, error) {
	return &Task{}, nil
}