package main

import (
	"fmt"
)

type Graph map[string]map[string]uint

var (
	processed []string
)

func main() {
	graph := makeGraph()
	costs := makeCosts()
	parents := makeParents()

	node := findLowestCostNode(costs, processed)
	for node != "" {
		cost := costs[node]
		neighbors := graph[node]
		for n := range neighbors {
			newCost := cost + neighbors[n]
			if costs[n] > newCost {
				costs[n] = newCost
				parents[n] = node
			}
		}
		processed = append(processed, node)
		node = findLowestCostNode(costs, processed)
	}
	fmt.Printf("the shortest path is %d\n", costs["fin"])
	fmt.Println("the path is:", parents)
}

func makeGraph() Graph {

	tree := Graph{
		"start": {
			"a": 5,
			"b": 2,
		},
		"a": {
			"c": 4,
			"d": 2,
		},
		"b": {
			"a": 8,
			"d": 7,
		},
		"c": {
			"d":   6,
			"fin": 3,
		},
		"d": {
			"fin": 1,
		},
		"fin": {},
	}
	return tree
}

func makeCosts() map[string]uint {
	costs := map[string]uint{
		"a":   5,
		"b":   2,
		"c":   ^uint(0),
		"d":   ^uint(0),
		"fin": ^uint(0),
	}
	return costs
}

func makeParents() map[string]string {
	parents := map[string]string{
		"a":   "start",
		"b":   "start",
		"fin": "",
	}
	return parents
}

func findLowestCostNode(costs map[string]uint, processed []string) string {
	var (
		lowestCostNode string
	)
	lowestCost := ^uint(0)

	for k := range costs {
		cost := costs[k]
		if cost < lowestCost && !sliceContainsString(processed, k) {
			lowestCost = cost
			lowestCostNode = k
		}
	}
	return lowestCostNode
}

func sliceContainsString(slice []string, val string) bool {
	if slice == nil {
		return false
	}
	for _, k := range slice {
		if k == val {
			return true
		}
	}
	return false
}
