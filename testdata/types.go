package foo

var (
	cond, cond2, cond3 bool

	num int

	slice []string

	Sink interface{}

	errVar error

	fnErr  func() (*int, error)
	action func() bool
)
