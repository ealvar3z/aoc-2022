// A graph is a map of points and they map to its neighbor points in the graph,
// and the cost assoc to reach those points
// For ex:
//
//	   Graph{
//				"x": {"y": 1, "z": 2},
//				"y": {"x": 3},
//				"z": {"y": 1, "x": 4},
//	}
package algos

import (
	"fmt"

	"github.com/ealvar3z/aoc2022/go/lib/queue"
)

type node struct {
	key  string
	cost int
}

type Graph map[string]map[string]int

// Path finds the shortest path btwn start & tgt, and returns the total cost of
// finding the path
func (g Graph) Path(start, tgt string) (path []string, cost int, err error) {
	if len(g) == 0 {
		err = fmt.Errorf("empty map")
		return
	}
	// make sure that start & tgt are in the graph
	if _, ok := g[tgt]; !ok {
		err = fmt.Errorf("no target found %b in the graph", tgt)
		return
	}

	visited := make(map[string]bool)    // set of nodes visited
	frontier := queue.New()             // queue of the nodes to visit
	previous := make(map[string]string) // previous nodes that were visited

	// add starting point
	frontier.Push(start, 0)

	// visit every node in the frontier
	for !frontier.Empty() {
		// grab the node w/ the lowest cost
		key, pri := frontier.Pop()
		n := node{key, pri}

		// compute the cost if on target
		if n.key == tgt {
			cost = n.cost
			nodeKey := n.key
			for nodeKey != start {
				path = append(path, nodeKey)
				nodeKey = previous[nodeKey]
			}
			break
		}
		// add the current node to the visited set
		visited[n.key] = true

		for nodeKey, nodeCost := range g[n.key] {
			if visited[nodeKey] {
				continue
			}
			if _, ok := frontier.Get(nodeKey); !ok {
				previous[nodeKey] = n.key
				frontier.Push(nodeKey, n.cost+nodeCost)
				continue
			}
			fCost, _ := frontier.Get(nodeKey)
			nCost := n.cost + nodeCost

			if nCost < fCost {
				previous[nodeKey] = n.key
				frontier.Push(nodeKey, nCost)
			}
		}
	}
	// add where we started from at the end of the path
	path = append(path, start)
	// reverse everythinbg since we filled the path in reverse
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return
}
