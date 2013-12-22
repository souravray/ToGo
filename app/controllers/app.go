package controllers
import  (
    // "fmt"
    // "myapp/app/db"
    "github.com/robfig/revel"
    // "myapp/app/models"
)

type App struct {
    ModelController
}

func (c App) Index() revel.Result {
    return c.Render()
}
