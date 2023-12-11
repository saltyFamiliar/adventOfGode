package day5

import (
	"2023/ergo"
	"fmt"
	"slices"
	"strings"
)

type gardenMapping struct {
	dst int
	src int
	rng int
}

type gardenRouter struct {
	routes     []gardenMapping
	minAddress int
	maxAddress int
}

func (gr *gardenRouter) addRoute(dst, src, rng int) {
	newRoute := gardenMapping{dst, src, rng}

	if len(gr.routes) == 0 {
		gr.routes = append(gr.routes, newRoute)
		gr.minAddress, gr.maxAddress = src, src+rng-1
		return
	}

	if src < gr.minAddress {
		gr.minAddress = src
	}

	if src+rng-1 > gr.maxAddress {
		gr.maxAddress = src + rng - 1
	}

	for i, route := range gr.routes {
		if route.src >= newRoute.src {
			gr.routes = slices.Insert(gr.routes, i, newRoute)
			return
		}
	}

	gr.routes = append(gr.routes, newRoute)
}

func (gr *gardenRouter) route(input int) int {
	if input > gr.maxAddress || input < gr.minAddress {
		return input
	}

	//27699552667
	//27795351500
	//for _, r := range gr.routes {
	//	if input < r.src {
	//		continue
	//	}
	//
	//	if input < r.src+r.rng {
	//		return r.dst + (input - r.src)
	//	}
	//}

	//18746385792
	//18344241958
	lowBound := 0
	highBound := len(gr.routes) - 1
	for i := (lowBound + highBound) / 2; lowBound <= highBound; i = (lowBound + highBound) / 2 {
		r := gr.routes[i]
		if input < r.src {
			highBound = i - 1
		} else if input > r.src+r.rng-1 {
			lowBound = i + 1
		} else {
			return r.dst + (input - r.src)
		}
	}

	return input
}

func routeToEnd(seed int, allRouters []gardenRouter) int {
	for _, gr := range allRouters {
		seed = gr.route(seed)
	}
	return seed
}

func Solve1() int {
	scanner := ergo.GetInputScanner("solutions/day5/input.txt")

	scanner.Scan()
	seedsLine := scanner.Text()
	_, seedsLine, _ = strings.Cut(seedsLine, ": ")
	seedsNums := strings.Fields(seedsLine)
	var seedsInts []int
	for _, n := range seedsNums {
		seedsInts = append(seedsInts, ergo.EzIntParse(n))
	}
	scanner.Scan()

	var allRouters []gardenRouter
	latestRouter := gardenRouter{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			allRouters = append(allRouters, latestRouter)
		} else if strings.Contains(line, ":") {
			latestRouter = gardenRouter{}
		} else {
			parts := strings.Fields(line)
			dst, src, rng := ergo.EzIntParse(parts[0]), ergo.EzIntParse(parts[1]), ergo.EzIntParse(parts[2])
			latestRouter.addRoute(dst, src, rng)
		}
	}
	allRouters = append(allRouters, latestRouter)

	lowestLocation := -1
	for _, seed := range seedsInts {
		loc := routeToEnd(seed, allRouters)
		if lowestLocation < 0 || loc < lowestLocation {
			lowestLocation = loc
		}
	}

	return lowestLocation
}

func Solve2() int {
	scanner := ergo.GetInputScanner("solutions/day5/input.txt")

	scanner.Scan()
	seedsLine := scanner.Text()
	_, seedsLine, _ = strings.Cut(seedsLine, ": ")
	seedsNums := strings.Fields(seedsLine)
	var seedsInts []int
	for _, n := range seedsNums {
		seedsInts = append(seedsInts, ergo.EzIntParse(n))
	}
	scanner.Scan()

	var allRouters []gardenRouter
	latestRouter := gardenRouter{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			allRouters = append(allRouters, latestRouter)
		} else if strings.Contains(line, ":") {
			latestRouter = gardenRouter{}
		} else {
			parts := strings.Fields(line)
			dst, src, rng := ergo.EzIntParse(parts[0]), ergo.EzIntParse(parts[1]), ergo.EzIntParse(parts[2])
			latestRouter.addRoute(dst, src, rng)
		}
	}
	allRouters = append(allRouters, latestRouter)

	locCh := make(chan int)
	workersActive := 0
	for i := 0; i < len(seedsInts); i += 2 {
		workersActive += 1
		go func(start, rng int) {
			lowestLocation := -1
			for j := start; j < start+rng; j++ {
				loc := routeToEnd(j, allRouters)
				if lowestLocation < 0 || loc < lowestLocation {
					lowestLocation = loc
				}
			}
			locCh <- lowestLocation
			fmt.Printf("%d to %d done\n", start, rng)
		}(seedsInts[i], seedsInts[i+1])
	}

	lowest := <-locCh
	workersActive -= 1
	for nextLow := range locCh {
		if nextLow < lowest {
			lowest = nextLow
		}
		workersActive -= 1
		if workersActive == 0 {
			break
		}
	}

	return lowest
}
