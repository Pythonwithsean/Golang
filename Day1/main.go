package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// in Golang all variables must begin with small letter unless you want to Export that var
	// When importing a package you need to refer to the
	fmt.Println(rand.Intn(10))
	fmt.Println(time.Now().Date())
	fmt.Println("Go is Better")
	var age int = 2
	fmt.Print(age + 2)
}
