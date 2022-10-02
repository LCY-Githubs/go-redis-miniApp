package webDemo

import (
	"fmt"
	"testing"
)

type Method interface {
	get() string
}

func get() string {
	return "sssada"
}
func TestWebDemo(t *testing.T) {
	// fmt.Println("sss")
	// fmt.Println(get())

	var a = 1
	fmt.Println(a)
	b := &a
	fmt.Println(b)
	fmt.Println(a)

}
