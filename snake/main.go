package main

import (
	"fmt"
	"math/rand"
	"os"
	"snake/cutil"
	"time"
)

const (
	WIDTH  = 20
	HIGHTH = 20
)

type Position struct {
	X int
	Y int
}

type Food struct {
	Pos Position
}

var f Food

func RandomFood() {
	f.Pos = Position{X: rand.Intn(WIDTH), Y: rand.Intn(HIGHTH)}
}

func InitFood() {
	rand.Seed(time.Now().UnixNano())
	f.Pos = Position{X: rand.Intn(WIDTH), Y: rand.Intn(HIGHTH)}
}

type Snake struct {
	size int
	dir  byte
	Pos  [WIDTH * HIGHTH]Position
}

var s Snake

func InitSnake() {
	s.size = 2
	s.dir = 'R'
	s.Pos[0] = Position{X: WIDTH / 2, Y: HIGHTH / 2}
	s.Pos[1] = Position{X: WIDTH/2 - 1, Y: HIGHTH / 2}
}
func InitMap() {
	fmt.Fprintln(os.Stderr, `
  #-----------------------------------------#
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  #-----------------------------------------#
`)
}

func ShowUI(x int, y int, ui byte) {
	cutil.GotoPosition(x, y)
	fmt.Fprintf(os.Stderr, "%c", ui)
}

func main() {
	cutil.Clrscr()
	InitSnake()
	InitFood()
	InitMap()
	ShowUI(f.Pos.X, f.Pos.Y, 's')
	fmt.Printf("%#v\n", f)
	for i := 0; i < s.size; i++ {
		if i == 0 {
			ShowUI(s.Pos[i].X, s.Pos[i].Y, '@')
		}
		ShowUI(s.Pos[i].X, s.Pos[i].Y, '*')
	}
}
