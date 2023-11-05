package glo

import "net/http"

func deglo[T any](route Route[T], repo T) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := route.Handler(w, r, repo)

		if err != nil {
			if gerr, ok := err.(GloErr); ok {
				WriteJSON(w, gerr.Status, map[string]string{"error": err.Error()})
				return
			}
			WriteJSON(w, 400, map[string]string{"error": err.Error()})
			return
		}
	}
}
