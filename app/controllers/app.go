package controllers
import  (
    "github.com/robfig/revel"
    "ToGo/app/routes"
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
