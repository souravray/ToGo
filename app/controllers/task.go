package controllers
import  (
    "github.com/robfig/revel"
    "ToGo/app/models"
    "ToGo/app/routes"
    // "github.com/jinzhu/gorm"
)

type Task struct {
    ModelController
}

func (c Task) Index() revel.Result {
	user:=models.User{}
	tasks := []models.Task{}
	if username, ok := c.Session["user"]; ok {
		if err:=c.Orm.Where("name = ?", username).First(&user).Error;  err==nil {
			c.Orm.Model(&user).Related(&tasks)
			user.Tasks = tasks
			return c.Render(user, tasks)
		}
	} 
   return c.Redirect(routes.App.Index())
}

func (c Task) Add(task string) revel.Result {
	user:=models.User{}
	if username, ok := c.Session["user"]; ok {
		if err:=c.Orm.Where("name = ?", username).First(&user).Error;  err==nil {
			newTask := models.Task{UserId: user.Id, Description: task}
			if error:= c.Orm.Save(&newTask).Error; error!= nil{
				c.Flash.Error ("Cannot add task")
			}
			return c.Redirect(routes.Task.Index())
		}
	} 
    return c.Redirect(routes.App.Index())
}

func (c Task) Done(taskId int64) revel.Result {
    user:=models.User{}
	if username, ok := c.Session["user"]; ok {
		if err:=c.Orm.Where("name = ?", username).First(&user).Error;  err==nil {
			task := models.Task{}
			if error:= c.Orm.Where("Id = ?", taskId).First(&task).Error; error!= nil{
				c.Flash.Error ("Cannot find task")
			} else {
				task.IsDone=true
			}
			if error:= c.Orm.Save(&task).Error; error!= nil{
				c.Flash.Error ("Cannot modify task")
			} 
			return c.Redirect(routes.Task.Index())
		}
	} 
    return c.Redirect(routes.App.Index())
}


func (c Task) Undone(taskId int64) revel.Result {
     user:=models.User{}
	if username, ok := c.Session["user"]; ok {
		if err:=c.Orm.Where("name = ?", username).First(&user).Error;  err==nil {
			task := models.Task{}
			if error:= c.Orm.Where("Id = ?", taskId).First(&task).Error; error!= nil{
				c.Flash.Error ("Cannot find task")
			} else {
				task.IsDone=false
			}
			if error:= c.Orm.Save(&task).Error; error!= nil{
				c.Flash.Error ("Cannot modify task")
			} 
			return c.Redirect(routes.Task.Index())
		}
	} 
    return c.Redirect(routes.App.Index())
}