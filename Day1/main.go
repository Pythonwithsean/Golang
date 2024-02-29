package main

import (
	"fmt"
)

func testing(num int, arr *[]string) int {

	fmt.Println(*arr)
	return num

}

func f(g func(n int, b int) int, c int) func(a int, b int) int {

	return g
}

func main() {

	// in Golang all variables must begin with small letter unless you want to Export that var
	// When importing a package you need to refer to the exported names any export names are not accessible outside its package
	// * in Golang returns the value the pointer is pointing too
	// p *int p is a pointer that points to integers and int is the base
	//* then the name of the variable is known as dereferencing

	var names = []string{"Sean", "Steven"}
	g := func(a int, b int) int {
		return a + b
	}
	sum := func(a, b int) int { return a + b }(3, 4)
	println(sum)
	k := f(g, 2)(2, 3)
	fmt.Println(k)
	testing(22, &names)
	x := (*int)(nil)

	fmt.Println(x)
	// fmt.Println(rand.Intn(10))
	// fmt.Println(time.Now().Date())
	// fmt.Println("Go is Better")
	// var age int = 2
	// fmt.Print(age + 2)
}
