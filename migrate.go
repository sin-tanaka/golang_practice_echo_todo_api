package main

import (
	"github.com/sin_tanaka/echo_todo_crud/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sin_tanaka/echo_todo_crud/db"
)


func main() {
	conn := db.Opendb()
	defer conn.Close()

	conn.CreateTable(&models.Task{})
	conn.AutoMigrate(&models.Task{})
}