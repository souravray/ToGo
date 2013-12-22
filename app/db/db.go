package db

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
)

var DB gorm.DB

func Init() {
    fmt.Println("----------------- Calling DB init ----------------")
    var err error
    DB, err = gorm.Open("mysql", "root:admin@/togo")
    
    if err != nil {     
        panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
    } else {
        fmt.Println("<<<<<<<<<<<<<<<< Success >>>>>>>>>>>", DB)
    }
}