package main

import "fmt"

func main() {
	var havedick bool
	var dicksize int
	var question string

	fmt.Println("do u have dick? (yes/no)")
	fmt.Scanln(&question)

	if question == "yes" {
		havedick = true
	} else {
		havedick = false
	}
	if havedick == false {
		fmt.Println("fuck off")
	} else {
		fmt.Println("enter ur size of dick:")
		fmt.Scanln(&dicksize)
		dicksize--
		fmt.Println("ur REAL size of dick:", dicksize, "cm")
		fmt.Println(holeorgasm(dicksize))
	}
	fmt.Println("lmao")
}

func holeorgasm(size int) string {
	if size > 11 && size < 18 {
		return ("hole will enjoy")
	} else {
		return ("hole wont enjoy")
	}
}
