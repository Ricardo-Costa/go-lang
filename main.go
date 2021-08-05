package main

import "github.com/Ricardo-Costa/go-lang/seeders"
import "github.com/Ricardo-Costa/go-lang/webserver"

func main() {
	seeders.LoadData()
	webserver := webserver.NewWebServer()
	webserver.Serve()
}
