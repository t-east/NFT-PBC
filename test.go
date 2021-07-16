package main

import (

	"fmt"

	"math/rand"

)
func main() {
	for i:=0;i<1;i++{
		ck := rand.Intn(5) +1
		fmt.Print(ck)
		fmt.Print("\n")
	}
}