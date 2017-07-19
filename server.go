package main

import (
	"html/template"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sin_tanaka/echo_todo_crud/view"
	"io"
	"os"
	"strings"
)

// https://echo.labstack.com/guide/templates を写経
type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func getIndex(c echo.Context) error {
	header := c.Request().Header //Headerの実装は、map[string][]string
	requests := map[string]string{}

	for key, value := range header { //range map[]でkey, valueを捜査出来る
		requests[key] = strings.Join(value, ", ")
	}
	// {name: {hoge: fuga}} のような構造で渡して、.nameで{hoge: fuga}にアクセスできる
	data := map[string]map[string]string{
		"requests": requests,
	}

	// {{define index}}しているテンプレートファイルに自動的にマッピングする
	return c.Render(http.StatusOK, "index", data)
}

func main() {
	e := echo.New()
	t := &Template{
		// テンプレートファイルパス
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = t

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/tasks", view.CreateTask)
	e.GET("/tasks", view.GetTasks)
	e.PUT("/tasks/:id", view.UpdateTask)
	e.DELETE("/tasks/:id", view.DeleteTask)

	e.GET("/", getIndex)

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT"))) // Heroku用のポート設定
}
