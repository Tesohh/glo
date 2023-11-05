package glo

import "net/http"

type Func[T any] func(w http.ResponseWriter, r *http.Request, repo T) error
type MWFunc[T any] func(f Func[T]) Func[T]
