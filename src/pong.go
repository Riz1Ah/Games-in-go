package main

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

func startPong(pixels []byte) {
	player1 := paddle{pos{100, 100}, color{255, 255, 255}, 20, 100}
	ball := ball{pos{300, 300}, 20, 0, 0, color{255, 255, 255}}

	player1.draw(pixels)
	ball.draw(pixels)

}
