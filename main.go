package main

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	startServer()
}

func startServer() {
	e := echo.New()
	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.tmpl")),
	}
	e.GET("/", handleHomePage)
	e.GET("/resume", handleResumePage)
	e.GET("/music", handleMusicPage)
	e.File("/static/main.css", "static/main.css")
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
	Url  string
	Name string
}

func handleMusicPage(c echo.Context) error {
	urls := []soundcloudUrl{
		// {url: "https://soundcloud.com/user-434601011/leather-britches", name: "Leather Britches"},

		{Url: `<iframe width="100%" height="300" scrolling="no" frameborder="no" allow="autoplay" src="https://w.soundcloud.com/player/?url=https%3A//api.soundcloud.com/tracks/1299581302&color=%23ff5500&auto_play=false&hide_related=true&show_comments=false&show_user=true&show_reposts=false&show_teaser=false&visual=true"></iframe><div style="font-size: 10px; color: #cccccc;line-break: anywhere;word-break: normal;overflow: hidden;white-space: nowrap;text-overflow: ellipsis; font-family: Interstate,Lucida Grande,Lucida Sans Unicode,Lucida Sans,Garuda,Verdana,Tahoma,sans-serif;font-weight: 100;"><a href="https://soundcloud.com/user-434601011" title="Andrew Willette" target="_blank" style="color: #cccccc; text-decoration: none;">Andrew Willette</a> · <a href="https://soundcloud.com/user-434601011/carrol-county-blues" title="Carrol County Blues" target="_blank" style="color: #cccccc; text-decoration: none;">Carrol County Blues</a></div>`, Name: "Raggedy Ann"},
	}
	data := map[string]interface{}{
		"urls": urls,
	}
	err := c.Render(http.StatusOK, "musicpage", data)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func handleResumePage(c echo.Context) error {
	err := c.Redirect(http.StatusPermanentRedirect, "https://andrewwillette.s3.us-east-2.amazonaws.com/newdir/resume.pdf")
	if err != nil {
		log.Println(err)
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
