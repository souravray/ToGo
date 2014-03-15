package controllers

import (
	"ToGo/app/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	r "github.com/robfig/revel"
)

var DB gorm.DB

func Init() {

	var err error
	DB, err = gorm.Open("mysql", "root:admin@/togo")

	if err != nil {
		panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
	} else {
		fmt.Println("Successfully connected to DB", DB)
	}
}

type ModelController struct {
	*r.Controller
	Orm gorm.DB
}

func (c *ModelController) Begin() r.Result {
	c.Orm = DB
	return nil
}

func (c *ModelController) AddTables() r.Result {
	c.Orm.CreateTable(models.Task{})
	c.Orm.CreateTable(models.User{})
	return nil
}

func (c *ModelController) EnterData() r.Result {
	user := models.User{
		Name:  "Jitin",
		Tasks: []models.Task{{Description: "Go for a walk.", IsDone: false}, {Description: "Write some Go code!", IsDone: false}},
	}
	c.Orm.Save(&user)
	return nil
}
