package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

var table [4][4]int
var freePositions [16][2]int

func main() {
	fillTheTable()

	result := game()
	printTable()
	if result {
		fmt.Println("Congratulations you won the game!")
	} else {
		fmt.Println("You lost!")
	}
}

func game() bool {
	for i := 0; i < 2; i++ {
		AddNumber()
	}
	for {
		printTable()
		fmt.Println()
		fmt.Println()
		fmt.Println("Make your move.")
		fmt.Println()
		fmt.Println("U - Up")
		fmt.Println("D - Down")
		fmt.Println("L - Left")
		fmt.Println("R - Right")
		var inp string
		fmt.Scan(&inp)
		var x bool
		switch inp {
		case "U":
			x = upMovement()
		case "D":
			x = downMovement()
		case "L":
			x = leftMovement()
		case "R":
			x = rightMovement()
		}

		if x == false {
			continue
		}

		AddNumber()
		fmt.Println()
		if count() == 16 && check() == false {
			return false
		}
		if isWon() {
			break
		}
	}
	return true
}

func isWon() bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if table[i][j] == 2048 {
				return true
			}
		}
	}
	return false
}

func upMovement() bool {
	x := goUp()
	y := addUp()
	if !x && !y {
		return false
	}
	goUp()
	return true
}

func downMovement() bool {
	x := goDown()
	y := addDown()
	if !x && !y {
		return false
	}
	goDown()
	return true
}

func leftMovement() bool {
	x := goLeft()
	y := addLeft()
	if !x && !y {
		return false
	}
	goLeft()
	return true
}

func rightMovement() bool {
	x := goRight()
	y := addRight()
	if !x && !y {
		return false
	}
	goRight()
	return true
}

func count() int {
	count := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if table[i][j] != 0 {
				count++
			}

		}
	}
	return count
}

func fillTheTable() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			table[i][j] = 0
		}
	}
}

func AddNumber() {
	newNum := 2
	x := rand.Intn(10) + 1
	if x == 1 {
		newNum = 4
	}
	for {
		if count() == 16 {
			break
		}
		m := rand.Intn(4)
		n := rand.Intn(4)
		if table[n][m] == 0 {
			table[n][m] = newNum
			break
		}
	}
}

func goUp() bool {
	i := 0
	x := false
	for i < 4 {
		pos := 0
		j := 0
		for j < 4 {
			if table[j][i] == 0 {
				j++
				continue
			} else {
				if pos == j {
					pos++
					j++
					continue
				}
				table[pos][i] = table[j][i]
				table[j][i] = 0
				pos++
				x = true
			}
		}
		i++
	}
	return x
}

func addUp() bool {
	i := 0
	x := false
	for i < 4 {
		j := 0
		for j < 4 {
			if j != 3 {
				if table[j][i] == table[j+1][i] && table[j][i] != 0 {
					table[j][i] = table[j][i] + table[j+1][i]
					table[j+1][i] = 0
					j += 2
					x = true
				} else {
					j++
				}
			} else {
				j++
			}
		}
		i++
	}
	return x
}

func goDown() bool {
	i := 3
	x := false
	for i >= 0 {
		pos := 3
		j := 3
		for j >= 0 {
			if table[j][i] == 0 {
				j--
				continue
			} else {
				if pos == j {
					pos--
					j--
					continue
				}
				table[pos][i] = table[j][i]
				table[j][i] = 0
				pos--
				x = true
			}
		}
		i--
	}
	return x
}

func addDown() bool {
	i := 0
	x := false
	for i < 4 {
		j := 3
		for j >= 0 {
			if j != 0 {
				if table[j][i] == table[j-1][i] && table[j][i] != 0 {
					table[j][i] = table[j][i] + table[j-1][i]
					table[j-1][i] = 0
					j -= 2
					x = true
				} else {
					j--
				}
			} else {
				j--
			}
		}
		i++
	}
	return x
}

func goLeft() bool {
	i := 0
	x := false
	for i < 4 {
		pos := 0
		j := 0

		for j < 4 {
			if table[i][j] == 0 {
				j++
				continue
			} else {
				if pos == j {
					pos++
					j++
					continue
				}
				table[i][pos] = table[i][j]
				table[i][j] = 0
				pos++
				x = true
			}
		}
		i++
	}
	return x
}

func addLeft() bool {
	i := 0
	x := false
	for i < 4 {
		j := 0
		for j < 4 {
			if j != 3 {
				if table[i][j] == table[i][j+1] && table[i][j] != 0 {
					table[i][j] = table[i][j] + table[i][j+1]
					table[i][j+1] = 0
					j += 2
					x = true
				} else {
					j++
				}
			} else {
				j++
			}
		}
		i++
	}
	return x
}

func goRight() bool {
	i := 0
	x := false
	for i < 4 {
		pos := 3
		j := 3
		for j >= 0 {
			if table[i][j] == 0 {
				j--
				continue
			} else {
				if pos == j {
					pos--
					j--
					continue
				}
				table[i][pos] = table[i][j]
				table[i][j] = 0
				pos--
				x = true
			}
		}
		i++
	}
	return x
}

func addRight() bool {
	i := 0
	x := false
	for i < 4 {
		j := 3
		for j >= 0 {
			if j != 0 {
				if table[i][j] == table[i][j-1] && table[i][j] != 0 {
					table[i][j] = table[i][j] + table[i][j-1]
					table[i][j-1] = 0
					j -= 2
					x = true
				} else {
					j--
				}
			} else {
				j--
			}
		}
		i++
	}
	return x
}

func printTable() {

	for i := 0; i < 4; i++ {
		fmt.Print("|")
		for j := 0; j < 4; j++ {
			x := " "
			if table[i][j] != 0 {
				x = strconv.Itoa(table[i][j])
			}
			fmt.Printf(" %v |", x)
		}
		fmt.Println()
		fmt.Println()
	}

}
func check() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if table[i][j] == table[i][j+1] {
				return true
			}
			if table[j][i] == table[j+1][i] {
				return true
			}
			if table[3-i][3-j] == table[3-i][3-j-1] {
				return true
			}
			if table[3-j][3-i] == table[3-j-1][3-i] {
				return true
			}
		}
	}
	return false
}
