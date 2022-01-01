package main

import "net/http"

func main() {
	sta := http.FileServer(http.Dir("static"))
	server := Newserver(":3000")
	server.Handle("GET", "/", HandleRoot)
	//server.Handle("GET", "/login", HandleLog)
	server.Handle("GET", "/login", HandleLog)
	http.Handle("/static/", http.StripPrefix("/static/", sta))
	server.Listen()
}
