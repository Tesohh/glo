package glo

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route[T any] struct {
	Handler Func[T]
	// separate methods by comma
	// example: "GET,POST,PATCH"
	Methods string
}

func NewRouter[T any](prefix string, repo T) Router[T] {
	return Router[T]{Prefix: prefix, r: mux.NewRouter(), repo: repo}
}

type Router[T any] struct {
	// can access all gorilla mux methods from this router
	r          *mux.Router
	repo       T
	Routes     map[string]Route[T]
	Prefix     string
	Middleware []Func[T]
}

func (r Router[T]) Serve(addr string) {
	for k, v := range r.Routes {
		r.r.HandleFunc(k, deglo(v, r.repo)).Methods(strings.Split(v.Methods, ",")...)
	}

	fmt.Println(
		"      _\n" +
			"     | |\n" +
			" __, | |  __\n" +
			"/  | |/  /  \\_\n" +
			"\\_/|/|__/\\__/\n" +
			" /|\n" +
			" \\|")

	fmt.Printf("Serving %v routes on address %s\n", len(r.Routes), addr)

	http.ListenAndServe(addr, r.r)
}

// fmt.Println(
// 	"	     ,dPYb,\n" +
// 		"             IP'`Yb\n" +
// 		"             I8  8I\n" +
// 		"             I8  8'\n" +
// 		"   ,gggg,gg  I8 dP    ,ggggg,\n" +
// 		"  dP\"  \"Y8I  I8dP    dP\"  \"Y8ggg\n" +
// 		" i8'    ,8I  I8P    i8'    ,8I\n" +
// 		",d8,   ,d8I ,d8b,_ ,d8,   ,d8'\n" +
// 		"P\"Y8888P\"8888P'\"Y88P\"Y8888P\"\n" +
// 		"       ,d8I'\n" +
// 		"     ,dP'8I\n" +
// 		"    ,8\"  8I\n" +
// 		"    I8   8I\n" +
// 		"    `8, ,8I\n" +
// 		"     `Y8P\"")
