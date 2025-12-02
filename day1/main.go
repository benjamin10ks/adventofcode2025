package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Dial struct {
	vals []int
	pos  int
	size int
}

func NewDial(size int) *Dial {
	vals := make([]int, size)
	for i := 0; i < size; i++ {
		vals[i] = i
	}
	return &Dial{vals: vals, pos: 0, size: size}
}

func (d *Dial) Move(steps int) (crossedZero bool) {
	startPos := d.pos
	endPos := (d.pos + steps) % d.size
	if endPos < 0 {
		endPos += d.size
	}

	if steps > 0 && startPos > endPos {
		crossedZero = true
	} else if steps < 0 && startPos < endPos {
		crossedZero = true
	}

	d.pos = endPos
	return crossedZero
}

func (d *Dial) Value() int {
	return d.vals[d.pos]
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	dial := NewDial(100)
	dial.pos = 50

	password := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		dir := line[0]
		amount, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println("invalid amount:", err)
			continue
		}

		var crossed bool
		switch dir {
		case 'L':
			amount = -amount
			crossed = dial.Move(amount)
		case 'R':
			crossed = dial.Move(amount)
		default:
			fmt.Println("invalid direction:", string(dir))
			continue
		}

		if crossed || dial.Value() == 0 {
			password++
		}
	}

	fmt.Println("Password:", password)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file", err)
	}
}
