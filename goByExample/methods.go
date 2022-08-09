package main

import "fmt"

type rectangle struct {
	width, height int
}

func area(val rectangle) int {
	return val.width * val.height
}
func areaP(val *rectangle) int {
	return val.width * val.height
}

func main() {
	fmt.Println(area(rectangle{width: 20, height: 10}))
	fmt.Println(areaP(&rectangle{width: 20, height: 10}))
}
