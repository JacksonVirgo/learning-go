package main

import (
	"fmt"
)

type strawberry struct {
	sweetness uint8
	bitterness uint8
}

func (s strawberry) String() string {
	return fmt.Sprintf("Sweetness: %d, Bitterness: %d", s.sweetness, s.bitterness)
}

func (s strawberry) isSweet() bool {
	return true
}

type fruit interface {
	isSweet() bool
}

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

	var s  = strawberry{1, 2}
	fmt.Println(s)
	fmt.Println(s.sweetness)

	fmt.Println(s.String())

	fruit_blah(s)
}

func echo(param string) (string, error) {
	var err error
	fmt.Println(param)

	if true {
		return param, nil
	}

	return param, err
}

func fruit_blah(f fruit) {
	fmt.Println(f.isSweet())
}