package main

import (
	"fmt"
	"go_function/internal"
)

func main() {
	internal.SimpleFunction()
	internal.ParamFunction("WBSM")

	internal.AnonimalFunction()
	internal.ParamAnonimalFunction()
	internal.AnonimalFunctionWithParam()
	funcI := internal.AnonimalFunctionWithReturn()

	fmt.Println("FUNC I: ", funcI())

	funcS := internal.Closures()

	fmt.Println("FUNC S: ", funcS())
	fmt.Println("FUNC S: ", funcS())

	internal.Variadics(1, 2, 3, 4, 5, 6, 7, 8, 9)

}
