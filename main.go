package main

import (
	"bufio"
	"bytes"
	"html/template"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type product struct {
	Img         string
	Name        string
	Price       string
	Stars       float64
	Reviews     int
	Description string
}

func subtr(a, b float64) float64 {
	return a - b
}

func list(e ...float64) []float64 {
	return e
}

func main() {
	initBasicIndexHtml()
	startServer()
}

func startServer() {
	e := echo.New()
	// Routes
	e.GET("/", renderHomePage)

	// Start server
	e.Logger.Fatal(e.Start(":80"))
}

func renderHomePage(c echo.Context) error {
	data := []product{
		{"images/1.png", "strawberries", "$2.00", 4.0, 251, "Lorem ipsum dolor sit amet, consectetur adipiscing elit."},
		{"images/2.png", "onions", "$2.80", 5.0, 123, "Morbi sit amet erat vitae purus consequat vehicula nec sit amet purus."},
		{"images/3.png", "tomatoes", "$3.10", 4.5, 235, "Curabitur tristique odio et nibh auctor, ut sollicitudin justo condimentum."},
		{"images/4.png", "courgette", "$1.20", 4.0, 251, "Phasellus at leo a purus consequat ornare ac aliquam quam."},
		{"images/5.png", "broccoli", "$3.80", 3.5, 123, "Maecenas sed ante sagittis, dapibus dui quis, hendrerit orci."},
		{"images/6.png", "potatoes", "$3.00", 2.5, 235, "Vivamus malesuada est et tellus porta, vel consectetur orci dapibus."},
	}

	allFiles := []string{"content.tmpl", "footer.tmpl", "header.tmpl", "page.tmpl"}

	var allPaths []string
	for _, tmpl := range allFiles {
		allPaths = append(allPaths, "./templates/"+tmpl)
	}

	templates := template.Must(template.New("").Funcs(template.FuncMap{"subtr": subtr, "list": list}).ParseFiles(allPaths...))

	var processed bytes.Buffer
	templates.ExecuteTemplate(&processed, "page", data)

	outputPath := "./static/index.html"
	f, _ := os.Create(outputPath)
	w := bufio.NewWriter(f)
	w.WriteString(string(processed.Bytes()))
	w.Flush()

	// return c.Render(http.StatusOK, "static/index.html", nil)
	println("hitting here")
	err := c.Render(http.StatusOK, "home", template.HTML(string(processed.Bytes())))
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}

func initBasicIndexHtml() {
	data := []product{
		{"images/1.png", "strawberries", "$2.00", 4.0, 251, "Lorem ipsum dolor sit amet, consectetur adipiscing elit."},
		{"images/2.png", "onions", "$2.80", 5.0, 123, "Morbi sit amet erat vitae purus consequat vehicula nec sit amet purus."},
		{"images/3.png", "tomatoes", "$3.10", 4.5, 235, "Curabitur tristique odio et nibh auctor, ut sollicitudin justo condimentum."},
		{"images/4.png", "courgette", "$1.20", 4.0, 251, "Phasellus at leo a purus consequat ornare ac aliquam quam."},
		{"images/5.png", "broccoli", "$3.80", 3.5, 123, "Maecenas sed ante sagittis, dapibus dui quis, hendrerit orci."},
		{"images/6.png", "potatoes", "$3.00", 2.5, 235, "Vivamus malesuada est et tellus porta, vel consectetur orci dapibus."},
	}

	allFiles := []string{"content.tmpl", "footer.tmpl", "header.tmpl", "page.tmpl"}

	var allPaths []string
	for _, tmpl := range allFiles {
		allPaths = append(allPaths, "./templates/"+tmpl)
	}

	templates := template.Must(template.New("").Funcs(template.FuncMap{"subtr": subtr, "list": list}).ParseFiles(allPaths...))

	var processed bytes.Buffer
	templates.ExecuteTemplate(&processed, "page", data)

	outputPath := "./static/index.html"
	f, _ := os.Create(outputPath)
	w := bufio.NewWriter(f)
	w.WriteString(string(processed.Bytes()))
	w.Flush()

}
