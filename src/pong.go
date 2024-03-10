package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type pos struct {
	x float32
	y float32
}

type ball struct {
	pos
	radius int
	vx     float32
	vy     float32
	color
}

func (ball *ball) draw(pixels []byte) {
	for y := -ball.radius; y < ball.radius; y++ {
		for x := -ball.radius; x < ball.radius; x++ {
			if x*x+y*y < ball.radius*ball.radius {
				setPixel(int(ball.x)+x, int(ball.y)+y, ball.color, pixels)
			}
		}
	}
}

func (ball *ball) Update(leftPaddle *paddle, rightPaddle *paddle) {
	ball.x += ball.vx
	ball.y += ball.vy
	if int(ball.y)-ball.radius < 0 || int(ball.y)+ball.radius > winHeight {
		ball.vy = -ball.vy
	}
	if int(ball.x)-ball.radius < 0 || int(ball.x)+ball.radius > winWidth {
		ball.x = 300
		ball.y = 300
	}

	if int(ball.x)-ball.radius < int(leftPaddle.x)+leftPaddle.w/2 &&
		int(ball.y) > int(leftPaddle.y)-leftPaddle.h/2 && int(ball.y) < int(leftPaddle.y)+leftPaddle.h/2 {
		ball.vx = -ball.vx
	} else if int(ball.x)+ball.radius > int(rightPaddle.x)-rightPaddle.w/2 &&
		int(ball.y) > int(rightPaddle.y)-rightPaddle.h/2 && int(ball.y) < int(rightPaddle.y)+rightPaddle.h/2 {
		ball.vx = -ball.vx
	}

}

type paddle struct {
	pos
	color
	w int
	h int
}

func (paddle *paddle) draw(pixels []byte) {
	startX := int(paddle.x) - paddle.w/2
	startY := int(paddle.y) - paddle.h/2

	for y := 0; y < paddle.h; y++ {
		for x := 0; x < paddle.w; x++ {
			setPixel(startX+x, startY+y, paddle.color, pixels)
		}
	}
}

func (paddle *paddle) Update(keyState []uint8) {
	if keyState[sdl.SCANCODE_UP] != 0 {
		paddle.y -= 5
	}
	if keyState[sdl.SCANCODE_DOWN] != 0 {
		paddle.y += 5
	}
}

func (paddle *paddle) aiUpdate(ball *ball) {
	paddle.y = ball.y
}

func clear(pixels []byte) {
	for i := range pixels {
		pixels[i] = 0
	}
}

// func findEdgeOfBall(param float32) float32 {
// 	return
// }

func initializePong() (paddle, ball, paddle) {
	player1 := paddle{pos{50, 100}, color{255, 255, 255}, 20, 100}
	ball := ball{pos{300, 300}, 20, 2, 2, color{255, 255, 255}}
	player2 := paddle{pos{750, 100}, color{255, 255, 255}, 20, 100}

	return player1, ball, player2

}
func startPong(player1 *paddle, ball *ball, player2 *paddle, pixels []byte, keyState []uint8) {
	// fmt.Println(player1)

	clear(pixels)

	player1.Update(keyState)
	player1.draw(pixels)
	player2.aiUpdate(ball)
	player2.draw(pixels)
	ball.Update(player1, player2)
	ball.draw(pixels)

}
