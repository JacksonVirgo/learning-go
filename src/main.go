package main

import (
	"fmt"
)

func main() {
	var intNum int = 10
	fmt.Println(intNum)

	value := 100
	fmt.Println(value)

	const PI float32 = 3.14
	fmt.Println(PI)

	var str, err = echo("Let's Fucking Go")
	if err != nil {
		fmt.Printf("Echoed: %s\n", str)	
	} else {
		fmt.Printf("Error: %s\n", str)
	}
}

func echo(param string) (string, error) {
	var err error
	fmt.Println(param)

	if true {
		return param, nil
	}

	return param, err
}