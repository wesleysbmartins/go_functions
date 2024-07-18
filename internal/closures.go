package internal

func Closures() func() string {

	var n int

	l := []string{"Closure result", "closure result two"}

	return func() string {
		toReturn := l[n]
		n++
		return toReturn
	}
}
