package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)


func twiceFunc (a rune, word string) bool {
	return strings.Count(word, string(a)) == 2
}

func thriceFunc (a rune, word string) bool {
	return strings.Count(word, string(a)) == 3
}

func common(a string, b string) string {
	var c string
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			c += string(a[i])
		}
	}
	return c
}


func part1(strArray []string){
	twice := 0
	thrice := 0

	for _, entry := range strArray  {
		hasTwice := false
		hasThrice := false
		for _, char := range entry{
			if !hasTwice {
				hasTwice = twiceFunc(char, entry)
			}

			if !hasThrice {
				hasThrice = thriceFunc(char, entry)
			}

			if hasTwice && hasThrice {
				break
			}

		}
		if hasTwice { twice += 1}
		if hasThrice {thrice += 1}
	}

	fmt.Println(twice * thrice)
}

func part2(strArray []string) {
	var id string
	for i, a := range strArray {
		for j, b := range strArray {
			if i != j {
				temp := common(a, b)
				if len(temp) > len(id) {
					id = temp
				}
			}
		}
	}
	fmt.Println(id)
}


func main(){
	b, err := ioutil.ReadFile("day2/day2.txt")
	if err != nil {
		fmt.Print(err)
	}
	strArray := strings.Fields(string(b))
	part1(strArray)
	part2(strArray)
}