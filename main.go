package main

import (
	"github.com/go-martini/martini"
	"github.com/GeertJohan/go.rice"
	"net/http"
)

func main() {
	m := martini.Classic()
	m.Use(martini.Static("static"))

	r := martini.NewRouter()
	m.Action(r.Handle)

	http.Handle("/", http.FileServer(rice.MustFindBox("static").HTTPBox()))
	http.ListenAndServe(":8080", m)

}
