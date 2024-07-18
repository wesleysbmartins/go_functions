package internal

import "fmt"

func AnonimalFunction() {
	// anonimal func create
	anonimalF := func() {
		fmt.Println("Anonimal Print!")
	}

	// anonimal func run
	anonimalF()
}

func ParamAnonimalFunction() {
	// anonimal func create
	anonimalF := func(param string) {
		fmt.Println("Anonimal Param: ", param)
	}

	// anonimal func run
	anonimalF("Anonimal Func Param")
}

func AnonimalFunctionWithReturn() func() int {
	n := 10

	return func() int {
		return n * 5
	}
}

func AnonimalFunctionWithParam() {
	n := func() int {
		return 10
	}

	r := func(i int) int {
		return i * 5
	}

	result := r(n())

	fmt.Println("Result: ", result)
}
