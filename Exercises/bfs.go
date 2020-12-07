package main

import (
	"fmt"
)

const size int = 10 + 5

var graph [size][]int

func bfs(start int) [size]int {
	var distance [size]int
	for i := 0; i < size; i++ {
		distance[i] = -1
	}
	distance[start] = 0
	queue := make([]int, 0)
	queue = append(queue, start)
	for len(queue) > 0 {
		frontIndex := 0
		front := queue[frontIndex]
		queue = queue[frontIndex+1:]
		for _, next := range graph[front] {
			if distance[next] == -1 || distance[next] > distance[front]+1 {
				distance[next] = distance[front] + 1
				queue = append(queue, next)
			}
		}
	}
	return distance
}

func main() {
	fmt.Println("breadth-first search!")

	graph[1] = append(graph[1], 2)
	graph[2] = append(graph[2], 1)

	graph[2] = append(graph[2], 3)
	graph[3] = append(graph[3], 2)

	// printing the adjacency list that we made using slice (graph[])
	for i := 1; i <= 5; i++ {
		if len(graph[i]) > 0 {
			for _, val := range graph[i] {
				fmt.Print(val, " ")
			}
			fmt.Println()
		}
	}
	// normal bfs, source is 1 here, finds the shortest path from 1 to all the other nodes
	distance := bfs(1)
	fmt.Println(distance)
}
