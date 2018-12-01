package day1

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)


func exists(a []int, freq int) bool {
	for _, n := range a{
		if freq == n {
			return true
		}
	}
	return false
}


func part1(strArray []string){
	freq := 0

	for _, entry := range strArray  {
		n, _ := strconv.Atoi(entry[1:])

		if entry[:1] == "+" {
			freq += n
		}else{
			freq -= n
		}
	}

	fmt.Println(freq)
}

func part2(strArray []string) {
	freq := 0
	var checked []int
	exist := false

	for !exist {

		for _, entry := range strArray {
			checked = append(checked, freq)
			n, _ := strconv.Atoi(entry[1:])

			if entry[:1] == "+" {
				freq += n
			} else {
				freq -= n
			}
			if exists(checked, freq) {
				exist = true
				break
			}
		}
	}

	fmt.Println(freq)
}


func main(){
	b, err := ioutil.ReadFile("day1.txt")
	if err != nil {
		fmt.Print(err)
	}
	strArray := strings.Fields(string(b))
	part1(strArray)
	part2(strArray)


}