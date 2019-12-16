package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	ui "github.com/logrusorgru/aurora"
)

// Effect are buffs or debuffs for fighters
type Effect struct {
	Atk         int
	Def         int
	Dodge       int
	Crit        int
	HP          int
	MultiTarget int
	TargetAlly  bool
}

// Fighter is figher
type Fighter struct {
	Name      string
	Rank      int
	Item      int
	Props     []int
	Buffs     []Effect
	BaseAtk   []int
	BaseDef   []int
	BaseAs    []int
	BaseHP    []int
	CD        []int
	SkillVal  []int
	SkillType int
}

// Shuffle int array
func Shuffle(vals []int) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(vals) > 0 {
		n := len(vals)
		randIndex := r.Intn(n)
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
		vals = vals[:n-1]
	}
}

// CheckMix checks if fighter trigger buffs
func CheckMix(fighters []Fighter) []Effect {
	pc := map[int]int{}
	for _, f := range fighters {
		for _, p := range f.Props {
			pc[p]++
		}
	}
	bf := Effect{}
	if pc[1] > 3 {
		bf.Atk += 10
		bf.TargetAlly = true
	}
	return []Effect{bf}
}

var fighterSettings = map[int]Fighter{
	1: Fighter{
		Name:  "甲",
		Props: []int{1, 2},
	},
	2: Fighter{
		Name:  "乙",
		Props: []int{1, 2},
	},
}
var ei = map[int]string{
	0:  "  ",
	1:  "甲",
	2:  "乙",
	3:  "丙",
	4:  "丁",
	5:  "戊",
	6:  "己",
	7:  "庚",
	8:  "辛",
	9:  "壬",
	10: "癸",
}

var poolSetting = map[int]int{
	1:  10,
	2:  10,
	3:  10,
	4:  10,
	5:  10,
	6:  10,
	7:  10,
	8:  10,
	9:  10,
	10: 10,
}
var reader = bufio.NewReader(os.Stdin)

var emptyRow = []int{0, 0, 0, 0, 0, 0, 0, 0}
var emptyBoard = [][]int{emptyRow, emptyRow}

func rankFormat(char string, rank int) ui.Value {
	if rank == 1 {
		return ui.White(char)
	}
	if rank == 2 {
		return ui.Cyan(char)
	}
	return ui.Bold(ui.BgBrightRed(char))
}
func cin() string {
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}
func coutError(txt string) {
	fmt.Println(ui.Bold(ui.Red(txt)))
}
func cout(txt string) {
	fmt.Print(txt)
}

type styleFunc func(arg interface{}) ui.Value

// https://en.wikipedia.org/wiki/Box-drawing_character
func drawBoxWith(temp [][]int, sf styleFunc) {
	fmt.Print(ui.Bold(sf("┏━━━━━━━━━━━━━━━━━━┓\n")))
	for i := 0; i < len(temp); i++ {
		fmt.Print(ui.Bold(sf("┃ ")))
		for j := 0; j < len(temp[i]); j++ {
			if temp[i][j] != 0 {
				fmt.Print(rankFormat(ei[temp[i][j]], 3))
			} else {
				fmt.Print(ui.Cyan("　"))
			}
		}
		fmt.Print(ui.Bold(sf(" ┃")))
		fmt.Print("\n")
	}
	fmt.Print(ui.Bold(sf("┗━━━━━━━━━━━━━━━━━━┛\n")))

}

func main() {
	fmt.Print(ui.Bold(ui.Cyan("Welcome To AutoGo.\nEnter You Name: ")))

	// temp := [][]int{
	// 	{0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 2, 3, 0, 0, 0},
	// 	{0, 0, 3, 9, 1, 0, 0, 0},
	// 	{0, 0, 5, 6, 7, 8, 0, 0},
	// }
	// init pool
	pool := []int{}
	for k, v := range poolSetting {
		for i := 0; i < v; i++ {
			pool = append(pool, k)
		}
	}
	myname, _ := reader.ReadString('\n')
	fmt.Println("Hello,", ui.Magenta(myname))

	holding := []int{}
	for {
		Shuffle(pool)
		shop := []int{}
		for i := 0; i < 5; i++ {
			shop = append(shop, pool[len(pool)-1])
			pool = pool[:len(pool)-1]
			if i != 4 {
				fmt.Print(" ")
			} else {
				fmt.Print("\n")
			}
		}

		drawBoxWith(emptyBoard, func(arg interface{}) ui.Value { return ui.Bold(ui.Yellow(arg)) })
		for {
			fmt.Print(ui.Bold(ui.Cyan("Holding: ")))
			for _, v := range holding {
				fmt.Print(ei[v])
				//holding = holding[:len(holding)-1]
				fmt.Print(" ")
			}
			fmt.Print("\n")
			fmt.Print(ui.Bold(ui.Cyan("   Shop: ")))
			for i, c := range shop {
				fmt.Print(ei[c])
				if i != 4 {
					fmt.Print(" ")
				} else {
					fmt.Print("\n")
				}
			}

			fmt.Println(ui.Bold(ui.Cyan("Enter: (b)uy, (u)p, (s)ell, (d), (p)ut, (m)ove")))

			cmd := cin()
			if cmd == "b" {

				if len(holding) == 8 {
					coutError("Your hand is full.")
				} else {
					cout("Which one? (1-5)")
					v, _ := strconv.Atoi(cin())
					if v == 0 || v > 5 || shop[v-1] == 0 {
						coutError("Invalid Input.")
					} else {
						holding = append(holding, shop[v-1])

						shop[v-1] = 0
					}
				}
			} else if cmd == "u" {
			} else if cmd == "s" {
			} else if cmd == "d" {
			} else if cmd == "p" {
			} else if cmd == "m" {
			} else if cmd == "next" {
				break
			} else {
				coutError("Invalid Input.")
			}
		}
		print("test")

		//ti := time.Now()
		//h, m, s := ti.Clock()
		//fmt.Println("Time: ", h, ":", m, ":", s)
	}
}
