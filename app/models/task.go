package models

type Task struct {
    Id           int64
    UserId       int64
    Description  string  `sql:"size:255"`
    IsDone       bool
}