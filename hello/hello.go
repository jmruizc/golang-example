package main

import (
	"fmt"
	"github.com/golang/example/stringutil"
	"math/rand"
	"time"
	"github.com/golang/example/functionsutil"
	"math"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Hello, my friend.")
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	fmt.Println("My favorite number is", rand.Intn(1000))
	fmt.Println("14 + 28 =:", functionsutil.SumIntegers(14, 28))

	a, b := functionsutil.Swapstrings("hello", "world")
	fmt.Println("Swap (hello world):", a +" "+ b)

	sin, cos := functionsutil.Sincos(float64(math.Pi))
	fmt.Println("Sin(Pi):", math.Trunc(sin))
	fmt.Println("Cos(Pi):", cos)
}