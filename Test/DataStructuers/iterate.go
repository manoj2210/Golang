package main

func Iterate(a string) string {
	for index := 0; index < 3; index++ {
		a += a
	}
	return a
}
