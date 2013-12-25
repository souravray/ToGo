package controllers
import  (
    // "fmt"
    // "myapp/app/db"
    "github.com/robfig/revel"
    "myapp/app/routes"
)

type App struct {
    ModelController
}

func (c App) Index() revel.Result {
    if _, ok := c.Session["user"]; ok {
        return c.Redirect(routes.Task.Index())
    } else {
        return c.Render()
    }
}
