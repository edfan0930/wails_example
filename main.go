package main

import (
	"fmt"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

type Robot struct {
	Name    string
	runtime *wails.Runtime
}

func NewRobot() *Robot {
	result := &Robot{
		Name: "Robbie",
	}
	return result
}

func (t *Robot) Hello(name string) string {
	return fmt.Sprintf("Hello %s! My name is %s", name, t.Name)
}

func (t *Robot) Rename(name string) string {
	t.Name = name
	return fmt.Sprintf("My name is now '%s'", t.Name)
}

func (t *Robot) privateMethod(name string) string {
	t.Name = name
	return fmt.Sprintf("My name is now '%s'", t.Name)
}

func (t *Robot) WailsInit(runtime *wails.Runtime) error {
	t.runtime = runtime
	return nil
}

func basic() string {
	return "Hello World!"
}

func twokey(say string) string {

	return say
}

func (t *Robot) TryEmit() string {

	t.runtime.Events.Emit("try", "try event")
	return "in try emit"
}

func main() {

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "wails_example",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(twokey)
	app.Bind(basic)
	app.Bind(NewRobot())
	app.Run()
}