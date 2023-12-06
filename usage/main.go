package main

import (
	"fmt"
	"net/http"

	"github.com/Tesohh/glo"
)

type repo struct {
	Users []string
}

func main() {
	rep := repo{Users: []string{""}}
	r := glo.NewRouter("/api", rep)
	r.Preprocessor = func(f glo.Func[repo]) glo.Func[repo] {
		return func(w http.ResponseWriter, r *http.Request, repo repo) error {
			fmt.Println("got a request brodie")
			return f(w, r, repo)
		}
	}

	r.Routes = map[string]glo.Route[repo]{
		"/hello": {
			Handler: func(w http.ResponseWriter, r *http.Request, repo repo) error {
				fmt.Fprintf(w, "Hello")
				return nil
			},
			Methods: "GET",
		},
		"/error": {
			Handler: func(w http.ResponseWriter, r *http.Request, repo repo) error {
				return fmt.Errorf("asddsdd")
			},
			Methods: "GET",
		},
	}

	r.Serve(":8080")
}
