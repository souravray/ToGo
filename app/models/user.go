package models

type User struct {
    Id           int64
    Name         string  `sql:"size:255"`
    Tasks            []Task         // Embedded structs
}
