package main

import (
	"os"
	"html/template"
)

type Inventory struct {
	Material string
	Count  uint
}
func main() {
	//tmpl.AddDir("./templates")
	//err := tmpl.Execute(os.Stdout, "hello/content", nil)
	//if err != nil {
	//	panic(err)
	//}
	tmpl, err := template.ParseGlob("./templates/**/*.tmpl")
	if err != nil {
		panic(err)
	}
	err = tmpl.ExecuteTemplate(os.Stdout, "hello/content", nil)
	if err != nil {
		panic(err)
	}
}
