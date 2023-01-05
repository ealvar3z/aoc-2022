package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/ealvar3z/aoc2022/go/lib/aoc"
	"github.com/ealvar3z/aoc2022/go/lib/numbers"
)

type Valve struct {
	idx      int
	name     string
	flowRate int
	inNode   map[int]int
	outNode  map[int]int
}

func (v Valve) String() string {
	return fmt.Sprintf(
		"{ Index: %d, Name: %s, Flow: %d, In: %v, Out: %v}\n", v.idx, v.name, v.flowRate, v.inNode, v.outNode)
}

func main() {
	file, complete := aoc.Setup(2022, 16, false)
	defer complete()

	input := parse(file)
	valves := closedValves(input)
	paths := shortestPath(valves)
	// aoc.PrintAnswer(1, partOne(valves, paths, 30))
	aoc.PrintAnswer(2, partTwo(valves, paths, 26))
}

func parse(fpath string) []*Valve {
	re := regexp.MustCompile("Valve (?P<id>[A-Z]{2}) has flow rate=(?P<flow_rate>\\d+); tunnel[s]? lead[s]? to valve[s]? (?P<connected_valve>[A-Z, ]+)")

	f, _ := os.Open(fpath)
	defer f.Close()

	s := bufio.NewScanner(f)

	vIndex := make(map[string]int)
	var edges [][2]string
	var valves []*Valve
	for s.Scan() {
		match := re.FindStringSubmatch(s.Text())
		name := match[1]
		flow, _ := strconv.ParseInt(match[2], 10, 32)
		nodes := strings.Split(match[3], ", ")

		vIndex[name] = len(valves)
		valves = append(valves, &Valve{
			idx:      len(valves),
			name:     name,
			flowRate: int(flow),
			inNode:   make(map[int]int),
			outNode:  make(map[int]int),
		})

		// connect the edges w/ their respective node
		//		v----------------v
		// e.g AA<-->BB<-->CC<-->DD<-->EE<-->FF<-->GG<-->HH
		//	    ^
		//		^---> II<-->JJ
		for _, node := range nodes {
			// connect the edges of inNodes & outNodes
			edges = append(edges, [2]string{name, node})
		}
	}
	for _, edge := range edges {
		valves[vIndex[edge[0]]].outNode[vIndex[edge[1]]] = 1
		valves[vIndex[edge[1]]].inNode[vIndex[edge[0]]] = 1
	}
	return valves
}

func closedValves(valves []*Valve) []*Valve {
	type indexDist struct {
		index int
		dist  int
	}

	// use a double-linked list to simulate
	// the two-way connection b/wn edges & nodes
	// we refer back to our earlier ascii graph
	// in the previous func
	links := list.New()
	for i := range valves {
		if valves[i].flowRate == 0 && valves[i].name != "AA" {
			// get the incoming edges and
			// evaluate them for possible removal
			incoming := make(map[int]int)
			links.PushBack(indexDist{index: i, dist: 0})
			seen := make(map[int]bool)
			for links.Len() != 0 { // while there are connected edges
				cur := links.Remove(links.Front()).(indexDist) // pop it for examination
				if seen[cur.index] {                           // if we've already seen it; do nothing
					continue
				}
				seen[cur.index] = true // reset

				for in := range valves[cur.index].inNode {
					if valves[in].flowRate == 0 && valves[in].name != "AA" {
						links.PushBack(indexDist{index: in, dist: cur.dist + 1})
					} else {
						if _, ok := incoming[in]; !ok || incoming[in] > cur.dist+1 {
							incoming[in] = cur.dist + 1
						}
					}
				}
			}
			// now we do outgoing edges and
			// evaluate them for possible removal
			outgoing := make(map[int]int)
			links.PushBack(indexDist{index: i, dist: 0})
			seen = make(map[int]bool)
			for links.Len() != 0 { // while there are connected edges
				cur := links.Remove(links.Front()).(indexDist) // pop it for examination
				if seen[cur.index] {                           // if we've already seen it; do nothing
					continue
				}
				seen[cur.index] = true // reset

				for out := range valves[cur.index].outNode {
					if valves[out].flowRate == 0 && valves[out].name != "AA" {
						links.PushBack(indexDist{index: out, dist: cur.dist + 1})
					} else {
						if _, ok := outgoing[out]; !ok || outgoing[out] > cur.dist+1 {
							outgoing[out] = cur.dist + 1
						}
					}
				}
			}
			// merge the incoming w/ the outcoming edges pairwise
			for in, inDist := range incoming {
				for out, outDist := range outgoing {
					if in == out {
						continue
					}
					if _, ok := valves[in].outNode[out]; !ok || valves[in].outNode[out] > inDist+outDist {
						valves[in].outNode[out] = inDist + outDist
					}
					if _, ok := valves[out].inNode[in]; !ok || valves[out].inNode[in] > inDist+outDist {
						valves[out].inNode[in] = inDist + outDist
					}
				}
			}
		}
	}
	// decrease our search space
	// after removing the (0) producing valves
	newIndexes := make(map[int]int)
	for i := range valves {
		if valves[i].flowRate == 0 && valves[i].name != "AA" {
			continue
		}
		newIndexes[i] = len(newIndexes)
	}
	var producingValves []*Valve
	for i := range valves {
		if valves[i].flowRate == 0 && valves[i].name != "AA" {
			continue
		}
		valves[i].idx = newIndexes[i]

		newIncoming := make(map[int]int)
		for in, distance := range valves[i].inNode {
			if _, ok := newIndexes[in]; !ok {
				continue
			}
			newIncoming[newIndexes[in]] = distance
		}
		valves[i].inNode = newIncoming

		newOutgoing := make(map[int]int)
		for out, distance := range valves[i].outNode {
			if _, ok := newIndexes[out]; !ok {
				continue
			}
			newOutgoing[newIndexes[out]] = distance
		}
		valves[i].outNode = newOutgoing

		producingValves = append(producingValves, valves[i])
	}
	return producingValves
}

// shortesPath computes paths b/wn all pairs of nodes (valves in our case)
// since most or all pairs of nodes are linked (i.e. connected by edges)
// Thus, we'll use Floyd-Warshall algorithm: (O(V^3))
// [ https://brilliant.org/wiki/floyd-warshall-algorithm/ ]
func shortestPath(valves []*Valve) [][]int {
	paths := make([][]int, len(valves))
	for i := 0; i < len(valves); i++ {
		paths[i] = make([]int, len(valves))
		for j := 0; j < len(valves); j++ {
			if i == j {
				paths[i][j] = 0
			} else {
				paths[i][j] = 10000 // unreachable value (see linked article)
			}
		}
	}
	for i := 0; i < len(valves); i++ {
		for out, path := range valves[i].outNode {
			paths[i][out] = path
		}
	}
	// here comes Floyd!!!
	for k := 0; k < len(valves); k++ {
		for i := 0; i < len(valves); i++ {
			for j := 0; j < len(valves); j++ {
				if paths[i][k]+paths[k][j] < paths[i][j] {
					paths[i][j] = paths[i][k] + paths[k][j]
				}
			}
		}
	}
	return paths
}

// we need to keep track of the steps
type Steps struct {
	cur    int
	time   int // time wasted opening a valve or moving to a valve
	opened int // bitmask flag to keep track of wether we opened a valve or not
}

// solve partOne
// Work out the steps to release the most pressure in 30 minutes. What is the
// most pressure you can release?
func partOne(valves []*Valve, paths [][]int, time int) int {
	var start int // starting index
	for _, v := range valves {
		if v.name == "AA" {
			start = v.idx
		}
	}
	return backtrack(valves, paths, Steps{
		cur:    start,
		time:   time,
		opened: 0,
	}, make(map[Steps]int))
}

// backtrack is a DFS helper function to exhaust searching all the nodes
func backtrack(valves []*Valve, paths [][]int, cur Steps, memoize map[Steps]int) int {
	if cur.time == 0 {
		return 0
	}

	// if we've already computed it
	// give it to me!
	if answer, alreadyComputed := memoize[cur]; alreadyComputed {
		return answer
	}
	steps, values := calcStepsValues(valves, paths, cur)

	var answer int
	for i, step := range steps {
		answer = numbers.Max(answer, backtrack(valves, paths, step, memoize)+values[i])
	}

	memoize[cur] = answer // stash the answer
	return answer
}

// now we need to keep track of
// both of our steps
type UsTwo struct {
	myCur   int
	myTime  int
	hisCur  int
	hisTime int
	opened  int
}

// solve partTwo
// With you and an elephant working together for 26 minutes, what is the most
// pressure you could release?
func partTwo(valves []*Valve, paths [][]int, time int) int {
	var start int
	for _, v := range valves {
		if v.name == "AA" {
			start = v.idx
		}
	}
	return _backtrack(valves, paths, UsTwo{
		myCur:   start,
		myTime:  time,
		hisCur:  start,
		hisTime: time,
		opened:  0,
	}, make(map[UsTwo]int))
}

// same as partOne, but pairwise
func _backtrack(valves []*Valve, paths [][]int, cur UsTwo, memoize map[UsTwo]int) int {
	curTime := numbers.Max(cur.myTime, cur.hisTime)
	if curTime == 0 {
		return 0
	}

	if answer, alreadyComputed := memoize[cur]; alreadyComputed {
		return answer
	}
	var mySteps []Steps
	var myValues []int
	if curTime == cur.myTime {
		mySteps, myValues = calcStepsValues(valves, paths, Steps{
			cur:    cur.myCur,
			time:   cur.myTime,
			opened: cur.opened,
		})
	}

	var hisSteps []Steps
	var hisValues []int
	if curTime == cur.hisTime {
		hisSteps, hisValues = calcStepsValues(valves, paths, Steps{
			cur:    cur.hisCur,
			time:   cur.hisTime,
			opened: cur.opened,
		})
	}
	var answer int
	if len(mySteps) != 0 && len(hisSteps) != 0 {
		for i, myStep := range mySteps {
			for j, hisStep := range hisSteps {
				if myStep.cur == hisStep.cur {
					// means we're both opening the same valve
					continue
				}
				answer = numbers.Max(answer, _backtrack(valves, paths, UsTwo{
					myCur:   myStep.cur,
					myTime:  myStep.time,
					hisCur:  hisStep.cur,
					hisTime: hisStep.time,
					opened:  myStep.opened | hisStep.opened,
				}, memoize)+myValues[i]+hisValues[j])
			}
		}
	} else if len(mySteps) != 0 {
		for i, myStep := range mySteps {
			answer = numbers.Max(answer, _backtrack(valves, paths, UsTwo{
				myCur:   myStep.cur,
				myTime:  myStep.time,
				hisCur:  cur.hisCur,
				hisTime: cur.hisTime,
				opened:  myStep.opened,
			}, memoize)+myValues[i])
		}
	} else if len(hisSteps) != 0 {
		for i, hisStep := range hisSteps {
			answer = numbers.Max(answer, _backtrack(valves, paths, UsTwo{
				myCur:   cur.myCur,
				myTime:  cur.myTime,
				hisCur:  hisStep.cur,
				hisTime: hisStep.time,
				opened:  hisStep.opened,
			}, memoize)+hisValues[i])
		}
	}
	memoize[cur] = answer
	return answer
}

// All of the valves begin closed. You start at valve AA, but it must be
// damaged or jammed or something: its flow rate is 0, so there's no point in
// opening it. However, you could spend one minute moving to valve BB and
// another minute opening it; doing so would release pressure during the
// remaining 28 minutes at a flow rate of 13, a total eventual pressure release
// of 28 * 13 = 364.

func calcStepsValues(valves []*Valve, paths [][]int, cur Steps) ([]Steps, []int) {
	var steps []Steps
	var values []int

	for i := 0; i < len(valves); i++ {
		if valves[i].flowRate != 0 && cur.opened&(1<<i) == 0 { // check if the i-th bit is set
			// no flow; not opened ... yet!
			timeRemaining := paths[cur.cur][i] + 1
			if cur.time < timeRemaining {
				// need more time
				continue
			}
			steps = append(steps, Steps{
				cur:    i,
				time:   cur.time - timeRemaining,
				opened: cur.opened | (1 << i), // set the i-th bit
			})
			values = append(values, valves[i].flowRate*(cur.time-timeRemaining))
		}
	}
	return steps, values
}
