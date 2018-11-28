package main

import "fmt"

type handler func(...interface{}) interface{}

type array []interface{}

func (arr array) Map(f handler) (ret array) {
	for _, el := range arr {
		ret = append(ret, f(el))
	}
	return
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(a.Map(func(a int) int { return 2 * a }))
}
