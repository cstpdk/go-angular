package main

import (
	"github.com/go-martini/martini"
	"github.com/GeertJohan/go.rice"
	"net/http"
)

func main() {
	m := martini.Classic()
	m.Use(martini.Static("static"))

	http.Handle("/", http.FileServer(rice.MustFindBox("static").HTTPBox()))
	r := martini.NewRouter()
	m.Action(r.Handle)

	http.ListenAndServe(":80", m)
	m.Run()
}
