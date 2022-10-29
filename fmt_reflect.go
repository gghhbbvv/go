package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello,\nGo!")
	fmt.Println("Hello,\tMe!")
	fmt.Println("Hello, \"Github!\"")
	fmt.Println("Byu, \\\\ GO!")
	fmt.Println('A')
	fmt.Println('\\')
	fmt.Println(42)
	fmt.Println(3.1415 * 4)
	fmt.Println(reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf(3.1415))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("Hello, Go!"))
}
