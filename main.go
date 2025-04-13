package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var brick_width int16 = 60
var brick_height int16 = 20
var brick_offset_x int16 = 2
var brick_offset_y int16 = 2
var window_width int16 = 622
var window_height int16 = 800

type Paddle struct {
	width  int16
	height int16
	x      int16
	y      int16
	vel    int16
}

var paddle = Paddle{
	width:  80,
	height: 30,
	x:      (window_width - 80) / 2,
	y:      (window_height - 40),
	vel:    4,
}

type Ball struct {
	size float32
	x    int16
	y    int16
	vel  int16
}

var ball = Ball{
	size: 10.0,
	x:    window_width / 2,
	y:    window_height - 75,
	vel:  4,
}

func ball_spawn() {
	var ball_ver int16 = ball.y + int16(ball.size)
	//var ball_hor int16 = ball.x + int16(ball.size)

	//ball behaviour
	if ball_ver <= 0 {
		ball.vel = 4
	}

	var ball_position = rl.Vector2{float32(ball.x), float32(ball.y)}
	var paddle_position = rl.Rectangle{float32(paddle.x), float32(paddle.y), float32(paddle.width), float32(paddle.height)}

	if rl.CheckCollisionCircleRec(ball_position, ball.size, paddle_position) {
		ball.vel = -4
	}

	ball.y += ball.vel
	//ball.vel += 1

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

func set_scene() {
	rl.ClearBackground(rl.RayWhite)

	for pos_x := 2; pos_x < int(window_width-2); pos_x += int(brick_width + brick_offset_x) {
		for pos_y := 100; pos_y < int((window_height-2)/2); pos_y += int(brick_height + brick_offset_y) {
			rl.DrawRectangle(int32(pos_x), int32(pos_y), int32(brick_width), int32(brick_height), rl.DarkGray)
		}
	}

	paddle_spawn()
	ball_spawn()
}

func main() {
	rl.InitWindow(int32(window_width), int32(window_height), "Breakout - 2D Game using Go")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		set_scene()
		rl.EndDrawing()
	}

}
