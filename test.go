package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	ui "github.com/logrusorgru/aurora"
)

// Buff are buffs for fighters
type Buff struct {
	AtkUp   int
	DefUp   int
	DodgeUp int
	CritUp  int
}

// Fighter is figher
type Fighter struct {
	Name      string
	Rank      int
	Item      int
	Props     []int
	Buffs     []Buff
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
func CheckMix(fighters []Fighter) Buff {
	pc := map[int]int{}
	for _, f := range fighters {
		for _, p := range f.Props {
			pc[p]++
		}
	}
	bf := Buff{}
	if pc[1] > 3 {
		bf.AtkUp += 10
	}
	return bf
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

func rankFormat(char string, rank int) ui.Value {
	if rank == 1 {
		return ui.White(char)
	}
	if rank == 2 {
		return ui.Cyan(char)
	}
	return ui.Bold(ui.BgBrightRed(char))
}
func main() {

	temp := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 2, 3, 0, 0, 0},
		{0, 0, 3, 9, 1, 0, 0, 0},
		{0, 0, 5, 6, 7, 8, 0, 0},
	}
	pool := []int{}
	for k, v := range poolSetting {
		for i := 0; i < v; i++ {
			pool = append(pool, k)
		}
	}

	fmt.Println("Hello,", ui.Magenta("Abby"))
	holding := []int{}
	for {
		Shuffle(pool)
		fmt.Print(ui.Bold(ui.Cyan("｜－－－－－－－－－\n")))
		for i := 0; i < 8; i++ {
			fmt.Print(ui.Bold(ui.Cyan("｜")))
			for j := 0; j < 8; j++ {
				if i <= 3 && temp[i][j] != 0 {
					fmt.Print(rankFormat(ei[temp[i][j]], 3))
				} else {
					fmt.Print(ui.Cyan("　"))
				}
			}
			fmt.Print(ui.Bold(ui.Cyan("｜")))
			fmt.Print("\n")
		}
		fmt.Print(ui.Bold(ui.Cyan("－－－－－－－－－｜\n")))

		fmt.Print(ui.Bold(ui.Cyan("Holding:")))
		for i := 0; i < len(holding); i++ {
			fmt.Print(holding[len(holding)-1])
			holding = holding[:len(holding)-1]
			fmt.Print(" ")
		}
		fmt.Print("\n")
		fmt.Println(ui.Bold(ui.Cyan("Shop:")))
		for i := 0; i < 5; i++ {
			fmt.Print(ei[pool[len(pool)-1]])
			pool = pool[:len(pool)-1]
			if i != 4 {
				fmt.Print(" ")
			} else {
				fmt.Print("\n")
			}
		}
		fmt.Println(ui.Bold(ui.Cyan("Enter:")))

		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
		v, _ := strconv.Atoi(text)
		holding = append(holding, v)
		t := time.Now()
		h, m, s := t.Clock()
		fmt.Println("Time: ", h, ":", m, ":", s)
	}
}
