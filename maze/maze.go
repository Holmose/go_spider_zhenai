package maze

import (
	"fmt"
	"log"
	"os"
)

/*
*
使用广度优先搜索走迷宫
用循环创建二维slice
使用slice来实现队列
用Fscanf读取文件
对Point的抽象
*/
func ReadMaze(filename string) [][]int {
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		log.Panicf("文件打开失败: %v", err)
	}
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	// 分配行
	maze := make([][]int, row)
	// 分配列
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			// 需要将换行符改成LF
			// Fscanf在遇到\n才结束，遇到\r时就会把\r替换成0
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

// 坐标点结构
type Point struct {
	I, J int
}

// 坐标运算 加
func (p Point) add(r Point) Point {
	return Point{p.I + r.I, p.J + r.J}
}

// 判断是否越界
func (p Point) At(grid [][]int) (int, bool) {
	if p.I < 0 || p.I >= len(grid) {
		return 0, false
	}
	if p.J < 0 || p.J >= len(grid[p.I]) {
		return 0, false
	}
	return grid[p.I][p.J], true
}

// 四个位置，上左下右
var dirs = [4]Point{
	{-1, 0}, // 上
	{0, -1}, // 左
	{1, 0},  // 下
	{0, 1},  // 右
}

func Walk(maze [][]int, start, end Point) [][]int {
	// 记录行走记录
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	// 队列
	Q := []Point{start}

END:
	// 队列为空结束
	for len(Q) > 0 {
		// 从队列中取出一个元素
		cur := Q[0]
		Q = Q[1:]

		// 上左下右进行搜索
		for _, dir := range dirs {
			// 获取新发现节点的坐标
			next := cur.add(dir)

			// maze at next is 0		地图的下一个坐标可以走
			// and steps at next is 0	行走历史中没有走过
			// adn next != start		下一节点不等于起点
			val, ok := next.At(maze)
			if !ok || val == 1 {
				continue
			}
			val, ok = next.At(steps)
			if !ok || val != 0 {
				continue
			}
			if next == start {
				continue
			}

			// 填写当前步骤数
			curSetps, _ := cur.At(steps)
			steps[next.I][next.J] = curSetps + 1

			// 走到终点退出
			if next == end {
				break END
			}

			// 将符合的坐标点加入队列中
			Q = append(Q, next)

		}
	}
	return steps
}
