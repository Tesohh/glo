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
	r := glo.NewRouter("", rep)
	r.Routes = map[string]glo.Route[repo]{
		"/hello": {
			Handler: func(w http.ResponseWriter, r *http.Request, repo repo) error {
				fmt.Fprintf(w, "Hello")
				return nil
			},
			Methods: "GET",
		},
	}

	r.Serve(":8080")
}
