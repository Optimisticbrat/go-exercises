package main

import "fmt"

func closestpoint(cordinates [][2]int, point [2]int) [2]int {
	minDist := 100000
	index := 0
	for i, coord := range cordinates {
		//	var origin = [2]int{0, 0}
		var x int
		var y int
		x = coord[0] - point[0]
		y = coord[1] - point[1]
		dist := x*x + y*y
		if dist < minDist {
			minDist = dist
			index = i
		}
		fmt.Println(i, coord, dist)
	}
	return cordinates[index]
}

func main() {
	var coordinates = [][2]int{{5, -2}, {4, 7}, {0, 1}, {-2, 2}, {-1, 6}, {6, 3}}
	var point = [2]int{3, 5}
	coords := closestpoint(coordinates, point)
	fmt.Println("Nearest coordinate is ", coords)
}
