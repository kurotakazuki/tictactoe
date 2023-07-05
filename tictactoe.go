package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Board struct {
	tokens []int // tokens[0] -> (0,0), tokens[1] -> (0,1), ...
}

//マルバツを入力する関
func (b *Board) put(x, y int, u string) {
	if u == "o" {
		b.tokens[x+3*y] = 1
	} else if u == "x" {
		b.tokens[x+3*y] = 2
	}
}

//マルかバツのどちらかを受け取る関数
func (b *Board) get(x, y int) string {
	if b.tokens[x+3*y] == 1 {
		return "o"
	} else if b.tokens[x+3*y] == 2 {
		return "x"
	} else {
		return "."
	}
} // マルバツを表示する関数
func (b *Board) show() {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			fmt.Print(b.get(row, col))
		}
		fmt.Print("\n")
	}
}

//マルバツの勝ち負けを判定する関数
func (b *Board) judge() string {
    winningPatterns := [8][3]int{
        {0, 1, 2}, // top row
        {3, 4, 5}, // middle row
        {6, 7, 8}, // bottom row
        {0, 3, 6}, // left column
        {1, 4, 7}, // middle column
        {2, 5, 8}, // right column
        {0, 4, 8}, // diagonal from left top to right bottom
        {2, 4, 6}, // diagonal from right top to left bottom
    }

    for _, pattern := range winningPatterns {
        if b.tokens[pattern[0]] != 0 && b.tokens[pattern[0]] == b.tokens[pattern[1]] && b.tokens[pattern[1]] == b.tokens[pattern[2]] {
            if b.tokens[pattern[0]] == 1 {
                return "o"
            } else if b.tokens[pattern[0]] == 2 {
                return "x"
            }
        }
    }

    for _, token := range b.tokens {
        if token == 0 {
            return "yet"
        }
    }

    return "draw"
}

func main() {
	i := 0
	b := &Board{tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	for i < 9 {
		fmt.Printf("Player %d: Input (x,y) ", (i%2)+1)
		var s string
		_, err := fmt.Scanln(&s)
		if err != nil {
			panic(err)
		}
		arr := strings.Split(s, ",")
		var player string
		if i%2 == 0 {
			player = "o"
		} else {
			player = "x"
		}
		row, _ := strconv.Atoi(arr[0])
		col, _ := strconv.Atoi(arr[1])
		b.put(row, col, player)
		b.show()
    	if b.judge() == "o"{
    		fmt.Printf("Player 1 won\n")
    		break
    	}else if b.judge() == "x"{
    		fmt.Printf("Player 2 won\n")
    		break
    	}
		i++
	}
}
