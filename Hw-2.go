package main

import (
	"errors"
	"fmt"
)

type Divided struct {
	a   float64
	b   float64
	res string
}

func (d Divided) New() float64 {
	return d.a / d.b
}

func (d Divided) Error() string {
	return fmt.Sprintf("%v(%e)", d.a, d.b)
}

var err error

func div(a, b float64) (float64, error) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("err in divided: ", e)
			err = e.(error)
		}
	}()
	if b == 0 {
		panic(errors.New("divided by zero"))
	}
	return a / b, nil

}

func main() {
	res, err := div(4, 0)
	fmt.Println("Division:", res, err)
}
