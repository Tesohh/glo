package glo

type GloErr struct {
	Err    error
	Status int
}

func (e GloErr) Error() string {
	return e.Err.Error()
}
