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
	// middleware that will be added to the routers mw chain
	// Middleware   []Func[T]
	// Preprocessor Func[T]
	ErrorHandler func(w http.ResponseWriter, err error)
}

func NewRouter[T any](prefix string, repo T) Router[T] {
	return Router[T]{Prefix: prefix, r: mux.NewRouter(), repo: repo, ErrorHandler: defaultErrorHandler}
}

type Router[T any] struct {
	// can access all gorilla mux methods from this router
	r            *mux.Router
	repo         T
	Routes       map[string]Route[T]
	Prefix       string
	Preprocessor MWFunc[T]
	// Middleware   []Func[T]
	// will trickle down to routes that dont have a ErrorHandler set
	ErrorHandler func(w http.ResponseWriter, err error)
}

func (r Router[T]) Serve(addr string) {
	for k, v := range r.Routes {
		if v.ErrorHandler == nil {
			v.ErrorHandler = r.ErrorHandler
		}
		if r.Preprocessor != nil {
			v.Handler = r.Preprocessor(v.Handler)
		}
		cleanPrefix := strings.TrimRight(r.Prefix, "/ \n")
		r.r.HandleFunc(cleanPrefix+k, deglo(v, r.repo)).Methods(strings.Split(v.Methods, ",")...)
	}

	fmt.Printf("Serving %v routes on address %s\n", len(r.Routes), addr)

	http.ListenAndServe(addr, r.r)
}

// fmt.Println(
// 	"      _\n" +
// 		"     | |\n" +
// 		" __, | |  __\n" +
// 		"/  | |/  /  \\_\n" +
// 		"\\_/|/|__/\\__/\n" +
// 		" /|\n" +
// 		" \\|")
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
