package main

import (
	"GoProjects/greetings"
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"strconv"

	//"rsc.io/quote"
	"log"
)

var c, python, java bool
var (
	c2, python2, java2            = true, false, "yep!"
	z                  complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	Shan  = "世界"
	Chi   = "10"
	Big   = 1 << 10
	Small = Big >> 99
)

func main() {

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings:")
	log.SetFlags(0)

	names := []string{"Alex", "Natali", "Den"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)

	//
	//fmt.Println("Hello and GO!!")
	//fmt.Println(quote.Go())
	//message:=greetings.Hello("Gladys")
	//fmt.Println(message)

	fmt.Println("My favorite number is ", rand.Intn(2))
	fmt.Printf("My favorite number is  %g. \n", math.Sqrt(121))
	fmt.Println(math.Pi)
	fmt.Println(add(42, 1567))
	fmt.Println(swap("Perenion", "Don"))
	fmt.Println(split(3))
	var i int
	var i2, j2 int = 17, 21

	fmt.Println(i, c, python, java)
	fmt.Println(i2, j2, c2, python2, java2)
	k := 57
	fmt.Println(k)

	var r rune
	fmt.Println(r)

	fmt.Printf("Type : %T Value: %v \n", z, z)

	var zero string = "0"
	ze, err := strconv.Atoi(zero)
	fmt.Printf("ze %v is  %T  type\n", ze, ze)
	fmt.Printf("SHAN %v CHI %v is  %T  type $ %T type\n", Shan, Chi, Shan, Chi)


	fmt.Println(Small)
	fmt.Println(Big)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

}

func add(x, y int) int {
	return x + y

}

func swap(a, b string) (string, string) {
	return b, a
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x*0.1
}
