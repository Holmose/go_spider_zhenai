package test

import (
	maze2 "PRO02/maze"
	"fmt"
	"testing"
)

func TestMaze(t *testing.T) {
	maze := maze2.ReadMaze("../maze/maze.in")
	t.Log("文件载入成功！")

	fmt.Println("----------地图--------------")
	for _, row := range maze {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
	fmt.Println("---------------------------")
	steps := maze2.Walk(maze, maze2.Point{0, 0}, maze2.Point{len(maze) - 1, len(maze[0]) - 1})

	if steps[len(maze)-1][len(maze[0])-1] != 0 {
		t.Log("成功找到终点路径图")
	} else {
		t.Log("没有找到终点路径图")
	}

	fmt.Println("----------路径--------------")
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
	fmt.Println("---------------------------")

}
