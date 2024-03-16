package main

import (
	"fmt"
	"os"
)

const (
	lignes, coulons = 9, 9
)

func main() {
	args := os.Args[1:]
	// Check Errors!
	for _, arg := range args {
		for _, char := range arg {
			if len(arg) != 9 || (char < '1' || char > '9') && char != '.' {
				fmt.Println("Error!")
				return
			}
		}
	}

	// Create and Remplir the table
	table := make([][]rune, lignes)
	for i := 0; i < lignes; i++ {
		table[i] = make([]rune, coulons)
		for j, char := range args[i] {
			if (char >= '1' && char <= '9') || char == '.' {
				table[i][j] = rune(char)
			} else {
				fmt.Println("Error!")
				return
			}
		}
	}

	// Solve and print
	if Solution(table) {
		PrintTable(table)
	} else {
		fmt.Println("Error!")
		return
	}
}

func Solution(table [][]rune) bool {
	for i := 0; i < lignes; i++ {
		for j := 0; j < coulons; j++ {
			if table[i][j] == '.' {
				for val := '1'; val <= '9'; val++ {
					if validation(table, i, j, val) {
						table[i][j] = val    // table[i][j] tbdlat donc table kaml tbedel o ykhess ytverifa a 0
						if Solution(table) { // nouvelle verification dyal nouveau table
							return true
						} else {
							table[i][j] = '.'
						}
					}
				}
				return false
			}
		}
	}
	return true
}

func validation(table [][]rune, row, col int, val rune) bool {
	// verify ligne et coulon
	for i := 0; i < 9; i++ {
		if table[row][i] == val || table[i][col] == val {
			return false
		}
	}

	// check 3*3
	depart_r := row - row%3
	depart_c := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if table[i+depart_r][j+depart_c] == val {
				return false
			}
		}
	}
	return true
}

func PrintTable(table [][]rune) {
	for i := 0; i < lignes; i++ {
		for j := 0; j < coulons; j++ {
			fmt.Print(string(table[i][j]))
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
