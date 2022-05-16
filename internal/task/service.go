package task

import "context"

type Service interface {
	Get(ctx context.Context, id string) (Task, error)
	// TODO: create offset, limit
	GetAll(ctx context.Context) ([]Task, error)
	Count(ctx context.Context) int
	Create(ctx context.Context, rq CreateTaskRequest) (Task, error)
	Update(ctx context.Context, id string, rq UpdateTaskRequest) (Task, error)
	Delete(ctx context.Context, id string)
}

type CreateTaskRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (r CreateTaskRequest) Validate() error {
	// implement validation
	return nil
}

type UpdateTaskRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (r UpdateTaskRequest) Validate() error {
	// implement validation
	return nil
}

type service struct {
	repo Repository
}

func NewService(repo repository) Service {
	return service{repo: repo}
}

func (s service) Get(ctx context.Context, id string) (Task, error) {
	panic("implement me")
}

func (s service) GetAll(ctx context.Context) ([]Task, error) {
	panic("implement me")
}

func (s service) Count(ctx context.Context) int {
	panic("implement me")
}

func (s service) Create(ctx context.Context, rq CreateTaskRequest) (Task, error) {
	panic("implement me")
}

func (s service) Update(ctx context.Context, id string, rq UpdateTaskRequest) (Task, error) {
	panic("implement me")
}

func (s service) Delete(ctx context.Context, id string) {
	panic("implement me")
}

