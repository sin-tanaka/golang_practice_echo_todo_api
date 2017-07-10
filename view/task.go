package view

import (
	"github.com/labstack/echo"
	"strconv"

	"net/http"

	"github.com/sin_tanaka/echo_todo_crud/models"
	// . "github.com/sin_tanaka/echo_todo_crud/models" // いわゆるstar importなので可読性は低い
)

var (
	// 本当はORMなどで、DB操作したいが
	// 一旦modelの配列を用意する
	tasks = map[int]*models.Task{}
	seq   = 1
)

//----------
// Handlers
//----------
// packageで呼び出す時の関数名はCamelCaseにすること

func CreateTask(c echo.Context) error {
	u := &models.Task{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	tasks[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func GetTasks(c echo.Context) error {
	//id, _ := strconv.Atoi(c.Param("id"))
	//return c.JSON(http.StatusOK, tasks[id])
	return c.JSON(http.StatusOK, tasks) // c.JSONはmapをそのまま返せる
}

func UpdateTask(c echo.Context) error {
	u := new(models.Task)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	tasks[id].Task = u.Task
	return c.JSON(http.StatusOK, tasks[id])
}

func DeleteTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(tasks, id)
	return c.NoContent(http.StatusNoContent)
}
