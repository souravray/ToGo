package controllers

import (
    "fmt"
    "github.com/jinzhu/gorm"
     _ "github.com/go-sql-driver/mysql"
    r "github.com/robfig/revel"
    "myapp/app/models"
)

var DB gorm.DB

func init() {

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
        Name: "jitin",
        Tasks: []models.Task{{Description: "Go for a walk.", IsDone: false}, {Description: "Write some go code", IsDone: false}},
    }
    c.Orm.Save(&user)
    return nil
}
