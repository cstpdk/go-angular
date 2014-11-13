package main

import (
	"github.com/go-martini/martini"
	"net/http"
)

func main() {
	m := martini.Classic()
	m.Use(martini.Static("static"))

	r := martini.NewRouter()
	m.Action(r.Handle)

	http.ListenAndServe(":80", m)
	m.Run()
}
