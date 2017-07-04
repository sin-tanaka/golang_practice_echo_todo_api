package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
	"fmt"
	"strings"
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

func getIndex(c echo.Context) error {
	header := c.Request().Header //Headerの実装は、map[string][]string
	html := "<div>Hello World. 下記はあなたのRequestHeaderです。<br/>"

	//range map[]でkey, valueを捜査出来る
	for key, value := range header{
		// Joinで配列をカンマ区切りのstringへ変換
		value := strings.Join(value, ", ")

		// Goで書式化文字列を扱いたいときはSprintf()
		// %vで良さげに型推定される
		html += fmt.Sprintf("%s: %v<br/>", key, value)
	}
	html += "</div>"
	return c.HTML(http.StatusOK, html)
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

	// ルートへアクセスした時
	e.GET("/", getIndex)

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT"))) // Heroku用のポート設定
	//e.Logger.Fatal(e.Start(":8000")) // DEBUG: ローカル用のポート設定
}
