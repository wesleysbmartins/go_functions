package internal

import (
	"fmt"
	"strconv"
	"strings"
)

func Variadics(numbers ...int) {
	values := []string{}

	for i, v := range numbers {
		values = append(values, fmt.Sprintf("Index: %s Value: %s", strconv.Itoa(i), strconv.Itoa(v)))
	}

	fmt.Println("Result:\n", strings.Join(values, "\n"))
}
