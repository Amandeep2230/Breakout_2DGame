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

var paddle_width int16 = 80
var paddle_height int16 = 30
var paddle_x int16 = (window_width - paddle_width) / 2
var paddle_y int16 = window_height - (paddle_height + 10)
var paddle_vel int16 = 4

var ball_size float32 = 10.0
var ball_x int16 = window_width / 2
var ball_y int16 = window_height - (paddle_height + 45)
var ball_vel int16 = 4

func ball_spawn() {
	var ball_bottom int16 = ball_y + int16(ball_size)

	//ball behaviour
	if ball_bottom >= window_height {
		ball_vel = 0
		ball_y = window_height - int16(ball_size)
	}

	ball_y += ball_vel
	ball_vel += 1

	rl.DrawCircle(int32(ball_x), int32(ball_y), ball_size, rl.Red)
}

func paddle_spawn() {
	var paddle_right int16 = paddle_x + paddle_width

	rl.DrawRectangle(int32(paddle_x), int32(paddle_y), int32(paddle_width), int32(paddle_height), rl.Black)

	//paddle behavior
	if rl.IsKeyDown(rl.KeyLeft) {
		paddle_x -= paddle_vel
	}

	if rl.IsKeyDown(rl.KeyRight) {
		paddle_x += paddle_vel
	}

	if paddle_x <= 0 {
		paddle_x = 2
	}

	if paddle_right >= window_width {
		paddle_x = window_width - (paddle_width + 2)
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
