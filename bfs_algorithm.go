package main

import (
	"fmt"
	"github.com/gammazero/deque"
)

func main() {
	tree := makeTree()
	visited := make(map[string]int)

	var q deque.Deque
	addDequ(tree, "dima", &q)

	for q.Len() != 0 {
		friendName := q.PopFront().(string)

		if _, ok := visited[friendName]; ok != true {
			visited[friendName] = 1
			if friendName == "artem" {
				fmt.Println(friendName, "found!!!")
			} else {
				fmt.Println(friendName, "is not artem")
				addDequ(tree, friendName, &q)
			}
		}
	}

}

func addDequ(tree map[string][]string, friendName string, q *deque.Deque) {
	for _, name := range tree[friendName] {
		q.PushBack(name)
	}
}

func makeTree() map[string][]string {
	tree := make(map[string][]string)
	tree["dima"] = []string{"alex", "andrey", "banan", "misha", "kristina"}
	tree["kristina"] = []string{"yulia"}
	tree["misha"] = []string{"ksusha"}
	tree["banan"] = []string{"artem", "tanya"}
	tree["andrey"] = []string{"artem"}

	return tree
}
