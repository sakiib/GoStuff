package main

import (
	"fmt"
	"github.com/sakiib/GoStuff/Packages/myPackages"
)

func main() {
	fmt.Println("Packages!")
	fmt.Println(myPackages.Min(10, 20), myPackages.Max(10, 20))
}