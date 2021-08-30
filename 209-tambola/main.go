// Generate housie tickets

// each ticket should have 3 rows and 9 columns
// each entry can have a blank or a random number from 1 - 90
// The first column can only have numbers from 1  - 10
// The second column can have numbers from 11 - 20
// and so on
// A row must have only 5 numbers
// numbers in a column have to be in increasing order.
// https://www.postermywall.com/index.php/posters/search?s=tambola%20ticket#65e9254f826c1a37d8491be4a59d42dc

// init a 2D slice [9][3]
// valu

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	board := makeBoard()
	// board = initArray(board)
	fmt.Println(board)

}

func randomNo(min int, max int) int {
	time.Sleep(1 * time.Nanosecond)
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(max-min) + min
	return n
}

func makeBoard() [][]string {
	numMap := make(map[int]int)
	board := make([][]string, 3)
	for i := range board {
		board[i] = make([]string, 9)
	}
	i, j := 0, 0

	for i < 3 {
		for j < 9 {
			// board[i][j] = fmt.Sprint(time.Now().UnixNano()) creating board with some value
			n := randomNo(j*10+1, (j+1)*10+1)
			if _, ok := numMap[n]; !ok {
				numMap[n]++
				board[i][j] = fmt.Sprint(n)
				j++
			}
		}
		j = 0
		i++
	}

	return board
}

func getNo(n string) int {
	num, _ := strconv.Atoi(n)
	return num
}

func sortBoard(board [][]string) [][]string {
	// ToDo: sort board
	for j := 0; j < 9; j++ {
		if getNo(board[0][j]) > getNo(board[1][j]) && getNo(board[1][j]) > getNo(board[2][j]) {
			temp := board[0][j]
			board[0][j] = board[2][j]
			board[2][j] = temp
		} else if getNo(board[0][j]) > getNo(board[1][j]) && getNo(board[2][j]) > getNo(board[1][j]) {
			temp := board[0][j]
			board[0][j] = board[1][j]
			board[1][j] = temp
		} else if getNo(board[1][j]) > getNo(board[2][j]) && getNo(board[1][j]) > getNo(board[0][j]) {
			if getNo(board[0][j]) > getNo(board[2][j]) {
				board[0][j], board[1][j], board[2][j] = board[2][j], board[0][j], board[1][j]
			} else {
				board[0][j], board[1][j], board[2][j] = board[0][j], board[2][j], board[1][j]
			}
		}
	}
	return board
}
