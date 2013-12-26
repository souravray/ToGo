package controllers
import  (
    "github.com/robfig/revel"
    "ToGo/app/models"
    "ToGo/app/routes"
    "github.com/jinzhu/gorm"
)

type Auth struct {
    ModelController
}

func (c Auth) Login(name string) revel.Result {
	user:=models.User{Name:name}
	if err:=c.Orm.Where("name = ?", name).First(&user).Error; 
		err!=nil && err!= gorm.RecordNotFound {
		c.Flash.Error("Database unavailable")
	} else {
		if c.Orm.NewRecord(user) {
			c.Orm.Save(&user)
		} 
		c.Session["user"] = name
		c.Flash.Success("Welcome, " + name)
		return c.Redirect(routes.Task.Index())
	}
    return  c.Redirect(routes.App.Index())
}

func (c Auth) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(routes.App.Index())
}