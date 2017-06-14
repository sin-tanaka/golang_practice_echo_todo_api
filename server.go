package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// https://echo.labstack.com/cookbook/crudの写経
// モデルはtodoへ変更

type (
	todo struct {
		ID   int    `json:"id"`
		Task string `json:"task"`
	}
)

var (
	// 本当はORMなどで、DBのsessionを操作したいが
	// 一旦todo構造体の配列を用意する
	// migrationできるORMがよい
	tasks = map[int]*todo{}
	seq   = 1
)

//----------
// Handlers
//----------

func createUser(c echo.Context) error {
	u := &todo{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	tasks[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, tasks[id])
}

func updateUser(c echo.Context) error {
	u := new(todo)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	tasks[id].Task = u.Task
	return c.JSON(http.StatusOK, tasks[id])
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(tasks, id)
	return c.NoContent(http.StatusNoContent)
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/tasks", createUser)
	e.GET("/tasks/:id", getUser)
	e.PUT("/tasks/:id", updateUser)
	e.DELETE("/tasks/:id", deleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}