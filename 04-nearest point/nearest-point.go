package main

import "fmt"

func nearcoords(cordinates [][2]int) [2]int {
	minDist := 1000000
	index := 0
	for i, coord := range cordinates {
		//	var origin = [2]int{0, 0}
		dist := coord[0]*coord[0] + coord[1]*coord[1]
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
	coords := nearcoords(coordinates)
	fmt.Println("Nearest coordinate is ", coords)
}
