package routers

import (
	"fmt"
	"testing"
)

type Usb interface {
	Open() string
	Close() string
}

type Mobile struct {
	name string
}

func (mobile Mobile) Open() string {
	fmt.Println("打开")
	fmt.Println(mobile.name)
	return "打开"
}

func (mobile Mobile) Close() string {
	fmt.Println("打开")
	fmt.Println(mobile.name)
	return "打开"
}

func add[T int | float32 | float64](a T, b T) T {
	return a + b
}

func TestInterface(t *testing.T) {
	var typeC Usb
	typeC = Mobile{"xiaomi"}
	typeC.Open()
	t2 := add(1.0, 1.2)
	t3 := add(10, 12)
	fmt.Println(t2)
	fmt.Println(t3)
}
