package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)
const sheetSize = 1000

type Claim struct {
	id string
	x int
	y int
	width int
	height int
}

func (c* Claim) setPos(pos string) {
	c.x, _ = strconv.Atoi(pos[:strings.Index(pos, ",")])
	c.y, _ = strconv.Atoi(pos[strings.Index(pos, ",") + 1 :strings.Index(pos, ":")  ])
}

func (c* Claim) setDim(dim string) {
	c.width, _ = strconv.Atoi(dim[:strings.Index(dim, "x")])
	c.height, _ = strconv.Atoi(dim[strings.Index(dim, "x") + 1:])
}

func (c* Claim) setId(id string) {
	c.id= id[1:]
}

func generateClaims(strArray []string)[]Claim{
	var claim Claim
	var claims []Claim
	for _, line := range strArray {
		claimOrder := strings.Fields(line)
		claim.setId(claimOrder[0])
		claim.setPos(claimOrder[2])
		claim.setDim(claimOrder[3])

		claims = append(claims, claim)

	}
	return claims
}

func doClaim (sheet* [sheetSize][sheetSize]string, claim Claim, sqrInch* int) {
	for x := claim.x; x < claim.width + claim.x; x++ {
		for y := claim.y; y < claim.height + claim.y; y++ {
			switch sheet[x][y] {
			case "":
				sheet[x][y] = claim.id
			case "X":
				break
			default:
				sheet[x][y] = "X"
				*sqrInch += 1
			}
		}
	}
}

func intact(claim Claim, sheet [sheetSize][sheetSize]string) bool{
	for x := claim.x; x < claim.x + claim.width; x++ {
		for y := claim.y; y < claim.y + claim.height; y++ {
			if sheet[x][y] != claim.id {
				return false
			}
		}
	}
	return true
}

func part1(claims []Claim, sheet* [sheetSize][sheetSize]string){
	sqrInch := 0
	for _, claim := range claims {
		doClaim(sheet, claim, &sqrInch)
	}

	fmt.Println(sqrInch)
}

func part2(claims []Claim, sheet [sheetSize][sheetSize]string) {
	for _, claim := range claims {
		if intact(claim, sheet){
			fmt.Println(claim.id)
			return
		}
	}
}


func main(){
	b, err := ioutil.ReadFile("day3/day3.txt")
	if err != nil {
		fmt.Print(err)
	}
	strArray := strings.FieldsFunc(string(b), func(r rune) bool {
		if r == '\n' {
			return true
		}
		return false
	})


	sheet := [sheetSize][sheetSize]string{}
	claims := generateClaims(strArray)

	part1(claims, &sheet)
	part2(claims, sheet)
}