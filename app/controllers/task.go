package controllers
import  (
    "github.com/robfig/revel"
    "myapp/app/models"
    "myapp/app/routes"
    "github.com/jinzhu/gorm"

type Task struct {
    ModelController
}

func (c Task) Index() revel.Result {
    return c.Render()
}

func (c Task) Add(task string) revel.Result {
    return c.Redirect(routes.Task.Index())

func (c Task) Done(taskId int64) revel.Result {
    return c.Redirect(routes.Task.Index())
}

func (c Task) Undone(taskId int64) revel.Result {
    return c.Redirect(routes.Task.Index())
}