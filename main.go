package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var brick_width int16 = 60
var brick_height int16 = 20
var brick_offset_x int16 = 2
var brick_offset_y int16 = 2
var window_width int16 = 622
var window_height int16 = 700
var rows int = 6
var cols int = 10
var padding int = 2
var gameOver bool = false
var gameCompleted bool = false

type Paddle struct {
	width  int16
	height int16
	x      int16
	y      int16
	vel    int16
}

var paddle = Paddle{
	width:  120,
	height: 20,
	x:      (window_width - 80) / 2,
	y:      (window_height - 40),
	vel:    9,
}

type Ball struct {
	size  float32
	x     int16
	y     int16
	vel_x int16
	vel_y int16
}

var ball = Ball{
	size:  10.0,
	x:     window_width / 2,
	y:     window_height - 75,
	vel_y: 6,
	vel_x: 0,
}

type Brick struct {
	pos    rl.Vector2
	width  float32
	height float32
	hit    bool
}

var bricks [][]Brick = create_bricks(2.0, 100.0, 2.0)
var score int
var max_score int = len(bricks) * len(bricks[1])

func ball_spawn() {
	var ball_ver int16 = ball.y + int16(ball.size)
	var ball_hor int16 = ball.x + int16(ball.size)

	//ball behaviour
	if ball_ver <= int16(ball.size) {
		ball.vel_y = 6
	}

	if ball_hor <= int16(ball.size) {
		ball.vel_x = 6
	} else if ball_hor >= window_width {
		ball.vel_x = -6
	}

	//end game if ball misses the paddle
	if ball_ver >= window_height {
		gameOver = true
	}

	//determine point of contact with paddle
	var hitPos float32 = float32(ball.x) - float32(paddle.x)
	//convert the contact into range between -0.5 and 0.5
	var relHit float32 = (hitPos / float32(paddle.width)) - 0.5

	var ball_position = rl.Vector2{float32(ball.x), float32(ball.y)}
	var paddle_position = rl.Rectangle{float32(paddle.x), float32(paddle.y), float32(paddle.width), float32(paddle.height)}

	if rl.CheckCollisionCircleRec(ball_position, ball.size, paddle_position) {

		if (-0.5 <= relHit) && (relHit < -0.3) {
			ball.vel_y = -6
			ball.vel_x = -6
		} else if (-0.3 <= relHit) && (relHit < -0.1) {
			ball.vel_y = -6
			ball.vel_x = -4
		} else if -0.1 <= relHit && relHit < 0.1 {
			ball.vel_y = -6
			ball.vel_x = 0
		} else if 0.1 <= relHit && relHit < 0.3 {
			ball.vel_y = -6
			ball.vel_x = 4
		} else if 0.3 <= relHit && relHit <= 0.5 {
			ball.vel_y = -6
			ball.vel_x = 6
		}

	}

	ball.y += ball.vel_y
	ball.x += ball.vel_x

	rl.DrawCircle(int32(ball.x), int32(ball.y), ball.size, rl.Red)
}

func paddle_spawn() {
	var paddle_right int16 = paddle.x + paddle.width

	rl.DrawRectangle(int32(paddle.x), int32(paddle.y), int32(paddle.width), int32(paddle.height), rl.Black)

	//paddle behavior
	if rl.IsKeyDown(rl.KeyLeft) {
		paddle.x -= paddle.vel
	}

	if rl.IsKeyDown(rl.KeyRight) {
		paddle.x += paddle.vel
	}

	if paddle.x <= 0 {
		paddle.x = 2
	}

	if paddle_right >= window_width {
		paddle.x = window_width - (paddle.width + 2)
	}

}

func create_bricks(startX, startY, padding float32) [][]Brick {

	grid := make([][]Brick, rows)
	for i := 0; i < rows; i++ {
		row := make([]Brick, cols)
		for j := 0; j < cols; j++ {
			x := 2 + (int16(j) * (brick_width + int16(padding)))
			y := 100 + (int16(i) * (brick_height + int16(padding)))

			row[j] = Brick{
				pos:    rl.NewVector2(float32(x), float32(y)),
				width:  float32(brick_width),
				height: float32(brick_height),
				hit:    true,
			}
		}
		grid[i] = row
	}
	return grid

}

func brick_spawn() {

	for _, row := range bricks {
		for _, brick := range row {
			if brick.hit {
				rl.DrawRectangle(
					int32(brick.pos.X),
					int32(brick.pos.Y),
					int32(brick.width),
					int32(brick.height),
					rl.Black)
			}
		}
	}

	checkBrickCollision()

}

func checkBrickCollision() {
	//ball collision behavior with bricks
	var ball_position = rl.Vector2{float32(ball.x), float32(ball.y)}

	for i := range bricks {
		for j := range bricks[i] {
			brick := &bricks[i][j]
			if brick.hit {
				brick_position := rl.NewRectangle(brick.pos.X, brick.pos.Y, brick.width, brick.height)
				if rl.CheckCollisionCircleRec(ball_position, ball.size, brick_position) {
					brick.hit = false
					//determine point of contact with paddle
					var hitPos float32 = float32(ball.x) - float32(brick.pos.X)
					//convert the contact into range between -0.5 and 0.5
					var relHit float32 = (hitPos / float32(brick.width)) - 0.5

					//ball behavior after collision
					if (-0.5 <= relHit) && (relHit < -0.1) {
						ball.vel_y = 6
						ball.vel_x = -6
					} else if (-0.1 <= relHit) && (relHit < 0.2) {
						ball.vel_y = 6
						ball.vel_x = 1
					} else if 0.2 <= relHit && relHit <= 0.5 {
						ball.vel_y = 6
						ball.vel_x = 6
					}
					score += 1

					if score == max_score {
						gameCompleted = true
					} else {
						continue
					}
				}
			}
		}
	}
}

func set_scene() {

	rl.DrawText(fmt.Sprintf("Score: %d", score), 2, 2, 20, rl.Black)
	brick_spawn()
	paddle_spawn()
	ball_spawn()

}

func main() {
	rl.InitWindow(int32(window_width), int32(window_height), "Breakout - 2D Game using Go")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	started := false

	for !rl.WindowShouldClose() {

		if rl.IsKeyPressed(rl.KeySpace) {
			started = true
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		if !started {
			rl.DrawText("Please press space key to start", int32(window_width/4), int32(window_height/2), 20, rl.Gray)
		} else {
			if !gameOver {
				if !gameCompleted {
					set_scene()
				} else {
					rl.DrawText("Game Completed!", int32(window_width/4)+20, int32(window_height/2), 40, rl.Gray)
					rl.DrawText("Thank you for playing!", int32(window_width/3), int32(window_height/2)+50, 20, rl.Gray)
				}
			} else {
				rl.DrawText("Game Over!", int32(window_width/3), int32(window_height/2), 40, rl.Gray)
				rl.DrawText(fmt.Sprintf("Final Score: %d", score), int32(window_width/3)+40, int32(window_height/2)+50, 20, rl.Gray)
			}
		}
		rl.EndDrawing()
	}

}
