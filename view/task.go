package view

import (
	"github.com/labstack/echo"
	"strconv"

	"net/http"

	"github.com/sin_tanaka/echo_todo_crud/models"
	"github.com/sin_tanaka/echo_todo_crud/db"
)

//----------
// HandlerID int
//----------
// packageで呼び出す時の関数名はCamelCaseにすること

func CreateTask(c echo.Context) error {
	conn := db.Opendb()
	defer conn.Close()

	conn.Create(&models.Task{
		Task: "do hogehoge.",
	})

	return c.JSON(http.StatusCreated, "{status: ok}")
}

func GetTasks(c echo.Context) error {
	conn := db.Opendb()
	defer conn.Close()

	tasks := []models.Task{}
	recs := conn.Find(&tasks)
	return c.JSON(http.StatusOK, recs) // c.JSONはmapをそのまま返せる
}

func UpdateTask(c echo.Context) error {
	conn := db.Opendb()
	defer conn.Close()

	id, _ := strconv.Atoi(c.Param("id"))
	task := models.Task{}
	conn.First(&task, id).Update("task", "testtest")
	return c.JSON(http.StatusOK, task)
}

func DeleteTask(c echo.Context) error {
	conn := db.Opendb()
	defer conn.Close()

	id, _ := strconv.Atoi(c.Param("id"))
	task := models.Task{}
	conn.Delete(&task, id)

	return c.JSON(http.StatusOK, id)
}
