package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"

	"github.com/ealvar3z/aoc2022/go/lib/aoc"
	"github.com/ealvar3z/aoc2022/go/lib/numbers"
)

type position struct {
	x int
	y int
}

type Sensor struct {
	pos    position
	beacon position
}

type Range struct {
	start int
	end   int
}

func main() {
	file, complete := aoc.Setup(2022, 15, false)
	defer complete()

	input := parse(file)
	aoc.PrintAnswer(1, partOne(input, 2000000))
	aoc.PrintAnswer(2, partTwo(input, 4000000))
}

func parse(fpath string) []Sensor {
	re := regexp.MustCompile("-?[0-9]+")

	f, _ := os.Open(fpath)
	defer f.Close()

	s := bufio.NewScanner(f)

	var sensors []Sensor
	for s.Scan() {
		match := re.FindAllString(s.Text(), -1)
		xSensor, _ := strconv.ParseInt(match[0], 10, 32)
		ySensor, _ := strconv.ParseInt(match[1], 10, 32)
		xBeacon, _ := strconv.ParseInt(match[2], 10, 32)
		yBeacon, _ := strconv.ParseInt(match[3], 10, 32)

		sensors = append(sensors, Sensor{
			pos: position{
				x: int(xSensor),
				y: int(ySensor),
			},
			beacon: position{
				x: int(xBeacon),
				y: int(yBeacon),
			},
		})
	}
	return sensors
}

func partOne(sensors []Sensor, yAxis int) int {
	beaconLoc := make(map[position]bool)

	// checking minx, maxx
	minx, maxx := math.MaxInt, math.MinInt
	for _, s := range sensors {
		beaconLoc[s.beacon] = true

		beaconDist := manhattan(s.pos, s.beacon)
		yDist := numbers.Abs(yAxis - s.pos.y)
		if yDist > beaconDist {
			// keep going
			continue
		}
		xDist := beaconDist - yDist

		minx, maxx = numbers.Min(minx, s.pos.x-xDist), numbers.Max(maxx, s.pos.x+xDist)
	}

	// check each pos b/wn [minx, maxx]
	// if it's near a sensor && that's the corresponding beacon
	var answer int
	for x := minx; x <= maxx; x++ {
		pos := position{x: x, y: yAxis}
		if beaconLoc[pos] {
			continue
		}

		for _, s := range sensors {
			if manhattan(pos, s.pos) <= manhattan(s.pos, s.beacon) {
				// count
				answer++
				break
			}

		}

	}
	return answer
}

// the searchArea is the x and y coordinates:
// each no lower than 0 and no larger than 4000000
func partTwo(sensors []Sensor, searchArea int) int {
	ranges := make([][]Range, searchArea+1) // one more to be inclusive in our search
	for y := 0; y <= searchArea; y++ {
		for _, s := range sensors {
			beaconDist := manhattan(s.pos, s.beacon)
			yDist := numbers.Abs(y - s.pos.y)
			if yDist > beaconDist {
				// keep going
				continue
			}
			xDist := beaconDist - yDist

			ranges[y] = append(ranges[y], Range{
				start: s.pos.x - xDist,
				end:   s.pos.x + xDist,
			})
		}
		if len(ranges[y]) == 0 {
			panic(fmt.Sprintf("there's no range at Y = %d", y))
		}
		// sort the ranges by their starting point (i.e. x)
		sort.Slice(ranges[y], func(i, j int) bool {
			return ranges[y][i].start < ranges[y][j].start
		})
		// merge them
		var lstIndex int
		for i := 1; i < len(ranges[y]); i++ {
			lst := ranges[y][lstIndex]
			cur := ranges[y][i]
			if lst.end+1 >= cur.start { // we found an overlap
				// merge
				ranges[y][lstIndex] = Range{
					start: lst.start,
					end:   numbers.Max(lst.end, cur.end),
				}
			} else {
				// no overlap, move the last index's pointer
				lstIndex++
				ranges[y][lstIndex] = cur
			}

		}
		ranges[y] = ranges[y][:lstIndex+1]

		// With this reduced search area, there is only a single
		// position that could have a beacon
		if len(ranges[y]) == 2 {
			// To isolate the distress beacon's signal, you need to
			// determine its tuning frequency, which can be found by
			// multiplying its x coordinate by 4000000 and then adding
			// its y coordinate
			return (ranges[y][1].start-1)*searchArea + y
		}

	}
	// we got zilch
	return 0
}

func manhattan(fst, snd position) int {
	return numbers.Abs(snd.x-fst.x) + numbers.Abs(snd.y-fst.y)
}
