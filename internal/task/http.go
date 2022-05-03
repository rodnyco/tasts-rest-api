package task

import (
	routing "github.com/go-ozzo/ozzo-routing"
)

func RegisterHandlers(r *routing.RouteGroup)  {
	resource := resource{}
	r.Get("/tasks/<id>", resource.get)
	r.Get("/tasks", resource.getAll)
	r.Post("/tasks", resource.create)
	r.Put("/tasks/<id>", resource.update)
	r.Delete("/tasks/<id>", resource.delete)
}

type resource struct {
	//TODO: add service
	//TODO: add logger
}

func (r resource) get(c *routing.Context) error {
	return c.Write("Get task handler")
}

func (r resource) getAll(c *routing.Context) error {
	return c.Write("Gel All tasks handler")
}

func (r resource) create(c *routing.Context) error {
	return c.Write("Create task handler")
}

func (r resource) update(c *routing.Context) error {
	return c.Write("Update task handler")
}

func (r resource) delete(c *routing.Context) error {
	return c.Write("Delete task handler")
}
