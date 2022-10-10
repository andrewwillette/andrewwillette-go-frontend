package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	startServer()
}

func startServer() {
	e := echo.New()
	e.GET("/", handleHomePage)
	e.GET("/resume", handleResumePage)
	e.GET("/music", handleMusicPage)
	e.File("/static/main.css", "static/main.css")
	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.tmpl")),
	}
	e.Renderer = t
	e.Logger.Fatal(e.Start(":80"))
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type soundcloudUrl struct {
	url  string
	name string
}

func handleMusicPage(c echo.Context) error {
	urls := []soundcloudUrl{
		{url: "https://soundcloud.com/user-434601011/leather-britches", name: "Leather Britches"},
		{url: "https://soundcloud.com/user-434601011/raggedy-ann", name: "Raggedy Ann"}}
	data := map[string]interface{}{
		"urls": urls,
	}
	err := c.Render(http.StatusOK, "musicpage", data)
	if err != nil {
		return err
	}
	return nil
}
func handleResumePage(c echo.Context) error {
	err := c.Render(http.StatusOK, "resumepage", nil)
	if err != nil {
		return err
	}
	return nil
}

func handleHomePage(c echo.Context) error {
	err := c.Render(http.StatusOK, "homepage", nil)
	if err != nil {
		return err
	}
	return nil
}

func getCss(c echo.Context) error {
	c.String(http.StatusOK, "OK")
	return nil
}
