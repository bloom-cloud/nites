package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

type GameObject struct {
	row,
	col,
	width,
	height,
	velRow,
	velCol,
	symbol int
}

var SPRITE_BLOCK = 0x2588 //'█'
var SPRITE_BALL = 0x25CF

var SCREEN tcell.Screen
var player1 *GameObject
var player2 *GameObject
var ball *GameObject
var gameObjects = []*GameObject{}
var debugLog string = "test"

const PADDLE_HEIGHT = 4
const INIT_VELOCITY_ROW = 1
const INIT_VELOCITY_COL = 2

func main() {
	initScreen()
	initGameObjects()
	chanInput := initUserInput()

	for {
		SCREEN.Clear()

		UpdateState()

		// wait for messages
		var key string
		select {
		case key = <-chanInput:
		default:
			key = ""
		}

		// Handle user input
		_, screenHeight := SCREEN.Size()
		if key == "Rune[q]" {
			SCREEN.Fini()
			os.Exit(0)
		} else if key == "Rune[w]" && player1.row > 0 {
			player1.row--
		} else if key == "Rune[s]" && (player1.row+player1.height) < screenHeight {
			player1.row++
		} else if key == "Up" && player2.row > 0 {
			player2.row--
		} else if key == "Down" && (player2.row+player2.height) < screenHeight {
			player2.row++
		}
		debugLog = fmt.Sprintf("%v", player1.row)

		// Draw state
		for _, item := range gameObjects {
			renderRect(item.row, item.col, item.width, item.height, rune(item.symbol))
		}
		// renderRect(ball.row, ball.col, ball.width, ball.height, rune(SPRITE_BALL))
		// renderRect(player1.row, player1.col, player1.width, player1.height, rune(SPRITE_BLOCK))
		// renderRect(player2.row, player2.col, player2.width, player2.height, rune(SPRITE_BLOCK))
		renderString(2, 3, debugLog)

		time.Sleep(75 * time.Millisecond)
		SCREEN.Show()
	}

}

func UpdateState() {
	for i := range gameObjects {
		obj := gameObjects[i]
		obj.row += obj.velRow
		obj.col += obj.velCol

		if CollidesWithWall(obj) {
			obj.velRow = -obj.velRow
		}

		if CollidesWithPaddle(ball, player1) || CollidesWithPaddle(ball, player2) {
			obj.velCol = -obj.velCol
		}
	}
}

func CollidesWithWall(obj *GameObject) bool {
	_, screenHeight := SCREEN.Size()

	// top or bottom collision
	if obj.row <= 0 || obj.row+obj.height >= screenHeight {
		return true
	}

	return false
}

func CollidesWithPaddle(ball, paddle *GameObject) bool {
	nextRow := ball.row + ball.velRow
	nextCol := ball.col + ball.velCol

	overlapsVertically := nextRow < paddle.row+paddle.height && nextRow+ball.height > paddle.row

	// left paddle collision
	if ball.velCol < 0 && nextCol <= paddle.col+paddle.width && nextCol+ball.width > paddle.col {
		return overlapsVertically
	}

	// right paddle collision
	if ball.velCol > 0 && nextCol+ball.width >= paddle.col && nextCol < paddle.col+paddle.width {
		return overlapsVertically
	}

	if IsGameOver() {
		winner := GetWinner()
		renderString(0, 0, "Game Over")
		renderString(0, 0, fmt.Sprintf("%s wins", winner))
	}

	return false
}

func IsGameOver() bool {
	return GetWinner() != ""
}

func GetWinner() string {
	screenWidth, _ := SCREEN.Size()
	if ball.col < 0 {
		return "Player 2"
	} else if ball.col >= screenWidth {
		return "Player 1"
	} else {
		return ""
	}
}

// background Process and send data using channel
func initUserInput() chan string {
	// make initializes new variables - same thing
	// _ = make([]int, 0)
	// _ := []int()

	// but will use make for channels
	inputChan := make(chan string)

	// go func() {
	// 	// anonimus function and invoke
	// }()
	go func() {
		for {
			switch ev := SCREEN.PollEvent().(type) {
			case *tcell.EventKey:
				inputChan <- ev.Name()
			case *tcell.EventResize:
				SCREEN.Sync()
			}
		}
	}()
	return inputChan
}

func renderString(row, col int, line string) {
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite)
	for _, c := range line {
		SCREEN.SetContent(col, row, c, nil, style)
		col++ // move right for next character
	}
}

func renderRect(row, col, width, height int, ch rune) {
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite)
	for dy := 0; dy < height; dy++ {
		for dx := 0; dx < width; dx++ {
			SCREEN.SetContent(col+dx, row+dy, ch, nil, style)
		}
	}
}

func initGameObjects() {
	w, h := SCREEN.Size()
	height := PADDLE_HEIGHT
	width := 1
	player1 = &GameObject{
		row:    h/2 - height/2,
		col:    0,
		width:  width,
		height: height,
		symbol: SPRITE_BLOCK,
	}
	player2 = &GameObject{
		row:    h/2 - height/2,
		col:    w - width,
		width:  width,
		height: height,
		symbol: SPRITE_BLOCK,
	}
	ball = &GameObject{
		row:    h / 2,
		col:    w / 2,
		width:  1,
		height: 1,
		velRow: INIT_VELOCITY_ROW,
		velCol: INIT_VELOCITY_COL,
		symbol: SPRITE_BALL,
	}
	gameObjects = []*GameObject{
		player1,
		player2,
		ball,
	}
}

func initScreen() {
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := screen.Init(); err != nil {
		panic(err)
	}
	SCREEN = screen
}

/*
# install
go get github.com/gdamore/tcell/v2
*/
